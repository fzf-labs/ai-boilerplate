CREATE TABLE public.user_notification_settings (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    notification_preferences jsonb DEFAULT '{}'::jsonb NOT NULL,
    dnd_start_time character varying(5),
    dnd_end_time character varying(5),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);
COMMENT ON TABLE public.user_notification_settings IS '用户通知设置表';
COMMENT ON COLUMN public.user_notification_settings.id IS 'id';
COMMENT ON COLUMN public.user_notification_settings.user_id IS '用户ID';
COMMENT ON COLUMN public.user_notification_settings.notification_preferences IS '通知偏好(JSON)';
COMMENT ON COLUMN public.user_notification_settings.dnd_start_time IS '勿扰开始时间';
COMMENT ON COLUMN public.user_notification_settings.dnd_end_time IS '勿扰结束时间';
COMMENT ON COLUMN public.user_notification_settings.created_at IS '创建时间';
COMMENT ON COLUMN public.user_notification_settings.updated_at IS '更新时间';
COMMENT ON COLUMN public.user_notification_settings.deleted_at IS '删除时间';
ALTER TABLE ONLY public.user_notification_settings ADD CONSTRAINT user_notification_settings_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX user_notification_settings_user_id_idx ON public.user_notification_settings USING btree (user_id);
