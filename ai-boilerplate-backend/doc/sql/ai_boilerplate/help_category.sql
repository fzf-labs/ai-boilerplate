CREATE TABLE public.help_category (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(100) NOT NULL,
    icon character varying(50),
    "order" integer DEFAULT 0 NOT NULL,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.help_category IS '帮助分类表';
COMMENT ON COLUMN public.help_category.id IS 'id';
COMMENT ON COLUMN public.help_category.name IS '分类名称';
COMMENT ON COLUMN public.help_category.icon IS '图标';
COMMENT ON COLUMN public.help_category."order" IS '排序';
COMMENT ON COLUMN public.help_category.status IS '状态(1启用 0禁用)';
COMMENT ON COLUMN public.help_category.created_at IS '创建时间';
COMMENT ON COLUMN public.help_category.updated_at IS '更新时间';
COMMENT ON COLUMN public.help_category.deleted_at IS '删除时间';
ALTER TABLE ONLY public.help_category ADD CONSTRAINT help_category_pkey PRIMARY KEY (id);
CREATE INDEX help_category_order_idx ON public.help_category USING btree ("order");
