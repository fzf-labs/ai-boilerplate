package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
	"github.com/samber/lo"
)

// GetSysNotifyMessageMyList 系统-通知消息-我的-列表
func (a *AdminV1SysNotifyMessageService) GetSysNotifyMessageMyList(ctx context.Context, req *pb.GetSysNotifyMessageMyListReq) (*pb.GetSysNotifyMessageMyListReply, error) {
	resp := &pb.GetSysNotifyMessageMyListReply{}
	tenantID := meta.GetMetadataFromClient(ctx, constant.XMdTenantID)
	adminID := meta.GetMetadataFromClient(ctx, constant.XMdAdminID)
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query: []*condition.QueryParam{
			{
				Field: "tenant_id",
				Value: tenantID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
			{
				Field: "receiver",
				Value: adminID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
		Order: []*condition.OrderParam{
			{
				Field: "send_time",
				Order: condition.DESC,
			},
		},
	}
	if req.GetSender() != "" {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "sender",
			Value: req.GetSender(),
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	if len(req.GetSendTime()) > 0 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "send_time",
			Value: req.GetSendTime()[0],
			Exp:   condition.GTE,
			Logic: condition.AND,
		})
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "send_time",
			Value: req.GetSendTime()[1],
			Exp:   condition.LTE,
			Logic: condition.AND,
		})
	}
	if req.GetReadStatus() == 1 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "read_time",
			Value: "",
			Exp:   condition.NEQ,
			Logic: condition.AND,
		})
	}
	if req.GetReadStatus() == -1 {
		param.Query = append(param.Query, &condition.QueryParam{
			Field: "read_time",
			Value: "",
			Exp:   condition.EQ,
			Logic: condition.AND,
		})
	}
	list, p, err := a.sysNotifyMessageRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
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
				ReceiverName:   adminIDToNickname[v.Receiver],
				ReceiverAvatar: adminIDToAvatar[v.Receiver],
				SenderName:     adminIDToNickname[v.Sender],
				SenderAvatar:   adminIDToAvatar[v.Sender],
			})
		}
	}
	return resp, nil
}
