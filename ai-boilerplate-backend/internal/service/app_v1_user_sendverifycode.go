package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// SendVerifyCode 发送验证码
func (a *AppV1UserService) SendVerifyCode(ctx context.Context, req *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeReply, error) {
	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}
	return nil, pb.ErrorReasonAPIThirdErr(pb.WithFmtMsg("短信验证码功能未接入"))
}
