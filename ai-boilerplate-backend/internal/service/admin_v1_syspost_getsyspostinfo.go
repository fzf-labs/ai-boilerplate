package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysPostInfo 系统-工作岗位-单条数据查询
func (a *AdminV1SysPostService) GetSysPostInfo(ctx context.Context, req *pb.GetSysPostInfoReq) (*pb.GetSysPostInfoReply, error) {
	resp := &pb.GetSysPostInfoReply{}
	data, err := a.sysPostRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysPostInfo{
		Id:        data.ID,
		TenantId:  data.TenantID,
		Name:      data.Name,
		Code:      data.Code,
		Remark:    data.Remark,
		Sort:      int32(data.Sort),
		Status:    int32(data.Status),
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
