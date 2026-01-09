CREATE TABLE public.sms_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sms_channel_id character varying(64) NOT NULL,
    sms_template_id character varying(64) NOT NULL,
    sms_params_content character varying(255) NOT NULL,
    mobile character varying(11) NOT NULL,
    user_id character varying(64),
    send_status character varying(64) NOT NULL,
    send_time timestamp with time zone NOT NULL,
    receive_status character varying(64) NOT NULL,
    receive_time timestamp with time zone,
    api_send_code character varying(64),
    api_send_msg character varying(255),
    api_request_id character varying(255),
    api_serial_no character varying(255),
    api_receive_code character varying(64),
    api_receive_msg character varying(255),
    created_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.sms_log IS '短信日志';
COMMENT ON COLUMN public.sms_log.id IS '编号';
COMMENT ON COLUMN public.sms_log.sms_channel_id IS '短信渠道编号';
COMMENT ON COLUMN public.sms_log.sms_template_id IS '模板编号';
COMMENT ON COLUMN public.sms_log.sms_params_content IS '短信参数内容';
COMMENT ON COLUMN public.sms_log.mobile IS '手机号';
COMMENT ON COLUMN public.sms_log.user_id IS '用户id';
COMMENT ON COLUMN public.sms_log.send_status IS '发送状态';
COMMENT ON COLUMN public.sms_log.send_time IS '发送时间';
COMMENT ON COLUMN public.sms_log.receive_status IS '接收状态';
COMMENT ON COLUMN public.sms_log.receive_time IS '接收时间';
COMMENT ON COLUMN public.sms_log.api_send_code IS '短信 API 发送结果的编码';
COMMENT ON COLUMN public.sms_log.api_send_msg IS '短信 API 发送失败的提示';
COMMENT ON COLUMN public.sms_log.api_request_id IS '短信 API 发送返回的唯一请求 ID';
COMMENT ON COLUMN public.sms_log.api_serial_no IS '短信 API 发送返回的序号';
COMMENT ON COLUMN public.sms_log.api_receive_code IS 'API 接收结果的编码';
COMMENT ON COLUMN public.sms_log.api_receive_msg IS 'API 接收结果的说明';
COMMENT ON COLUMN public.sms_log.created_at IS '创建时间';
ALTER TABLE ONLY public.sms_log ADD CONSTRAINT sms_log_pkey PRIMARY KEY (id);
CREATE INDEX sms_log_mobile_idx ON public.sms_log USING btree (mobile);
