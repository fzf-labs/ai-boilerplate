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
