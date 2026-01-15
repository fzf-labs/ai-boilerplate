package service

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// benefitValue 权益值
type benefitValue struct {
	Value string
	Num   string
}

// GetMembershipBenefitsCompare 获取会员权益对比（一次返回所有会员类型的权益）
func (s *AppV1MembershipService) GetMembershipBenefitsCompare(ctx context.Context, req *pb.GetMembershipBenefitsCompareReq) (*pb.GetMembershipBenefitsCompareReply, error) {
	membershipTypes := []string{"normal", "vip", "svip"}

	// 一次性获取所有会员类型的权益
	allBenefits, err := s.membershipBenefitRepo.FindMultiCacheByMembershipTypes(ctx, membershipTypes)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 按会员类型分组权益数据，同时收集所有权益 key
	benefitsByType := map[string]map[string]benefitValue{
		"normal": {},
		"vip":    {},
		"svip":   {},
	}
	benefitKeySet := make(map[string]struct{})

	for _, b := range allBenefits {
		benefitsByType[b.MembershipType][b.BenefitKey] = benefitValue{
			Value: b.BenefitValue,
			Num:   b.BenefitNum,
		}
		benefitKeySet[b.BenefitKey] = struct{}{}
	}

	// 提取所有权益 key
	allBenefitKeys := make([]string, 0, len(benefitKeySet))
	for key := range benefitKeySet {
		allBenefitKeys = append(allBenefitKeys, key)
	}

	if len(allBenefitKeys) == 0 {
		return &pb.GetMembershipBenefitsCompareReply{Items: nil}, nil
	}

	// 获取权益类型信息（名称、描述、排序）
	typeList, err := s.membershipBenefitTypeRepo.FindMultiCacheByBenefitKeys(ctx, allBenefitKeys)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 构建对比列表
	items := make([]*pb.MembershipBenefitCompareItem, 0, len(typeList))
	for _, t := range typeList {
		items = append(items, &pb.MembershipBenefitCompareItem{
			BenefitKey:  t.BenefitKey,
			BenefitName: t.BenefitName,
			BenefitDesc: t.BenefitDesc,
			Sort:        t.Sort,
			Normal:      toBenefitValue(benefitsByType["normal"][t.BenefitKey]),
			Vip:         toBenefitValue(benefitsByType["vip"][t.BenefitKey]),
			Svip:        toBenefitValue(benefitsByType["svip"][t.BenefitKey]),
		})
	}

	// 按 Sort 字段排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Sort < items[j].Sort
	})

	return &pb.GetMembershipBenefitsCompareReply{Items: items}, nil
}

// toBenefitValue 转换权益值
func toBenefitValue(bv benefitValue) *pb.MembershipBenefitValue {
	if bv.Value == "" && bv.Num == "" {
		return &pb.MembershipBenefitValue{Supported: false}
	}
	return &pb.MembershipBenefitValue{
		Supported: true,
		Value:     bv.Value,
		Num:       bv.Num,
	}
}
