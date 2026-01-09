CREATE TABLE public.sms_template (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sms_channel_id character varying(64) NOT NULL,
    template_type smallint NOT NULL,
    template_code character varying(64) NOT NULL,
    template_name character varying(64) NOT NULL,
    template_content character varying(500) NOT NULL,
    template_params jsonb,
    remark character varying(255),
    api_template_id character varying(64) NOT NULL,
    status smallint DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.sms_template IS '短信模板';
COMMENT ON COLUMN public.sms_template.id IS 'id';
COMMENT ON COLUMN public.sms_template.sms_channel_id IS '短信渠道编号';
COMMENT ON COLUMN public.sms_template.template_type IS '模板类型';
COMMENT ON COLUMN public.sms_template.template_code IS '模板编码';
COMMENT ON COLUMN public.sms_template.template_name IS '模板名称';
COMMENT ON COLUMN public.sms_template.template_content IS '模板内容';
COMMENT ON COLUMN public.sms_template.template_params IS '模板参数';
COMMENT ON COLUMN public.sms_template.remark IS '备注';
COMMENT ON COLUMN public.sms_template.api_template_id IS '短信供应商的模板编号';
COMMENT ON COLUMN public.sms_template.status IS '状态(-1禁用,1开启)';
COMMENT ON COLUMN public.sms_template.created_at IS '创建时间';
COMMENT ON COLUMN public.sms_template.updated_at IS '更新时间';
COMMENT ON COLUMN public.sms_template.deleted_at IS '删除时间';
ALTER TABLE ONLY public.sms_template ADD CONSTRAINT sms_template_pkey PRIMARY KEY (id);
CREATE INDEX sms_template_sms_channel_id_idx ON public.sms_template USING btree (sms_channel_id);
