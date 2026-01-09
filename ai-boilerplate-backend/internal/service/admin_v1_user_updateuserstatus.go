package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateUserStatus 用户表-更新状态
func (a *AdminV1UserService) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusReq) (*pb.UpdateUserStatusReply, error) {
	resp := &pb.UpdateUserStatusReply{}
	data, err := a.userRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.userRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.userRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
