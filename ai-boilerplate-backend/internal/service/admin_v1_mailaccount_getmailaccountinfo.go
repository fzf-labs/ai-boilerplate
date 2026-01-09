package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetMailAccountInfo 邮箱账号表-单条数据查询
func (a *AdminV1MailAccountService) GetMailAccountInfo(ctx context.Context, req *pb.GetMailAccountInfoReq) (*pb.GetMailAccountInfoReply, error) {
	resp := &pb.GetMailAccountInfoReply{}
	data, err := a.mailAccountRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.MailAccountInfo{
		Id:        data.ID,
		Mail:      data.Mail,
		Username:  data.Username,
		Password:  data.Password,
		Host:      data.Host,
		Port:      data.Port,
		SslEnable: data.SslEnable,
		Status:    data.Status,
		Remark:    data.Remark,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
