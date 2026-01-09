package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetWxGzhMaterialStats 公众号素材表-统计
func (a *AdminV1WxGzhMaterialService) GetWxGzhMaterialStats(ctx context.Context, req *pb.GetWxGzhMaterialStatsReq) (*pb.GetWxGzhMaterialStatsReply, error) {
	resp := &pb.GetWxGzhMaterialStatsReply{
		TotalCount: 0,
		ImageCount: 0,
		VoiceCount: 0,
		VideoCount: 0,
	}
	stats, err := a.wxGzhMaterialRepo.GetWxGzhMaterialStats(ctx, req.GetAppId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.ImageCount = stats["image"]
	resp.VoiceCount = stats["voice"]
	resp.VideoCount = stats["video"]
	resp.TotalCount = resp.ImageCount + resp.VoiceCount + resp.VideoCount
	return resp, nil
}
