package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateWxGzhAccount 公众号账号表-创建一条数据
func (a *AdminV1WxGzhAccountService) CreateWxGzhAccount(ctx context.Context, req *pb.CreateWxGzhAccountReq) (*pb.CreateWxGzhAccountReply, error) {
	resp := &pb.CreateWxGzhAccountReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.wxGzhAccountRepo.NewData()
	data.TenantID = tenantID
	data.Name = req.GetName()
	data.Account = req.GetAccount()
	data.AppID = req.GetAppId()
	data.AppSecret = req.GetAppSecret()
	data.URL = req.GetURL()
	data.Token = req.GetToken()
	data.EncodingAesKey = req.GetEncodingAesKey()
	data.QrCodeURL = req.GetQrCodeURL()
	data.Remark = req.GetRemark()
	err := a.wxGzhAccountRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
