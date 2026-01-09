package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiIndexVideoRecordStatus AI 视频表-更新状态
func (a *AdminV1AiIndexVideoService) UpdateAiIndexVideoRecordStatus(_ context.Context, _ *pb.UpdateAiIndexVideoRecordStatusReq) (*pb.UpdateAiIndexVideoRecordStatusReply, error) {
	resp := &pb.UpdateAiIndexVideoRecordStatusReply{}
	// TODO
	return resp, nil
}
