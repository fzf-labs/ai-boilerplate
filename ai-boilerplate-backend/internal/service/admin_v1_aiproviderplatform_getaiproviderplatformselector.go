package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiProviderPlatformSelector AI 配置平台表-获取平台选择器
func (a *AdminV1AiProviderPlatformService) GetAiProviderPlatformSelector(_ context.Context, _ *pb.GetAiProviderPlatformSelectorReq) (*pb.GetAiProviderPlatformSelectorReply, error) {
	resp := &pb.GetAiProviderPlatformSelectorReply{
		List: []*pb.AiProviderPlatformSelectorItem{
			{
				Label: "OpenAI",
				Value: "openai",
			},
			{
				Label: "Gemini",
				Value: "gemini",
			},
			{
				Label: "Claude",
				Value: "claude",
			},
			{
				Label: "阿里云百炼",
				Value: "aliyun",
			},
			{
				Label: "火山引擎",
				Value: "volcengine",
			},
		},
	}
	return resp, nil
}
