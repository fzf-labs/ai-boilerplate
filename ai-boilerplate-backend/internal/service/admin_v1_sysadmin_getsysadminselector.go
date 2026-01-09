package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysAdminSelector 系统-用户-选择器
func (a *AdminV1SysAdminService) GetSysAdminSelector(ctx context.Context, req *pb.GetSysAdminSelectorReq) (*pb.GetSysAdminSelectorReply, error) {
	_ = req

	resp := &pb.GetSysAdminSelectorReply{
		List: []*pb.GetSysAdminSelectorItem{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	list, _, err := a.sysAdminRepo.FindMultiCacheByCondition(ctx, &condition.Req{
		Query: []*condition.QueryParam{
			{
				Field: "tenant_id",
				Value: tenantID,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "updated_at",
				Order: condition.ASC,
			},
		},
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, item := range list {
		resp.List = append(resp.List, &pb.GetSysAdminSelectorItem{
			Id:       item.ID,
			Nickname: item.Nickname,
		})
	}
	return resp, nil
}
