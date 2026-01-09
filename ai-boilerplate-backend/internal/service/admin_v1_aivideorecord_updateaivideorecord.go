package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
	"gorm.io/datatypes"
)

// UpdateAiVideoRecord AI 视频表-更新一条数据
func (a *AdminV1AiVideoRecordService) UpdateAiVideoRecord(ctx context.Context, req *pb.UpdateAiVideoRecordReq) (*pb.UpdateAiVideoRecordReply, error) {
	resp := &pb.UpdateAiVideoRecordReply{}
	data, err := a.aiVideoRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiVideoRecordRepo.DeepCopy(data)
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
	err = a.aiVideoRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
