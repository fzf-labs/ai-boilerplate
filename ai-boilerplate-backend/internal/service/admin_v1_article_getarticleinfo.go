package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetArticleInfo 内容-文章-单条数据查询
func (a *AdminV1ArticleService) GetArticleInfo(ctx context.Context, req *pb.GetArticleInfoReq) (*pb.GetArticleInfoReply, error) {
	resp := &pb.GetArticleInfoReply{}

	data, err := a.articleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	publishTime := ""
	if data.PublishTime.Valid {
		publishTime = timeutil.RFC3339(data.PublishTime.Time)
	}

	resp.Info = &pb.ArticleInfo{
		Id:              data.ID,
		Title:           data.Title,
		Summary:         data.Summary,
		CoverImage:      data.CoverImage,
		ContentMarkdown: data.ContentMarkdown,
		Status:          int32(data.Status),
		PublishTime:     publishTime,
		Tags:            []string(data.Tags),
		IsRecommend:     data.IsRecommend,
		IsHot:           data.IsHot,
		ViewCount:       data.ViewCount,
		LikeCount:       data.LikeCount,
		CreatedAt:       timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:       timeutil.RFC3339(data.UpdatedAt),
	}

	return resp, nil
}
