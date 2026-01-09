package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetSysNotifyMessageMyUnreadCount 系统-通知消息-我的-未读数量
func (a *AdminV1SysNotifyMessageService) GetSysNotifyMessageMyUnreadCount(ctx context.Context, req *pb.GetSysNotifyMessageMyUnreadCountReq) (*pb.GetSysNotifyMessageMyUnreadCountReply, error) {
	_ = req

	resp := &pb.GetSysNotifyMessageMyUnreadCountReply{}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	list, err := a.sysNotifyMessageRepo.FindMultiCacheByReceiverReadTime(ctx, adminID, "")
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Count = int32(len(list))
	return resp, nil
}
