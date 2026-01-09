package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// UpdateFileConfig 文件配置表-更新一条数据
func (a *AdminV1FileConfigService) UpdateFileConfig(ctx context.Context, req *pb.UpdateFileConfigReq) (*pb.UpdateFileConfigReply, error) {
	resp := &pb.UpdateFileConfigReply{}
	data, err := a.fileConfigRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	configJSON, err := jsonutil.Marshal(req.GetConfig())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	oldData := a.fileConfigRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Storage = req.GetStorage()
	data.Remark = req.GetRemark()
	data.Config = configJSON
	err = a.fileConfigRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
