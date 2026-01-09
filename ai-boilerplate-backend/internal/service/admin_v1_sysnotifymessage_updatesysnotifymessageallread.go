package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/samber/lo"
)

// UpdateSysNotifyMessageAllRead 系统-通知消息-我的-全部已读
func (a *AdminV1SysNotifyMessageService) UpdateSysNotifyMessageAllRead(ctx context.Context, req *pb.UpdateSysNotifyMessageAllReadReq) (*pb.UpdateSysNotifyMessageAllReadReply, error) {
	_ = req

	resp := &pb.UpdateSysNotifyMessageAllReadReply{}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	list, err := a.sysNotifyMessageRepo.FindMultiCacheByReceiverReadTime(ctx, adminID, "")
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) == 0 {
		return resp, nil
	}
	ids := lo.Map(list, func(item *ai_boilerplate_model.SysNotifyMessage, _ int) string {
		return item.ID
	})
	err = a.sysNotifyMessageRepo.UpdateBatchByIDS(ctx, ids, map[string]interface{}{
		"read_time": timeutil.RFC3339(time.Now()),
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	err = a.sysNotifyMessageRepo.DeleteIndexCache(ctx, list...)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
