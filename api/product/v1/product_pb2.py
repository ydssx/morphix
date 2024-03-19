# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/product/v1/product.proto
# Protobuf Python Version: 5.26.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x61pi/product/v1/product.proto\x12\tproductv1\x1a\x1cgoogle/api/annotations.proto\"+\n\x17GetProductsStockRequest\x12\x10\n\x03ids\x18\x01 \x03(\x03R\x03ids\"\x9e\x01\n\x18GetProductsStockResponse\x12G\n\x06stocks\x18\x01 \x03(\x0b\x32/.productv1.GetProductsStockResponse.StocksEntryR\x06stocks\x1a\x39\n\x0bStocksEntry\x12\x10\n\x03key\x18\x01 \x01(\x03R\x03key\x12\x14\n\x05value\x18\x02 \x01(\x05R\x05value:\x02\x38\x01\"(\n\x16GetProductStockRequest\x12\x0e\n\x02id\x18\x01 \x01(\x03R\x02id\"/\n\x17GetProductStockResponse\x12\x14\n\x05stock\x18\x01 \x01(\x05R\x05stock\"A\n\x19UpdateProductStockRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x05stock\x18\x02 \x01(\x05R\x05stock\"\x1c\n\x1aUpdateProductStockResponse\"{\n\x07Product\x12\x0e\n\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x14\n\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n\x05stock\x18\x05 \x01(\x05R\x05stock\"x\n\x14\x43reateProductRequest\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x02 \x01(\tR\x0b\x64\x65scription\x12\x14\n\x05price\x18\x03 \x01(\x02R\x05price\x12\x14\n\x05stock\x18\x04 \x01(\x05R\x05stock\"\'\n\x15\x43reateProductResponse\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"z\n\x12GetProductsRequest\x12\x1b\n\tpage_size\x18\x01 \x01(\x03R\x08pageSize\x12\x12\n\x04page\x18\x02 \x01(\x03R\x04page\x12\x12\n\x04name\x18\x03 \x01(\tR\x04name\x12\x1f\n\x0bproduct_ids\x18\x04 \x03(\x03R\nproductIds\"[\n\x13GetProductsResponse\x12.\n\x08products\x18\x01 \x03(\x0b\x32\x12.productv1.ProductR\x08products\x12\x14\n\x05total\x18\x02 \x01(\x03R\x05total\"#\n\x11GetProductRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x88\x01\n\x14UpdateProductRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12 \n\x0b\x64\x65scription\x18\x03 \x01(\tR\x0b\x64\x65scription\x12\x14\n\x05price\x18\x04 \x01(\x02R\x05price\x12\x14\n\x05stock\x18\x05 \x01(\x05R\x05stock\"\x17\n\x15UpdateProductResponse\"&\n\x14\x44\x65leteProductRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x17\n\x15\x44\x65leteProductResponse2\xb9\x07\n\x0eProductService\x12o\n\rCreateProduct\x12\x1f.productv1.CreateProductRequest\x1a .productv1.CreateProductResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\"\x10/api/v1/products:\x01*\x12\x66\n\x0bGetProducts\x12\x1d.productv1.GetProductsRequest\x1a\x1e.productv1.GetProductsResponse\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/api/v1/products\x12]\n\nGetProduct\x12\x1c.productv1.GetProductRequest\x1a\x12.productv1.Product\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/api/v1/products/{id}\x12t\n\rUpdateProduct\x12\x1f.productv1.UpdateProductRequest\x1a .productv1.UpdateProductResponse\" \x82\xd3\xe4\x93\x02\x1a\x1a\x15/api/v1/products/{id}:\x01*\x12q\n\rDeleteProduct\x12\x1f.productv1.DeleteProductRequest\x1a .productv1.DeleteProductResponse\"\x1d\x82\xd3\xe4\x93\x02\x17*\x15/api/v1/products/{id}\x12}\n\x0fGetProductStock\x12!.productv1.GetProductStockRequest\x1a\".productv1.GetProductStockResponse\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/api/v1/products/{id}/stock\x12\x89\x01\n\x12UpdateProductStock\x12$.productv1.UpdateProductStockRequest\x1a%.productv1.UpdateProductStockResponse\"&\x82\xd3\xe4\x93\x02 \x1a\x1b/api/v1/products/{id}/stock:\x01*\x12{\n\x10GetProductsStock\x12\".productv1.GetProductsStockRequest\x1a#.productv1.GetProductsStockResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/api/v1/products/stockB3Z1github.com/ydssx/morphix/api/product/v1;productv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.product.v1.product_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z1github.com/ydssx/morphix/api/product/v1;productv1'
  _globals['_GETPRODUCTSSTOCKRESPONSE_STOCKSENTRY']._loaded_options = None
  _globals['_GETPRODUCTSSTOCKRESPONSE_STOCKSENTRY']._serialized_options = b'8\001'
  _globals['_PRODUCTSERVICE'].methods_by_name['CreateProduct']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['CreateProduct']._serialized_options = b'\202\323\344\223\002\025\"\020/api/v1/products:\001*'
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProducts']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProducts']._serialized_options = b'\202\323\344\223\002\022\022\020/api/v1/products'
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProduct']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProduct']._serialized_options = b'\202\323\344\223\002\027\022\025/api/v1/products/{id}'
  _globals['_PRODUCTSERVICE'].methods_by_name['UpdateProduct']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['UpdateProduct']._serialized_options = b'\202\323\344\223\002\032\032\025/api/v1/products/{id}:\001*'
  _globals['_PRODUCTSERVICE'].methods_by_name['DeleteProduct']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['DeleteProduct']._serialized_options = b'\202\323\344\223\002\027*\025/api/v1/products/{id}'
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProductStock']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProductStock']._serialized_options = b'\202\323\344\223\002\035\022\033/api/v1/products/{id}/stock'
  _globals['_PRODUCTSERVICE'].methods_by_name['UpdateProductStock']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['UpdateProductStock']._serialized_options = b'\202\323\344\223\002 \032\033/api/v1/products/{id}/stock:\001*'
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProductsStock']._loaded_options = None
  _globals['_PRODUCTSERVICE'].methods_by_name['GetProductsStock']._serialized_options = b'\202\323\344\223\002\030\022\026/api/v1/products/stock'
  _globals['_GETPRODUCTSSTOCKREQUEST']._serialized_start=73
  _globals['_GETPRODUCTSSTOCKREQUEST']._serialized_end=116
  _globals['_GETPRODUCTSSTOCKRESPONSE']._serialized_start=119
  _globals['_GETPRODUCTSSTOCKRESPONSE']._serialized_end=277
  _globals['_GETPRODUCTSSTOCKRESPONSE_STOCKSENTRY']._serialized_start=220
  _globals['_GETPRODUCTSSTOCKRESPONSE_STOCKSENTRY']._serialized_end=277
  _globals['_GETPRODUCTSTOCKREQUEST']._serialized_start=279
  _globals['_GETPRODUCTSTOCKREQUEST']._serialized_end=319
  _globals['_GETPRODUCTSTOCKRESPONSE']._serialized_start=321
  _globals['_GETPRODUCTSTOCKRESPONSE']._serialized_end=368
  _globals['_UPDATEPRODUCTSTOCKREQUEST']._serialized_start=370
  _globals['_UPDATEPRODUCTSTOCKREQUEST']._serialized_end=435
  _globals['_UPDATEPRODUCTSTOCKRESPONSE']._serialized_start=437
  _globals['_UPDATEPRODUCTSTOCKRESPONSE']._serialized_end=465
  _globals['_PRODUCT']._serialized_start=467
  _globals['_PRODUCT']._serialized_end=590
  _globals['_CREATEPRODUCTREQUEST']._serialized_start=592
  _globals['_CREATEPRODUCTREQUEST']._serialized_end=712
  _globals['_CREATEPRODUCTRESPONSE']._serialized_start=714
  _globals['_CREATEPRODUCTRESPONSE']._serialized_end=753
  _globals['_GETPRODUCTSREQUEST']._serialized_start=755
  _globals['_GETPRODUCTSREQUEST']._serialized_end=877
  _globals['_GETPRODUCTSRESPONSE']._serialized_start=879
  _globals['_GETPRODUCTSRESPONSE']._serialized_end=970
  _globals['_GETPRODUCTREQUEST']._serialized_start=972
  _globals['_GETPRODUCTREQUEST']._serialized_end=1007
  _globals['_UPDATEPRODUCTREQUEST']._serialized_start=1010
  _globals['_UPDATEPRODUCTREQUEST']._serialized_end=1146
  _globals['_UPDATEPRODUCTRESPONSE']._serialized_start=1148
  _globals['_UPDATEPRODUCTRESPONSE']._serialized_end=1171
  _globals['_DELETEPRODUCTREQUEST']._serialized_start=1173
  _globals['_DELETEPRODUCTREQUEST']._serialized_end=1211
  _globals['_DELETEPRODUCTRESPONSE']._serialized_start=1213
  _globals['_DELETEPRODUCTRESPONSE']._serialized_end=1236
  _globals['_PRODUCTSERVICE']._serialized_start=1239
  _globals['_PRODUCTSERVICE']._serialized_end=2192
# @@protoc_insertion_point(module_scope)
