package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMallProductStatus 商品表-更新状态
func (a *AdminV1MallProductService) UpdateMallProductStatus(ctx context.Context, req *pb.UpdateMallProductStatusReq) (*pb.UpdateMallProductStatusReply, error) {
	resp := &pb.UpdateMallProductStatusReply{}
	data, err := a.mallProductRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mallProductRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.mallProductRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
