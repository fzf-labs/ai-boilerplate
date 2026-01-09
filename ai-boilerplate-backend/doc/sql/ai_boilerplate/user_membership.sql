CREATE TABLE public.user_membership (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    membership_type character varying(20) NOT NULL,
    expired_at timestamp with time zone,
    auto_renew integer DEFAULT 0,
    auto_renew_days integer,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_membership IS '用户会员关系表';
COMMENT ON COLUMN public.user_membership.id IS 'id';
COMMENT ON COLUMN public.user_membership.user_id IS '用户ID';
COMMENT ON COLUMN public.user_membership.membership_type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.user_membership.expired_at IS '到期时间(普通会员为NULL,表示永不过期)';
COMMENT ON COLUMN public.user_membership.auto_renew IS '是否自动续费(0否,1是)';
COMMENT ON COLUMN public.user_membership.auto_renew_days IS '自动续费天数';
COMMENT ON COLUMN public.user_membership.status IS '状态(-1禁用,1正常)';
COMMENT ON COLUMN public.user_membership.created_at IS '创建时间';
COMMENT ON COLUMN public.user_membership.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_membership.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_membership ADD CONSTRAINT user_membership_pkey PRIMARY KEY (id);
CREATE INDEX user_membership_expired_at_idx ON public.user_membership USING btree (expired_at);
CREATE INDEX user_membership_membership_type_code_idx ON public.user_membership USING btree (membership_type);
CREATE UNIQUE INDEX user_membership_user_id_idx ON public.user_membership USING btree (user_id);
