package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysMenuInfo 菜单和权限规则表-单条数据查询
func (a *AdminV1SysMenuService) GetSysMenuInfo(ctx context.Context, req *pb.GetSysMenuInfoReq) (*pb.GetSysMenuInfoReply, error) {
	resp := &pb.GetSysMenuInfoReply{}
	data, err := a.sysMenuRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysMenuInfo{
		Id:            data.ID,
		Pid:           data.Pid,
		Name:          data.Name,
		Type:          data.Type,
		Path:          data.Path,
		Permission:    data.Permission,
		Icon:          data.Icon,
		Component:     data.Component,
		ComponentName: data.ComponentName,
		Sort:          int32(data.Sort),
		Status:        int32(data.Status),
		CreatedAt:     timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:     timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
