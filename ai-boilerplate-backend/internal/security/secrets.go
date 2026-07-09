package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	CredentialCipherPrefix = "enc:v1:"
	SensitiveMask          = "******"
)

var (
	credentialKeyEnvNames = []string{
		"AI_BOILERPLATE_CREDENTIAL_KEY",
		"THIRD_PARTY_CREDENTIAL_KEY",
		"THIRD_PARTY_SECRET_KEY",
	}
	querySecretPattern = regexp.MustCompile(`(?i)(password|api[_-]?key|api[_-]?secret|secret[_-]?key|access[_-]?key|access[_-]?secret|secret[_-]?access[_-]?key|app[_-]?secret|encoding[_-]?aes[_-]?key|token|authorization)=([^&\s]+)`)
	pairSecretPattern  = regexp.MustCompile(`(?i)(["']?(?:password|oldPassword|newPassword|confirmPassword|apiKey|api_key|apiSecret|api_secret|secretKey|secret_key|accessKey|access_key|accessSecret|access_secret|secretAccessKey|secret_access_key|appSecret|app_secret|encodingAesKey|encoding_aes_key|token|authorization)["']?\s*[:=]\s*["']?)([^"',}\]\s]+)(["']?)`)
)

type jsonTransformMode int

const (
	jsonTransformEncrypt jsonTransformMode = iota
	jsonTransformEncryptUpdate
	jsonTransformDecrypt
	jsonTransformRedact
)

func CredentialKeyEnvNames() []string {
	names := make([]string, len(credentialKeyEnvNames))
	copy(names, credentialKeyEnvNames)
	return names
}

func IsEncryptedSecret(value string) bool {
	return strings.HasPrefix(value, CredentialCipherPrefix)
}

func IsMaskedSecretPlaceholder(value string) bool {
	return strings.Contains(strings.TrimSpace(value), SensitiveMask)
}

func EncryptSecret(value string) (string, error) {
	if value == "" || IsEncryptedSecret(value) {
		return value, nil
	}
	if IsMaskedSecretPlaceholder(value) {
		return "", errors.New("masked credential cannot be encrypted without an existing value")
	}
	key, err := credentialKey()
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nil, nonce, []byte(value), nil)
	payload := append(nonce, ciphertext...)
	return CredentialCipherPrefix + base64.RawURLEncoding.EncodeToString(payload), nil
}

