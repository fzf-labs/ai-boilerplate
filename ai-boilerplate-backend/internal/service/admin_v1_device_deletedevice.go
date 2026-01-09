package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteDevice 设备表-删除一条数据
func (a *AdminV1DeviceService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceReq) (*pb.DeleteDeviceReply, error) {
	resp := &pb.DeleteDeviceReply{}
	err := a.deviceRepo.DeleteOneCacheBySn(ctx, req.GetSn())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
