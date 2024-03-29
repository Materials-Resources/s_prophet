// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: receiving/v1alpha0/receiving.proto

package receiving_v1alpha0

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

type GetReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReceiptRequest) Reset() {
	*x = GetReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptRequest) ProtoMessage() {}

func (x *GetReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptRequest.ProtoReflect.Descriptor instead.
func (*GetReceiptRequest) Descriptor() ([]byte, []int) {
	return file_receiving_v1alpha0_receiving_proto_rawDescGZIP(), []int{0}
}

func (x *GetReceiptRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Products []*GetReceiptResponse_Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetReceiptResponse) Reset() {
	*x = GetReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptResponse) ProtoMessage() {}

func (x *GetReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptResponse.ProtoReflect.Descriptor instead.
func (*GetReceiptResponse) Descriptor() ([]byte, []int) {
	return file_receiving_v1alpha0_receiving_proto_rawDescGZIP(), []int{1}
}

func (x *GetReceiptResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetReceiptResponse) GetProducts() []*GetReceiptResponse_Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetReceiptResponse_Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sn              string                              `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty"`
	Name            string                              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	UnitsReceived   float64                             `protobuf:"fixed64,4,opt,name=units_received,json=unitsReceived,proto3" json:"units_received,omitempty"`
	UnitLabel       string                              `protobuf:"bytes,5,opt,name=unit_label,json=unitLabel,proto3" json:"unit_label,omitempty"`
	PrimaryBin      string                              `protobuf:"bytes,6,opt,name=primary_bin,json=primaryBin,proto3" json:"primary_bin,omitempty"`
	AllocatedOrders []*GetReceiptResponse_Product_Order `protobuf:"bytes,7,rep,name=allocated_orders,json=allocatedOrders,proto3" json:"allocated_orders,omitempty"`
}

func (x *GetReceiptResponse_Product) Reset() {
	*x = GetReceiptResponse_Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptResponse_Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptResponse_Product) ProtoMessage() {}

func (x *GetReceiptResponse_Product) ProtoReflect() protoreflect.Message {
	mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptResponse_Product.ProtoReflect.Descriptor instead.
func (*GetReceiptResponse_Product) Descriptor() ([]byte, []int) {
	return file_receiving_v1alpha0_receiving_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetReceiptResponse_Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetReceiptResponse_Product) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *GetReceiptResponse_Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetReceiptResponse_Product) GetUnitsReceived() float64 {
	if x != nil {
		return x.UnitsReceived
	}
	return 0
}

func (x *GetReceiptResponse_Product) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

func (x *GetReceiptResponse_Product) GetPrimaryBin() string {
	if x != nil {
		return x.PrimaryBin
	}
	return ""
}

func (x *GetReceiptResponse_Product) GetAllocatedOrders() []*GetReceiptResponse_Product_Order {
	if x != nil {
		return x.AllocatedOrders
	}
	return nil
}

type GetReceiptResponse_Product_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                 //Order ID
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                             //Order Name
	UnitsAllocated float64 `protobuf:"fixed64,3,opt,name=units_allocated,json=unitsAllocated,proto3" json:"units_allocated,omitempty"` //Quantity Allocated
	UnitLabel      string  `protobuf:"bytes,4,opt,name=unit_label,json=unitLabel,proto3" json:"unit_label,omitempty"`
}

func (x *GetReceiptResponse_Product_Order) Reset() {
	*x = GetReceiptResponse_Product_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptResponse_Product_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptResponse_Product_Order) ProtoMessage() {}

func (x *GetReceiptResponse_Product_Order) ProtoReflect() protoreflect.Message {
	mi := &file_receiving_v1alpha0_receiving_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptResponse_Product_Order.ProtoReflect.Descriptor instead.
func (*GetReceiptResponse_Product_Order) Descriptor() ([]byte, []int) {
	return file_receiving_v1alpha0_receiving_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *GetReceiptResponse_Product_Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetReceiptResponse_Product_Order) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetReceiptResponse_Product_Order) GetUnitsAllocated() float64 {
	if x != nil {
		return x.UnitsAllocated
	}
	return 0
}

func (x *GetReceiptResponse_Product_Order) GetUnitLabel() string {
	if x != nil {
		return x.UnitLabel
	}
	return ""
}

var File_receiving_v1alpha0_receiving_proto protoreflect.FileDescriptor

var file_receiving_v1alpha0_receiving_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x30, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x30, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xed, 0x03,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x1a, 0xfa, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x62, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x42, 0x69, 0x6e, 0x12, 0x5f, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x73, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0x6f, 0x0a,
	0x10, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x25, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51,
	0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_receiving_v1alpha0_receiving_proto_rawDescOnce sync.Once
	file_receiving_v1alpha0_receiving_proto_rawDescData = file_receiving_v1alpha0_receiving_proto_rawDesc
)

func file_receiving_v1alpha0_receiving_proto_rawDescGZIP() []byte {
	file_receiving_v1alpha0_receiving_proto_rawDescOnce.Do(func() {
		file_receiving_v1alpha0_receiving_proto_rawDescData = protoimpl.X.CompressGZIP(file_receiving_v1alpha0_receiving_proto_rawDescData)
	})
	return file_receiving_v1alpha0_receiving_proto_rawDescData
}

var file_receiving_v1alpha0_receiving_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_receiving_v1alpha0_receiving_proto_goTypes = []interface{}{
	(*GetReceiptRequest)(nil),                // 0: receiving.v1alpha0.GetReceiptRequest
	(*GetReceiptResponse)(nil),               // 1: receiving.v1alpha0.GetReceiptResponse
	(*GetReceiptResponse_Product)(nil),       // 2: receiving.v1alpha0.GetReceiptResponse.Product
	(*GetReceiptResponse_Product_Order)(nil), // 3: receiving.v1alpha0.GetReceiptResponse.Product.Order
}
var file_receiving_v1alpha0_receiving_proto_depIdxs = []int32{
	2, // 0: receiving.v1alpha0.GetReceiptResponse.products:type_name -> receiving.v1alpha0.GetReceiptResponse.Product
	3, // 1: receiving.v1alpha0.GetReceiptResponse.Product.allocated_orders:type_name -> receiving.v1alpha0.GetReceiptResponse.Product.Order
	0, // 2: receiving.v1alpha0.ReceivingService.GetReceipt:input_type -> receiving.v1alpha0.GetReceiptRequest
	1, // 3: receiving.v1alpha0.ReceivingService.GetReceipt:output_type -> receiving.v1alpha0.GetReceiptResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_receiving_v1alpha0_receiving_proto_init() }
func file_receiving_v1alpha0_receiving_proto_init() {
	if File_receiving_v1alpha0_receiving_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_receiving_v1alpha0_receiving_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptRequest); i {
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
		file_receiving_v1alpha0_receiving_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptResponse); i {
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
		file_receiving_v1alpha0_receiving_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptResponse_Product); i {
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
		file_receiving_v1alpha0_receiving_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptResponse_Product_Order); i {
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
			RawDescriptor: file_receiving_v1alpha0_receiving_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_receiving_v1alpha0_receiving_proto_goTypes,
		DependencyIndexes: file_receiving_v1alpha0_receiving_proto_depIdxs,
		MessageInfos:      file_receiving_v1alpha0_receiving_proto_msgTypes,
	}.Build()
	File_receiving_v1alpha0_receiving_proto = out.File
	file_receiving_v1alpha0_receiving_proto_rawDesc = nil
	file_receiving_v1alpha0_receiving_proto_goTypes = nil
	file_receiving_v1alpha0_receiving_proto_depIdxs = nil
}
