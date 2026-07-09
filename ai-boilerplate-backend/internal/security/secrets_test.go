package security

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

const testCredentialKey = "0123456789abcdef0123456789abcdef"

func TestEncryptSecretRoundTrip(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", testCredentialKey)

	const plaintext = "ak_live_1234567890"
	ciphertext, err := EncryptSecret(plaintext)
	if err != nil {
		t.Fatalf("EncryptSecret() error = %v", err)
	}
	if !IsEncryptedSecret(ciphertext) {
		t.Fatalf("ciphertext should use prefix %q, got %q", CredentialCipherPrefix, ciphertext)
	}
	if strings.Contains(ciphertext, plaintext) {
		t.Fatalf("ciphertext leaked plaintext: %q", ciphertext)
	}
	decrypted, err := DecryptSecret(ciphertext)
	if err != nil {
		t.Fatalf("DecryptSecret() error = %v", err)
	}
	if decrypted != plaintext {
		t.Fatalf("DecryptSecret() = %q, want %q", decrypted, plaintext)
	}
}

func TestPrepareSecretForUpdateKeepsMaskedValue(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", testCredentialKey)

	oldSecret, err := EncryptSecret("old-secret-value")
	if err != nil {
		t.Fatalf("EncryptSecret() error = %v", err)
	}
	got, err := PrepareSecretForUpdate("old-******alue", oldSecret)
	if err != nil {
		t.Fatalf("PrepareSecretForUpdate() error = %v", err)
	}
	if got != oldSecret {
		t.Fatalf("masked update should keep old ciphertext")
	}
}

func TestJSONSecretTransforms(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", testCredentialKey)

	raw := []byte(`{"aliyun":{"accessKey":"ak_live_1234","secretKey":"sk_live_5678","bucket":"demo"}}`)
	encrypted, err := EncryptJSONSecrets(raw)
	if err != nil {
		t.Fatalf("EncryptJSONSecrets() error = %v", err)
	}
	if strings.Contains(string(encrypted), "ak_live_1234") || strings.Contains(string(encrypted), "sk_live_5678") {
		t.Fatalf("encrypted JSON leaked plaintext: %s", encrypted)
	}

	decrypted, err := DecryptJSONSecrets(encrypted)
	if err != nil {
		t.Fatalf("DecryptJSONSecrets() error = %v", err)
	}
	var got map[string]map[string]string
	if err := json.Unmarshal(decrypted, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if got["aliyun"]["accessKey"] != "ak_live_1234" || got["aliyun"]["secretKey"] != "sk_live_5678" {
		t.Fatalf("decrypted JSON mismatch: %#v", got)
	}

	redacted := string(RedactJSONBytes(encrypted))
	if strings.Contains(redacted, "ak_live_1234") || strings.Contains(redacted, "sk_live_5678") {
		t.Fatalf("redacted JSON leaked plaintext: %s", redacted)
	}
	if !strings.Contains(redacted, SensitiveMask) {
		t.Fatalf("redacted JSON should contain mask, got %s", redacted)
	}
}

func TestPrepareJSONSecretsForUpdateKeepsMaskedValue(t *testing.T) {
	t.Setenv("AI_BOILERPLATE_CREDENTIAL_KEY", testCredentialKey)

	oldRaw, err := EncryptJSONSecrets([]byte(`{"qiniu":{"accessKey":"old-ak","secretKey":"old-sk","bucket":"demo"}}`))
	if err != nil {
		t.Fatalf("EncryptJSONSecrets() error = %v", err)
	}
	newRaw := []byte(`{"qiniu":{"accessKey":"******","secretKey":"old-******ld","bucket":"demo-2"}}`)
	prepared, err := PrepareJSONSecretsForUpdate(newRaw, oldRaw)
	if err != nil {
		t.Fatalf("PrepareJSONSecretsForUpdate() error = %v", err)
	}
	decrypted, err := DecryptJSONSecrets(prepared)
	if err != nil {
		t.Fatalf("DecryptJSONSecrets() error = %v", err)
	}
	if !strings.Contains(string(decrypted), `"accessKey":"old-ak"`) || !strings.Contains(string(decrypted), `"secretKey":"old-sk"`) {
		t.Fatalf("masked update did not preserve old secrets: %s", decrypted)
	}
	if !strings.Contains(string(decrypted), `"bucket":"demo-2"`) {
		t.Fatalf("non-secret field was not updated: %s", decrypted)
	}
}

func TestRedactTextAndLogger(t *testing.T) {
	raw := `{"password":"secret-value","nested":{"apiKey":"ak_live_123456"}}`
	redacted := RedactText(raw)
	if strings.Contains(redacted, "secret-value") || strings.Contains(redacted, "ak_live_123456") {
		t.Fatalf("RedactText leaked secret: %s", redacted)
	}

	capture := &captureLogger{}
	logger := NewRedactingLogger(capture)
	if err := logger.Log(log.LevelInfo, "apiKey", "ak_live_123456", "msg", "token=secret-token"); err != nil {
		t.Fatalf("Log() error = %v", err)
	}
	got := strings.TrimSpace(strings.Join(capture.values, " "))
	if strings.Contains(got, "ak_live_123456") || strings.Contains(got, "secret-token") {
		t.Fatalf("redacting logger leaked secret: %s", got)
	}
}

type captureLogger struct {
	values []string
}

func (l *captureLogger) Log(_ log.Level, keyvals ...any) error {
	for _, item := range keyvals {
		l.values = append(l.values, fmt.Sprint(item))
	}
	return nil
}
