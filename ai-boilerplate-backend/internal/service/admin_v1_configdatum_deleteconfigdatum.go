package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteConfigDatum 配置管理-删除一条数据
func (a *AdminV1ConfigDatumService) DeleteConfigDatum(ctx context.Context, req *pb.DeleteConfigDatumReq) (*pb.DeleteConfigDatumReply, error) {
	resp := &pb.DeleteConfigDatumReply{}
	err := a.configDatumRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
