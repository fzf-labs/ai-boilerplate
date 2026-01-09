package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// ListHelpFaqs 帮助中心常见问题表-列表数据查询
func (a *AppV1HelpFaqService) ListHelpFaqs(ctx context.Context, req *pb.ListHelpFaqsReq) (*pb.ListHelpFaqsReply, error) {
	resp := &pb.ListHelpFaqsReply{
		Total: 0,
		List:  []*pb.HelpFaqInfo{},
	}

	// 构建查询条件
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    make([]*condition.QueryParam, 0),
		Order: []*condition.OrderParam{
			{Field: "order_num", Order: condition.ASC},
		},
	}

	// 只查询启用的FAQ
	param.Query = append(param.Query, &condition.QueryParam{
		Field: "status",
		Value: 1,
		Exp:   condition.EQ,
		Logic: condition.AND,
	})

	// 如果有分类ID筛选
	if req.GetCategoryId() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "category_id",
			Value: req.GetCategoryId(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	// 如果有关键词搜索
	if req.GetKeyword() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "question",
			Value: "%" + req.GetKeyword() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}

	// 查询数据
	list, p, err := a.helpFaqRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Total = p.Total
	for _, v := range list {
		resp.List = append(resp.List, &pb.HelpFaqInfo{
			Id:             v.ID,
			CategoryId:     v.CategoryID,
			Question:       v.Question,
			Answer:         v.Answer,
			OrderNum:       v.OrderNum,
			ViewCount:      v.ViewCount,
			HelpfulCount:   v.HelpfulCount,
			UnhelpfulCount: v.UnhelpfulCount,
			IsHot:          v.IsHot,
			Status:         v.Status,
			CreatedAt:      v.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      v.UpdatedAt.Format(time.RFC3339),
		})
	}

	return resp, nil
}
