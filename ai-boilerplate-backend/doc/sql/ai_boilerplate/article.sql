CREATE TABLE public.article (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title character varying(200) NOT NULL,
    summary character varying(500) DEFAULT ''::character varying NOT NULL,
    cover_image character varying(500) DEFAULT ''::character varying NOT NULL,
    content_markdown text NOT NULL,
    status smallint DEFAULT 0 NOT NULL,
    publish_time timestamp with time zone,
    tags text[] DEFAULT '{}'::text[] NOT NULL,
    is_recommend boolean DEFAULT false NOT NULL,
    is_hot boolean DEFAULT false NOT NULL,
    view_count bigint DEFAULT 0 NOT NULL,
    like_count bigint DEFAULT 0 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.article IS '内容-文章';
COMMENT ON COLUMN public.article.id IS 'id';
COMMENT ON COLUMN public.article.title IS '标题';
COMMENT ON COLUMN public.article.summary IS '摘要';
COMMENT ON COLUMN public.article.cover_image IS '封面图';
COMMENT ON COLUMN public.article.content_markdown IS 'Markdown 内容';
COMMENT ON COLUMN public.article.status IS '状态(-1下线,0草稿,1已发布)';
COMMENT ON COLUMN public.article.publish_time IS '发布时间';
COMMENT ON COLUMN public.article.tags IS '标签';
COMMENT ON COLUMN public.article.is_recommend IS '是否推荐';
COMMENT ON COLUMN public.article.is_hot IS '是否热门';
COMMENT ON COLUMN public.article.view_count IS '浏览量';
COMMENT ON COLUMN public.article.like_count IS '点赞数';
COMMENT ON COLUMN public.article.created_at IS '创建时间';
COMMENT ON COLUMN public.article.updated_at IS '更新时间';
COMMENT ON COLUMN public.article.deleted_at IS '删除时间';
ALTER TABLE ONLY public.article ADD CONSTRAINT article_pkey PRIMARY KEY (id);
CREATE INDEX article_status_idx ON public.article USING btree (status);
CREATE INDEX article_publish_time_idx ON public.article USING btree (publish_time);
CREATE INDEX article_title_idx ON public.article USING btree (title);

