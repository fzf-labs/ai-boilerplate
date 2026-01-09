package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSelfAppReleaseStatus 自应用版本发布表-更新状态
func (a *AdminV1SelfAppReleaseService) UpdateSelfAppReleaseStatus(ctx context.Context, req *pb.UpdateSelfAppReleaseStatusReq) (*pb.UpdateSelfAppReleaseStatusReply, error) {
	resp := &pb.UpdateSelfAppReleaseStatusReply{}
	data, err := a.selfAppReleaseRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.selfAppReleaseRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.selfAppReleaseRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
