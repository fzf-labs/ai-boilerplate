package service

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSmsChannelSelector 短信渠道-选择器
func (a *AdminV1SmsChannelService) GetSmsChannelSelector(ctx context.Context, _ *pb.GetSmsChannelSelectorReq) (*pb.GetSmsChannelSelectorReply, error) {
	resp := &pb.GetSmsChannelSelectorReply{
		List: []*pb.SmsChannelSelector{},
	}
	list, _, err := a.smsChannelRepo.FindMultiByCondition(ctx, &condition.Req{})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 时间排序 创建时间倒序
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.Before(list[j].CreatedAt)
	})
	for _, v := range list {
		resp.List = append(resp.List, &pb.SmsChannelSelector{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return resp, nil
}
