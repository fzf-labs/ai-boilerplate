package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiVideoRecordInfo AI 视频表-单条数据查询
func (a *AdminV1AiVideoRecordService) GetAiVideoRecordInfo(ctx context.Context, req *pb.GetAiVideoRecordInfoReq) (*pb.GetAiVideoRecordInfoReply, error) {
	resp := &pb.GetAiVideoRecordInfoReply{}
	data, err := a.aiVideoRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiVideoRecordInfo{
		Id:           data.ID,
		AdminId:      data.AdminID,
		Prompt:       data.Prompt,
		Platform:     data.Platform,
		ModelId:      data.ModelID,
		Model:        data.Model,
		Status:       data.Status,
		FinishTime:   data.FinishTime.Time.Format(time.RFC3339),
		ErrorMessage: data.ErrorMessage,
		PublicStatus: data.PublicStatus,
		VideoURL:     data.VideoURL,
		Options:      data.Options.String(),
		TaskId:       data.TaskID,
		CreatedAt:    data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
