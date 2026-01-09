CREATE TABLE public.ai_prompt (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64),
    name character varying(128) NOT NULL,
    "desc" character varying(255) NOT NULL,
    prompt text NOT NULL,
    sort integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_prompt IS 'AI 提示词';
COMMENT ON COLUMN public.ai_prompt.id IS 'Id';
COMMENT ON COLUMN public.ai_prompt.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_prompt.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_prompt.name IS '名称';
COMMENT ON COLUMN public.ai_prompt."desc" IS '描述';
COMMENT ON COLUMN public.ai_prompt.prompt IS '提示词';
COMMENT ON COLUMN public.ai_prompt.sort IS '排序';
COMMENT ON COLUMN public.ai_prompt.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_prompt.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_prompt.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_prompt ADD CONSTRAINT ai_prompt_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_prompt_tenant_id ON public.ai_prompt USING btree (tenant_id);
CREATE INDEX idx_ai_prompt_user_id ON public.ai_prompt USING btree (admin_id);
