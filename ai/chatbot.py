from langchain_community.chat_models import ChatOllama
from langchain.memory import ChatMessageHistory
from langchain.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_community.document_loaders import WebBaseLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_community.embeddings import OllamaEmbeddings
from langchain.vectorstores.chroma import Chroma
from langchain_core.messages import HumanMessage, AIMessage
from langchain.chains.combine_documents import create_stuff_documents_chain
from langchain.chains.retrieval import create_retrieval_chain
from langchain.chains.history_aware_retriever import create_history_aware_retriever
from langchain.tools.retriever import create_retriever_tool
from langchain import hub
from langchain.agents import AgentExecutor, create_structured_chat_agent

prompt = ChatPromptTemplate.from_messages(
    [
        (
            "system",
            "You are a helpful assistant. Answer the user's questions in Chinese based on the below context:\n\n{context}",
        ),
        MessagesPlaceholder(variable_name="chat_history"),
        ("user", "{input}"),
    ]
)


class ChatBot:
    def __init__(self, model="llama3:instruct"):
        self.chatbot = ChatOllama(model=model)
        # self.chat_history = ChatMessageHistory()
        # self.chat_history.clear()
        self.chat_history = []
        self.retriever = None
        self.chain = prompt | self.chatbot
        self.agent = None

    def chat(self, message: str):
        self.chat_history.append(HumanMessage(message))
        docs = self.retriever.invoke(message)
        response = self.chain.invoke(
            {
                "chat_history": self.chat_history,
                "context": docs,
                "input": message,
            }
        )

        self.chat_history.append(AIMessage(content=response["answer"]))
        return response

    def load_documents(self, url: str):
        """
        Load documents from a web page

        Parameters
        ----------
        url : str
            The URL of the web page
        """
        loader = WebBaseLoader(url)
        docs = loader.load()

        text_splitter = RecursiveCharacterTextSplitter(chunk_size=500, chunk_overlap=0)
        all_splits = text_splitter.split_documents(docs)

        # Create a Chroma vector store from the split documents
        vectorstore = Chroma.from_documents(
            documents=all_splits,
            embedding=OllamaEmbeddings(model="llama3:instruct"),
        )

        # Create a retriever from the vector store
        retriever = vectorstore.as_retriever(k=4)
        self.retriever = retriever

        retriever_chain = create_history_aware_retriever(
            self.chatbot, retriever, prompt
        )

        document_chain = create_stuff_documents_chain(self.chatbot, prompt)

        self.chain = create_retrieval_chain(retriever_chain, document_chain)

    def agent_chat(self, message: str):
        agent_executor = self.create_agent()

        self.chat_history.append(HumanMessage(message))
        response = agent_executor.invoke(
            {
                "chat_history": self.chat_history,
                "input": message,
            }
        )
        # self.chat_history.append(AIMessage(content=response["action_input"]))
        return response

    def create_agent(self):
        prompt = hub.pull("hwchase17/structured-chat-agent")
        retriever_tool = create_retriever_tool(
            self.retriever,
            "langsmith_search",
            "Search for information about LangSmith. For any questions about LangSmith, you must use this tool!",
        )

        tools = [retriever_tool]
        agent = create_structured_chat_agent(self.chatbot, tools, prompt)
        agent_executor = AgentExecutor(
            agent=agent, tools=tools, verbose=True, handle_parsing_errors=True
        )
        return agent_executor


if __name__ == "__main__":
    bot = ChatBot(model="llama3:instruct")
    bot.load_documents("https://docs.smith.langchain.com/user_guide")
    while True:
        message = input("User: ")
        if message == "quit":
            break
        response = bot.agent_chat(message)
        print(f"ChatBot: {response['output']}")
