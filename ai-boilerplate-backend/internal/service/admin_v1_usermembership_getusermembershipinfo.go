package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetUserMembershipInfo 用户会员关系表-单条数据查询
func (a *AdminV1UserMembershipService) GetUserMembershipInfo(ctx context.Context, req *pb.GetUserMembershipInfoReq) (*pb.GetUserMembershipInfoReply, error) {
	resp := &pb.GetUserMembershipInfoReply{}
	data, err := a.userMembershipRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.UserMembershipInfo{
		Id:             data.ID,
		UserId:         data.UserID,
		MembershipType: data.MembershipType,
		ExpiredAt:      data.ExpiredAt.Time.Format(time.RFC3339),
		AutoRenew:      data.AutoRenew,
		AutoRenewDays:  data.AutoRenewDays,
		Status:         data.Status,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
