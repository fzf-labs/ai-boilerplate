package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiAudioRecord AI 音乐表-删除一条数据
func (a *AdminV1AiAudioRecordService) DeleteAiAudioRecord(ctx context.Context, req *pb.DeleteAiAudioRecordReq) (*pb.DeleteAiAudioRecordReply, error) {
	resp := &pb.DeleteAiAudioRecordReply{}
	err := a.aiAudioRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
