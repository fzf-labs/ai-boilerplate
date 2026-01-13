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

// CreateAiIndexVideoRecord AI 视频表-创建一条数据
func (a *AdminV1AiIndexVideoService) CreateAiIndexVideoRecord(ctx context.Context, req *pb.CreateAiIndexVideoRecordReq) (*pb.CreateAiIndexVideoRecordReply, error) {
	resp := &pb.CreateAiIndexVideoRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	if adminID == "" {
		adminID = req.GetAdminId()
	}
	data := a.aiVideoRecordRepo.NewData()
	data.TenantID = tenantID
	data.AdminID = adminID
	data.Prompt = req.GetPrompt()
	data.Platform = req.GetPlatform()
	data.ModelID = req.GetModelId()
	data.Model = req.GetModel()
	data.Status = req.GetStatus()
	if req.GetFinishTime() != "" {
		data.FinishTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetFinishTime()).StdTime())
	}
	data.ErrorMessage = req.GetErrorMessage()
	data.PublicStatus = req.GetPublicStatus()
	data.VideoURL = req.GetVideoURL()
	data.Options = datatypes.JSON(req.GetOptions())
	data.TaskID = req.GetTaskId()
	err := a.aiVideoRecordRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
