package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiIndexVideoRecord AI 视频表-更新一条数据
func (a *AdminV1AiIndexVideoService) UpdateAiIndexVideoRecord(_ context.Context, _ *pb.UpdateAiIndexVideoRecordReq) (*pb.UpdateAiIndexVideoRecordReply, error) {
	resp := &pb.UpdateAiIndexVideoRecordReply{}
	// TODO
	return resp, nil
}
