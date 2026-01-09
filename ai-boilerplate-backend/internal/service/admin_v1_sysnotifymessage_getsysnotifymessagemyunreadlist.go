package service

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/samber/lo"
)

// GetSysNotifyMessageMyUnreadList 系统-通知消息-我的-未读列表
func (a *AdminV1SysNotifyMessageService) GetSysNotifyMessageMyUnreadList(ctx context.Context, req *pb.GetSysNotifyMessageMyUnreadListReq) (*pb.GetSysNotifyMessageMyUnreadListReply, error) {
	_ = req

	resp := &pb.GetSysNotifyMessageMyUnreadListReply{
		List: make([]*pb.SysNotifyMessageInfo, 0),
	}
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	list, err := a.sysNotifyMessageRepo.FindMultiCacheByReceiverReadTime(ctx, adminID, "")
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	if len(list) == 0 {
		return resp, nil
	}
	// 发送时间排序
	sort.Slice(list, func(i, j int) bool {
		return list[i].SendTime > list[j].SendTime
	})
	if len(list) > 0 {
		adminIDs := make([]string, 0)
		adminIDToNickname := make(map[string]string)
		adminIDToAvatar := make(map[string]string)
		for _, v := range list {
			adminIDs = append(adminIDs, v.Receiver, v.Sender)
		}
		adminIDs = lo.Uniq(adminIDs)
		admins, err := a.sysAdminRepo.FindMultiCacheByIDS(ctx, adminIDs)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}
		for _, v := range admins {
			adminIDToNickname[v.ID] = v.Nickname
			adminIDToAvatar[v.ID] = v.Avatar
		}
		for _, v := range list {
			resp.List = append(resp.List, &pb.SysNotifyMessageInfo{
				Id:             v.ID,
				TenantId:       v.TenantID,
				Type:           v.Type,
				Subject:        v.Subject,
				Content:        v.Content,
				Sender:         v.Sender,
				Receiver:       v.Receiver,
				SendTime:       v.SendTime,
				ReadTime:       v.ReadTime,
				Extend:         v.Extend.String(),
				SenderName:     adminIDToNickname[v.Sender],
				SenderAvatar:   adminIDToAvatar[v.Sender],
				ReceiverName:   adminIDToNickname[v.Receiver],
				ReceiverAvatar: adminIDToAvatar[v.Receiver],
			})
		}
	}
	return resp, nil
}
