package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// CreateUserMembership 用户会员关系表-创建一条数据
func (a *AdminV1UserMembershipService) CreateUserMembership(ctx context.Context, req *pb.CreateUserMembershipReq) (*pb.CreateUserMembershipReply, error) {
	resp := &pb.CreateUserMembershipReply{}
	data := a.userMembershipRepo.NewData()
	data.UserID = req.GetUserId()
	data.MembershipType = req.GetMembershipType()
	data.ExpiredAt = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetExpiredAt()).StdTime())
	data.AutoRenew = req.GetAutoRenew()
	data.AutoRenewDays = req.GetAutoRenewDays()
	data.Status = req.GetStatus()
	err := a.userMembershipRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
