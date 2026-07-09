package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMailAccountList 邮箱账号表-列表数据查询
func (a *AdminV1MailAccountService) GetMailAccountList(ctx context.Context, req *pb.GetMailAccountListReq) (*pb.GetMailAccountListReply, error) {
	resp := &pb.GetMailAccountListReply{
		Total: 0,
		List:  []*pb.MailAccountInfo{},
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
	if req.GetMail() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "mail",
			Value: req.GetMail(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetUsername() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "username",
			Value: req.GetUsername(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.mailAccountRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, mailAccountInfoFromModel(v))
		}
	}
	return resp, nil
}
