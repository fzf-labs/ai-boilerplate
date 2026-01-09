CREATE TABLE public.membership (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    type character varying(20) NOT NULL,
    description character varying(255),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership IS '会员类型配置表';
COMMENT ON COLUMN public.membership.id IS 'id';
COMMENT ON COLUMN public.membership.name IS '会员类型名称';
COMMENT ON COLUMN public.membership.type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.membership.description IS '会员类型描述';
COMMENT ON COLUMN public.membership.sort IS '排序';
COMMENT ON COLUMN public.membership.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership.created_at IS '创建时间';
COMMENT ON COLUMN public.membership.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership ADD CONSTRAINT membership_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX membership_code_idx ON public.membership USING btree (type);
CREATE INDEX membership_sort_idx ON public.membership USING btree (sort);
