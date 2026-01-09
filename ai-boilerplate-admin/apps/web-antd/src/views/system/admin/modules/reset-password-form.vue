<script lang="ts" setup>
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { updateSysAdminPassword } from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import { useResetPasswordFormSchema } from '../data';

const emit = defineEmits(['success']);
const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useResetPasswordFormSchema(),
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
      await updateSysAdminPassword({
        body: {
          id: data.id as string,
          password: data.newPassword as string,
        },
      });
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
      return;
    }
    // 加载数据
    const data = modalApi.getData<SysAdminInfo>();
    if (!data || !data.id) {
      return;
    }
    // 设置到 values
    await formApi.setValues(data);
  },
});
</script>

<template>
  <Modal title="重置密码">
    <Form class="mx-4" />
  </Modal>
</template>
