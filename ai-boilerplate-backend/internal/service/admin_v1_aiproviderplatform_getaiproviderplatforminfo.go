package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiProviderPlatformInfo AI 配置平台表-单条数据查询
func (a *AdminV1AiProviderPlatformService) GetAiProviderPlatformInfo(ctx context.Context, req *pb.GetAiProviderPlatformInfoReq) (*pb.GetAiProviderPlatformInfoReply, error) {
	resp := &pb.GetAiProviderPlatformInfoReply{}
	data, err := a.aiProviderPlatformRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = aiProviderPlatformInfoFromModel(data)
	return resp, nil
}
