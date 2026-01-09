package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateWxGzhAutoReply 公众号消息自动回复表-更新一条数据
func (a *AdminV1WxGzhAutoReplyService) UpdateWxGzhAutoReply(ctx context.Context, req *pb.UpdateWxGzhAutoReplyReq) (*pb.UpdateWxGzhAutoReplyReply, error) {
	resp := &pb.UpdateWxGzhAutoReplyReply{}
	data, err := a.wxGzhAutoReplyRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.wxGzhAutoReplyRepo.DeepCopy(data)
	data.AppID = req.GetAppId()
	data.Type = req.GetType()
	data.RequestKeyword = req.GetRequestKeyword()
	data.RequestKeywordMatch = req.GetRequestKeywordMatch()
	data.ResponseMessageType = req.GetResponseMessageType()
	data.ResponseContent = req.GetResponseContent()
	data.ResponseMediaID = req.GetResponseMediaId()
	err = a.wxGzhAutoReplyRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
