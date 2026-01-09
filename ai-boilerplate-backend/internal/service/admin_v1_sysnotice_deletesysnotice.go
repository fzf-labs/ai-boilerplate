package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysNotice 系统-公告-删除一条数据
func (a *AdminV1SysNoticeService) DeleteSysNotice(ctx context.Context, req *pb.DeleteSysNoticeReq) (*pb.DeleteSysNoticeReply, error) {
	resp := &pb.DeleteSysNoticeReply{}
	err := a.sysNoticeRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
