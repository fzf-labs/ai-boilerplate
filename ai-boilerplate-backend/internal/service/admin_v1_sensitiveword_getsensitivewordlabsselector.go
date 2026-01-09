package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSensitiveWordLabsSelector 敏感词-标签选择器
func (a *AdminV1SensitiveWordService) GetSensitiveWordLabsSelector(_ context.Context, _ *pb.GetSensitiveWordLabsSelectorReq) (*pb.GetSensitiveWordLabsSelectorReply, error) {
	resp := &pb.GetSensitiveWordLabsSelectorReply{
		List: []*pb.SensitiveWordLabsSelector{
			{
				Key:   "politics",
				Value: "政治",
			},
			{
				Key:   "sex",
				Value: "性",
			},
			{
				Key:   "age",
				Value: "年龄",
			},
			{
				Key:   "education",
				Value: "教育",
			},
			{
				Key:   "occupation",
				Value: "职业",
			},
			{
				Key:   "income",
				Value: "收入",
			},
			{
				Key:   "other",
				Value: "其他",
			},
		},
	}
	return resp, nil
}
