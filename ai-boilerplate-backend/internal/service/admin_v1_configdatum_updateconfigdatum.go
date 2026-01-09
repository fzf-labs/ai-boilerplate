package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"gorm.io/datatypes"
)

// UpdateConfigDatum 配置管理-更新一条数据
func (a *AdminV1ConfigDatumService) UpdateConfigDatum(ctx context.Context, req *pb.UpdateConfigDatumReq) (*pb.UpdateConfigDatumReply, error) {
	resp := &pb.UpdateConfigDatumReply{}
	data, err := a.configDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.configDatumRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Key = req.GetKey()
	data.Value = datatypes.JSON(req.GetValue())
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err = a.configDatumRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
