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

// CreateAiVideoRecord AI 视频表-创建一条数据
func (a *AdminV1AiVideoRecordService) CreateAiVideoRecord(ctx context.Context, req *pb.CreateAiVideoRecordReq) (*pb.CreateAiVideoRecordReply, error) {
	resp := &pb.CreateAiVideoRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data := a.aiVideoRecordRepo.NewData()
	data.TenantID = tenantID
	data.AdminID = req.GetAdminId()
	data.Prompt = req.GetPrompt()
	data.Platform = req.GetPlatform()
	data.ModelID = req.GetModelId()
	data.Model = req.GetModel()
	data.Status = req.GetStatus()
	data.FinishTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetFinishTime()).StdTime())
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
