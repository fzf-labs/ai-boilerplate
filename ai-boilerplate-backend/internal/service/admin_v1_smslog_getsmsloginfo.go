package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSmsLogInfo 短信日志-单条数据查询
func (a *AdminV1SmsLogService) GetSmsLogInfo(ctx context.Context, req *pb.GetSmsLogInfoReq) (*pb.GetSmsLogInfoReply, error) {
	resp := &pb.GetSmsLogInfoReply{}
	data, err := a.smsLogRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SmsLogInfo{
		Id:               data.ID,
		SmsChannelId:     data.SmsChannelID,
		SmsTemplateId:    data.SmsTemplateID,
		SmsParamsContent: data.SmsParamsContent,
		Mobile:           data.Mobile,
		UserId:           data.UserID,
		SendStatus:       data.SendStatus,
		SendTime:         timeutil.RFC3339(data.SendTime),
		ReceiveStatus:    data.ReceiveStatus,
		ReceiveTime:      timeutil.RFC3339(data.ReceiveTime.Time),
		APISendCode:      data.APISendCode,
		APISendMsg:       data.APISendMsg,
		APIRequestID:     data.APIRequestID,
		APISerialNo:      data.APISerialNo,
		APIReceiveCode:   data.APIReceiveCode,
		APIReceiveMsg:    data.APIReceiveMsg,
		CreatedAt:        timeutil.RFC3339(data.CreatedAt),
	}
	return resp, nil
}
