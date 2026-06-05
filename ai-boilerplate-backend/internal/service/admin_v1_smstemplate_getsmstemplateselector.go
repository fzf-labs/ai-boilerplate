package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSmsTemplateSelector 短信模板-模板选择器
func (a *AdminV1SmsTemplateService) GetSmsTemplateSelector(ctx context.Context, _ *pb.GetSmsTemplateSelectorReq) (*pb.GetSmsTemplateSelectorReply, error) {
	resp := &pb.GetSmsTemplateSelectorReply{
		List: []*pb.SmsTemplateSelector{},
	}
	list, _, err := a.smsTemplateRepo.FindMultiCacheByCondition(ctx, &condition.Req{})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.SmsTemplateSelector{
			Id:           v.ID,
			TemplateName: v.TemplateName,
		})
	}
	return resp, nil
}
