package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxGzhAccountSelector 公众号账号表-公众号选择器
func (a *AdminV1WxGzhAccountService) GetWxGzhAccountSelector(ctx context.Context, _ *pb.GetWxGzhAccountSelectorReq) (*pb.GetWxGzhAccountSelectorReply, error) {
	resp := &pb.GetWxGzhAccountSelectorReply{
		List: []*pb.WxGzhAccountSelector{},
	}
	list, _, err := a.wxGzhAccountRepo.FindMultiCacheByCondition(ctx, &condition.Req{})
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		resp.List = append(resp.List, &pb.WxGzhAccountSelector{
			Id:    item.ID,
			Name:  item.Name,
			AppId: item.AppID,
		})
	}
	return resp, nil
}
