package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateWxGzhMessage 公众号消息表 -创建一条数据
func (a *AdminV1WxGzhMessageService) CreateWxGzhMessage(ctx context.Context, req *pb.CreateWxGzhMessageReq) (*pb.CreateWxGzhMessageReply, error) {
	resp := &pb.CreateWxGzhMessageReply{}
	data := a.wxGzhMessageRepo.NewData()
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
	err := a.wxGzhMessageRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
