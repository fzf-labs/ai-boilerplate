package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiProviderPlatformList AI 配置平台表-列表数据查询
func (a *AdminV1AiProviderPlatformService) GetAiProviderPlatformList(ctx context.Context, req *pb.GetAiProviderPlatformListReq) (*pb.GetAiProviderPlatformListReply, error) {
	resp := &pb.GetAiProviderPlatformListReply{
		Total: 0,
		List:  []*pb.AiProviderPlatformInfo{},
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
	if req.GetPlatform() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "platform",
			Value: req.GetPlatform(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.aiProviderPlatformRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiProviderPlatformInfo{
				Id:        v.ID,
				Platform:  v.Platform,
				Name:      v.Name,
				APIURL:    v.APIURL,
				APIKey:    v.APIKey,
				DocURL:    v.DocURL,
				Sort:      v.Sort,
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
