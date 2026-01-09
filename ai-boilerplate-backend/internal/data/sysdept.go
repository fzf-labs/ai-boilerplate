package data

import (
	"context"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

func NewSysDeptRepo(
	logger log.Logger,
	data *Data,
	sysDeptRepo *ai_boilerplate_repo.SysDeptRepo,
) *SysDeptRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysDept"))
	return &SysDeptRepo{
		log:         l,
		data:        data,
		SysDeptRepo: sysDeptRepo,
	}
}

type SysDeptRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysDeptRepo
}

// DeptIDToName 根据deptIDs获取deptName
func (r *SysDeptRepo) DeptIDToName(ctx context.Context, deptIDs []string) (map[string]string, error) {
	deptIDs = lo.Filter(deptIDs, func(item string, _ int) bool {
		return item != ""
	})
	deptIDs = lo.Uniq(deptIDs)
	if len(deptIDs) == 0 {
		return map[string]string{}, nil
	}
	deptMap, err := r.FindMultiCacheByIDS(ctx, deptIDs)
	if err != nil {
		return nil, err
	}
	deptNameMap := make(map[string]string)
	for _, dept := range deptMap {
		deptNameMap[dept.ID] = dept.Name
	}
	return deptNameMap, nil
}
