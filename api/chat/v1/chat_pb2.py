# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/chat/v1/chat.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x61pi/chat/v1/chat.proto\x12\x04\x63hat\x1a\x1cgoogle/api/annotations.proto\"d\n\rClientMessage\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12!\n\x0cmessage_text\x18\x02 \x01(\tR\x0bmessageText\x12\x17\n\x07img_url\x18\x03 \x01(\tR\x06imgUrl\"O\n\rServerMessage\x12\x1b\n\tsender_id\x18\x01 \x01(\tR\x08senderId\x12!\n\x0cmessage_text\x18\x02 \x01(\tR\x0bmessageText\"I\n\x0b\x43hatMessage\x12\x17\n\x07user_id\x18\x01 \x01(\tR\x06userId\x12!\n\x0cmessage_text\x18\x02 \x01(\tR\x0bmessageText2\xf7\x01\n\x0b\x43hatService\x12W\n\x0bSendMessage\x12\x13.chat.ClientMessage\x1a\x13.chat.ServerMessage\"\x1c\x82\xd3\xe4\x93\x02\x16\"\x11/api/v1/chat/send:\x01*(\x01\x12\x30\n\x04\x43hat\x12\x11.chat.ChatMessage\x1a\x11.chat.ChatMessage(\x01\x30\x01\x12]\n\x0eReceiveMessage\x12\x13.chat.ClientMessage\x1a\x13.chat.ServerMessage\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/chat/receive:\x01*0\x01\x42+Z)github.com/ydssx/morphix/api/chat/v1;chatb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.chat.v1.chat_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/ydssx/morphix/api/chat/v1;chat'
  _globals['_CHATSERVICE'].methods_by_name['SendMessage']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['SendMessage']._serialized_options = b'\202\323\344\223\002\026\"\021/api/v1/chat/send:\001*'
  _globals['_CHATSERVICE'].methods_by_name['ReceiveMessage']._loaded_options = None
  _globals['_CHATSERVICE'].methods_by_name['ReceiveMessage']._serialized_options = b'\202\323\344\223\002\031\"\024/api/v1/chat/receive:\001*'
  _globals['_CLIENTMESSAGE']._serialized_start=62
  _globals['_CLIENTMESSAGE']._serialized_end=162
  _globals['_SERVERMESSAGE']._serialized_start=164
  _globals['_SERVERMESSAGE']._serialized_end=243
  _globals['_CHATMESSAGE']._serialized_start=245
  _globals['_CHATMESSAGE']._serialized_end=318
  _globals['_CHATSERVICE']._serialized_start=321
  _globals['_CHATSERVICE']._serialized_end=568
# @@protoc_insertion_point(module_scope)
