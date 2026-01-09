package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSysDept 系统-部门-创建一条数据
func (a *AdminV1SysDeptService) CreateSysDept(ctx context.Context, req *pb.CreateSysDeptReq) (*pb.CreateSysDeptReply, error) {
	resp := &pb.CreateSysDeptReply{}
	data := a.sysDeptRepo.NewData()
	data.Pid = req.GetPid()
	data.Name = req.GetName()
	data.AdminID = req.GetAdminId()
	data.Status = int16(req.GetStatus())
	data.Sort = int64(req.GetSort())
	err := a.sysDeptRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
