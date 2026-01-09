package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysDept 系统-部门-更新一条数据
func (a *AdminV1SysDeptService) UpdateSysDept(ctx context.Context, req *pb.UpdateSysDeptReq) (*pb.UpdateSysDeptReply, error) {
	resp := &pb.UpdateSysDeptReply{}
	data, err := a.sysDeptRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysDeptRepo.DeepCopy(data)
	data.Pid = req.GetPid()
	data.Name = req.GetName()
	data.AdminID = req.GetAdminId()
	data.Status = int16(req.GetStatus())
	data.Sort = int64(req.GetSort())
	err = a.sysDeptRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
