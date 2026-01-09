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
