package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetFileConfigInfo 文件配置表-单条数据查询
func (a *AdminV1FileConfigService) GetFileConfigInfo(ctx context.Context, req *pb.GetFileConfigInfoReq) (*pb.GetFileConfigInfoReply, error) {
	resp := &pb.GetFileConfigInfoReply{}
	data, err := a.fileConfigRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	config := &pb.StorageConfig{}
	if data.Config.String() != "" {
		err = jsonutil.Unmarshal(data.Config, config)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	resp.Info = &pb.FileConfigInfo{
		Id:        data.ID,
		Name:      data.Name,
		Storage:   data.Storage,
		Remark:    data.Remark,
		Master:    data.Master,
		Config:    config,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
