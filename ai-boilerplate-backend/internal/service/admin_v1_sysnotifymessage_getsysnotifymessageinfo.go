package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
)

// GetSysNotifyMessageInfo 系统-通知消息-单条数据查询
func (a *AdminV1SysNotifyMessageService) GetSysNotifyMessageInfo(ctx context.Context, req *pb.GetSysNotifyMessageInfoReq) (*pb.GetSysNotifyMessageInfoReply, error) {
	resp := &pb.GetSysNotifyMessageInfoReply{}
	data, err := a.sysNotifyMessageRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	adminIDToNickname := make(map[string]string)
	adminIDToAvatar := make(map[string]string)
	admins, err := a.sysAdminRepo.FindMultiCacheByIDS(ctx, []string{data.Receiver})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	for _, v := range admins {
		adminIDToNickname[v.ID] = v.Nickname
		adminIDToAvatar[v.ID] = v.Avatar
	}
	resp.Info = &pb.SysNotifyMessageInfo{
		Id:             data.ID,
		TenantId:       data.TenantID,
		Type:           data.Type,
		Subject:        data.Subject,
		Content:        data.Content,
		Receiver:       data.Receiver,
		SendTime:       data.SendTime,
		ReadTime:       data.ReadTime,
		Extend:         data.Extend.String(),
		ReceiverName:   adminIDToNickname[data.Receiver],
		ReceiverAvatar: adminIDToAvatar[data.Receiver],
		SenderName:     adminIDToNickname[data.Sender],
		SenderAvatar:   adminIDToAvatar[data.Sender],
	}
	return resp, nil
}
