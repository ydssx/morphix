import os

from langchain import hub
from langchain.agents import (
    AgentExecutor,
    create_structured_chat_agent,
    create_tool_calling_agent,
    create_react_agent,
    tool,
)
from langchain.chains import RetrievalQA
from langchain.chains.combine_documents import create_stuff_documents_chain
from langchain.chains.history_aware_retriever import create_history_aware_retriever
from langchain.chains.retrieval import create_retrieval_chain
from langchain.chains.summarize import load_summarize_chain
from langchain.memory import ChatMessageHistory, FileChatMessageHistory
from langchain.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain.tools.retriever import create_retriever_tool
from langchain.vectorstores.chroma import Chroma
from langchain_community.chat_models import ChatOllama
from langchain_community.document_loaders import WebBaseLoader
from langchain_community.embeddings import OllamaEmbeddings
from langchain_community.tools.tavily_search import TavilySearchResults
from langchain_core.messages import AIMessage, HumanMessage
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables.history import RunnableWithMessageHistory
from langchain_text_splitters import RecursiveCharacterTextSplitter

prompt = ChatPromptTemplate.from_messages(
    [
        (
            "system",
            "You are a helpful assistant. Answer the user's questions in Simplified Chinese based on the below context:\n\n{context}",
        ),
        MessagesPlaceholder(variable_name="chat_history"),
        ("user", "{input}"),
        MessagesPlaceholder(variable_name="agent_scratchpad", optional=True),
    ]
)


