package service

import (
	"strings"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/testfixture"
	"gorm.io/datatypes"
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

func TestThirdPartyInfoFromModelMasksSecrets(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", "0123456789abcdef0123456789abcdef")

	encryptedAPIKey, err := security.EncryptSecret("ai-provider-api-key")
	if err != nil {
		t.Fatalf("EncryptSecret() error = %v", err)
	}
	aiInfo := aiProviderPlatformInfoFromModel(&ai_boilerplate_model.AiProviderPlatform{APIKey: encryptedAPIKey})
	if aiInfo.APIKey == "ai-provider-api-key" || strings.Contains(aiInfo.APIKey, "provider") {
		t.Fatalf("AI provider APIKey should be masked, got %q", aiInfo.APIKey)
	}

	mailInfo := mailAccountInfoFromModel(&ai_boilerplate_model.MailAccount{Password: "smtp-password"})
	if mailInfo.Password == "smtp-password" || !strings.Contains(mailInfo.Password, sensitiveMask) {
		t.Fatalf("mail password should be masked, got %q", mailInfo.Password)
	}

	wxInfo := wxGzhAccountInfoFromModel(&ai_boilerplate_model.WxGzhAccount{
		AppSecret:      "wx-app-secret",
		Token:          "wx-token",
		EncodingAesKey: "wx-encoding-aes-key",
	})
	if wxInfo.AppSecret == "wx-app-secret" || wxInfo.Token == "wx-token" || wxInfo.EncodingAesKey == "wx-encoding-aes-key" {
		t.Fatalf("wx secrets should be masked: %#v", wxInfo)
	}
}

func TestStorageConfigFromJSONMasksSecrets(t *testing.T) {
	configJSON := datatypes.JSON([]byte(`{"aliyun":{"accessKey":"oss-access-key","secretKey":"oss-secret-key","bucket":"demo"}}`))
	config, err := storageConfigFromJSON(configJSON, true)
	if err != nil {
		t.Fatalf("storageConfigFromJSON() error = %v", err)
	}
	if config.Aliyun.AccessKey == "oss-access-key" || config.Aliyun.SecretKey == "oss-secret-key" {
		t.Fatalf("storage config secrets should be masked: %#v", config.Aliyun)
	}
	if config.Aliyun.Bucket != "demo" {
		t.Fatalf("non-secret field should be preserved, got %q", config.Aliyun.Bucket)
	}
}
