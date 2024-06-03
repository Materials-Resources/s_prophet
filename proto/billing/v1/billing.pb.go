// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: billing/v1/billing.proto

package billing_v1

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

type SimplifiedInvoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId    string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status     string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  string  `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Total      float64 `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	AmountPaid float64 `protobuf:"fixed64,6,opt,name=amount_paid,json=amountPaid,proto3" json:"amount_paid,omitempty"`
}

func (x *SimplifiedInvoice) Reset() {
	*x = SimplifiedInvoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_v1_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimplifiedInvoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimplifiedInvoice) ProtoMessage() {}

func (x *SimplifiedInvoice) ProtoReflect() protoreflect.Message {
	mi := &file_billing_v1_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimplifiedInvoice.ProtoReflect.Descriptor instead.
func (*SimplifiedInvoice) Descriptor() ([]byte, []int) {
	return file_billing_v1_billing_proto_rawDescGZIP(), []int{0}
}

func (x *SimplifiedInvoice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimplifiedInvoice) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *SimplifiedInvoice) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SimplifiedInvoice) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SimplifiedInvoice) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SimplifiedInvoice) GetAmountPaid() float64 {
	if x != nil {
		return x.AmountPaid
	}
	return 0
}

type GetInvoicesByOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetInvoicesByOrderRequest) Reset() {
	*x = GetInvoicesByOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_v1_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoicesByOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoicesByOrderRequest) ProtoMessage() {}

func (x *GetInvoicesByOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_v1_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoicesByOrderRequest.ProtoReflect.Descriptor instead.
func (*GetInvoicesByOrderRequest) Descriptor() ([]byte, []int) {
	return file_billing_v1_billing_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvoicesByOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetInvoicesByOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoices []*SimplifiedInvoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *GetInvoicesByOrderResponse) Reset() {
	*x = GetInvoicesByOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_v1_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvoicesByOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoicesByOrderResponse) ProtoMessage() {}

func (x *GetInvoicesByOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_billing_v1_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoicesByOrderResponse.ProtoReflect.Descriptor instead.
func (*GetInvoicesByOrderResponse) Descriptor() ([]byte, []int) {
	return file_billing_v1_billing_proto_rawDescGZIP(), []int{2}
}

func (x *GetInvoicesByOrderResponse) GetInvoices() []*SimplifiedInvoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

var File_billing_v1_billing_proto protoreflect.FileDescriptor

var file_billing_v1_billing_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xac, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x61, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x32, 0x75, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x25,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_billing_v1_billing_proto_rawDescOnce sync.Once
	file_billing_v1_billing_proto_rawDescData = file_billing_v1_billing_proto_rawDesc
)

func file_billing_v1_billing_proto_rawDescGZIP() []byte {
	file_billing_v1_billing_proto_rawDescOnce.Do(func() {
		file_billing_v1_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_billing_v1_billing_proto_rawDescData)
	})
	return file_billing_v1_billing_proto_rawDescData
}

var file_billing_v1_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_billing_v1_billing_proto_goTypes = []interface{}{
	(*SimplifiedInvoice)(nil),          // 0: billing.v1.SimplifiedInvoice
	(*GetInvoicesByOrderRequest)(nil),  // 1: billing.v1.GetInvoicesByOrderRequest
	(*GetInvoicesByOrderResponse)(nil), // 2: billing.v1.GetInvoicesByOrderResponse
}
var file_billing_v1_billing_proto_depIdxs = []int32{
	0, // 0: billing.v1.GetInvoicesByOrderResponse.invoices:type_name -> billing.v1.SimplifiedInvoice
	1, // 1: billing.v1.BillingService.GetInvoicesByOrder:input_type -> billing.v1.GetInvoicesByOrderRequest
	2, // 2: billing.v1.BillingService.GetInvoicesByOrder:output_type -> billing.v1.GetInvoicesByOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_billing_v1_billing_proto_init() }
func file_billing_v1_billing_proto_init() {
	if File_billing_v1_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_billing_v1_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimplifiedInvoice); i {
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
		file_billing_v1_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoicesByOrderRequest); i {
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
		file_billing_v1_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvoicesByOrderResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_billing_v1_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_billing_v1_billing_proto_goTypes,
		DependencyIndexes: file_billing_v1_billing_proto_depIdxs,
		MessageInfos:      file_billing_v1_billing_proto_msgTypes,
	}.Build()
	File_billing_v1_billing_proto = out.File
	file_billing_v1_billing_proto_rawDesc = nil
	file_billing_v1_billing_proto_goTypes = nil
	file_billing_v1_billing_proto_depIdxs = nil
}
