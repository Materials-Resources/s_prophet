// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: receiving/v1alpha0/receiving.proto

package receiving_v1alpha0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ReceivingService_GetReceipt_FullMethodName = "/receiving.v1alpha0.ReceivingService/GetReceipt"
)

// ReceivingServiceClient is the client API for ReceivingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceivingServiceClient interface {
	GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error)
}

type receivingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReceivingServiceClient(cc grpc.ClientConnInterface) ReceivingServiceClient {
	return &receivingServiceClient{cc}
}

func (c *receivingServiceClient) GetReceipt(ctx context.Context, in *GetReceiptRequest, opts ...grpc.CallOption) (*GetReceiptResponse, error) {
	out := new(GetReceiptResponse)
	err := c.cc.Invoke(ctx, ReceivingService_GetReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceivingServiceServer is the server API for ReceivingService service.
// All implementations should embed UnimplementedReceivingServiceServer
// for forward compatibility
type ReceivingServiceServer interface {
	GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error)
}

// UnimplementedReceivingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReceivingServiceServer struct {
}

func (UnimplementedReceivingServiceServer) GetReceipt(context.Context, *GetReceiptRequest) (*GetReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}

// UnsafeReceivingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceivingServiceServer will
// result in compilation errors.
type UnsafeReceivingServiceServer interface {
	mustEmbedUnimplementedReceivingServiceServer()
}

func RegisterReceivingServiceServer(s grpc.ServiceRegistrar, srv ReceivingServiceServer) {
	s.RegisterService(&ReceivingService_ServiceDesc, srv)
}

func _ReceivingService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceivingServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReceivingService_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceivingServiceServer).GetReceipt(ctx, req.(*GetReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReceivingService_ServiceDesc is the grpc.ServiceDesc for ReceivingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReceivingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "receiving.v1alpha0.ReceivingService",
	HandlerType: (*ReceivingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReceipt",
			Handler:    _ReceivingService_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiving/v1alpha0/receiving.proto",
}