package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
)

// UpdateSysRole 系统-角色-更新一条数据
func (a *AdminV1SysRoleService) UpdateSysRole(ctx context.Context, req *pb.UpdateSysRoleReq) (*pb.UpdateSysRoleReply, error) {
	resp := &pb.UpdateSysRoleReply{}
	data, err := a.sysRoleRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysRoleRepo.DeepCopy(data)
	menuIDs, err := jsonutil.Marshal(req.GetMenuIds())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	data.Name = req.GetName()
	data.Remark = req.GetRemark()
	data.DataScope = req.GetDataScope()
	data.MenuIds = menuIDs
	data.Sort = int64(req.GetSort())
	data.Status = int16(req.GetStatus())
	err = a.sysRoleRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
