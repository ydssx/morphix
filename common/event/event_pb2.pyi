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

class _Subject:
    ValueType = typing.NewType("ValueType", builtins.int)
    V: typing_extensions.TypeAlias = ValueType

class _SubjectEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[_Subject.ValueType], builtins.type):
    DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
    PaymentCompleted: _Subject.ValueType  # 0
    """支付完成"""
    OrderCreated: _Subject.ValueType  # 1
    """创建订单"""
    CancelPayment: _Subject.ValueType  # 2
    """取消支付"""

class Subject(_Subject, metaclass=_SubjectEnumTypeWrapper): ...

PaymentCompleted: Subject.ValueType  # 0
"""支付完成"""
OrderCreated: Subject.ValueType  # 1
"""创建订单"""
CancelPayment: Subject.ValueType  # 2
"""取消支付"""
global___Subject = Subject

@typing_extensions.final
class Event(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    SUBJECT_FIELD_NUMBER: builtins.int
    PAYMENT_COMPLETED_FIELD_NUMBER: builtins.int
    CANCEL_PAYMENT_FIELD_NUMBER: builtins.int
    subject: global___Subject.ValueType
    @property
    def payment_completed(self) -> global___PayloadPaymentCompleted: ...
    @property
    def cancel_payment(self) -> global___PayloadCancelPayment: ...
    def __init__(
        self,
        *,
        subject: global___Subject.ValueType = ...,
        payment_completed: global___PayloadPaymentCompleted | None = ...,
        cancel_payment: global___PayloadCancelPayment | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["cancel_payment", b"cancel_payment", "payload", b"payload", "payment_completed", b"payment_completed"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["cancel_payment", b"cancel_payment", "payload", b"payload", "payment_completed", b"payment_completed", "subject", b"subject"]) -> None: ...
    def WhichOneof(self, oneof_group: typing_extensions.Literal["payload", b"payload"]) -> typing_extensions.Literal["payment_completed", "cancel_payment"] | None: ...

global___Event = Event

@typing_extensions.final
class PayloadPaymentCompleted(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    AMOUNT_FIELD_NUMBER: builtins.int
    ORDER_ID_FIELD_NUMBER: builtins.int
    SUBJECT_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    amount: builtins.float
    order_id: builtins.int
    subject: global___Subject.ValueType
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
        amount: builtins.float = ...,
        order_id: builtins.int = ...,
        subject: global___Subject.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["amount", b"amount", "order_id", b"order_id", "subject", b"subject", "user_id", b"user_id"]) -> None: ...

global___PayloadPaymentCompleted = PayloadPaymentCompleted

@typing_extensions.final
class PayloadCancelPayment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ORDER_ID_FIELD_NUMBER: builtins.int
    order_id: builtins.int
    def __init__(
        self,
        *,
        order_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["order_id", b"order_id"]) -> None: ...

global___PayloadCancelPayment = PayloadCancelPayment