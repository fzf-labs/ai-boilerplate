package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetWxGzhAccountList 公众号账号表-列表数据查询
func (a *AdminV1WxGzhAccountService) GetWxGzhAccountList(ctx context.Context, req *pb.GetWxGzhAccountListReq) (*pb.GetWxGzhAccountListReply, error) {
	resp := &pb.GetWxGzhAccountListReply{
		Total: 0,
		List:  []*pb.WxGzhAccountInfo{},
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
	if req.GetAccount() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "account",
			Value: req.GetAccount(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetAppId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "app_id",
			Value: req.GetAppId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.wxGzhAccountRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, wxGzhAccountInfoFromModel(v))
		}
	}
	return resp, nil
}
