package data

import (
	"context"
	"fmt"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/gopkg/jwt"
	"github.com/fzf-labs/goutil/cryptutil"
	"github.com/fzf-labs/goutil/uuidutil"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

func NewUserRepo(
	logger log.Logger,
	data *Data,
	userRepo *ai_boilerplate_repo.UserRepo,
) *UserRepo {
	l := log.NewHelper(log.With(logger, "module", "data/user"))
	jwt := jwt.NewJwt(&jwt.Config{
		AccessSecret: data.cfg.GetBusiness()["jwt"].GetFields()["parent"].GetStructValue().GetFields()["accessSecret"].GetStringValue(),
		AccessExpire: int64(data.cfg.GetBusiness()["jwt"].GetFields()["parent"].GetStructValue().GetFields()["accessExpire"].GetNumberValue()),
		RefreshAfter: int64(data.cfg.GetBusiness()["jwt"].GetFields()["parent"].GetStructValue().GetFields()["refreshAfter"].GetNumberValue()),
		Issuer:       data.cfg.GetBusiness()["jwt"].GetFields()["parent"].GetStructValue().GetFields()["issuer"].GetStringValue(),
	}, jwt.NewRueidisCache(data.rueidis))
	return &UserRepo{
		log:      l,
		data:     data,
		UserRepo: userRepo,
		jwt:      jwt,
	}
}

type UserRepo struct {
	log  *log.Helper
	data *Data
	jwt  *jwt.Jwt
	*ai_boilerplate_repo.UserRepo
}

// GenerateNicknameByPhone 手机尾号生成昵称
func (u *UserRepo) GenerateNicknameByPhone(phone string) string {
	return fmt.Sprintf("用户%s", phone[len(phone)-4:])
}

// GenerateSalt 生成 salt
func (u *UserRepo) GenerateSalt() string {
	return uuidutil.KSUIDByTime()
}

// GeneratePassword 生成 password
func (u *UserRepo) GeneratePassword(salt string, password string) (string, error) {
	password, err := cryptutil.Encrypt(salt + password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// VerifyPassword 验证 password
func (u *UserRepo) VerifyPassword(salt string, password string, hash string) bool {
	return cryptutil.Compare(hash, salt+password) == nil
}

// GenerateToken 生成token
func (u *UserRepo) GenerateToken(ctx context.Context, userID string, wxGzhUserID string, wxGzhXcxID string) (*jwt.Token, error) {
	_ = ctx

	token, _, err := u.jwt.GenerateToken(map[string]any{
		"uid":         userID,
		"wxGzhUserId": wxGzhUserID,
		"wxGzhXcxId":  wxGzhXcxID,
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CheckToken 检查token
func (u *UserRepo) CheckToken(ctx context.Context, token string) (map[string]any, error) {
	_ = ctx

	claims, err := u.jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// JwtTokenClear 清理当前 user 的 token
func (u *UserRepo) JwtTokenClear(ctx context.Context, userID string) error {
	_ = ctx

	return u.jwt.JwtTokenClear(userID)
}

// RefreshToken 刷新token
func (u *UserRepo) RefreshToken(ctx context.Context, refreshToken string) (*jwt.Token, error) {
	_ = ctx

	claims, err := u.jwt.ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}
	if ok, token := u.jwt.JwtBlackTokenCheck(claims); ok && token != nil {
		return token, nil
	}
	token, _, err := u.jwt.RefreshToken(claims)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// BuildDeletedUserFields 构建账号注销后的用户脱敏字段。
func (u *UserRepo) BuildDeletedUserFields(userID string) map[string]interface{} {
	deletedPhone := fmt.Sprintf("deleted_%s", userID)
	return map[string]interface{}{
		"phone":          deletedPhone,
		"password":       "",
		"salt":           u.GenerateSalt(),
		"nickname":       "已注销用户",
		"gender":         int32(0),
		"avatar":         "",
		"profile":        "",
		"other":          nil,
		"wx_gzh_user_id": "",
		"wx_gzh_xcx_id":  "",
		"status":         int32(constant.StatusDisable),
	}
}

// UserIdToNickname 根据userIds查询用户昵称
func (u *UserRepo) UserIDToNickname(ctx context.Context, userIDs []string) (map[string]string, error) {
	resp := make(map[string]string)
	userIDs = lo.Filter(userIDs, func(item string, _ int) bool {
		return item != ""
	})
	userIDs = lo.Uniq(userIDs)
	if len(userIDs) == 0 {
		return resp, nil
	}
	users, err := u.FindMultiCacheByIDS(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		resp[user.ID] = user.Nickname
	}
	return resp, nil
}
