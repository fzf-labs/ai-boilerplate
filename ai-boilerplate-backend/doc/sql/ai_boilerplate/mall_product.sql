CREATE TABLE public.mall_product (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    product_type character varying(20) NOT NULL,
    product_name character varying(100) NOT NULL,
    product_desc character varying(500),
    product_images jsonb,
    product_detail jsonb,
    product_config jsonb,
    original_price numeric(10,2) NOT NULL,
    current_price numeric(10,2) NOT NULL,
    stock_quantity integer DEFAULT '-1'::integer,
    sold_quantity integer DEFAULT 0,
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mall_product IS '商品表';
COMMENT ON COLUMN public.mall_product.id IS 'id';
COMMENT ON COLUMN public.mall_product.product_type IS '商品类型(membership:会员,service:增值服务,goods:商品)';
COMMENT ON COLUMN public.mall_product.product_name IS '商品名称';
COMMENT ON COLUMN public.mall_product.product_desc IS '商品描述';
COMMENT ON COLUMN public.mall_product.product_images IS '商品图片(多个用逗号分隔)';
COMMENT ON COLUMN public.mall_product.product_detail IS '商品详情(JSON格式,包含特色功能等)';
COMMENT ON COLUMN public.mall_product.product_config IS '商品配置(JSON格式)';
COMMENT ON COLUMN public.mall_product.original_price IS '原价';
COMMENT ON COLUMN public.mall_product.current_price IS '现价';
COMMENT ON COLUMN public.mall_product.stock_quantity IS '库存数量(-1表示无限库存)';
COMMENT ON COLUMN public.mall_product.sold_quantity IS '已售数量';
COMMENT ON COLUMN public.mall_product.sort IS '排序';
COMMENT ON COLUMN public.mall_product.status IS '状态(-1下架,0待上架,1在售,2售罄)';
COMMENT ON COLUMN public.mall_product.created_at IS '创建时间';
COMMENT ON COLUMN public.mall_product.updated_at IS '更新时间';
COMMENT ON COLUMN public.mall_product.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mall_product ADD CONSTRAINT mall_product_pkey PRIMARY KEY (id);
CREATE INDEX mall_product_product_type_idx ON public.mall_product USING btree (product_type);
CREATE INDEX mall_product_sort_idx ON public.mall_product USING btree (sort);
CREATE INDEX mall_product_status_idx ON public.mall_product USING btree (status);
