package service

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/uuidutil"
	"gorm.io/datatypes"
)

const (
	smsSendStatusSuccess = "success"
	smsSendStatusFailed  = "failed"
	smsReceiveStatusWait = "waiting"
	smsMask              = "******"

	defaultSmsVerifyChannelID     = "mock-sms-channel"
	defaultSmsVerifyTemplateID    = "mock-sms-verify-code"
	defaultSmsVerifyTemplateCode  = "sms_bind"
	defaultSmsVerifyTemplateName  = "短信验证码"
	defaultSmsVerifyTemplateText  = "您的验证码是{code}，5分钟内有效。"
	defaultSmsVerifyAPITemplateID = "mock-sms-verify-code"
)

type smsProvider interface {
	Send(ctx context.Context, req smsProviderRequest) (*smsProviderResult, error)
}

type smsProviderRequest struct {
	Phone           string
	TemplateID      string
	TemplateCode    string
	TemplateContent string
	Params          map[string]string
	Operator        string
	APIKey          string
	APISecret       string
}

type smsProviderResult struct {
	SendCode  string
	SendMsg   string
	RequestID string
	SerialNo  string
}

type mockSMSProvider struct{}

func (m mockSMSProvider) Send(_ context.Context, req smsProviderRequest) (*smsProviderResult, error) {
	result := &smsProviderResult{
		SendCode:  "MOCK_OK",
		SendMsg:   "mock sms provider accepted",
		RequestID: uuidutil.GenUUID(),
		SerialNo:  uuidutil.GenUUID(),
	}
	if strings.EqualFold(strings.TrimSpace(req.APIKey), "fail") || strings.EqualFold(strings.TrimSpace(req.Params["mock_error"]), "true") {
		result.SendCode = "MOCK_FAILED"
		result.SendMsg = "mock sms provider rejected message"
		return result, errors.New("mock sms provider rejected message")
	}
	return result, nil
}

func selectSMSProvider(_ string) smsProvider {
	return mockSMSProvider{}
}

func sendSMSTemplate(ctx context.Context, channel *ai_boilerplate_model.SmsChannel, template *ai_boilerplate_model.SmsTemplate, phone string, params map[string]string) (*smsProviderResult, error) {
	if params == nil {
		params = map[string]string{}
	}
	renderedContent := renderSMSTemplateContent(template.TemplateContent, params)
	provider := selectSMSProvider(channel.Operator)
	return provider.Send(ctx, smsProviderRequest{
		Phone:           phone,
		TemplateID:      template.APITemplateID,
		TemplateCode:    template.TemplateCode,
		TemplateContent: renderedContent,
		Params:          params,
		Operator:        channel.Operator,
		APIKey:          channel.APIKey,
		APISecret:       channel.APISecret,
	})
}

func recordSMSSendLog(
	ctx context.Context,
	smsLogRepo *data.SmsLogRepo,
	channel *ai_boilerplate_model.SmsChannel,
	template *ai_boilerplate_model.SmsTemplate,
	phone string,
	userID string,
	params map[string]string,
	result *smsProviderResult,
	sendErr error,
) error {
	if result == nil {
		result = &smsProviderResult{
			SendCode:  "UNKNOWN",
			RequestID: uuidutil.GenUUID(),
		}
	}
	sendStatus := smsSendStatusSuccess
	if sendErr != nil {
		sendStatus = smsSendStatusFailed
		if result.SendMsg == "" {
			result.SendMsg = sendErr.Error()
		}
	}

	now := time.Now()
	logData := smsLogRepo.NewData()
	logData.ID = uuidutil.GenUUID()
	logData.SmsChannelID = channel.ID
	logData.SmsTemplateID = template.ID
	logData.SmsParamsContent = safeSMSParamsContent(params)
	logData.Mobile = phone
	logData.UserID = userID
	logData.SendStatus = sendStatus
	logData.SendTime = now
	logData.ReceiveStatus = smsReceiveStatusWait
	logData.APISendCode = truncateSMSLogValue(result.SendCode)
	logData.APISendMsg = truncateSMSLogValue(sanitizeSMSLogMessage(result.SendMsg, channel, params))
	logData.APIRequestID = truncateSMSLogValue(result.RequestID)
	logData.APISerialNo = truncateSMSLogValue(result.SerialNo)
	logData.CreatedAt = now
	return smsLogRepo.CreateOneCache(ctx, logData)
}

