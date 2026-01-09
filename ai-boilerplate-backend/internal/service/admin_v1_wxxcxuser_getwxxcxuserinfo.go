package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxXcxUserInfo 小程序用户表-单条数据查询
func (a *AdminV1WxXcxUserService) GetWxXcxUserInfo(ctx context.Context, req *pb.GetWxXcxUserInfoReq) (*pb.GetWxXcxUserInfoReply, error) {
	resp := &pb.GetWxXcxUserInfoReply{}
	data, err := a.wxXcxUserRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxXcxUserInfo{
		Id:        data.ID,
		AppId:     data.AppID,
		Openid:    data.Openid,
		Unionid:   data.Unionid,
		Nickname:  data.Nickname,
		AvatarURL: data.AvatarURL,
		Language:  data.Language,
		Country:   data.Country,
		Province:  data.Province,
		City:      data.City,
		Remark:    data.Remark,
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
