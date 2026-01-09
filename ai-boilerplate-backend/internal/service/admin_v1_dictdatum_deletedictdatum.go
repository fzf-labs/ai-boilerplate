package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteDictDatum 字典数据表-删除一条数据
func (a *AdminV1DictDatumService) DeleteDictDatum(ctx context.Context, req *pb.DeleteDictDatumReq) (*pb.DeleteDictDatumReply, error) {
	resp := &pb.DeleteDictDatumReply{}
	err := a.dictDatumRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
