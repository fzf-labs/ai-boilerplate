CREATE TABLE public.sys_notice (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    content character varying(200) NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_notice IS '系统-公告';
COMMENT ON COLUMN public.sys_notice.id IS 'id';
COMMENT ON COLUMN public.sys_notice.tenant_id IS '租户id';
COMMENT ON COLUMN public.sys_notice.type IS '类型';
COMMENT ON COLUMN public.sys_notice.title IS '标题';
COMMENT ON COLUMN public.sys_notice.content IS '内容';
COMMENT ON COLUMN public.sys_notice.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_notice.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_notice.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_notice.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_notice ADD CONSTRAINT sys_notice_pkey PRIMARY KEY (id);
