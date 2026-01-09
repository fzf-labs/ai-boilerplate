package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteMailLog 邮件日志表-删除一条数据
func (a *AdminV1MailLogService) DeleteMailLog(ctx context.Context, req *pb.DeleteMailLogReq) (*pb.DeleteMailLogReply, error) {
	resp := &pb.DeleteMailLogReply{}
	err := a.mailLogRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
