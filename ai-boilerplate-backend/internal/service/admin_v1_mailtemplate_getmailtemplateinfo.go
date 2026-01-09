package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMailTemplateInfo 邮件模版表-单条数据查询
func (a *AdminV1MailTemplateService) GetMailTemplateInfo(ctx context.Context, req *pb.GetMailTemplateInfoReq) (*pb.GetMailTemplateInfoReply, error) {
	resp := &pb.GetMailTemplateInfoReply{}
	data, err := a.mailTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MailTemplateInfo{
		Id:        data.ID,
		Name:      data.Name,
		Code:      data.Code,
		AccountId: data.AccountID,
		Nickname:  data.Nickname,
		Title:     data.Title,
		Content:   data.Content,
		Params:    string(data.Params),
		Remark:    data.Remark,
		Status:    data.Status,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
