CREATE TABLE public.mall_activation_code (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    product_type character varying(20) NOT NULL,
    product_id character varying(64) NOT NULL,
    batch_no character varying(64) NOT NULL,
    code character varying(32) NOT NULL,
    valid_st timestamp with time zone NOT NULL,
    valid_ed timestamp with time zone NOT NULL,
    activated_at timestamp with time zone,
    user_id character varying(64),
    user_change jsonb,
    platform character varying(20),
    platform_sold_at timestamp with time zone,
    platform_order_no character varying(100),
    platform_buyer_id character varying(100),
    platform_buyer_name character varying(100),
    remark character varying(255),
    status integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mall_activation_code IS '激活码管理表';
COMMENT ON COLUMN public.mall_activation_code.id IS 'id';
COMMENT ON COLUMN public.mall_activation_code.product_type IS '商品类型(membership:会员,service:服务)';
COMMENT ON COLUMN public.mall_activation_code.product_id IS '商品ID';
COMMENT ON COLUMN public.mall_activation_code.batch_no IS '批次号';
COMMENT ON COLUMN public.mall_activation_code.code IS '激活码';
COMMENT ON COLUMN public.mall_activation_code.valid_st IS '激活码有效期开始时间';
COMMENT ON COLUMN public.mall_activation_code.valid_ed IS '激活码有效期截止时间';
COMMENT ON COLUMN public.mall_activation_code.activated_at IS '激活时间';
COMMENT ON COLUMN public.mall_activation_code.user_id IS '用户ID';
COMMENT ON COLUMN public.mall_activation_code.user_change IS '用户属性变化';
COMMENT ON COLUMN public.mall_activation_code.platform IS '平台';
COMMENT ON COLUMN public.mall_activation_code.platform_sold_at IS '平台售出时间';
COMMENT ON COLUMN public.mall_activation_code.platform_order_no IS '平台订单号';
COMMENT ON COLUMN public.mall_activation_code.platform_buyer_id IS '平台买家ID';
COMMENT ON COLUMN public.mall_activation_code.platform_buyer_name IS '平台买家昵称';
COMMENT ON COLUMN public.mall_activation_code.remark IS '备注';
COMMENT ON COLUMN public.mall_activation_code.status IS '状态(-2已退款,-1禁用,0库存,1已售出,2已激活,3已过期)';
COMMENT ON COLUMN public.mall_activation_code.created_at IS '创建时间';
COMMENT ON COLUMN public.mall_activation_code.updated_at IS '更新时间';
COMMENT ON COLUMN public.mall_activation_code.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mall_activation_code ADD CONSTRAINT mall_activation_code_pkey PRIMARY KEY (id);
CREATE INDEX mall_activation_code_activated_at_idx ON public.mall_activation_code USING btree (activated_at);
CREATE INDEX mall_activation_code_activated_user_id_idx ON public.mall_activation_code USING btree (user_id);
CREATE UNIQUE INDEX mall_activation_code_idx ON public.mall_activation_code USING btree (code);
CREATE INDEX mall_activation_code_status_idx ON public.mall_activation_code USING btree (status);
