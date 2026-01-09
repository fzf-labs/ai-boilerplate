package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMallActivationCode 激活码管理表-删除一条数据
func (a *AdminV1MallActivationCodeService) DeleteMallActivationCode(ctx context.Context, req *pb.DeleteMallActivationCodeReq) (*pb.DeleteMallActivationCodeReply, error) {
	resp := &pb.DeleteMallActivationCodeReply{}
	err := a.mallActivationCodeRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
