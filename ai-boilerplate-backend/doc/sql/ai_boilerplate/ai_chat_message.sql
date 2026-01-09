CREATE TABLE public.ai_chat_message (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64),
    conversation_id character varying(64) NOT NULL,
    reply_id character varying(64),
    admin_id character varying(64) NOT NULL,
    role_id character varying(64),
    type character varying(16) NOT NULL,
    model character varying(32) NOT NULL,
    model_id character varying(64) NOT NULL,
    content text NOT NULL,
    use_context boolean DEFAULT false NOT NULL,
    segment_ids character varying(2048),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_chat_message IS 'AI 聊天消息表';
COMMENT ON COLUMN public.ai_chat_message.id IS '消息编号';
COMMENT ON COLUMN public.ai_chat_message.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_chat_message.conversation_id IS '对话编号';
COMMENT ON COLUMN public.ai_chat_message.reply_id IS '回复编号';
COMMENT ON COLUMN public.ai_chat_message.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_chat_message.role_id IS '角色编号';
COMMENT ON COLUMN public.ai_chat_message.type IS '消息类型';
COMMENT ON COLUMN public.ai_chat_message.model IS '模型标识';
COMMENT ON COLUMN public.ai_chat_message.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_chat_message.content IS '消息内容';
COMMENT ON COLUMN public.ai_chat_message.use_context IS '是否携带上下文';
COMMENT ON COLUMN public.ai_chat_message.segment_ids IS '段落编号数组';
COMMENT ON COLUMN public.ai_chat_message.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_chat_message.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_chat_message.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_chat_message ADD CONSTRAINT ai_chat_message_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_chat_message_conversation_id ON public.ai_chat_message USING btree (conversation_id);
CREATE INDEX idx_ai_chat_message_tenant_id ON public.ai_chat_message USING btree (tenant_id);
CREATE INDEX idx_ai_chat_message_user_id ON public.ai_chat_message USING btree (admin_id);
