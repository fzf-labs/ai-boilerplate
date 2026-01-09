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
