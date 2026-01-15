package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserMembershipInfo 获取用户会员基础信息
func (s *AppV1MembershipService) GetUserMembershipInfo(ctx context.Context, req *pb.GetUserMembershipInfoReq) (*pb.GetUserMembershipInfoReply, error) {
	// 从上下文获取用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	// 查询用户会员信息
	userMembership, err := s.userMembershipRepo.GetUserActualMembershipInfo(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 查询当前实际会员类型信息
	membership, err := s.membershipRepo.FindOneCacheByType(ctx, userMembership.MembershipType)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 构建响应
	reply := &pb.GetUserMembershipInfoReply{
		MembershipType:        userMembership.MembershipType,
		MembershipName:        membership.Name,
		MembershipDescription: membership.Description,
		Status:                userMembership.Status,
		AutoRenew:             userMembership.AutoRenew,
		AutoRenewDays:         userMembership.AutoRenewDays,
		CreatedAt:             userMembership.CreatedAt.Format(time.RFC3339),
		IsExpired:             userMembership.ExpiredAt.Time.Before(time.Now()),
		ExpiredAt:             timeutil.RFC3339(userMembership.ExpiredAt.Time),
	}
	return reply, nil
}
