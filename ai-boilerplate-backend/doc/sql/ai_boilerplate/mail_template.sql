CREATE TABLE public.mail_template (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(64) NOT NULL,
    code character varying(64) NOT NULL,
    account_id character varying(64) NOT NULL,
    nickname character varying(255),
    title character varying(255) NOT NULL,
    content text NOT NULL,
    params jsonb,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_template IS '邮件模版表';
COMMENT ON COLUMN public.mail_template.id IS 'id';
COMMENT ON COLUMN public.mail_template.name IS '模板名称';
COMMENT ON COLUMN public.mail_template.code IS '模板编码';
COMMENT ON COLUMN public.mail_template.account_id IS '发送的邮箱账号编号';
COMMENT ON COLUMN public.mail_template.nickname IS '发送人名称';
COMMENT ON COLUMN public.mail_template.title IS '模板标题';
COMMENT ON COLUMN public.mail_template.content IS '模板内容';
COMMENT ON COLUMN public.mail_template.params IS '参数数组';
COMMENT ON COLUMN public.mail_template.remark IS '备注';
COMMENT ON COLUMN public.mail_template.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.mail_template.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_template.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_template.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_template ADD CONSTRAINT mail_template_pkey PRIMARY KEY (id);
CREATE INDEX mail_template_account_id_idx ON public.mail_template USING btree (account_id);
