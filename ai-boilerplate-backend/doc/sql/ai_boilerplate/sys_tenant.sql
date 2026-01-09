CREATE TABLE public.sys_tenant (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    remark character varying(200),
    admin_id character varying(50),
    expire_time timestamp with time zone,
    "menuIds" jsonb,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_tenant IS '系统-租户';
COMMENT ON COLUMN public.sys_tenant.id IS 'id';
COMMENT ON COLUMN public.sys_tenant.name IS '名称';
COMMENT ON COLUMN public.sys_tenant.remark IS '备注';
COMMENT ON COLUMN public.sys_tenant.admin_id IS '租户管理员Id';
COMMENT ON COLUMN public.sys_tenant.expire_time IS '过期时间';
COMMENT ON COLUMN public.sys_tenant."menuIds" IS '菜单';
COMMENT ON COLUMN public.sys_tenant.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_tenant.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_tenant.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_tenant.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_tenant ADD CONSTRAINT sys_tenant_pkey PRIMARY KEY (id);
CREATE INDEX sys_tenant_name_idx ON public.sys_tenant USING btree (name);
