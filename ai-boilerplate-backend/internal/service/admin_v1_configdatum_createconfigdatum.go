package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"gorm.io/datatypes"
)

// CreateConfigDatum 配置管理-创建一条数据
func (a *AdminV1ConfigDatumService) CreateConfigDatum(ctx context.Context, req *pb.CreateConfigDatumReq) (*pb.CreateConfigDatumReply, error) {
	resp := &pb.CreateConfigDatumReply{}
	data := a.configDatumRepo.NewData()
	data.Name = req.GetName()
	data.Key = req.GetKey()
	data.Value = datatypes.JSON(req.GetValue())
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err := a.configDatumRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
