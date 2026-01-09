package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhAccountInfo 公众号账号表-单条数据查询
func (a *AdminV1WxGzhAccountService) GetWxGzhAccountInfo(ctx context.Context, req *pb.GetWxGzhAccountInfoReq) (*pb.GetWxGzhAccountInfoReply, error) {
	resp := &pb.GetWxGzhAccountInfoReply{}
	data, err := a.wxGzhAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.WxGzhAccountInfo{
		Id:             data.ID,
		Name:           data.Name,
		Account:        data.Account,
		AppId:          data.AppID,
		AppSecret:      data.AppSecret,
		URL:            data.URL,
		Token:          data.Token,
		EncodingAesKey: data.EncodingAesKey,
		QrCodeURL:      data.QrCodeURL,
		Remark:         data.Remark,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
