package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateDictType 字典类型表-更新一条数据
func (a *AdminV1DictTypeService) UpdateDictType(ctx context.Context, req *pb.UpdateDictTypeReq) (*pb.UpdateDictTypeReply, error) {
	resp := &pb.UpdateDictTypeReply{}
	data, err := a.dictTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.dictTypeRepo.DeepCopy(data)
	data.Type = req.GetType()
	data.Name = req.GetName()
	data.Status = int16(req.GetStatus())
	data.Remark = req.GetRemark()
	err = a.dictTypeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
