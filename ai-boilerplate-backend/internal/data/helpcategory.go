package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/fzf-labs/godb/orm/condition"
	"github.com/go-kratos/kratos/v2/log"
)

func NewHelpCategoryRepo(
	logger log.Logger,
	data *Data,
	helpCategoryRepo *ai_boilerplate_repo.HelpCategoryRepo,
) *HelpCategoryRepo {
	l := log.NewHelper(log.With(logger, "module", "data/helpCategory"))
	return &HelpCategoryRepo{
		log:              l,
		data:             data,
		HelpCategoryRepo: helpCategoryRepo,
	}
}

type HelpCategoryRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.HelpCategoryRepo
}

// FindAllEnabled 查询所有启用的帮助分类
func (r *HelpCategoryRepo) FindAllEnabled(ctx context.Context) ([]*ai_boilerplate_model.HelpCategory, error) {
	condReq := &condition.Req{
		Query: []*condition.QueryParam{
			{Field: "status", Value: 1, Exp: condition.EQ, Logic: condition.AND},
		},
		Order: []*condition.OrderParam{
			{Field: "order", Order: condition.ASC},
		},
	}
	list, _, err := r.FindMultiByCondition(ctx, condReq)
	return list, err
}
