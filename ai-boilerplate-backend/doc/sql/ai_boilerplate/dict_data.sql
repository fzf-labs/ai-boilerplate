CREATE TABLE public.dict_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type character varying(100) NOT NULL,
    label character varying(100) NOT NULL,
    key character varying(100) NOT NULL,
    value character varying(100) NOT NULL,
    remark character varying(500),
    css_color character varying(100),
    css_class character varying(100),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.dict_data IS '字典数据表';
COMMENT ON COLUMN public.dict_data.id IS 'id';
COMMENT ON COLUMN public.dict_data.type IS '字典类型';
COMMENT ON COLUMN public.dict_data.label IS '字典标签';
COMMENT ON COLUMN public.dict_data.key IS '字典键';
COMMENT ON COLUMN public.dict_data.value IS '字典值';
COMMENT ON COLUMN public.dict_data.remark IS '备注';
COMMENT ON COLUMN public.dict_data.css_color IS 'css 颜色';
COMMENT ON COLUMN public.dict_data.css_class IS 'css 样式';
COMMENT ON COLUMN public.dict_data.status IS '状态（0正常 1停用）';
COMMENT ON COLUMN public.dict_data.created_at IS '创建时间';
COMMENT ON COLUMN public.dict_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.dict_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.dict_data ADD CONSTRAINT dict_data_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX dict_data_dict_type_key_key_idx ON public.dict_data USING btree (type, key);
