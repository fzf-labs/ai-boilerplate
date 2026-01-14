CREATE TABLE public.membership_benefit_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    benefit_key character varying(100) NOT NULL,
    benefit_name character varying(255) NOT NULL,
    benefit_icon character varying(500),
    benefit_desc character varying(500),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership_benefit_type IS '会员权益类型表';
COMMENT ON COLUMN public.membership_benefit_type.id IS 'id';
COMMENT ON COLUMN public.membership_benefit_type.benefit_key IS '权益标识';
COMMENT ON COLUMN public.membership_benefit_type.benefit_name IS '权益名称';
COMMENT ON COLUMN public.membership_benefit_type.benefit_icon IS '权益图标';
COMMENT ON COLUMN public.membership_benefit_type.benefit_desc IS '权益描述';
COMMENT ON COLUMN public.membership_benefit_type.sort IS '排序';
COMMENT ON COLUMN public.membership_benefit_type.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership_benefit_type.created_at IS '创建时间';
COMMENT ON COLUMN public.membership_benefit_type.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership_benefit_type.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership_benefit_type ADD CONSTRAINT membership_benefit_type_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX membership_benefit_type_key_idx ON public.membership_benefit_type USING btree (benefit_key);
CREATE INDEX membership_benefit_type_sort_idx ON public.membership_benefit_type USING btree (sort);
