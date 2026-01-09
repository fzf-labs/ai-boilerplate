package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSelfAppRelease 自应用版本发布表-删除一条数据
func (a *AdminV1SelfAppReleaseService) DeleteSelfAppRelease(ctx context.Context, req *pb.DeleteSelfAppReleaseReq) (*pb.DeleteSelfAppReleaseReply, error) {
	resp := &pb.DeleteSelfAppReleaseReply{}
	err := a.selfAppReleaseRepo.DeleteOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
