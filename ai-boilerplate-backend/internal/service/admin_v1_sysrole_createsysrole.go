package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// CreateSysRole 系统-角色-创建一条数据
func (a *AdminV1SysRoleService) CreateSysRole(ctx context.Context, req *pb.CreateSysRoleReq) (*pb.CreateSysRoleReply, error) {
	resp := &pb.CreateSysRoleReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	menuIDs, err := jsonutil.Marshal(req.GetMenuIds())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	data := a.sysRoleRepo.NewData()
	data.TenantID = tenantID
	data.Name = req.GetName()
	data.Remark = req.GetRemark()
	data.DataScope = req.GetDataScope()
	data.MenuIds = menuIDs
	data.Sort = int64(req.GetSort())
	data.Status = int16(req.GetStatus())
	err = a.sysRoleRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
