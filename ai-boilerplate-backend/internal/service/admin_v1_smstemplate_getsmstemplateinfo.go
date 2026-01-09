package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSmsTemplateInfo 短信模板-单条数据查询
func (a *AdminV1SmsTemplateService) GetSmsTemplateInfo(ctx context.Context, req *pb.GetSmsTemplateInfoReq) (*pb.GetSmsTemplateInfoReply, error) {
	resp := &pb.GetSmsTemplateInfoReply{}
	data, err := a.smsTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SmsTemplateInfo{
		Id:              data.ID,
		SmsChannelId:    data.SmsChannelID,
		TemplateType:    int32(data.TemplateType),
		TemplateCode:    data.TemplateCode,
		TemplateName:    data.TemplateName,
		TemplateContent: data.TemplateContent,
		TemplateParams:  data.TemplateParams.String(),
		Remark:          data.Remark,
		ApiTemplateId:   data.APITemplateID,
		Status:          int32(data.Status),
		CreatedAt:       timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:       timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
