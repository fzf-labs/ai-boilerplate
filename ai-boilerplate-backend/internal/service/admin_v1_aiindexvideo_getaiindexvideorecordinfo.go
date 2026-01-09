package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiIndexVideoRecordInfo AI 视频表-单条数据查询
func (a *AdminV1AiIndexVideoService) GetAiIndexVideoRecordInfo(_ context.Context, _ *pb.GetAiIndexVideoRecordInfoReq) (*pb.GetAiIndexVideoRecordInfoReply, error) {
	resp := &pb.GetAiIndexVideoRecordInfoReply{}
	// TODO
	return resp, nil
}
