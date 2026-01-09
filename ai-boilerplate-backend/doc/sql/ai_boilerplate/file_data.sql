CREATE TABLE public.file_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    config_id character varying(64),
    name character varying(255),
    path character varying(512) NOT NULL,
    url character varying(1024) NOT NULL,
    ext character varying(32),
    size integer NOT NULL,
    status integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.file_data IS '文件表';
COMMENT ON COLUMN public.file_data.id IS '文件编号';
COMMENT ON COLUMN public.file_data.config_id IS '配置编号';
COMMENT ON COLUMN public.file_data.name IS '文件名';
COMMENT ON COLUMN public.file_data.path IS '文件路径';
COMMENT ON COLUMN public.file_data.url IS '文件 URL';
COMMENT ON COLUMN public.file_data.ext IS '文件类型';
COMMENT ON COLUMN public.file_data.size IS '文件大小';
COMMENT ON COLUMN public.file_data.status IS '状态（-1失败,1未知,2 成功）';
COMMENT ON COLUMN public.file_data.created_at IS '创建时间';
COMMENT ON COLUMN public.file_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.file_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.file_data ADD CONSTRAINT file_data_pkey PRIMARY KEY (id);
