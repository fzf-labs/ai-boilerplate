package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetWxGzhMenuInfo 公众号菜单表-单条数据查询
func (a *AdminV1WxGzhMenuService) GetWxGzhMenuInfo(ctx context.Context, req *pb.GetWxGzhMenuInfoReq) (*pb.GetWxGzhMenuInfoReply, error) {
	resp := &pb.GetWxGzhMenuInfoReply{
		Info: &pb.WxGzhMenuInfo{},
	}
	data, err := a.wxGzhMenuRepo.FindOneCacheByAppID(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	selfmenuInfo := &pb.SelfmenuInfo{}
	if data.SelfmenuInfo.String() != "" {
		err = jsonutil.Unmarshal(data.SelfmenuInfo, selfmenuInfo)
		if err != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
		}
	}
	resp.Info = &pb.WxGzhMenuInfo{
		Id:           data.ID,
		AppId:        data.AppID,
		IsMenuOpen:   data.IsMenuOpen,
		SelfmenuInfo: selfmenuInfo,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
