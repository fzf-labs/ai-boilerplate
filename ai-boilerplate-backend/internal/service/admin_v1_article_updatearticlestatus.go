package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// UpdateArticleStatus 内容-文章-更新状态(发布/下线/草稿)
func (a *AdminV1ArticleService) UpdateArticleStatus(ctx context.Context, req *pb.UpdateArticleStatusReq) (*pb.UpdateArticleStatusReply, error) {
	resp := &pb.UpdateArticleStatusReply{}

	data, err := a.articleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	oldData := a.articleRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())

	if req.GetStatus() == 1 {
		if req.GetPublishTime() != "" {
			data.PublishTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetPublishTime()).StdTime())
		} else if !data.PublishTime.Valid {
			data.PublishTime = timeutil.TimeToSQLNullTime(time.Now())
		}
	} else {
		data.PublishTime = sql.NullTime{}
	}

	if err := a.articleRepo.UpdateOneCacheWithZero(ctx, data, oldData); err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	return resp, nil
}
