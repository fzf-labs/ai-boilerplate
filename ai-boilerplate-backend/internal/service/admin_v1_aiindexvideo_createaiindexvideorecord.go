package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateAiIndexVideoRecord AI 视频表-创建一条数据
func (a *AdminV1AiIndexVideoService) CreateAiIndexVideoRecord(_ context.Context, _ *pb.CreateAiIndexVideoRecordReq) (*pb.CreateAiIndexVideoRecordReply, error) {
	resp := &pb.CreateAiIndexVideoRecordReply{}
	// TODO
	return resp, nil
}
