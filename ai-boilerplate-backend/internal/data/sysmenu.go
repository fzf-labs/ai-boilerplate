package data

import (
	"context"
	"sort"

	pb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_repo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

func NewSysMenuRepo(
	logger log.Logger,
	data *Data,
	sysMenuRepo *ai_boilerplate_repo.SysMenuRepo,
) *SysMenuRepo {
	l := log.NewHelper(log.With(logger, "module", "data/sysMenu"))
	return &SysMenuRepo{
		log:         l,
		data:        data,
		SysMenuRepo: sysMenuRepo,
	}
}

type SysMenuRepo struct {
	log  *log.Helper
	data *Data
	*ai_boilerplate_repo.SysMenuRepo
}

// TraverseMenuTree 遍历菜单树，构建树形结构
func (s *SysMenuRepo) TraverseMenuTree(_ context.Context, menus []*ai_boilerplate_model.SysMenu) ([]*pb.SysMenuItem, error) {
	// menus 只要目录和菜单的,并且只需要状态是开启的
	menus = lo.Filter(menus, func(menu *ai_boilerplate_model.SysMenu, _ int) bool {
		return lo.Contains([]string{constant.SysMenuTypeDir.String(), constant.SysMenuTypeMenu.String()}, menu.Type) && menu.Status == int16(constant.StatusEnable)
	})
	// 构建菜单树
	var menuTree []*pb.SysMenuItem
	// 创建一个映射，用于存储菜单ID到菜单项的映射
	menuMap := make(map[string]*pb.SysMenuItem)
	// 第一步：将所有菜单项添加到映射中
	for _, menu := range menus {
		item := &pb.SysMenuItem{
			Id:            menu.ID,
			Pid:           menu.Pid,
			Name:          menu.Name,
			Type:          menu.Type,
			Path:          menu.Path,
			Permission:    menu.Permission,
			Icon:          menu.Icon,
			Component:     menu.Component,
			ComponentName: menu.ComponentName,
			Sort:          int32(menu.Sort),
			Status:        int32(menu.Status),
			Children:      []*pb.SysMenuItem{},
		}
		menuMap[menu.ID] = item
	}
	// 第二步：构建树形结构
	for _, menu := range menus {
		if menu.Pid == "" {
			// 顶级菜单直接添加到结果中
			menuTree = append(menuTree, menuMap[menu.ID])
		} else {
			// 将子菜单添加到父菜单的children中
			if parent, exists := menuMap[menu.Pid]; exists {
				parent.Children = append(parent.Children, menuMap[menu.ID])
			}
		}
	}
	return menuTree, nil
}

// TraversePermissions 遍历菜单的权限,获取权限列表
func (s *SysMenuRepo) TraversePermissions(_ context.Context, menus []*ai_boilerplate_model.SysMenu) ([]string, error) {
	var permissions []string
	for _, menu := range menus {
		if menu.Permission != "" {
			permissions = append(permissions, menu.Permission)
		}
	}
	// 去重
	permissions = lo.Uniq(permissions)
	// 排序
	sort.Strings(permissions)
	return permissions, nil
}
