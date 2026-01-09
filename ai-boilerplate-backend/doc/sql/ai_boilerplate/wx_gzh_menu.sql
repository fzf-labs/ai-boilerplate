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
