CREATE TABLE public.sys_menu (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid character varying(64) NOT NULL,
    name character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    path character varying(100) NOT NULL,
    permission character varying(64),
    icon character varying(64),
    component character varying(100),
    component_name character varying(100),
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_menu IS '菜单和权限规则表';
COMMENT ON COLUMN public.sys_menu.id IS 'id';
COMMENT ON COLUMN public.sys_menu.pid IS '上级菜单';
COMMENT ON COLUMN public.sys_menu.name IS '菜单名称';
COMMENT ON COLUMN public.sys_menu.type IS '菜单类型(dir,menu,button)';
COMMENT ON COLUMN public.sys_menu.path IS '路由路径';
COMMENT ON COLUMN public.sys_menu.permission IS '权限标识';
COMMENT ON COLUMN public.sys_menu.icon IS '图标';
COMMENT ON COLUMN public.sys_menu.component IS '组件路径';
COMMENT ON COLUMN public.sys_menu.component_name IS '组件名';
COMMENT ON COLUMN public.sys_menu.sort IS '权重(排序)';
COMMENT ON COLUMN public.sys_menu.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_menu.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_menu.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_menu.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_menu ADD CONSTRAINT sys_menu_pkey PRIMARY KEY (id);
CREATE INDEX sys_menu_pid_idx ON public.sys_menu USING btree (pid);
