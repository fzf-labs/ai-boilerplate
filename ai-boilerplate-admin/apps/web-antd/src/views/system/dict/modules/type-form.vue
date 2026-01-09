<script lang="ts" setup>
import type { DictTypeInfo } from '#/api/v1/dict-type';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createDictType,
  getDictTypeInfo,
  updateDictType,
} from '#/api/v1/dict-type';
import { $t } from '#/locales';

import { useTypeFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<DictTypeInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['字典类型'])
    : $t('ui.actionTitle.create', ['字典类型']);
});

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useTypeFormSchema(),
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
    const data = (await formApi.getValues()) as DictTypeInfo;
    try {
      await (formData.value?.id
        ? updateDictType({ body: data })
        : createDictType({ body: data }));
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: $t('ui.actionMessage.operationSuccess'),
        key: 'action_process_msg',
      });
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<DictTypeInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getDictTypeInfo({ params: { id: data.id } });
      formData.value = res.info;
      // 设置到 values
      if (formData.value) {
        await formApi.setValues(formData.value);
      }
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
