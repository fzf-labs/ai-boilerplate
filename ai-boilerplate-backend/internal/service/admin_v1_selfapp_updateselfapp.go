package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// UpdateSelfApp 自应用信息表-更新一条数据
func (a *AdminV1SelfAppService) UpdateSelfApp(ctx context.Context, req *pb.UpdateSelfAppReq) (*pb.UpdateSelfAppReply, error) {
	resp := &pb.UpdateSelfAppReply{}
	data, err := a.selfAppRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	oldData := a.selfAppRepo.DeepCopy(data)
	data.PackageName = req.GetPackageName()
	data.Name = req.GetName()
	data.Description = req.GetDescription()
	data.Status = req.GetStatus()
	err = a.selfAppRepo.UpdateOneCacheWithZero(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