func parseSMSTemplateParamKeys(raw datatypes.JSON) ([]string, error) {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" {
		return nil, nil
	}

	var paramMap map[string]string
	if err := json.Unmarshal([]byte(trimmed), &paramMap); err == nil {
		return dedupeSMSParamKeysFromMap(paramMap), nil
	}

	var paramList []string
	if err := json.Unmarshal([]byte(trimmed), &paramList); err != nil {
		return nil, err
	}
	return dedupeSMSParamKeys(paramList), nil
}

func dedupeSMSParamKeysFromMap(params map[string]string) []string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	return dedupeSMSParamKeys(keys)
}

func dedupeSMSParamKeys(params []string) []string {
	cleaned := make([]string, 0, len(params))
	seen := make(map[string]struct{}, len(params))
	for _, param := range params {
		param = strings.TrimSpace(param)
		if param == "" {
			continue
		}
		if _, ok := seen[param]; ok {
			continue
		}
		seen[param] = struct{}{}
		cleaned = append(cleaned, param)
	}
	return cleaned
}

func renderSMSTemplateContent(content string, params map[string]string) string {
	rendered := content
	for key, value := range params {
		rendered = strings.ReplaceAll(rendered, "{"+key+"}", value)
	}
	return rendered
}

func safeSMSParamsContent(params map[string]string) string {
	if params == nil {
		params = map[string]string{}
	}
	safeParams := make(map[string]string, len(params))
	for key, value := range params {
		if isSensitiveSMSParam(key) {
			safeParams[key] = smsMask
			continue
		}
		safeParams[key] = value
	}
	raw, err := json.Marshal(safeParams)
	if err != nil {
		return "{}"
	}
	return truncateSMSLogValue(string(raw))
}

func sanitizeSMSLogMessage(message string, channel *ai_boilerplate_model.SmsChannel, params map[string]string) string {
	sanitized := message
	if channel != nil {
		sanitized = replaceSecret(sanitized, channel.APIKey)
		sanitized = replaceSecret(sanitized, channel.APISecret)
	}
	for key, value := range params {
		if isSensitiveSMSParam(key) {
			sanitized = replaceSecret(sanitized, value)
		}
	}
	return sanitized
}

func replaceSecret(value, secret string) string {
	secret = strings.TrimSpace(secret)
	if secret == "" {
		return value
	}
	return strings.ReplaceAll(value, secret, smsMask)
}

func isSensitiveSMSParam(key string) bool {
	normalized := strings.ToLower(strings.TrimSpace(key))
	if normalized == "" {
		return false
	}
	for _, part := range []string{"code", "secret", "key", "password", "token", "验证码", "密钥"} {
		if strings.Contains(normalized, part) {
			return true
		}
	}
	return false
}

func truncateSMSLogValue(value string) string {
	const maxSMSLogValueLength = 255
	runes := []rune(value)
	if len(runes) <= maxSMSLogValueLength {
		return value
	}
	return string(runes[:maxSMSLogValueLength])
}

func defaultSmsVerifyChannel() *ai_boilerplate_model.SmsChannel {
	now := time.Now()
	return &ai_boilerplate_model.SmsChannel{
		ID:        defaultSmsVerifyChannelID,
		Name:      "Mock SMS",
		Operator:  "MOCK",
		APIKey:    "mock",
		Status:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func defaultSmsVerifyTemplate() *ai_boilerplate_model.SmsTemplate {
	now := time.Now()
	return &ai_boilerplate_model.SmsTemplate{
		ID:              defaultSmsVerifyTemplateID,
		SmsChannelID:    defaultSmsVerifyChannelID,
		TemplateType:    1,
		TemplateCode:    defaultSmsVerifyTemplateCode,
		TemplateName:    defaultSmsVerifyTemplateName,
		TemplateContent: defaultSmsVerifyTemplateText,
		APITemplateID:   defaultSmsVerifyAPITemplateID,
		Status:          1,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
