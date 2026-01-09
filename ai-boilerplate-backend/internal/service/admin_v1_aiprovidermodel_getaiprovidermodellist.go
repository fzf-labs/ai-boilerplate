package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
)

// GetAiProviderModelList AI 配置模型表-列表数据查询
func (a *AdminV1AiProviderModelService) GetAiProviderModelList(ctx context.Context, req *pb.GetAiProviderModelListReq) (*pb.GetAiProviderModelListReply, error) {
	resp := &pb.GetAiProviderModelListReply{
		Total: 0,
		List:  []*pb.AiProviderModelInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.aiProviderModelRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			modelConfig := &pb.ModelConfig{}
			err = jsonutil.Unmarshal(v.ModelConfig, modelConfig)
			if err != nil {
				return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
			}
			resp.List = append(resp.List, &pb.AiProviderModelInfo{
				Id:          v.ID,
				PlatformId:  v.PlatformID,
				ModelType:   v.ModelType,
				ModelId:     v.ModelID,
				ModelName:   v.ModelName,
				ModelConfig: modelConfig,
				Sort:        v.Sort,
				Status:      v.Status,
				CreatedAt:   v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