func DecryptSecret(value string) (string, error) {
	if value == "" || !IsEncryptedSecret(value) {
		return value, nil
	}
	key, err := credentialKey()
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	payload, err := base64.RawURLEncoding.DecodeString(strings.TrimPrefix(value, CredentialCipherPrefix))
	if err != nil {
		return "", err
	}
	if len(payload) < gcm.NonceSize() {
		return "", errors.New("credential ciphertext is too short")
	}
	nonce := payload[:gcm.NonceSize()]
	ciphertext := payload[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func PrepareSecretForUpdate(newValue string, oldValue string) (string, error) {
	if strings.TrimSpace(newValue) == "" || IsMaskedSecretPlaceholder(newValue) {
		return oldValue, nil
	}
	return EncryptSecret(newValue)
}

func MaskSecret(value string) string {
	if value == "" {
		return ""
	}
	if IsEncryptedSecret(value) {
		plaintext, err := DecryptSecret(value)
		if err != nil {
			return SensitiveMask
		}
		value = plaintext
	}
	runes := []rune(value)
	if len(runes) <= 8 {
		return SensitiveMask
	}
	return string(runes[:4]) + SensitiveMask + string(runes[len(runes)-4:])
}

func IsSensitiveField(name string) bool {
	switch canonicalFieldName(name) {
	case "password",
		"oldpassword",
		"newpassword",
		"confirmpassword",
		"apikey",
		"apisecret",
		"secret",
		"secretkey",
		"accesskey",
		"accesssecret",
		"accesskeysecret",
		"secretaccesskey",
		"tmpsecretkey",
		"tmpsecretid",
		"appsecret",
		"encodingaeskey",
		"aeskey",
		"token",
		"refreshtoken",
		"sessiontoken",
		"authorization",
		"cookie",
		"setcookie",
		"credential",
		"credentials",
		"privatekey",
		"clientsecret":
		return true
	default:
		return false
	}
}

func EncryptJSONSecrets(raw []byte) ([]byte, error) {
	return transformJSONBytes(raw, nil, jsonTransformEncrypt)
}

func PrepareJSONSecretsForUpdate(newRaw []byte, oldRaw []byte) ([]byte, error) {
	return transformJSONBytes(newRaw, oldRaw, jsonTransformEncryptUpdate)
}

func DecryptJSONSecrets(raw []byte) ([]byte, error) {
	return transformJSONBytes(raw, nil, jsonTransformDecrypt)
}

func RedactJSONBytes(raw []byte) []byte {
	out, err := transformJSONBytes(raw, nil, jsonTransformRedact)
	if err != nil {
		return []byte(RedactText(string(raw)))
	}
	return out
}

func RedactText(text string) string {
	if strings.TrimSpace(text) == "" {
		return text
	}
	raw := []byte(text)
	if json.Valid(raw) {
		return string(RedactJSONBytes(raw))
	}
	text = querySecretPattern.ReplaceAllStringFunc(text, func(match string) string {
		idx := strings.Index(match, "=")
		if idx < 0 {
			return match
		}
		return match[:idx+1] + SensitiveMask
	})
	text = pairSecretPattern.ReplaceAllString(text, "${1}"+SensitiveMask+"${3}")
	return text
}

func RedactValue(key string, value any) any {
	if IsSensitiveField(key) {
		return SensitiveMask
	}
	switch v := value.(type) {
	case string:
		return RedactText(v)
	case []byte:
		return []byte(RedactText(string(v)))
	case error:
		return RedactText(v.Error())
	default:
		return value
	}
}

func transformJSONBytes(raw []byte, oldRaw []byte, mode jsonTransformMode) ([]byte, error) {
	if len(bytes.TrimSpace(raw)) == 0 {
		return raw, nil
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return nil, err
	}
	var oldValue any
	if len(bytes.TrimSpace(oldRaw)) > 0 {
		_ = json.Unmarshal(oldRaw, &oldValue)
	}
	transformed, err := transformJSONValue("", value, oldValue, mode)
	if err != nil {
		return nil, err
	}
	return json.Marshal(transformed)
}

func transformJSONValue(key string, value any, oldValue any, mode jsonTransformMode) (any, error) {
	if IsSensitiveField(key) {
		return transformSensitiveValue(value, oldValue, mode)
	}
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		oldMap, _ := oldValue.(map[string]any)
		for itemKey, itemValue := range typed {
			var oldItem any
			if oldMap != nil {
				oldItem = oldMap[itemKey]
			}
			next, err := transformJSONValue(itemKey, itemValue, oldItem, mode)
			if err != nil {
				return nil, err
			}
			out[itemKey] = next
		}
		return out, nil
	case []any:
		out := make([]any, len(typed))
		oldSlice, _ := oldValue.([]any)
		for i, itemValue := range typed {
			var oldItem any
			if i < len(oldSlice) {
				oldItem = oldSlice[i]
			}
			next, err := transformJSONValue(key, itemValue, oldItem, mode)
			if err != nil {
				return nil, err
			}
			out[i] = next
		}
		return out, nil
	default:
		return value, nil
	}
}

func transformSensitiveValue(value any, oldValue any, mode jsonTransformMode) (any, error) {
	switch mode {
	case jsonTransformRedact:
		switch typed := value.(type) {
		case string:
			return MaskSecret(typed), nil
		case nil:
			return nil, nil
		default:
			return SensitiveMask, nil
		}
	case jsonTransformDecrypt:
		typed, ok := value.(string)
		if !ok {
			return value, nil
		}
		return DecryptSecret(typed)
	case jsonTransformEncrypt, jsonTransformEncryptUpdate:
		typed, ok := value.(string)
		if !ok {
			return value, nil
		}
		if mode == jsonTransformEncryptUpdate {
			if strings.TrimSpace(typed) == "" || IsMaskedSecretPlaceholder(typed) {
				if oldString, ok := oldValue.(string); ok {
					return oldString, nil
				}
				return typed, nil
			}
		}
		return EncryptSecret(typed)
	default:
		return value, nil
	}
}

func credentialKey() ([]byte, error) {
	var configured string
	for _, name := range credentialKeyEnvNames {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			configured = value
			break
		}
	}
	if configured == "" {
		return nil, fmt.Errorf("credential encryption key is not configured; set %s", strings.Join(credentialKeyEnvNames, " or "))
	}
	if decoded, err := base64.StdEncoding.DecodeString(configured); err == nil && len(decoded) == 32 {
		return decoded, nil
	}
	if decoded, err := base64.RawStdEncoding.DecodeString(configured); err == nil && len(decoded) == 32 {
		return decoded, nil
	}
	if len(configured) < 32 {
		return nil, errors.New("credential encryption key must be at least 32 characters or base64-encoded 32 bytes")
	}
	sum := sha256.Sum256([]byte(configured))
	return sum[:], nil
}

func canonicalFieldName(name string) string {
	var b strings.Builder
	b.Grow(len(name))
	for _, r := range name {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}
