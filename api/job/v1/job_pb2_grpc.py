# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.job.v1 import job_pb2 as api_dot_job_dot_v1_dot_job__pb2


class JobServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Enqueue = channel.unary_unary(
                '/job.v1.JobService/Enqueue',
                request_serializer=api_dot_job_dot_v1_dot_job__pb2.EnqueueRequest.SerializeToString,
                response_deserializer=api_dot_job_dot_v1_dot_job__pb2.EnqueueResponse.FromString,
                )
        self.QueryTasks = channel.unary_unary(
                '/job.v1.JobService/QueryTasks',
                request_serializer=api_dot_job_dot_v1_dot_job__pb2.QueryTasksRequest.SerializeToString,
                response_deserializer=api_dot_job_dot_v1_dot_job__pb2.QueryTasksResponse.FromString,
                )


class JobServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Enqueue(self, request, context):
        """Enqueue a job
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryTasks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JobServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Enqueue': grpc.unary_unary_rpc_method_handler(
                    servicer.Enqueue,
                    request_deserializer=api_dot_job_dot_v1_dot_job__pb2.EnqueueRequest.FromString,
                    response_serializer=api_dot_job_dot_v1_dot_job__pb2.EnqueueResponse.SerializeToString,
            ),
            'QueryTasks': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryTasks,
                    request_deserializer=api_dot_job_dot_v1_dot_job__pb2.QueryTasksRequest.FromString,
                    response_serializer=api_dot_job_dot_v1_dot_job__pb2.QueryTasksResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'job.v1.JobService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class JobService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Enqueue(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/job.v1.JobService/Enqueue',
            api_dot_job_dot_v1_dot_job__pb2.EnqueueRequest.SerializeToString,
            api_dot_job_dot_v1_dot_job__pb2.EnqueueResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryTasks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/job.v1.JobService/QueryTasks',
            api_dot_job_dot_v1_dot_job__pb2.QueryTasksRequest.SerializeToString,
            api_dot_job_dot_v1_dot_job__pb2.QueryTasksResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)