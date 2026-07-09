package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

// SendSmsTemplateMsg 短信模板-发送短信
func (a *AdminV1SmsTemplateService) SendSmsTemplateMsg(ctx context.Context, req *pb.SendSmsTemplateMsgReq) (*pb.SendSmsTemplateMsgReply, error) {
	template, err := a.smsTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if template == nil || template.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if template.Status != int16(constant.StatusEnable) {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("短信模板未启用"))
	}

	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}

	channel, err := a.smsChannelRepo.FindOneCacheByID(ctx, template.SmsChannelID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if channel == nil || channel.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound(pb.WithFmtMsg("短信渠道不存在"))
	}
	if channel.Status != int16(constant.StatusEnable) {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("短信渠道未启用"))
	}

	sendParams := req.GetParams()
	if sendParams == nil {
		sendParams = map[string]string{}
	}
	if err := validateAdminSMSTemplateParams(template.TemplateParams, sendParams); err != nil {
		return nil, err
	}

	result, sendErr := sendSMSTemplate(ctx, channel, template, phone, sendParams)
	if err := recordSMSSendLog(ctx, a.smsLogRepo, channel, template, phone, "", sendParams, result, sendErr); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if sendErr != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("短信发送失败"))
	}
	return &pb.SendSmsTemplateMsgReply{}, nil
}

func validateAdminSMSTemplateParams(rawTemplateParams []byte, sendParams map[string]string) error {
	requiredParams, err := parseSMSTemplateParamKeys(rawTemplateParams)
	if err != nil {
		return pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	missingParams := make([]string, 0)
	for _, key := range requiredParams {
		if strings.TrimSpace(sendParams[key]) == "" {
			missingParams = append(missingParams, key)
		}
	}
	if len(missingParams) > 0 {
		return pb.ErrorReasonParamError(pb.WithFmtMsg("缺少短信模板参数: %s", strings.Join(missingParams, ",")))
	}
	return nil
}
