package service

import (
	"context"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/menu/request"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/dtoutil"
	"github.com/fzf-labs/goutil/jsonutil"
)

// StoreWxGzhMenu 公众号菜单表-创建/更新一条数据
func (a *AdminV1WxGzhMenuService) StoreWxGzhMenu(ctx context.Context, req *pb.StoreWxGzhMenuReq) (*pb.StoreWxGzhMenuReply, error) {
	resp := &pb.StoreWxGzhMenuReply{}
	selfmenuInfo, err := jsonutil.Marshal(req.GetSelfmenuInfo())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	data, err := a.wxGzhMenuRepo.FindOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		data = a.wxGzhMenuRepo.NewData()
		data.AppID = req.GetAppId()
		data.SelfmenuInfo = selfmenuInfo
		err = a.wxGzhMenuRepo.CreateOneCache(ctx, data)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	} else {
		oldData := a.wxGzhMenuRepo.DeepCopy(data)
		data.SelfmenuInfo = selfmenuInfo
		err = a.wxGzhMenuRepo.UpdateOneCacheWithZero(ctx, data, oldData)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	wxGzhAccount, err := a.wxGzhAccountRepo.FindOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	officialAccount, err := a.wxGzhAccountRepo.NewOfficialAccountClient(wxGzhAccount.AppID, wxGzhAccount.AppSecret, wxGzhAccount.Token, wxGzhAccount.EncodingAesKey)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	buttons := make([]*request.Button, 0)
	err = dtoutil.Copy(&buttons, req.GetSelfmenuInfo().GetButton())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	responseMenuCreate, err := officialAccount.Menu.Create(ctx, buttons)
	if err != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	jsonutil.Dump(responseMenuCreate)
	if responseMenuCreate.ErrCode != 0 {
		return nil, pb.ErrorReasonAPIThirdErr()
	}
	return resp, nil
}
