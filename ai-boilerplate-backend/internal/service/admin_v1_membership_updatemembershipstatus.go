package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembershipStatus 会员类型配置表-更新状态
func (a *AdminV1MembershipService) UpdateMembershipStatus(ctx context.Context, req *pb.UpdateMembershipStatusReq) (*pb.UpdateMembershipStatusReply, error) {
	resp := &pb.UpdateMembershipStatusReply{}
	data, err := a.membershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.membershipRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
