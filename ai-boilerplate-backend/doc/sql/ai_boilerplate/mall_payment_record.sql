CREATE TABLE public.mall_payment_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    order_id uuid NOT NULL,
    transaction_id character varying(128) NOT NULL,
    payment_channel character varying(50) NOT NULL,
    payment_method character varying(50) NOT NULL,
    amount numeric(10,2) NOT NULL,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_status integer DEFAULT 0,
    third_party_order_no character varying(128),
    third_party_transaction_id character varying(128),
    callback_data jsonb,
    callback_time timestamp with time zone,
    error_code character varying(50),
    error_message character varying(500),
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mall_payment_record IS '支付记录表';
COMMENT ON COLUMN public.mall_payment_record.id IS 'id';
COMMENT ON COLUMN public.mall_payment_record.order_id IS '订单ID';
COMMENT ON COLUMN public.mall_payment_record.transaction_id IS '交易流水号';
COMMENT ON COLUMN public.mall_payment_record.payment_channel IS '支付渠道(wechat,alipay)';
COMMENT ON COLUMN public.mall_payment_record.payment_method IS '支付方式(mini_program,h5,native,jsapi)';
COMMENT ON COLUMN public.mall_payment_record.amount IS '支付金额';
COMMENT ON COLUMN public.mall_payment_record.currency IS '币种';
COMMENT ON COLUMN public.mall_payment_record.payment_status IS '支付状态(0待支付,1支付成功,2支付失败,3已退款)';
COMMENT ON COLUMN public.mall_payment_record.third_party_order_no IS '第三方订单号';
COMMENT ON COLUMN public.mall_payment_record.third_party_transaction_id IS '第三方交易号';
COMMENT ON COLUMN public.mall_payment_record.callback_data IS '回调数据';
COMMENT ON COLUMN public.mall_payment_record.callback_time IS '回调时间';
COMMENT ON COLUMN public.mall_payment_record.error_code IS '错误代码';
COMMENT ON COLUMN public.mall_payment_record.error_message IS '错误信息';
COMMENT ON COLUMN public.mall_payment_record.status IS '状态(-1无效,1正常)';
COMMENT ON COLUMN public.mall_payment_record.created_at IS '创建时间';
COMMENT ON COLUMN public.mall_payment_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.mall_payment_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mall_payment_record ADD CONSTRAINT mall_payment_record_pkey PRIMARY KEY (id);
CREATE INDEX mall_payment_record_order_id_idx ON public.mall_payment_record USING btree (order_id);
CREATE INDEX mall_payment_record_payment_status_idx ON public.mall_payment_record USING btree (payment_status);
CREATE INDEX mall_payment_record_third_party_order_no_idx ON public.mall_payment_record USING btree (third_party_order_no);
CREATE UNIQUE INDEX mall_payment_record_transaction_id_idx ON public.mall_payment_record USING btree (transaction_id);
