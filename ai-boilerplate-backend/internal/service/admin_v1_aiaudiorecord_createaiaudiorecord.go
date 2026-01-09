package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateAiAudioRecord AI 音乐表-创建一条数据
func (a *AdminV1AiAudioRecordService) CreateAiAudioRecord(ctx context.Context, req *pb.CreateAiAudioRecordReq) (*pb.CreateAiAudioRecordReply, error) {
	resp := &pb.CreateAiAudioRecordReply{}
	data := a.aiAudioRecordRepo.NewData()
	data.TenantID = req.GetTenantId()
	data.AdminID = req.GetAdminId()
	data.Title = req.GetTitle()
	data.Lyric = req.GetLyric()
	data.ImageURL = req.GetImageURL()
	data.AudioURL = req.GetAudioURL()
	data.Status = req.GetStatus()
	data.Description = req.GetDescription()
	data.Prompt = req.GetPrompt()
	data.Platform = req.GetPlatform()
	data.ModelID = req.GetModelId()
	data.Model = req.GetModel()
	data.GenerateMode = req.GetGenerateMode()
	data.Tags = req.GetTags()
	data.Duration = req.GetDuration()
	data.PublicStatus = req.GetPublicStatus()
	data.TaskID = req.GetTaskId()
	data.ErrorMessage = req.GetErrorMessage()
	err := a.aiAudioRecordRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
