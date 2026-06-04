package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// SendSmsTemplateMsg 短信模板-发送短信
func (a *AdminV1SmsTemplateService) SendSmsTemplateMsg(_ context.Context, _ *pb.SendSmsTemplateMsgReq) (*pb.SendSmsTemplateMsgReply, error) {
	return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("短信模板发送功能未接入"))
}
