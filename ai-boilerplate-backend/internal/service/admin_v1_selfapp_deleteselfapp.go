package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSelfApp 自应用信息表-删除一条数据
func (a *AdminV1SelfAppService) DeleteSelfApp(ctx context.Context, req *pb.DeleteSelfAppReq) (*pb.DeleteSelfAppReply, error) {
	resp := &pb.DeleteSelfAppReply{}
	err := a.selfAppRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
