<script lang="ts" setup>
import type { AiChatRoleApi } from '#/api/ai/manager/chatRole';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createAiChatRole,
  getAiChatRoleInfo,
  updateAiChatRole,
} from '#/api/ai/manager/chatRole';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<AiChatRoleApi.AiChatRoleInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['AI聊天角色'])
    : $t('ui.actionTitle.create', ['AI聊天角色']);
});

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 120,
  },
  layout: 'horizontal',
  schema: useFormSchema(),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    // 提交表单
    const data = (await formApi.getValues()) as
      | AiChatRoleApi.CreateAiChatRoleReq
      | AiChatRoleApi.UpdateAiChatRoleReq;
    try {
      await (formData.value?.id
        ? updateAiChatRole({
            id: formData.value.id,
            ...data,
          } as AiChatRoleApi.UpdateAiChatRoleReq)
        : createAiChatRole(data as AiChatRoleApi.CreateAiChatRoleReq));
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success($t('ui.actionMessage.operationSuccess'));
    } finally {
      modalApi.unlock();
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<AiChatRoleApi.AiChatRoleInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const result = await getAiChatRoleInfo(data.id);
      formData.value = result.info;
      // 设置到 values
      await formApi.setValues(formData.value);
    } finally {
      modalApi.unlock();
    }
  },
});
</script>

<template>
  <Modal class="w-2/5" :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
