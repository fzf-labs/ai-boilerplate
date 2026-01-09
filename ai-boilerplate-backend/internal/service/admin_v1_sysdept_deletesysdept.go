package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysDept 系统-部门-删除一条数据
func (a *AdminV1SysDeptService) DeleteSysDept(ctx context.Context, req *pb.DeleteSysDeptReq) (*pb.DeleteSysDeptReply, error) {
	resp := &pb.DeleteSysDeptReply{}
	err := a.sysDeptRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
