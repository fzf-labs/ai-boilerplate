package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateWxGzhMessage 公众号消息表 -更新一条数据
func (a *AdminV1WxGzhMessageService) UpdateWxGzhMessage(ctx context.Context, req *pb.UpdateWxGzhMessageReq) (*pb.UpdateWxGzhMessageReply, error) {
	resp := &pb.UpdateWxGzhMessageReply{}
	data, err := a.wxGzhMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.wxGzhMessageRepo.DeepCopy(data)
	data.AppID = req.GetAppId()
	data.MsgID = req.GetMsgId()
	data.UserID = req.GetUserId()
	data.Openid = req.GetOpenid()
	data.MessageType = req.GetMessageType()
	data.SendFrom = req.GetSendFrom()
	data.Content = req.GetContent()
	data.MediaID = req.GetMediaId()
	data.MediaURL = req.GetMediaURL()
	data.Recognition = req.GetRecognition()
	data.Format = req.GetFormat()
	data.Title = req.GetTitle()
	data.Description = req.GetDescription()
	data.ThumbMediaID = req.GetThumbMediaId()
	data.ThumbMediaURL = req.GetThumbMediaURL()
	data.URL = req.GetURL()
	data.LocationX = req.GetLocationX()
	data.LocationY = req.GetLocationY()
	data.Scale = req.GetScale()
	data.Label = req.GetLabel()
	data.Articles = req.GetArticles()
	data.MusicURL = req.GetMusicURL()
	data.HqMusicURL = req.GetHqMusicURL()
	data.Event = req.GetEvent()
	data.EventKey = req.GetEventKey()
	err = a.wxGzhMessageRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
