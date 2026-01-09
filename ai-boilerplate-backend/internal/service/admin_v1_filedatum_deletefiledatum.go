package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteFileDatum 文件表-删除一条数据
func (a *AdminV1FileDatumService) DeleteFileDatum(ctx context.Context, req *pb.DeleteFileDatumReq) (*pb.DeleteFileDatumReply, error) {
	resp := &pb.DeleteFileDatumReply{}
	err := a.fileDatumRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
