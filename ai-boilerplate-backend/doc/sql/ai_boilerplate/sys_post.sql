CREATE TABLE public.sys_post (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64),
    name character varying(50) NOT NULL,
    code character varying(32),
    remark character varying(255),
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_post IS '系统-工作岗位';
COMMENT ON COLUMN public.sys_post.id IS '编号';
COMMENT ON COLUMN public.sys_post.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_post.name IS '岗位名称';
COMMENT ON COLUMN public.sys_post.code IS '岗位编码';
COMMENT ON COLUMN public.sys_post.remark IS '备注';
COMMENT ON COLUMN public.sys_post.sort IS '排序值';
COMMENT ON COLUMN public.sys_post.status IS '0=禁用 1=开启 ';
COMMENT ON COLUMN public.sys_post.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_post.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_post.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_post ADD CONSTRAINT sys_post_pkey PRIMARY KEY (id);
CREATE INDEX sys_post_tenant_id_idx ON public.sys_post USING btree (tenant_id);
