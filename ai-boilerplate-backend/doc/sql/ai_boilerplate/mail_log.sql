CREATE TABLE public.mail_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    account_id character varying(64) NOT NULL,
    from_mail character varying(255) NOT NULL,
    to_mail character varying(255) NOT NULL,
    template_id character varying(64) NOT NULL,
    template_code character varying(64) NOT NULL,
    template_nickname character varying(255),
    template_title character varying(255) NOT NULL,
    template_content text NOT NULL,
    template_params character varying(255) NOT NULL,
    send_status integer NOT NULL,
    send_time timestamp with time zone NOT NULL,
    send_message_id character varying(255),
    send_exception character varying(4096),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.mail_log IS '邮件日志表';
COMMENT ON COLUMN public.mail_log.id IS 'id';
COMMENT ON COLUMN public.mail_log.account_id IS '邮箱账号编号';
COMMENT ON COLUMN public.mail_log.from_mail IS '发送邮箱地址';
COMMENT ON COLUMN public.mail_log.to_mail IS '接收邮箱地址';
COMMENT ON COLUMN public.mail_log.template_id IS '模板编号';
COMMENT ON COLUMN public.mail_log.template_code IS '模板编码';
COMMENT ON COLUMN public.mail_log.template_nickname IS '模版发送人名称';
COMMENT ON COLUMN public.mail_log.template_title IS '邮件标题';
COMMENT ON COLUMN public.mail_log.template_content IS '邮件内容';
COMMENT ON COLUMN public.mail_log.template_params IS '邮件参数';
COMMENT ON COLUMN public.mail_log.send_status IS '发送状态';
COMMENT ON COLUMN public.mail_log.send_time IS '发送时间';
COMMENT ON COLUMN public.mail_log.send_message_id IS '发送返回的消息 ID';
COMMENT ON COLUMN public.mail_log.send_exception IS '发送异常';
COMMENT ON COLUMN public.mail_log.created_at IS '创建时间';
COMMENT ON COLUMN public.mail_log.updated_at IS '更新时间';
COMMENT ON COLUMN public.mail_log.deleted_at IS '删除时间';
ALTER TABLE ONLY public.mail_log ADD CONSTRAINT mail_log_pkey PRIMARY KEY (id);
