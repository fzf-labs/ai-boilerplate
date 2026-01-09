package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetAiPromptList AI 提示词-列表数据查询
func (a *AdminV1AiPromptService) GetAiPromptList(ctx context.Context, req *pb.GetAiPromptListReq) (*pb.GetAiPromptListReply, error) {
	resp := &pb.GetAiPromptListReply{
		Total: 0,
		List:  []*pb.AiPromptInfo{},
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
	list, p, err := a.aiPromptRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.AiPromptInfo{
				Id:        v.ID,
				TenantId:  v.TenantID,
				AdminId:   v.AdminID,
				Name:      v.Name,
				Desc:      v.Desc,
				Prompt:    v.Prompt,
				Sort:      v.Sort,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
