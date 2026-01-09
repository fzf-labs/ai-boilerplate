package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSysAPI 系统-接口-更新一条数据
func (a *AdminV1SysAPIService) UpdateSysAPI(ctx context.Context, req *pb.UpdateSysAPIReq) (*pb.UpdateSysAPIReply, error) {
	resp := &pb.UpdateSysAPIReply{}
	data, err := a.sysAPIRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.sysAPIRepo.DeepCopy(data)
	// TODO
	err = a.sysAPIRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
