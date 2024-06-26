# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/order/v1/order.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61pi/order/v1/order.proto\x12\x07orderv1\x1a\x1cgoogle/api/annotations.proto\"\xb3\x01\n\x05Order\x12\x19\n\x08order_id\x18\x01 \x01(\x05R\x07orderId\x12\x1f\n\x0b\x63ustomer_id\x18\x02 \x01(\x05R\ncustomerId\x12\x16\n\x06\x61mount\x18\x03 \x01(\x02R\x06\x61mount\x12,\n\x06status\x18\x04 \x01(\x0e\x32\x14.orderv1.OrderStatusR\x06status\x12(\n\x05items\x18\x05 \x03(\x0b\x32\x12.orderv1.OrderItemR\x05items\"\x7f\n\tOrderItem\x12\x1d\n\nproduct_id\x18\x01 \x01(\x05R\tproductId\x12!\n\x0cproduct_name\x18\x02 \x01(\tR\x0bproductName\x12\x1a\n\x08quantity\x18\x03 \x01(\x05R\x08quantity\x12\x14\n\x05price\x18\x04 \x01(\x02R\x05price\"_\n\x12\x43reateOrderRequest\x12\x1f\n\x0b\x63ustomer_id\x18\x01 \x01(\x03R\ncustomerId\x12(\n\x05items\x18\x02 \x03(\x0b\x32\x12.orderv1.OrderItemR\x05items\";\n\x13\x43reateOrderResponse\x12$\n\x05order\x18\x01 \x01(\x0b\x32\x0e.orderv1.OrderR\x05order\"4\n\x0fGetOrderRequest\x12!\n\x0corder_number\x18\x01 \x01(\tR\x0borderNumber\"8\n\x10GetOrderResponse\x12$\n\x05order\x18\x01 \x01(\x0b\x32\x0e.orderv1.OrderR\x05order\"k\n\x18UpdateOrderStatusRequest\x12!\n\x0corder_number\x18\x01 \x01(\tR\x0borderNumber\x12,\n\x06status\x18\x02 \x01(\x0e\x32\x14.orderv1.OrderStatusR\x06status\"A\n\x19UpdateOrderStatusResponse\x12$\n\x05order\x18\x01 \x01(\x0b\x32\x0e.orderv1.OrderR\x05order\"7\n\x12\x44\x65leteOrderRequest\x12!\n\x0corder_number\x18\x01 \x01(\tR\x0borderNumber\"\x15\n\x13\x44\x65leteOrderResponse\"\x93\x01\n\x11ListOrdersRequest\x12\x1f\n\x0b\x63ustomer_id\x18\x01 \x01(\x03R\ncustomerId\x12,\n\x06status\x18\x02 \x01(\x0e\x32\x14.orderv1.OrderStatusR\x06status\x12\x1b\n\tpage_size\x18\x03 \x01(\x05R\x08pageSize\x12\x12\n\x04page\x18\x04 \x01(\x05R\x04page\"<\n\x12ListOrdersResponse\x12&\n\x06orders\x18\x01 \x03(\x0b\x32\x0e.orderv1.OrderR\x06orders\"[\n\x0fPayOrderRequest\x12!\n\x0corder_number\x18\x01 \x01(\tR\x0borderNumber\x12%\n\x0epayment_method\x18\x02 \x01(\tR\rpaymentMethod\"3\n\x10PayOrderResponse\x12\x1f\n\x0bpayment_url\x18\x01 \x01(\tR\npaymentUrl\"O\n\x12\x43\x61ncelOrderRequest\x12!\n\x0corder_number\x18\x01 \x01(\tR\x0borderNumber\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\"\x15\n\x13\x43\x61ncelOrderResponse*k\n\x0bOrderStatus\x12\x0b\n\x07PENDING\x10\x00\x12\x0e\n\nPROCESSING\x10\x01\x12\r\n\tCOMPLETED\x10\x02\x12\x0c\n\x08\x43\x41NCELED\x10\x03\x12\n\n\x06\x46\x41ILED\x10\x04\x12\x08\n\x04PAID\x10\x05\x12\x0c\n\x08REFUNDED\x10\x06\x32\xf5\x05\n\x0cOrderService\x12\x63\n\x0b\x43reateOrder\x12\x1b.orderv1.CreateOrderRequest\x1a\x1c.orderv1.CreateOrderResponse\"\x19\x82\xd3\xe4\x93\x02\x13\"\x0e/api/v1/orders:\x01*\x12\x66\n\x08GetOrder\x12\x18.orderv1.GetOrderRequest\x1a\x19.orderv1.GetOrderResponse\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/api/v1/orders/{order_number}\x12|\n\x11UpdateOrderStatus\x12!.orderv1.UpdateOrderStatusRequest\x1a\".orderv1.UpdateOrderStatusResponse\" \x82\xd3\xe4\x93\x02\x1a\x1a\x15/api/v1/orders/status:\x01*\x12^\n\x08PayOrder\x12\x18.orderv1.PayOrderRequest\x1a\x19.orderv1.PayOrderResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x12/api/v1/orders/pay:\x01*\x12o\n\x0b\x44\x65leteOrder\x12\x1b.orderv1.DeleteOrderRequest\x1a\x1c.orderv1.DeleteOrderResponse\"%\x82\xd3\xe4\x93\x02\x1f*\x1d/api/v1/orders/{order_number}\x12]\n\nListOrders\x12\x1a.orderv1.ListOrdersRequest\x1a\x1b.orderv1.ListOrdersResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/api/v1/orders\x12j\n\x0b\x43\x61ncelOrder\x12\x1b.orderv1.CancelOrderRequest\x1a\x1c.orderv1.CancelOrderResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/api/v1/orders/cancel:\x01*B/Z-github.com/ydssx/morphix/api/order/v1;orderv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.order.v1.order_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z-github.com/ydssx/morphix/api/order/v1;orderv1'
  _globals['_ORDERSERVICE'].methods_by_name['CreateOrder']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['CreateOrder']._serialized_options = b'\202\323\344\223\002\023\"\016/api/v1/orders:\001*'
  _globals['_ORDERSERVICE'].methods_by_name['GetOrder']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['GetOrder']._serialized_options = b'\202\323\344\223\002\037\022\035/api/v1/orders/{order_number}'
  _globals['_ORDERSERVICE'].methods_by_name['UpdateOrderStatus']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['UpdateOrderStatus']._serialized_options = b'\202\323\344\223\002\032\032\025/api/v1/orders/status:\001*'
  _globals['_ORDERSERVICE'].methods_by_name['PayOrder']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['PayOrder']._serialized_options = b'\202\323\344\223\002\027\"\022/api/v1/orders/pay:\001*'
  _globals['_ORDERSERVICE'].methods_by_name['DeleteOrder']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['DeleteOrder']._serialized_options = b'\202\323\344\223\002\037*\035/api/v1/orders/{order_number}'
  _globals['_ORDERSERVICE'].methods_by_name['ListOrders']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['ListOrders']._serialized_options = b'\202\323\344\223\002\020\022\016/api/v1/orders'
  _globals['_ORDERSERVICE'].methods_by_name['CancelOrder']._loaded_options = None
  _globals['_ORDERSERVICE'].methods_by_name['CancelOrder']._serialized_options = b'\202\323\344\223\002\032\"\025/api/v1/orders/cancel:\001*'
  _globals['_ORDERSTATUS']._serialized_start=1366
  _globals['_ORDERSTATUS']._serialized_end=1473
  _globals['_ORDER']._serialized_start=68
  _globals['_ORDER']._serialized_end=247
  _globals['_ORDERITEM']._serialized_start=249
  _globals['_ORDERITEM']._serialized_end=376
  _globals['_CREATEORDERREQUEST']._serialized_start=378
  _globals['_CREATEORDERREQUEST']._serialized_end=473
  _globals['_CREATEORDERRESPONSE']._serialized_start=475
  _globals['_CREATEORDERRESPONSE']._serialized_end=534
  _globals['_GETORDERREQUEST']._serialized_start=536
  _globals['_GETORDERREQUEST']._serialized_end=588
  _globals['_GETORDERRESPONSE']._serialized_start=590
  _globals['_GETORDERRESPONSE']._serialized_end=646
  _globals['_UPDATEORDERSTATUSREQUEST']._serialized_start=648
  _globals['_UPDATEORDERSTATUSREQUEST']._serialized_end=755
  _globals['_UPDATEORDERSTATUSRESPONSE']._serialized_start=757
  _globals['_UPDATEORDERSTATUSRESPONSE']._serialized_end=822
  _globals['_DELETEORDERREQUEST']._serialized_start=824
  _globals['_DELETEORDERREQUEST']._serialized_end=879
  _globals['_DELETEORDERRESPONSE']._serialized_start=881
  _globals['_DELETEORDERRESPONSE']._serialized_end=902
  _globals['_LISTORDERSREQUEST']._serialized_start=905
  _globals['_LISTORDERSREQUEST']._serialized_end=1052
  _globals['_LISTORDERSRESPONSE']._serialized_start=1054
  _globals['_LISTORDERSRESPONSE']._serialized_end=1114
  _globals['_PAYORDERREQUEST']._serialized_start=1116
  _globals['_PAYORDERREQUEST']._serialized_end=1207
  _globals['_PAYORDERRESPONSE']._serialized_start=1209
  _globals['_PAYORDERRESPONSE']._serialized_end=1260
  _globals['_CANCELORDERREQUEST']._serialized_start=1262
  _globals['_CANCELORDERREQUEST']._serialized_end=1341
  _globals['_CANCELORDERRESPONSE']._serialized_start=1343
  _globals['_CANCELORDERRESPONSE']._serialized_end=1364
  _globals['_ORDERSERVICE']._serialized_start=1476
  _globals['_ORDERSERVICE']._serialized_end=2233
# @@protoc_insertion_point(module_scope)
