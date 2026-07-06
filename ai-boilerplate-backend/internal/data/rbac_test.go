package data

import (
	"context"
	"reflect"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
)

func TestSysMenuRepoTraversePermissionsDeduplicatesAndSorts(t *testing.T) {
	repo := &SysMenuRepo{}

	got, err := repo.TraversePermissions(context.Background(), []*ai_boilerplate_model.SysMenu{
		{ID: "menu-1", Permission: "sys:role:list"},
		{ID: "menu-2", Permission: "mall:order:pay"},
		{ID: "menu-3", Permission: "sys:role:list"},
		{ID: "menu-4", Permission: ""},
	})
	if err != nil {
		t.Fatalf("TraversePermissions() error = %v", err)
	}

	want := []string{"mall:order:pay", "sys:role:list"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TraversePermissions() = %#v, want %#v", got, want)
	}
}

func TestSysMenuRepoTraverseMenuTreeFiltersDisabledAndButtonNodes(t *testing.T) {
	repo := &SysMenuRepo{}

	tree, err := repo.TraverseMenuTree(context.Background(), []*ai_boilerplate_model.SysMenu{
		{
			ID:     "root",
			Name:   "Root",
			Type:   constant.SysMenuTypeDir.String(),
			Status: int16(constant.StatusEnable),
		},
		{
			ID:     "child",
			Pid:    "root",
			Name:   "Child",
			Type:   constant.SysMenuTypeMenu.String(),
			Status: int16(constant.StatusEnable),
		},
		{
			ID:     "button",
			Pid:    "child",
			Name:   "Button",
			Type:   constant.SysMenuTypeButton.String(),
			Status: int16(constant.StatusEnable),
		},
		{
			ID:     "disabled",
			Pid:    "root",
			Name:   "Disabled",
			Type:   constant.SysMenuTypeMenu.String(),
			Status: int16(constant.StatusDisable),
		},
	})
	if err != nil {
		t.Fatalf("TraverseMenuTree() error = %v", err)
	}
	if len(tree) != 1 || tree[0].Id != "root" {
		t.Fatalf("TraverseMenuTree() root = %#v, want root only", tree)
	}
	if len(tree[0].Children) != 1 || tree[0].Children[0].Id != "child" {
		t.Fatalf("TraverseMenuTree() children = %#v, want child only", tree[0].Children)
	}
	if len(tree[0].Children[0].Children) != 0 {
		t.Fatalf("TraverseMenuTree() should filter button children, got %#v", tree[0].Children[0].Children)
	}
}
