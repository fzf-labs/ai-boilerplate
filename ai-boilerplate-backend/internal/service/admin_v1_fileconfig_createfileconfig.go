package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// CreateFileConfig 文件配置表-创建一条数据
func (a *AdminV1FileConfigService) CreateFileConfig(ctx context.Context, req *pb.CreateFileConfigReq) (*pb.CreateFileConfigReply, error) {
	resp := &pb.CreateFileConfigReply{}
	configJSON, err := jsonutil.Marshal(req.GetConfig())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	data := a.fileConfigRepo.NewData()
	data.Name = req.GetName()
	data.Storage = req.GetStorage()
	data.Remark = req.GetRemark()
	data.Master = false
	data.Config = configJSON
	err = a.fileConfigRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
