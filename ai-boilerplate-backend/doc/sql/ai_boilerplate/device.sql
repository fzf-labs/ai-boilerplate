CREATE TABLE public.device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    sn character varying(128) NOT NULL,
    name character varying(225),
    "desc" character varying(225),
    brand character varying(225),
    model character varying(225),
    network character varying(225),
    imei character varying(225),
    cpu character varying(125),
    mac character varying(125),
    app_version character varying(125),
    android_version character varying(125),
    ram_size numeric,
    ddr_size numeric,
    certificate character varying(225),
    secure_key character varying(225),
    registry_time timestamp with time zone,
    push jsonb,
    status integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.device IS '设备表';
COMMENT ON COLUMN public.device.id IS 'ID';
COMMENT ON COLUMN public.device.sn IS '设备ID';
COMMENT ON COLUMN public.device.name IS '设备名称';
COMMENT ON COLUMN public.device."desc" IS '描述';
COMMENT ON COLUMN public.device.brand IS '设备品牌';
COMMENT ON COLUMN public.device.model IS '设备型号';
COMMENT ON COLUMN public.device.network IS '入网型号';
COMMENT ON COLUMN public.device.imei IS 'IMEI';
COMMENT ON COLUMN public.device.cpu IS 'cpu型号';
COMMENT ON COLUMN public.device.mac IS 'mac地址';
COMMENT ON COLUMN public.device.app_version IS 'app版本';
COMMENT ON COLUMN public.device.android_version IS '安卓版本';
COMMENT ON COLUMN public.device.ram_size IS 'RAM大小';
COMMENT ON COLUMN public.device.ddr_size IS 'DDR大小';
COMMENT ON COLUMN public.device.certificate IS '设备证书';
COMMENT ON COLUMN public.device.secure_key IS '设备密钥';
COMMENT ON COLUMN public.device.registry_time IS '激活时间';
COMMENT ON COLUMN public.device.push IS '推送';
COMMENT ON COLUMN public.device.status IS '状态';
COMMENT ON COLUMN public.device.created_at IS '创建时间';
COMMENT ON COLUMN public.device.updated_at IS '更新时间';
COMMENT ON COLUMN public.device.deleted_at IS '删除时间';
ALTER TABLE ONLY public.device ADD CONSTRAINT device_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX device_sn_idx ON public.device USING btree (sn);
