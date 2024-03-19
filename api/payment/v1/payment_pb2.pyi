"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

class _PaymentMethod:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _PaymentMethodEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_PaymentMethod.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    UNKNOWN: _PaymentMethod.ValueType  # 0
    ALIPAY: _PaymentMethod.ValueType  # 1
    WECHAT: _PaymentMethod.ValueType  # 2
    PAYPAL: _PaymentMethod.ValueType  # 3

class PaymentMethod(_PaymentMethod, metaclass=_PaymentMethodEnumTypeWrapper):
    """支付方式"""

UNKNOWN: PaymentMethod.ValueType  # 0
ALIPAY: PaymentMethod.ValueType  # 1
WECHAT: PaymentMethod.ValueType  # 2
PAYPAL: PaymentMethod.ValueType  # 3
global___PaymentMethod = PaymentMethod

class _PaymentStatus:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _PaymentStatusEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_PaymentStatus.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    PENDING: _PaymentStatus.ValueType  # 0
    SUCCESS: _PaymentStatus.ValueType  # 1
    FAILED: _PaymentStatus.ValueType  # 2

class PaymentStatus(_PaymentStatus, metaclass=_PaymentStatusEnumTypeWrapper):
    """支付状态"""

PENDING: PaymentStatus.ValueType  # 0
SUCCESS: PaymentStatus.ValueType  # 1
FAILED: PaymentStatus.ValueType  # 2
global___PaymentStatus = PaymentStatus

@typing_extensions.final
class MakePaymentRequest(google.protobuf.message.Message):
    """支付请求"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    AMOUNT_FIELD_NUMBER: builtins.int
    CURRENCY_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    amount: builtins.float
    currency: builtins.str
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
        amount: builtins.float = ...,
        currency: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["amount", b"amount", "currency", b"currency", "order_number", b"order_number"]) -> None: ...

global___MakePaymentRequest = MakePaymentRequest

@typing_extensions.final
class PaymentResponse(google.protobuf.message.Message):
    """支付响应"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    PAYMENT_URL_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    status: builtins.str
    payment_url: builtins.str
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
        status: builtins.str = ...,
        payment_url: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["order_number", b"order_number", "payment_url", b"payment_url", "status", b"status"]) -> None: ...

global___PaymentResponse = PaymentResponse

@typing_extensions.final
class GetPaymentRequest(google.protobuf.message.Message):
    """查询支付请求"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["order_number", b"order_number"]) -> None: ...

global___GetPaymentRequest = GetPaymentRequest

@typing_extensions.final
class GetPaymentResponse(google.protobuf.message.Message):
    """查询支付响应"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    AMOUNT_FIELD_NUMBER: builtins.int
    CURRENCY_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    amount: builtins.float
    currency: builtins.str
    status: global___PaymentStatus.ValueType
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
        amount: builtins.float = ...,
        currency: builtins.str = ...,
        status: global___PaymentStatus.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["amount", b"amount", "currency", b"currency", "order_number", b"order_number", "status", b"status"]) -> None: ...

global___GetPaymentResponse = GetPaymentResponse

@typing_extensions.final
class CancelPaymentRequest(google.protobuf.message.Message):
    """取消支付请求"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    order_number: builtins.int
    def __init__(
        self,
        *,
        order_number: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["order_number", b"order_number"]) -> None: ...

global___CancelPaymentRequest = CancelPaymentRequest

@typing_extensions.final
class CancelPaymentResponse(google.protobuf.message.Message):
    """取消支付响应"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    STATUS_FIELD_NUMBER: builtins.int
    status: builtins.str
    """其他字段..."""
    def __init__(
        self,
        *,
        status: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["status", b"status"]) -> None: ...

global___CancelPaymentResponse = CancelPaymentResponse

@typing_extensions.final
class RefundRequest(google.protobuf.message.Message):
    """退款请求"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    AMOUNT_FIELD_NUMBER: builtins.int
    CURRENCY_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    amount: builtins.float
    currency: builtins.str
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
        amount: builtins.float = ...,
        currency: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["amount", b"amount", "currency", b"currency", "order_number", b"order_number"]) -> None: ...

global___RefundRequest = RefundRequest

@typing_extensions.final
class RefundResponse(google.protobuf.message.Message):
    """退款响应"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_NUMBER_FIELD_NUMBER: builtins.int
    STATUS_FIELD_NUMBER: builtins.int
    order_number: builtins.str
    status: builtins.str
    def __init__(
        self,
        *,
        order_number: builtins.str = ...,
        status: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["order_number", b"order_number", "status", b"status"]) -> None: ...

global___RefundResponse = RefundResponse