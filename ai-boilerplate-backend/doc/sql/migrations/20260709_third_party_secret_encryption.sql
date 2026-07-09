-- Expand third-party credential columns before enabling application-side
-- AES-GCM encryption. Existing plaintext values remain readable and are
-- encrypted on the next create/update through the backend.

ALTER TABLE public.ai_provider_platform
  ALTER COLUMN api_key TYPE character varying(1024);

ALTER TABLE public.sms_channel
  ALTER COLUMN api_key TYPE character varying(1024),
  ALTER COLUMN api_secret TYPE character varying(1024);

ALTER TABLE public.mail_account
  ALTER COLUMN password TYPE character varying(1024);

ALTER TABLE public.wx_gzh_account
  ALTER COLUMN app_secret TYPE character varying(1024),
  ALTER COLUMN token TYPE character varying(1024),
  ALTER COLUMN encoding_aes_key TYPE character varying(1024);
