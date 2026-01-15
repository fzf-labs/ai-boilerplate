package data

import (
	"context"
	"encoding/json"

	v1 "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

func NewMallProductRepo(
	logger log.Logger,
	data *Data,
	mallProductRepo *ai_boilerplate_repo.MallProductRepo,
) *MallProductRepo {
	l := log.NewHelper(log.With(logger, "module", "data/mallProduct"))
	return &MallProductRepo{
		log:             l,
		data:            data,
		MallProductRepo: mallProductRepo,
	}
}

type MallProductRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.MallProductRepo
}

// ProductIDToProductName 根据productIDs查询商品名称
func (m *MallProductRepo) ProductIDToProductName(ctx context.Context, productIDs []string) (map[string]string, error) {
	resp := make(map[string]string)
	productIDs = lo.Filter(productIDs, func(item string, _ int) bool {
		return item != ""
	})
	productIDs = lo.Uniq(productIDs)
	if len(productIDs) == 0 {
		return resp, nil
	}
	result, err := m.FindMultiCacheByIDS(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		resp[v.ID] = v.ProductName
	}
	return resp, nil
}

// GetMembershipProductConfig 解析会员产品配置
func (m *MallProductRepo) GetMembershipProductConfig(product *ai_boilerplate_model.MallProduct) (*v1.ProductConfig, error) {
	var cfg v1.ProductConfig
	if product.ProductConfig.String() != "" {
		if err := json.Unmarshal(product.ProductConfig, &cfg); err != nil {
			return nil, err
		}
	}
	return &cfg, nil
}
