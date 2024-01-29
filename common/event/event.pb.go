// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: common/event/event.proto

package event

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subject int32

const (
	Subject_PaymentCompleted Subject = 0 //支付完成
	Subject_OrderCreated     Subject = 1 //创建订单
	Subject_CancelPayment    Subject = 2 //取消支付
)

// Enum value maps for Subject.
var (
	Subject_name = map[int32]string{
		0: "PaymentCompleted",
		1: "OrderCreated",
		2: "CancelPayment",
	}
	Subject_value = map[string]int32{
		"PaymentCompleted": 0,
		"OrderCreated":     1,
		"CancelPayment":    2,
	}
)

func (x Subject) Enum() *Subject {
	p := new(Subject)
	*p = x
	return p
}

func (x Subject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Subject) Descriptor() protoreflect.EnumDescriptor {
	return file_common_event_event_proto_enumTypes[0].Descriptor()
}

func (Subject) Type() protoreflect.EnumType {
	return &file_common_event_event_proto_enumTypes[0]
}

func (x Subject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Subject.Descriptor instead.
func (Subject) EnumDescriptor() ([]byte, []int) {
	return file_common_event_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject Subject `protobuf:"varint,1,opt,name=subject,proto3,enum=event.Subject" json:"subject,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Event_PaymentCompleted
	//	*Event_CancelPayment
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_common_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_common_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetSubject() Subject {
	if x != nil {
		return x.Subject
	}
	return Subject_PaymentCompleted
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetPaymentCompleted() *PayloadPaymentCompleted {
	if x, ok := x.GetPayload().(*Event_PaymentCompleted); ok {
		return x.PaymentCompleted
	}
	return nil
}

func (x *Event) GetCancelPayment() *PayloadCancelPayment {
	if x, ok := x.GetPayload().(*Event_CancelPayment); ok {
		return x.CancelPayment
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_PaymentCompleted struct {
	PaymentCompleted *PayloadPaymentCompleted `protobuf:"bytes,2,opt,name=payment_completed,json=paymentCompleted,proto3,oneof"`
}

type Event_CancelPayment struct {
	CancelPayment *PayloadCancelPayment `protobuf:"bytes,3,opt,name=cancel_payment,json=cancelPayment,proto3,oneof"`
}

func (*Event_PaymentCompleted) isEvent_Payload() {}

func (*Event_CancelPayment) isEvent_Payload() {}

type PayloadPaymentCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount  float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	OrderId int64   `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Subject Subject `protobuf:"varint,4,opt,name=subject,proto3,enum=event.Subject" json:"subject,omitempty"`
}

func (x *PayloadPaymentCompleted) Reset() {
	*x = PayloadPaymentCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadPaymentCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadPaymentCompleted) ProtoMessage() {}

func (x *PayloadPaymentCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_common_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadPaymentCompleted.ProtoReflect.Descriptor instead.
func (*PayloadPaymentCompleted) Descriptor() ([]byte, []int) {
	return file_common_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadPaymentCompleted) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PayloadPaymentCompleted) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayloadPaymentCompleted) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PayloadPaymentCompleted) GetSubject() Subject {
	if x != nil {
		return x.Subject
	}
	return Subject_PaymentCompleted
}

type PayloadCancelPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PayloadCancelPayment) Reset() {
	*x = PayloadCancelPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadCancelPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadCancelPayment) ProtoMessage() {}

func (x *PayloadCancelPayment) ProtoReflect() protoreflect.Message {
	mi := &file_common_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadCancelPayment.ProtoReflect.Descriptor instead.
func (*PayloadCancelPayment) Descriptor() ([]byte, []int) {
	return file_common_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *PayloadCancelPayment) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_common_event_event_proto protoreflect.FileDescriptor

var file_common_event_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4d, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x31, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x44, 0x0a, 0x07, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x64, 0x73, 0x73, 0x78, 0x2f, 0x6d, 0x6f, 0x72, 0x70, 0x68, 0x69, 0x78, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_event_event_proto_rawDescOnce sync.Once
	file_common_event_event_proto_rawDescData = file_common_event_event_proto_rawDesc
)

func file_common_event_event_proto_rawDescGZIP() []byte {
	file_common_event_event_proto_rawDescOnce.Do(func() {
		file_common_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_event_event_proto_rawDescData)
	})
	return file_common_event_event_proto_rawDescData
}

var file_common_event_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_event_event_proto_goTypes = []interface{}{
	(Subject)(0),                    // 0: event.Subject
	(*Event)(nil),                   // 1: event.Event
	(*PayloadPaymentCompleted)(nil), // 2: event.PayloadPaymentCompleted
	(*PayloadCancelPayment)(nil),    // 3: event.PayloadCancelPayment
}
var file_common_event_event_proto_depIdxs = []int32{
	0, // 0: event.Event.subject:type_name -> event.Subject
	2, // 1: event.Event.payment_completed:type_name -> event.PayloadPaymentCompleted
	3, // 2: event.Event.cancel_payment:type_name -> event.PayloadCancelPayment
	0, // 3: event.PayloadPaymentCompleted.subject:type_name -> event.Subject
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_event_event_proto_init() }
func file_common_event_event_proto_init() {
	if File_common_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadPaymentCompleted); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadCancelPayment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_event_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_PaymentCompleted)(nil),
		(*Event_CancelPayment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_event_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_event_event_proto_goTypes,
		DependencyIndexes: file_common_event_event_proto_depIdxs,
		EnumInfos:         file_common_event_event_proto_enumTypes,
		MessageInfos:      file_common_event_event_proto_msgTypes,
	}.Build()
	File_common_event_event_proto = out.File
	file_common_event_event_proto_rawDesc = nil
	file_common_event_event_proto_goTypes = nil
	file_common_event_event_proto_depIdxs = nil
}
