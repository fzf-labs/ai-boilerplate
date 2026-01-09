package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// UpdateUserMembership 用户会员关系表-更新一条数据
func (a *AdminV1UserMembershipService) UpdateUserMembership(ctx context.Context, req *pb.UpdateUserMembershipReq) (*pb.UpdateUserMembershipReply, error) {
	resp := &pb.UpdateUserMembershipReply{}
	data, err := a.userMembershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.userMembershipRepo.DeepCopy(data)
	data.UserID = req.GetUserId()
	data.MembershipType = req.GetMembershipType()
	data.ExpiredAt = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetExpiredAt()).StdTime())
	data.AutoRenew = req.GetAutoRenew()
	data.AutoRenewDays = req.GetAutoRenewDays()
	data.Status = req.GetStatus()
	err = a.userMembershipRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
