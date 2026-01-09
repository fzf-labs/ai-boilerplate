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
