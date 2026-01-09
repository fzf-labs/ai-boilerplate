package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMailTemplate 邮件模版表-删除一条数据
func (a *AdminV1MailTemplateService) DeleteMailTemplate(ctx context.Context, req *pb.DeleteMailTemplateReq) (*pb.DeleteMailTemplateReply, error) {
	resp := &pb.DeleteMailTemplateReply{}
	err := a.mailTemplateRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
