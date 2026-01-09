package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateMailAccount 邮箱账号表-更新一条数据
func (a *AdminV1MailAccountService) UpdateMailAccount(ctx context.Context, req *pb.UpdateMailAccountReq) (*pb.UpdateMailAccountReply, error) {
	resp := &pb.UpdateMailAccountReply{}
	data, err := a.mailAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mailAccountRepo.DeepCopy(data)
	data.Mail = req.GetMail()
	data.Username = req.GetUsername()
	data.Password = req.GetPassword()
	data.Host = req.GetHost()
	data.Port = req.GetPort()
	data.SslEnable = req.GetSslEnable()
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err = a.mailAccountRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
