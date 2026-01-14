package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// PurchaseProduct 购买商品
func (a *AppV1MallProductService) PurchaseProduct(ctx context.Context, req *pb.PurchaseProductReq) (*pb.PurchaseProductReply, error) {
	resp := &pb.PurchaseProductReply{}
	// TODO
	return resp, nil
}
