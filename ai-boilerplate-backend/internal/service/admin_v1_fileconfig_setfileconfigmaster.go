package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// SetFileConfigMaster 文件配置表-设置主配置
func (a *AdminV1FileConfigService) SetFileConfigMaster(ctx context.Context, req *pb.SetFileConfigMasterReq) (*pb.SetFileConfigMasterReply, error) {
	resp := &pb.SetFileConfigMasterReply{}
	// 查询所有的配置
	allConfigs, _, err := a.fileConfigRepo.FindMultiCacheByCondition(ctx, &condition.Req{})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	// 设置主配置
	for _, config := range allConfigs {
		oldData := a.fileConfigRepo.DeepCopy(config)
		if config.ID == req.GetId() {
			config.Master = true
		} else {
			config.Master = false
		}
		err = a.fileConfigRepo.UpdateOneCacheWithZero(ctx, config, oldData)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
	}
	return resp, nil
}
