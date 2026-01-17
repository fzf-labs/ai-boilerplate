package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetContentDetail 获取内容详情
func (s *AppV1HomeService) GetContentDetail(ctx context.Context, req *pb.GetContentDetailReq) (*pb.GetContentDetailReply, error) {
	resp := &pb.GetContentDetailReply{}

	data, err := s.articleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if data.Status != 1 {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if !data.PublishTime.Valid || data.PublishTime.Time.After(time.Now()) {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}

	contentHTML, err := renderMarkdownToSafeHTML(data.ContentMarkdown)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	publishTime := ""
	if data.PublishTime.Valid {
		publishTime = timeutil.RFC3339(data.PublishTime.Time)
	}

	resp.Info = &pb.ContentDetail{
		Id:          data.ID,
		Title:       data.Title,
		Content:     contentHTML,
		CoverImage:  data.CoverImage,
		PublishTime: publishTime,
		Tags:        []string(data.Tags),
		ViewCount:   data.ViewCount,
		LikeCount:   data.LikeCount,
	}
	return resp, nil
}
