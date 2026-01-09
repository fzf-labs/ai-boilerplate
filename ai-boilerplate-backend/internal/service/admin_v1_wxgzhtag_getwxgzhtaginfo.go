package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhTagInfo 公众号标签表-单条数据查询
func (a *AdminV1WxGzhTagService) GetWxGzhTagInfo(ctx context.Context, req *pb.GetWxGzhTagInfoReq) (*pb.GetWxGzhTagInfoReply, error) {
	resp := &pb.GetWxGzhTagInfoReply{}
	data, err := a.wxGzhTagRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxGzhTagInfo{
		Id:        data.ID,
		AppId:     data.AppID,
		TagId:     data.TagID,
		Name:      data.Name,
		Count:     data.Count,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
