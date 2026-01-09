package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateMailAccount 邮箱账号表-创建一条数据
func (a *AdminV1MailAccountService) CreateMailAccount(ctx context.Context, req *pb.CreateMailAccountReq) (*pb.CreateMailAccountReply, error) {
	resp := &pb.CreateMailAccountReply{}
	data := a.mailAccountRepo.NewData()
	data.Mail = req.GetMail()
	data.Username = req.GetUsername()
	data.Password = req.GetPassword()
	data.Host = req.GetHost()
	data.Port = req.GetPort()
	data.SslEnable = req.GetSslEnable()
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err := a.mailAccountRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
