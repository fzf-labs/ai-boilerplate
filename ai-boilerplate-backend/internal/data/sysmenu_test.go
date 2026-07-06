package data

import (
	"context"
	"reflect"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
)

func TestTraverseMenuTreeFiltersDisabledAndButtonMenus(t *testing.T) {
	t.Parallel()

	repo := &SysMenuRepo{}
	got, err := repo.TraverseMenuTree(context.Background(), testfixture.SysMenus())
	if err != nil {
		t.Fatalf("TraverseMenuTree() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("root menu count = %d, want 1: %#v", len(got), got)
	}
	if got[0].Id != "menu-root" {
		t.Fatalf("root menu id = %q", got[0].Id)
	}
	if len(got[0].Children) != 1 {
		t.Fatalf("root children count = %d, want enabled menu child only: %#v", len(got[0].Children), got[0].Children)
	}
	if got[0].Children[0].Id != "menu-child" {
		t.Fatalf("child menu id = %q", got[0].Children[0].Id)
	}
}

func TestTraversePermissionsDeduplicatesAndSorts(t *testing.T) {
	t.Parallel()

	repo := &SysMenuRepo{}
	got, err := repo.TraversePermissions(context.Background(), testfixture.SysMenus())
	if err != nil {
		t.Fatalf("TraversePermissions() error = %v", err)
	}

	want := []string{"dashboard:create", "dashboard:view"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("permissions = %#v, want %#v", got, want)
	}
}
