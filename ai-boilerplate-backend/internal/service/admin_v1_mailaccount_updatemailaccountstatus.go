package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMailAccountStatus 邮箱账号表-更新状态
func (a *AdminV1MailAccountService) UpdateMailAccountStatus(ctx context.Context, req *pb.UpdateMailAccountStatusReq) (*pb.UpdateMailAccountStatusReply, error) {
	resp := &pb.UpdateMailAccountStatusReply{}
	data, err := a.mailAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mailAccountRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.mailAccountRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
