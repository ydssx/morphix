# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/media/v1/media.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61pi/media/v1/media.proto\x12\x07mediav1\"N\n\x12VideoUploadRequest\x12\x1c\n\tvideoData\x18\x01 \x01(\x0cR\tvideoData\x12\x1a\n\x08\x66ilename\x18\x02 \x01(\tR\x08\x66ilename\"/\n\x13VideoUploadResponse\x12\x18\n\x07videoId\x18\x01 \x01(\tR\x07videoId\"W\n\x15VideoTranscodeRequest\x12\x18\n\x07videoId\x18\x01 \x01(\tR\x07videoId\x12$\n\routputFormats\x18\x02 \x03(\tR\routputFormats\"H\n\x16VideoTranscodeResponse\x12.\n\x12transcodedVideoIds\x18\x01 \x03(\tR\x12transcodedVideoIds\"0\n\x14VideoPlaybackRequest\x12\x18\n\x07videoId\x18\x01 \x01(\tR\x07videoId\"9\n\x15VideoPlaybackResponse\x12 \n\x0bplaybackUrl\x18\x01 \x01(\tR\x0bplaybackUrl\"P\n\x16VideoScreenshotRequest\x12\x18\n\x07videoId\x18\x01 \x01(\tR\x07videoId\x12\x1c\n\ttimestamp\x18\x02 \x01(\x01R\ttimestamp\"?\n\x17VideoScreenshotResponse\x12$\n\rscreenshotUrl\x18\x01 \x01(\tR\rscreenshotUrl2\xcc\x02\n\x0cMediaService\x12H\n\x0bUploadVideo\x12\x1b.mediav1.VideoUploadRequest\x1a\x1c.mediav1.VideoUploadResponse\x12Q\n\x0eTranscodeVideo\x12\x1e.mediav1.VideoTranscodeRequest\x1a\x1f.mediav1.VideoTranscodeResponse\x12J\n\tPlayVideo\x12\x1d.mediav1.VideoPlaybackRequest\x1a\x1e.mediav1.VideoPlaybackResponse\x12S\n\x0eTakeScreenshot\x12\x1f.mediav1.VideoScreenshotRequest\x1a .mediav1.VideoScreenshotResponseB/Z-github.com/ydssx/morphix/api/media/v1;mediav1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.media.v1.media_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z-github.com/ydssx/morphix/api/media/v1;mediav1'
  _globals['_VIDEOUPLOADREQUEST']._serialized_start=37
  _globals['_VIDEOUPLOADREQUEST']._serialized_end=115
  _globals['_VIDEOUPLOADRESPONSE']._serialized_start=117
  _globals['_VIDEOUPLOADRESPONSE']._serialized_end=164
  _globals['_VIDEOTRANSCODEREQUEST']._serialized_start=166
  _globals['_VIDEOTRANSCODEREQUEST']._serialized_end=253
  _globals['_VIDEOTRANSCODERESPONSE']._serialized_start=255
  _globals['_VIDEOTRANSCODERESPONSE']._serialized_end=327
  _globals['_VIDEOPLAYBACKREQUEST']._serialized_start=329
  _globals['_VIDEOPLAYBACKREQUEST']._serialized_end=377
  _globals['_VIDEOPLAYBACKRESPONSE']._serialized_start=379
  _globals['_VIDEOPLAYBACKRESPONSE']._serialized_end=436
  _globals['_VIDEOSCREENSHOTREQUEST']._serialized_start=438
  _globals['_VIDEOSCREENSHOTREQUEST']._serialized_end=518
  _globals['_VIDEOSCREENSHOTRESPONSE']._serialized_start=520
  _globals['_VIDEOSCREENSHOTRESPONSE']._serialized_end=583
  _globals['_MEDIASERVICE']._serialized_start=586
  _globals['_MEDIASERVICE']._serialized_end=918
# @@protoc_insertion_point(module_scope)
