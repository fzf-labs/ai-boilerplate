package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetAiWriteRecordInfo AI 写作表-单条数据查询
func (a *AdminV1AiWriteRecordService) GetAiWriteRecordInfo(ctx context.Context, req *pb.GetAiWriteRecordInfoReq) (*pb.GetAiWriteRecordInfoReply, error) {
	resp := &pb.GetAiWriteRecordInfoReply{}
	data, err := a.aiWriteRecordRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.AiWriteRecordInfo{
		Id:               data.ID,
		AdminId:          data.AdminID,
		Type:             data.Type,
		Platform:         data.Platform,
		ModelId:          data.ModelID,
		Model:            data.Model,
		Prompt:           data.Prompt,
		GeneratedContent: data.GeneratedContent,
		OriginalContent:  data.OriginalContent,
		Length:           data.Length,
		Format:           data.Format,
		Tone:             data.Tone,
		Language:         data.Language,
		ErrorMessage:     data.ErrorMessage,
		CreatedAt:        data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
