package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiIndexAudioRecordList AI 音乐表-列表数据查询
func (a *AdminV1AiIndexAudioService) GetAiIndexAudioRecordList(_ context.Context, _ *pb.GetAiIndexAudioRecordListReq) (*pb.GetAiIndexAudioRecordListReply, error) {
	resp := &pb.GetAiIndexAudioRecordListReply{}
	// TODO
	return resp, nil
}
