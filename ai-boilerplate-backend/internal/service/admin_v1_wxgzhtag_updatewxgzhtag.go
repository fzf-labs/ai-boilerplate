package service

import (
	"context"
	"strconv"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateWxGzhTag 公众号标签表-更新一条数据
func (a *AdminV1WxGzhTagService) UpdateWxGzhTag(ctx context.Context, req *pb.UpdateWxGzhTagReq) (*pb.UpdateWxGzhTagReply, error) {
	resp := &pb.UpdateWxGzhTagReply{}
	data, err := a.wxGzhTagRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	// 查询账号信息
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, data.AppID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 微信公众号更新标签
	_, err = officialAccount.UserTag.Update(ctx, strconv.Itoa(int(data.TagID)), req.GetName())
	if err != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	// 数据更新
	oldData := a.wxGzhTagRepo.DeepCopy(data)
	data.Name = req.GetName()
	err = a.wxGzhTagRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
