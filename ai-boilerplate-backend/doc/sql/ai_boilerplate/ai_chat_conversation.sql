CREATE TABLE public.ai_chat_conversation (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    title character varying(256) NOT NULL,
    pinned boolean NOT NULL,
    pinned_time timestamp with time zone,
    prompt_setting jsonb,
    model_setting jsonb,
    knowledge_setting jsonb,
    mcp_setting jsonb,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_chat_conversation IS 'AI 聊天对话表';
COMMENT ON COLUMN public.ai_chat_conversation.id IS '对话编号';
COMMENT ON COLUMN public.ai_chat_conversation.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_chat_conversation.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_chat_conversation.title IS '对话标题';
COMMENT ON COLUMN public.ai_chat_conversation.pinned IS '是否置顶';
COMMENT ON COLUMN public.ai_chat_conversation.pinned_time IS '置顶时间';
COMMENT ON COLUMN public.ai_chat_conversation.prompt_setting IS '提示词设置';
COMMENT ON COLUMN public.ai_chat_conversation.model_setting IS '模型设置';
COMMENT ON COLUMN public.ai_chat_conversation.knowledge_setting IS '知识库设置';
COMMENT ON COLUMN public.ai_chat_conversation.mcp_setting IS 'mcp设置';
COMMENT ON COLUMN public.ai_chat_conversation.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_chat_conversation.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_chat_conversation.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_chat_conversation ADD CONSTRAINT ai_chat_conversation_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_chat_conversation_tenant_id ON public.ai_chat_conversation USING btree (tenant_id);
CREATE INDEX idx_ai_chat_conversation_user_id ON public.ai_chat_conversation USING btree (admin_id);
