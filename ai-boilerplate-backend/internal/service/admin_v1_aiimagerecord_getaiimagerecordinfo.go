package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiImageRecordInfo AI 绘画表-单条数据查询
func (a *AdminV1AiImageRecordService) GetAiImageRecordInfo(ctx context.Context, req *pb.GetAiImageRecordInfoReq) (*pb.GetAiImageRecordInfoReply, error) {
	resp := &pb.GetAiImageRecordInfoReply{}
	data, err := a.aiImageRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiImageRecordInfo{
		Id:           data.ID,
		AdminId:      data.AdminID,
		Prompt:       data.Prompt,
		Platform:     data.Platform,
		ModelId:      data.ModelID,
		Model:        data.Model,
		Width:        data.Width,
		Height:       data.Height,
		Status:       data.Status,
		FinishTime:   data.FinishTime.Time.Format(time.RFC3339),
		ErrorMessage: data.ErrorMessage,
		PublicStatus: data.PublicStatus,
		PicURL:       data.PicURL,
		Options:      data.Options.String(),
		TaskId:       data.TaskID,
		Buttons:      data.Buttons,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
