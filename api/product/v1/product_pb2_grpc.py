# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.product.v1 import product_pb2 as api_dot_product_dot_v1_dot_product__pb2


class ProductServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateProduct = channel.unary_unary(
                '/productv1.ProductService/CreateProduct',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.CreateProductRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.CreateProductResponse.FromString,
                )
        self.GetProducts = channel.unary_unary(
                '/productv1.ProductService/GetProducts',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsResponse.FromString,
                )
        self.GetProduct = channel.unary_unary(
                '/productv1.ProductService/GetProduct',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.Product.FromString,
                )
        self.UpdateProduct = channel.unary_unary(
                '/productv1.ProductService/UpdateProduct',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductResponse.FromString,
                )
        self.DeleteProduct = channel.unary_unary(
                '/productv1.ProductService/DeleteProduct',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.DeleteProductRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.DeleteProductResponse.FromString,
                )
        self.GetProductStock = channel.unary_unary(
                '/productv1.ProductService/GetProductStock',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductStockRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductStockResponse.FromString,
                )
        self.UpdateProductStock = channel.unary_unary(
                '/productv1.ProductService/UpdateProductStock',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockResponse.FromString,
                )
        self.GetProductsStock = channel.unary_unary(
                '/productv1.ProductService/GetProductsStock',
                request_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsStockRequest.SerializeToString,
                response_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsStockResponse.FromString,
                )


class ProductServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateProduct(self, request, context):
        """创建产品
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProducts(self, request, context):
        """获取产品列表
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProduct(self, request, context):
        """获取单个产品
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateProduct(self, request, context):
        """更新产品信息
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteProduct(self, request, context):
        """删除产品
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProductStock(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateProductStock(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProductsStock(self, request, context):
        """获取产品库存
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProductServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateProduct,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.CreateProductRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.CreateProductResponse.SerializeToString,
            ),
            'GetProducts': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProducts,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsResponse.SerializeToString,
            ),
            'GetProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProduct,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.Product.SerializeToString,
            ),
            'UpdateProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateProduct,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductResponse.SerializeToString,
            ),
            'DeleteProduct': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteProduct,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.DeleteProductRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.DeleteProductResponse.SerializeToString,
            ),
            'GetProductStock': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProductStock,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductStockRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductStockResponse.SerializeToString,
            ),
            'UpdateProductStock': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateProductStock,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockResponse.SerializeToString,
            ),
            'GetProductsStock': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProductsStock,
                    request_deserializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsStockRequest.FromString,
                    response_serializer=api_dot_product_dot_v1_dot_product__pb2.GetProductsStockResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'productv1.ProductService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ProductService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/CreateProduct',
            api_dot_product_dot_v1_dot_product__pb2.CreateProductRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.CreateProductResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProducts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/GetProducts',
            api_dot_product_dot_v1_dot_product__pb2.GetProductsRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.GetProductsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/GetProduct',
            api_dot_product_dot_v1_dot_product__pb2.GetProductRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.Product.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/UpdateProduct',
            api_dot_product_dot_v1_dot_product__pb2.UpdateProductRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.UpdateProductResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteProduct(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/DeleteProduct',
            api_dot_product_dot_v1_dot_product__pb2.DeleteProductRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.DeleteProductResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProductStock(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/GetProductStock',
            api_dot_product_dot_v1_dot_product__pb2.GetProductStockRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.GetProductStockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateProductStock(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/UpdateProductStock',
            api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.UpdateProductStockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetProductsStock(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/productv1.ProductService/GetProductsStock',
            api_dot_product_dot_v1_dot_product__pb2.GetProductsStockRequest.SerializeToString,
            api_dot_product_dot_v1_dot_product__pb2.GetProductsStockResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
