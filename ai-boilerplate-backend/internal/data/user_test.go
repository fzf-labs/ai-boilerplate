package data

import (
	"strings"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/constant"
)

func TestBuildDeletedUserFields(t *testing.T) {
	t.Parallel()

	repo := &UserRepo{}
	const userID = "0190e7f4-1c17-70ee-85e5-d6a963df4821"

	got := repo.BuildDeletedUserFields(userID)

	phone, ok := got["phone"].(string)
	if !ok || !strings.Contains(phone, userID) || !strings.HasPrefix(phone, "deleted_") {
		t.Fatalf("unexpected deleted phone: %#v", got["phone"])
	}
	if got["password"] != "" {
		t.Fatalf("password should be cleared, got %#v", got["password"])
	}
	if got["nickname"] != "已注销用户" {
		t.Fatalf("nickname should be anonymized, got %#v", got["nickname"])
	}
	if got["avatar"] != "" || got["profile"] != "" || got["other"] != nil {
		t.Fatalf("profile fields should be cleared: %#v", got)
	}
	if got["wx_gzh_user_id"] != "" || got["wx_gzh_xcx_id"] != "" {
		t.Fatalf("third-party bindings should be cleared: %#v", got)
	}
	if got["status"] != int32(constant.StatusDisable) {
		t.Fatalf("status = %#v, want disabled", got["status"])
	}
	if salt, ok := got["salt"].(string); !ok || salt == "" {
		t.Fatalf("salt should be regenerated, got %#v", got["salt"])
	}
}
