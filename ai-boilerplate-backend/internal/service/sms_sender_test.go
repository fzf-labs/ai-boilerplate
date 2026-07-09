package service

import (
	"context"
	"strings"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
	"gorm.io/datatypes"
)

func TestParseSMSTemplateParamKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  datatypes.JSON
		want map[string]struct{}
	}{
		{
			name: "map",
			raw:  datatypes.JSON([]byte(`{"code":"","minutes":""}`)),
			want: map[string]struct{}{"code": {}, "minutes": {}},
		},
		{
			name: "list",
			raw:  datatypes.JSON([]byte(`["code","code","ttl",""]`)),
			want: map[string]struct{}{"code": {}, "ttl": {}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := parseSMSTemplateParamKeys(tt.raw)
			if err != nil {
				t.Fatalf("parseSMSTemplateParamKeys() error = %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("keys = %#v, want %d keys", got, len(tt.want))
			}
			for _, key := range got {
				if _, ok := tt.want[key]; !ok {
					t.Fatalf("unexpected key %q in %#v", key, got)
				}
			}
		})
	}
}

func TestSafeSMSParamsContentMasksCode(t *testing.T) {
	t.Parallel()

	got := safeSMSParamsContent(map[string]string{
		"code": "123456",
		"name": "fixture",
	})
	if strings.Contains(got, "123456") {
		t.Fatalf("safeSMSParamsContent leaked code: %s", got)
	}
	if !strings.Contains(got, smsMask) || !strings.Contains(got, "fixture") {
		t.Fatalf("safeSMSParamsContent = %s, want masked code and normal params", got)
	}
}

func TestSanitizeSMSLogMessageMasksSecrets(t *testing.T) {
	t.Parallel()

	channel := testfixture.SmsChannel()
	got := sanitizeSMSLogMessage(
		"failed with ak_fixture_1234567890 sk_fixture_0987654321 and code 123456",
		channel,
		map[string]string{"code": "123456"},
	)
	if strings.Contains(got, channel.APIKey) || strings.Contains(got, channel.APISecret) || strings.Contains(got, "123456") {
		t.Fatalf("sanitizeSMSLogMessage leaked secret: %s", got)
	}
}

func TestMockSMSProviderFailure(t *testing.T) {
	t.Parallel()

	result, err := mockSMSProvider{}.Send(context.Background(), smsProviderRequest{
		APIKey: "fail",
		Params: map[string]string{
			"code": "123456",
		},
	})
	if err == nil {
		t.Fatalf("expected mock provider failure")
	}
	if result == nil || result.SendCode != "MOCK_FAILED" {
		t.Fatalf("result = %#v, want MOCK_FAILED", result)
	}
	if strings.Contains(err.Error(), "123456") {
		t.Fatalf("mock provider error leaked code: %v", err)
	}
}

func TestRenderSMSTemplateContent(t *testing.T) {
	t.Parallel()

	got := renderSMSTemplateContent("验证码{code}，{minutes}分钟有效", map[string]string{
		"code":    "123456",
		"minutes": "5",
	})
	if got != "验证码123456，5分钟有效" {
		t.Fatalf("renderSMSTemplateContent() = %q", got)
	}
}
