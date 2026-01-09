package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// GetContentList 获取内容列表
func (s *AppV1HomeService) GetContentList(_ context.Context, req *pb.GetContentListReq) (*pb.GetContentListReply, error) {
	resp := &pb.GetContentListReply{}
	// TODO: 从数据库获取内容列表数据
	// 这里返回示例数据
	resp.List = []*pb.ContentInfo{
		{
			Id:          1,
			Title:       "示例内容标题1",
			Summary:     "这是示例内容的摘要信息",
			CoverImage:  "https://example.com/cover1.jpg",
			PublishTime: "2024-01-01 10:00:00",
			Tags:        []string{"推荐", "热门"},
			IsRecommend: true,
			IsHot:       true,
		},
		{
			Id:          2,
			Title:       "示例内容标题2",
			Summary:     "这是另一个示例内容的摘要信息",
			CoverImage:  "https://example.com/cover2.jpg",
			PublishTime: "2024-01-02 10:00:00",
			Tags:        []string{"新闻"},
			IsRecommend: false,
			IsHot:       false,
		},
	}
	resp.Total = 2
	resp.Page = req.GetPage()
	resp.PageSize = req.GetPageSize()
	return resp, nil
}
