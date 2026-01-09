package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSmsChannelList 短信渠道-列表数据查询
func (a *AdminV1SmsChannelService) GetSmsChannelList(ctx context.Context, req *pb.GetSmsChannelListReq) (*pb.GetSmsChannelListReply, error) {
	resp := &pb.GetSmsChannelListReply{
		Total: 0,
		List:  []*pb.SmsChannelInfo{},
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
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetOperator() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "operator",
			Value: req.GetOperator(),
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
			Logic: condition.AND,
		})
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "created_at",
			Value: req.GetCreatedAt()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	list, p, err := a.smsChannelRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SmsChannelInfo{
				Id:           v.ID,
				Name:         v.Name,
				Operator:     v.Operator,
				Remark:       v.Remark,
				APIKey:       v.APIKey,
				APISecret:    v.APISecret,
				CallbackURL:  v.CallbackURL,
				Status:       int32(v.Status),
				CreatedAt:    v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
				OperatorName: constant.SmsChannelCodeToName[v.Operator],
			})
		}
	}
	return resp, nil
}
