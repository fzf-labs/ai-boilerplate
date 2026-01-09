<script lang="ts" setup>
import { computed } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { registerDevice } from '#/api/v1/device';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const getTitle = computed(() => {
  return '注册设备';
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
    const data = await formApi.getValues();
    try {
      // 注册设备
      await registerDevice({ body: { sn: data.sn } });
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: '设备注册成功',
        key: 'action_process_msg',
      });
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal :title="getTitle" class="w-full max-w-2xl">
    <Form class="mx-2" />
  </Modal>
</template>
