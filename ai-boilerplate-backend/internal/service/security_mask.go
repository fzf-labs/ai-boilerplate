package service

import (
	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/timeutil"
)

const sensitiveMask = "******"

func maskSensitiveValue(value string) string {
	if value == "" {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= 8 {
		return sensitiveMask
	}
	return string(runes[:4]) + sensitiveMask + string(runes[len(runes)-4:])
}

func smsChannelInfoFromModel(data *ai_boilerplate_model.SmsChannel) *pb.SmsChannelInfo {
	return &pb.SmsChannelInfo{
		Id:           data.ID,
		Name:         data.Name,
		Operator:     data.Operator,
		Remark:       data.Remark,
		APIKey:       maskSensitiveValue(data.APIKey),
		APISecret:    maskSensitiveValue(data.APISecret),
		CallbackURL:  data.CallbackURL,
		Status:       int32(data.Status),
		CreatedAt:    timeutil.RFC3339(data.CreatedAt),
		UpdatedAt:    timeutil.RFC3339(data.UpdatedAt),
		OperatorName: constant.SmsChannelCodeToName[data.Operator],
	}
}
