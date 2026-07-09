package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetMailAccountInfo 邮箱账号表-单条数据查询
func (a *AdminV1MailAccountService) GetMailAccountInfo(ctx context.Context, req *pb.GetMailAccountInfoReq) (*pb.GetMailAccountInfoReply, error) {
	resp := &pb.GetMailAccountInfoReply{}
	data, err := a.mailAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = mailAccountInfoFromModel(data)
	return resp, nil
}
