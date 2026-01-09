package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSmsTemplateStatus 短信模板-更新状态
func (a *AdminV1SmsTemplateService) UpdateSmsTemplateStatus(ctx context.Context, req *pb.UpdateSmsTemplateStatusReq) (*pb.UpdateSmsTemplateStatusReply, error) {
	resp := &pb.UpdateSmsTemplateStatusReply{}
	data, err := a.smsTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.smsTemplateRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.smsTemplateRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
