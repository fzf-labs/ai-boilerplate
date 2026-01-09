package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhAutoReplyInfo 公众号消息自动回复表-单条数据查询
func (a *AdminV1WxGzhAutoReplyService) GetWxGzhAutoReplyInfo(ctx context.Context, req *pb.GetWxGzhAutoReplyInfoReq) (*pb.GetWxGzhAutoReplyInfoReply, error) {
	resp := &pb.GetWxGzhAutoReplyInfoReply{}
	data, err := a.wxGzhAutoReplyRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxGzhAutoReplyInfo{
		Id:                  data.ID,
		AppId:               data.AppID,
		Type:                data.Type,
		RequestKeyword:      data.RequestKeyword,
		RequestKeywordMatch: data.RequestKeywordMatch,
		ResponseMessageType: data.ResponseMessageType,
		ResponseContent:     data.ResponseContent,
		ResponseMediaId:     data.ResponseMediaID,
		Status:              data.Status,
		CreatedAt:           data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
