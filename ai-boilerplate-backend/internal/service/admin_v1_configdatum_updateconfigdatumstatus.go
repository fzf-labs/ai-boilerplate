package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateConfigDatumStatus 配置管理-更新状态
func (a *AdminV1ConfigDatumService) UpdateConfigDatumStatus(ctx context.Context, req *pb.UpdateConfigDatumStatusReq) (*pb.UpdateConfigDatumStatusReply, error) {
	resp := &pb.UpdateConfigDatumStatusReply{}
	data, err := a.configDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.configDatumRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.configDatumRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
