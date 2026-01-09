package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMailTemplateStatus 邮件模版表-更新状态
func (a *AdminV1MailTemplateService) UpdateMailTemplateStatus(ctx context.Context, req *pb.UpdateMailTemplateStatusReq) (*pb.UpdateMailTemplateStatusReply, error) {
	resp := &pb.UpdateMailTemplateStatusReply{}
	data, err := a.mailTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mailTemplateRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.mailTemplateRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
