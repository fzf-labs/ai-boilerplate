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

	// 收集所有 benefit_key 并查询权益类型表
	benefitTypeMap := make(map[string]struct {
		Name string
		Desc string
		Icon string
	})
	if len(benefits) > 0 {
		keys := make([]string, 0, len(benefits))
		for _, b := range benefits {
			keys = append(keys, b.BenefitKey)
		}
		typeList, err := s.membershipBenefitTypeRepo.FindMultiCacheByBenefitKeys(ctx, keys)
		if err == nil {
			for _, t := range typeList {
				benefitTypeMap[t.BenefitKey] = struct {
					Name string
					Desc string
					Icon string
				}{Name: t.BenefitName, Desc: t.BenefitDesc, Icon: t.BenefitIcon}
			}
		}
	}

	// 构建响应
	reply := &pb.GetMembershipBenefitsReply{
		Benefits: make([]*pb.MembershipBenefit, 0, len(benefits)),
	}

	for _, benefit := range benefits {
		info := &pb.MembershipBenefit{
			BenefitKey:   benefit.BenefitKey,
			BenefitValue: benefit.BenefitValue,
			BenefitNum:   benefit.BenefitNum,
			Sort:         benefit.Sort,
		}
		// 从权益类型表获取名称和描述
		if typeInfo, ok := benefitTypeMap[benefit.BenefitKey]; ok {
			info.BenefitName = typeInfo.Name
			info.BenefitDesc = typeInfo.Desc
		}
		reply.Benefits = append(reply.Benefits, info)
	}

	return reply, nil
}
