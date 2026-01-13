package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// DeleteAiIndexVideoRecord AI 视频表-删除一条数据
func (a *AdminV1AiIndexVideoService) DeleteAiIndexVideoRecord(ctx context.Context, req *pb.DeleteAiIndexVideoRecordReq) (*pb.DeleteAiIndexVideoRecordReply, error) {
	resp := &pb.DeleteAiIndexVideoRecordReply{}
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
	err = a.aiVideoRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
