CREATE TABLE public.self_app (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    package_name character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.self_app IS '自应用信息表';
COMMENT ON COLUMN public.self_app.id IS 'ID';
COMMENT ON COLUMN public.self_app.package_name IS '包名';
COMMENT ON COLUMN public.self_app.name IS '应用名称';
COMMENT ON COLUMN public.self_app.description IS '应用描述';
COMMENT ON COLUMN public.self_app.status IS '状态(-1禁用 1启用)';
COMMENT ON COLUMN public.self_app.created_at IS '创建时间';
COMMENT ON COLUMN public.self_app.updated_at IS '更新时间';
COMMENT ON COLUMN public.self_app.deleted_at IS '删除时间';
ALTER TABLE ONLY public.self_app ADD CONSTRAINT self_app_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX self_app_package_name_idx ON public.self_app USING btree (package_name);
