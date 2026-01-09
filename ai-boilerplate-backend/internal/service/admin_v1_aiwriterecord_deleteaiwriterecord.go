package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiWriteRecord AI 写作表-删除一条数据
func (a *AdminV1AiWriteRecordService) DeleteAiWriteRecord(ctx context.Context, req *pb.DeleteAiWriteRecordReq) (*pb.DeleteAiWriteRecordReply, error) {
	resp := &pb.DeleteAiWriteRecordReply{}
	err := a.aiWriteRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
