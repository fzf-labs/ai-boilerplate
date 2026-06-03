package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/kratos-contrib/meta"
)

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

	userMembership, err := a.userMembershipRepo.FindOneCacheByUserID(ctx, dataUser.ID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if err := a.commonRepo.Transaction(ctx, func(tx *ai_boilerplate_dao.Query) error {
		if userMembership != nil && userMembership.ID != "" {
			if err := a.userMembershipRepo.DeleteOneCacheByIDTx(ctx, tx, userMembership.ID); err != nil {
				return err
			}
		}
		return a.userRepo.DeleteOneCacheByIDTx(ctx, tx, dataUser.ID)
	}); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
