# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.quote.v1 import quote_pb2 as api_dot_quote_dot_v1_dot_quote__pb2


class QuoteServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateQuote = channel.unary_unary(
                '/quote.QuoteService/CreateQuote',
                request_serializer=api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteRequest.SerializeToString,
                response_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteResponse.FromString,
                )
        self.GetQuotes = channel.unary_unary(
                '/quote.QuoteService/GetQuotes',
                request_serializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesRequest.SerializeToString,
                response_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesResponse.FromString,
                )
        self.GetQuote = channel.unary_unary(
                '/quote.QuoteService/GetQuote',
                request_serializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuoteRequest.SerializeToString,
                response_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.Quote.FromString,
                )
        self.GetUserCoupons = channel.unary_unary(
                '/quote.QuoteService/GetUserCoupons',
                request_serializer=api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsRequest.SerializeToString,
                response_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsResponse.FromString,
                )
        self.UseCoupon = channel.unary_unary(
                '/quote.QuoteService/UseCoupon',
                request_serializer=api_dot_quote_dot_v1_dot_quote__pb2.UseCouponRequest.SerializeToString,
                response_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.UseCouponResponse.FromString,
                )


class QuoteServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateQuote(self, request, context):
        """创建报价
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetQuotes(self, request, context):
        """获取报价列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetQuote(self, request, context):
        """获取单个报价
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserCoupons(self, request, context):
        """获取用户拥有的优惠券列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UseCoupon(self, request, context):
        """使用优惠券
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_QuoteServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateQuote': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateQuote,
                    request_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteRequest.FromString,
                    response_serializer=api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteResponse.SerializeToString,
            ),
            'GetQuotes': grpc.unary_unary_rpc_method_handler(
                    servicer.GetQuotes,
                    request_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesRequest.FromString,
                    response_serializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesResponse.SerializeToString,
            ),
            'GetQuote': grpc.unary_unary_rpc_method_handler(
                    servicer.GetQuote,
                    request_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.GetQuoteRequest.FromString,
                    response_serializer=api_dot_quote_dot_v1_dot_quote__pb2.Quote.SerializeToString,
            ),
            'GetUserCoupons': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserCoupons,
                    request_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsRequest.FromString,
                    response_serializer=api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsResponse.SerializeToString,
            ),
            'UseCoupon': grpc.unary_unary_rpc_method_handler(
                    servicer.UseCoupon,
                    request_deserializer=api_dot_quote_dot_v1_dot_quote__pb2.UseCouponRequest.FromString,
                    response_serializer=api_dot_quote_dot_v1_dot_quote__pb2.UseCouponResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'quote.QuoteService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class QuoteService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateQuote(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/quote.QuoteService/CreateQuote',
            api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteRequest.SerializeToString,
            api_dot_quote_dot_v1_dot_quote__pb2.CreateQuoteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetQuotes(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/quote.QuoteService/GetQuotes',
            api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesRequest.SerializeToString,
            api_dot_quote_dot_v1_dot_quote__pb2.GetQuotesResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetQuote(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/quote.QuoteService/GetQuote',
            api_dot_quote_dot_v1_dot_quote__pb2.GetQuoteRequest.SerializeToString,
            api_dot_quote_dot_v1_dot_quote__pb2.Quote.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserCoupons(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/quote.QuoteService/GetUserCoupons',
            api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsRequest.SerializeToString,
            api_dot_quote_dot_v1_dot_quote__pb2.GetUserCouponsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UseCoupon(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/quote.QuoteService/UseCoupon',
            api_dot_quote_dot_v1_dot_quote__pb2.UseCouponRequest.SerializeToString,
            api_dot_quote_dot_v1_dot_quote__pb2.UseCouponResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)