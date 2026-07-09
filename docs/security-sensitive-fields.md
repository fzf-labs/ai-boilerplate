# Sensitive Field Handling

Backend third-party credentials must never be stored, returned, or logged in
plaintext.

## Encryption Key

- Primary environment variable: `AI_BOILERPLATE_CREDENTIAL_KEY`.
- Compatibility aliases: `THIRD_PARTY_CREDENTIAL_KEY`,
  `THIRD_PARTY_SECRET_KEY`.
- The value must be at least 32 characters, or base64-encoded 32 bytes. In
  production this value should come from the deployment secret store or KMS
  material injection.
- Stored ciphertext uses the `enc:v1:` prefix. Legacy plaintext values remain
  readable, but new create/update writes encrypt sensitive values.

## Sensitive Fields

The backend treats these names as sensitive for encryption or log redaction:

| Area | Fields |
| --- | --- |
| AI provider | `ai_provider_platform.api_key`, `APIKey`, `apiKey` |
| SMS channel | `sms_channel.api_key`, `sms_channel.api_secret`, `APIKey`, `APISecret` |
| Mail account | `mail_account.password`, `password` |
| WeChat official account | `wx_gzh_account.app_secret`, `token`, `encoding_aes_key`, `appSecret`, `encodingAesKey` |
| OSS/file storage config | `accessKey`, `secretKey`, `accessSecret`, `secretAccessKey`, `tmpSecretKey` |
| Auth/log payloads | `Authorization`, `token`, `refreshToken`, `sessionToken`, `cookie`, `setCookie` |

## API Rules

- Query APIs must return masked values by default. They must not return full
  ciphertext or plaintext credentials.
- Update APIs may receive a masked placeholder such as `******` or
  `abcd******wxyz`; this means "keep the existing secret".
- Resetting a credential requires sending a new value. The response still only
  returns the masked value.

## Development Rules

- Use `internal/security.EncryptSecret` or the existing data repo wrappers for
  new credential fields.
- Decrypt only at the boundary where a third-party SDK or SMTP/OSS client needs
  the plaintext value.
- Add new field names to `internal/security.IsSensitiveField` before logging or
  storing new credential shapes.
- Do not write request, response, error, or operation logs with raw credentials.
  Use `internal/security.RedactText` or `RedactJSONBytes` when adding manual
  logging paths.
