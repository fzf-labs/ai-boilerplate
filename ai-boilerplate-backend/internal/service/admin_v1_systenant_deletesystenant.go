package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysTenant 系统-租户-删除一条数据
func (a *AdminV1SysTenantService) DeleteSysTenant(ctx context.Context, req *pb.DeleteSysTenantReq) (*pb.DeleteSysTenantReply, error) {
	resp := &pb.DeleteSysTenantReply{}
	err := a.sysTenantRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
