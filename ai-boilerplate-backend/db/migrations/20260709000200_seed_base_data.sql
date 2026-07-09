-- +goose Up
INSERT INTO public.sys_menu (
    id, pid, name, type, path, permission, icon, component, component_name,
    sort, status, created_at, updated_at
) VALUES
    ('00000000-0000-0000-0000-000000000010', '', 'Dashboard', 'menu', '/dashboard', 'dashboard:view', 'dashboard', '/dashboard/index', 'Dashboard', 10, 1, now(), now()),
    ('00000000-0000-0000-0000-000000000020', '', 'System', 'dir', '/system', 'system:view', 'settings', '', 'System', 20, 1, now(), now()),
    ('00000000-0000-0000-0000-000000000021', '00000000-0000-0000-0000-000000000020', 'Admins', 'menu', '/system/admins', 'sys:admin:list', 'user', '/system/admins/index', 'SysAdmin', 21, 1, now(), now()),
    ('00000000-0000-0000-0000-000000000022', '00000000-0000-0000-0000-000000000020', 'Roles', 'menu', '/system/roles', 'sys:role:list', 'shield', '/system/roles/index', 'SysRole', 22, 1, now(), now())
ON CONFLICT (id) DO UPDATE SET
    pid = EXCLUDED.pid,
    name = EXCLUDED.name,
    type = EXCLUDED.type,
    path = EXCLUDED.path,
    permission = EXCLUDED.permission,
    icon = EXCLUDED.icon,
    component = EXCLUDED.component,
    component_name = EXCLUDED.component_name,
    sort = EXCLUDED.sort,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.sys_role (
    id, tenant_id, name, remark, "dataScope", "menuIds", sort, status, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000003',
    '00000000-0000-0000-0000-000000000001',
    'Super Admin',
    'Default full-access role for local bootstrap.',
    'all',
    '["00000000-0000-0000-0000-000000000010","00000000-0000-0000-0000-000000000020","00000000-0000-0000-0000-000000000021","00000000-0000-0000-0000-000000000022"]'::jsonb,
    1,
    1,
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    tenant_id = EXCLUDED.tenant_id,
    name = EXCLUDED.name,
    remark = EXCLUDED.remark,
    "dataScope" = EXCLUDED."dataScope",
    "menuIds" = EXCLUDED."menuIds",
    sort = EXCLUDED.sort,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.sys_tenant (
    id, name, remark, admin_id, expire_time, "menuIds", status, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000001',
    'Default Tenant',
    'Default tenant for local bootstrap.',
    '00000000-0000-0000-0000-000000000002',
    '2099-12-31 23:59:59+00',
    '["00000000-0000-0000-0000-000000000010","00000000-0000-0000-0000-000000000020","00000000-0000-0000-0000-000000000021","00000000-0000-0000-0000-000000000022"]'::jsonb,
    1,
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    remark = EXCLUDED.remark,
    admin_id = EXCLUDED.admin_id,
    expire_time = EXCLUDED.expire_time,
    "menuIds" = EXCLUDED."menuIds",
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.sys_dept (
    id, pid, name, admin_id, status, sort, created_at, updated_at, tenant_id
) VALUES (
    '00000000-0000-0000-0000-000000000004',
    '',
    'Default Department',
    '00000000-0000-0000-0000-000000000002',
    1,
    1,
    now(),
    now(),
    '00000000-0000-0000-0000-000000000001'
)
ON CONFLICT (id) DO UPDATE SET
    pid = EXCLUDED.pid,
    name = EXCLUDED.name,
    admin_id = EXCLUDED.admin_id,
    status = EXCLUDED.status,
    sort = EXCLUDED.sort,
    tenant_id = EXCLUDED.tenant_id,
    updated_at = now();

INSERT INTO public.sys_post (
    id, tenant_id, name, code, remark, sort, status, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000005',
    '00000000-0000-0000-0000-000000000001',
    'Super Admin',
    'super_admin',
    'Default local bootstrap post.',
    1,
    1,
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    tenant_id = EXCLUDED.tenant_id,
    name = EXCLUDED.name,
    code = EXCLUDED.code,
    remark = EXCLUDED.remark,
    sort = EXCLUDED.sort,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.sys_admin (
    id, tenant_id, username, password, nickname, avatar, email, sex, mobile,
    role_id, dept_id, post_id, status, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000002',
    '00000000-0000-0000-0000-000000000001',
    'admin',
    '$2y$10$2qeogpFIUO8jfa.dDjAeUOY8SCMsExd4WO5PmUJRm.xjdCd61R4mm',
    'Admin',
    '',
    'admin@example.com',
    0,
    '',
    '00000000-0000-0000-0000-000000000003',
    '00000000-0000-0000-0000-000000000004',
    '00000000-0000-0000-0000-000000000005',
    1,
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    tenant_id = EXCLUDED.tenant_id,
    username = EXCLUDED.username,
    password = EXCLUDED.password,
    nickname = EXCLUDED.nickname,
    avatar = EXCLUDED.avatar,
    email = EXCLUDED.email,
    sex = EXCLUDED.sex,
    mobile = EXCLUDED.mobile,
    role_id = EXCLUDED.role_id,
    dept_id = EXCLUDED.dept_id,
    post_id = EXCLUDED.post_id,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.config_data (
    id, name, key, value, remark, status, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000030',
    'Site Name',
    'site_name',
    '"AI Boilerplate"'::jsonb,
    'Default site name.',
    1,
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    key = EXCLUDED.key,
    value = EXCLUDED.value,
    remark = EXCLUDED.remark,
    status = EXCLUDED.status,
    updated_at = now();

INSERT INTO public.dict_type (
    id, name, type, status, remark, created_at, updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000040',
    'System Status',
    'sys_status',
    1,
    'Default enable/disable status dictionary.',
    now(),
    now()
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    type = EXCLUDED.type,
    status = EXCLUDED.status,
    remark = EXCLUDED.remark,
    updated_at = now();

INSERT INTO public.dict_data (
    id, type, label, key, value, remark, css_color, css_class, status, created_at, updated_at
) VALUES
    ('00000000-0000-0000-0000-000000000041', 'sys_status', 'Enabled', 'enable', '1', 'Enabled status.', 'green', '', 1, now(), now()),
    ('00000000-0000-0000-0000-000000000042', 'sys_status', 'Disabled', 'disable', '-1', 'Disabled status.', 'red', '', 1, now(), now())
ON CONFLICT (id) DO UPDATE SET
    type = EXCLUDED.type,
    label = EXCLUDED.label,
    key = EXCLUDED.key,
    value = EXCLUDED.value,
    remark = EXCLUDED.remark,
    css_color = EXCLUDED.css_color,
    css_class = EXCLUDED.css_class,
    status = EXCLUDED.status,
    updated_at = now();

-- +goose Down
DELETE FROM public.dict_data
WHERE id IN (
    '00000000-0000-0000-0000-000000000041',
    '00000000-0000-0000-0000-000000000042'
);

DELETE FROM public.dict_type
WHERE id = '00000000-0000-0000-0000-000000000040';

DELETE FROM public.config_data
WHERE id = '00000000-0000-0000-0000-000000000030';

DELETE FROM public.sys_admin
WHERE id = '00000000-0000-0000-0000-000000000002';

DELETE FROM public.sys_post
WHERE id = '00000000-0000-0000-0000-000000000005';

DELETE FROM public.sys_dept
WHERE id = '00000000-0000-0000-0000-000000000004';

DELETE FROM public.sys_tenant
WHERE id = '00000000-0000-0000-0000-000000000001';

DELETE FROM public.sys_role
WHERE id = '00000000-0000-0000-0000-000000000003';

DELETE FROM public.sys_menu
WHERE id IN (
    '00000000-0000-0000-0000-000000000010',
    '00000000-0000-0000-0000-000000000020',
    '00000000-0000-0000-0000-000000000021',
    '00000000-0000-0000-0000-000000000022'
);
