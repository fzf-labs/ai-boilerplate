package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSmsTemplateList 短信模板-列表数据查询
func (a *AdminV1SmsTemplateService) GetSmsTemplateList(ctx context.Context, req *pb.GetSmsTemplateListReq) (*pb.GetSmsTemplateListReply, error) {
	resp := &pb.GetSmsTemplateListReply{
		Total: 0,
		List:  []*pb.SmsTemplateInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.smsTemplateRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		smsChannelIDs := make([]string, 0)
		for _, v := range list {
			smsChannelIDs = append(smsChannelIDs, v.SmsChannelID)
		}
		smsChannelIDToName, err := a.smsChannelRepo.IDToName(ctx, smsChannelIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.SmsTemplateInfo{
				Id:              v.ID,
				SmsChannelId:    v.SmsChannelID,
				TemplateType:    int32(v.TemplateType),
				TemplateCode:    v.TemplateCode,
				TemplateName:    v.TemplateName,
				TemplateContent: v.TemplateContent,
				TemplateParams:  v.TemplateParams.String(),
				Remark:          v.Remark,
				ApiTemplateId:   v.APITemplateID,
				Status:          int32(v.Status),
				CreatedAt:       v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:       v.UpdatedAt.Format(time.RFC3339),
				SmsChannelName:  smsChannelIDToName[v.SmsChannelID],
			})
		}
	}
	return resp, nil
}
