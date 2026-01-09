CREATE TABLE public.mail_account (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    mail character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    host character varying(255) NOT NULL,
    port integer NOT NULL,
    ssl_enable boolean,
    remark character varying(255),
    status integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_account IS '邮箱账号表';
COMMENT ON COLUMN public.mail_account.id IS 'id';
COMMENT ON COLUMN public.mail_account.mail IS '邮箱';
COMMENT ON COLUMN public.mail_account.username IS '用户名';
COMMENT ON COLUMN public.mail_account.password IS '密码';
COMMENT ON COLUMN public.mail_account.host IS 'SMTP 服务器域名';
COMMENT ON COLUMN public.mail_account.port IS 'SMTP 服务器端口';
COMMENT ON COLUMN public.mail_account.ssl_enable IS '是否开启 SSL';
COMMENT ON COLUMN public.mail_account.remark IS '备注';
COMMENT ON COLUMN public.mail_account.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.mail_account.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_account.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_account.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_account ADD CONSTRAINT mail_account_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX mail_account_mail_idx ON public.mail_account USING btree (mail);
CREATE INDEX mail_account_status_idx ON public.mail_account USING btree (status);
