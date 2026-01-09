CREATE TABLE public.ai_image_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    prompt text NOT NULL,
    platform character varying(64) NOT NULL,
    model_id character varying(64),
    model character varying(64) NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    status integer NOT NULL,
    finish_time timestamp with time zone,
    error_message character varying(1024),
    public_status boolean DEFAULT false NOT NULL,
    pic_url character varying(2048),
    options jsonb,
    task_id character varying(255),
    buttons character varying(2048),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_image_record IS 'AI 绘画表';
COMMENT ON COLUMN public.ai_image_record.id IS '编号';
COMMENT ON COLUMN public.ai_image_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_image_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_image_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_image_record.platform IS '平台';
COMMENT ON COLUMN public.ai_image_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_image_record.model IS '模型';
COMMENT ON COLUMN public.ai_image_record.width IS '图片宽度';
COMMENT ON COLUMN public.ai_image_record.height IS '图片高度';
COMMENT ON COLUMN public.ai_image_record.status IS '绘画状态';
COMMENT ON COLUMN public.ai_image_record.finish_time IS '完成时间';
COMMENT ON COLUMN public.ai_image_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_image_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_image_record.pic_url IS '图片地址';
COMMENT ON COLUMN public.ai_image_record.options IS '绘制参数';
COMMENT ON COLUMN public.ai_image_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_image_record.buttons IS 'mj buttons 按钮';
COMMENT ON COLUMN public.ai_image_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_image_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_image_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_image_record ADD CONSTRAINT ai_image_record_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_image_record_platform ON public.ai_image_record USING btree (platform);
CREATE INDEX idx_ai_image_record_tenant_id ON public.ai_image_record USING btree (tenant_id);
CREATE INDEX idx_ai_image_record_user_id ON public.ai_image_record USING btree (admin_id);
