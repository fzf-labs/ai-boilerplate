package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateWxGzhAutoReply 公众号消息自动回复表-创建一条数据
func (a *AdminV1WxGzhAutoReplyService) CreateWxGzhAutoReply(ctx context.Context, req *pb.CreateWxGzhAutoReplyReq) (*pb.CreateWxGzhAutoReplyReply, error) {
	resp := &pb.CreateWxGzhAutoReplyReply{}
	data := a.wxGzhAutoReplyRepo.NewData()
	data.AppID = req.GetAppId()
	data.Type = req.GetType()
	data.RequestKeyword = req.GetRequestKeyword()
	data.RequestKeywordMatch = req.GetRequestKeywordMatch()
	data.ResponseMessageType = req.GetResponseMessageType()
	data.ResponseContent = req.GetResponseContent()
	data.ResponseMediaID = req.GetResponseMediaId()
	data.Status = req.GetStatus() // 默认开启
	err := a.wxGzhAutoReplyRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
