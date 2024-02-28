// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: catalog/v1alpha0/catalog.proto

package catalog_v1alpha0

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
	CatalogService_ListProduct_FullMethodName               = "/catalog.v1alpha0.CatalogService/ListProduct"
	CatalogService_GetProduct_FullMethodName                = "/catalog.v1alpha0.CatalogService/GetProduct"
	CatalogService_CreateProduct_FullMethodName             = "/catalog.v1alpha0.CatalogService/CreateProduct"
	CatalogService_UpdateProduct_FullMethodName             = "/catalog.v1alpha0.CatalogService/UpdateProduct"
	CatalogService_DeleteProduct_FullMethodName             = "/catalog.v1alpha0.CatalogService/DeleteProduct"
	CatalogService_CreateProductSupplier_FullMethodName     = "/catalog.v1alpha0.CatalogService/CreateProductSupplier"
	CatalogService_UpdateProductSupplier_FullMethodName     = "/catalog.v1alpha0.CatalogService/UpdateProductSupplier"
	CatalogService_SetPrimaryProductSupplier_FullMethodName = "/catalog.v1alpha0.CatalogService/SetPrimaryProductSupplier"
	CatalogService_ListGroup_FullMethodName                 = "/catalog.v1alpha0.CatalogService/ListGroup"
	CatalogService_GetGroup_FullMethodName                  = "/catalog.v1alpha0.CatalogService/GetGroup"
	CatalogService_CreateGroup_FullMethodName               = "/catalog.v1alpha0.CatalogService/CreateGroup"
	CatalogService_UpdateGroup_FullMethodName               = "/catalog.v1alpha0.CatalogService/UpdateGroup"
	CatalogService_GetProductBySupplier_FullMethodName      = "/catalog.v1alpha0.CatalogService/GetProductBySupplier"
)

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogServiceClient interface {
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	CreateProductSupplier(ctx context.Context, in *CreateProductSupplierRequest, opts ...grpc.CallOption) (*CreateProductSupplierResponse, error)
	UpdateProductSupplier(ctx context.Context, in *UpdateProductSupplierRequest, opts ...grpc.CallOption) (*UpdateProductSupplierResponse, error)
	SetPrimaryProductSupplier(ctx context.Context, in *SetPrimaryProductSupplierRequest, opts ...grpc.CallOption) (*SetPrimaryProductSupplierResponse, error)
	ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
	GetProductBySupplier(ctx context.Context, in *GetProductBySupplierRequest, opts ...grpc.CallOption) (*GetProductBySupplierResponse, error)
}

type catalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogServiceClient(cc grpc.ClientConnInterface) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, CatalogService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateProductSupplier(ctx context.Context, in *CreateProductSupplierRequest, opts ...grpc.CallOption) (*CreateProductSupplierResponse, error) {
	out := new(CreateProductSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateProductSupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateProductSupplier(ctx context.Context, in *UpdateProductSupplierRequest, opts ...grpc.CallOption) (*UpdateProductSupplierResponse, error) {
	out := new(UpdateProductSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateProductSupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) SetPrimaryProductSupplier(ctx context.Context, in *SetPrimaryProductSupplierRequest, opts ...grpc.CallOption) (*SetPrimaryProductSupplierResponse, error) {
	out := new(SetPrimaryProductSupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_SetPrimaryProductSupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error) {
	out := new(ListGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_ListGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	out := new(UpdateGroupResponse)
	err := c.cc.Invoke(ctx, CatalogService_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) GetProductBySupplier(ctx context.Context, in *GetProductBySupplierRequest, opts ...grpc.CallOption) (*GetProductBySupplierResponse, error) {
	out := new(GetProductBySupplierResponse)
	err := c.cc.Invoke(ctx, CatalogService_GetProductBySupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServiceServer is the server API for CatalogService service.
// All implementations should embed UnimplementedCatalogServiceServer
// for forward compatibility
type CatalogServiceServer interface {
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	CreateProductSupplier(context.Context, *CreateProductSupplierRequest) (*CreateProductSupplierResponse, error)
	UpdateProductSupplier(context.Context, *UpdateProductSupplierRequest) (*UpdateProductSupplierResponse, error)
	SetPrimaryProductSupplier(context.Context, *SetPrimaryProductSupplierRequest) (*SetPrimaryProductSupplierResponse, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
	GetProductBySupplier(context.Context, *GetProductBySupplierRequest) (*GetProductBySupplierResponse, error)
}

// UnimplementedCatalogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (UnimplementedCatalogServiceServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedCatalogServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedCatalogServiceServer) CreateProductSupplier(context.Context, *CreateProductSupplierRequest) (*CreateProductSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateProductSupplier(context.Context, *UpdateProductSupplierRequest) (*UpdateProductSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) SetPrimaryProductSupplier(context.Context, *SetPrimaryProductSupplierRequest) (*SetPrimaryProductSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryProductSupplier not implemented")
}
func (UnimplementedCatalogServiceServer) ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroup not implemented")
}
func (UnimplementedCatalogServiceServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedCatalogServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedCatalogServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedCatalogServiceServer) GetProductBySupplier(context.Context, *GetProductBySupplierRequest) (*GetProductBySupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBySupplier not implemented")
}

// UnsafeCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServiceServer will
// result in compilation errors.
type UnsafeCatalogServiceServer interface {
	mustEmbedUnimplementedCatalogServiceServer()
}

func RegisterCatalogServiceServer(s grpc.ServiceRegistrar, srv CatalogServiceServer) {
	s.RegisterService(&CatalogService_ServiceDesc, srv)
}

func _CatalogService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateProductSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateProductSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateProductSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateProductSupplier(ctx, req.(*CreateProductSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateProductSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateProductSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateProductSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateProductSupplier(ctx, req.(*UpdateProductSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_SetPrimaryProductSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrimaryProductSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).SetPrimaryProductSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_SetPrimaryProductSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).SetPrimaryProductSupplier(ctx, req.(*SetPrimaryProductSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).ListGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_ListGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).ListGroup(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_GetProductBySupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBySupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).GetProductBySupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CatalogService_GetProductBySupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).GetProductBySupplier(ctx, req.(*GetProductBySupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogService_ServiceDesc is the grpc.ServiceDesc for CatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.v1alpha0.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProduct",
			Handler:    _CatalogService_ListProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _CatalogService_GetProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _CatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _CatalogService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _CatalogService_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateProductSupplier",
			Handler:    _CatalogService_CreateProductSupplier_Handler,
		},
		{
			MethodName: "UpdateProductSupplier",
			Handler:    _CatalogService_UpdateProductSupplier_Handler,
		},
		{
			MethodName: "SetPrimaryProductSupplier",
			Handler:    _CatalogService_SetPrimaryProductSupplier_Handler,
		},
		{
			MethodName: "ListGroup",
			Handler:    _CatalogService_ListGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _CatalogService_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _CatalogService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _CatalogService_UpdateGroup_Handler,
		},
		{
			MethodName: "GetProductBySupplier",
			Handler:    _CatalogService_GetProductBySupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog/v1alpha0/catalog.proto",
}
