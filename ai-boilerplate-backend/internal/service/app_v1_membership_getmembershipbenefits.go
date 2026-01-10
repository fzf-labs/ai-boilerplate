package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetMembershipBenefits 获取会员权益列表
func (s *AppV1MembershipService) GetMembershipBenefits(ctx context.Context, req *pb.GetMembershipBenefitsReq) (*pb.GetMembershipBenefitsReply, error) {
	membershipType := req.MembershipType

	// 如果没有传会员类型,则获取当前用户的会员类型
	if membershipType == "" {
		userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)

		// 查询用户会员信息
		userMembership, err := s.userMembershipRepo.FindOneCacheByUserID(ctx, userID)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		membershipType = userMembership.MembershipType
	}

	// 查询会员权益列表
	benefits, err := s.membershipBenefitRepo.FindMultiByMembershipType(ctx, membershipType)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 构建响应
	reply := &pb.GetMembershipBenefitsReply{
		Benefits: make([]*pb.MembershipBenefit, 0, len(benefits)),
	}

	for _, benefit := range benefits {
		reply.Benefits = append(reply.Benefits, &pb.MembershipBenefit{
			BenefitKey:   benefit.BenefitKey,
			BenefitName:  benefit.BenefitName,
			BenefitDesc:  benefit.BenefitDesc,
			BenefitValue: benefit.BenefitValue,
			BenefitNum:   benefit.BenefitNum,
			Sort:         benefit.Sort,
		})
	}

	return reply, nil
}
