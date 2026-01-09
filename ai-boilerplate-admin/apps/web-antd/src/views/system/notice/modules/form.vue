<script lang="ts" setup>
import type {
  CreateSysNoticeReq,
  SysNoticeInfo,
  UpdateSysNoticeReq,
} from '#/api/v1/sys-notice';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createSysNotice,
  getSysNoticeInfo,
  updateSysNotice,
} from '#/api/v1/sys-notice';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<SysNoticeInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['公告'])
    : $t('ui.actionTitle.create', ['公告']);
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
    const data = (await formApi.getValues()) as SysNoticeInfo;
    try {
      await (formData.value?.id
        ? updateSysNotice({ body: data as UpdateSysNoticeReq })
        : createSysNotice({ body: data as CreateSysNoticeReq }));
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
    const data = modalApi.getData<SysNoticeInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getSysNoticeInfo({ params: { id: data.id } });
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
  <Modal :title="getTitle" class="w-1/2">
    <Form class="mx-4" />
  </Modal>
</template>
