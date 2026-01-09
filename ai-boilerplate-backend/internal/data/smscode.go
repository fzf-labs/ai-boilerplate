// Package data 提供数据访问层实现
//
// SmsCodeRepo 是一个专门的短信验证码管理仓库，支持多种场景（登录、注册、绑定、重置密码）
//
// 使用示例：
//
//	// 1. 检查发送频率限制
//	if err := repo.CheckSmsCodeFrequency(ctx, SmsCodeSceneLogin, "13800138000"); err != nil {
//		return err
//	}
//
//	// 2. 生成验证码
//	code, err := repo.GenerateSmsCode(SmsCodeSceneLogin)
//	if err != nil {
//		return err
//	}
//
//	// 3. 存储验证码
//	codeData := &SmsCodeData{
//		Scene:  SmsCodeSceneLogin,
//		CodeId: "login_" + uuid.New().String(),
//		Code:   code,
//		Phone:  "13800138000",
//		Uid:    "",
//	}
//	if err := repo.SetSmsCode(ctx, codeData); err != nil {
//		return err
//	}
//
//	// 4. 通过回调函数发送短信
//	err = repo.SendSmsCode(ctx, codeData, func(ctx context.Context, data *SmsCodeData) error {
//		// 在这里实现实际的短信发送逻辑
//		return sendSmsToUser(data.Phone, data.Code)
//	})
//
//	// 5. 验证验证码
//	verifiedData, err := repo.CheckSmsCode(ctx, SmsCodeSceneLogin, codeId, userInputCode)
//	if err != nil {
//		return err // 验证失败
//	}
//
//	// 6. 消费验证码（验证成功后删除）
//	if err := repo.ConsumeSmsCode(ctx, SmsCodeSceneLogin, codeId); err != nil {
//		// 不影响主流程，只记录日志
//	}
package data

import (
	"context"
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/goutil/jsonutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/redis/rueidis"
)

// SmsCodeScene 短信验证码场景
type SmsCodeScene string

const (
	SmsCodeSceneLogin    SmsCodeScene = "sms_login"    // 短信登录
	SmsCodeSceneRegister SmsCodeScene = "sms_register" // 短信注册
	SmsCodeSceneBind     SmsCodeScene = "sms_bind"     // 短信绑定
	SmsCodeSceneReset    SmsCodeScene = "sms_reset"    // 短信重置密码
)

// 验证码生成相关常量
const (
	digitCharset = "0123456789" // 数字字符集
)

// SmsCodeConfig 短信验证码配置
type SmsCodeConfig struct {
	Scene        SmsCodeScene  // 使用场景
	HourlyLimit  int           // 每小时发送限制
	DailyLimit   int           // 每天发送限制
	CodeTTL      time.Duration // 验证码有效期
	FrequencyTTL time.Duration // 频率限制TTL
	CodeLength   int           // 验证码长度（只支持数字）
}

// SmsCodeData 短信验证码数据结构
type SmsCodeData struct {
	Scene  SmsCodeScene `json:"scene"`  // 使用场景
	CodeID string       `json:"codeId"` // 验证码ID
	Code   string       `json:"code"`   // 验证码
	Phone  string       `json:"phone"`  // 手机号
	UID    string       `json:"uid"`    // 用户ID(可选)
}

func NewSmsCodeRepo(
	logger log.Logger,
	data *Data,
) *SmsCodeRepo {
	l := log.NewHelper(log.With(logger, "module", "data/smsCode"))
	return &SmsCodeRepo{
		log:  l,
		data: data,
	}
}

type SmsCodeRepo struct {
	log  *log.Helper
	data *Data
}

// GetSmsConfig 获取短信验证码配置
func (s *SmsCodeRepo) GetSmsConfig(scene SmsCodeScene) SmsCodeConfig {
	baseConfig := SmsCodeConfig{
		Scene:        scene,
		HourlyLimit:  3,
		DailyLimit:   10,
		CodeTTL:      5 * time.Minute,
		FrequencyTTL: 24 * time.Hour,
		CodeLength:   6, // 6位数字验证码
	}
	// 根据场景调整配置
	switch scene {
	case SmsCodeSceneLogin:
		baseConfig.HourlyLimit = 5
		baseConfig.DailyLimit = 15
	case SmsCodeSceneRegister:
		baseConfig.HourlyLimit = 3
		baseConfig.DailyLimit = 10
	case SmsCodeSceneBind:
		baseConfig.HourlyLimit = 3
		baseConfig.DailyLimit = 5
	case SmsCodeSceneReset:
		baseConfig.HourlyLimit = 2
		baseConfig.DailyLimit = 5
		baseConfig.CodeTTL = 10 * time.Minute // 重置密码验证码有效期更长
	}

	return baseConfig
}

