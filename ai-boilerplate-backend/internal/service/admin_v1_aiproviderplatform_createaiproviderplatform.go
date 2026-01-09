package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateAiProviderPlatform AI 配置平台表-创建一条数据
func (a *AdminV1AiProviderPlatformService) CreateAiProviderPlatform(ctx context.Context, req *pb.CreateAiProviderPlatformReq) (*pb.CreateAiProviderPlatformReply, error) {
	resp := &pb.CreateAiProviderPlatformReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.aiProviderPlatformRepo.NewData()
	data.TenantID = tenantID
	data.Platform = req.GetPlatform()
	data.Name = req.GetName()
	data.APIURL = req.GetAPIURL()
	data.APIKey = req.GetAPIKey()
	data.DocURL = req.GetDocURL()
	data.Sort = req.GetSort()
	data.Status = req.GetStatus()
	err := a.aiProviderPlatformRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
