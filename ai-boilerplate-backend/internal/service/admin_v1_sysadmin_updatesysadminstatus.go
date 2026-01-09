package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysAdminStatus 系统-用户-更新状态
func (a *AdminV1SysAdminService) UpdateSysAdminStatus(ctx context.Context, req *pb.UpdateSysAdminStatusReq) (*pb.UpdateSysAdminStatusReply, error) {
	resp := &pb.UpdateSysAdminStatusReply{}
	data, err := a.sysAdminRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysAdminRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.sysAdminRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
