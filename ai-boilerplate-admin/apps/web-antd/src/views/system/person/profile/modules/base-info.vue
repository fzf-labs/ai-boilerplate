<script setup lang="ts">
import type { Recordable } from '@vben/types';

import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { watch } from 'vue';

import { $t } from '@vben/locales';

import { message } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import { sysAuthUpdateAdminInfo } from '#/api/v1/sys-auth';

const props = defineProps<{
  profile?: SysAdminInfo;
}>();
const emit = defineEmits<{
  (e: 'success'): void;
}>();

const [Form, formApi] = useVbenForm({
  commonConfig: {
    labelWidth: 70,
  },
  schema: [
    {
      label: '用户昵称',
      fieldName: 'nickname',
      component: 'Input',
      componentProps: {
        placeholder: '请输入用户昵称',
      },
      rules: 'required',
    },
    {
      label: '用户性别',
      fieldName: 'sex',
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '男', value: 1 },
          { label: '女', value: 2 },
        ],
        buttonStyle: 'solid',
        optionType: 'button',
      },
      rules: z.number(),
    },
  ],
  resetButtonOptions: {
    show: false,
  },
  submitButtonOptions: {
    content: '更新信息',
  },
  handleSubmit,
});

async function handleSubmit(values: Recordable<any>) {
  try {
    // 提交表单，头像使用原值
    await sysAuthUpdateAdminInfo({
      body: {
        ...values,
        avatar: props.profile?.avatar || '',
      },
    });
    // 关闭并提示
    emit('success');
    message.success($t('ui.actionMessage.operationSuccess'));
  } catch (error) {
    console.error('更新个人信息失败:', error);
    message.error('更新个人信息失败，请重试');
  }
}

/** 监听 profile 变化 */
watch(
  () => props.profile,
  (newProfile) => {
    if (newProfile) {
      formApi.setValues(newProfile);
    }
  },
  { immediate: true },
);
</script>

<template>
  <div class="mt-4 md:w-full lg:w-1/2 2xl:w-2/5">
    <Form />
  </div>
</template>
