CREATE TABLE public.ai_provider_model (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    platform_id character varying(64) NOT NULL,
    model_type character varying(64) NOT NULL,
    model_id character varying(64) NOT NULL,
    model_name character varying(64),
    model_config jsonb,
    sort integer NOT NULL,
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_provider_model IS 'AI 配置模型表';
COMMENT ON COLUMN public.ai_provider_model.id IS '编号';
COMMENT ON COLUMN public.ai_provider_model.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_model.platform_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_model.model_type IS '模型类型';
COMMENT ON COLUMN public.ai_provider_model.model_id IS '模型id';
COMMENT ON COLUMN public.ai_provider_model.model_name IS '模型名字';
COMMENT ON COLUMN public.ai_provider_model.model_config IS '配置';
COMMENT ON COLUMN public.ai_provider_model.sort IS '排序';
COMMENT ON COLUMN public.ai_provider_model.status IS '状态';
COMMENT ON COLUMN public.ai_provider_model.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_provider_model.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_provider_model.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_provider_model ADD CONSTRAINT ai_provider_model_pkey PRIMARY KEY (id);
