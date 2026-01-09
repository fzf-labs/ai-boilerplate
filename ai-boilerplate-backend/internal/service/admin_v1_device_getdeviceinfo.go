package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetDeviceInfo 设备表-单条数据查询
func (a *AdminV1DeviceService) GetDeviceInfo(ctx context.Context, req *pb.GetDeviceInfoReq) (*pb.GetDeviceInfoReply, error) {
	resp := &pb.GetDeviceInfoReply{}
	data, err := a.deviceRepo.FindOneCacheBySn(ctx, req.GetSn())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	devicePush := &pb.DevicePush{}
	if data.Push.String() != "" {
		if unmarshalErr := jsonutil.Unmarshal(data.Push, devicePush); unmarshalErr != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(unmarshalErr))
		}
	}
	deviceOnline, err := a.deviceHeartbeatRepo.IsDeviceOnline(ctx, req.GetSn())
	if err != nil {
		return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	resp.Info = &pb.DeviceInfo{
		Id:             data.ID,
		Sn:             data.Sn,
		Name:           data.Name,
		Desc:           data.Desc,
		Brand:          data.Brand,
		Model:          data.Model,
		Network:        data.Network,
		Imei:           data.Imei,
		CPU:            data.CPU,
		Mac:            data.Mac,
		AppVersion:     data.AppVersion,
		AndroidVersion: data.AndroidVersion,
		RAMSize:        data.RAMSize,
		DdrSize:        data.DdrSize,
		Certificate:    data.Certificate,
		SecureKey:      data.SecureKey,
		RegistryTime:   data.RegistryTime.Time.Format(time.RFC3339),
		Push:           devicePush,
		Status:         data.Status,
		CreatedAt:      data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      data.UpdatedAt.Format(time.RFC3339),
		Online:         deviceOnline,
	}
	return resp, nil
}
