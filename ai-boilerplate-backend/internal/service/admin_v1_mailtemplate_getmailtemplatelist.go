package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMailTemplateList 邮件模版表-列表数据查询
func (a *AdminV1MailTemplateService) GetMailTemplateList(ctx context.Context, req *pb.GetMailTemplateListReq) (*pb.GetMailTemplateListReply, error) {
	resp := &pb.GetMailTemplateListReply{
		Total: 0,
		List:  []*pb.MailTemplateInfo{},
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
	if req.GetCode() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "code",
			Value: req.GetCode(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: req.GetName(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetAccountId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "account_id",
			Value: req.GetAccountId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[0],
			Exp:   condition.GTE,
			Logic: condition.OR,
		})
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.OR,
		})
	}
	list, p, err := a.mailTemplateRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MailTemplateInfo{
				Id:        v.ID,
				Name:      v.Name,
				Code:      v.Code,
				AccountId: v.AccountID,
				Nickname:  v.Nickname,
				Title:     v.Title,
				Content:   v.Content,
				Params:    string(v.Params),
				Remark:    v.Remark,
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
