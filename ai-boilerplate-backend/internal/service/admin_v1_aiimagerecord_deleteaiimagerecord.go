package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiImageRecord AI 绘画表-删除一条数据
func (a *AdminV1AiImageRecordService) DeleteAiImageRecord(ctx context.Context, req *pb.DeleteAiImageRecordReq) (*pb.DeleteAiImageRecordReply, error) {
	resp := &pb.DeleteAiImageRecordReply{}
	err := a.aiImageRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
