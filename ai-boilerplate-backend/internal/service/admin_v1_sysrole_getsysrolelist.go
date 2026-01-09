package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysRoleList 系统-角色-列表数据查询
func (a *AdminV1SysRoleService) GetSysRoleList(ctx context.Context, req *pb.GetSysRoleListReq) (*pb.GetSysRoleListReply, error) {
	resp := &pb.GetSysRoleListReply{
		Total: 0,
		List:  []*pb.SysRoleInfo{},
	}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
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
	if req.GetName() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "name",
			Value: "%" + req.GetName() + "%",
			Exp:   condition.LIKE,
			Logic: condition.AND,
		})
	}
	if req.GetStatus() != 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "status",
			Value: req.GetStatus(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetCreatedAt()) > 0 {
		param.Query = append(param.Query,
			&condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[0],
				Exp:   condition.GTE,
				Logic: condition.AND,
			}, &condition.QueryParam{
				Field: "created_at",
				Value: req.GetCreatedAt()[1],
				Exp:   condition.LTE,
				Logic: condition.AND,
			})
	}
	list, p, err := a.sysRoleRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			menuIDs := make([]string, 0)
			if v.MenuIds.String() != "" {
				if unmarshalErr := jsonutil.Unmarshal(v.MenuIds, &menuIDs); unmarshalErr != nil {
					return nil, pb.ErrorReasonDataFormattingError(pb.WithError(unmarshalErr))
				}
			}
			resp.List = append(resp.List, &pb.SysRoleInfo{
				Id:        v.ID,
				TenantId:  v.TenantID,
				Name:      v.Name,
				Remark:    v.Remark,
				DataScope: v.DataScope,
				MenuIds:   menuIDs,
				Sort:      int32(v.Sort),
				Status:    int32(v.Status),
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
