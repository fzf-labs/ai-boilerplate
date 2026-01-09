package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
	"gorm.io/datatypes"
)

// UpdateAiImageRecord AI 绘画表-更新一条数据
func (a *AdminV1AiImageRecordService) UpdateAiImageRecord(ctx context.Context, req *pb.UpdateAiImageRecordReq) (*pb.UpdateAiImageRecordReply, error) {
	resp := &pb.UpdateAiImageRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data, err := a.aiImageRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiImageRecordRepo.DeepCopy(data)
	data.TenantID = tenantID
	data.AdminID = req.GetAdminId()
	data.Prompt = req.GetPrompt()
	data.Platform = req.GetPlatform()
	data.ModelID = req.GetModelId()
	data.Model = req.GetModel()
	data.Width = req.GetWidth()
	data.Height = req.GetHeight()
	data.Status = req.GetStatus()
	data.FinishTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetFinishTime()).StdTime())
	data.ErrorMessage = req.GetErrorMessage()
	data.PublicStatus = req.GetPublicStatus()
	data.PicURL = req.GetPicURL()
	data.Options = datatypes.JSON(req.GetOptions())
	data.TaskID = req.GetTaskId()
	data.Buttons = req.GetButtons()
	err = a.aiImageRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
