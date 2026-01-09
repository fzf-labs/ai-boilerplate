package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
)

// GetFileConfigSelect 文件配置表-获取所有选择器
func (a *AdminV1FileConfigService) GetFileConfigSelect(ctx context.Context, _ *pb.GetFileConfigSelectReq) (*pb.GetFileConfigSelectReply, error) {
	resp := &pb.GetFileConfigSelectReply{}
	list, _, err := a.fileConfigRepo.FindMultiCacheByCondition(ctx, &condition.Req{})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range list {
		resp.List = append(resp.List, &pb.FileConfigSelect{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return resp, nil
}
