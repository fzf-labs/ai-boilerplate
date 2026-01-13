package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetAiIndexVideoRecordInfo AI 视频表-单条数据查询
func (a *AdminV1AiIndexVideoService) GetAiIndexVideoRecordInfo(ctx context.Context, req *pb.GetAiIndexVideoRecordInfoReq) (*pb.GetAiIndexVideoRecordInfoReply, error) {
	resp := &pb.GetAiIndexVideoRecordInfoReply{}
	data, err := a.aiVideoRecordRepo.FindOneCacheByID(ctx, req.GetId())
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
	resp.Info = &pb.AiIndexVideoRecordInfo{
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
