package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// ListHelpCategories 帮助分类表-列表数据查询
func (a *AppV1HelpCategoryService) ListHelpCategories(ctx context.Context, _ *pb.ListHelpCategoriesReq) (*pb.ListHelpCategoriesReply, error) {
	resp := &pb.ListHelpCategoriesReply{}

	// 查询所有启用的分类
	categories, err := a.helpCategoryRepo.FindAllEnabled(ctx)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	// 转换为响应格式
	resp.List = make([]*pb.HelpCategoryInfo, 0, len(categories))
	for _, category := range categories {
		resp.List = append(resp.List, &pb.HelpCategoryInfo{
			Id:        category.ID,
			Name:      category.Name,
			Icon:      category.Icon,
			Order:     category.Order,
			Status:    category.Status,
			CreatedAt: category.CreatedAt.Format(time.RFC3339),
			UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
		})
	}

	return resp, nil
}
