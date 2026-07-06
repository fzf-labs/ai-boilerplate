package auth

import (
	"context"
	"testing"

	adminpb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	apppb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

func TestAdminAuthWhiteListIncludesLoginAndLogoutOnly(t *testing.T) {
	matcher := whiteListMatcher(AdminPrefixPathToWhiteList)

	if matcher(context.Background(), adminpb.OperationSysAuthSysAuthLogin) {
		t.Fatal("admin login should not require auth middleware")
	}
	if matcher(context.Background(), adminpb.OperationSysAuthSysAuthLogout) {
		t.Fatal("admin logout should not require auth middleware")
	}
	if !matcher(context.Background(), adminpb.OperationSysAuthSysAuthPermission) {
		t.Fatal("admin permission endpoint should require auth middleware")
	}
}

func TestAppAuthWhiteListIncludesLoginAndVerifyCode(t *testing.T) {
	matcher := whiteListMatcher(AppPrefixPathToWhiteList)

	if matcher(context.Background(), apppb.OperationUserLogin) {
		t.Fatal("app login should not require auth middleware")
	}
	if matcher(context.Background(), apppb.OperationUserSendVerifyCode) {
		t.Fatal("app send verify code should not require auth middleware")
	}
	if !matcher(context.Background(), apppb.OperationUserDeleteAccount) {
		t.Fatal("app delete account should require auth middleware")
	}
}
