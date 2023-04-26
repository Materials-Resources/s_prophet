package grpc_service

import (
	"context"
	"github.com/materials-resources/ProphetService/pkg/database"
	rpc "github.com/materials-resources/ProphetService/proto/product/v1alpha0"
)

type ProductServer struct {
	rpc.UnimplementedProductServiceServer
	ProductHandler *database.ProductHandler
}

func (s *ProductServer) GetProduct(ctx context.Context, req *rpc.GetProductRequest) (*rpc.GetProductResponse, error) {
	hRes, err := s.ProductHandler.SelectProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &rpc.GetProductResponse{
		Id:             hRes.InvMastUid,
		ProductSn:      hRes.ItemId,
		Name:           hRes.ItemDesc,
		Description:    hRes.ExtendedDesc,
		ProductGroupId: hRes.DefaultProductGroup,
		Price:          hRes.Price1,
	}, nil
}

func (s *ProductServer) GetProductGroups(ctx context.Context, req *rpc.GetProductGroupsRequest) (*rpc.GetProductGroupsResponse, error) {
	grpcResponse := new(rpc.GetProductGroupsResponse)

	qr, err := s.ProductHandler.SelectProductGroups(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range *qr {
		grpcResponse.ProductGroups = append(grpcResponse.ProductGroups, &rpc.ProductGroup{
			Id:          r.ProductGroupId,
			Description: r.ProductGroupDesc,
		})

	}

	return grpcResponse, nil
}
func (s *ProductServer) CreateProductGroup(ctx context.Context, req *rpc.CreateProductGroupRequest) (*rpc.CreateProductGroupResponse, error) {

	return nil, nil
}
