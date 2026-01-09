package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/gopkg/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

func NewSysAdminRepo(
	logger log.Logger,
	data *Data,
	sysAdminRepo *ai_boilerplate_repo.SysAdminRepo,
) *SysAdminRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysAdmin"))
	jwt := jwt.NewJwt(&jwt.Config{
		AccessSecret: data.cfg.GetBusiness()["jwt"].GetFields()["admin"].GetStructValue().GetFields()["accessSecret"].GetStringValue(),
		AccessExpire: int64(data.cfg.GetBusiness()["jwt"].GetFields()["admin"].GetStructValue().GetFields()["accessExpire"].GetNumberValue()),
		RefreshAfter: int64(data.cfg.GetBusiness()["jwt"].GetFields()["admin"].GetStructValue().GetFields()["refreshAfter"].GetNumberValue()),
		Issuer:       data.cfg.GetBusiness()["jwt"].GetFields()["admin"].GetStructValue().GetFields()["issuer"].GetStringValue(),
	}, jwt.NewRueidisCache(data.rueidis))
	return &SysAdminRepo{
		log:          l,
		data:         data,
		SysAdminRepo: sysAdminRepo,
		jwt:          jwt,
	}
}

type SysAdminRepo struct {
	log  *log.Helper
	jwt  *jwt.Jwt
	data *Data
	*ai_boilerplate_repo.SysAdminRepo
}

// GenerateToken 生成token
func (r *SysAdminRepo) GenerateToken(ctx context.Context, sysAdmin *ai_boilerplate_model.SysAdmin) (*jwt.Token, error) {
	_ = ctx

	token, _, err := r.jwt.GenerateToken(map[string]any{
		"uid":       sysAdmin.ID,
		"nickname":  sysAdmin.Nickname,
		"tenant_id": sysAdmin.TenantID,
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CheckToken 检查token
func (r *SysAdminRepo) CheckToken(ctx context.Context, token string) (map[string]any, error) {
	_ = ctx

	claims, err := r.jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// AdminIdToNickname 根据adminId获取adminName
func (r *SysAdminRepo) AdminIDToNickname(ctx context.Context, adminIDs []string) (map[string]string, error) {
	adminIDs = lo.Filter(adminIDs, func(item string, _ int) bool {
		return item != ""
	})
	adminIDs = lo.Uniq(adminIDs)
	if len(adminIDs) == 0 {
		return map[string]string{}, nil
	}
	adminMap, err := r.FindMultiCacheByIDS(ctx, adminIDs)
	if err != nil {
		return nil, err
	}
	adminNameMap := make(map[string]string)
	for _, admin := range adminMap {
		adminNameMap[admin.ID] = admin.Nickname
	}
	return adminNameMap, nil
}
