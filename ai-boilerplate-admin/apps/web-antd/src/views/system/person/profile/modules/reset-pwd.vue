<script setup lang="ts">
import type { Recordable } from '@vben/types';

import { $t } from '@vben/locales';

import { message } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import { sysAuthUpdateAdminPassword } from '#/api/v1/sys-auth';

const [Form, formApi] = useVbenForm({
  commonConfig: {
    labelWidth: 70,
  },
  schema: [
    {
      component: 'InputPassword',
      fieldName: 'oldPassword',
      label: '旧密码',
      rules: z
        .string({ message: '请输入密码' })
        .min(5, '密码长度不能少于 5 个字符')
        .max(20, '密码长度不能超过 20 个字符'),
    },
    {
      component: 'InputPassword',
      dependencies: {
        rules(values) {
          return z
            .string({ message: '请输入新密码' })
            .min(5, '密码长度不能少于 5 个字符')
            .max(20, '密码长度不能超过 20 个字符')
            .refine(
              (value) => value !== values.oldPassword,
              '新旧密码不能相同',
            );
        },
        triggerFields: ['newPassword', 'oldPassword'],
      },
      fieldName: 'newPassword',
      label: '新密码',
      rules: 'required',
    },
    {
      component: 'InputPassword',
      dependencies: {
        rules(values) {
          return z
            .string({ message: '请输入确认密码' })
            .min(5, '密码长度不能少于 5 个字符')
            .max(20, '密码长度不能超过 20 个字符')
            .refine(
              (value) => value === values.newPassword,
              '新密码和确认密码不一致',
            );
        },
        triggerFields: ['newPassword', 'confirmPassword'],
      },
      fieldName: 'confirmPassword',
      label: '确认密码',
      rules: 'required',
    },
  ],
  resetButtonOptions: {
    show: false,
  },
  submitButtonOptions: {
    content: '修改密码',
  },
  handleSubmit,
});

async function handleSubmit(values: Recordable<any>) {
  try {
    // 提交表单
    await sysAuthUpdateAdminPassword({
      body: {
        oldPassword: values.oldPassword,
        newPassword: values.newPassword,
      },
    });
    message.success($t('ui.actionMessage.operationSuccess'));
    // 重置表单
    formApi.resetForm();
  } catch (error) {
    console.error(error);
  }
}
</script>

<template>
  <div class="mt-4 md:w-full lg:w-1/2 2xl:w-2/5">
    <Form />
  </div>
</template>
