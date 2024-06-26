from typing import List

from fastapi import FastAPI
from langchain.pydantic_v1 import BaseModel, Field
from langchain_core.messages import BaseMessage
from langserve import add_routes

from ai.chatbot import ChatBot

app = FastAPI(
    title="LangChain Server",
    version="1.0",
    description="A simple API server using LangChain's Runnable interfaces",
)


class Input(BaseModel):
    input: str
    context: str = Field(default="")
    chat_history: List[BaseMessage] = Field(
        ...,
        extra={"widget": {"type": "chat", "input": "location"}},
    )


class Output(BaseModel):
    output: str


chain = ChatBot().create_agent()
add_routes(
    app,
    chain.with_types(input_type=Input, output_type=Output),
    path="/agent",
)

if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="localhost", port=8000)
