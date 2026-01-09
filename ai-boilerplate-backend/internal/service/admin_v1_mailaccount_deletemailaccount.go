package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMailAccount 邮箱账号表-删除一条数据
func (a *AdminV1MailAccountService) DeleteMailAccount(ctx context.Context, req *pb.DeleteMailAccountReq) (*pb.DeleteMailAccountReply, error) {
	resp := &pb.DeleteMailAccountReply{}
	err := a.mailAccountRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
