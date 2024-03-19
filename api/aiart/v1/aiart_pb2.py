# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/aiart/v1/aiart.proto
# Protobuf Python Version: 5.26.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61pi/aiart/v1/aiart.proto\x12\x07\x61iartv1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x96\x05\n\x14GenerateImageRequest\x12%\n\x0eoriginal_image\x18\x01 \x01(\x0cR\roriginalImage\x12\x16\n\x06prompt\x18\x02 \x01(\tR\x06prompt\x12\x1d\n\nimage_size\x18\x03 \x01(\x05R\timageSize\x12%\n\x0eguidance_scale\x18\x04 \x01(\x01R\rguidanceScale\x12g\n\x15generation_parameters\x18\x05 \x01(\x0b\x32\x32.aiartv1.GenerateImageRequest.GenerationParametersR\x14generationParameters\x12\x65\n\x15image_generation_mode\x18\x06 \x01(\x0e\x32\x31.aiartv1.GenerateImageRequest.ImageGenerationModeR\x13imageGenerationMode\x1a\x85\x01\n\x14GenerationParameters\x12\x1d\n\nmodel_name\x18\x01 \x01(\tR\tmodelName\x12%\n\x0esampling_steps\x18\x02 \x01(\x05R\rsamplingSteps\x12\'\n\x0fsampling_method\x18\x03 \x01(\tR\x0esamplingMethod\x1aS\n\x12GuidanceParameters\x12\x16\n\x06prompt\x18\x01 \x01(\tR\x06prompt\x12%\n\x0eguidance_scale\x18\x02 \x01(\x01R\rguidanceScale\"L\n\x13ImageGenerationMode\x12\x19\n\x15GENERATE_MODE_DEFAULT\x10\x00\x12\x1a\n\x16GENERATE_MODE_GUIDANCE\x10\x01\"2\n\x15GenerateImageResponse\x12\x19\n\x08image_id\x18\x01 \x01(\tR\x07imageId\"\x85\x01\n\x14GenerationParameters\x12\x1d\n\nmodel_name\x18\x01 \x01(\tR\tmodelName\x12%\n\x0esampling_steps\x18\x02 \x01(\x05R\rsamplingSteps\x12\'\n\x0fsampling_method\x18\x03 \x01(\tR\x0esamplingMethod\"5\n\x18GetGenerateStatusRequest\x12\x19\n\x08image_id\x18\x01 \x01(\tR\x07imageId\"I\n\x16GenerateStatusResponse\x12/\n\x06status\x18\x01 \x01(\x0e\x32\x17.aiartv1.GenerateStatusR\x06status\"5\n\x18GetGeneratedImageRequest\x12\x19\n\x08image_id\x18\x01 \x01(\tR\x07imageId\":\n\x19GetGeneratedImageResponse\x12\x1d\n\nimage_data\x18\x01 \x01(\x0cR\timageData\"Z\n\x14GetModelInfoResponse\x12\x1d\n\nmodel_name\x18\x01 \x01(\tR\tmodelName\x12#\n\rmodel_version\x18\x02 \x01(\tR\x0cmodelVersion\"\xea\x01\n\x13ImageToImageRequest\x12%\n\x0eoriginal_image\x18\x01 \x01(\x0cR\roriginalImage\x12\x16\n\x06prompt\x18\x02 \x01(\tR\x06prompt\x12\x1d\n\nimage_size\x18\x03 \x01(\x05R\timageSize\x12%\n\x0eguidance_scale\x18\x04 \x01(\x01R\rguidanceScale\x12%\n\x0esampling_steps\x18\x05 \x01(\x05R\rsamplingSteps\x12\'\n\x0fsampling_method\x18\x06 \x01(\tR\x0esamplingMethod\"1\n\x14ImageToImageResponse\x12\x19\n\x08image_id\x18\x01 \x01(\tR\x07imageId*m\n\x0eGenerateStatus\x12\x0b\n\x07PENDING\x10\x00\x12\x0b\n\x07RUNNING\x10\x01\x12\r\n\tCOMPLETED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x12\x0c\n\x08\x43\x41NCELED\x10\x04\x12\x0b\n\x07TIMEOUT\x10\x05\x12\x0b\n\x07UNKNOWN\x10\x06\x32\xae\x04\n\nArtService\x12n\n\rGenerateImage\x12\x1d.aiartv1.GenerateImageRequest\x1a\x1e.aiartv1.GenerateImageResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/images/generate:\x01*\x12{\n\x11GetGenerateStatus\x12!.aiartv1.GetGenerateStatusRequest\x1a\x1f.aiartv1.GenerateStatusResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/images/generate_status\x12y\n\x11GetGeneratedImage\x12!.aiartv1.GetGeneratedImageRequest\x1a\".aiartv1.GetGeneratedImageResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/v1/images/{image_id}\x12\x45\n\x0cGetModelInfo\x12\x16.google.protobuf.Empty\x1a\x1d.aiartv1.GetModelInfoResponse\x12q\n\x0cImageToImage\x12\x1c.aiartv1.ImageToImageRequest\x1a\x1d.aiartv1.ImageToImageResponse\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/v1/images/image_to_image:\x01*B/Z-github.com/ydssx/morphix/api/aiart/v1;aiartv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.aiart.v1.aiart_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z-github.com/ydssx/morphix/api/aiart/v1;aiartv1'
  _globals['_ARTSERVICE'].methods_by_name['GenerateImage']._loaded_options = None
  _globals['_ARTSERVICE'].methods_by_name['GenerateImage']._serialized_options = b'\202\323\344\223\002\030\"\023/v1/images/generate:\001*'
  _globals['_ARTSERVICE'].methods_by_name['GetGenerateStatus']._loaded_options = None
  _globals['_ARTSERVICE'].methods_by_name['GetGenerateStatus']._serialized_options = b'\202\323\344\223\002\034\022\032/v1/images/generate_status'
  _globals['_ARTSERVICE'].methods_by_name['GetGeneratedImage']._loaded_options = None
  _globals['_ARTSERVICE'].methods_by_name['GetGeneratedImage']._serialized_options = b'\202\323\344\223\002\027\022\025/v1/images/{image_id}'
  _globals['_ARTSERVICE'].methods_by_name['ImageToImage']._loaded_options = None
  _globals['_ARTSERVICE'].methods_by_name['ImageToImage']._serialized_options = b'\202\323\344\223\002\036\"\031/v1/images/image_to_image:\001*'
  _globals['_GENERATESTATUS']._serialized_start=1574
  _globals['_GENERATESTATUS']._serialized_end=1683
  _globals['_GENERATEIMAGEREQUEST']._serialized_start=97
  _globals['_GENERATEIMAGEREQUEST']._serialized_end=759
  _globals['_GENERATEIMAGEREQUEST_GENERATIONPARAMETERS']._serialized_start=463
  _globals['_GENERATEIMAGEREQUEST_GENERATIONPARAMETERS']._serialized_end=596
  _globals['_GENERATEIMAGEREQUEST_GUIDANCEPARAMETERS']._serialized_start=598
  _globals['_GENERATEIMAGEREQUEST_GUIDANCEPARAMETERS']._serialized_end=681
  _globals['_GENERATEIMAGEREQUEST_IMAGEGENERATIONMODE']._serialized_start=683
  _globals['_GENERATEIMAGEREQUEST_IMAGEGENERATIONMODE']._serialized_end=759
  _globals['_GENERATEIMAGERESPONSE']._serialized_start=761
  _globals['_GENERATEIMAGERESPONSE']._serialized_end=811
  _globals['_GENERATIONPARAMETERS']._serialized_start=463
  _globals['_GENERATIONPARAMETERS']._serialized_end=596
  _globals['_GETGENERATESTATUSREQUEST']._serialized_start=949
  _globals['_GETGENERATESTATUSREQUEST']._serialized_end=1002
  _globals['_GENERATESTATUSRESPONSE']._serialized_start=1004
  _globals['_GENERATESTATUSRESPONSE']._serialized_end=1077
  _globals['_GETGENERATEDIMAGEREQUEST']._serialized_start=1079
  _globals['_GETGENERATEDIMAGEREQUEST']._serialized_end=1132
  _globals['_GETGENERATEDIMAGERESPONSE']._serialized_start=1134
  _globals['_GETGENERATEDIMAGERESPONSE']._serialized_end=1192
  _globals['_GETMODELINFORESPONSE']._serialized_start=1194
  _globals['_GETMODELINFORESPONSE']._serialized_end=1284
  _globals['_IMAGETOIMAGEREQUEST']._serialized_start=1287
  _globals['_IMAGETOIMAGEREQUEST']._serialized_end=1521
  _globals['_IMAGETOIMAGERESPONSE']._serialized_start=1523
  _globals['_IMAGETOIMAGERESPONSE']._serialized_end=1572
  _globals['_ARTSERVICE']._serialized_start=1686
  _globals['_ARTSERVICE']._serialized_end=2244
# @@protoc_insertion_point(module_scope)
