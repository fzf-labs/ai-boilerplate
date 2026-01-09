CREATE TABLE public.ai_write_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    type integer,
    platform character varying(255) NOT NULL,
    model_id character varying(64) NOT NULL,
    model character varying(255) NOT NULL,
    prompt text NOT NULL,
    generated_content text,
    original_content text,
    length integer,
    format integer,
    tone integer,
    language integer,
    error_message character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_write_record IS 'AI 写作表';
COMMENT ON COLUMN public.ai_write_record.id IS '编号';
COMMENT ON COLUMN public.ai_write_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_write_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_write_record.type IS '写作类型';
COMMENT ON COLUMN public.ai_write_record.platform IS '平台';
COMMENT ON COLUMN public.ai_write_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_write_record.model IS '模型';
COMMENT ON COLUMN public.ai_write_record.prompt IS '生成内容提示';
COMMENT ON COLUMN public.ai_write_record.generated_content IS '生成的内容';
COMMENT ON COLUMN public.ai_write_record.original_content IS '原文';
COMMENT ON COLUMN public.ai_write_record.length IS '长度提示词';
COMMENT ON COLUMN public.ai_write_record.format IS '格式提示词';
COMMENT ON COLUMN public.ai_write_record.tone IS '语气提示词';
COMMENT ON COLUMN public.ai_write_record.language IS '语言提示词';
COMMENT ON COLUMN public.ai_write_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_write_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_write_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_write_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_write_record ADD CONSTRAINT ai_write_record_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_write_record_platform ON public.ai_write_record USING btree (platform);
CREATE INDEX idx_ai_write_record_tenant_id ON public.ai_write_record USING btree (tenant_id);
CREATE INDEX idx_ai_write_record_user_id ON public.ai_write_record USING btree (admin_id);
