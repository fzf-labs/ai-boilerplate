package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysRoleInfo 系统-角色-单条数据查询
func (a *AdminV1SysRoleService) GetSysRoleInfo(ctx context.Context, req *pb.GetSysRoleInfoReq) (*pb.GetSysRoleInfoReply, error) {
	resp := &pb.GetSysRoleInfoReply{}
	data, err := a.sysRoleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	menuIDs := make([]string, 0)
	if data.MenuIds.String() != "" {
		if unmarshalErr := jsonutil.Unmarshal(data.MenuIds, &menuIDs); unmarshalErr != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(unmarshalErr))
		}
	}
	resp.Info = &pb.SysRoleInfo{
		Id:        data.ID,
		TenantId:  data.TenantID,
		Name:      data.Name,
		Remark:    data.Remark,
		DataScope: data.DataScope,
		MenuIds:   menuIDs,
		Sort:      int32(data.Sort),
		Status:    int32(data.Status),
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
