# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.media.v1 import media_pb2 as api_dot_media_dot_v1_dot_media__pb2


class MediaServiceStub(object):
    """定义媒体服务接口
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UploadVideo = channel.unary_unary(
                '/mediav1.MediaService/UploadVideo',
                request_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoUploadRequest.SerializeToString,
                response_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoUploadResponse.FromString,
                )
        self.TranscodeVideo = channel.unary_unary(
                '/mediav1.MediaService/TranscodeVideo',
                request_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeRequest.SerializeToString,
                response_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeResponse.FromString,
                )
        self.PlayVideo = channel.unary_unary(
                '/mediav1.MediaService/PlayVideo',
                request_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackRequest.SerializeToString,
                response_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackResponse.FromString,
                )
        self.TakeScreenshot = channel.unary_unary(
                '/mediav1.MediaService/TakeScreenshot',
                request_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotRequest.SerializeToString,
                response_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotResponse.FromString,
                )


class MediaServiceServicer(object):
    """定义媒体服务接口
    """

    def UploadVideo(self, request, context):
        """视频上传
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TranscodeVideo(self, request, context):
        """视频转码
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PlayVideo(self, request, context):
        """视频播放
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TakeScreenshot(self, request, context):
        """视频截图
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MediaServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UploadVideo': grpc.unary_unary_rpc_method_handler(
                    servicer.UploadVideo,
                    request_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoUploadRequest.FromString,
                    response_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoUploadResponse.SerializeToString,
            ),
            'TranscodeVideo': grpc.unary_unary_rpc_method_handler(
                    servicer.TranscodeVideo,
                    request_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeRequest.FromString,
                    response_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeResponse.SerializeToString,
            ),
            'PlayVideo': grpc.unary_unary_rpc_method_handler(
                    servicer.PlayVideo,
                    request_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackRequest.FromString,
                    response_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackResponse.SerializeToString,
            ),
            'TakeScreenshot': grpc.unary_unary_rpc_method_handler(
                    servicer.TakeScreenshot,
                    request_deserializer=api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotRequest.FromString,
                    response_serializer=api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mediav1.MediaService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MediaService(object):
    """定义媒体服务接口
    """

    @staticmethod
    def UploadVideo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mediav1.MediaService/UploadVideo',
            api_dot_media_dot_v1_dot_media__pb2.VideoUploadRequest.SerializeToString,
            api_dot_media_dot_v1_dot_media__pb2.VideoUploadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TranscodeVideo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mediav1.MediaService/TranscodeVideo',
            api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeRequest.SerializeToString,
            api_dot_media_dot_v1_dot_media__pb2.VideoTranscodeResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PlayVideo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mediav1.MediaService/PlayVideo',
            api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackRequest.SerializeToString,
            api_dot_media_dot_v1_dot_media__pb2.VideoPlaybackResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TakeScreenshot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mediav1.MediaService/TakeScreenshot',
            api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotRequest.SerializeToString,
            api_dot_media_dot_v1_dot_media__pb2.VideoScreenshotResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
