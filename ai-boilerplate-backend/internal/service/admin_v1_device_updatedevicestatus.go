package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateDeviceStatus 设备表-更新状态
func (a *AdminV1DeviceService) UpdateDeviceStatus(ctx context.Context, req *pb.UpdateDeviceStatusReq) (*pb.UpdateDeviceStatusReply, error) {
	resp := &pb.UpdateDeviceStatusReply{}
	data, err := a.deviceRepo.FindOneCacheBySn(ctx, req.GetSn())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.deviceRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.deviceRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
