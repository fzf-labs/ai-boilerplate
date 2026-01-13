package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// DeleteAiIndexImageRecord AI 绘画表-删除一条数据
func (a *AdminV1AiIndexImageService) DeleteAiIndexImageRecord(ctx context.Context, req *pb.DeleteAiIndexImageRecordReq) (*pb.DeleteAiIndexImageRecordReply, error) {
	resp := &pb.DeleteAiIndexImageRecordReply{}
	data, err := a.aiImageRecordRepo.FindOneCacheByID(ctx, req.GetId())
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
	err = a.aiImageRecordRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
