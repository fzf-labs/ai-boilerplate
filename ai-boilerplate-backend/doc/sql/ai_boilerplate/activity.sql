CREATE TABLE public.activity (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    title character varying(200) NOT NULL,
    image_url character varying(500) NOT NULL,
    link_url character varying(500) NOT NULL,
    link_type character varying(32) NOT NULL,
    sort integer DEFAULT 0 NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.activity IS '内容-活动';
COMMENT ON COLUMN public.activity.id IS 'id';
COMMENT ON COLUMN public.activity.tenant_id IS '租户id';
COMMENT ON COLUMN public.activity.title IS '标题';
COMMENT ON COLUMN public.activity.image_url IS '图片URL';
COMMENT ON COLUMN public.activity.link_url IS '跳转链接';
COMMENT ON COLUMN public.activity.link_type IS '跳转类型';
COMMENT ON COLUMN public.activity.sort IS '排序';
COMMENT ON COLUMN public.activity.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.activity.start_time IS '开始时间';
COMMENT ON COLUMN public.activity.end_time IS '结束时间';
COMMENT ON COLUMN public.activity.created_at IS '创建时间';
COMMENT ON COLUMN public.activity.updated_at IS '更新时间';
COMMENT ON COLUMN public.activity.deleted_at IS '删除时间';
ALTER TABLE ONLY public.activity ADD CONSTRAINT activity_pkey PRIMARY KEY (id);
CREATE INDEX activity_tenant_id_idx ON public.activity USING btree (tenant_id);
CREATE INDEX activity_status_idx ON public.activity USING btree (status);
CREATE INDEX activity_sort_idx ON public.activity USING btree (sort);
CREATE INDEX activity_start_time_idx ON public.activity USING btree (start_time);
CREATE INDEX activity_end_time_idx ON public.activity USING btree (end_time);
