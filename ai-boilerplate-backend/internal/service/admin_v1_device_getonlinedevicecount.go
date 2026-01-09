package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetOnlineDeviceCount 设备表-在线设备数量统计
func (a *AdminV1DeviceService) GetOnlineDeviceCount(ctx context.Context, _ *pb.GetOnlineDeviceCountReq) (*pb.GetOnlineDeviceCountReply, error) {
	resp := &pb.GetOnlineDeviceCountReply{}
	// 从 Redis 获取在线设备数量
	count, err := a.deviceHeartbeatRepo.GetOnlineDeviceCount(ctx)
	if err != nil {
		return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	resp.Count = count
	return resp, nil
}
