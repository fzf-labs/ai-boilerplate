package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"gorm.io/datatypes"
)

// UpdateMailTemplate 邮件模版表-更新一条数据
func (a *AdminV1MailTemplateService) UpdateMailTemplate(ctx context.Context, req *pb.UpdateMailTemplateReq) (*pb.UpdateMailTemplateReply, error) {
	resp := &pb.UpdateMailTemplateReply{}
	data, err := a.mailTemplateRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.mailTemplateRepo.DeepCopy(data)
	data.Name = req.GetName()
	data.Code = req.GetCode()
	data.AccountID = req.GetAccountId()
	data.Nickname = req.GetNickname()
	data.Title = req.GetTitle()
	data.Content = req.GetContent()
	data.Params = datatypes.JSON(req.GetParams())
	data.Remark = req.GetRemark()
	data.Status = req.GetStatus()
	err = a.mailTemplateRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
