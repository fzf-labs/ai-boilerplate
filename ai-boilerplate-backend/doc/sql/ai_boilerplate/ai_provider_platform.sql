CREATE TABLE public.ai_provider_platform (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    platform character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    api_url character varying(500),
    api_key character varying(500),
    doc_url character varying(500),
    sort integer,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_provider_platform IS 'AI 配置平台表';
COMMENT ON COLUMN public.ai_provider_platform.id IS '编号';
COMMENT ON COLUMN public.ai_provider_platform.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_platform.platform IS '平台';
COMMENT ON COLUMN public.ai_provider_platform.name IS '名称';
COMMENT ON COLUMN public.ai_provider_platform.api_url IS 'API 地址';
COMMENT ON COLUMN public.ai_provider_platform.api_key IS 'API KEY';
COMMENT ON COLUMN public.ai_provider_platform.doc_url IS '文档地址';
COMMENT ON COLUMN public.ai_provider_platform.sort IS '排序';
COMMENT ON COLUMN public.ai_provider_platform.status IS '状态';
COMMENT ON COLUMN public.ai_provider_platform.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_provider_platform.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_provider_platform.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_provider_platform ADD CONSTRAINT ai_provider_platform_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX ai_conf_platform_pkey ON public.ai_provider_platform USING btree (id);
