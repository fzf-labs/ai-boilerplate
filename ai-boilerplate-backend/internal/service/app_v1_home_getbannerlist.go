package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

// GetBannerList 获取轮播图列表
func (s *AppV1HomeService) GetBannerList(_ context.Context, _ *pb.GetBannerListReq) (*pb.GetBannerListReply, error) {
	resp := &pb.GetBannerListReply{}
	// TODO: 从数据库或配置中获取轮播图数据
	// 这里返回示例数据
	resp.List = []*pb.BannerInfo{
		{
			Id:       1,
			ImageUrl: "https://example.com/banner1.jpg",
			LinkUrl:  "/pages/detail/index?id=1",
			Title:    "欢迎使用",
			Sort:     1,
		},
		{
			Id:       2,
			ImageUrl: "https://example.com/banner2.jpg",
			LinkUrl:  "/pages/detail/index?id=2",
			Title:    "新功能上线",
			Sort:     2,
		},
	}
	return resp, nil
}
