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
