package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/gopkg/jwt"
	"github.com/go-kratos/kratos/v2/log"
)

func NewDeviceRepo(
	logger log.Logger,
	data *Data,
	deviceRepo *ai_boilerplate_repo.DeviceRepo,
) *DeviceRepo {
	l := log.NewHelper(log.With(logger, "module", "data/device"))
	jwt := jwt.NewJwt(&jwt.Config{
		AccessSecret: data.cfg.GetBusiness()["jwt"].GetFields()["kid"].GetStructValue().GetFields()["accessSecret"].GetStringValue(),
		AccessExpire: int64(data.cfg.GetBusiness()["jwt"].GetFields()["kid"].GetStructValue().GetFields()["accessExpire"].GetNumberValue()),
		RefreshAfter: int64(data.cfg.GetBusiness()["jwt"].GetFields()["kid"].GetStructValue().GetFields()["refreshAfter"].GetNumberValue()),
		Issuer:       data.cfg.GetBusiness()["jwt"].GetFields()["kid"].GetStructValue().GetFields()["issuer"].GetStringValue(),
	}, jwt.NewRueidisCache(data.rueidis))
	return &DeviceRepo{
		log:        l,
		data:       data,
		DeviceRepo: deviceRepo,
		jwt:        jwt,
	}
}

type DeviceRepo struct {
	log  *log.Helper
	data *Data
	jwt  *jwt.Jwt
	*ai_boilerplate_repo.DeviceRepo
}

// GenerateToken 生成token
func (r *DeviceRepo) GenerateToken(_ context.Context, sn string) (*jwt.Token, error) {
	token, _, err := r.jwt.GenerateToken(map[string]any{
		"uid": sn,
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CheckToken 检查token
func (r *DeviceRepo) CheckToken(_ context.Context, token string) (map[string]any, error) {
	claims, err := r.jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
