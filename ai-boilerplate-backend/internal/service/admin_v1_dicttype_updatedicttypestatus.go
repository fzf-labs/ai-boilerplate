package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateDictTypeStatus 字典类型表-更新状态
func (a *AdminV1DictTypeService) UpdateDictTypeStatus(ctx context.Context, req *pb.UpdateDictTypeStatusReq) (*pb.UpdateDictTypeStatusReply, error) {
	resp := &pb.UpdateDictTypeStatusReply{}
	data, err := a.dictTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.dictTypeRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.dictTypeRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
