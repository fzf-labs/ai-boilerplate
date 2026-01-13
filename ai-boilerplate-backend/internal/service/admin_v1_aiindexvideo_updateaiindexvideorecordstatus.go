package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateAiIndexVideoRecordStatus AI 视频表-更新状态
func (a *AdminV1AiIndexVideoService) UpdateAiIndexVideoRecordStatus(ctx context.Context, req *pb.UpdateAiIndexVideoRecordStatusReq) (*pb.UpdateAiIndexVideoRecordStatusReply, error) {
	resp := &pb.UpdateAiIndexVideoRecordStatusReply{}
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
	oldData := a.aiVideoRecordRepo.DeepCopy(data)
	data.Status = req.GetStatus()
	err = a.aiVideoRecordRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
