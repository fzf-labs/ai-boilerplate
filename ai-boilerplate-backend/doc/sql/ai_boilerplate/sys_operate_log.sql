CREATE TABLE public.sys_operate_log (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    tenant_id character varying NOT NULL,
    trace_id character varying,
    admin_id uuid NOT NULL,
    ip character varying(32) NOT NULL,
    uri character varying(200) NOT NULL,
    useragent character varying(255),
    header json,
    req json,
    resp json,
    created_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.sys_operate_log IS '系统-日志';
COMMENT ON COLUMN public.sys_operate_log.id IS 'id';
COMMENT ON COLUMN public.sys_operate_log.tenant_id IS '租户Id';
COMMENT ON COLUMN public.sys_operate_log.trace_id IS '链路Id';
COMMENT ON COLUMN public.sys_operate_log.admin_id IS '管理员ID';
COMMENT ON COLUMN public.sys_operate_log.ip IS 'ip';
COMMENT ON COLUMN public.sys_operate_log.uri IS '请求路径';
COMMENT ON COLUMN public.sys_operate_log.useragent IS '浏览器标识';
COMMENT ON COLUMN public.sys_operate_log.header IS 'header';
COMMENT ON COLUMN public.sys_operate_log.req IS '请求数据';
COMMENT ON COLUMN public.sys_operate_log.resp IS '响应数据';
COMMENT ON COLUMN public.sys_operate_log.created_at IS '创建时间';
ALTER TABLE ONLY public.sys_operate_log ADD CONSTRAINT sys_operate_log_pkey PRIMARY KEY (id);
CREATE INDEX sys_operate_log_admin_id_idx ON public.sys_operate_log USING btree (admin_id);
CREATE INDEX sys_operate_log_tenant_id_idx ON public.sys_operate_log USING btree (tenant_id);
CREATE INDEX sys_operate_log_trace_id_idx ON public.sys_operate_log USING btree (trace_id);
