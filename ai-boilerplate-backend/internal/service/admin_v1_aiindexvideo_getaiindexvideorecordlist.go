package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiIndexVideoRecordList AI 视频表-列表数据查询
func (a *AdminV1AiIndexVideoService) GetAiIndexVideoRecordList(_ context.Context, _ *pb.GetAiIndexVideoRecordListReq) (*pb.GetAiIndexVideoRecordListReply, error) {
	resp := &pb.GetAiIndexVideoRecordListReply{}
	// TODO
	return resp, nil
}
