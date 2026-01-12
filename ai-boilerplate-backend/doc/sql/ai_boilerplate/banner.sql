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
