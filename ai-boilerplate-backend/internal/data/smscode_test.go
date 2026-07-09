package data

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"
)

func TestSmsCodeConfigByScene(t *testing.T) {
	t.Parallel()

	repo := &SmsCodeRepo{}

	tests := []struct {
		name        string
		scene       SmsCodeScene
		hourlyLimit int
		dailyLimit  int
		codeTTL     time.Duration
	}{
		{name: "login", scene: SmsCodeSceneLogin, hourlyLimit: 5, dailyLimit: 15, codeTTL: 5 * time.Minute},
		{name: "register", scene: SmsCodeSceneRegister, hourlyLimit: 3, dailyLimit: 10, codeTTL: 5 * time.Minute},
		{name: "bind", scene: SmsCodeSceneBind, hourlyLimit: 3, dailyLimit: 5, codeTTL: 5 * time.Minute},
		{name: "reset", scene: SmsCodeSceneReset, hourlyLimit: 2, dailyLimit: 5, codeTTL: 10 * time.Minute},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := repo.GetSmsConfig(tt.scene)
			if got.Scene != tt.scene {
				t.Fatalf("Scene = %q, want %q", got.Scene, tt.scene)
			}
			if got.HourlyLimit != tt.hourlyLimit || got.DailyLimit != tt.dailyLimit {
				t.Fatalf("limits = hourly %d daily %d, want hourly %d daily %d", got.HourlyLimit, got.DailyLimit, tt.hourlyLimit, tt.dailyLimit)
			}
			if got.CodeTTL != tt.codeTTL {
				t.Fatalf("CodeTTL = %s, want %s", got.CodeTTL, tt.codeTTL)
			}
			if got.CodeLength != 6 {
				t.Fatalf("CodeLength = %d, want 6", got.CodeLength)
			}
			if got.MaxCheckAttempts != 5 {
				t.Fatalf("MaxCheckAttempts = %d, want 5", got.MaxCheckAttempts)
			}
		})
	}
}

func TestGenerateSmsCodeData(t *testing.T) {
	t.Parallel()

	repo := &SmsCodeRepo{}
	got, err := repo.GenerateSmsCodeData(SmsCodeSceneRegister, "13800138000", "fixture-user")
	if err != nil {
		t.Fatalf("GenerateSmsCodeData() error = %v", err)
	}
	if got.Scene != SmsCodeSceneRegister || got.Phone != "13800138000" || got.UID != "fixture-user" {
		t.Fatalf("unexpected sms code data: %#v", got)
	}
	if got.CodeID == "" {
		t.Fatalf("CodeID should be generated")
	}
	if !regexp.MustCompile(`^\d{6}$`).MatchString(got.Code) {
		t.Fatalf("Code = %q, want 6 digits", got.Code)
	}
}

func TestSendSmsCodeDelegatesCallback(t *testing.T) {
	t.Parallel()

	repo := &SmsCodeRepo{}
	codeData := &SmsCodeData{
		Scene:  SmsCodeSceneLogin,
		CodeID: "fixture-code-id",
		Code:   "123456",
		Phone:  "13800138000",
	}
	sentinel := errors.New("provider rejected message")
	called := false

	err := repo.SendSmsCode(context.Background(), codeData, func(ctx context.Context, data *SmsCodeData) error {
		if ctx == nil {
			t.Fatalf("context should be passed to callback")
		}
		if data != codeData {
			t.Fatalf("callback data = %#v, want original pointer", data)
		}
		called = true
		return sentinel
	})

	if !called {
		t.Fatalf("callback was not called")
	}
	if !errors.Is(err, sentinel) {
		t.Fatalf("SendSmsCode() error = %v, want %v", err, sentinel)
	}
}
