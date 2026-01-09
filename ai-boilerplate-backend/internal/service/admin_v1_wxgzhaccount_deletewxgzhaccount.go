package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteWxGzhAccount 公众号账号表-删除一条数据
func (a *AdminV1WxGzhAccountService) DeleteWxGzhAccount(ctx context.Context, req *pb.DeleteWxGzhAccountReq) (*pb.DeleteWxGzhAccountReply, error) {
	resp := &pb.DeleteWxGzhAccountReply{}
	err := a.wxGzhAccountRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
