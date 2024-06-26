# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.chat.v1 import chat_pb2 as api_dot_chat_dot_v1_dot_chat__pb2


class ChatServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendMessage = channel.stream_unary(
                '/chat.ChatService/SendMessage',
                request_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.SerializeToString,
                response_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.FromString,
                )
        self.Chat = channel.stream_stream(
                '/chat.ChatService/Chat',
                request_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.SerializeToString,
                response_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.FromString,
                )
        self.ReceiveMessage = channel.unary_stream(
                '/chat.ChatService/ReceiveMessage',
                request_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.SerializeToString,
                response_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.FromString,
                )


class ChatServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendMessage(self, request_iterator, context):
        """客户端到服务器的流，用于发送消息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Chat(self, request_iterator, context):
        """双向流，用于实现聊天
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReceiveMessage(self, request, context):
        """服务器到客户端的流，用于接收消息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ChatServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendMessage': grpc.stream_unary_rpc_method_handler(
                    servicer.SendMessage,
                    request_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.FromString,
                    response_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.SerializeToString,
            ),
            'Chat': grpc.stream_stream_rpc_method_handler(
                    servicer.Chat,
                    request_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.FromString,
                    response_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.SerializeToString,
            ),
            'ReceiveMessage': grpc.unary_stream_rpc_method_handler(
                    servicer.ReceiveMessage,
                    request_deserializer=api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.FromString,
                    response_serializer=api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'chat.ChatService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ChatService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendMessage(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/chat.ChatService/SendMessage',
            api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.SerializeToString,
            api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Chat(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/chat.ChatService/Chat',
            api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.SerializeToString,
            api_dot_chat_dot_v1_dot_chat__pb2.ChatMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ReceiveMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/chat.ChatService/ReceiveMessage',
            api_dot_chat_dot_v1_dot_chat__pb2.ClientMessage.SerializeToString,
            api_dot_chat_dot_v1_dot_chat__pb2.ServerMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
