package service

import (
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

func smsChannelInfoFromModel(v *ai_boilerplate_model.SmsChannel) *pb.SmsChannelInfo {
	if v == nil {
		return &pb.SmsChannelInfo{}
	}
	return &pb.SmsChannelInfo{
		Id:           v.ID,
		Name:         v.Name,
		Operator:     v.Operator,
		Remark:       v.Remark,
		APIKey:       v.APIKey,
		APISecret:    maskSecret(v.APISecret),
		CallbackURL:  v.CallbackURL,
		Status:       int32(v.Status),
		CreatedAt:    formatRFC3339(v.CreatedAt),
		UpdatedAt:    formatRFC3339(v.UpdatedAt),
		OperatorName: constant.SmsChannelCodeToName[v.Operator],
	}
}

func maskSecret(secret string) string {
	if secret == "" {
		return ""
	}
	if len(secret) <= 4 {
		return "****"
	}
	if len(secret) <= 8 {
		return secret[:2] + "****" + secret[len(secret)-2:]
	}
	return secret[:4] + "****" + secret[len(secret)-4:]
}

func formatRFC3339(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
