package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiProviderPlatform AI 配置平台表-更新一条数据
func (a *AdminV1AiProviderPlatformService) UpdateAiProviderPlatform(ctx context.Context, req *pb.UpdateAiProviderPlatformReq) (*pb.UpdateAiProviderPlatformReply, error) {
	resp := &pb.UpdateAiProviderPlatformReply{}
	data, err := a.aiProviderPlatformRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiProviderPlatformRepo.DeepCopy(data)
	data.Platform = req.GetPlatform()
	data.Name = req.GetName()
	data.APIURL = req.GetAPIURL()
	data.APIKey = req.GetAPIKey()
	data.DocURL = req.GetDocURL()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err = a.aiProviderPlatformRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
