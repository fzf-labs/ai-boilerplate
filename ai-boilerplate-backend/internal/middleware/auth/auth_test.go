package auth

import (
	"context"
	"testing"

	adminpb "github.com/fzf-labs/ai-boilerplate-backend/api/admin/v1"
	apppb "github.com/fzf-labs/ai-boilerplate-backend/api/app/v1"
)

func TestWhiteListMatcher(t *testing.T) {
	t.Parallel()

	adminMatcher := whiteListMatcher(AdminPrefixPathToWhiteList)
	if adminMatcher(context.Background(), adminpb.OperationSysAuthSysAuthLogin) {
		t.Fatalf("admin login should be skipped by auth matcher")
	}
	if adminMatcher(context.Background(), adminpb.OperationSysAuthSysAuthLogout) {
		t.Fatalf("admin logout should be skipped by auth matcher")
	}
	if !adminMatcher(context.Background(), adminpb.OperationSysAuthSysAuthPermission) {
		t.Fatalf("admin permission endpoint should require auth")
	}
	if adminMatcher(context.Background(), "/unknown.Service/Method") {
		t.Fatalf("unknown operation should not match admin auth selector")
	}

	appMatcher := whiteListMatcher(AppPrefixPathToWhiteList)
	if appMatcher(context.Background(), apppb.OperationUserLogin) {
		t.Fatalf("app login should be skipped by auth matcher")
	}
	if !appMatcher(context.Background(), apppb.OperationUserDeleteAccount) {
		t.Fatalf("app delete account endpoint should require auth")
	}
}
