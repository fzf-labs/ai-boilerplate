<script lang="ts" setup>
import type { AiProviderPlatformApi } from '#/api/ai/manager/providerPlatform';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createAiProviderPlatform,
  getAiProviderPlatformInfo,
  updateAiProviderPlatform,
} from '#/api/ai/manager/providerPlatform';
import { $t } from '#/locales';

import { usePlatformFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<AiProviderPlatformApi.AiProviderPlatformInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['供应商平台'])
    : $t('ui.actionTitle.create', ['供应商平台']);
});

/** 处理平台选择变化 */
function handlePlatformChange(_value: string, label: string) {
  // 只在新增模式下自动填充 name
  if (!formData.value?.id) {
    formApi.setFieldValue('name', label);
  }
}

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 100,
  },
  layout: 'horizontal',
  schema: usePlatformFormSchema(handlePlatformChange),
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
      | AiProviderPlatformApi.CreateAiProviderPlatformReq
      | AiProviderPlatformApi.UpdateAiProviderPlatformReq;
    try {
      await (formData.value?.id
        ? updateAiProviderPlatform({
            ...data,
            id: formData.value.id,
          } as AiProviderPlatformApi.UpdateAiProviderPlatformReq)
        : createAiProviderPlatform(
            data as AiProviderPlatformApi.CreateAiProviderPlatformReq,
          ));
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
    const data =
      modalApi.getData<AiProviderPlatformApi.AiProviderPlatformInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const response = await getAiProviderPlatformInfo(data.id);
      formData.value = response.info;
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
