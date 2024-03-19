# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.sms.v1 import sms_pb2 as api_dot_sms_dot_v1_dot_sms__pb2


class SMSServiceStub(object):
    """短信服务接口
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendSMS = channel.unary_unary(
                '/smsv1.SMSService/SendSMS',
                request_serializer=api_dot_sms_dot_v1_dot_sms__pb2.SendSMSRequest.SerializeToString,
                response_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.SendSMSResponse.FromString,
                )
        self.CheckSMSStatus = channel.unary_unary(
                '/smsv1.SMSService/CheckSMSStatus',
                request_serializer=api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusRequest.SerializeToString,
                response_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusResponse.FromString,
                )
        self.ManageSMSTemplate = channel.unary_unary(
                '/smsv1.SMSService/ManageSMSTemplate',
                request_serializer=api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementRequest.SerializeToString,
                response_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementResponse.FromString,
                )
        self.ManageSMSSignature = channel.unary_unary(
                '/smsv1.SMSService/ManageSMSSignature',
                request_serializer=api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementRequest.SerializeToString,
                response_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementResponse.FromString,
                )


class SMSServiceServicer(object):
    """短信服务接口
    """

    def SendSMS(self, request, context):
        """发送短信
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CheckSMSStatus(self, request, context):
        """查询短信状态
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ManageSMSTemplate(self, request, context):
        """管理短信模板
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ManageSMSSignature(self, request, context):
        """管理短信签名
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SMSServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendSMS': grpc.unary_unary_rpc_method_handler(
                    servicer.SendSMS,
                    request_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.SendSMSRequest.FromString,
                    response_serializer=api_dot_sms_dot_v1_dot_sms__pb2.SendSMSResponse.SerializeToString,
            ),
            'CheckSMSStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.CheckSMSStatus,
                    request_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusRequest.FromString,
                    response_serializer=api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusResponse.SerializeToString,
            ),
            'ManageSMSTemplate': grpc.unary_unary_rpc_method_handler(
                    servicer.ManageSMSTemplate,
                    request_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementRequest.FromString,
                    response_serializer=api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementResponse.SerializeToString,
            ),
            'ManageSMSSignature': grpc.unary_unary_rpc_method_handler(
                    servicer.ManageSMSSignature,
                    request_deserializer=api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementRequest.FromString,
                    response_serializer=api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'smsv1.SMSService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SMSService(object):
    """短信服务接口
    """

    @staticmethod
    def SendSMS(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/smsv1.SMSService/SendSMS',
            api_dot_sms_dot_v1_dot_sms__pb2.SendSMSRequest.SerializeToString,
            api_dot_sms_dot_v1_dot_sms__pb2.SendSMSResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CheckSMSStatus(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/smsv1.SMSService/CheckSMSStatus',
            api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusRequest.SerializeToString,
            api_dot_sms_dot_v1_dot_sms__pb2.QuerySMSStatusResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ManageSMSTemplate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/smsv1.SMSService/ManageSMSTemplate',
            api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementRequest.SerializeToString,
            api_dot_sms_dot_v1_dot_sms__pb2.TemplateManagementResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ManageSMSSignature(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/smsv1.SMSService/ManageSMSSignature',
            api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementRequest.SerializeToString,
            api_dot_sms_dot_v1_dot_sms__pb2.SignatureManagementResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
