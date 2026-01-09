CREATE TABLE public.help_feedback (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    category character varying(50) NOT NULL,
    description text NOT NULL,
    images jsonb,
    contact character varying(100),
    status integer DEFAULT 0 NOT NULL,
    reply text,
    replied_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.help_feedback IS '用户反馈表';
COMMENT ON COLUMN public.help_feedback.id IS 'id';
COMMENT ON COLUMN public.help_feedback.user_id IS '用户ID';
COMMENT ON COLUMN public.help_feedback.category IS '问题分类';
COMMENT ON COLUMN public.help_feedback.description IS '问题描述';
COMMENT ON COLUMN public.help_feedback.images IS '图片列表';
COMMENT ON COLUMN public.help_feedback.contact IS '联系方式';
COMMENT ON COLUMN public.help_feedback.status IS '状态(0待处理 1已处理 2已关闭)';
COMMENT ON COLUMN public.help_feedback.reply IS '回复内容';
COMMENT ON COLUMN public.help_feedback.replied_at IS '回复时间';
COMMENT ON COLUMN public.help_feedback.created_at IS '创建时间';
COMMENT ON COLUMN public.help_feedback.updated_at IS '更新时间';
COMMENT ON COLUMN public.help_feedback.deleted_at IS '删除时间';
ALTER TABLE ONLY public.help_feedback ADD CONSTRAINT help_feedback_pkey PRIMARY KEY (id);
CREATE INDEX help_feedback_user_id_idx ON public.help_feedback USING btree (user_id);
CREATE INDEX help_feedback_status_idx ON public.help_feedback USING btree (status);
