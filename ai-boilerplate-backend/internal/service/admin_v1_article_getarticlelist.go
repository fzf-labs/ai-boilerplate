package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetArticleList 内容-文章-列表数据查询
func (a *AdminV1ArticleService) GetArticleList(ctx context.Context, req *pb.GetArticleListReq) (*pb.GetArticleListReply, error) {
	resp := &pb.GetArticleListReply{
		Total: 0,
		List:  []*pb.ArticleInfo{},
	}

	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    make([]*condition.QueryParam, 0),
		Order: []*condition.OrderParam{
			{Field: "publish_time", Order: condition.DESC},
			{Field: "created_at", Order: condition.DESC},
		},
	}

	if req.Status != nil {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}

	if req.GetKeyword() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "title",
			Value: "%" + req.GetKeyword() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}

	list, p, err := a.articleRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Total = p.Total
	for _, v := range list {
		publishTime := ""
		if v.PublishTime.Valid {
			publishTime = timeutil.RFC3339(v.PublishTime.Time)
		}
		resp.List = append(resp.List, &pb.ArticleInfo{
			Id:              v.ID,
			Title:           v.Title,
			Summary:         v.Summary,
			CoverImage:      v.CoverImage,
			ContentMarkdown: v.ContentMarkdown,
			Status:          int32(v.Status),
			PublishTime:     publishTime,
			Tags:            []string(v.Tags),
			IsRecommend:     v.IsRecommend,
			IsHot:           v.IsHot,
			ViewCount:       v.ViewCount,
			LikeCount:       v.LikeCount,
			CreatedAt:       v.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       v.UpdatedAt.Format(time.RFC3339),
		})
	}

	return resp, nil
}
