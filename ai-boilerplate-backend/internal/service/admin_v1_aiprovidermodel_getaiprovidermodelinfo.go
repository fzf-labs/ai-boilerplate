package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetAiProviderModelInfo AI 配置模型表-单条数据查询
func (a *AdminV1AiProviderModelService) GetAiProviderModelInfo(ctx context.Context, req *pb.GetAiProviderModelInfoReq) (*pb.GetAiProviderModelInfoReply, error) {
	resp := &pb.GetAiProviderModelInfoReply{}
	data, err := a.aiProviderModelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	modelConfig := &pb.ModelConfig{}
	err = jsonutil.Unmarshal(data.ModelConfig, modelConfig)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiProviderModelInfo{
		Id:          data.ID,
		PlatformId:  data.PlatformID,
		ModelType:   data.ModelType,
		ModelId:     data.ModelID,
		ModelName:   data.ModelName,
		ModelConfig: modelConfig,
		Sort:        data.Sort,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
