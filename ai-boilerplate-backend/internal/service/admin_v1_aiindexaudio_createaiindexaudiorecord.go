package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateAiIndexAudioRecord AI 音乐表-创建一条数据
func (a *AdminV1AiIndexAudioService) CreateAiIndexAudioRecord(ctx context.Context, req *pb.CreateAiIndexAudioRecordReq) (*pb.CreateAiIndexAudioRecordReply, error) {
	resp := &pb.CreateAiIndexAudioRecordReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	if tenantID == "" {
		tenantID = req.GetTenantId()
	}
	if adminID == "" {
		adminID = req.GetAdminId()
	}
	data := a.aiAudioRecordRepo.NewData()
	data.TenantID = tenantID
	data.AdminID = adminID
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
