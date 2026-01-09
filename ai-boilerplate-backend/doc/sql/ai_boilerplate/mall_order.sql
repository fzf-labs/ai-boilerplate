CREATE TABLE public.mall_order (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id character varying(64) NOT NULL,
    product_type character varying(20) NOT NULL,
    product_id character varying(64) NOT NULL,
    original_amount numeric(10,2) NOT NULL,
    discount_amount numeric(10,2) DEFAULT 0.00,
    actual_amount numeric(10,2) NOT NULL,
    refund_amount numeric(10,2) NOT NULL,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_method character varying(50),
    payment_status integer DEFAULT 0,
    payment_time timestamp with time zone,
    delivery_time timestamp with time zone,
    expired_time timestamp with time zone,
    remark character varying(500),
    status character varying DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mall_order IS '订单信息表';
COMMENT ON COLUMN public.mall_order.id IS 'id';
COMMENT ON COLUMN public.mall_order.user_id IS '用户ID';
COMMENT ON COLUMN public.mall_order.product_type IS '商品类型(membership:会员,service:服务,goods:商品)';
COMMENT ON COLUMN public.mall_order.product_id IS '商品ID';
COMMENT ON COLUMN public.mall_order.original_amount IS '原价';
COMMENT ON COLUMN public.mall_order.discount_amount IS '优惠金额';
COMMENT ON COLUMN public.mall_order.actual_amount IS '实付金额';
COMMENT ON COLUMN public.mall_order.refund_amount IS '退款金额';
COMMENT ON COLUMN public.mall_order.currency IS '币种';
COMMENT ON COLUMN public.mall_order.payment_method IS '支付方式(微信,支付宝)';
COMMENT ON COLUMN public.mall_order.payment_status IS '支付状态(0待支付,1已支付,2支付失败,3已退款)';
COMMENT ON COLUMN public.mall_order.payment_time IS '支付时间';
COMMENT ON COLUMN public.mall_order.delivery_time IS '确认时间';
COMMENT ON COLUMN public.mall_order.expired_time IS '订单过期时间';
COMMENT ON COLUMN public.mall_order.remark IS '备注';
COMMENT ON COLUMN public.mall_order.status IS '状态(待付款pendingPayment,待发货pendingDelivery,待收货pendingReceipt,已完成completed,已取消canceled,已退款refunded)';
COMMENT ON COLUMN public.mall_order.created_at IS '创建时间';
COMMENT ON COLUMN public.mall_order.updated_at IS '更新时间';
COMMENT ON COLUMN public.mall_order.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mall_order ADD CONSTRAINT mall_order_pkey PRIMARY KEY (id);
CREATE INDEX mall_order_info_created_at_idx ON public.mall_order USING btree (created_at);
CREATE UNIQUE INDEX mall_order_info_order_no_idx ON public.mall_order USING btree (product_id);
CREATE INDEX mall_order_info_payment_status_idx ON public.mall_order USING btree (payment_status);
CREATE UNIQUE INDEX mall_order_info_pkey ON public.mall_order USING btree (id);
CREATE INDEX mall_order_info_user_id_idx ON public.mall_order USING btree (user_id);
