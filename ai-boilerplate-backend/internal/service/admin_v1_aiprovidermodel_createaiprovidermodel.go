package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateAiProviderModel AI 配置模型表-创建一条数据
func (a *AdminV1AiProviderModelService) CreateAiProviderModel(ctx context.Context, req *pb.CreateAiProviderModelReq) (*pb.CreateAiProviderModelReply, error) {
	resp := &pb.CreateAiProviderModelReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	modelConfig, err := jsonutil.Marshal(req.GetModelConfig())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	data := a.aiProviderModelRepo.NewData()
	data.TenantID = tenantID
	data.PlatformID = req.GetPlatformId()
	data.ModelType = req.GetModelType()
	data.ModelID = req.GetModelId()
	data.ModelName = req.GetModelName()
	data.ModelConfig = modelConfig
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.aiProviderModelRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
