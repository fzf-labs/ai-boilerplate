package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysAdmin 系统-用户-删除一条数据
func (a *AdminV1SysAdminService) DeleteSysAdmin(ctx context.Context, req *pb.DeleteSysAdminReq) (*pb.DeleteSysAdminReply, error) {
	resp := &pb.DeleteSysAdminReply{}
	err := a.sysAdminRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
