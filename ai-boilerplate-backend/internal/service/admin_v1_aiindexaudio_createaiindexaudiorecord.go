package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateAiIndexAudioRecord AI 音乐表-创建一条数据
func (a *AdminV1AiIndexAudioService) CreateAiIndexAudioRecord(_ context.Context, _ *pb.CreateAiIndexAudioRecordReq) (*pb.CreateAiIndexAudioRecordReply, error) {
	resp := &pb.CreateAiIndexAudioRecordReply{}
	// TODO
	return resp, nil
}
