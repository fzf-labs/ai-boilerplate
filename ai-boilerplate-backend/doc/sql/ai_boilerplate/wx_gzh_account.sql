CREATE TABLE public.wx_gzh_account (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    name character varying(100) NOT NULL,
    account character varying(100) NOT NULL,
    app_id character varying(100) NOT NULL,
    app_secret character varying(100) NOT NULL,
    url character varying(1000),
    token character varying(64),
    encoding_aes_key character varying(64),
    qr_code_url character varying(1000),
    remark character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_account IS '公众号账号表';
COMMENT ON COLUMN public.wx_gzh_account.id IS '编号';
COMMENT ON COLUMN public.wx_gzh_account.tenant_id IS '租户编号';
COMMENT ON COLUMN public.wx_gzh_account.name IS '公众号名称';
COMMENT ON COLUMN public.wx_gzh_account.account IS '公众号账号';
COMMENT ON COLUMN public.wx_gzh_account.app_id IS '公众号appid';
COMMENT ON COLUMN public.wx_gzh_account.app_secret IS '公众号密钥';
COMMENT ON COLUMN public.wx_gzh_account.url IS '公众号url';
COMMENT ON COLUMN public.wx_gzh_account.token IS '公众号token';
COMMENT ON COLUMN public.wx_gzh_account.encoding_aes_key IS '加密密钥';
COMMENT ON COLUMN public.wx_gzh_account.qr_code_url IS '二维码图片URL';
COMMENT ON COLUMN public.wx_gzh_account.remark IS '备注';
COMMENT ON COLUMN public.wx_gzh_account.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_account.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_account.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_account ADD CONSTRAINT wx_gzh_account_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_gzh_account_app_id_idx ON public.wx_gzh_account USING btree (app_id);
