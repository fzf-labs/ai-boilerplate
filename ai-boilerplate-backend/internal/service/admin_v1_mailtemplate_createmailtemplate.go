package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"gorm.io/datatypes"
)

// CreateMailTemplate 邮件模版表-创建一条数据
func (a *AdminV1MailTemplateService) CreateMailTemplate(ctx context.Context, req *pb.CreateMailTemplateReq) (*pb.CreateMailTemplateReply, error) {
	resp := &pb.CreateMailTemplateReply{}
	data := a.mailTemplateRepo.NewData()
	data.Name = req.GetName()
	data.Code = req.GetCode()
	data.AccountID = req.GetAccountId()
	data.Nickname = req.GetNickname()
	data.Title = req.GetTitle()
	data.Content = req.GetContent()
	data.Params = datatypes.JSON(req.GetParams())
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err := a.mailTemplateRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
