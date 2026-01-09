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
