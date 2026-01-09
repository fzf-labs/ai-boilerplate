CREATE TABLE public.sys_admin (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying,
    username character varying NOT NULL,
    password character varying NOT NULL,
    nickname character varying,
    avatar character varying,
    email character varying,
    sex smallint,
    mobile character varying,
    role_id character varying,
    dept_id character varying,
    post_id character varying,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_admin IS '系统-管理员';
COMMENT ON COLUMN public.sys_admin.id IS 'id';
COMMENT ON COLUMN public.sys_admin.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_admin.username IS '用户名';
COMMENT ON COLUMN public.sys_admin.password IS '密码';
COMMENT ON COLUMN public.sys_admin.nickname IS '昵称';
COMMENT ON COLUMN public.sys_admin.avatar IS '头像';
COMMENT ON COLUMN public.sys_admin.email IS '邮件';
COMMENT ON COLUMN public.sys_admin.sex IS '性别';
COMMENT ON COLUMN public.sys_admin.mobile IS '手机号';
COMMENT ON COLUMN public.sys_admin.role_id IS '角色Id';
COMMENT ON COLUMN public.sys_admin.dept_id IS '部门';
COMMENT ON COLUMN public.sys_admin.post_id IS '岗位';
COMMENT ON COLUMN public.sys_admin.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_admin.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_admin.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_admin.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_admin ADD CONSTRAINT sys_admin_pkey PRIMARY KEY (id);
CREATE INDEX sys_admin_dept_id_idx ON public.sys_admin USING btree (dept_id);
CREATE INDEX sys_admin_post_id_idx ON public.sys_admin USING btree (post_id);
CREATE INDEX sys_admin_role_id_idx ON public.sys_admin USING btree (role_id);
CREATE UNIQUE INDEX sys_admin_username_idx ON public.sys_admin USING btree (username);
