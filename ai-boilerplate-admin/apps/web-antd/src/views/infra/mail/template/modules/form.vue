<script lang="ts" setup>
import type {
  CreateMailTemplateReq,
  MailTemplateInfo,
  UpdateMailTemplateReq,
} from '#/api/v1/mail-template';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createMailTemplate,
  getMailTemplateInfo,
  updateMailTemplate,
} from '#/api/v1/mail-template';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<MailTemplateInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['邮件模板'])
    : $t('ui.actionTitle.create', ['邮件模板']);
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
      await (formData.value?.id
        ? updateMailTemplate({ body: data as UpdateMailTemplateReq })
        : createMailTemplate({ body: data as CreateMailTemplateReq }));
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
    const data = modalApi.getData<MailTemplateInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getMailTemplateInfo({ params: { id: data.id } });
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
