CREATE TABLE public.dict_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(100) NOT NULL,
    type character varying(100) NOT NULL,
    status smallint NOT NULL,
    remark character varying(500),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.dict_type IS '字典类型表';
COMMENT ON COLUMN public.dict_type.id IS 'id';
COMMENT ON COLUMN public.dict_type.name IS '字典名称';
COMMENT ON COLUMN public.dict_type.type IS '字典类型';
COMMENT ON COLUMN public.dict_type.status IS '状态（1正常-1停用）';
COMMENT ON COLUMN public.dict_type.remark IS '备注';
COMMENT ON COLUMN public.dict_type.created_at IS '创建时间';
COMMENT ON COLUMN public.dict_type.updated_at IS '更新时间';
COMMENT ON COLUMN public.dict_type.deleted_at IS '删除时间';
ALTER TABLE ONLY public.dict_type ADD CONSTRAINT dict_type_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX dict_type_key_idx ON public.dict_type USING btree (type);
