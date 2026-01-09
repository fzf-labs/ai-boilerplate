package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSensitiveWord 敏感词-删除一条数据
func (a *AdminV1SensitiveWordService) DeleteSensitiveWord(ctx context.Context, req *pb.DeleteSensitiveWordReq) (*pb.DeleteSensitiveWordReply, error) {
	resp := &pb.DeleteSensitiveWordReply{}
	err := a.sensitiveWordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
