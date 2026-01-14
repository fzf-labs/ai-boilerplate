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

// TODO: 替换为实际的知识库 API
async function getKnowledgeOptions() {
  // 临时返回空数组，等待后端知识库接口实现
  return [];
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
      fieldName: 'knowledgeIds',
      label: '知识库',
      componentProps: {
        api: getKnowledgeOptions,
        labelField: 'label',
        valueField: 'value',
        mode: 'multiple',
        allowClear: true,
        placeholder: '请选择知识库（支持多选）',
      },
      help: '选择对话使用的知识库，可多选。AI 将基于所选知识库内容回答问题。',
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
      knowledgeIds?: string[];
    };
    try {
      const knowledgeSetting = {
        knowledgeIds: values.knowledgeIds || [],
      };

      await updateAiIndexChatConversation({
        body: {
          id: values.id || '',
          title: formData.value?.title || '',
          promptSetting: formData.value?.promptSetting,
          modelSetting: formData.value?.modelSetting,
          knowledgeSetting,
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
        knowledgeIds: formData.value?.knowledgeSetting?.knowledgeIds || [],
      });
    } finally {
      modalApi.unlock();
    }
  },
});
</script>

<template>
  <Modal class="w-2/5" title="知识库设置">
    <Form class="mx-4" />
    <template #footer>
      <div class="text-muted-foreground mb-2 text-sm">
        💡 提示：暂无可用知识库，等待后端接口实现后即可使用
      </div>
    </template>
  </Modal>
</template>
