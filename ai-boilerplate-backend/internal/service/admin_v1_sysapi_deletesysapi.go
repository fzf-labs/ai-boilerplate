package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// DeleteSysAPI 系统-接口-删除一条数据
func (a *AdminV1SysAPIService) DeleteSysAPI(ctx context.Context, req *pb.DeleteSysAPIReq) (*pb.DeleteSysAPIReply, error) {
	resp := &pb.DeleteSysAPIReply{}
	err := a.sysAPIRepo.DeleteMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
