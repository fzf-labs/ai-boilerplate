package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSmsTemplate 短信模板-删除一条数据
func (a *AdminV1SmsTemplateService) DeleteSmsTemplate(ctx context.Context, req *pb.DeleteSmsTemplateReq) (*pb.DeleteSmsTemplateReply, error) {
	resp := &pb.DeleteSmsTemplateReply{}
	err := a.smsTemplateRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
