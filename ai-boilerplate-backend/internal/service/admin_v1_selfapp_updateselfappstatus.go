package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSelfAppStatus 自应用信息表-更新状态
func (a *AdminV1SelfAppService) UpdateSelfAppStatus(ctx context.Context, req *pb.UpdateSelfAppStatusReq) (*pb.UpdateSelfAppStatusReply, error) {
	resp := &pb.UpdateSelfAppStatusReply{}
	data, err := a.selfAppRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.selfAppRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.selfAppRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
