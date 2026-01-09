CREATE TABLE public.membership_benefit (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    membership_type character varying(20) NOT NULL,
    benefit_key character varying(100) NOT NULL,
    benefit_name character varying(255) NOT NULL,
    benefit_desc character varying(500),
    benefit_value character varying(100),
    benefit_num character varying(100),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership_benefit IS '会员权益配置表';
COMMENT ON COLUMN public.membership_benefit.id IS 'id';
COMMENT ON COLUMN public.membership_benefit.membership_type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.membership_benefit.benefit_key IS '权益标识';
COMMENT ON COLUMN public.membership_benefit.benefit_name IS '权益名称';
COMMENT ON COLUMN public.membership_benefit.benefit_desc IS '权益描述';
COMMENT ON COLUMN public.membership_benefit.benefit_value IS '权益值';
COMMENT ON COLUMN public.membership_benefit.benefit_num IS '权益次数';
COMMENT ON COLUMN public.membership_benefit.sort IS '排序';
COMMENT ON COLUMN public.membership_benefit.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership_benefit.created_at IS '创建时间';
COMMENT ON COLUMN public.membership_benefit.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership_benefit.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership_benefit ADD CONSTRAINT membership_benefit_pkey PRIMARY KEY (id);
CREATE INDEX membership_benefit_membership_code_idx ON public.membership_benefit USING btree (membership_type);
CREATE INDEX membership_benefit_sort_order_idx ON public.membership_benefit USING btree (sort);
CREATE UNIQUE INDEX membership_benefit_type_key_idx ON public.membership_benefit USING btree (membership_type, benefit_key);
