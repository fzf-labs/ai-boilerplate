package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiProviderModel AI 配置模型表-删除一条数据
func (a *AdminV1AiProviderModelService) DeleteAiProviderModel(ctx context.Context, req *pb.DeleteAiProviderModelReq) (*pb.DeleteAiProviderModelReply, error) {
	resp := &pb.DeleteAiProviderModelReply{}
	err := a.aiProviderModelRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
