package service

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/uuidutil"
)

// SendMailTemplateMsg 邮件模板-发送邮件
func (a *AdminV1MailTemplateService) SendMailTemplateMsg(ctx context.Context, req *pb.SendMailTemplateMsgReq) (*pb.SendMailTemplateMsgReply, error) {
	template, err := a.mailTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if template == nil || template.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if template.Status != int32(constant.StatusEnable) {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("邮件模板未启用"))
	}

	mail := strings.TrimSpace(req.GetMail())
	if mail == "" {
		return nil, pb.ErrorReasonParamError()
	}

	account, err := a.mailAccountRepo.FindOneCacheByID(ctx, template.AccountID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if account == nil || account.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if account.Status != int32(constant.StatusEnable) {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("邮箱账号未启用"))
	}

	requiredParams, err := parseMailTemplateParams(template.Params)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	sendParams := req.GetParams()
	if sendParams == nil {
		sendParams = map[string]string{}
	}
	missingParams := make([]string, 0)
	for _, key := range requiredParams {
		if strings.TrimSpace(sendParams[key]) == "" {
			missingParams = append(missingParams, key)
		}
	}
	if len(missingParams) > 0 {
		return nil, pb.ErrorReasonParamError(pb.WithFmtMsg("缺少模板参数: %s", strings.Join(missingParams, ",")))
	}

	title, err := renderMailTemplateString(template.Title, sendParams)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	content, err := renderMailTemplateString(template.Content, sendParams)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}

	messageID, sendErr := sendSMTPMail(account, template.Nickname, mail, title, content)
	paramsJSON, _ := json.Marshal(sendParams)

	logData := a.mailLogRepo.NewData()
	logData.ID = uuidutil.GenUUID()
	logData.AccountID = account.ID
	logData.FromMail = account.Mail
	logData.ToMail = mail
	logData.TemplateID = template.ID
	logData.TemplateCode = template.Code
	logData.TemplateNickname = template.Nickname
	logData.TemplateTitle = title
	logData.TemplateContent = content
	logData.TemplateParams = string(paramsJSON)
	logData.SendTime = time.Now()
	logData.SendMessageID = messageID

	if sendErr != nil {
		logData.SendStatus = -1
		logData.SendException = sendErr.Error()
		if err := a.mailLogRepo.CreateOneCache(ctx, logData); err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("邮件发送失败: %v", sendErr))
	}

	logData.SendStatus = 1
	if err := a.mailLogRepo.CreateOneCache(ctx, logData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return &pb.SendMailTemplateMsgReply{}, nil
}
