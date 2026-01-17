package service

import (
	"context"
	"time"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/lib/pq"
)

// CreateArticle 内容-文章-创建一条数据
func (a *AdminV1ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleReq) (*pb.CreateArticleReply, error) {
	resp := &pb.CreateArticleReply{}

	data := a.articleRepo.NewData()
	data.Title = req.GetTitle()
	data.Summary = req.GetSummary()
	data.CoverImage = req.GetCoverImage()
	data.ContentMarkdown = req.GetContentMarkdown()
	data.Tags = pq.StringArray(req.GetTags())
	data.IsRecommend = req.GetIsRecommend()
	data.IsHot = req.GetIsHot()

	status := req.GetStatus()
	data.Status = int16(status)

	if status == 1 {
		if req.GetPublishTime() != "" {
			data.PublishTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetPublishTime()).StdTime())
		} else {
			data.PublishTime = timeutil.TimeToSQLNullTime(time.Now())
		}
	}

	if err := a.articleRepo.CreateOneCache(ctx, data); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Id = data.ID
	return resp, nil
}
