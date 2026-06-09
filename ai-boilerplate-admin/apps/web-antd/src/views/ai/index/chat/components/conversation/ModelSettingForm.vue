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
import { getAiProviderModelList } from '#/api/v1/ai-provider-model';
import { $t } from '#/locales';

const emit = defineEmits(['success']);
const formData = ref<AiIndexChatApi.AiIndexChatConversationItem>();

async function getChatModelOptions() {
  const { list } = await getAiProviderModelList({
    params: { page: 1, pageSize: 200 },
  });
  return (list || [])
    .filter((item) => item.modelType === 'text')
    .map((item) => ({
      label: item.modelName || item.modelId,
      value: item.modelId,
    }));
}

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 120,
  },
  layout: 'horizontal',
  schema: [
    {
      component: 'Input',
      fieldName: 'id',
      dependencies: {
        triggerFields: [''],
        show: () => false,
      },
    },
    {
      component: 'ApiSelect',
      fieldName: 'modelId',
      label: '模型',
      componentProps: {
        api: getChatModelOptions,
        labelField: 'label',
        valueField: 'value',
        allowClear: true,
        placeholder: '请选择模型',
      },
      rules: 'required',
    },
    {
      fieldName: 'temperature',
      label: '温度参数',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入温度参数',
        class: 'w-full',
        precision: 2,
        min: 0,
        max: 2,
      },
      rules: 'required',
    },
    {
      fieldName: 'topP',
      label: 'Top P',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入 Top P 参数',
        class: 'w-full',
        precision: 2,
        min: 0,
        max: 1,
      },
    },
    {
      fieldName: 'maxTokens',
      label: '最大 Token 数',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入最大 Token 数',
        class: 'w-full',
        min: 0,
        max: 8192,
      },
      rules: 'required',
    },
    {
      fieldName: 'maxContexts',
      label: '上下文数量',
      component: 'InputNumber',
      componentProps: {
        controlsPosition: 'right',
        placeholder: '请输入上下文数量',
        class: 'w-full',
        min: 0,
        max: 20,
      },
      rules: 'required',
    },
  ],
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    const values = (await formApi.getValues()) as {
      id?: string;
      maxContexts?: number;
      maxTokens?: number;
      modelId?: string;
      temperature?: number;
      topP?: number;
    };
    try {
      const modelSetting = {
        ...formData.value?.modelSetting,
        modelId: values.modelId || '',
        temperature: values.temperature ?? 0,
        top_p: values.topP ?? 0,
        max_tokens: values.maxTokens ?? 0,
        max_contexts: values.maxContexts ?? 0,
      };

      await updateAiIndexChatConversation({
        body: {
          id: values.id || '',
          title: formData.value?.title || '',
          promptSetting: formData.value?.promptSetting,
          modelSetting,
          knowledgeSetting: formData.value?.knowledgeSetting,
          mcpSetting: formData.value?.mcpSetting,
        },
      });

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
      await formApi.setValues({
        id: formData.value?.id,
        modelId: formData.value?.modelSetting?.modelId,
        temperature: formData.value?.modelSetting?.temperature,
        topP: formData.value?.modelSetting?.top_p,
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
  <Modal class="w-2/5" title="模型设置">
    <Form class="mx-4" />
  </Modal>
</template>
