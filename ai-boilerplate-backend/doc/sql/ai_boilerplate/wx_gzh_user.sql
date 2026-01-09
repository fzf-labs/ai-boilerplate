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
