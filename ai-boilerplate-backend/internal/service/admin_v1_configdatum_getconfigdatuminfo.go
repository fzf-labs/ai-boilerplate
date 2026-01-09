package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetConfigDatumInfo 配置管理-单条数据查询
func (a *AdminV1ConfigDatumService) GetConfigDatumInfo(ctx context.Context, req *pb.GetConfigDatumInfoReq) (*pb.GetConfigDatumInfoReply, error) {
	resp := &pb.GetConfigDatumInfoReply{}
	data, err := a.configDatumRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.ConfigDatumInfo{
		Id:        data.ID,
		Name:      data.Name,
		Key:       data.Key,
		Value:     string(data.Value),
		Remark:    data.Remark,
		Status:    data.Status,
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
