package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiProviderModelStatus AI 配置模型表-更新状态
func (a *AdminV1AiProviderModelService) UpdateAiProviderModelStatus(ctx context.Context, req *pb.UpdateAiProviderModelStatusReq) (*pb.UpdateAiProviderModelStatusReply, error) {
	resp := &pb.UpdateAiProviderModelStatusReply{}
	data, err := a.aiProviderModelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiProviderModelRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiProviderModelRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
