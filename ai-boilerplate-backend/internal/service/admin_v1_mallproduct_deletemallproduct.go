package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMallProduct 商品表-删除一条数据
func (a *AdminV1MallProductService) DeleteMallProduct(ctx context.Context, req *pb.DeleteMallProductReq) (*pb.DeleteMallProductReply, error) {
	resp := &pb.DeleteMallProductReply{}
	err := a.mallProductRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
