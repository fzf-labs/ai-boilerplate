package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// SysAuthUpdateAdminInfo Auth-更新用户信息
func (a *AdminV1SysAuthService) SysAuthUpdateAdminInfo(ctx context.Context, req *pb.SysAuthUpdateAdminInfoReq) (*pb.SysAuthUpdateAdminInfoReply, error) {
	resp := &pb.SysAuthUpdateAdminInfoReply{}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	admin, err := a.sysAdminRepo.FindOneCacheByID(ctx, adminID)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if admin == nil || admin.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysAdminRepo.DeepCopy(admin)
	admin.Nickname = req.GetNickname()
	admin.Avatar = req.GetAvatar()
	admin.Sex = int16(req.GetSex())
	err = a.sysAdminRepo.UpdateOneCacheWithZero(ctx, admin, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
