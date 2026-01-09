package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateDictType 字典类型表-创建一条数据
func (a *AdminV1DictTypeService) CreateDictType(ctx context.Context, req *pb.CreateDictTypeReq) (*pb.CreateDictTypeReply, error) {
	resp := &pb.CreateDictTypeReply{}
	data := a.dictTypeRepo.NewData()
	data.Type = req.GetType()
	data.Name = req.GetName()
	data.Status = int16(req.GetStatus())
	data.Remark = req.GetRemark()
	err := a.dictTypeRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
