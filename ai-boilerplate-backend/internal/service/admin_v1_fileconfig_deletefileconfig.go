package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteFileConfig 文件配置表-删除一条数据
func (a *AdminV1FileConfigService) DeleteFileConfig(ctx context.Context, req *pb.DeleteFileConfigReq) (*pb.DeleteFileConfigReply, error) {
	resp := &pb.DeleteFileConfigReply{}
	err := a.fileConfigRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
