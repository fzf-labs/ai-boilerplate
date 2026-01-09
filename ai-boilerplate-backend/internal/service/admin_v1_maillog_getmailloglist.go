package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMailLogList 邮件日志表-列表数据查询
func (a *AdminV1MailLogService) GetMailLogList(ctx context.Context, req *pb.GetMailLogListReq) (*pb.GetMailLogListReply, error) {
	resp := &pb.GetMailLogListReply{
		Total: 0,
		List:  []*pb.MailLogInfo{},
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
	list, p, err := a.mailLogRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.MailLogInfo{
				Id:               v.ID,
				AccountId:        v.AccountID,
				FromMail:         v.FromMail,
				ToMail:           v.ToMail,
				TemplateId:       v.TemplateID,
				TemplateCode:     v.TemplateCode,
				TemplateNickname: v.TemplateNickname,
				TemplateTitle:    v.TemplateTitle,
				TemplateContent:  v.TemplateContent,
				TemplateParams:   v.TemplateParams,
				SendStatus:       v.SendStatus,
				SendTime:         timeutil.RFC3339(v.SendTime),
				SendMessageId:    v.SendMessageID,
				SendException:    v.SendException,
				CreatedAt:        v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:        v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
