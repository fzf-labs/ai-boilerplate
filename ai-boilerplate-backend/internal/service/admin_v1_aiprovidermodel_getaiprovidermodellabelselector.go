package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiProviderModelLabelSelector AI 配置模型表-标签选择器
func (a *AdminV1AiProviderModelService) GetAiProviderModelLabelSelector(_ context.Context, _ *pb.GetAiProviderModelLabelSelectorReq) (*pb.GetAiProviderModelLabelSelectorReply, error) {
	resp := &pb.GetAiProviderModelLabelSelectorReply{
		List: []*pb.AiProviderModelLabelSelectorItem{
			{
				Label: "OpenAI",
				Value: "openai",
			},
		},
	}
	return resp, nil
}
