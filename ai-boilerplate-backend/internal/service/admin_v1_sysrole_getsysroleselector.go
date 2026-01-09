package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysRoleSelector 系统-角色-选择器
func (a *AdminV1SysRoleService) GetSysRoleSelector(ctx context.Context, req *pb.GetSysRoleSelectorReq) (*pb.GetSysRoleSelectorReply, error) {
	_ = req

	resp := &pb.GetSysRoleSelectorReply{
		List: []*pb.GetSysRoleSelectorItem{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Query: []*condition.QueryParam{
			{
				Field: "tenant_id",
				Value: tenantID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, _, err := a.sysRoleRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.GetSysRoleSelectorItem{
				Id:   v.ID,
				Name: v.Name,
			})
		}
	}
	return resp, nil
}
