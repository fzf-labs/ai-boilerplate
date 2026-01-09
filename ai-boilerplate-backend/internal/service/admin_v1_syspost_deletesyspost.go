package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysPost 系统-工作岗位-删除一条数据
func (a *AdminV1SysPostService) DeleteSysPost(ctx context.Context, req *pb.DeleteSysPostReq) (*pb.DeleteSysPostReply, error) {
	resp := &pb.DeleteSysPostReply{}
	err := a.sysPostRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
