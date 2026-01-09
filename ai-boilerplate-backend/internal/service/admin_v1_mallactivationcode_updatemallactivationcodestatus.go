package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMallActivationCodeStatus 激活码管理表-更新状态
func (a *AdminV1MallActivationCodeService) UpdateMallActivationCodeStatus(ctx context.Context, req *pb.UpdateMallActivationCodeStatusReq) (*pb.UpdateMallActivationCodeStatusReply, error) {
	resp := &pb.UpdateMallActivationCodeStatusReply{}
	data, err := a.mallActivationCodeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mallActivationCodeRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.mallActivationCodeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
