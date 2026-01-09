package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiIndexImageRecord AI 绘画表-删除一条数据
func (a *AdminV1AiIndexImageService) DeleteAiIndexImageRecord(_ context.Context, _ *pb.DeleteAiIndexImageRecordReq) (*pb.DeleteAiIndexImageRecordReply, error) {
	resp := &pb.DeleteAiIndexImageRecordReply{}
	// TODO
	return resp, nil
}
