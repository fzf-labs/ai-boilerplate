<script lang="ts" setup>
import type { AiProviderModelApi } from '#/api/v1/ai-provider-model';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createAiProviderModel,
  getAiProviderModelInfo,
  updateAiProviderModel,
} from '#/api/v1/ai-provider-model';
import { $t } from '#/locales';

import { useModelFormSchema } from './model-data';

const emit = defineEmits(['success']);
const formData = ref<Partial<AiProviderModelApi.AiProviderModelInfo>>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['模型'])
    : $t('ui.actionTitle.create', ['模型']);
});

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 100,
  },
  layout: 'horizontal',
  schema: useModelFormSchema(),
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
      | AiProviderModelApi.CreateAiProviderModelReq
      | AiProviderModelApi.UpdateAiProviderModelReq;
    try {
      await (formData.value?.id
        ? updateAiProviderModel({
            ...data,
            id: formData.value.id,
          } as AiProviderModelApi.UpdateAiProviderModelReq)
        : createAiProviderModel(
            data as AiProviderModelApi.CreateAiProviderModelReq,
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
      modalApi.getData<Partial<AiProviderModelApi.AiProviderModelInfo>>();
    if (!data) {
      return;
    }

    // 如果有 id，说明是编辑模式，加载详细信息
    if (data.id) {
      modalApi.lock();
      try {
        const response = await getAiProviderModelInfo(data.id);
        formData.value = response.info;
        // 设置到 values
        await formApi.setValues(formData.value);
      } finally {
        modalApi.unlock();
      }
    } else {
      // 新增模式，设置 platformId
      formData.value = data;
      if (data.platformId) {
        await formApi.setValues({ platformId: data.platformId });
      }
    }
  },
});
</script>

<template>
  <Modal class="w-2/5" :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
