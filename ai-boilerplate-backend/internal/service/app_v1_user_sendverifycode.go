package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
)

// SendVerifyCode 发送验证码
func (a *AppV1UserService) SendVerifyCode(ctx context.Context, req *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeReply, error) {
	resp := &pb.SendVerifyCodeReply{}
	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}
	if err := a.smsCodeRepo.CheckSmsCodeFrequency(ctx, data.SmsCodeSceneBind, phone); err != nil {
		return nil, err
	}
	codeData, err := a.smsCodeRepo.GenerateSmsCodeData(data.SmsCodeSceneBind, phone, "")
	if err != nil {
		return nil, pb.ErrorReasonDataProcessingError(pb.WithError(err))
	}
	codeData.CodeID = phone
	if err := a.smsCodeRepo.SetSmsCode(ctx, codeData); err != nil {
		return nil, err
	}
	if err := a.smsCodeRepo.SetSmsCodeFrequency(ctx, data.SmsCodeSceneBind, phone); err != nil {
		return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	if err := a.smsCodeRepo.SendSmsCode(ctx, codeData, func(_ context.Context, codeData *data.SmsCodeData) error {
		a.log.Infof("sms verify code generated scene=%s phone=%s code=%s", codeData.Scene, codeData.Phone, codeData.Code)
		return nil
	}); err != nil {
		return nil, pb.ErrorReasonAPIThirdErr(pb.WithError(err))
	}
	return resp, nil
}
