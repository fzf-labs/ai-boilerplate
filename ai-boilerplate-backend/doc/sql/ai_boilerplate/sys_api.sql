CREATE TABLE public.sys_api (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    permission_id uuid NOT NULL,
    method character varying(32) NOT NULL,
    path character varying(255) NOT NULL,
    "desc" character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_api IS '系统-接口';
COMMENT ON COLUMN public.sys_api.id IS 'id';
COMMENT ON COLUMN public.sys_api.permission_id IS '权限Id';
COMMENT ON COLUMN public.sys_api.method IS '方法';
COMMENT ON COLUMN public.sys_api.path IS '路径';
COMMENT ON COLUMN public.sys_api."desc" IS '描述';
COMMENT ON COLUMN public.sys_api.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_api.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_api.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_api ADD CONSTRAINT sys_api_pkey PRIMARY KEY (id);
CREATE INDEX sys_api_permission_id_idx ON public.sys_api USING btree (permission_id);
