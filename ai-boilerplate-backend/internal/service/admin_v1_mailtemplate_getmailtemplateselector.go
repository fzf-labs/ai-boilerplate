package service

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetMailTemplateSelector 邮件模版表-选择器
func (a *AdminV1MailTemplateService) GetMailTemplateSelector(ctx context.Context, req *pb.GetMailTemplateSelectorReq) (*pb.GetMailTemplateSelectorReply, error) {
	resp := &pb.GetMailTemplateSelectorReply{
		List: []*pb.MailTemplateSelectorItem{},
	}
	var mailTemplates []*ai_boilerplate_model.MailTemplate
	var err error
	if req.GetAccountId() == "" {
		mailTemplates, _, err = a.mailTemplateRepo.FindMultiCacheByCondition(ctx, &condition.Req{})
		if err != nil {
			return nil, err
		}
	} else {
		mailTemplates, err = a.mailTemplateRepo.FindMultiCacheByAccountID(ctx, req.GetAccountId())
		if err != nil {
			return nil, err
		}
	}
	// 创建时间排序
	sort.Slice(mailTemplates, func(i, j int) bool {
		return mailTemplates[i].CreatedAt.After(mailTemplates[j].CreatedAt)
	})
	for _, mailTemplate := range mailTemplates {
		resp.List = append(resp.List, &pb.MailTemplateSelectorItem{
			Id:   mailTemplate.ID,
			Name: mailTemplate.Name,
		})
	}
	return resp, nil
}
