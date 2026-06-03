package service

import (
	"context"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// BindPhone 绑定手机号
func (a *AppV1UserService) BindPhone(ctx context.Context, req *pb.BindPhoneReq) (*pb.BindPhoneReply, error) {
	resp := &pb.BindPhoneReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	dataUser, err := a.userRepo.FindOneCacheByID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if dataUser == nil || dataUser.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}

	phone := strings.TrimSpace(req.GetPhone())
	if phone == "" {
		return nil, pb.ErrorReasonParamError()
	}
	codeData, err := a.smsCodeRepo.CheckSmsCode(ctx, data.SmsCodeSceneBind, phone, strings.TrimSpace(req.GetCode()))
	if err != nil {
		return nil, err
	}
	if codeData.Phone != phone {
		return nil, pb.ErrorReasonSmsCodeInvalid()
	}

	existUser, err := a.userRepo.FindOneCacheByPhone(ctx, phone)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if existUser != nil && existUser.ID != "" && existUser.ID != dataUser.ID {
		return nil, pb.ErrorReasonAccountAlreadyExists()
	}

	oldData := a.userRepo.DeepCopy(dataUser)
	dataUser.Phone = phone
	if dataUser.Nickname == "" {
		dataUser.Nickname = a.userRepo.GenerateNicknameByPhone(phone)
	}
	if err := a.userRepo.UpdateOneCacheWithZero(ctx, dataUser, oldData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	_ = a.smsCodeRepo.ClearSmsCode(ctx, data.SmsCodeSceneBind, phone)
	return resp, nil
}
