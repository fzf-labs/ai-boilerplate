package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiIndexAudioRecordInfo AI 音乐表-单条数据查询
func (a *AdminV1AiIndexAudioService) GetAiIndexAudioRecordInfo(_ context.Context, _ *pb.GetAiIndexAudioRecordInfoReq) (*pb.GetAiIndexAudioRecordInfoReply, error) {
	resp := &pb.GetAiIndexAudioRecordInfoReply{}
	// TODO
	return resp, nil
}
