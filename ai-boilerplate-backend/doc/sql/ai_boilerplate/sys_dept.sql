CREATE TABLE public.sys_dept (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid character varying(64),
    name character varying(64) NOT NULL,
    admin_id character varying(64),
    status smallint DEFAULT 1 NOT NULL,
    sort bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    tenant_id character varying(64)
);
COMMENT ON TABLE public.sys_dept IS '系统-部门';
COMMENT ON COLUMN public.sys_dept.id IS '编号';
COMMENT ON COLUMN public.sys_dept.pid IS '父级id';
COMMENT ON COLUMN public.sys_dept.name IS '部门简称';
COMMENT ON COLUMN public.sys_dept.admin_id IS '负责人Id';
COMMENT ON COLUMN public.sys_dept.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_dept.sort IS '排序值';
COMMENT ON COLUMN public.sys_dept.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_dept.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_dept.deleted_at IS '删除时间';
COMMENT ON COLUMN public.sys_dept.tenant_id IS '租户Id';
ALTER TABLE ONLY public.sys_dept ADD CONSTRAINT sys_dept_pkey PRIMARY KEY (id);
CREATE INDEX sys_dept_pid_idx ON public.sys_dept USING btree (pid);
