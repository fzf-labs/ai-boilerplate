package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateMembership 会员类型配置表-创建一条数据
func (a *AdminV1MembershipService) CreateMembership(ctx context.Context, req *pb.CreateMembershipReq) (*pb.CreateMembershipReply, error) {
	resp := &pb.CreateMembershipReply{}
	data := a.membershipRepo.NewData()
	data.Name = req.GetName()
	data.Type = req.GetType()
	data.Description = req.GetDescription()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err := a.membershipRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
