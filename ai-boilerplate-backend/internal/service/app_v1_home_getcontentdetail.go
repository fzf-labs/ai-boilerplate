package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// GetContentDetail 获取内容详情
func (s *AppV1HomeService) GetContentDetail(_ context.Context, req *pb.GetContentDetailReq) (*pb.GetContentDetailReply, error) {
	resp := &pb.GetContentDetailReply{}
	// TODO: 从数据库获取内容详情数据
	// 这里返回示例数据
	resp.Info = &pb.ContentDetail{
		Id:          req.GetId(),
		Title:       "示例内容标题",
		Content:     "<p>这是示例内容的详细信息，支持HTML格式。</p>",
		CoverImage:  "https://example.com/cover.jpg",
		PublishTime: "2024-01-01 10:00:00",
		Tags:        []string{"推荐", "热门"},
		ViewCount:   100,
		LikeCount:   50,
	}
	return resp, nil
}
