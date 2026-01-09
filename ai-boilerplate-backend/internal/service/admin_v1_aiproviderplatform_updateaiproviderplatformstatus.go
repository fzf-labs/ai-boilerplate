package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiProviderPlatformStatus AI 配置平台表-更新状态
func (a *AdminV1AiProviderPlatformService) UpdateAiProviderPlatformStatus(ctx context.Context, req *pb.UpdateAiProviderPlatformStatusReq) (*pb.UpdateAiProviderPlatformStatusReply, error) {
	resp := &pb.UpdateAiProviderPlatformStatusReply{}
	data, err := a.aiProviderPlatformRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiProviderPlatformRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiProviderPlatformRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
