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
