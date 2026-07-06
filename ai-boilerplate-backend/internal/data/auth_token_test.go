package data

import (
	"context"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
	"github.com/fzf-labs/gopkg/jwt"
)

func TestSysAdminTokenRoundTrip(t *testing.T) {
	t.Parallel()

	repo := &SysAdminRepo{jwt: newTestJWT()}
	admin := testfixture.SysAdmin()

	token, err := repo.GenerateToken(context.Background(), admin)
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}
	if token.Token == "" || token.ExpiredAt <= token.RefreshAt {
		t.Fatalf("unexpected token metadata: %#v", token)
	}

	claims, err := repo.CheckToken(context.Background(), token.Token)
	if err != nil {
		t.Fatalf("CheckToken() error = %v", err)
	}
	if claims["uid"] != admin.ID {
		t.Fatalf("uid claim = %#v, want %q", claims["uid"], admin.ID)
	}
	if claims["nickname"] != admin.Nickname {
		t.Fatalf("nickname claim = %#v, want %q", claims["nickname"], admin.Nickname)
	}
	if claims["tenant_id"] != admin.TenantID {
		t.Fatalf("tenant_id claim = %#v, want %q", claims["tenant_id"], admin.TenantID)
	}
	if claims[jwt.JwtID] == "" || claims[jwt.JwtExpired] == nil || claims[jwt.JwtRefresh] == nil {
		t.Fatalf("standard token claims missing: %#v", claims)
	}
}

func TestUserTokenRoundTrip(t *testing.T) {
	t.Parallel()

	repo := &UserRepo{jwt: newTestJWT()}

	token, err := repo.GenerateToken(context.Background(), testfixture.UserID, "wx-gzh-001", "wx-xcx-001")
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}

	claims, err := repo.CheckToken(context.Background(), token.Token)
	if err != nil {
		t.Fatalf("CheckToken() error = %v", err)
	}
	if claims["uid"] != testfixture.UserID {
		t.Fatalf("uid claim = %#v, want %q", claims["uid"], testfixture.UserID)
	}
	if claims["wxGzhUserId"] != "wx-gzh-001" {
		t.Fatalf("wxGzhUserId claim = %#v", claims["wxGzhUserId"])
	}
	if claims["wxGzhXcxId"] != "wx-xcx-001" {
		t.Fatalf("wxGzhXcxId claim = %#v", claims["wxGzhXcxId"])
	}
}

func newTestJWT() *jwt.Jwt {
	return jwt.NewJwt(&jwt.Config{
		AccessSecret: "unit-test-secret",
		AccessExpire: 3600,
		RefreshAfter: 60,
		Issuer:       "unit-test",
	}, nil)
}
