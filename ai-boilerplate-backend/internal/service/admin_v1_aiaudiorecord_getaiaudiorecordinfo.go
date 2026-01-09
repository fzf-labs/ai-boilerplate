package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiAudioRecordInfo AI 音乐表-单条数据查询
func (a *AdminV1AiAudioRecordService) GetAiAudioRecordInfo(ctx context.Context, req *pb.GetAiAudioRecordInfoReq) (*pb.GetAiAudioRecordInfoReply, error) {
	resp := &pb.GetAiAudioRecordInfoReply{}
	data, err := a.aiAudioRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiAudioRecordInfo{
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
