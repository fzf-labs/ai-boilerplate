package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetContentList 获取内容列表
func (s *AppV1HomeService) GetContentList(ctx context.Context, req *pb.GetContentListReq) (*pb.GetContentListReply, error) {
	resp := &pb.GetContentListReply{
		List:     []*pb.ContentInfo{},
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	}

	now := time.Now()
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "status",
				Value: 1,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "publish_time",
				Value: timeutil.RFC3339(now),
				Exp:   condition.LTE,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "publish_time",
				Order: condition.DESC,
			},
		},
	}

	list, p, err := s.articleRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	resp.Total = int64(p.Total)
	for _, v := range list {
		publishTime := ""
		if v.PublishTime.Valid {
			publishTime = timeutil.RFC3339(v.PublishTime.Time)
		}
		resp.List = append(resp.List, &pb.ContentInfo{
			Id:          v.ID,
			Title:       v.Title,
			Summary:     v.Summary,
			CoverImage:  v.CoverImage,
			PublishTime: publishTime,
			Tags:        []string(v.Tags),
			IsRecommend: v.IsRecommend,
			IsHot:       v.IsHot,
		})
	}

	resp.Page = req.GetPage()
	resp.PageSize = req.GetPageSize()
	return resp, nil
}
