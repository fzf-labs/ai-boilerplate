package service

import (
	"strings"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
)

func TestMaskSensitiveValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "empty", value: "", want: ""},
		{name: "short", value: "secret", want: "******"},
		{name: "long", value: "abcd1234wxyz", want: "abcd******wxyz"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := maskSensitiveValue(tt.value); got != tt.want {
				t.Fatalf("maskSensitiveValue(%q) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestSmsChannelInfoFromModelMasksSecrets(t *testing.T) {
	t.Parallel()

	channel := testfixture.SmsChannel()
	info := smsChannelInfoFromModel(channel)

	if info.APIKey == channel.APIKey || info.APISecret == channel.APISecret {
		t.Fatalf("sensitive fields were not masked: %#v", info)
	}
	if !strings.HasPrefix(info.APIKey, "ak_f") || !strings.HasSuffix(info.APIKey, "7890") {
		t.Fatalf("APIKey mask should preserve a small prefix and suffix, got %q", info.APIKey)
	}
	if strings.Contains(info.APISecret, "098765") {
		t.Fatalf("APISecret mask leaks middle secret: %q", info.APISecret)
	}
	if info.OperatorName == "" {
		t.Fatalf("operator name should be populated")
	}
}
