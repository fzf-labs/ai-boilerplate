package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateAiAudioRecord AI 音乐表-更新一条数据
func (a *AdminV1AiAudioRecordService) UpdateAiAudioRecord(ctx context.Context, req *pb.UpdateAiAudioRecordReq) (*pb.UpdateAiAudioRecordReply, error) {
	resp := &pb.UpdateAiAudioRecordReply{}
	data, err := a.aiAudioRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.aiAudioRecordRepo.DeepCopy(data)
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
	err = a.aiAudioRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
