CREATE TABLE public.file_config (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    storage character varying(32) NOT NULL,
    remark character varying(255),
    master boolean NOT NULL,
    config jsonb,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.file_config IS '文件配置表';
COMMENT ON COLUMN public.file_config.id IS '编号';
COMMENT ON COLUMN public.file_config.name IS '配置名';
COMMENT ON COLUMN public.file_config.storage IS '存储器';
COMMENT ON COLUMN public.file_config.remark IS '备注';
COMMENT ON COLUMN public.file_config.master IS '是否为主配置';
COMMENT ON COLUMN public.file_config.config IS '存储配置';
COMMENT ON COLUMN public.file_config.created_at IS '创建时间';
COMMENT ON COLUMN public.file_config.updated_at IS '更新时间';
COMMENT ON COLUMN public.file_config.deleted_at IS '删除时间';
ALTER TABLE ONLY public.file_config ADD CONSTRAINT file_config_pkey PRIMARY KEY (id);
CREATE INDEX file_config_master_idx ON public.file_config USING btree (master);
