package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_dao"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
)

func NewWxGzhMaterialRepo(
	logger log.Logger,
	data *Data,
	wxGzhMaterialRepo *ai_boilerplate_repo.WxGzhMaterialRepo,
) *WxGzhMaterialRepo {
	l := log.NewHelper(log.With(logger, "module", "data/wxGzhMaterial"))
	return &WxGzhMaterialRepo{
		log:               l,
		data:              data,
		WxGzhMaterialRepo: wxGzhMaterialRepo,
	}
}

type WxGzhMaterialRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.WxGzhMaterialRepo
}

type WxGzhMaterialStatsItem struct {
	Type  string `gorm:"column:type"`
	Count int32  `gorm:"column:count"`
}

// GetWxGzhMaterialStats 获取素材统计信息
func (r *WxGzhMaterialRepo) GetWxGzhMaterialStats(ctx context.Context, appID string) (map[string]int32, error) {
	result := make(map[string]int32)
	var tmp []*WxGzhMaterialStatsItem
	dao := ai_boilerplate_dao.Use(r.data.gorm).WxGzhMaterial
	err := dao.WithContext(ctx).Select(dao.Type, dao.ID.Count().As("count")).Where(dao.AppID.Eq(appID)).Group(dao.Type).Scan(&tmp)
	if err != nil {
		return nil, err
	}
	for _, v := range tmp {
		result[v.Type] = v.Count
	}
	return result, nil
}
