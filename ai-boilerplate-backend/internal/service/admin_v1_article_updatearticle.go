package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/lib/pq"
)

// UpdateArticle 内容-文章-更新一条数据
func (a *AdminV1ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleReq) (*pb.UpdateArticleReply, error) {
	resp := &pb.UpdateArticleReply{}

	data, err := a.articleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	oldData := a.articleRepo.DeepCopy(data)
	data.Title = req.GetTitle()
	data.Summary = req.GetSummary()
	data.CoverImage = req.GetCoverImage()
	data.ContentMarkdown = req.GetContentMarkdown()
	data.Tags = pq.StringArray(req.GetTags())
	data.IsRecommend = req.GetIsRecommend()
	data.IsHot = req.GetIsHot()

	if err := a.articleRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
