package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiVideoRecord AI 视频表-删除一条数据
func (a *AdminV1AiVideoRecordService) DeleteAiVideoRecord(ctx context.Context, req *pb.DeleteAiVideoRecordReq) (*pb.DeleteAiVideoRecordReply, error) {
	resp := &pb.DeleteAiVideoRecordReply{}
	err := a.aiVideoRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
