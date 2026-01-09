package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSensitiveWordList 敏感词-列表数据查询
func (a *AdminV1SensitiveWordService) GetSensitiveWordList(ctx context.Context, req *pb.GetSensitiveWordListReq) (*pb.GetSensitiveWordListReply, error) {
	resp := &pb.GetSensitiveWordListReply{
		Total: 0,
		List:  []*pb.SensitiveWordInfo{},
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
	list, p, err := a.sensitiveWordRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SensitiveWordInfo{
				Id:        v.ID,
				Word:      v.Word,
				Lab:       v.Lab,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
