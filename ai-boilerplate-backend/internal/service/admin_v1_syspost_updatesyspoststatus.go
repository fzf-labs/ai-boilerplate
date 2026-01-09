package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysPostStatus 系统-工作岗位-更新状态
func (a *AdminV1SysPostService) UpdateSysPostStatus(ctx context.Context, req *pb.UpdateSysPostStatusReq) (*pb.UpdateSysPostStatusReply, error) {
	resp := &pb.UpdateSysPostStatusReply{}
	data, err := a.sysPostRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysPostRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.sysPostRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
