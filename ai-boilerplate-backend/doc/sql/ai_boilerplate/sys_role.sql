CREATE TABLE public.sys_role (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    name character varying(50) NOT NULL,
    remark character varying(200),
    "dataScope" character varying(64) NOT NULL,
    "menuIds" jsonb,
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_role IS '系统-角色';
COMMENT ON COLUMN public.sys_role.id IS '编号';
COMMENT ON COLUMN public.sys_role.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_role.name IS '名称';
COMMENT ON COLUMN public.sys_role.remark IS '备注';
COMMENT ON COLUMN public.sys_role."dataScope" IS '数据范围';
COMMENT ON COLUMN public.sys_role."menuIds" IS '菜单';
COMMENT ON COLUMN public.sys_role.sort IS '排序值';
COMMENT ON COLUMN public.sys_role.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_role.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_role.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_role.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_role ADD CONSTRAINT sys_role_pkey PRIMARY KEY (id);
CREATE INDEX sys_role_tenant_id_idx ON public.sys_role USING btree (tenant_id);
