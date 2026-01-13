<script lang="ts" setup>
import type * as AiIndexChatApi from '#/api/v1/ai-index-chat';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  getAiIndexChatConversationItem,
  updateAiIndexChatConversation,
} from '#/api/v1/ai-index-chat';
import { $t } from '#/locales';

import { useFormSchema } from '../../data';

const emit = defineEmits(['success']);
const formData = ref<AiIndexChatApi.AiIndexChatConversationItem>();

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
    const values = (await formApi.getValues()) as {
      id?: string;
      systemMessage?: string;
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
      const promptSetting = {
        ...(formData.value?.promptSetting || {}),
        prompt: values.systemMessage,
      };
      await updateAiIndexChatConversation({
        body: {
          id: values.id || '',
          promptSetting,
          modelSetting,
        },
      });

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
    const data = modalApi.getData<AiIndexChatApi.AiIndexChatConversationItem>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getAiIndexChatConversationItem({
        params: { id: data.id as string },
      });
      formData.value = res.info;
      // 设置到 values
      await formApi.setValues({
        id: formData.value?.id,
        systemMessage: formData.value?.promptSetting?.prompt,
        modelId: formData.value?.modelSetting?.modelId,
        temperature: formData.value?.modelSetting?.temperature,
        maxTokens: formData.value?.modelSetting?.max_tokens,
        maxContexts: formData.value?.modelSetting?.max_contexts,
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
