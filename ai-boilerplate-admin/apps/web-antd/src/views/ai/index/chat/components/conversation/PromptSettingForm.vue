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

const emit = defineEmits(['success']);
const formData = ref<AiIndexChatApi.AiIndexChatConversationItem>();

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
      fieldName: 'promptName',
      label: '提示词名称',
      component: 'Input',
      componentProps: {
        placeholder: '请输入提示词名称',
      },
    },
    {
      fieldName: 'promptDesc',
      label: '提示词描述',
      component: 'Textarea',
      componentProps: {
        rows: 2,
        placeholder: '请输入提示词描述',
      },
    },
    {
      fieldName: 'systemMessage',
      label: '提示词内容',
      component: 'Textarea',
      componentProps: {
        rows: 8,
        placeholder: '请输入提示词内容（系统角色设定）',
      },
      help: '定义 AI 的角色、行为风格和回复规则',
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
      promptName?: string;
      promptDesc?: string;
      systemMessage?: string;
    };
    try {
      const promptSetting = {
        ...(formData.value?.promptSetting || {}),
        name: values.promptName || '',
        desc: values.promptDesc || '',
        prompt: values.systemMessage || '',
      };

      await updateAiIndexChatConversation({
        body: {
          id: values.id || '',
          title: formData.value?.title || '',
          promptSetting,
          modelSetting: formData.value?.modelSetting,
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
        promptName: formData.value?.promptSetting?.name,
        promptDesc: formData.value?.promptSetting?.desc,
        systemMessage: formData.value?.promptSetting?.prompt,
      });
    } finally {
      modalApi.unlock();
    }
  },
});
</script>

<template>
  <Modal class="w-2/5" title="提示词设置">
    <Form class="mx-4" />
  </Modal>
</template>
