package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysDeptInfo 系统-部门-单条数据查询
func (a *AdminV1SysDeptService) GetSysDeptInfo(ctx context.Context, req *pb.GetSysDeptInfoReq) (*pb.GetSysDeptInfoReply, error) {
	resp := &pb.GetSysDeptInfoReply{}
	data, err := a.sysDeptRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	adminIDToNickname, err := a.sysAdminRepo.AdminIDToNickname(ctx, []string{data.AdminID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SysDeptInfo{
		Id:        data.ID,
		Pid:       data.Pid,
		Name:      data.Name,
		AdminId:   data.AdminID,
		Status:    int32(data.Status),
		Sort:      int32(data.Sort),
		CreatedAt: timeutil.RFC3339(data.CreatedAt),
		UpdatedAt: timeutil.RFC3339(data.UpdatedAt),
		AdminName: adminIDToNickname[data.AdminID],
	}
	return resp, nil
}
