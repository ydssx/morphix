"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import api.product.v1.product_pb2
import collections.abc
import grpc
import grpc.aio
import typing

_T = typing.TypeVar('_T')

class _MaybeAsyncIterator(collections.abc.AsyncIterator[_T], collections.abc.Iterator[_T], metaclass=abc.ABCMeta):
    ...

class _ServicerContext(grpc.ServicerContext, grpc.aio.ServicerContext):  # type: ignore
    ...

class ProductServiceStub:
    def __init__(self, channel: typing.Union[grpc.Channel, grpc.aio.Channel]) -> None: ...
    CreateProduct: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.CreateProductRequest,
        api.product.v1.product_pb2.CreateProductResponse,
    ]
    """创建产品"""
    GetProducts: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductsRequest,
        api.product.v1.product_pb2.GetProductsResponse,
    ]
    """获取产品列表"""
    GetProduct: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductRequest,
        api.product.v1.product_pb2.Product,
    ]
    """获取单个产品"""
    UpdateProduct: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.UpdateProductRequest,
        api.product.v1.product_pb2.UpdateProductResponse,
    ]
    """更新产品信息"""
    DeleteProduct: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.DeleteProductRequest,
        api.product.v1.product_pb2.DeleteProductResponse,
    ]
    """删除产品"""
    GetProductStock: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductStockRequest,
        api.product.v1.product_pb2.GetProductStockResponse,
    ]
    UpdateProductStock: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.UpdateProductStockRequest,
        api.product.v1.product_pb2.UpdateProductStockResponse,
    ]
    GetProductsStock: grpc.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductsStockRequest,
        api.product.v1.product_pb2.GetProductsStockResponse,
    ]
    """获取产品库存"""

class ProductServiceAsyncStub:
    CreateProduct: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.CreateProductRequest,
        api.product.v1.product_pb2.CreateProductResponse,
    ]
    """创建产品"""
    GetProducts: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductsRequest,
        api.product.v1.product_pb2.GetProductsResponse,
    ]
    """获取产品列表"""
    GetProduct: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductRequest,
        api.product.v1.product_pb2.Product,
    ]
    """获取单个产品"""
    UpdateProduct: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.UpdateProductRequest,
        api.product.v1.product_pb2.UpdateProductResponse,
    ]
    """更新产品信息"""
    DeleteProduct: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.DeleteProductRequest,
        api.product.v1.product_pb2.DeleteProductResponse,
    ]
    """删除产品"""
    GetProductStock: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductStockRequest,
        api.product.v1.product_pb2.GetProductStockResponse,
    ]
    UpdateProductStock: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.UpdateProductStockRequest,
        api.product.v1.product_pb2.UpdateProductStockResponse,
    ]
    GetProductsStock: grpc.aio.UnaryUnaryMultiCallable[
        api.product.v1.product_pb2.GetProductsStockRequest,
        api.product.v1.product_pb2.GetProductsStockResponse,
    ]
    """获取产品库存"""

class ProductServiceServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def CreateProduct(
        self,
        request: api.product.v1.product_pb2.CreateProductRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.CreateProductResponse, collections.abc.Awaitable[api.product.v1.product_pb2.CreateProductResponse]]:
        """创建产品"""
    @abc.abstractmethod
    def GetProducts(
        self,
        request: api.product.v1.product_pb2.GetProductsRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.GetProductsResponse, collections.abc.Awaitable[api.product.v1.product_pb2.GetProductsResponse]]:
        """获取产品列表"""
    @abc.abstractmethod
    def GetProduct(
        self,
        request: api.product.v1.product_pb2.GetProductRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.Product, collections.abc.Awaitable[api.product.v1.product_pb2.Product]]:
        """获取单个产品"""
    @abc.abstractmethod
    def UpdateProduct(
        self,
        request: api.product.v1.product_pb2.UpdateProductRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.UpdateProductResponse, collections.abc.Awaitable[api.product.v1.product_pb2.UpdateProductResponse]]:
        """更新产品信息"""
    @abc.abstractmethod
    def DeleteProduct(
        self,
        request: api.product.v1.product_pb2.DeleteProductRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.DeleteProductResponse, collections.abc.Awaitable[api.product.v1.product_pb2.DeleteProductResponse]]:
        """删除产品"""
    @abc.abstractmethod
    def GetProductStock(
        self,
        request: api.product.v1.product_pb2.GetProductStockRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.GetProductStockResponse, collections.abc.Awaitable[api.product.v1.product_pb2.GetProductStockResponse]]: ...
    @abc.abstractmethod
    def UpdateProductStock(
        self,
        request: api.product.v1.product_pb2.UpdateProductStockRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.UpdateProductStockResponse, collections.abc.Awaitable[api.product.v1.product_pb2.UpdateProductStockResponse]]: ...
    @abc.abstractmethod
    def GetProductsStock(
        self,
        request: api.product.v1.product_pb2.GetProductsStockRequest,
        context: _ServicerContext,
    ) -> typing.Union[api.product.v1.product_pb2.GetProductsStockResponse, collections.abc.Awaitable[api.product.v1.product_pb2.GetProductsStockResponse]]:
        """获取产品库存"""

def add_ProductServiceServicer_to_server(servicer: ProductServiceServicer, server: typing.Union[grpc.Server, grpc.aio.Server]) -> None: ...
