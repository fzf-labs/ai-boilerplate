package service

import (
	"context"

	"github.com/dromara/carbon/v2"
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/fzf-labs/goutil/timeutil"
)

// UpdateSysTenant 系统-租户-更新一条数据
func (a *AdminV1SysTenantService) UpdateSysTenant(ctx context.Context, req *pb.UpdateSysTenantReq) (*pb.UpdateSysTenantReply, error) {
	resp := &pb.UpdateSysTenantReply{}
	data, err := a.sysTenantRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if data == nil || data.ID == "" {
		return nil, pb.ErrorReasonDataRecordNotFound()
	}
	menuIDs, err := jsonutil.Marshal(req.GetMenuIds())
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	oldData := a.sysTenantRepo.DeepCopy(data)
	data.MenuIds = menuIDs
	data.Name = req.GetName()
	data.Remark = req.GetRemark()
	data.ExpireTime = timeutil.TimeToSQLNullTime(carbon.Parse(req.GetExpireTime()).StdTime())
	data.Status = int16(req.GetStatus())
	err = a.sysTenantRepo.UpdateOneCache(ctx, data, oldData)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
