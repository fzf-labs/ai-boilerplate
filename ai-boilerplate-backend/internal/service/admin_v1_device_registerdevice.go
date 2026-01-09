package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/timeutil"
)

// RegisterDevice 设备表-注册设备
func (a *AdminV1DeviceService) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceReq) (*pb.RegisterDeviceReply, error) {
	resp := &pb.RegisterDeviceReply{}
	data, err := a.deviceRepo.FindOneCacheBySn(ctx, req.GetSn())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data != nil && data.ID != "" {
		return nil, pb.ErrorReasonDataDuplicateRecord()
	}
	// 注册设备
	err = a.deviceRepo.CreateOneCache(ctx, &ai_boilerplate_model.Device{
		Sn:           req.GetSn(),
		RegistryTime: timeutil.NowSQLNullTime(),
		Status:       int32(constant.DeviceStatusEnable),
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
