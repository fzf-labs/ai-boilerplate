package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetUserMembershipInfo 获取用户会员基础信息
func (s *AppV1MembershipService) GetUserMembershipInfo(ctx context.Context, req *pb.GetUserMembershipInfoReq) (*pb.GetUserMembershipInfoReply, error) {
	// 从上下文获取用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)

	// 查询用户会员信息
	userMembership, err := s.userMembershipRepo.FindOneCacheByUserID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 判断当前实际的会员类型（过期后降为普通会员）
	currentMembershipType := userMembership.MembershipType
	isExpired := false

	// 如果不是普通会员，检查是否已过期
	if userMembership.MembershipType != constant.MembershipTypeNormal.String() {
		if userMembership.ExpiredAt.Valid && userMembership.ExpiredAt.Time.Before(time.Now()) {
			// 会员已过期，降级为普通会员
			currentMembershipType = constant.MembershipTypeNormal.String()
			isExpired = true
		}
	}

	// 查询当前实际会员类型信息
	membership, err := s.membershipRepo.FindOneCacheByType(ctx, currentMembershipType)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 构建响应
	reply := &pb.GetUserMembershipInfoReply{
		MembershipType:        currentMembershipType,
		MembershipName:        membership.Name,
		MembershipDescription: membership.Description,
		Status:                userMembership.Status,
		AutoRenew:             userMembership.AutoRenew,
		AutoRenewDays:         userMembership.AutoRenewDays,
		CreatedAt:             userMembership.CreatedAt.Format(time.RFC3339),
		IsExpired:             isExpired,
	}

	// 处理到期时间（仅在非普通会员时显示原到期时间）
	if userMembership.ExpiredAt.Valid {
		reply.ExpiredAt = userMembership.ExpiredAt.Time.Format(time.RFC3339)
	}

	return reply, nil
}
