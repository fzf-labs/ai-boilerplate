CREATE TABLE public.sys_notify_message (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying(64) NOT NULL,
    type character varying(64) NOT NULL,
    subject character varying(200) NOT NULL,
    content character varying(200) NOT NULL,
    sender character varying(64) NOT NULL,
    receiver character varying(64) NOT NULL,
    send_time character varying(64) NOT NULL,
    read_time character varying(64) DEFAULT ''::character varying,
    extend jsonb
);
COMMENT ON TABLE public.sys_notify_message IS '系统-通知消息';
COMMENT ON COLUMN public.sys_notify_message.id IS 'id';
COMMENT ON COLUMN public.sys_notify_message.tenant_id IS '租户id';
COMMENT ON COLUMN public.sys_notify_message.type IS '消息类型';
COMMENT ON COLUMN public.sys_notify_message.subject IS '主题';
COMMENT ON COLUMN public.sys_notify_message.content IS '内容';
COMMENT ON COLUMN public.sys_notify_message.sender IS '发送人';
COMMENT ON COLUMN public.sys_notify_message.receiver IS '接收人';
COMMENT ON COLUMN public.sys_notify_message.send_time IS '发送时间';
COMMENT ON COLUMN public.sys_notify_message.read_time IS '阅读时间';
COMMENT ON COLUMN public.sys_notify_message.extend IS '扩展';
ALTER TABLE ONLY public.sys_notify_message ADD CONSTRAINT sys_notify_message_pkey PRIMARY KEY (id);
CREATE INDEX sys_notify_message_receiver_read_time_idx ON public.sys_notify_message USING btree (receiver, read_time);
