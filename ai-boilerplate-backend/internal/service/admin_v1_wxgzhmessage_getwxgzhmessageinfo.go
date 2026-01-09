package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhMessageInfo 公众号消息表 -单条数据查询
func (a *AdminV1WxGzhMessageService) GetWxGzhMessageInfo(ctx context.Context, req *pb.GetWxGzhMessageInfoReq) (*pb.GetWxGzhMessageInfoReply, error) {
	resp := &pb.GetWxGzhMessageInfoReply{}
	data, err := a.wxGzhMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxGzhMessageInfo{
		Id:            data.ID,
		AppId:         data.AppID,
		MsgId:         data.MsgID,
		UserId:        data.UserID,
		Openid:        data.Openid,
		MessageType:   data.MessageType,
		SendFrom:      data.SendFrom,
		Content:       data.Content,
		MediaId:       data.MediaID,
		MediaURL:      data.MediaURL,
		Recognition:   data.Recognition,
		Format:        data.Format,
		Title:         data.Title,
		Description:   data.Description,
		ThumbMediaId:  data.ThumbMediaID,
		ThumbMediaURL: data.ThumbMediaURL,
		URL:           data.URL,
		LocationX:     data.LocationX,
		LocationY:     data.LocationY,
		Scale:         data.Scale,
		Label:         data.Label,
		Articles:      data.Articles,
		MusicURL:      data.MusicURL,
		HqMusicURL:    data.HqMusicURL,
		Event:         data.Event,
		EventKey:      data.EventKey,
		CreatedAt:     data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
