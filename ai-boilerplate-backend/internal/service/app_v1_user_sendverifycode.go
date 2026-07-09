package service

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
)

// SendVerifyCode 发送验证码
func (a *AppV1UserService) SendVerifyCode(ctx context.Context, req *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeReply, error) {
	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}

	scene := data.SmsCodeSceneBind
	if err := a.smsCodeRepo.CheckSmsCodeFrequency(ctx, scene, phone); err != nil {
		return nil, err
	}

	channel, template, err := a.getSmsVerifyTemplate(ctx, scene)
	if err != nil {
		return nil, err
	}

	codeData, err := a.smsCodeRepo.GenerateSmsCodeData(scene, phone, "")
	if err != nil {
		return nil, pb.ErrorReasonDataProcessingError(pb.WithError(err))
	}
	codeData.CodeID = phone

	config := a.smsCodeRepo.GetSmsConfig(scene)
	params := map[string]string{
		"code":    codeData.Code,
		"ttl":     fmt.Sprintf("%.0f", config.CodeTTL.Minutes()),
		"minutes": fmt.Sprintf("%.0f", config.CodeTTL.Minutes()),
	}
	if err := validateSMSTemplateParams(template, params); err != nil {
		return nil, err
	}

	if err := a.smsCodeRepo.SetSmsCode(ctx, codeData); err != nil {
		return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}

	result, sendErr := sendSMSTemplate(ctx, channel, template, phone, params)
	if sendErr != nil {
		_ = a.smsCodeRepo.ClearSmsCode(ctx, scene, phone)
		if err := recordSMSSendLog(ctx, a.smsLogRepo, channel, template, phone, "", params, result, sendErr); err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("短信发送失败"))
	}

	if err := a.smsCodeRepo.SetSmsCodeFrequency(ctx, scene, phone); err != nil {
		return nil, err
	}
	if err := recordSMSSendLog(ctx, a.smsLogRepo, channel, template, phone, "", params, result, nil); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	return &pb.SendVerifyCodeReply{}, nil
}

func (a *AppV1UserService) getSmsVerifyTemplate(ctx context.Context, scene data.SmsCodeScene) (*ai_boilerplate_model.SmsChannel, *ai_boilerplate_model.SmsTemplate, error) {
	template, err := a.findEnabledSmsVerifyTemplate(ctx, string(scene))
	if err != nil {
		return nil, nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if template == nil {
		return defaultSmsVerifyChannel(), defaultSmsVerifyTemplate(), nil
	}

	channel, err := a.smsChannelRepo.FindOneCacheByID(ctx, template.SmsChannelID)
	if err != nil {
		return nil, nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if channel == nil || channel.ID == "" {
		return nil, nil, pb.ErrorReasonDataRecordNotFound(pb.WithFmtMsg("短信渠道不存在"))
	}
	if channel.Status != int16(constant.StatusEnable) {
		return nil, nil, pb.ErrorReasonParamError(pb.WithFmtMsg("短信渠道未启用"))
	}
	return channel, template, nil
}

func (a *AppV1UserService) findEnabledSmsVerifyTemplate(ctx context.Context, scene string) (*ai_boilerplate_model.SmsTemplate, error) {
	template, err := a.findEnabledSmsTemplateByCondition(ctx, []*condition.QueryParam{
		{Field: "template_type", Value: 1, Exp: condition.EQ, Logic: condition.AND},
		{Field: "template_code", Value: scene, Exp: condition.EQ, Logic: condition.AND},
		{Field: "status", Value: int16(constant.StatusEnable), Exp: condition.EQ, Logic: condition.AND},
	})
	if err != nil || template != nil {
		return template, err
	}
	return a.findEnabledSmsTemplateByCondition(ctx, []*condition.QueryParam{
		{Field: "template_type", Value: 1, Exp: condition.EQ, Logic: condition.AND},
		{Field: "status", Value: int16(constant.StatusEnable), Exp: condition.EQ, Logic: condition.AND},
	})
}

func (a *AppV1UserService) findEnabledSmsTemplateByCondition(ctx context.Context, query []*condition.QueryParam) (*ai_boilerplate_model.SmsTemplate, error) {
	list, _, err := a.smsTemplateRepo.FindMultiCacheByCondition(ctx, &condition.Req{
		Page:     1,
		PageSize: 1,
		Query:    query,
		Order: []*condition.OrderParam{
			{Field: "created_at", Order: condition.DESC},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return list[0], nil
}

func validateSMSTemplateParams(template *ai_boilerplate_model.SmsTemplate, params map[string]string) error {
	requiredParams, err := parseSMSTemplateParamKeys(template.TemplateParams)
	if err != nil {
		return pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	missingParams := make([]string, 0)
	for _, key := range requiredParams {
		if strings.TrimSpace(params[key]) == "" {
			missingParams = append(missingParams, key)
		}
	}
	if len(missingParams) > 0 {
		return pb.ErrorReasonParamError(pb.WithFmtMsg("缺少短信模板参数: %s", strings.Join(missingParams, ",")))
	}
	return nil
}
