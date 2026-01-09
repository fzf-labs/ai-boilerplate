package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// UpdateSmsTemplate 短信模板-更新一条数据
func (a *AdminV1SmsTemplateService) UpdateSmsTemplate(ctx context.Context, req *pb.UpdateSmsTemplateReq) (*pb.UpdateSmsTemplateReply, error) {
	resp := &pb.UpdateSmsTemplateReply{}
	templateParams := a.smsTemplateRepo.ParseTemplateParams(req.GetTemplateContent())
	templateParamsJSON, err := jsonutil.Marshal(templateParams)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data, err := a.smsTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.smsTemplateRepo.DeepCopy(data)
	data.SmsChannelID = req.GetSmsChannelId()
	data.TemplateType = int16(req.GetTemplateType())
	data.TemplateCode = req.GetTemplateCode()
	data.TemplateName = req.GetTemplateName()
	data.TemplateContent = req.GetTemplateContent()
	data.TemplateParams = templateParamsJSON
	data.APITemplateID = req.GetApiTemplateId()
	data.Remark = req.GetRemark()
	data.Status = int16(req.GetStatus())
	err = a.smsTemplateRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
