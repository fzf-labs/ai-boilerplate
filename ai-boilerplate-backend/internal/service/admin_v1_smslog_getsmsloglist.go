package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSmsLogList 短信日志-列表数据查询
func (a *AdminV1SmsLogService) GetSmsLogList(ctx context.Context, req *pb.GetSmsLogListReq) (*pb.GetSmsLogListReply, error) {
	resp := &pb.GetSmsLogListReply{
		Total: 0,
		List:  []*pb.SmsLogInfo{},
	}
	param := &condition.Req{
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
		Query:    []*condition.QueryParam{},
		Order: []*condition.OrderParam{
			{
				Field: "created_at",
				Order: condition.DESC,
			},
		},
	}
	list, p, err := a.smsLogRepo.FindMultiCacheByCondition(ctx, param)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Total = p.Total
	if len(list) > 0 {
		for _, v := range list {
			resp.List = append(resp.List, &pb.SmsLogInfo{
				Id:               v.ID,
				SmsChannelId:     v.SmsChannelID,
				SmsTemplateId:    v.SmsTemplateID,
				SmsParamsContent: v.SmsParamsContent,
				Mobile:           v.Mobile,
				UserId:           v.UserID,
				SendStatus:       v.SendStatus,
				SendTime:         timeutil.RFC3339(v.SendTime),
				ReceiveStatus:    v.ReceiveStatus,
				ReceiveTime:      timeutil.RFC3339(v.ReceiveTime.Time),
				APISendCode:      v.APISendCode,
				APISendMsg:       v.APISendMsg,
				APIRequestID:     v.APIRequestID,
				APISerialNo:      v.APISerialNo,
				APIReceiveCode:   v.APIReceiveCode,
				APIReceiveMsg:    v.APIReceiveMsg,
				CreatedAt:        v.CreatedAt.Format(time.RFC3339),
			})
		}
	}
	return resp, nil
}
