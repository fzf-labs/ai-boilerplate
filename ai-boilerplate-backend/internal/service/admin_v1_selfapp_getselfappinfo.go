package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSelfAppInfo 自应用信息表-单条数据查询
func (a *AdminV1SelfAppService) GetSelfAppInfo(ctx context.Context, req *pb.GetSelfAppInfoReq) (*pb.GetSelfAppInfoReply, error) {
	resp := &pb.GetSelfAppInfoReply{}
	data, err := a.selfAppRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SelfAppInfo{
		Id:          data.ID,
		PackageName: data.PackageName,
		Name:        data.Name,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   data.UpdatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
