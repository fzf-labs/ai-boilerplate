package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSysTenantInfo 系统-租户-单条数据查询
func (a *AdminV1SysTenantService) GetSysTenantInfo(ctx context.Context, req *pb.GetSysTenantInfoReq) (*pb.GetSysTenantInfoReply, error) {
	resp := &pb.GetSysTenantInfoReply{}
	data, err := a.sysTenantRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	adminName, err := a.sysAdminRepo.AdminIDToNickname(ctx, []string{data.AdminID})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	menuIDs := make([]string, 0)
	if data.MenuIds.String() != "" {
		err = jsonutil.Unmarshal(data.MenuIds, &menuIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
		}
	}
	resp.Info = &pb.SysTenantInfo{
		Id:         data.ID,
		Name:       data.Name,
		Remark:     data.Remark,
		AdminId:    data.AdminID,
		ExpireTime: timeutil.RFC3339(data.ExpireTime.Time),
		MenuIds:    menuIDs,
		Status:     int32(data.Status),
		CreatedAt:  timeutil.RFC3339(data.CreatedAt.In(time.Local)),
		UpdatedAt:  timeutil.RFC3339(data.UpdatedAt.In(time.Local)),
		AdminName:  adminName[data.AdminID],
	}
	return resp, nil
}
