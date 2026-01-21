package service

import (
	"encoding/json"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/datatypes"
)

func NewAppV1UserNotificationSettingService(
	logger log.Logger,
	userNotificationSettingRepo *data.UserNotificationSettingRepo,
) *AppV1UserNotificationSettingService {
	l := log.NewHelper(log.With(logger, "module", "service/userNotificationSetting"))
	return &AppV1UserNotificationSettingService{
		log:                         l,
		userNotificationSettingRepo: userNotificationSettingRepo,
	}
}

type AppV1UserNotificationSettingService struct {
	pb.UnimplementedUserNotificationSettingServer
	log                         *log.Helper
	userNotificationSettingRepo *data.UserNotificationSettingRepo
}

type notificationCategoryMeta struct {
	Key         string
	Title       string
	Description string
}

var defaultNotificationCategories = []notificationCategoryMeta{
	{
		Key:         "activity",
		Title:       "活动通知",
		Description: "活动与运营提醒",
	},
	{
		Key:         "service",
		Title:       "服务通知",
		Description: "服务进度与售后提醒",
	},
	{
		Key:         "interaction",
		Title:       "互动通知",
		Description: "评论与互动提醒",
	},
	{
		Key:         "direct_message",
		Title:       "私信通知",
		Description: "新的私信与回复提醒",
	},
}

func defaultNotificationPreferences() map[string]bool {
	preferences := make(map[string]bool, len(defaultNotificationCategories))
	for _, meta := range defaultNotificationCategories {
		preferences[meta.Key] = true
	}
	return preferences
}

func decodeNotificationPreferences(raw datatypes.JSON) map[string]bool {
	if len(raw) == 0 {
		return map[string]bool{}
	}
	var preferences map[string]bool
	if err := json.Unmarshal(raw, &preferences); err != nil {
		return map[string]bool{}
	}
	return preferences
}

func encodeNotificationPreferences(preferences map[string]bool) datatypes.JSON {
	if len(preferences) == 0 {
		return datatypes.JSON([]byte("{}"))
	}
	payload, err := json.Marshal(preferences)
	if err != nil {
		return datatypes.JSON([]byte("{}"))
	}
	return datatypes.JSON(payload)
}

func buildNotificationCategories(preferences map[string]bool) []*pb.NotificationCategory {
	categories := make([]*pb.NotificationCategory, 0, len(defaultNotificationCategories)+len(preferences))
	seen := make(map[string]struct{}, len(defaultNotificationCategories))
	for _, meta := range defaultNotificationCategories {
		enabled := preferences[meta.Key]
		categories = append(categories, &pb.NotificationCategory{
			Key:         meta.Key,
			Title:       meta.Title,
			Description: meta.Description,
			Enabled:     enabled,
		})
		seen[meta.Key] = struct{}{}
	}
	for key, enabled := range preferences {
		if _, ok := seen[key]; ok {
			continue
		}
		categories = append(categories, &pb.NotificationCategory{
			Key:     key,
			Title:   key,
			Enabled: enabled,
		})
	}
	return categories
}
