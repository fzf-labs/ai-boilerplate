package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/cryptutil"
)

// UpdateSysAdminPassword 系统-用户-重置密码
func (a *AdminV1SysAdminService) UpdateSysAdminPassword(ctx context.Context, req *pb.UpdateSysAdminPasswordReq) (*pb.UpdateSysAdminPasswordReply, error) {
	resp := &pb.UpdateSysAdminPasswordReply{}
	password, err := cryptutil.Encrypt(req.GetPassword())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysAdminRepo.DeepCopy(data)
	data.Password = password
	err = a.sysAdminRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
