package service

import (
	"context"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/fzf-labs/goutil/uuidutil"
	"github.com/fzf-labs/kratos-contrib/meta"
)

// GetNotificationSettings 获取用户通知设置
func (a *AppV1UserNotificationSettingService) GetNotificationSettings(ctx context.Context, _ *pb.GetNotificationSettingsReq) (*pb.GetNotificationSettingsReply, error) {
	resp := &pb.GetNotificationSettingsReply{}

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

	// 如果不存在，创建默认设置
	if len(settings) == 0 {
		defaultSettings := &ai_boilerplate_model.UserNotificationSetting{
			ID:                   uuidutil.GenUUID(),
			UserID:               userID,
			SystemNotification:   true,
			ActivityNotification: true,
			OrderNotification:    true,
			DndStartTime:         "22:00",
			DndEndTime:           "08:00",
		}

		err = a.userNotificationSettingRepo.CreateOne(ctx, defaultSettings)
		if err != nil {
			return nil, pb.ErrorReasonDataSQLError(pb.WithError(err))
		}

		resp.Settings = &pb.NotificationSettingsInfo{
			Id:                   defaultSettings.ID,
			UserId:               defaultSettings.UserID,
			SystemNotification:   defaultSettings.SystemNotification,
			ActivityNotification: defaultSettings.ActivityNotification,
			OrderNotification:    defaultSettings.OrderNotification,
			DndStartTime:         defaultSettings.DndStartTime,
			DndEndTime:           defaultSettings.DndEndTime,
			CreatedAt:            defaultSettings.CreatedAt.Format(time.RFC3339),
			UpdatedAt:            defaultSettings.UpdatedAt.Format(time.RFC3339),
		}
	} else {
		setting := settings[0]
		resp.Settings = &pb.NotificationSettingsInfo{
			Id:                   setting.ID,
			UserId:               setting.UserID,
			SystemNotification:   setting.SystemNotification,
			ActivityNotification: setting.ActivityNotification,
			OrderNotification:    setting.OrderNotification,
			DndStartTime:         setting.DndStartTime,
			DndEndTime:           setting.DndEndTime,
			CreatedAt:            setting.CreatedAt.Format(time.RFC3339),
			UpdatedAt:            setting.UpdatedAt.Format(time.RFC3339),
		}
	}

	return resp, nil
}
