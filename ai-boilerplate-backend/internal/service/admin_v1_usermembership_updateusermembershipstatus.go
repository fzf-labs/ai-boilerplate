package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateUserMembershipStatus 用户会员关系表-更新状态
func (a *AdminV1UserMembershipService) UpdateUserMembershipStatus(ctx context.Context, req *pb.UpdateUserMembershipStatusReq) (*pb.UpdateUserMembershipStatusReply, error) {
	resp := &pb.UpdateUserMembershipStatusReply{}
	data, err := a.userMembershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.userMembershipRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.userMembershipRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
