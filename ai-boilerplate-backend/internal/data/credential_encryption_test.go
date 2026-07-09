package data

import (
	"strings"
	"testing"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/ai-boilerplate-backend/internal/security"
	"gorm.io/datatypes"
)

func TestModelCredentialEncryption(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", "0123456789abcdef0123456789abcdef")

	aiProvider := &ai_boilerplate_model.AiProviderPlatform{APIKey: "ai-provider-api-key"}
	if err := encryptAiProviderPlatformSecrets(aiProvider); err != nil {
		t.Fatalf("encryptAiProviderPlatformSecrets() error = %v", err)
	}
	assertEncryptedSecret(t, aiProvider.APIKey, "ai-provider-api-key")

	smsChannel := &ai_boilerplate_model.SmsChannel{APIKey: "sms-api-key", APISecret: "sms-api-secret"}
	if err := encryptSmsChannelSecrets(smsChannel); err != nil {
		t.Fatalf("encryptSmsChannelSecrets() error = %v", err)
	}
	assertEncryptedSecret(t, smsChannel.APIKey, "sms-api-key")
	assertEncryptedSecret(t, smsChannel.APISecret, "sms-api-secret")

	mailAccount := &ai_boilerplate_model.MailAccount{Password: "smtp-password"}
	if err := encryptMailAccountSecrets(mailAccount); err != nil {
		t.Fatalf("encryptMailAccountSecrets() error = %v", err)
	}
	assertEncryptedSecret(t, mailAccount.Password, "smtp-password")

	wxAccount := &ai_boilerplate_model.WxGzhAccount{
		AppSecret:      "wx-app-secret",
		Token:          "wx-token",
		EncodingAesKey: "wx-encoding-aes-key",
	}
	if err := encryptWxGzhAccountSecrets(wxAccount); err != nil {
		t.Fatalf("encryptWxGzhAccountSecrets() error = %v", err)
	}
	assertEncryptedSecret(t, wxAccount.AppSecret, "wx-app-secret")
	assertEncryptedSecret(t, wxAccount.Token, "wx-token")
	assertEncryptedSecret(t, wxAccount.EncodingAesKey, "wx-encoding-aes-key")
}

func TestPrepareModelCredentialUpdateKeepsMaskedValue(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", "0123456789abcdef0123456789abcdef")

	oldPassword, err := security.EncryptSecret("old-password")
	if err != nil {
		t.Fatalf("EncryptSecret() error = %v", err)
	}
	newData := &ai_boilerplate_model.MailAccount{Password: "old-******word"}
	oldData := &ai_boilerplate_model.MailAccount{Password: oldPassword}
	if err := prepareMailAccountSecretsForUpdate(newData, oldData); err != nil {
		t.Fatalf("prepareMailAccountSecretsForUpdate() error = %v", err)
	}
	if newData.Password != oldPassword {
		t.Fatalf("masked password update should keep old ciphertext")
	}
}

func TestFileConfigCredentialJSONEncryption(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", "0123456789abcdef0123456789abcdef")

	fileConfig := &ai_boilerplate_model.FileConfig{
		Config: datatypes.JSON([]byte(`{"aliyun":{"accessKey":"oss-access-key","secretKey":"oss-secret-key","bucket":"demo"}}`)),
	}
	if err := encryptFileConfigSecrets(fileConfig); err != nil {
		t.Fatalf("encryptFileConfigSecrets() error = %v", err)
	}
	if strings.Contains(fileConfig.Config.String(), "oss-access-key") || strings.Contains(fileConfig.Config.String(), "oss-secret-key") {
		t.Fatalf("encrypted file config leaked plaintext: %s", fileConfig.Config.String())
	}

	decrypted, err := security.DecryptJSONSecrets(fileConfig.Config)
	if err != nil {
		t.Fatalf("DecryptJSONSecrets() error = %v", err)
	}
	if !strings.Contains(string(decrypted), `"accessKey":"oss-access-key"`) || !strings.Contains(string(decrypted), `"secretKey":"oss-secret-key"`) {
		t.Fatalf("decrypted file config mismatch: %s", decrypted)
	}
}

func assertEncryptedSecret(t *testing.T, ciphertext string, plaintext string) {
	t.Helper()

	if !security.IsEncryptedSecret(ciphertext) {
		t.Fatalf("secret should be encrypted, got %q", ciphertext)
	}
	if strings.Contains(ciphertext, plaintext) {
		t.Fatalf("ciphertext leaked plaintext %q: %q", plaintext, ciphertext)
	}
	decrypted, err := security.DecryptSecret(ciphertext)
	if err != nil {
		t.Fatalf("DecryptSecret() error = %v", err)
	}
	if decrypted != plaintext {
		t.Fatalf("DecryptSecret() = %q, want %q", decrypted, plaintext)
	}
}
