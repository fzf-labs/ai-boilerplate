package data

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	appjwt "github.com/fzf-labs/gopkg/jwt"
)

func TestUserRepoPasswordAndToken(t *testing.T) {
	repo := &UserRepo{jwt: testJWT("parent")}

	hash, err := repo.GeneratePassword("fixture-salt:", "fixture-password")
	if err != nil {
		t.Fatalf("GeneratePassword() error = %v", err)
	}
	if hash == "" || hash == "fixture-salt:fixture-password" {
		t.Fatalf("GeneratePassword() returned unsafe hash %q", hash)
	}
	if !repo.VerifyPassword("fixture-salt:", "fixture-password", hash) {
		t.Fatal("VerifyPassword() should accept matching fixture password")
	}
	if repo.VerifyPassword("fixture-salt:", "wrong-password", hash) {
		t.Fatal("VerifyPassword() should reject mismatched fixture password")
	}

	token, err := repo.GenerateToken(context.Background(), "user-fixture", "wx-gzh-fixture", "wx-xcx-fixture")
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}
	assertTokenWindow(t, token.ExpiredAt, token.RefreshAt)

	claims, err := repo.CheckToken(context.Background(), token.Token)
	if err != nil {
		t.Fatalf("CheckToken() error = %v", err)
	}
	if got := claims["uid"]; got != "user-fixture" {
		t.Fatalf("uid claim = %v, want user-fixture", got)
	}
	if got := claims["wxGzhUserId"]; got != "wx-gzh-fixture" {
		t.Fatalf("wxGzhUserId claim = %v, want wx-gzh-fixture", got)
	}
	if got := claims["wxGzhXcxId"]; got != "wx-xcx-fixture" {
		t.Fatalf("wxGzhXcxId claim = %v, want wx-xcx-fixture", got)
	}
}

func TestSysAdminRepoTokenCarriesAdminClaims(t *testing.T) {
	repo := &SysAdminRepo{jwt: testJWT("admin")}
	admin := &ai_boilerplate_model.SysAdmin{
		ID:       "admin-fixture",
		Nickname: "Fixture Admin",
		TenantID: "tenant-fixture",
	}

	token, err := repo.GenerateToken(context.Background(), admin)
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}
	assertTokenWindow(t, token.ExpiredAt, token.RefreshAt)

	claims, err := repo.CheckToken(context.Background(), token.Token)
	if err != nil {
		t.Fatalf("CheckToken() error = %v", err)
	}
	if got := claims["uid"]; got != admin.ID {
		t.Fatalf("uid claim = %v, want %s", got, admin.ID)
	}
	if got := claims["nickname"]; got != admin.Nickname {
		t.Fatalf("nickname claim = %v, want %s", got, admin.Nickname)
	}
	if got := claims["tenant_id"]; got != admin.TenantID {
		t.Fatalf("tenant_id claim = %v, want %s", got, admin.TenantID)
	}
}

func TestSmsCodeRepoGeneratesNumericFixtureCode(t *testing.T) {
	repo := &SmsCodeRepo{}

	cfg := repo.GetSmsConfig(SmsCodeSceneReset)
	if cfg.CodeLength != 6 {
		t.Fatalf("reset code length = %d, want 6", cfg.CodeLength)
	}
	if cfg.CodeTTL != 10*time.Minute {
		t.Fatalf("reset code ttl = %s, want 10m", cfg.CodeTTL)
	}

	code, err := repo.GenerateSmsCode(SmsCodeSceneLogin)
	if err != nil {
		t.Fatalf("GenerateSmsCode() error = %v", err)
	}
	if !regexp.MustCompile(`^\d{6}$`).MatchString(code) {
		t.Fatalf("GenerateSmsCode() = %q, want six digits", code)
	}

	data, err := repo.GenerateSmsCodeData(SmsCodeSceneBind, "13800138000", "user-fixture")
	if err != nil {
		t.Fatalf("GenerateSmsCodeData() error = %v", err)
	}
	if data.Scene != SmsCodeSceneBind || data.Phone != "13800138000" || data.UID != "user-fixture" {
		t.Fatalf("GenerateSmsCodeData() = %+v, want bind fixture data", data)
	}
	if data.CodeID == "" || !regexp.MustCompile(`^\d{6}$`).MatchString(data.Code) {
		t.Fatalf("GenerateSmsCodeData() produced invalid code payload %+v", data)
	}
}

func testJWT(issuer string) *appjwt.Jwt {
	return appjwt.NewJwt(&appjwt.Config{
		AccessSecret: "unit-test-secret-" + issuer,
		AccessExpire: 3600,
		RefreshAfter: 60,
		Issuer:       issuer,
	}, nil)
}

func assertTokenWindow(t *testing.T, expiredAt int64, refreshAt int64) {
	t.Helper()
	if expiredAt <= 0 || refreshAt <= 0 {
		t.Fatalf("token times must be set, expiredAt=%d refreshAt=%d", expiredAt, refreshAt)
	}
	if expiredAt <= refreshAt {
		t.Fatalf("expiredAt=%d must be after refreshAt=%d", expiredAt, refreshAt)
	}
}
