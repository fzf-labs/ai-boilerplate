package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteAiProviderPlatform AI 配置平台表-删除一条数据
func (a *AdminV1AiProviderPlatformService) DeleteAiProviderPlatform(ctx context.Context, req *pb.DeleteAiProviderPlatformReq) (*pb.DeleteAiProviderPlatformReply, error) {
	resp := &pb.DeleteAiProviderPlatformReply{}
	err := a.aiProviderPlatformRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
