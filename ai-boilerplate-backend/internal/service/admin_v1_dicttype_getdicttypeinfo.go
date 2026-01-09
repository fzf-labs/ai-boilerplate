package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetDictTypeInfo 字典类型表-单条数据查询
func (a *AdminV1DictTypeService) GetDictTypeInfo(ctx context.Context, req *pb.GetDictTypeInfoReq) (*pb.GetDictTypeInfoReply, error) {
	resp := &pb.GetDictTypeInfoReply{}
	data, err := a.dictTypeRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.DictTypeInfo{
		Id:        data.ID,
		Name:      data.Name,
		Type:      data.Type,
		Status:    int32(data.Status),
		Remark:    data.Remark,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
