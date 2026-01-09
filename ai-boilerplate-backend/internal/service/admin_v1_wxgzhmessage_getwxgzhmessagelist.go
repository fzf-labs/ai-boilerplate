package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxGzhMessageList 公众号消息表 -列表数据查询
func (a *AdminV1WxGzhMessageService) GetWxGzhMessageList(ctx context.Context, req *pb.GetWxGzhMessageListReq) (*pb.GetWxGzhMessageListReply, error) {
	resp := &pb.GetWxGzhMessageListReply{
		Total: 0,
		List:  []*pb.WxGzhMessageInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.wxGzhMessageRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.WxGzhMessageInfo{
				Id:            v.ID,
				AppId:         v.AppID,
				MsgId:         v.MsgID,
				UserId:        v.UserID,
				Openid:        v.Openid,
				MessageType:   v.MessageType,
				SendFrom:      v.SendFrom,
				Content:       v.Content,
				MediaId:       v.MediaID,
				MediaURL:      v.MediaURL,
				Recognition:   v.Recognition,
				Format:        v.Format,
				Title:         v.Title,
				Description:   v.Description,
				ThumbMediaId:  v.ThumbMediaID,
				ThumbMediaURL: v.ThumbMediaURL,
				URL:           v.URL,
				LocationX:     v.LocationX,
				LocationY:     v.LocationY,
				Scale:         v.Scale,
				Label:         v.Label,
				Articles:      v.Articles,
				MusicURL:      v.MusicURL,
				HqMusicURL:    v.HqMusicURL,
				Event:         v.Event,
				EventKey:      v.EventKey,
				CreatedAt:     v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:     v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
