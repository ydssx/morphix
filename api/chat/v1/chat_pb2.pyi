"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class ClientMessage(google.protobuf.message.Message):
    """客户端发送的消息"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    MESSAGE_TEXT_FIELD_NUMBER: builtins.int
    IMG_URL_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    message_text: builtins.str
    img_url: builtins.str
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
        message_text: builtins.str = ...,
        img_url: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["img_url", b"img_url", "message_text", b"message_text", "user_id", b"user_id"]) -> None: ...

global___ClientMessage = ClientMessage

@typing_extensions.final
class ServerMessage(google.protobuf.message.Message):
    """服务器发送的消息"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SENDER_ID_FIELD_NUMBER: builtins.int
    MESSAGE_TEXT_FIELD_NUMBER: builtins.int
    sender_id: builtins.str
    message_text: builtins.str
    def __init__(
        self,
        *,
        sender_id: builtins.str = ...,
        message_text: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["message_text", b"message_text", "sender_id", b"sender_id"]) -> None: ...

global___ServerMessage = ServerMessage

@typing_extensions.final
class ChatMessage(google.protobuf.message.Message):
    """客户端和服务器之间的聊天消息"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    MESSAGE_TEXT_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    message_text: builtins.str
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
        message_text: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["message_text", b"message_text", "user_id", b"user_id"]) -> None: ...

global___ChatMessage = ChatMessage
