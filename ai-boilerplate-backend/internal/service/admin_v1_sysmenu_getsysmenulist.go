package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetSysMenuList 菜单和权限规则表-列表数据查询
func (a *AdminV1SysMenuService) GetSysMenuList(ctx context.Context, _ *pb.GetSysMenuListReq) (*pb.GetSysMenuListReply, error) {
	resp := &pb.GetSysMenuListReply{
		Total: 0,
		List:  []*pb.SysMenuInfo{},
	}
	param := &condition.Req{
		Query: []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "sort",
				Order: condition.ASC,
			},
		},
	}
	list, _, err := a.sysMenuRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysMenuInfo{
				Id:            v.ID,
				Pid:           v.Pid,
				Name:          v.Name,
				Type:          v.Type,
				Path:          v.Path,
				Permission:    v.Permission,
				Icon:          v.Icon,
				Component:     v.Component,
				ComponentName: v.ComponentName,
				Sort:          int32(v.Sort),
				Status:        int32(v.Status),
				CreatedAt:     v.CreatedAt.Format(time.RFC3339),
				UpdatedAt:     v.UpdatedAt.Format(time.RFC3339),
			})
		}
	}
	resp.Total = int32(len(list))
	return resp, nil
}
