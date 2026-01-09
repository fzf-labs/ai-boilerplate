package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateWxGzhTag 公众号标签表-创建一条数据
func (a *AdminV1WxGzhTagService) CreateWxGzhTag(ctx context.Context, req *pb.CreateWxGzhTagReq) (*pb.CreateWxGzhTagReply, error) {
	resp := &pb.CreateWxGzhTagReply{}
	// 查询账号信息
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 微信公众号创建标签
	wxTag, err := officialAccount.UserTag.Create(ctx, req.GetName())
	if err != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	// 数据入库
	data := a.wxGzhTagRepo.NewData()
	data.AppID = req.GetAppId()
	data.TagID = int32(wxTag.Tag.ID)
	data.Name = req.GetName()
	data.Count = int32(wxTag.Tag.Count)
	err = a.wxGzhTagRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
