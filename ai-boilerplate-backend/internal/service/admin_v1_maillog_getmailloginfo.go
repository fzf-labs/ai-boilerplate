package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMailLogInfo 邮件日志表-单条数据查询
func (a *AdminV1MailLogService) GetMailLogInfo(ctx context.Context, req *pb.GetMailLogInfoReq) (*pb.GetMailLogInfoReply, error) {
	resp := &pb.GetMailLogInfoReply{}
	data, err := a.mailLogRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MailLogInfo{
		Id:               data.ID,
		AccountId:        data.AccountID,
		FromMail:         data.FromMail,
		ToMail:           data.ToMail,
		TemplateId:       data.TemplateID,
		TemplateCode:     data.TemplateCode,
		TemplateNickname: data.TemplateNickname,
		TemplateTitle:    data.TemplateTitle,
		TemplateContent:  data.TemplateContent,
		TemplateParams:   data.TemplateParams,
		SendStatus:       data.SendStatus,
		SendTime:         timeutil.RFC3339(data.SendTime),
		SendMessageId:    data.SendMessageID,
		SendException:    data.SendException,
		CreatedAt:        timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:        timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
