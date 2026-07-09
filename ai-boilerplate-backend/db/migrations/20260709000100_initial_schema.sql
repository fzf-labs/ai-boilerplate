-- +goose Up

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Source: doc/sql/ai_boilerplate/activity.sql
CREATE TABLE public.activity (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    image_url character varying(500) NOT NULL,
    link_url character varying(500) NOT NULL,
    link_type character varying(32) NOT NULL,
    sort integer DEFAULT 0 NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.activity IS '内容-活动';
COMMENT ON COLUMN public.activity.id IS 'id';
COMMENT ON COLUMN public.activity.tenant_id IS '租户id';
COMMENT ON COLUMN public.activity.title IS '标题';
COMMENT ON COLUMN public.activity.image_url IS '图片URL';
COMMENT ON COLUMN public.activity.link_url IS '跳转链接';
COMMENT ON COLUMN public.activity.link_type IS '跳转类型';
COMMENT ON COLUMN public.activity.sort IS '排序';
COMMENT ON COLUMN public.activity.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.activity.start_time IS '开始时间';
COMMENT ON COLUMN public.activity.end_time IS '结束时间';
COMMENT ON COLUMN public.activity.created_at IS '创建时间';
COMMENT ON COLUMN public.activity.updated_at IS '更新时间';
COMMENT ON COLUMN public.activity.deleted_at IS '删除时间';
ALTER TABLE ONLY public.activity ADD CONSTRAINT activity_pkey PRIMARY KEY (id);
CREATE INDEX activity_tenant_id_idx ON public.activity USING btree (tenant_id);
CREATE INDEX activity_status_idx ON public.activity USING btree (status);
CREATE INDEX activity_sort_idx ON public.activity USING btree (sort);
CREATE INDEX activity_start_time_idx ON public.activity USING btree (start_time);
CREATE INDEX activity_end_time_idx ON public.activity USING btree (end_time);

-- Source: doc/sql/ai_boilerplate/ai_audio_record.sql
CREATE TABLE public.ai_audio_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    lyric text,
    image_url character varying(600),
    audio_url character varying(600),
    status integer NOT NULL,
    description text,
    prompt text,
    platform character varying(64) NOT NULL,
    model_id character varying(64) NOT NULL,
    model character varying(50) NOT NULL,
    generate_mode integer NOT NULL,
    tags character varying(600),
    duration double precision,
    public_status boolean DEFAULT false NOT NULL,
    task_id character varying(255),
    error_message character varying(1024),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_audio_record IS 'AI 音乐表';
COMMENT ON COLUMN public.ai_audio_record.id IS '编号';
COMMENT ON COLUMN public.ai_audio_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_audio_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_audio_record.title IS '音乐名称';
COMMENT ON COLUMN public.ai_audio_record.lyric IS '歌词';
COMMENT ON COLUMN public.ai_audio_record.image_url IS '图片地址';
COMMENT ON COLUMN public.ai_audio_record.audio_url IS '音频地址';
COMMENT ON COLUMN public.ai_audio_record.status IS '音频状态';
COMMENT ON COLUMN public.ai_audio_record.description IS '描述词';
COMMENT ON COLUMN public.ai_audio_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_audio_record.platform IS '模型平台';
COMMENT ON COLUMN public.ai_audio_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_audio_record.model IS '模型';
COMMENT ON COLUMN public.ai_audio_record.generate_mode IS '生成模式';
COMMENT ON COLUMN public.ai_audio_record.tags IS '风格标签';
COMMENT ON COLUMN public.ai_audio_record.duration IS '时长';
COMMENT ON COLUMN public.ai_audio_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_audio_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_audio_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_audio_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_audio_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_audio_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_audio_record ADD CONSTRAINT ai_audio_record_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX ai_music_record_pkey ON public.ai_audio_record USING btree (id);
CREATE INDEX idx_ai_music_record_platform ON public.ai_audio_record USING btree (platform);
CREATE INDEX idx_ai_music_record_tenant_id ON public.ai_audio_record USING btree (tenant_id);
CREATE INDEX idx_ai_music_record_user_id ON public.ai_audio_record USING btree (admin_id);

-- Source: doc/sql/ai_boilerplate/ai_chat_conversation.sql
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

-- Source: doc/sql/ai_boilerplate/ai_chat_message.sql
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

-- Source: doc/sql/ai_boilerplate/ai_image_record.sql
CREATE TABLE public.ai_image_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    prompt text NOT NULL,
    platform character varying(64) NOT NULL,
    model_id character varying(64),
    model character varying(64) NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    status integer NOT NULL,
    finish_time timestamp with time zone,
    error_message character varying(1024),
    public_status boolean DEFAULT false NOT NULL,
    pic_url character varying(2048),
    options jsonb,
    task_id character varying(255),
    buttons character varying(2048),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_image_record IS 'AI 绘画表';
COMMENT ON COLUMN public.ai_image_record.id IS '编号';
COMMENT ON COLUMN public.ai_image_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_image_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_image_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_image_record.platform IS '平台';
COMMENT ON COLUMN public.ai_image_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_image_record.model IS '模型';
COMMENT ON COLUMN public.ai_image_record.width IS '图片宽度';
COMMENT ON COLUMN public.ai_image_record.height IS '图片高度';
COMMENT ON COLUMN public.ai_image_record.status IS '绘画状态';
COMMENT ON COLUMN public.ai_image_record.finish_time IS '完成时间';
COMMENT ON COLUMN public.ai_image_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_image_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_image_record.pic_url IS '图片地址';
COMMENT ON COLUMN public.ai_image_record.options IS '绘制参数';
COMMENT ON COLUMN public.ai_image_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_image_record.buttons IS 'mj buttons 按钮';
COMMENT ON COLUMN public.ai_image_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_image_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_image_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_image_record ADD CONSTRAINT ai_image_record_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_image_record_platform ON public.ai_image_record USING btree (platform);
CREATE INDEX idx_ai_image_record_tenant_id ON public.ai_image_record USING btree (tenant_id);
CREATE INDEX idx_ai_image_record_user_id ON public.ai_image_record USING btree (admin_id);

-- Source: doc/sql/ai_boilerplate/ai_prompt.sql
CREATE TABLE public.ai_prompt (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64),
    name character varying(128) NOT NULL,
    "desc" character varying(255) NOT NULL,
    prompt text NOT NULL,
    sort integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_prompt IS 'AI 提示词';
COMMENT ON COLUMN public.ai_prompt.id IS 'Id';
COMMENT ON COLUMN public.ai_prompt.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_prompt.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_prompt.name IS '名称';
COMMENT ON COLUMN public.ai_prompt."desc" IS '描述';
COMMENT ON COLUMN public.ai_prompt.prompt IS '提示词';
COMMENT ON COLUMN public.ai_prompt.sort IS '排序';
COMMENT ON COLUMN public.ai_prompt.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_prompt.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_prompt.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_prompt ADD CONSTRAINT ai_prompt_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_prompt_tenant_id ON public.ai_prompt USING btree (tenant_id);
CREATE INDEX idx_ai_prompt_user_id ON public.ai_prompt USING btree (admin_id);

-- Source: doc/sql/ai_boilerplate/ai_provider_model.sql
CREATE TABLE public.ai_provider_model (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    platform_id character varying(64) NOT NULL,
    model_type character varying(64) NOT NULL,
    model_id character varying(64) NOT NULL,
    model_name character varying(64),
    model_config jsonb,
    sort integer NOT NULL,
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_provider_model IS 'AI 配置模型表';
COMMENT ON COLUMN public.ai_provider_model.id IS '编号';
COMMENT ON COLUMN public.ai_provider_model.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_model.platform_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_model.model_type IS '模型类型';
COMMENT ON COLUMN public.ai_provider_model.model_id IS '模型id';
COMMENT ON COLUMN public.ai_provider_model.model_name IS '模型名字';
COMMENT ON COLUMN public.ai_provider_model.model_config IS '配置';
COMMENT ON COLUMN public.ai_provider_model.sort IS '排序';
COMMENT ON COLUMN public.ai_provider_model.status IS '状态';
COMMENT ON COLUMN public.ai_provider_model.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_provider_model.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_provider_model.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_provider_model ADD CONSTRAINT ai_provider_model_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/ai_provider_platform.sql
CREATE TABLE public.ai_provider_platform (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    platform character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    api_url character varying(500),
    api_key character varying(500),
    doc_url character varying(500),
    sort integer,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_provider_platform IS 'AI 配置平台表';
COMMENT ON COLUMN public.ai_provider_platform.id IS '编号';
COMMENT ON COLUMN public.ai_provider_platform.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_provider_platform.platform IS '平台';
COMMENT ON COLUMN public.ai_provider_platform.name IS '名称';
COMMENT ON COLUMN public.ai_provider_platform.api_url IS 'API 地址';
COMMENT ON COLUMN public.ai_provider_platform.api_key IS 'API KEY';
COMMENT ON COLUMN public.ai_provider_platform.doc_url IS '文档地址';
COMMENT ON COLUMN public.ai_provider_platform.sort IS '排序';
COMMENT ON COLUMN public.ai_provider_platform.status IS '状态';
COMMENT ON COLUMN public.ai_provider_platform.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_provider_platform.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_provider_platform.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_provider_platform ADD CONSTRAINT ai_provider_platform_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX ai_conf_platform_pkey ON public.ai_provider_platform USING btree (id);

-- Source: doc/sql/ai_boilerplate/ai_video_record.sql
CREATE TABLE public.ai_video_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    prompt text NOT NULL,
    platform character varying(64) NOT NULL,
    model_id character varying(64),
    model character varying(64) NOT NULL,
    status integer NOT NULL,
    finish_time timestamp with time zone,
    error_message character varying(1024),
    public_status boolean DEFAULT false NOT NULL,
    video_url character varying(512),
    options jsonb,
    task_id character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_video_record IS 'AI 视频表';
COMMENT ON COLUMN public.ai_video_record.id IS '编号';
COMMENT ON COLUMN public.ai_video_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_video_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_video_record.prompt IS '提示词';
COMMENT ON COLUMN public.ai_video_record.platform IS '平台';
COMMENT ON COLUMN public.ai_video_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_video_record.model IS '模型';
COMMENT ON COLUMN public.ai_video_record.status IS '状态';
COMMENT ON COLUMN public.ai_video_record.finish_time IS '完成时间';
COMMENT ON COLUMN public.ai_video_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_video_record.public_status IS '是否发布';
COMMENT ON COLUMN public.ai_video_record.video_url IS '视频地址';
COMMENT ON COLUMN public.ai_video_record.options IS '绘制参数';
COMMENT ON COLUMN public.ai_video_record.task_id IS '任务编号';
COMMENT ON COLUMN public.ai_video_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_video_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_video_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_video_record ADD CONSTRAINT ai_video_record_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/ai_write_record.sql
CREATE TABLE public.ai_write_record (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    admin_id character varying(64) NOT NULL,
    type integer,
    platform character varying(255) NOT NULL,
    model_id character varying(64) NOT NULL,
    model character varying(255) NOT NULL,
    prompt text NOT NULL,
    generated_content text,
    original_content text,
    length integer,
    format integer,
    tone integer,
    language integer,
    error_message character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.ai_write_record IS 'AI 写作表';
COMMENT ON COLUMN public.ai_write_record.id IS '编号';
COMMENT ON COLUMN public.ai_write_record.tenant_id IS '租户编号';
COMMENT ON COLUMN public.ai_write_record.admin_id IS '用户编号';
COMMENT ON COLUMN public.ai_write_record.type IS '写作类型';
COMMENT ON COLUMN public.ai_write_record.platform IS '平台';
COMMENT ON COLUMN public.ai_write_record.model_id IS '模型编号';
COMMENT ON COLUMN public.ai_write_record.model IS '模型';
COMMENT ON COLUMN public.ai_write_record.prompt IS '生成内容提示';
COMMENT ON COLUMN public.ai_write_record.generated_content IS '生成的内容';
COMMENT ON COLUMN public.ai_write_record.original_content IS '原文';
COMMENT ON COLUMN public.ai_write_record.length IS '长度提示词';
COMMENT ON COLUMN public.ai_write_record.format IS '格式提示词';
COMMENT ON COLUMN public.ai_write_record.tone IS '语气提示词';
COMMENT ON COLUMN public.ai_write_record.language IS '语言提示词';
COMMENT ON COLUMN public.ai_write_record.error_message IS '错误信息';
COMMENT ON COLUMN public.ai_write_record.created_at IS '创建时间';
COMMENT ON COLUMN public.ai_write_record.updated_at IS '更新时间';
COMMENT ON COLUMN public.ai_write_record.deleted_at IS '删除时间';
ALTER TABLE ONLY public.ai_write_record ADD CONSTRAINT ai_write_record_pkey PRIMARY KEY (id);
CREATE INDEX idx_ai_write_record_platform ON public.ai_write_record USING btree (platform);
CREATE INDEX idx_ai_write_record_tenant_id ON public.ai_write_record USING btree (tenant_id);
CREATE INDEX idx_ai_write_record_user_id ON public.ai_write_record USING btree (admin_id);

-- Source: doc/sql/ai_boilerplate/article.sql
CREATE TABLE public.article (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title character varying(200) NOT NULL,
    summary character varying(500) DEFAULT ''::character varying NOT NULL,
    cover_image character varying(500) DEFAULT ''::character varying NOT NULL,
    content_markdown text NOT NULL,
    status smallint DEFAULT 0 NOT NULL,
    publish_time timestamp with time zone,
    tags text[] DEFAULT '{}'::text[] NOT NULL,
    is_recommend boolean DEFAULT false NOT NULL,
    is_hot boolean DEFAULT false NOT NULL,
    view_count bigint DEFAULT 0 NOT NULL,
    like_count bigint DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.article IS '内容-文章';
COMMENT ON COLUMN public.article.id IS 'id';
COMMENT ON COLUMN public.article.title IS '标题';
COMMENT ON COLUMN public.article.summary IS '摘要';
COMMENT ON COLUMN public.article.cover_image IS '封面图';
COMMENT ON COLUMN public.article.content_markdown IS 'Markdown 内容';
COMMENT ON COLUMN public.article.status IS '状态(-1下线,0草稿,1已发布)';
COMMENT ON COLUMN public.article.publish_time IS '发布时间';
COMMENT ON COLUMN public.article.tags IS '标签';
COMMENT ON COLUMN public.article.is_recommend IS '是否推荐';
COMMENT ON COLUMN public.article.is_hot IS '是否热门';
COMMENT ON COLUMN public.article.view_count IS '浏览量';
COMMENT ON COLUMN public.article.like_count IS '点赞数';
COMMENT ON COLUMN public.article.created_at IS '创建时间';
COMMENT ON COLUMN public.article.updated_at IS '更新时间';
COMMENT ON COLUMN public.article.deleted_at IS '删除时间';
ALTER TABLE ONLY public.article ADD CONSTRAINT article_pkey PRIMARY KEY (id);
CREATE INDEX article_status_idx ON public.article USING btree (status);
CREATE INDEX article_publish_time_idx ON public.article USING btree (publish_time);
CREATE INDEX article_title_idx ON public.article USING btree (title);


-- Source: doc/sql/ai_boilerplate/banner.sql
CREATE TABLE public.banner (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    image_url character varying(500) NOT NULL,
    link_url character varying(500) NOT NULL,
    link_type character varying(32) NOT NULL,
    position character varying(64) NOT NULL,
    platform character varying(32) NOT NULL,
    sort integer DEFAULT 0 NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.banner IS '通用-轮播图';
COMMENT ON COLUMN public.banner.id IS 'id';
COMMENT ON COLUMN public.banner.tenant_id IS '租户id';
COMMENT ON COLUMN public.banner.title IS '标题';
COMMENT ON COLUMN public.banner.image_url IS '图片URL';
COMMENT ON COLUMN public.banner.link_url IS '跳转链接';
COMMENT ON COLUMN public.banner.link_type IS '跳转类型';
COMMENT ON COLUMN public.banner.position IS '展示位置';
COMMENT ON COLUMN public.banner.platform IS '平台';
COMMENT ON COLUMN public.banner.sort IS '排序';
COMMENT ON COLUMN public.banner.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.banner.start_time IS '开始时间';
COMMENT ON COLUMN public.banner.end_time IS '结束时间';
COMMENT ON COLUMN public.banner.created_at IS '创建时间';
COMMENT ON COLUMN public.banner.updated_at IS '更新时间';
COMMENT ON COLUMN public.banner.deleted_at IS '删除时间';
ALTER TABLE ONLY public.banner ADD CONSTRAINT banner_pkey PRIMARY KEY (id);
CREATE INDEX banner_tenant_id_idx ON public.banner USING btree (tenant_id);
CREATE INDEX banner_status_idx ON public.banner USING btree (status);
CREATE INDEX banner_position_idx ON public.banner USING btree (position);
CREATE INDEX banner_platform_idx ON public.banner USING btree (platform);
CREATE INDEX banner_sort_idx ON public.banner USING btree (sort);

-- Source: doc/sql/ai_boilerplate/config_data.sql
CREATE TABLE public.config_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    key character varying(64) NOT NULL,
    value jsonb,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.config_data IS '配置管理';
COMMENT ON COLUMN public.config_data.id IS 'id';
COMMENT ON COLUMN public.config_data.name IS '名称';
COMMENT ON COLUMN public.config_data.key IS '健';
COMMENT ON COLUMN public.config_data.value IS '值';
COMMENT ON COLUMN public.config_data.remark IS '备注';
COMMENT ON COLUMN public.config_data.status IS '状态';
COMMENT ON COLUMN public.config_data.created_at IS '创建时间';
COMMENT ON COLUMN public.config_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.config_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.config_data ADD CONSTRAINT config_data_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX config_data_key_idx ON public.config_data USING btree (key);

-- Source: doc/sql/ai_boilerplate/device.sql
CREATE TABLE public.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sn character varying(128) NOT NULL,
    name character varying(225),
    "desc" character varying(225),
    brand character varying(225),
    model character varying(225),
    network character varying(225),
    imei character varying(225),
    cpu character varying(125),
    mac character varying(125),
    app_version character varying(125),
    android_version character varying(125),
    ram_size numeric,
    ddr_size numeric,
    certificate character varying(225),
    secure_key character varying(225),
    registry_time timestamp with time zone,
    push jsonb,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.device IS '设备表';
COMMENT ON COLUMN public.device.id IS 'ID';
COMMENT ON COLUMN public.device.sn IS '设备ID';
COMMENT ON COLUMN public.device.name IS '设备名称';
COMMENT ON COLUMN public.device."desc" IS '描述';
COMMENT ON COLUMN public.device.brand IS '设备品牌';
COMMENT ON COLUMN public.device.model IS '设备型号';
COMMENT ON COLUMN public.device.network IS '入网型号';
COMMENT ON COLUMN public.device.imei IS 'IMEI';
COMMENT ON COLUMN public.device.cpu IS 'cpu型号';
COMMENT ON COLUMN public.device.mac IS 'mac地址';
COMMENT ON COLUMN public.device.app_version IS 'app版本';
COMMENT ON COLUMN public.device.android_version IS '安卓版本';
COMMENT ON COLUMN public.device.ram_size IS 'RAM大小';
COMMENT ON COLUMN public.device.ddr_size IS 'DDR大小';
COMMENT ON COLUMN public.device.certificate IS '设备证书';
COMMENT ON COLUMN public.device.secure_key IS '设备密钥';
COMMENT ON COLUMN public.device.registry_time IS '激活时间';
COMMENT ON COLUMN public.device.push IS '推送';
COMMENT ON COLUMN public.device.status IS '状态';
COMMENT ON COLUMN public.device.created_at IS '创建时间';
COMMENT ON COLUMN public.device.updated_at IS '更新时间';
COMMENT ON COLUMN public.device.deleted_at IS '删除时间';
ALTER TABLE ONLY public.device ADD CONSTRAINT device_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX device_sn_idx ON public.device USING btree (sn);

-- Source: doc/sql/ai_boilerplate/dict_data.sql
CREATE TABLE public.dict_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    type character varying(100) NOT NULL,
    label character varying(100) NOT NULL,
    key character varying(100) NOT NULL,
    value character varying(100) NOT NULL,
    remark character varying(500),
    css_color character varying(100),
    css_class character varying(100),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.dict_data IS '字典数据表';
COMMENT ON COLUMN public.dict_data.id IS 'id';
COMMENT ON COLUMN public.dict_data.type IS '字典类型';
COMMENT ON COLUMN public.dict_data.label IS '字典标签';
COMMENT ON COLUMN public.dict_data.key IS '字典键';
COMMENT ON COLUMN public.dict_data.value IS '字典值';
COMMENT ON COLUMN public.dict_data.remark IS '备注';
COMMENT ON COLUMN public.dict_data.css_color IS 'css 颜色';
COMMENT ON COLUMN public.dict_data.css_class IS 'css 样式';
COMMENT ON COLUMN public.dict_data.status IS '状态（0正常 1停用）';
COMMENT ON COLUMN public.dict_data.created_at IS '创建时间';
COMMENT ON COLUMN public.dict_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.dict_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.dict_data ADD CONSTRAINT dict_data_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX dict_data_dict_type_key_key_idx ON public.dict_data USING btree (type, key);

-- Source: doc/sql/ai_boilerplate/dict_type.sql
CREATE TABLE public.dict_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(100) NOT NULL,
    type character varying(100) NOT NULL,
    status smallint NOT NULL,
    remark character varying(500),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.dict_type IS '字典类型表';
COMMENT ON COLUMN public.dict_type.id IS 'id';
COMMENT ON COLUMN public.dict_type.name IS '字典名称';
COMMENT ON COLUMN public.dict_type.type IS '字典类型';
COMMENT ON COLUMN public.dict_type.status IS '状态（1正常-1停用）';
COMMENT ON COLUMN public.dict_type.remark IS '备注';
COMMENT ON COLUMN public.dict_type.created_at IS '创建时间';
COMMENT ON COLUMN public.dict_type.updated_at IS '更新时间';
COMMENT ON COLUMN public.dict_type.deleted_at IS '删除时间';
ALTER TABLE ONLY public.dict_type ADD CONSTRAINT dict_type_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX dict_type_key_idx ON public.dict_type USING btree (type);

-- Source: doc/sql/ai_boilerplate/file_config.sql
CREATE TABLE public.file_config (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    storage character varying(32) NOT NULL,
    remark character varying(255),
    master boolean NOT NULL,
    config jsonb,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.file_config IS '文件配置表';
COMMENT ON COLUMN public.file_config.id IS '编号';
COMMENT ON COLUMN public.file_config.name IS '配置名';
COMMENT ON COLUMN public.file_config.storage IS '存储器';
COMMENT ON COLUMN public.file_config.remark IS '备注';
COMMENT ON COLUMN public.file_config.master IS '是否为主配置';
COMMENT ON COLUMN public.file_config.config IS '存储配置';
COMMENT ON COLUMN public.file_config.created_at IS '创建时间';
COMMENT ON COLUMN public.file_config.updated_at IS '更新时间';
COMMENT ON COLUMN public.file_config.deleted_at IS '删除时间';
ALTER TABLE ONLY public.file_config ADD CONSTRAINT file_config_pkey PRIMARY KEY (id);
CREATE INDEX file_config_master_idx ON public.file_config USING btree (master);

-- Source: doc/sql/ai_boilerplate/file_data.sql
CREATE TABLE public.file_data (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    config_id character varying(64),
    name character varying(255),
    path character varying(512) NOT NULL,
    url character varying(1024) NOT NULL,
    ext character varying(32),
    size integer NOT NULL,
    status integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.file_data IS '文件表';
COMMENT ON COLUMN public.file_data.id IS '文件编号';
COMMENT ON COLUMN public.file_data.config_id IS '配置编号';
COMMENT ON COLUMN public.file_data.name IS '文件名';
COMMENT ON COLUMN public.file_data.path IS '文件路径';
COMMENT ON COLUMN public.file_data.url IS '文件 URL';
COMMENT ON COLUMN public.file_data.ext IS '文件类型';
COMMENT ON COLUMN public.file_data.size IS '文件大小';
COMMENT ON COLUMN public.file_data.status IS '状态（-1失败,1未知,2 成功）';
COMMENT ON COLUMN public.file_data.created_at IS '创建时间';
COMMENT ON COLUMN public.file_data.updated_at IS '更新时间';
COMMENT ON COLUMN public.file_data.deleted_at IS '删除时间';
ALTER TABLE ONLY public.file_data ADD CONSTRAINT file_data_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/help_category.sql
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

-- Source: doc/sql/ai_boilerplate/help_faq.sql
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

-- Source: doc/sql/ai_boilerplate/help_feedback.sql
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

-- Source: doc/sql/ai_boilerplate/mail_account.sql
CREATE TABLE public.mail_account (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    mail character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    host character varying(255) NOT NULL,
    port integer NOT NULL,
    ssl_enable boolean,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_account IS '邮箱账号表';
COMMENT ON COLUMN public.mail_account.id IS 'id';
COMMENT ON COLUMN public.mail_account.mail IS '邮箱';
COMMENT ON COLUMN public.mail_account.username IS '用户名';
COMMENT ON COLUMN public.mail_account.password IS '密码';
COMMENT ON COLUMN public.mail_account.host IS 'SMTP 服务器域名';
COMMENT ON COLUMN public.mail_account.port IS 'SMTP 服务器端口';
COMMENT ON COLUMN public.mail_account.ssl_enable IS '是否开启 SSL';
COMMENT ON COLUMN public.mail_account.remark IS '备注';
COMMENT ON COLUMN public.mail_account.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.mail_account.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_account.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_account.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_account ADD CONSTRAINT mail_account_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX mail_account_mail_idx ON public.mail_account USING btree (mail);
CREATE INDEX mail_account_status_idx ON public.mail_account USING btree (status);

-- Source: doc/sql/ai_boilerplate/mail_log.sql
CREATE TABLE public.mail_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    account_id character varying(64) NOT NULL,
    from_mail character varying(255) NOT NULL,
    to_mail character varying(255) NOT NULL,
    template_id character varying(64) NOT NULL,
    template_code character varying(64) NOT NULL,
    template_nickname character varying(255),
    template_title character varying(255) NOT NULL,
    template_content text NOT NULL,
    template_params character varying(255) NOT NULL,
    send_status integer NOT NULL,
    send_time timestamp with time zone NOT NULL,
    send_message_id character varying(255),
    send_exception character varying(4096),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_log IS '邮件日志表';
COMMENT ON COLUMN public.mail_log.id IS 'id';
COMMENT ON COLUMN public.mail_log.account_id IS '邮箱账号编号';
COMMENT ON COLUMN public.mail_log.from_mail IS '发送邮箱地址';
COMMENT ON COLUMN public.mail_log.to_mail IS '接收邮箱地址';
COMMENT ON COLUMN public.mail_log.template_id IS '模板编号';
COMMENT ON COLUMN public.mail_log.template_code IS '模板编码';
COMMENT ON COLUMN public.mail_log.template_nickname IS '模版发送人名称';
COMMENT ON COLUMN public.mail_log.template_title IS '邮件标题';
COMMENT ON COLUMN public.mail_log.template_content IS '邮件内容';
COMMENT ON COLUMN public.mail_log.template_params IS '邮件参数';
COMMENT ON COLUMN public.mail_log.send_status IS '发送状态';
COMMENT ON COLUMN public.mail_log.send_time IS '发送时间';
COMMENT ON COLUMN public.mail_log.send_message_id IS '发送返回的消息 ID';
COMMENT ON COLUMN public.mail_log.send_exception IS '发送异常';
COMMENT ON COLUMN public.mail_log.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_log.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_log.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_log ADD CONSTRAINT mail_log_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/mail_template.sql
CREATE TABLE public.mail_template (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    code character varying(64) NOT NULL,
    account_id character varying(64) NOT NULL,
    nickname character varying(255),
    title character varying(255) NOT NULL,
    content text NOT NULL,
    params jsonb,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_template IS '邮件模版表';
COMMENT ON COLUMN public.mail_template.id IS 'id';
COMMENT ON COLUMN public.mail_template.name IS '模板名称';
COMMENT ON COLUMN public.mail_template.code IS '模板编码';
COMMENT ON COLUMN public.mail_template.account_id IS '发送的邮箱账号编号';
COMMENT ON COLUMN public.mail_template.nickname IS '发送人名称';
COMMENT ON COLUMN public.mail_template.title IS '模板标题';
COMMENT ON COLUMN public.mail_template.content IS '模板内容';
COMMENT ON COLUMN public.mail_template.params IS '参数数组';
COMMENT ON COLUMN public.mail_template.remark IS '备注';
COMMENT ON COLUMN public.mail_template.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.mail_template.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_template.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_template.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_template ADD CONSTRAINT mail_template_pkey PRIMARY KEY (id);
CREATE INDEX mail_template_account_id_idx ON public.mail_template USING btree (account_id);

-- Source: doc/sql/ai_boilerplate/mall_activation_code.sql
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

-- Source: doc/sql/ai_boilerplate/mall_order.sql
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

-- Source: doc/sql/ai_boilerplate/mall_payment_record.sql
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

-- Source: doc/sql/ai_boilerplate/mall_product.sql
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

-- Source: doc/sql/ai_boilerplate/membership.sql
CREATE TABLE public.membership (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    type character varying(20) NOT NULL,
    description character varying(255),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership IS '会员类型配置表';
COMMENT ON COLUMN public.membership.id IS 'id';
COMMENT ON COLUMN public.membership.name IS '会员类型名称';
COMMENT ON COLUMN public.membership.type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.membership.description IS '会员类型描述';
COMMENT ON COLUMN public.membership.sort IS '排序';
COMMENT ON COLUMN public.membership.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership.created_at IS '创建时间';
COMMENT ON COLUMN public.membership.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership ADD CONSTRAINT membership_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX membership_code_idx ON public.membership USING btree (type);
CREATE INDEX membership_sort_idx ON public.membership USING btree (sort);

-- Source: doc/sql/ai_boilerplate/membership_benefit.sql
CREATE TABLE public.membership_benefit (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    membership_type character varying(20) NOT NULL,
    benefit_key character varying(100) NOT NULL,
    benefit_value character varying(100),
    benefit_num character varying(100),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership_benefit IS '会员权益配置表';
COMMENT ON COLUMN public.membership_benefit.id IS 'id';
COMMENT ON COLUMN public.membership_benefit.membership_type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.membership_benefit.benefit_key IS '权益标识(关联membership_benefit_type.benefit_key)';
COMMENT ON COLUMN public.membership_benefit.benefit_value IS '权益值';
COMMENT ON COLUMN public.membership_benefit.benefit_num IS '权益次数';
COMMENT ON COLUMN public.membership_benefit.sort IS '排序';
COMMENT ON COLUMN public.membership_benefit.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership_benefit.created_at IS '创建时间';
COMMENT ON COLUMN public.membership_benefit.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership_benefit.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership_benefit ADD CONSTRAINT membership_benefit_pkey PRIMARY KEY (id);
CREATE INDEX membership_benefit_membership_code_idx ON public.membership_benefit USING btree (membership_type);
CREATE INDEX membership_benefit_sort_order_idx ON public.membership_benefit USING btree (sort);
CREATE UNIQUE INDEX membership_benefit_membership_type_benefit_key_idx ON public.membership_benefit USING btree (membership_type, benefit_key);

-- Source: doc/sql/ai_boilerplate/membership_benefit_type.sql
CREATE TABLE public.membership_benefit_type (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    benefit_key character varying(100) NOT NULL,
    benefit_name character varying(255) NOT NULL,
    benefit_icon character varying(500),
    benefit_desc character varying(500),
    sort integer DEFAULT 0,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.membership_benefit_type IS '会员权益类型表';
COMMENT ON COLUMN public.membership_benefit_type.id IS 'id';
COMMENT ON COLUMN public.membership_benefit_type.benefit_key IS '权益标识';
COMMENT ON COLUMN public.membership_benefit_type.benefit_name IS '权益名称';
COMMENT ON COLUMN public.membership_benefit_type.benefit_icon IS '权益图标';
COMMENT ON COLUMN public.membership_benefit_type.benefit_desc IS '权益描述';
COMMENT ON COLUMN public.membership_benefit_type.sort IS '排序';
COMMENT ON COLUMN public.membership_benefit_type.status IS '状态(-1禁用,1启用)';
COMMENT ON COLUMN public.membership_benefit_type.created_at IS '创建时间';
COMMENT ON COLUMN public.membership_benefit_type.updated_at IS '更新时间';
COMMENT ON COLUMN public.membership_benefit_type.deleted_at IS '删除时间';
ALTER TABLE ONLY public.membership_benefit_type ADD CONSTRAINT membership_benefit_type_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX membership_benefit_type_key_idx ON public.membership_benefit_type USING btree (benefit_key);
CREATE INDEX membership_benefit_type_sort_idx ON public.membership_benefit_type USING btree (sort);

-- Source: doc/sql/ai_boilerplate/self_app.sql
CREATE TABLE public.self_app (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    package_name character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.self_app IS '自应用信息表';
COMMENT ON COLUMN public.self_app.id IS 'ID';
COMMENT ON COLUMN public.self_app.package_name IS '包名';
COMMENT ON COLUMN public.self_app.name IS '应用名称';
COMMENT ON COLUMN public.self_app.description IS '应用描述';
COMMENT ON COLUMN public.self_app.status IS '状态(-1禁用 1启用)';
COMMENT ON COLUMN public.self_app.created_at IS '创建时间';
COMMENT ON COLUMN public.self_app.updated_at IS '更新时间';
COMMENT ON COLUMN public.self_app.deleted_at IS '删除时间';
ALTER TABLE ONLY public.self_app ADD CONSTRAINT self_app_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX self_app_package_name_idx ON public.self_app USING btree (package_name);

-- Source: doc/sql/ai_boilerplate/self_app_release.sql
CREATE TABLE public.self_app_release (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    channel character varying(32) NOT NULL,
    package_name character varying(255) NOT NULL,
    build_num integer NOT NULL,
    version character varying(32),
    update_type integer DEFAULT 2 NOT NULL,
    title character varying(255) NOT NULL,
    changelog text,
    package_url character varying(500) NOT NULL,
    package_size numeric DEFAULT 0,
    package_md5 character varying(32),
    min_os_version character varying(32),
    publish_time timestamp with time zone NOT NULL,
    gray_strategy integer DEFAULT 1 NOT NULL,
    gray_sns jsonb,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.self_app_release IS '自应用版本发布表';
COMMENT ON COLUMN public.self_app_release.id IS 'ID';
COMMENT ON COLUMN public.self_app_release.channel IS '发布渠道';
COMMENT ON COLUMN public.self_app_release.package_name IS '包名';
COMMENT ON COLUMN public.self_app_release.build_num IS 'build值';
COMMENT ON COLUMN public.self_app_release.version IS '版本号';
COMMENT ON COLUMN public.self_app_release.update_type IS '更新类型(1强制 2提示 3静默)';
COMMENT ON COLUMN public.self_app_release.title IS '更新标题';
COMMENT ON COLUMN public.self_app_release.changelog IS '更新日志';
COMMENT ON COLUMN public.self_app_release.package_url IS '安装包地址';
COMMENT ON COLUMN public.self_app_release.package_size IS '安装包大小';
COMMENT ON COLUMN public.self_app_release.package_md5 IS '安装包MD5';
COMMENT ON COLUMN public.self_app_release.min_os_version IS '最低系统版本';
COMMENT ON COLUMN public.self_app_release.publish_time IS '发布时间';
COMMENT ON COLUMN public.self_app_release.gray_strategy IS '灰度策略(1全量 2自定义设备)';
COMMENT ON COLUMN public.self_app_release.gray_sns IS '灰度设备';
COMMENT ON COLUMN public.self_app_release.status IS '状态(-1禁用 1启用)';
COMMENT ON COLUMN public.self_app_release.created_at IS '创建时间';
COMMENT ON COLUMN public.self_app_release.updated_at IS '更新时间';
COMMENT ON COLUMN public.self_app_release.deleted_at IS '删除时间';
ALTER TABLE ONLY public.self_app_release ADD CONSTRAINT self_app_release_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX self_app_release_package_name_channel_build_num_idx ON public.self_app_release USING btree (package_name, channel, build_num);
CREATE INDEX self_app_release_package_name_channel_gray_strategy_idx ON public.self_app_release USING btree (package_name, channel, gray_strategy);

-- Source: doc/sql/ai_boilerplate/sensitive_word.sql
CREATE TABLE public.sensitive_word (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    word character varying NOT NULL,
    lab character varying,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sensitive_word IS '敏感词';
COMMENT ON COLUMN public.sensitive_word.id IS 'id';
COMMENT ON COLUMN public.sensitive_word.word IS '敏感词';
COMMENT ON COLUMN public.sensitive_word.lab IS '标签';
COMMENT ON COLUMN public.sensitive_word.created_at IS '创建时间';
COMMENT ON COLUMN public.sensitive_word.updated_at IS '更新时间';
COMMENT ON COLUMN public.sensitive_word.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sensitive_word ADD CONSTRAINT sensitive_word_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX sensitive_word_word_idx ON public.sensitive_word USING btree (word);

-- Source: doc/sql/ai_boilerplate/sms_channel.sql
CREATE TABLE public.sms_channel (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    operator character varying(64) NOT NULL,
    remark character varying(255),
    api_key character varying(128) NOT NULL,
    api_secret character varying(128),
    callback_url character varying(255),
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sms_channel IS '短信渠道';
COMMENT ON COLUMN public.sms_channel.id IS 'id';
COMMENT ON COLUMN public.sms_channel.name IS '渠道名称';
COMMENT ON COLUMN public.sms_channel.operator IS '运营商';
COMMENT ON COLUMN public.sms_channel.remark IS '备注';
COMMENT ON COLUMN public.sms_channel.api_key IS '短信 API 的账号';
COMMENT ON COLUMN public.sms_channel.api_secret IS '短信 API 的秘钥';
COMMENT ON COLUMN public.sms_channel.callback_url IS '短信发送回调 URL';
COMMENT ON COLUMN public.sms_channel.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sms_channel.created_at IS '创建时间';
COMMENT ON COLUMN public.sms_channel.updated_at IS '更新时间';
COMMENT ON COLUMN public.sms_channel.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sms_channel ADD CONSTRAINT sms_channel_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/sms_log.sql
CREATE TABLE public.sms_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sms_channel_id character varying(64) NOT NULL,
    sms_template_id character varying(64) NOT NULL,
    sms_params_content character varying(255) NOT NULL,
    mobile character varying(11) NOT NULL,
    user_id character varying(64),
    send_status character varying(64) NOT NULL,
    send_time timestamp with time zone NOT NULL,
    receive_status character varying(64) NOT NULL,
    receive_time timestamp with time zone,
    api_send_code character varying(64),
    api_send_msg character varying(255),
    api_request_id character varying(255),
    api_serial_no character varying(255),
    api_receive_code character varying(64),
    api_receive_msg character varying(255),
    created_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.sms_log IS '短信日志';
COMMENT ON COLUMN public.sms_log.id IS '编号';
COMMENT ON COLUMN public.sms_log.sms_channel_id IS '短信渠道编号';
COMMENT ON COLUMN public.sms_log.sms_template_id IS '模板编号';
COMMENT ON COLUMN public.sms_log.sms_params_content IS '短信参数内容';
COMMENT ON COLUMN public.sms_log.mobile IS '手机号';
COMMENT ON COLUMN public.sms_log.user_id IS '用户id';
COMMENT ON COLUMN public.sms_log.send_status IS '发送状态';
COMMENT ON COLUMN public.sms_log.send_time IS '发送时间';
COMMENT ON COLUMN public.sms_log.receive_status IS '接收状态';
COMMENT ON COLUMN public.sms_log.receive_time IS '接收时间';
COMMENT ON COLUMN public.sms_log.api_send_code IS '短信 API 发送结果的编码';
COMMENT ON COLUMN public.sms_log.api_send_msg IS '短信 API 发送失败的提示';
COMMENT ON COLUMN public.sms_log.api_request_id IS '短信 API 发送返回的唯一请求 ID';
COMMENT ON COLUMN public.sms_log.api_serial_no IS '短信 API 发送返回的序号';
COMMENT ON COLUMN public.sms_log.api_receive_code IS 'API 接收结果的编码';
COMMENT ON COLUMN public.sms_log.api_receive_msg IS 'API 接收结果的说明';
COMMENT ON COLUMN public.sms_log.created_at IS '创建时间';
ALTER TABLE ONLY public.sms_log ADD CONSTRAINT sms_log_pkey PRIMARY KEY (id);
CREATE INDEX sms_log_mobile_idx ON public.sms_log USING btree (mobile);

-- Source: doc/sql/ai_boilerplate/sms_template.sql
CREATE TABLE public.sms_template (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sms_channel_id character varying(64) NOT NULL,
    template_type smallint NOT NULL,
    template_code character varying(64) NOT NULL,
    template_name character varying(64) NOT NULL,
    template_content character varying(500) NOT NULL,
    template_params jsonb,
    remark character varying(255),
    api_template_id character varying(64) NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sms_template IS '短信模板';
COMMENT ON COLUMN public.sms_template.id IS 'id';
COMMENT ON COLUMN public.sms_template.sms_channel_id IS '短信渠道编号';
COMMENT ON COLUMN public.sms_template.template_type IS '模板类型';
COMMENT ON COLUMN public.sms_template.template_code IS '模板编码';
COMMENT ON COLUMN public.sms_template.template_name IS '模板名称';
COMMENT ON COLUMN public.sms_template.template_content IS '模板内容';
COMMENT ON COLUMN public.sms_template.template_params IS '模板参数';
COMMENT ON COLUMN public.sms_template.remark IS '备注';
COMMENT ON COLUMN public.sms_template.api_template_id IS '短信供应商的模板编号';
COMMENT ON COLUMN public.sms_template.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sms_template.created_at IS '创建时间';
COMMENT ON COLUMN public.sms_template.updated_at IS '更新时间';
COMMENT ON COLUMN public.sms_template.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sms_template ADD CONSTRAINT sms_template_pkey PRIMARY KEY (id);
CREATE INDEX sms_template_sms_channel_id_idx ON public.sms_template USING btree (sms_channel_id);

-- Source: doc/sql/ai_boilerplate/sys_admin.sql
CREATE TABLE public.sys_admin (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying,
    username character varying NOT NULL,
    password character varying NOT NULL,
    nickname character varying,
    avatar character varying,
    email character varying,
    sex smallint,
    mobile character varying,
    role_id character varying,
    dept_id character varying,
    post_id character varying,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_admin IS '系统-管理员';
COMMENT ON COLUMN public.sys_admin.id IS 'id';
COMMENT ON COLUMN public.sys_admin.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_admin.username IS '用户名';
COMMENT ON COLUMN public.sys_admin.password IS '密码';
COMMENT ON COLUMN public.sys_admin.nickname IS '昵称';
COMMENT ON COLUMN public.sys_admin.avatar IS '头像';
COMMENT ON COLUMN public.sys_admin.email IS '邮件';
COMMENT ON COLUMN public.sys_admin.sex IS '性别';
COMMENT ON COLUMN public.sys_admin.mobile IS '手机号';
COMMENT ON COLUMN public.sys_admin.role_id IS '角色Id';
COMMENT ON COLUMN public.sys_admin.dept_id IS '部门';
COMMENT ON COLUMN public.sys_admin.post_id IS '岗位';
COMMENT ON COLUMN public.sys_admin.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_admin.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_admin.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_admin.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_admin ADD CONSTRAINT sys_admin_pkey PRIMARY KEY (id);
CREATE INDEX sys_admin_dept_id_idx ON public.sys_admin USING btree (dept_id);
CREATE INDEX sys_admin_post_id_idx ON public.sys_admin USING btree (post_id);
CREATE INDEX sys_admin_role_id_idx ON public.sys_admin USING btree (role_id);
CREATE UNIQUE INDEX sys_admin_username_idx ON public.sys_admin USING btree (username);

-- Source: doc/sql/ai_boilerplate/sys_api.sql
CREATE TABLE public.sys_api (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    permission_id uuid NOT NULL,
    method character varying(32) NOT NULL,
    path character varying(255) NOT NULL,
    "desc" character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_api IS '系统-接口';
COMMENT ON COLUMN public.sys_api.id IS 'id';
COMMENT ON COLUMN public.sys_api.permission_id IS '权限Id';
COMMENT ON COLUMN public.sys_api.method IS '方法';
COMMENT ON COLUMN public.sys_api.path IS '路径';
COMMENT ON COLUMN public.sys_api."desc" IS '描述';
COMMENT ON COLUMN public.sys_api.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_api.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_api.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_api ADD CONSTRAINT sys_api_pkey PRIMARY KEY (id);
CREATE INDEX sys_api_permission_id_idx ON public.sys_api USING btree (permission_id);

-- Source: doc/sql/ai_boilerplate/sys_dept.sql
CREATE TABLE public.sys_dept (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid character varying(64),
    name character varying(64) NOT NULL,
    admin_id character varying(64),
    status smallint DEFAULT 1 NOT NULL,
    sort bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    tenant_id character varying(64)
);
COMMENT ON TABLE public.sys_dept IS '系统-部门';
COMMENT ON COLUMN public.sys_dept.id IS '编号';
COMMENT ON COLUMN public.sys_dept.pid IS '父级id';
COMMENT ON COLUMN public.sys_dept.name IS '部门简称';
COMMENT ON COLUMN public.sys_dept.admin_id IS '负责人Id';
COMMENT ON COLUMN public.sys_dept.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_dept.sort IS '排序值';
COMMENT ON COLUMN public.sys_dept.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_dept.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_dept.deleted_at IS '删除时间';
COMMENT ON COLUMN public.sys_dept.tenant_id IS '租户Id';
ALTER TABLE ONLY public.sys_dept ADD CONSTRAINT sys_dept_pkey PRIMARY KEY (id);
CREATE INDEX sys_dept_pid_idx ON public.sys_dept USING btree (pid);

-- Source: doc/sql/ai_boilerplate/sys_menu.sql
CREATE TABLE public.sys_menu (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    pid character varying(64) NOT NULL,
    name character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    path character varying(100) NOT NULL,
    permission character varying(64),
    icon character varying(64),
    component character varying(100),
    component_name character varying(100),
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_menu IS '菜单和权限规则表';
COMMENT ON COLUMN public.sys_menu.id IS 'id';
COMMENT ON COLUMN public.sys_menu.pid IS '上级菜单';
COMMENT ON COLUMN public.sys_menu.name IS '菜单名称';
COMMENT ON COLUMN public.sys_menu.type IS '菜单类型(dir,menu,button)';
COMMENT ON COLUMN public.sys_menu.path IS '路由路径';
COMMENT ON COLUMN public.sys_menu.permission IS '权限标识';
COMMENT ON COLUMN public.sys_menu.icon IS '图标';
COMMENT ON COLUMN public.sys_menu.component IS '组件路径';
COMMENT ON COLUMN public.sys_menu.component_name IS '组件名';
COMMENT ON COLUMN public.sys_menu.sort IS '权重(排序)';
COMMENT ON COLUMN public.sys_menu.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_menu.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_menu.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_menu.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_menu ADD CONSTRAINT sys_menu_pkey PRIMARY KEY (id);
CREATE INDEX sys_menu_pid_idx ON public.sys_menu USING btree (pid);

-- Source: doc/sql/ai_boilerplate/sys_notice.sql
CREATE TABLE public.sys_notice (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    content character varying(200) NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_notice IS '系统-公告';
COMMENT ON COLUMN public.sys_notice.id IS 'id';
COMMENT ON COLUMN public.sys_notice.tenant_id IS '租户id';
COMMENT ON COLUMN public.sys_notice.type IS '类型';
COMMENT ON COLUMN public.sys_notice.title IS '标题';
COMMENT ON COLUMN public.sys_notice.content IS '内容';
COMMENT ON COLUMN public.sys_notice.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_notice.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_notice.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_notice.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_notice ADD CONSTRAINT sys_notice_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/sys_notify_message.sql
CREATE TABLE public.sys_notify_message (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    subject character varying(200) NOT NULL,
    content character varying(200) NOT NULL,
    sender character varying(64) NOT NULL,
    receiver character varying(64) NOT NULL,
    send_time character varying(64) NOT NULL,
    read_time character varying(64) DEFAULT ''::character varying,
    extend jsonb
);
COMMENT ON TABLE public.sys_notify_message IS '系统-通知消息';
COMMENT ON COLUMN public.sys_notify_message.id IS 'id';
COMMENT ON COLUMN public.sys_notify_message.tenant_id IS '租户id';
COMMENT ON COLUMN public.sys_notify_message.type IS '消息类型';
COMMENT ON COLUMN public.sys_notify_message.subject IS '主题';
COMMENT ON COLUMN public.sys_notify_message.content IS '内容';
COMMENT ON COLUMN public.sys_notify_message.sender IS '发送人';
COMMENT ON COLUMN public.sys_notify_message.receiver IS '接收人';
COMMENT ON COLUMN public.sys_notify_message.send_time IS '发送时间';
COMMENT ON COLUMN public.sys_notify_message.read_time IS '阅读时间';
COMMENT ON COLUMN public.sys_notify_message.extend IS '扩展';
ALTER TABLE ONLY public.sys_notify_message ADD CONSTRAINT sys_notify_message_pkey PRIMARY KEY (id);
CREATE INDEX sys_notify_message_receiver_read_time_idx ON public.sys_notify_message USING btree (receiver, read_time);

-- Source: doc/sql/ai_boilerplate/sys_operate_log.sql
CREATE TABLE public.sys_operate_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying NOT NULL,
    trace_id character varying,
    admin_id uuid NOT NULL,
    ip character varying(32) NOT NULL,
    uri character varying(200) NOT NULL,
    useragent character varying(255),
    header json,
    req json,
    resp json,
    created_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.sys_operate_log IS '系统-日志';
COMMENT ON COLUMN public.sys_operate_log.id IS 'id';
COMMENT ON COLUMN public.sys_operate_log.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_operate_log.trace_id IS '链路Id';
COMMENT ON COLUMN public.sys_operate_log.admin_id IS '管理员ID';
COMMENT ON COLUMN public.sys_operate_log.ip IS 'ip';
COMMENT ON COLUMN public.sys_operate_log.uri IS '请求路径';
COMMENT ON COLUMN public.sys_operate_log.useragent IS '浏览器标识';
COMMENT ON COLUMN public.sys_operate_log.header IS 'header';
COMMENT ON COLUMN public.sys_operate_log.req IS '请求数据';
COMMENT ON COLUMN public.sys_operate_log.resp IS '响应数据';
COMMENT ON COLUMN public.sys_operate_log.created_at IS '创建时间';
ALTER TABLE ONLY public.sys_operate_log ADD CONSTRAINT sys_operate_log_pkey PRIMARY KEY (id);
CREATE INDEX sys_operate_log_admin_id_idx ON public.sys_operate_log USING btree (admin_id);
CREATE INDEX sys_operate_log_tenant_id_idx ON public.sys_operate_log USING btree (tenant_id);
CREATE INDEX sys_operate_log_trace_id_idx ON public.sys_operate_log USING btree (trace_id);

-- Source: doc/sql/ai_boilerplate/sys_post.sql
CREATE TABLE public.sys_post (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64),
    name character varying(50) NOT NULL,
    code character varying(32),
    remark character varying(255),
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_post IS '系统-工作岗位';
COMMENT ON COLUMN public.sys_post.id IS '编号';
COMMENT ON COLUMN public.sys_post.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_post.name IS '岗位名称';
COMMENT ON COLUMN public.sys_post.code IS '岗位编码';
COMMENT ON COLUMN public.sys_post.remark IS '备注';
COMMENT ON COLUMN public.sys_post.sort IS '排序值';
COMMENT ON COLUMN public.sys_post.status IS '0=禁用 1=开启 ';
COMMENT ON COLUMN public.sys_post.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_post.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_post.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_post ADD CONSTRAINT sys_post_pkey PRIMARY KEY (id);
CREATE INDEX sys_post_tenant_id_idx ON public.sys_post USING btree (tenant_id);

-- Source: doc/sql/ai_boilerplate/sys_role.sql
CREATE TABLE public.sys_role (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    name character varying(50) NOT NULL,
    remark character varying(200),
    "dataScope" character varying(64) NOT NULL,
    "menuIds" jsonb,
    sort bigint NOT NULL,
    status smallint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_role IS '系统-角色';
COMMENT ON COLUMN public.sys_role.id IS '编号';
COMMENT ON COLUMN public.sys_role.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_role.name IS '名称';
COMMENT ON COLUMN public.sys_role.remark IS '备注';
COMMENT ON COLUMN public.sys_role."dataScope" IS '数据范围';
COMMENT ON COLUMN public.sys_role."menuIds" IS '菜单';
COMMENT ON COLUMN public.sys_role.sort IS '排序值';
COMMENT ON COLUMN public.sys_role.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_role.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_role.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_role.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_role ADD CONSTRAINT sys_role_pkey PRIMARY KEY (id);
CREATE INDEX sys_role_tenant_id_idx ON public.sys_role USING btree (tenant_id);

-- Source: doc/sql/ai_boilerplate/sys_tenant.sql
CREATE TABLE public.sys_tenant (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(50) NOT NULL,
    remark character varying(200),
    admin_id character varying(50),
    expire_time timestamp with time zone,
    "menuIds" jsonb,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sys_tenant IS '系统-租户';
COMMENT ON COLUMN public.sys_tenant.id IS 'id';
COMMENT ON COLUMN public.sys_tenant.name IS '名称';
COMMENT ON COLUMN public.sys_tenant.remark IS '备注';
COMMENT ON COLUMN public.sys_tenant.admin_id IS '租户管理员Id';
COMMENT ON COLUMN public.sys_tenant.expire_time IS '过期时间';
COMMENT ON COLUMN public.sys_tenant."menuIds" IS '菜单';
COMMENT ON COLUMN public.sys_tenant.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sys_tenant.created_at IS '创建时间';
COMMENT ON COLUMN public.sys_tenant.updated_at IS '更新时间';
COMMENT ON COLUMN public.sys_tenant.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sys_tenant ADD CONSTRAINT sys_tenant_pkey PRIMARY KEY (id);
CREATE INDEX sys_tenant_name_idx ON public.sys_tenant USING btree (name);

-- Source: doc/sql/ai_boilerplate/user.sql
CREATE TABLE public."user" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    phone character varying NOT NULL,
    password character varying,
    salt character varying NOT NULL,
    nickname character varying,
    gender integer DEFAULT 0,
    avatar character varying,
    profile character varying,
    other jsonb,
    wx_gzh_user_id character varying(64),
    wx_gzh_xcx_id character varying(64),
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public."user" IS '用户表';
COMMENT ON COLUMN public."user".id IS 'id';
COMMENT ON COLUMN public."user".phone IS '手机';
COMMENT ON COLUMN public."user".password IS '密码';
COMMENT ON COLUMN public."user".salt IS '盐值';
COMMENT ON COLUMN public."user".nickname IS '昵称';
COMMENT ON COLUMN public."user".gender IS '性别（0未知 1男 2女）';
COMMENT ON COLUMN public."user".avatar IS '头像';
COMMENT ON COLUMN public."user".profile IS '简介';
COMMENT ON COLUMN public."user".other IS '其他';
COMMENT ON COLUMN public."user".wx_gzh_user_id IS '公众号用户Id';
COMMENT ON COLUMN public."user".wx_gzh_xcx_id IS '小程序用户Id';
COMMENT ON COLUMN public."user".status IS '状态';
COMMENT ON COLUMN public."user".created_at IS '创建时间';
COMMENT ON COLUMN public."user".updated_at IS '更新时间';
COMMENT ON COLUMN public."user".deleted_at IS '删除时间';
ALTER TABLE ONLY public."user" ADD CONSTRAINT user_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX user_phone_idx ON public."user" USING btree (phone);
CREATE INDEX user_wx_gzh_user_id_idx ON public."user" USING btree (wx_gzh_user_id);

-- Source: doc/sql/ai_boilerplate/user_bind_device.sql
CREATE TABLE public.user_bind_device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id character varying(64) NOT NULL,
    sn character varying(64) NOT NULL,
    identity character varying(64) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.user_bind_device IS '用户绑定设备表';
COMMENT ON COLUMN public.user_bind_device.id IS 'id';
COMMENT ON COLUMN public.user_bind_device.user_id IS '用户Id';
COMMENT ON COLUMN public.user_bind_device.sn IS 'sn';
COMMENT ON COLUMN public.user_bind_device.identity IS '身份';
COMMENT ON COLUMN public.user_bind_device.created_at IS '创建时间';
COMMENT ON COLUMN public.user_bind_device.updated_at IS '更新时间';
ALTER TABLE ONLY public.user_bind_device ADD CONSTRAINT user_bind_device_pkey PRIMARY KEY (id);
CREATE INDEX user_bind_device_sn_idx ON public.user_bind_device USING btree (sn);
CREATE UNIQUE INDEX user_bind_device_user_id_sn_idx ON public.user_bind_device USING btree (user_id, sn);

-- Source: doc/sql/ai_boilerplate/user_membership.sql
CREATE TABLE public.user_membership (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    membership_type character varying(20) NOT NULL,
    expired_at timestamp with time zone,
    auto_renew integer DEFAULT 0,
    auto_renew_days integer,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_membership IS '用户会员关系表';
COMMENT ON COLUMN public.user_membership.id IS 'id';
COMMENT ON COLUMN public.user_membership.user_id IS '用户ID';
COMMENT ON COLUMN public.user_membership.membership_type IS '会员类型编码(normal,vip,svip)';
COMMENT ON COLUMN public.user_membership.expired_at IS '到期时间(普通会员为NULL,表示永不过期)';
COMMENT ON COLUMN public.user_membership.auto_renew IS '是否自动续费(0否,1是)';
COMMENT ON COLUMN public.user_membership.auto_renew_days IS '自动续费天数';
COMMENT ON COLUMN public.user_membership.status IS '状态(-1禁用,1正常)';
COMMENT ON COLUMN public.user_membership.created_at IS '创建时间';
COMMENT ON COLUMN public.user_membership.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_membership.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_membership ADD CONSTRAINT user_membership_pkey PRIMARY KEY (id);
CREATE INDEX user_membership_expired_at_idx ON public.user_membership USING btree (expired_at);
CREATE INDEX user_membership_membership_type_code_idx ON public.user_membership USING btree (membership_type);
CREATE UNIQUE INDEX user_membership_user_id_idx ON public.user_membership USING btree (user_id);

-- Source: doc/sql/ai_boilerplate/user_membership_change.sql
CREATE TABLE public.user_membership_change (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    source_type character varying(20) NOT NULL,
    source_id character varying(128) NOT NULL,
    before_membership_type character varying(20),
    after_membership_type character varying(20),
    before_expired_at timestamp with time zone,
    after_expired_at timestamp with time zone,
    duration_days integer,
    remark character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_membership_change IS '用户会员变更记录表';
COMMENT ON COLUMN public.user_membership_change.id IS 'id';
COMMENT ON COLUMN public.user_membership_change.user_id IS '用户ID';
COMMENT ON COLUMN public.user_membership_change.source_type IS '来源类型(order,activation_code,admin)';
COMMENT ON COLUMN public.user_membership_change.source_id IS '来源ID(订单ID/激活码)';
COMMENT ON COLUMN public.user_membership_change.before_membership_type IS '变更前会员类型';
COMMENT ON COLUMN public.user_membership_change.after_membership_type IS '变更后会员类型';
COMMENT ON COLUMN public.user_membership_change.before_expired_at IS '变更前到期时间';
COMMENT ON COLUMN public.user_membership_change.after_expired_at IS '变更后到期时间';
COMMENT ON COLUMN public.user_membership_change.duration_days IS '变更时长(天)';
COMMENT ON COLUMN public.user_membership_change.remark IS '备注';
COMMENT ON COLUMN public.user_membership_change.created_at IS '创建时间';
COMMENT ON COLUMN public.user_membership_change.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_membership_change.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_membership_change ADD CONSTRAINT user_membership_change_pkey PRIMARY KEY (id);
CREATE INDEX user_membership_change_user_id_created_at_idx ON public.user_membership_change USING btree (user_id, created_at);
CREATE INDEX user_membership_change_source_type_idx ON public.user_membership_change USING btree (source_type);

-- Source: doc/sql/ai_boilerplate/user_message.sql
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

-- Source: doc/sql/ai_boilerplate/user_notification_settings.sql
CREATE TABLE public.user_notification_settings (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    notification_preferences jsonb DEFAULT '{}'::jsonb NOT NULL,
    dnd_start_time character varying(5),
    dnd_end_time character varying(5),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_notification_settings IS '用户通知设置表';
COMMENT ON COLUMN public.user_notification_settings.id IS 'id';
COMMENT ON COLUMN public.user_notification_settings.user_id IS '用户ID';
COMMENT ON COLUMN public.user_notification_settings.notification_preferences IS '通知偏好(JSON)';
COMMENT ON COLUMN public.user_notification_settings.dnd_start_time IS '勿扰开始时间';
COMMENT ON COLUMN public.user_notification_settings.dnd_end_time IS '勿扰结束时间';
COMMENT ON COLUMN public.user_notification_settings.created_at IS '创建时间';
COMMENT ON COLUMN public.user_notification_settings.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_notification_settings.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_notification_settings ADD CONSTRAINT user_notification_settings_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX user_notification_settings_user_id_idx ON public.user_notification_settings USING btree (user_id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_account.sql
CREATE TABLE public.wx_gzh_account (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    name character varying(100) NOT NULL,
    account character varying(100) NOT NULL,
    app_id character varying(100) NOT NULL,
    app_secret character varying(100) NOT NULL,
    url character varying(1000),
    token character varying(64),
    encoding_aes_key character varying(64),
    qr_code_url character varying(1000),
    remark character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_account IS '公众号账号表';
COMMENT ON COLUMN public.wx_gzh_account.id IS '编号';
COMMENT ON COLUMN public.wx_gzh_account.tenant_id IS '租户编号';
COMMENT ON COLUMN public.wx_gzh_account.name IS '公众号名称';
COMMENT ON COLUMN public.wx_gzh_account.account IS '公众号账号';
COMMENT ON COLUMN public.wx_gzh_account.app_id IS '公众号appid';
COMMENT ON COLUMN public.wx_gzh_account.app_secret IS '公众号密钥';
COMMENT ON COLUMN public.wx_gzh_account.url IS '公众号url';
COMMENT ON COLUMN public.wx_gzh_account.token IS '公众号token';
COMMENT ON COLUMN public.wx_gzh_account.encoding_aes_key IS '加密密钥';
COMMENT ON COLUMN public.wx_gzh_account.qr_code_url IS '二维码图片URL';
COMMENT ON COLUMN public.wx_gzh_account.remark IS '备注';
COMMENT ON COLUMN public.wx_gzh_account.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_account.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_account.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_account ADD CONSTRAINT wx_gzh_account_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_gzh_account_app_id_idx ON public.wx_gzh_account USING btree (app_id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_auto_reply.sql
CREATE TABLE public.wx_gzh_auto_reply (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    type integer NOT NULL,
    request_keyword character varying(255),
    request_keyword_match integer,
    response_message_type character varying(32) NOT NULL,
    response_content text,
    response_media_id character varying(1000),
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_auto_reply IS '公众号消息自动回复表';
COMMENT ON COLUMN public.wx_gzh_auto_reply.id IS '主键';
COMMENT ON COLUMN public.wx_gzh_auto_reply.app_id IS '公众号 appId';
COMMENT ON COLUMN public.wx_gzh_auto_reply.type IS '回复类型(关键词回复,收到消息回复,被关注回复)';
COMMENT ON COLUMN public.wx_gzh_auto_reply.request_keyword IS '请求的关键字';
COMMENT ON COLUMN public.wx_gzh_auto_reply.request_keyword_match IS '请求的关键字匹配类型';
COMMENT ON COLUMN public.wx_gzh_auto_reply.response_message_type IS '回复的消息类型';
COMMENT ON COLUMN public.wx_gzh_auto_reply.response_content IS '回复的消息内容';
COMMENT ON COLUMN public.wx_gzh_auto_reply.response_media_id IS '回复的媒体文件 id';
COMMENT ON COLUMN public.wx_gzh_auto_reply.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.wx_gzh_auto_reply.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_auto_reply.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_auto_reply.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_auto_reply ADD CONSTRAINT wx_gzh_auto_reply_pkey PRIMARY KEY (id);
CREATE INDEX wx_gzh_auto_reply_app_id_type_idx ON public.wx_gzh_auto_reply USING btree (app_id, type);

-- Source: doc/sql/ai_boilerplate/wx_gzh_material.sql
CREATE TABLE public.wx_gzh_material (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    type character varying(32) NOT NULL,
    media_id character varying(128) NOT NULL,
    tags jsonb,
    update_time timestamp with time zone NOT NULL,
    name character varying(255),
    url character varying(1000),
    cover_url character varying(1000),
    description character varying(1000),
    newcat character varying(1000),
    newsubcat character varying(1000),
    vid character varying(255),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_material IS '公众号素材表';
COMMENT ON COLUMN public.wx_gzh_material.id IS '主键';
COMMENT ON COLUMN public.wx_gzh_material.app_id IS '公众号 appId';
COMMENT ON COLUMN public.wx_gzh_material.type IS '微信-素材的类型，图片（image）、视频（video）、语音 （voice）';
COMMENT ON COLUMN public.wx_gzh_material.media_id IS '微信-消息ID';
COMMENT ON COLUMN public.wx_gzh_material.tags IS '微信-标签';
COMMENT ON COLUMN public.wx_gzh_material.update_time IS '微信-更新日期';
COMMENT ON COLUMN public.wx_gzh_material.name IS '微信-图片、语音、视频素材的名字';
COMMENT ON COLUMN public.wx_gzh_material.url IS '微信-图片、语音、视频素材URL(图片,视频是微信的地址,音频是服务端的地址)';
COMMENT ON COLUMN public.wx_gzh_material.cover_url IS '微信-视频封面 URL';
COMMENT ON COLUMN public.wx_gzh_material.description IS '微信-视频描述';
COMMENT ON COLUMN public.wx_gzh_material.newcat IS '微信-视频分类';
COMMENT ON COLUMN public.wx_gzh_material.newsubcat IS '微信-视频子分类';
COMMENT ON COLUMN public.wx_gzh_material.vid IS '微信-视频 ID';
COMMENT ON COLUMN public.wx_gzh_material.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_material.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_material.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_material ADD CONSTRAINT wx_gzh_material_pkey PRIMARY KEY (id);
CREATE INDEX wx_gzh_material_app_id_type_idx ON public.wx_gzh_material USING btree (app_id, type);
CREATE INDEX wx_gzh_material_media_id_idx ON public.wx_gzh_material USING btree (media_id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_menu.sql
CREATE TABLE public.wx_gzh_menu (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    is_menu_open integer,
    selfmenu_info jsonb,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_menu IS '公众号菜单表';
COMMENT ON COLUMN public.wx_gzh_menu.id IS '主键';
COMMENT ON COLUMN public.wx_gzh_menu.app_id IS '微信公众号 appid';
COMMENT ON COLUMN public.wx_gzh_menu.is_menu_open IS '菜单是否开启，0代表未开启，1代表开启';
COMMENT ON COLUMN public.wx_gzh_menu.selfmenu_info IS '菜单信息';
COMMENT ON COLUMN public.wx_gzh_menu.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_menu.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_menu.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_menu ADD CONSTRAINT wx_gzh_menu_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_gzh_menu_app_id_idx ON public.wx_gzh_menu USING btree (app_id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_message.sql
CREATE TABLE public.wx_gzh_message (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    msg_id character varying(100),
    user_id character varying(100) NOT NULL,
    openid character varying(100) NOT NULL,
    message_type character varying(32) NOT NULL,
    send_from integer NOT NULL,
    content character varying(1024),
    media_id character varying(128),
    media_url character varying(1024),
    recognition character varying(1024),
    format character varying(16),
    title character varying(128),
    description character varying(256),
    thumb_media_id character varying(128),
    thumb_media_url character varying(1024),
    url character varying(500),
    location_x double precision,
    location_y double precision,
    scale double precision,
    label character varying(128),
    articles character varying(1024),
    music_url character varying(1024),
    hq_music_url character varying(1024),
    event character varying(64),
    event_key character varying(64),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_message IS '公众号消息表 ';
COMMENT ON COLUMN public.wx_gzh_message.id IS '主键';
COMMENT ON COLUMN public.wx_gzh_message.app_id IS '公众号 appId';
COMMENT ON COLUMN public.wx_gzh_message.msg_id IS '微信公众号的消息编号';
COMMENT ON COLUMN public.wx_gzh_message.user_id IS '公众号粉丝的编号';
COMMENT ON COLUMN public.wx_gzh_message.openid IS '公众号粉丝标志';
COMMENT ON COLUMN public.wx_gzh_message.message_type IS '消息类型';
COMMENT ON COLUMN public.wx_gzh_message.send_from IS '消息来源';
COMMENT ON COLUMN public.wx_gzh_message.content IS '消息内容';
COMMENT ON COLUMN public.wx_gzh_message.media_id IS '媒体文件 id';
COMMENT ON COLUMN public.wx_gzh_message.media_url IS '媒体文件 URL';
COMMENT ON COLUMN public.wx_gzh_message.recognition IS '语音识别后文本';
COMMENT ON COLUMN public.wx_gzh_message.format IS '语音格式';
COMMENT ON COLUMN public.wx_gzh_message.title IS '标题';
COMMENT ON COLUMN public.wx_gzh_message.description IS '描述';
COMMENT ON COLUMN public.wx_gzh_message.thumb_media_id IS '缩略图的媒体 id';
COMMENT ON COLUMN public.wx_gzh_message.thumb_media_url IS '缩略图的媒体 URL';
COMMENT ON COLUMN public.wx_gzh_message.url IS '点击图文消息跳转链接';
COMMENT ON COLUMN public.wx_gzh_message.location_x IS '地理位置维度';
COMMENT ON COLUMN public.wx_gzh_message.location_y IS '地理位置经度';
COMMENT ON COLUMN public.wx_gzh_message.scale IS '地图缩放大小';
COMMENT ON COLUMN public.wx_gzh_message.label IS '详细地址';
COMMENT ON COLUMN public.wx_gzh_message.articles IS '图文消息数组';
COMMENT ON COLUMN public.wx_gzh_message.music_url IS '音乐链接';
COMMENT ON COLUMN public.wx_gzh_message.hq_music_url IS '高质量音乐链接';
COMMENT ON COLUMN public.wx_gzh_message.event IS '事件类型';
COMMENT ON COLUMN public.wx_gzh_message.event_key IS '事件 Key';
COMMENT ON COLUMN public.wx_gzh_message.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_message.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_message.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_message ADD CONSTRAINT wx_gzh_message_pkey PRIMARY KEY (id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_tag.sql
CREATE TABLE public.wx_gzh_tag (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    tag_id integer,
    name character varying(32),
    count integer DEFAULT 0,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_tag IS '公众号标签表';
COMMENT ON COLUMN public.wx_gzh_tag.id IS '主键';
COMMENT ON COLUMN public.wx_gzh_tag.app_id IS '公众号 appId';
COMMENT ON COLUMN public.wx_gzh_tag.tag_id IS '公众号标签 id';
COMMENT ON COLUMN public.wx_gzh_tag.name IS '标签名称';
COMMENT ON COLUMN public.wx_gzh_tag.count IS '粉丝数量';
COMMENT ON COLUMN public.wx_gzh_tag.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_tag.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_tag.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_tag ADD CONSTRAINT wx_gzh_tag_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_gzh_tag_app_id_tag_id_idx ON public.wx_gzh_tag USING btree (app_id, tag_id);

-- Source: doc/sql/ai_boilerplate/wx_gzh_user.sql
CREATE TABLE public.wx_gzh_user (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    openid character varying(100) NOT NULL,
    unionid character varying(100),
    subscribe_status integer,
    subscribe_time integer,
    nickname character varying(64),
    avatar_url character varying(1024),
    language character varying(30),
    country character varying(30),
    province character varying(30),
    city character varying(30),
    tag_ids character varying(255),
    remark character varying(128),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_gzh_user IS '公众号粉丝表';
COMMENT ON COLUMN public.wx_gzh_user.id IS '编号';
COMMENT ON COLUMN public.wx_gzh_user.app_id IS '微信公众号 appid';
COMMENT ON COLUMN public.wx_gzh_user.openid IS '用户标识';
COMMENT ON COLUMN public.wx_gzh_user.unionid IS '微信生态唯一标识';
COMMENT ON COLUMN public.wx_gzh_user.subscribe_status IS '关注状态';
COMMENT ON COLUMN public.wx_gzh_user.subscribe_time IS '关注时间';
COMMENT ON COLUMN public.wx_gzh_user.nickname IS '昵称';
COMMENT ON COLUMN public.wx_gzh_user.avatar_url IS '头像地址';
COMMENT ON COLUMN public.wx_gzh_user.language IS '语言';
COMMENT ON COLUMN public.wx_gzh_user.country IS '国家';
COMMENT ON COLUMN public.wx_gzh_user.province IS '省份';
COMMENT ON COLUMN public.wx_gzh_user.city IS '城市';
COMMENT ON COLUMN public.wx_gzh_user.tag_ids IS '标签编号数组';
COMMENT ON COLUMN public.wx_gzh_user.remark IS '备注';
COMMENT ON COLUMN public.wx_gzh_user.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_gzh_user.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_gzh_user.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_gzh_user ADD CONSTRAINT wx_gzh_user_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_gzh_user_app_id_openid_idx ON public.wx_gzh_user USING btree (app_id, openid);

-- Source: doc/sql/ai_boilerplate/wx_xcx_user.sql
CREATE TABLE public.wx_xcx_user (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    app_id character varying(128) NOT NULL,
    openid character varying(100) NOT NULL,
    unionid character varying(100),
    nickname character varying(64),
    avatar_url character varying(1024),
    language character varying(30),
    country character varying(30),
    province character varying(30),
    city character varying(30),
    remark character varying(128),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.wx_xcx_user IS '小程序用户表';
COMMENT ON COLUMN public.wx_xcx_user.id IS '编号';
COMMENT ON COLUMN public.wx_xcx_user.app_id IS '微信小程序 appid';
COMMENT ON COLUMN public.wx_xcx_user.openid IS '用户标识';
COMMENT ON COLUMN public.wx_xcx_user.unionid IS '微信生态唯一标识';
COMMENT ON COLUMN public.wx_xcx_user.nickname IS '昵称';
COMMENT ON COLUMN public.wx_xcx_user.avatar_url IS '头像地址';
COMMENT ON COLUMN public.wx_xcx_user.language IS '语言';
COMMENT ON COLUMN public.wx_xcx_user.country IS '国家';
COMMENT ON COLUMN public.wx_xcx_user.province IS '省份';
COMMENT ON COLUMN public.wx_xcx_user.city IS '城市';
COMMENT ON COLUMN public.wx_xcx_user.remark IS '备注';
COMMENT ON COLUMN public.wx_xcx_user.created_at IS '创建时间';
COMMENT ON COLUMN public.wx_xcx_user.updated_at IS '更新时间';
COMMENT ON COLUMN public.wx_xcx_user.deleted_at IS '删除时间';
ALTER TABLE ONLY public.wx_xcx_user ADD CONSTRAINT wx_xcx_user_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX wx_xcx_user_app_id_openid_idx ON public.wx_xcx_user USING btree (app_id, openid);

-- +goose Down


DROP TABLE IF EXISTS public.wx_xcx_user CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_user CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_tag CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_message CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_menu CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_material CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_auto_reply CASCADE;
DROP TABLE IF EXISTS public.wx_gzh_account CASCADE;
DROP TABLE IF EXISTS public.user_notification_settings CASCADE;
DROP TABLE IF EXISTS public.user_message CASCADE;
DROP TABLE IF EXISTS public.user_membership_change CASCADE;
DROP TABLE IF EXISTS public.user_membership CASCADE;
DROP TABLE IF EXISTS public.user_bind_device CASCADE;
DROP TABLE IF EXISTS public."user" CASCADE;
DROP TABLE IF EXISTS public.sys_tenant CASCADE;
DROP TABLE IF EXISTS public.sys_role CASCADE;
DROP TABLE IF EXISTS public.sys_post CASCADE;
DROP TABLE IF EXISTS public.sys_operate_log CASCADE;
DROP TABLE IF EXISTS public.sys_notify_message CASCADE;
DROP TABLE IF EXISTS public.sys_notice CASCADE;
DROP TABLE IF EXISTS public.sys_menu CASCADE;
DROP TABLE IF EXISTS public.sys_dept CASCADE;
DROP TABLE IF EXISTS public.sys_api CASCADE;
DROP TABLE IF EXISTS public.sys_admin CASCADE;
DROP TABLE IF EXISTS public.sms_template CASCADE;
DROP TABLE IF EXISTS public.sms_log CASCADE;
DROP TABLE IF EXISTS public.sms_channel CASCADE;
DROP TABLE IF EXISTS public.sensitive_word CASCADE;
DROP TABLE IF EXISTS public.self_app_release CASCADE;
DROP TABLE IF EXISTS public.self_app CASCADE;
DROP TABLE IF EXISTS public.membership_benefit_type CASCADE;
DROP TABLE IF EXISTS public.membership_benefit CASCADE;
DROP TABLE IF EXISTS public.membership CASCADE;
DROP TABLE IF EXISTS public.mall_product CASCADE;
DROP TABLE IF EXISTS public.mall_payment_record CASCADE;
DROP TABLE IF EXISTS public.mall_order CASCADE;
DROP TABLE IF EXISTS public.mall_activation_code CASCADE;
DROP TABLE IF EXISTS public.mail_template CASCADE;
DROP TABLE IF EXISTS public.mail_log CASCADE;
DROP TABLE IF EXISTS public.mail_account CASCADE;
DROP TABLE IF EXISTS public.help_feedback CASCADE;
DROP TABLE IF EXISTS public.help_faq CASCADE;
DROP TABLE IF EXISTS public.help_category CASCADE;
DROP TABLE IF EXISTS public.file_data CASCADE;
DROP TABLE IF EXISTS public.file_config CASCADE;
DROP TABLE IF EXISTS public.dict_type CASCADE;
DROP TABLE IF EXISTS public.dict_data CASCADE;
DROP TABLE IF EXISTS public.device CASCADE;
DROP TABLE IF EXISTS public.config_data CASCADE;
DROP TABLE IF EXISTS public.banner CASCADE;
DROP TABLE IF EXISTS public.article CASCADE;
DROP TABLE IF EXISTS public.ai_write_record CASCADE;
DROP TABLE IF EXISTS public.ai_video_record CASCADE;
DROP TABLE IF EXISTS public.ai_provider_platform CASCADE;
DROP TABLE IF EXISTS public.ai_provider_model CASCADE;
DROP TABLE IF EXISTS public.ai_prompt CASCADE;
DROP TABLE IF EXISTS public.ai_image_record CASCADE;
DROP TABLE IF EXISTS public.ai_chat_message CASCADE;
DROP TABLE IF EXISTS public.ai_chat_conversation CASCADE;
DROP TABLE IF EXISTS public.ai_audio_record CASCADE;
DROP TABLE IF EXISTS public.activity CASCADE;