// CheckSmsCodeFrequency 检查短信验证码发送频率限制
func (s *SmsCodeRepo) CheckSmsCodeFrequency(ctx context.Context, scene SmsCodeScene, phone string) error {
	now := time.Now()
	date := now.Format("2006-01-02")
	hour := now.Hour()
	config := s.GetSmsConfig(scene)
	cacheKey := constant.UserSmsCodeFrequency.Key(phone, date, string(scene))
	// 检查小时限制
	hourCountStr, err := s.data.rueidis.Do(ctx, s.data.rueidis.B().Hget().Key(cacheKey).Field(strconv.Itoa(hour)).Build()).AsInt64()
	hourCount := int(hourCountStr)
	if err != nil && !rueidis.IsRedisNil(err) {
		return pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	if hourCount >= config.HourlyLimit {
		return pb.ErrorReasonSmsFrequencyLimit(pb.WithFmtMsg(fmt.Sprintf("一小时内最多发送%d次短信验证码", config.HourlyLimit)))
	}
	// 检查日限制
	dayCountStr, err := s.data.rueidis.Do(ctx, s.data.rueidis.B().Hget().Key(cacheKey).Field("day").Build()).AsInt64()
	dayCount := int(dayCountStr)
	if err != nil && !rueidis.IsRedisNil(err) {
		return pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	if dayCount >= config.DailyLimit {
		return pb.ErrorReasonSmsFrequencyLimit(pb.WithFmtMsg(fmt.Sprintf("一天内最多发送%d次短信验证码", config.DailyLimit)))
	}
	return nil
}

// SetSmsCodeFrequency 设置短信验证码发送频率计数
func (s *SmsCodeRepo) SetSmsCodeFrequency(ctx context.Context, scene SmsCodeScene, phone string) error {
	now := time.Now()
	date := now.Format("2006-01-02")
	hour := now.Hour()
	config := s.GetSmsConfig(scene)
	cacheKey := constant.UserSmsCodeFrequency.Key(phone, date, string(scene))

	// 设置小时和天级别计数
	_ = s.data.rueidis.Do(ctx, s.data.rueidis.B().Hincrby().Key(cacheKey).Field(strconv.Itoa(hour)).Increment(1).Build()).Error()
	_ = s.data.rueidis.Do(ctx, s.data.rueidis.B().Hincrby().Key(cacheKey).Field("day").Increment(1).Build()).Error()

	// 设置过期时间
	_ = s.data.rueidis.Do(ctx, s.data.rueidis.B().Expire().Key(cacheKey).Seconds(int64(config.FrequencyTTL.Seconds())).Build()).Error()

	return nil
}

// GenerateSmsCode 生成短信验证码（只包含数字）
// 优化版本：使用更高效的随机数生成方式，减少系统调用次数
func (s *SmsCodeRepo) GenerateSmsCode(scene SmsCodeScene) (string, error) {
	config := s.GetSmsConfig(scene)
	// 直接使用 byte slice 构建结果，避免 strings.Builder 的额外开销
	result := make([]byte, config.CodeLength)
	// 一次性生成足够的随机字节，减少系统调用
	randomBytes := make([]byte, config.CodeLength)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate secure random bytes: %w", err)
	}
	// 将随机字节映射为数字字符
	for i := 0; i < config.CodeLength; i++ {
		// 使用模运算确保索引在有效范围内
		result[i] = digitCharset[randomBytes[i]%10] // 直接使用 %10 因为是数字字符集
	}
	return string(result), nil
}

// GenerateSmsCode 生成短信验证码
func (s *SmsCodeRepo) GenerateSmsCodeData(scene SmsCodeScene, phone string, uid string) (*SmsCodeData, error) {
	code, err := s.GenerateSmsCode(scene)
	if err != nil {
		return nil, err
	}
	return &SmsCodeData{
		Scene:  scene,
		CodeID: uuid.New().String(),
		Code:   code,
		Phone:  phone,
		UID:    uid,
	}, nil
}

// 发送短信验证码
func (s *SmsCodeRepo) SendSmsCode(ctx context.Context, data *SmsCodeData, fn func(ctx context.Context, data *SmsCodeData) error) error {
	return fn(ctx, data)
}

// SetSmsCode 存储短信验证码
func (s *SmsCodeRepo) SetSmsCode(ctx context.Context, data *SmsCodeData) error {
	config := s.GetSmsConfig(data.Scene)
	cacheValue, err := jsonutil.Marshal(data)
	if err != nil {
		return pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	cacheKey := constant.UserSmsCode.Key(string(data.Scene), data.CodeID)
	return s.data.rueidis.Do(ctx, s.data.rueidis.B().Set().Key(cacheKey).Value(string(cacheValue)).Ex(config.CodeTTL).Build()).Error()
}

// CheckSmsCode 验证短信验证码
func (s *SmsCodeRepo) CheckSmsCode(ctx context.Context, scene SmsCodeScene, codeID string, inputCode string) (*SmsCodeData, error) {
	cacheKey := constant.UserSmsCode.Key(string(scene), codeID)
	cacheValue, err := s.data.rueidis.Do(ctx, s.data.rueidis.B().Get().Key(cacheKey).Build()).ToString()
	if err != nil {
		if rueidis.IsRedisNil(err) {
			return nil, pb.ErrorReasonSmsCodeInvalid(pb.WithFmtMsg("验证码不存在或已过期"))
		}
		return nil, pb.ErrorReasonDataRedisErr(pb.WithError(err))
	}
	var data SmsCodeData
	err = jsonutil.Unmarshal([]byte(cacheValue), &data)
	if err != nil {
		return nil, pb.ErrorReasonDataFormattingError(pb.WithError(err))
	}
	// 检查验证码是否匹配
	if data.Code != inputCode {
		return nil, pb.ErrorReasonSmsCodeInvalid(pb.WithFmtMsg("验证码错误"))
	}
	// 检查场景是否匹配
	if data.Scene != scene {
		return nil, pb.ErrorReasonSmsCodeInvalid(pb.WithFmtMsg("验证码场景不匹配"))
	}
	return &data, nil
}

// ClearSmsCode 清除短信验证码
func (s *SmsCodeRepo) ClearSmsCode(ctx context.Context, scene SmsCodeScene, codeID string) error {
	cacheKey := constant.UserSmsCode.Key(string(scene), codeID)
	return s.data.rueidis.Do(ctx, s.data.rueidis.B().Del().Key(cacheKey).Build()).Error()
}
