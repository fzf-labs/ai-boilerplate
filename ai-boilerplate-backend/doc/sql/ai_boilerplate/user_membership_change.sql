CREATE TABLE public.user_membership_change (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    source_type character varying(20) NOT NULL,
    source_id character varying(128) NOT NULL,
    before_membership_type character varying(20),
    after_membership_type character varying(20),
    before_expired_at timestamp with time zone,
    after_expired_at timestamp with time zone,
    duration_days integer,
    remark character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_membership_change IS '用户会员变更记录表';
COMMENT ON COLUMN public.user_membership_change.id IS 'id';
COMMENT ON COLUMN public.user_membership_change.user_id IS '用户ID';
COMMENT ON COLUMN public.user_membership_change.source_type IS '来源类型(order,activation_code,admin)';
COMMENT ON COLUMN public.user_membership_change.source_id IS '来源ID(订单ID/激活码)';
COMMENT ON COLUMN public.user_membership_change.before_membership_type IS '变更前会员类型';
COMMENT ON COLUMN public.user_membership_change.after_membership_type IS '变更后会员类型';
COMMENT ON COLUMN public.user_membership_change.before_expired_at IS '变更前到期时间';
COMMENT ON COLUMN public.user_membership_change.after_expired_at IS '变更后到期时间';
COMMENT ON COLUMN public.user_membership_change.duration_days IS '变更时长(天)';
COMMENT ON COLUMN public.user_membership_change.remark IS '备注';
COMMENT ON COLUMN public.user_membership_change.created_at IS '创建时间';
COMMENT ON COLUMN public.user_membership_change.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_membership_change.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_membership_change ADD CONSTRAINT user_membership_change_pkey PRIMARY KEY (id);
CREATE INDEX user_membership_change_user_id_created_at_idx ON public.user_membership_change USING btree (user_id, created_at);
CREATE INDEX user_membership_change_source_type_idx ON public.user_membership_change USING btree (source_type);
