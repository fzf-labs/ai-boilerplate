CREATE TABLE public.user_message (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    message_id uuid NOT NULL,
    user_id uuid NOT NULL,
    category character varying(32) NOT NULL,
    title character varying(200) NOT NULL,
    summary character varying(200),
    cover_url character varying(500),
    content text NOT NULL,
    link_url character varying(500),
    audience_type character varying(32) NOT NULL,
    audience_value jsonb,
    sent_at timestamp with time zone NOT NULL,
    read_at timestamp with time zone,
    admin_id uuid,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_message IS 'App-用户消息';
COMMENT ON COLUMN public.user_message.id IS 'id';
COMMENT ON COLUMN public.user_message.message_id IS '消息批次id';
COMMENT ON COLUMN public.user_message.user_id IS '用户id';
COMMENT ON COLUMN public.user_message.category IS '消息分类(transaction/system/service)';
COMMENT ON COLUMN public.user_message.title IS '标题';
COMMENT ON COLUMN public.user_message.summary IS '摘要';
COMMENT ON COLUMN public.user_message.cover_url IS '封面图';
COMMENT ON COLUMN public.user_message.content IS '内容';
COMMENT ON COLUMN public.user_message.link_url IS '跳转链接';
COMMENT ON COLUMN public.user_message.audience_type IS '投放范围(all/segment/users)';
COMMENT ON COLUMN public.user_message.audience_value IS '投放条件/用户列表';
COMMENT ON COLUMN public.user_message.sent_at IS '发送时间';
COMMENT ON COLUMN public.user_message.read_at IS '阅读时间';
COMMENT ON COLUMN public.user_message.admin_id IS '创建人';
COMMENT ON COLUMN public.user_message.created_at IS '创建时间';
COMMENT ON COLUMN public.user_message.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_message.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_message ADD CONSTRAINT user_message_pkey PRIMARY KEY (id);
CREATE INDEX user_message_user_category_read_sent_idx ON public.user_message USING btree (user_id, category, read_at, sent_at);
CREATE INDEX user_message_message_id_idx ON public.user_message USING btree (message_id);
