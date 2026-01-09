package service

import (
	"context"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/timeutil"
)

// GetSmsChannelInfo 短信渠道-单条数据查询
func (a *AdminV1SmsChannelService) GetSmsChannelInfo(ctx context.Context, req *pb.GetSmsChannelInfoReq) (*pb.GetSmsChannelInfoReply, error) {
	resp := &pb.GetSmsChannelInfoReply{}
	data, err := a.smsChannelRepo.FindOneCacheByID(ctx, req.GetId())
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}
	resp.Info = &pb.SmsChannelInfo{
		Id:           data.ID,
		Name:         data.Name,
		Operator:     data.Operator,
		Remark:       data.Remark,
		APIKey:       data.APIKey,
		APISecret:    data.APISecret,
		CallbackURL:  data.CallbackURL,
		Status:       int32(data.Status),
		CreatedAt:    timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:    timeutil.RFC3339(data.UpdatedAt),
		OperatorName: constant.SmsChannelCodeToName[data.Operator],
	}
	return resp, nil
}
