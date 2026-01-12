package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateBannerStatus 通用-轮播图-更新状态
func (a *AdminV1BannerService) UpdateBannerStatus(ctx context.Context, req *pb.UpdateBannerStatusReq) (*pb.UpdateBannerStatusReply, error) {
	resp := &pb.UpdateBannerStatusReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	data, err := a.bannerRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	if data.TenantID != tenantID {
		return nil, pb.ErrorReasonAccountNoDataPermission()
	}
	oldData := a.bannerRepo.DeepCopy(data)
	data.Status = int16(req.GetStatus())
	err = a.bannerRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
