package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhUserInfo 公众号粉丝表-单条数据查询
func (a *AdminV1WxGzhUserService) GetWxGzhUserInfo(ctx context.Context, req *pb.GetWxGzhUserInfoReq) (*pb.GetWxGzhUserInfoReply, error) {
	resp := &pb.GetWxGzhUserInfoReply{}
	data, err := a.wxGzhUserRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxGzhUserInfo{
		Id:              data.ID,
		AppId:           data.AppID,
		Openid:          data.Openid,
		Unionid:         data.Unionid,
		SubscribeStatus: data.SubscribeStatus,
		Nickname:        data.Nickname,
		AvatarUrl:       data.AvatarURL,
		Language:        data.Language,
		Country:         data.Country,
		Province:        data.Province,
		City:            data.City,
		TagIds:          data.TagIds,
		Remark:          data.Remark,
		CreatedAt:       data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
