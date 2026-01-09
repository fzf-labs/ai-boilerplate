package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateSysNotifyMessageRead 系统-通知消息-我的-指定消息已读
func (a *AdminV1SysNotifyMessageService) UpdateSysNotifyMessageRead(ctx context.Context, req *pb.UpdateSysNotifyMessageReadReq) (*pb.UpdateSysNotifyMessageReadReply, error) {
	resp := &pb.UpdateSysNotifyMessageReadReply{}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	data, err := a.sysNotifyMessageRepo.FindMultiCacheByIDS(ctx, req.GetIds())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(data) == 0 {
		return resp, nil
	}
	ids := make([]string, 0)
	for _, v := range data {
		if v.Receiver != adminID {
			return nil, pb.ErrorReasonAccountNoDataPermission()
		}
		ids = append(ids, v.ID)
	}
	err = a.sysNotifyMessageRepo.UpdateBatchByIDS(ctx, ids, map[string]interface{}{
		"read_time": timeutil.RFC3339(time.Now()),
	})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	err = a.sysNotifyMessageRepo.DeleteIndexCache(ctx, data...)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	return resp, nil
}
