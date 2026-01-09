<script lang="ts" setup>
import type { SysDeptInfo } from '#/api/v1/sys-dept';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createSysDept,
  getSysDeptInfo,
  updateSysDept,
} from '#/api/v1/sys-dept';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<SysDeptInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['部门'])
    : $t('ui.actionTitle.create', ['部门']);
});

const [Form, formApi] = useVbenForm({
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
    const data = (await formApi.getValues()) as SysDeptInfo;
    try {
      await (formData.value?.id
        ? updateSysDept({ body: { ...data, id: formData.value.id } })
        : createSysDept({ body: data }));
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
    let data = modalApi.getData<SysDeptInfo>();
    if (!data) {
      return;
    }
    if (data.id) {
      modalApi.lock();
      try {
        const res = await getSysDeptInfo({ params: { id: data.id } });
        data = res.info;
      } finally {
        modalApi.lock(false);
      }
    }
    // 设置到 values
    formData.value = data;
    await formApi.setValues(formData.value);
  },
});
</script>

<template>
  <Modal :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
