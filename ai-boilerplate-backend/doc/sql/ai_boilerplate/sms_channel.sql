CREATE TABLE public.sms_channel (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    operator character varying(64) NOT NULL,
    remark character varying(255),
    api_key character varying(128) NOT NULL,
    api_secret character varying(128),
    callback_url character varying(255),
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sms_channel IS '短信渠道';
COMMENT ON COLUMN public.sms_channel.id IS 'id';
COMMENT ON COLUMN public.sms_channel.name IS '渠道名称';
COMMENT ON COLUMN public.sms_channel.operator IS '运营商';
COMMENT ON COLUMN public.sms_channel.remark IS '备注';
COMMENT ON COLUMN public.sms_channel.api_key IS '短信 API 的账号';
COMMENT ON COLUMN public.sms_channel.api_secret IS '短信 API 的秘钥';
COMMENT ON COLUMN public.sms_channel.callback_url IS '短信发送回调 URL';
COMMENT ON COLUMN public.sms_channel.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sms_channel.created_at IS '创建时间';
COMMENT ON COLUMN public.sms_channel.updated_at IS '更新时间';
COMMENT ON COLUMN public.sms_channel.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sms_channel ADD CONSTRAINT sms_channel_pkey PRIMARY KEY (id);
