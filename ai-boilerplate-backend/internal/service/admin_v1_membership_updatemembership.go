package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMembership 会员类型配置表-更新一条数据
func (a *AdminV1MembershipService) UpdateMembership(ctx context.Context, req *pb.UpdateMembershipReq) (*pb.UpdateMembershipReply, error) {
	resp := &pb.UpdateMembershipReply{}
	data, err := a.membershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.membershipRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Type = req.GetType()
	data.Description = req.GetDescription()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.membershipRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
