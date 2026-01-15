package data

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserMembershipRepo(
	logger log.Logger,
	data *Data,
	userMembershipRepo *ai_boilerplate_repo.UserMembershipRepo,
) *UserMembershipRepo {
	l := log.NewHelper(log.With(logger, "module", "data/userMembership"))
	return &UserMembershipRepo{
		log:                l,
		data:               data,
		UserMembershipRepo: userMembershipRepo,
	}
}

type UserMembershipRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.UserMembershipRepo
}

// membershipTypeLevel 会员类型等级映射（normal=普通用户，不是付费会员）
var membershipTypeLevel = map[string]int{
	constant.MembershipTypeNormal.String(): 0, // 普通用户（未开通会员）
	constant.MembershipTypeVip.String():    1, // VIP会员
	constant.MembershipTypeSvip.String():   2, // SVIP会员
}

// membershipConversionRate 付费会员升级折算比例
// key格式: "oldType->newType", value: 折算比例(旧会员1天折算新会员多少天)
// 注意: normal是普通用户，没有剩余天数概念，不需要配置折算比例
var membershipConversionRate = map[string]float64{
	"vip->svip": 0.5, // VIP会员1天 = SVIP会员0.5天
}

var (
	ErrMembershipTypeInvalid         = errors.New("invalid membership type")
	ErrMembershipDowngradeNotAllowed = errors.New("membership downgrade is not allowed")
)

// CalcMembershipChange 计算会员的类型和到期时间变更
// 允许的变更：
//   - normal -> vip/svip: 普通用户开通会员，从现在开始计算
//   - vip -> vip: VIP续期，在到期时间基础上延期
//   - svip -> svip: SVIP续期，在到期时间基础上延期
//   - vip -> svip: VIP升级SVIP，剩余VIP天数折算后加上新增天数
//
// 不允许的变更：svip -> vip, vip -> normal, svip -> normal（降级）
func (r *UserMembershipRepo) CalcMembershipChange(ctx context.Context, oldMembershipType string, oldExpiredAt time.Time, addMembershipType string, addDurationDays int) (string, time.Time, error) {
	// 普通用户没有剩余天数概念，到期时间设置为空
	if oldMembershipType == constant.MembershipTypeNormal.String() {
		oldExpiredAt = time.Time{}
	}
	// 验证会员类型是否有效
	oldLevel, oldOk := membershipTypeLevel[oldMembershipType]
	newLevel, newOk := membershipTypeLevel[addMembershipType]
	if !oldOk || !newOk {
		return "", time.Time{}, ErrMembershipTypeInvalid
	}

	// 不允许购买普通用户类型（normal不是商品）
	if addMembershipType == constant.MembershipTypeNormal.String() {
		return "", time.Time{}, ErrMembershipTypeInvalid
	}

	// 不允许降级（svip不能变vip，付费会员不能变普通用户）
	if newLevel < oldLevel {
		return "", time.Time{}, ErrMembershipDowngradeNotAllowed
	}

	now := time.Now()

	// 场景1: 普通用户开通会员（normal -> vip/svip）
	// 普通用户没有剩余天数概念，直接从现在开始计算
	if oldMembershipType == constant.MembershipTypeNormal.String() {
		newExpiredAt := now.AddDate(0, 0, addDurationDays)
		return addMembershipType, newExpiredAt, nil
	}

	// 场景2: 同类型续期（vip -> vip 或 svip -> svip）
	if oldMembershipType == addMembershipType {
		var newExpiredAt time.Time
		// 如果已过期，从现在开始计算；否则在到期时间基础上延期
		if oldExpiredAt.Before(now) || oldExpiredAt.IsZero() {
			newExpiredAt = now.AddDate(0, 0, addDurationDays)
		} else {
			newExpiredAt = oldExpiredAt.AddDate(0, 0, addDurationDays)
		}
		return addMembershipType, newExpiredAt, nil
	}

	// 场景3: VIP升级SVIP（vip -> svip）
	// 将VIP剩余天数按比例折算成SVIP天数，再加上新购买的天数
	var remainingDays float64 = 0

	// 计算VIP剩余天数
	if oldExpiredAt.After(now) {
		remainingDuration := oldExpiredAt.Sub(now)
		remainingDays = remainingDuration.Hours() / 24
	}

	// 获取折算比例
	conversionKey := oldMembershipType + "->" + addMembershipType
	rate, ok := membershipConversionRate[conversionKey]
	if !ok {
		// 如果没有配置折算比例，默认1:1
		rate = 1.0
	}

	// 折算后的天数 + 新增天数（向上取整，保护用户权益）
	convertedDays := remainingDays * rate
	totalDays := int(math.Ceil(convertedDays)) + addDurationDays

	// 从现在开始计算新的到期时间
	newExpiredAt := now.AddDate(0, 0, totalDays)

	return addMembershipType, newExpiredAt, nil
}

// GetUserActualMembershipInfo 获取用户的实际会员信息
func (r *UserMembershipRepo) GetUserActualMembershipInfo(ctx context.Context, userID string) (*ai_boilerplate_model.UserMembership, error) {
	defaultMembership := &ai_boilerplate_model.UserMembership{
		UserID:         userID,
		MembershipType: constant.MembershipTypeNormal.String(),
		Status:         int32(constant.StatusEnable),
	}
	userMembership, err := r.FindOneCacheByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	// 如果用户会员信息不存在，则创建普通会员
	if userMembership == nil || userMembership.ID == "" {
		err = r.CreateOneCache(ctx, defaultMembership)
		if err != nil {
			return nil, err
		}
		return defaultMembership, nil
	}
	// 计算用户的实际会员信息
	// 如果不是普通会员，检查是否已过期
	if userMembership.MembershipType != constant.MembershipTypeNormal.String() && userMembership.ExpiredAt.Valid && userMembership.ExpiredAt.Time.Before(time.Now()) {
		return defaultMembership, nil
	}
	return userMembership, nil
}
