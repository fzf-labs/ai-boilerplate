package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// CreateSmsTemplate 短信模板-创建一条数据
func (a *AdminV1SmsTemplateService) CreateSmsTemplate(ctx context.Context, req *pb.CreateSmsTemplateReq) (*pb.CreateSmsTemplateReply, error) {
	resp := &pb.CreateSmsTemplateReply{}
	templateParams := a.smsTemplateRepo.ParseTemplateParams(req.GetTemplateContent())
	templateParamsJSON, err := jsonutil.Marshal(templateParams)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := a.smsTemplateRepo.NewData()
	data.SmsChannelID = req.GetSmsChannelId()
	data.TemplateType = int16(req.GetTemplateType())
	data.TemplateCode = req.GetTemplateCode()
	data.TemplateName = req.GetTemplateName()
	data.TemplateContent = req.GetTemplateContent()
	data.TemplateParams = templateParamsJSON
	data.APITemplateID = req.GetApiTemplateId()
	data.Remark = req.GetRemark()
	data.Status = int16(req.GetStatus())
	err = a.smsTemplateRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
