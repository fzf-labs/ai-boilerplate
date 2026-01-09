CREATE TABLE public.ai_video_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    prompt text NOT NULL,
    platform character varying(64) NOT NULL,
    model_id character varying(64),
    model character varying(64) NOT NULL,
    status integer NOT NULL,
    finish_time timestamp with time zone,
    error_message character varying(1024),
    public_status boolean DEFAULT false NOT NULL,
    video_url character varying(512),
    options jsonb,
    task_id character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_video_record IS 'AI 视频表';
COMMENT ON COLUMN public.ai_video_record.id IS '编号';
COMMENT ON COLUMN public.ai_video_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_video_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_video_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_video_record.platform IS '平台';
COMMENT ON COLUMN public.ai_video_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_video_record.model IS '模型';
COMMENT ON COLUMN public.ai_video_record.status IS '状态';
COMMENT ON COLUMN public.ai_video_record.finish_time IS '完成时间';
COMMENT ON COLUMN public.ai_video_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_video_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_video_record.video_url IS '视频地址';
COMMENT ON COLUMN public.ai_video_record.options IS '绘制参数';
COMMENT ON COLUMN public.ai_video_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_video_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_video_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_video_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_video_record ADD CONSTRAINT ai_video_record_pkey PRIMARY KEY (id);
