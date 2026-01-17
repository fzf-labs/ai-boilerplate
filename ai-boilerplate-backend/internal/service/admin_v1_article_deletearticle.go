package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteArticle 内容-文章-删除一条数据
func (a *AdminV1ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.DeleteArticleReply, error) {
	resp := &pb.DeleteArticleReply{}

	if err := a.articleRepo.DeleteOneCacheByID(ctx, req.GetId()); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
