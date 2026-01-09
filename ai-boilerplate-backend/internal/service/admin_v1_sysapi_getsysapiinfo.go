package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSysAPIInfo 系统-接口-单条数据查询
func (a *AdminV1SysAPIService) GetSysAPIInfo(ctx context.Context, req *pb.GetSysAPIInfoReq) (*pb.GetSysAPIInfoReply, error) {
	resp := &pb.GetSysAPIInfoReply{}
	data, err := a.sysAPIRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysAPIInfo{
		Id: data.ID,
		// TODO
	}
	return resp, nil
}
