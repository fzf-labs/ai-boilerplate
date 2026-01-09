package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateWxGzhAccount 公众号账号表-更新一条数据
func (a *AdminV1WxGzhAccountService) UpdateWxGzhAccount(ctx context.Context, req *pb.UpdateWxGzhAccountReq) (*pb.UpdateWxGzhAccountReply, error) {
	resp := &pb.UpdateWxGzhAccountReply{}
	data, err := a.wxGzhAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.wxGzhAccountRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Account = req.GetAccount()
	data.AppID = req.GetAppId()
	data.AppSecret = req.GetAppSecret()
	data.URL = req.GetURL()
	data.Token = req.GetToken()
	data.EncodingAesKey = req.GetEncodingAesKey()
	data.QrCodeURL = req.GetQrCodeURL()
	data.Remark = req.GetRemark()
	err = a.wxGzhAccountRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