class ChatBot:
    def __init__(self, model="llama3:instruct"):
        self.model = model
        self.chatbot = ChatOllama(model=model)
        self.chat_history = FileChatMessageHistory("chat_history.json")
        self.retriever = None
        self.chain = prompt | self.chatbot
        self.agent = None

    def chat(self, message: str, stream: bool = True):
        """
        Generate an answer based on the given message.

        Parameters
        ----------
        message : str
            The message to be used for generating an answer
        stream : bool, optional
            Whether to generate an answer in a streaming fashion,
            by default True

        Yields
        -------
        str
            The generated answer, either in a single string or
            in a series of strings
        """
        # docs = self.retriever.invoke(message)
        input_ = {
            "chat_history": self.chat_history.messages,
            "context": [],
            "input": message,
        }
        # If the chain is a RunnableWithMessageHistory, then we don't
        # need to pass the chat history to the input. The RunnableWithMessageHistory
        # will handle the history internally.
        if isinstance(self.chain, RunnableWithMessageHistory):
            input_.pop("chat_history")

        if stream:
            # Generate an answer in a streaming fashion
            # using the Chain's stream method
            response = self.chain.stream(
                input_,
                {"configurable": {"session_id": "bar"}},
            )
            for r in response:
                # Yield each response as it comes
                yield r.content
        else:
            # Generate an answer in a single string
            # using the Chain's invoke method
            response = self.chain.invoke(
                input_,
                {"configurable": {"session_id": "bar"}},
            )
            # Extract the answer from the Chain's response
            output = response.content or response["answer"]
            # Yield the single answer
            yield output

    def create_retriever(self, url: str):
        """
        Create a retriever based on the documents from the given URL

        Parameters
        ----------
        url : str
            The URL of the documents to use for training the retriever

        Returns
        -------
        Retriever
            The created retriever
        """
        loader = WebBaseLoader(url)
        docs = loader.load()

        # Split the documents into smaller chunks for training
        text_splitter = RecursiveCharacterTextSplitter(chunk_size=500, chunk_overlap=0)
        all_splits = text_splitter.split_documents(docs)

        # Create a Chroma vector store from the split documents
        vectorstore = Chroma.from_documents(
            documents=all_splits,
            # Use the Ollama embeddings model to generate embeddings for the documents
            embedding=OllamaEmbeddings(model=self.model),
        )

        # Create a retriever from the vector store
        retriever = vectorstore.as_retriever(k=4)
        self.retriever = retriever

        return retriever

    def agent_chat(self, message: str, with_history: bool = True):
        # agent_executor = self.create_agent(with_history)

        response = self.agent.invoke(
            {
                # "chat_history": self.chat_history,
                "input": message,
            },
            {"configurable": {"session_id": "unused"}},
        )
        self.chat_history.add_user_message(message)

        # self.chat_history.append(AIMessage(content=response["action_input"]))
        return response

    def create_agent(self, with_history: bool = True):
        """
        Create the structured chat agent to be used for generating answers.

        Returns:
            AgentExecutor: An instance of the AgentExecutor class which is responsible for executing the agent
        """
        prompt = hub.pull("hwchase17/react")
        # Create a tool to search for information about LangSmith
        retriever_tool = create_retriever_tool(
            self.retriever,
            "langsmith_search",
            "Search for information about LangSmith. For any questions about LangSmith, you must use this tool!",
        )

        search = TavilySearchResults(max_results=1)
        tools = [
            # retriever_tool,
            search,
        ]

        @tool
        def magic_function(input: int) -> int:
            """Applies a magic function to an input."""
            return input + 2

        # tools = [magic_function]

        # Create the structured chat agent
        agent = create_react_agent(self.chatbot, tools, prompt)

        # Create an agent executor to execute the agent
        agent_executor = AgentExecutor(
            agent=agent, tools=tools, verbose=True, handle_parsing_errors=True
        )
        if with_history:
            agent_executor = RunnableWithMessageHistory(
                agent_executor,
                lambda session_id: self.chat_history,
                input_messages_key="input",
                output_messages_key="answer",
                history_messages_key="chat_history",
            )
        self.agent = agent_executor
        return agent_executor

    def manage_history(
        self, message: str, max_history: int = 100, sumary_length: int = 100
    ) -> None:
        """
        Manage the length of the chat history.

        If the length of the chat history is greater than the maximum history,
        the oldest messages will be removed from the history.

        Args:
            message (str):
                The message to be appended to the chat history.
            max_history (int, optional):
                The maximum number of messages to keep in the history. Defaults to 100.
            sumary_length (int, optional):
                The maximum length of the chat history summary. Defaults to 100.
        """
        if len(self.chat_history.messages) > max_history:
            # Remove the oldest messages from the chat history
            self.chat_history.messages = self.chat_history.messages[-max_history:]

    def create_chain(self, retriever=None, with_history=True):
        """
        Create a chain of functions to be executed on the user input. By default,
        the chain includes the following functions:

            1. The prompt function which generates a question based on the user input
            2. The chatbot function which generates an answer based on the user input

        If a retriever is provided, the chain is modified as follows:

            1. The prompt function is still executed
            2. The retriever is executed which retrieves documents from the web
            3. The chatbot is executed on the user input and the retrieved documents

        If with_history is True, the retriever is wrapped with a history-aware
        retriever which takes into account the user's chat history when retrieving
        documents.

        Parameters
        ----------
        retriever : function, optional
            A function that retrieves documents from the web, by default None
        with_history : bool, optional
            Whether to wrap the retriever with a history-aware retriever, by default True

        Returns
        -------
        Chain
            A chain of functions that can be executed in series
        """
        # Create a basic chain with the prompt and chatbot functions
        self.chain = prompt | self.chatbot

        # If a retriever function is provided, modify the chain
        if retriever is not None:

            # Create a chain that retrieves documents from the web
            document_chain = create_stuff_documents_chain(self.chatbot, prompt)

            # If with_history is True, wrap the retriever with a history-aware retriever
            if with_history:
                retriever = create_history_aware_retriever(
                    self.chatbot, retriever, prompt
                )

            # Create a chain that executes the retriever and the chatbot
            self.chain = create_retrieval_chain(retriever, document_chain)
        else:
            if with_history:
                self.chain = RunnableWithMessageHistory(
                    self.chain,
                    self.get_session_history,
                    input_messages_key="input",
                    history_messages_key="chat_history",
                    output_messages_key="output",
                )

        return self.chain

    def get_prompt(self, usecase: str = None):
        if usecase is None or usecase == "chat":
            prompt = hub.pull("langsmith/chat-prompt")
        elif usecase == "code":
            prompt = hub.pull("langsmith/code-prompt")
        elif usecase == "rag":
            prompt = hub.pull("rlm/rag-prompt")
        elif usecase == "summarize":
            prompt = hub.pull("langsmith/summarize-prompt")
        elif usecase == "react":
            prompt = hub.pull("hwchase17/react")
        else:
            raise ValueError(f"Unknown usecase: {usecase}")
        return prompt

    def get_session_history(self, session_id):
        return self.chat_history


class PromptManager:
    def __init__(self, prompt: str):
        self.prompt = prompt

    def get_prompt(self):
        return self.prompt


if __name__ == "__main__":
    bot = ChatBot(model="llama3:instruct")
    # retriever = bot.create_retriever("https://docs.smith.langchain.com/user_guide")
    chain = bot.create_chain(with_history=True)
    while True:
        message = input("User: ")
        if message == "quit":
            break
        response = bot.chat(message, stream=True)
        print("ChatBot:")
        for i in response:
            print(i, end="", flush=True)
        print("")
