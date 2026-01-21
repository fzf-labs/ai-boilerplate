package service

import (
	"context"
	"errors"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// UpdateNotificationSettings 更新用户通知设置
func (a *AppV1UserNotificationSettingService) UpdateNotificationSettings(ctx context.Context, req *pb.UpdateNotificationSettingsReq) (*pb.UpdateNotificationSettingsReply, error) {
	resp := &pb.UpdateNotificationSettingsReply{}

	// 从上下文获取用户ID
	userID := meta.GetMetadataFromClient(ctx, constant.XMdUserID)

	// 查询用户的通知设置
	conditionReq := &condition.Req{
		Query: []*condition.QueryParam{
			{
				Field: "user_id",
				Value: userID,
				Exp:   condition.EQ,
				Logic: condition.AND,
			},
		},
	}

	settings, _, err := a.userNotificationSettingRepo.FindMultiByCondition(ctx, conditionReq)
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	if len(settings) == 0 {
		return nil, pb.ErrorReasonDataRecordNotFound(pb.WithError(errors.New("通知设置不存在")))
	}

	// 更新设置
	setting := settings[0]
	preferences := make(map[string]bool, len(req.GetPreferences()))
	for _, pref := range req.GetPreferences() {
		if pref.GetKey() == "" {
			continue
		}
		preferences[pref.GetKey()] = pref.GetEnabled()
	}
	if len(preferences) == 0 {
		return nil, pb.ErrorReasonParamError()
	}
	setting.NotificationPreferences = encodeNotificationPreferences(preferences)

	if req.DndStartTime != "" {
		setting.DndStartTime = req.DndStartTime
	}
	if req.DndEndTime != "" {
		setting.DndEndTime = req.DndEndTime
	}

	// 使用 UpsertOneByFields 更新记录（基于 ID 字段）
	err = a.userNotificationSettingRepo.UpsertOneByFields(ctx, setting, []string{"id"})
	if err != nil {
		return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
	}

	return resp, nil
}
