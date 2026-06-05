package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysAPIInfo 系统-接口-单条数据查询
func (a *AdminV1SysAPIService) GetSysAPIInfo(ctx context.Context, req *pb.GetSysAPIInfoReq) (*pb.GetSysAPIInfoReply, error) {
	resp := &pb.GetSysAPIInfoReply{}
	data, err := a.sysAPIRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	resp.Info = &pb.SysAPIInfo{
		Id:           data.ID,
		PermissionId: data.PermissionID,
		Method:       data.Method,
		Path:         data.Path,
		Desc:         data.Desc,
		CreatedAt:    timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:    timeutil.RFC3339(data.UpdatedAt),
	}
	return resp, nil
}
