# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.order.v1 import order_pb2 as api_dot_order_dot_v1_dot_order__pb2


class OrderServiceStub(object):
    """订单管理服务接口
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateOrder = channel.unary_unary(
                '/orderv1.OrderService/CreateOrder',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.CreateOrderRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.CreateOrderResponse.FromString,
                )
        self.GetOrder = channel.unary_unary(
                '/orderv1.OrderService/GetOrder',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.GetOrderRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.GetOrderResponse.FromString,
                )
        self.UpdateOrderStatus = channel.unary_unary(
                '/orderv1.OrderService/UpdateOrderStatus',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusResponse.FromString,
                )
        self.PayOrder = channel.unary_unary(
                '/orderv1.OrderService/PayOrder',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.PayOrderRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.PayOrderResponse.FromString,
                )
        self.DeleteOrder = channel.unary_unary(
                '/orderv1.OrderService/DeleteOrder',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.DeleteOrderRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.DeleteOrderResponse.FromString,
                )
        self.ListOrders = channel.unary_unary(
                '/orderv1.OrderService/ListOrders',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.ListOrdersRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.ListOrdersResponse.FromString,
                )
        self.CancelOrder = channel.unary_unary(
                '/orderv1.OrderService/CancelOrder',
                request_serializer=api_dot_order_dot_v1_dot_order__pb2.CancelOrderRequest.SerializeToString,
                response_deserializer=api_dot_order_dot_v1_dot_order__pb2.CancelOrderResponse.FromString,
                )


class OrderServiceServicer(object):
    """订单管理服务接口
    """

    def CreateOrder(self, request, context):
        """创建订单
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrder(self, request, context):
        """查询订单
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateOrderStatus(self, request, context):
        """更新订单状态
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PayOrder(self, request, context):
        """支付订单
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteOrder(self, request, context):
        """删除订单
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListOrders(self, request, context):
        """查询订单列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CancelOrder(self, request, context):
        """取消订单
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OrderServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateOrder,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.CreateOrderRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.CreateOrderResponse.SerializeToString,
            ),
            'GetOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOrder,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.GetOrderRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.GetOrderResponse.SerializeToString,
            ),
            'UpdateOrderStatus': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateOrderStatus,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusResponse.SerializeToString,
            ),
            'PayOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.PayOrder,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.PayOrderRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.PayOrderResponse.SerializeToString,
            ),
            'DeleteOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteOrder,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.DeleteOrderRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.DeleteOrderResponse.SerializeToString,
            ),
            'ListOrders': grpc.unary_unary_rpc_method_handler(
                    servicer.ListOrders,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.ListOrdersRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.ListOrdersResponse.SerializeToString,
            ),
            'CancelOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.CancelOrder,
                    request_deserializer=api_dot_order_dot_v1_dot_order__pb2.CancelOrderRequest.FromString,
                    response_serializer=api_dot_order_dot_v1_dot_order__pb2.CancelOrderResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'orderv1.OrderService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OrderService(object):
    """订单管理服务接口
    """

    @staticmethod
    def CreateOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/CreateOrder',
            api_dot_order_dot_v1_dot_order__pb2.CreateOrderRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.CreateOrderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/GetOrder',
            api_dot_order_dot_v1_dot_order__pb2.GetOrderRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.GetOrderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateOrderStatus(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/UpdateOrderStatus',
            api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.UpdateOrderStatusResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PayOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/PayOrder',
            api_dot_order_dot_v1_dot_order__pb2.PayOrderRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.PayOrderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/DeleteOrder',
            api_dot_order_dot_v1_dot_order__pb2.DeleteOrderRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.DeleteOrderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListOrders(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/ListOrders',
            api_dot_order_dot_v1_dot_order__pb2.ListOrdersRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.ListOrdersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CancelOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/orderv1.OrderService/CancelOrder',
            api_dot_order_dot_v1_dot_order__pb2.CancelOrderRequest.SerializeToString,
            api_dot_order_dot_v1_dot_order__pb2.CancelOrderResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
