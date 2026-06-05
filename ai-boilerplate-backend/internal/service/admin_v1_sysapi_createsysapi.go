package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSysAPI 系统-接口-创建一条数据
func (a *AdminV1SysAPIService) CreateSysAPI(ctx context.Context, req *pb.CreateSysAPIReq) (*pb.CreateSysAPIReply, error) {
	resp := &pb.CreateSysAPIReply{}
	data := a.sysAPIRepo.NewData()
	data.PermissionID = req.GetPermissionId()
	data.Method = req.GetMethod()
	data.Path = req.GetPath()
	data.Desc = req.GetDesc()
	err := a.sysAPIRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
