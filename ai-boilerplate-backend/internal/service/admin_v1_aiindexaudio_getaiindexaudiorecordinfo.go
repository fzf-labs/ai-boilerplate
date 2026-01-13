package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexAudioRecordInfo AI 音乐表-单条数据查询
func (a *AdminV1AiIndexAudioService) GetAiIndexAudioRecordInfo(ctx context.Context, req *pb.GetAiIndexAudioRecordInfoReq) (*pb.GetAiIndexAudioRecordInfoReply, error) {
	resp := &pb.GetAiIndexAudioRecordInfoReply{}
	data, err := a.aiAudioRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	if tenantID != "" && data.TenantID != tenantID {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if adminID != "" && data.AdminID != adminID {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	resp.Info = &pb.AiIndexAudioRecordInfo{
		Id:           data.ID,
		TenantId:     data.TenantID,
		AdminId:      data.AdminID,
		Title:        data.Title,
		Lyric:        data.Lyric,
		ImageURL:     data.ImageURL,
		AudioURL:     data.AudioURL,
		Status:       data.Status,
		Description:  data.Description,
		Prompt:       data.Prompt,
		Platform:     data.Platform,
		ModelId:      data.ModelID,
		Model:        data.Model,
		GenerateMode: data.GenerateMode,
		Tags:         data.Tags,
		Duration:     data.Duration,
		PublicStatus: data.PublicStatus,
		TaskId:       data.TaskID,
		ErrorMessage: data.ErrorMessage,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
