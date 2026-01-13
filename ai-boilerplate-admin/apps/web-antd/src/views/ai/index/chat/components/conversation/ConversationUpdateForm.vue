<script lang="ts" setup>
import type { AiChatConversationApi } from '#/api/ai/chat';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  getChatConversationMy,
  updateChatConversationMy,
} from '#/api/ai/chat';
import { $t } from '#/locales';

import { useFormSchema } from '../../data';

const emit = defineEmits(['success']);
const formData = ref<AiChatConversationApi.ChatConversation>();

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 140,
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
    const values =
      (await formApi.getValues()) as AiChatConversationApi.ChatConversation & {
        modelId?: string;
        temperature?: number;
        maxTokens?: number;
        maxContexts?: number;
      };
    try {
      const modelSetting = {
        ...(formData.value?.modelSetting || {}),
        modelId: values.modelId || formData.value?.modelSetting?.modelId,
        temperature:
          values.temperature ?? formData.value?.modelSetting?.temperature,
        max_tokens:
          values.maxTokens ?? formData.value?.modelSetting?.max_tokens,
        max_contexts:
          values.maxContexts ?? formData.value?.modelSetting?.max_contexts,
      };
      await updateChatConversationMy({
        id: values.id,
        systemMessage: values.systemMessage,
        modelSetting,
      } as AiChatConversationApi.ChatConversation);

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
    const data = modalApi.getData<AiChatConversationApi.ChatConversation>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      formData.value = await getChatConversationMy(data.id as string);
      // 设置到 values
      await formApi.setValues({
        id: formData.value.id,
        systemMessage: formData.value.systemMessage,
        modelId: formData.value.modelSetting?.modelId,
        temperature: formData.value.modelSetting?.temperature,
        maxTokens: formData.value.modelSetting?.max_tokens,
        maxContexts: formData.value.modelSetting?.max_contexts,
      });
    } finally {
      modalApi.unlock();
    }
  },
});
</script>

<template>
  <Modal class="w-2/5" title="设定">
    <Form class="mx-4" />
  </Modal>
</template>
