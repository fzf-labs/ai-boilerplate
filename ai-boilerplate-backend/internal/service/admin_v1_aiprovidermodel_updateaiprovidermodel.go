package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// UpdateAiProviderModel AI 配置模型表-更新一条数据
func (a *AdminV1AiProviderModelService) UpdateAiProviderModel(ctx context.Context, req *pb.UpdateAiProviderModelReq) (*pb.UpdateAiProviderModelReply, error) {
	resp := &pb.UpdateAiProviderModelReply{}
	data, err := a.aiProviderModelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	modelConfig, err := jsonutil.Marshal(req.GetModelConfig())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	oldData := a.aiProviderModelRepo.DeepCopy(data)
	data.PlatformID = req.GetPlatformId()
	data.ModelType = req.GetModelType()
	data.ModelID = req.GetModelId()
	data.ModelName = req.GetModelName()
	data.ModelConfig = modelConfig
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.aiProviderModelRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
