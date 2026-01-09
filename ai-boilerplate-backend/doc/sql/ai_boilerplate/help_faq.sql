CREATE TABLE public.help_faq (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    category_id uuid NOT NULL,
    question character varying(500) NOT NULL,
    answer text NOT NULL,
    "order" integer DEFAULT 0 NOT NULL,
    view_count integer DEFAULT 0 NOT NULL,
    helpful_count integer DEFAULT 0 NOT NULL,
    unhelpful_count integer DEFAULT 0 NOT NULL,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.help_faq IS '常见问题表';
COMMENT ON COLUMN public.help_faq.id IS 'id';
COMMENT ON COLUMN public.help_faq.category_id IS '分类ID';
COMMENT ON COLUMN public.help_faq.question IS '问题';
COMMENT ON COLUMN public.help_faq.answer IS '答案';
COMMENT ON COLUMN public.help_faq."order" IS '排序';
COMMENT ON COLUMN public.help_faq.view_count IS '查看次数';
COMMENT ON COLUMN public.help_faq.helpful_count IS '有帮助次数';
COMMENT ON COLUMN public.help_faq.unhelpful_count IS '无帮助次数';
COMMENT ON COLUMN public.help_faq.status IS '状态(1启用 0禁用)';
COMMENT ON COLUMN public.help_faq.created_at IS '创建时间';
COMMENT ON COLUMN public.help_faq.updated_at IS '更新时间';
COMMENT ON COLUMN public.help_faq.deleted_at IS '删除时间';
ALTER TABLE ONLY public.help_faq ADD CONSTRAINT help_faq_pkey PRIMARY KEY (id);
CREATE INDEX help_faq_category_id_idx ON public.help_faq USING btree (category_id);
CREATE INDEX help_faq_order_idx ON public.help_faq USING btree ("order");
