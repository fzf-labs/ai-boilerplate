package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhTagSelector 公众号标签表-选择器
func (a *AdminV1WxGzhTagService) GetWxGzhTagSelector(ctx context.Context, req *pb.GetWxGzhTagSelectorReq) (*pb.GetWxGzhTagSelectorReply, error) {
	resp := &pb.GetWxGzhTagSelectorReply{
		List: []*pb.WxGzhTagSelector{},
	}
	list, err := a.wxGzhTagRepo.FindMultiCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.WxGzhTagSelector{
			Id:    v.ID,
			TagId: v.TagID,
			Name:  v.Name,
		})
	}
	return resp, nil
}
