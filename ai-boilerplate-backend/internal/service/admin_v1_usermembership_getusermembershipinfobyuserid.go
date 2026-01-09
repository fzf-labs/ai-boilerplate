package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetUserMembershipInfoByUserId 用户会员关系表-根据用户ID查询单条数据
//
//nolint:revive // gRPC method name comes from proto.
func (a *AdminV1UserMembershipService) GetUserMembershipInfoByUserId(ctx context.Context, req *pb.GetUserMembershipInfoByUserIdReq) (*pb.GetUserMembershipInfoByUserIdReply, error) {
	resp := &pb.GetUserMembershipInfoByUserIdReply{
		Info: &pb.UserMembershipInfo{},
	}
	data, err := a.userMembershipRepo.FindOneCacheByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return resp, nil
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
