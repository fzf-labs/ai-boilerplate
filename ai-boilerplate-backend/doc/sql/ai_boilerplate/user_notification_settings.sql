CREATE TABLE public.user_notification_settings (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    system_notification boolean DEFAULT true NOT NULL,
    activity_notification boolean DEFAULT true NOT NULL,
    order_notification boolean DEFAULT true NOT NULL,
    message_notification boolean DEFAULT true NOT NULL,
    dnd_enabled boolean DEFAULT false NOT NULL,
    dnd_start_time character varying(5),
    dnd_end_time character varying(5),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_notification_settings IS '用户通知设置表';
COMMENT ON COLUMN public.user_notification_settings.id IS 'id';
COMMENT ON COLUMN public.user_notification_settings.user_id IS '用户ID';
COMMENT ON COLUMN public.user_notification_settings.system_notification IS '系统通知';
COMMENT ON COLUMN public.user_notification_settings.activity_notification IS '活动通知';
COMMENT ON COLUMN public.user_notification_settings.order_notification IS '订单通知';
COMMENT ON COLUMN public.user_notification_settings.message_notification IS '消息通知';
COMMENT ON COLUMN public.user_notification_settings.dnd_enabled IS '勿扰模式启用';
COMMENT ON COLUMN public.user_notification_settings.dnd_start_time IS '勿扰开始时间';
COMMENT ON COLUMN public.user_notification_settings.dnd_end_time IS '勿扰结束时间';
COMMENT ON COLUMN public.user_notification_settings.created_at IS '创建时间';
COMMENT ON COLUMN public.user_notification_settings.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_notification_settings.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_notification_settings ADD CONSTRAINT user_notification_settings_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX user_notification_settings_user_id_idx ON public.user_notification_settings USING btree (user_id);
