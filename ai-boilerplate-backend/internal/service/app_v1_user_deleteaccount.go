package service

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/kratos-contrib/meta"
)

const deletedUserReferencePrefix = "deleted_user:"

func deletedUserReference(userID string) string {
	return fmt.Sprintf("%s%s", deletedUserReferencePrefix, userID)
}

func (a *AppV1UserService) anonymizeRetainedUserReferences(ctx context.Context, tx *ai_boilerplate_dao.Query, user *ai_boilerplate_model.User) error {
	userID := user.ID
	deletedUserID := deletedUserReference(userID)
	strippedUserID := strings.ReplaceAll(userID, "-", "")
	if len(strippedUserID) < 8 {
		strippedUserID += "00000000"
	}
	deletedSmsMobile := "del" + strippedUserID[:8]

	orders, err := a.mallOrderRepo.FindMultiCacheByUserID(ctx, userID)
	if err != nil {
		return err
	}
	for _, order := range orders {
		oldOrder := a.mallOrderRepo.DeepCopy(order)
		order.UserID = deletedUserID
		if err := a.mallOrderRepo.UpdateOneCacheWithZeroByTx(ctx, tx, order, oldOrder); err != nil {
			return err
		}
	}

	activationCodes, err := a.mallActivationCodeRepo.FindMultiCacheByUserID(ctx, userID)
	if err != nil {
		return err
	}
	for _, activationCode := range activationCodes {
		oldActivationCode := a.mallActivationCodeRepo.DeepCopy(activationCode)
		activationCode.UserID = deletedUserID
		if err := a.mallActivationCodeRepo.UpdateOneCacheWithZeroByTx(ctx, tx, activationCode, oldActivationCode); err != nil {
			return err
		}
	}
	if user.Phone != "" {
		smsLogs, err := a.smsLogRepo.FindMultiCacheByMobile(ctx, user.Phone)
		if err != nil {
			return err
		}
		for _, smsLog := range smsLogs {
			oldSmsLog := a.smsLogRepo.DeepCopy(smsLog)
			smsLog.Mobile = deletedSmsMobile
			smsLog.UserID = deletedUserID
			if err := a.smsLogRepo.UpdateOneCacheWithZeroByTx(ctx, tx, smsLog, oldSmsLog); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteAccount 注销账号
func (a *AppV1UserService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountReq) (*pb.DeleteAccountReply, error) {
	resp := &pb.DeleteAccountReply{}
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)
	dataUser, err := a.userRepo.FindOneCacheByID(ctx, userID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if dataUser == nil || dataUser.ID == "" {
		return nil, pb.ErrorReasonAccountNotFound()
	}
	if !a.userRepo.VerifyPassword(dataUser.Salt, req.GetPassword(), dataUser.Password) {
		return nil, pb.ErrorReasonAccountPasswordError()
	}

	if err := a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		if err := a.anonymizeRetainedUserReferences(ctx, tx, dataUser); err != nil {
			return err
		}
		if err := a.userMessageRepo.DeleteMultiCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		if err := a.helpFeedbackRepo.DeleteMultiCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		if err := a.userNotificationSettingRepo.DeleteOneCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		if err := a.userBindDeviceRepo.DeleteMultiCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		if err := a.userMembershipChangeRepo.DeleteMultiCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		if err := a.userMembershipRepo.DeleteOneCacheByUserIDTx(ctx, tx, dataUser.ID); err != nil {
			return err
		}
		oldUser := a.userRepo.DeepCopy(dataUser)
		deletedUser := a.userRepo.DeepCopy(dataUser)
		for field, value := range a.userRepo.BuildDeletedUserFields(dataUser.ID) {
			switch field {
			case "phone":
				deletedUser.Phone = value.(string)
			case "password":
				deletedUser.Password = value.(string)
			case "salt":
				deletedUser.Salt = value.(string)
			case "nickname":
				deletedUser.Nickname = value.(string)
			case "gender":
				deletedUser.Gender = value.(int32)
			case "avatar":
				deletedUser.Avatar = value.(string)
			case "profile":
				deletedUser.Profile = value.(string)
			case "other":
				deletedUser.Other = nil
			case "wx_gzh_user_id":
				deletedUser.WxGzhUserID = value.(string)
			case "wx_gzh_xcx_id":
				deletedUser.WxGzhXcxID = value.(string)
			case "status":
				deletedUser.Status = value.(int32)
			}
		}
		if err := a.userRepo.UpdateOneCacheWithZeroByTx(ctx, tx, deletedUser, oldUser); err != nil {
			return err
		}
		return a.userRepo.DeleteOneCacheByIDTx(ctx, tx, dataUser.ID)
	}); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if err := a.userRepo.JwtTokenClear(ctx, dataUser.ID); err != nil {
		a.log.Errorf("clear user token after delete account failed: %v", err)
	}
	return resp, nil
}
