import streamlit as st

from ai.chatbot import ChatBot
from streamlit_multipage import MultiPage
from streamlit_option_menu import option_menu

# st.title("Chat Bot")


@st.cache_resource
def init_model(session_id):
    bot = ChatBot(session_id)
    llm = bot.create_chain(with_history=True)
    return bot


hsi = st.sidebar.title("会话管理")

new_session_name = st.sidebar.text_input('输入新会话名称：')
create_session_button = st.sidebar.button('创建新会话')

# 如果用户点击创建新会话按钮，则创建新的会话
if create_session_button and new_session_name:
    session_list = new_session_name
    print(session_list)
    ChatBot.create_session(session_list)
    
sessions = ChatBot.get_session_ids()
session_id = st.sidebar.selectbox('选择会话', sessions)


def setup(session_id):
    # st.sidebar.title(session_id)
    # bot.get_session_history(session_id)
    bot = init_model(session_id)

    msgs= bot.get_session_history(session_id).messages
    print(len(msgs))
    for message in msgs:
        if message.type == "AIMessageChunk":
            message.type = "assistant"
        with st.chat_message(message.type):
            st.markdown(message.content)

    if prompt := st.chat_input("请输入消息"):
        # Display user message in chat message container
        with st.chat_message("user"):
            st.markdown(prompt)
        with st.chat_message("assistant"):
            resp = bot.chat(prompt)
            response = st.write_stream(resp)

if session_id:
    setup(session_id)
