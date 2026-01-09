package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteDictType 字典类型表-删除一条数据
func (a *AdminV1DictTypeService) DeleteDictType(ctx context.Context, req *pb.DeleteDictTypeReq) (*pb.DeleteDictTypeReply, error) {
	resp := &pb.DeleteDictTypeReply{}
	err := a.dictTypeRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
