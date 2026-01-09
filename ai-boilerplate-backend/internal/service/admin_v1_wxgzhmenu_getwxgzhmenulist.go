package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetWxGzhMenuList 公众号菜单表-列表数据查询
func (a *AdminV1WxGzhMenuService) GetWxGzhMenuList(ctx context.Context, req *pb.GetWxGzhMenuListReq) (*pb.GetWxGzhMenuListReply, error) {
	resp := &pb.GetWxGzhMenuListReply{
		Total: 0,
		List:  []*pb.WxGzhMenuInfo{},
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
	list, p, err := a.wxGzhMenuRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			selfmenuInfo := &pb.SelfmenuInfo{}
			err := jsonutil.Unmarshal(v.SelfmenuInfo, selfmenuInfo)
			if err != nil {
				return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
			}
			resp.List = append(resp.List, &pb.WxGzhMenuInfo{
				Id:           v.ID,
				AppId:        v.AppID,
				IsMenuOpen:   v.IsMenuOpen,
				SelfmenuInfo: selfmenuInfo,
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
