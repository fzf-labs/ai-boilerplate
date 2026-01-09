CREATE TABLE public.user_bind_device (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id character varying(64) NOT NULL,
    sn character varying(64) NOT NULL,
    identity character varying(64) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);
COMMENT ON TABLE public.user_bind_device IS '用户绑定设备表';
COMMENT ON COLUMN public.user_bind_device.id IS 'id';
COMMENT ON COLUMN public.user_bind_device.user_id IS '用户Id';
COMMENT ON COLUMN public.user_bind_device.sn IS 'sn';
COMMENT ON COLUMN public.user_bind_device.identity IS '身份';
COMMENT ON COLUMN public.user_bind_device.created_at IS '创建时间';
COMMENT ON COLUMN public.user_bind_device.updated_at IS '更新时间';
ALTER TABLE ONLY public.user_bind_device ADD CONSTRAINT user_bind_device_pkey PRIMARY KEY (id);
CREATE INDEX user_bind_device_sn_idx ON public.user_bind_device USING btree (sn);
CREATE UNIQUE INDEX user_bind_device_user_id_sn_idx ON public.user_bind_device USING btree (user_id, sn);
