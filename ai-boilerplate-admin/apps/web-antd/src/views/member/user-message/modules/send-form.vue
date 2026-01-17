<script lang="ts" setup>
import type { SendUserMessageReq } from '#/api/v1/user-message';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { sendUserMessage } from '#/api/v1/user-message';

import { useSendFormSchema } from '../data';

const emit = defineEmits(['success']);

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useSendFormSchema(),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }

    const values = await formApi.getValues();
    const payload: SendUserMessageReq = {
      category: values.category,
      title: values.title,
      summary: values.summary,
      coverURL: values.coverURL,
      content: values.content,
      linkURL: values.linkURL,
      audienceType: values.audienceType,
    };

    if (values.audienceType === 'users') {
      const userIds = parseUserIds(values.userIdsText);
      if (userIds.length === 0) {
        message.warning('请输入用户ID列表');
        return;
      }
      payload.userIds = userIds;
    }

    if (values.audienceType === 'segment') {
      const activeWithinDays = Number(values.activeWithinDays || 0);
      if (!values.membershipType && activeWithinDays <= 0) {
        message.warning('请选择会员类型或填写活跃天数');
        return;
      }
      if (values.membershipType) {
        payload.membershipType = values.membershipType;
      }
      if (activeWithinDays > 0) {
        payload.activeWithinDays = activeWithinDays;
      }
    }

    modalApi.lock();
    try {
      const res = await sendUserMessage({ body: payload });
      message.success(`已发送 ${res.total ?? 0} 条消息`);
      emit('success');
      await modalApi.close();
      await formApi.resetForm();
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      return;
    }
    await formApi.resetForm();
  },
});

function parseUserIds(value?: string): string[] {
  if (!value) {
    return [];
  }
  const chunks = value.split(/[\s,]+/).map((item) => item.trim());
  return [...new Set(chunks.filter(Boolean))];
}
</script>

<template>
  <Modal title="发送消息" class="w-full max-w-2xl">
    <Form class="mx-4" />
  </Modal>
</template>
