package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// CreateSelfApp 自应用信息表-创建一条数据
func (a *AdminV1SelfAppService) CreateSelfApp(ctx context.Context, req *pb.CreateSelfAppReq) (*pb.CreateSelfAppReply, error) {
	resp := &pb.CreateSelfAppReply{}
	data := a.selfAppRepo.NewData()
	data.PackageName = req.GetPackageName()
	data.Name = req.GetName()
	data.Description = req.GetDescription()
	data.Status = req.GetStatus()
	err := a.selfAppRepo.CreateOneCache(ctx, data)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Id = data.ID
	return resp, nil
}
