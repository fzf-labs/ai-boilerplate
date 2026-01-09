CREATE TABLE public.config_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    key character varying(64) NOT NULL,
    value jsonb,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.config_data IS '配置管理';
COMMENT ON COLUMN public.config_data.id IS 'id';
COMMENT ON COLUMN public.config_data.name IS '名称';
COMMENT ON COLUMN public.config_data.key IS '健';
COMMENT ON COLUMN public.config_data.value IS '值';
COMMENT ON COLUMN public.config_data.remark IS '备注';
COMMENT ON COLUMN public.config_data.status IS '状态';
COMMENT ON COLUMN public.config_data.created_at IS '创建时间';
COMMENT ON COLUMN public.config_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.config_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.config_data ADD CONSTRAINT config_data_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX config_data_key_idx ON public.config_data USING btree (key);
