CREATE TABLE public.ai_audio_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    lyric text,
    image_url character varying(600),
    audio_url character varying(600),
    status integer NOT NULL,
    description text,
    prompt text,
    platform character varying(64) NOT NULL,
    model_id character varying(64) NOT NULL,
    model character varying(50) NOT NULL,
    generate_mode integer NOT NULL,
    tags character varying(600),
    duration double precision,
    public_status boolean DEFAULT false NOT NULL,
    task_id character varying(255),
    error_message character varying(1024),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_audio_record IS 'AI 音乐表';
COMMENT ON COLUMN public.ai_audio_record.id IS '编号';
COMMENT ON COLUMN public.ai_audio_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_audio_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_audio_record.title IS '音乐名称';
COMMENT ON COLUMN public.ai_audio_record.lyric IS '歌词';
COMMENT ON COLUMN public.ai_audio_record.image_url IS '图片地址';
COMMENT ON COLUMN public.ai_audio_record.audio_url IS '音频地址';
COMMENT ON COLUMN public.ai_audio_record.status IS '音频状态';
COMMENT ON COLUMN public.ai_audio_record.description IS '描述词';
COMMENT ON COLUMN public.ai_audio_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_audio_record.platform IS '模型平台';
COMMENT ON COLUMN public.ai_audio_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_audio_record.model IS '模型';
COMMENT ON COLUMN public.ai_audio_record.generate_mode IS '生成模式';
COMMENT ON COLUMN public.ai_audio_record.tags IS '风格标签';
COMMENT ON COLUMN public.ai_audio_record.duration IS '时长';
COMMENT ON COLUMN public.ai_audio_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_audio_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_audio_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_audio_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_audio_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_audio_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_audio_record ADD CONSTRAINT ai_audio_record_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX ai_music_record_pkey ON public.ai_audio_record USING btree (id);
CREATE INDEX idx_ai_music_record_platform ON public.ai_audio_record USING btree (platform);
CREATE INDEX idx_ai_music_record_tenant_id ON public.ai_audio_record USING btree (tenant_id);
CREATE INDEX idx_ai_music_record_user_id ON public.ai_audio_record USING btree (admin_id);
