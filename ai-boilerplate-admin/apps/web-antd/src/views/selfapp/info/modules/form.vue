<script lang="ts" setup>
import type {
  CreateSelfAppReq,
  SelfAppInfo,
  UpdateSelfAppReq,
} from '#/api/v1/self-app';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createSelfApp,
  getSelfAppInfo,
  updateSelfApp,
} from '#/api/v1/self-app';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<SelfAppInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['自应用'])
    : $t('ui.actionTitle.create', ['自应用']);
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
    const rawData = (await formApi.getValues()) as any;

    // 处理数据格式转换
    const data: CreateSelfAppReq | UpdateSelfAppReq = {
      ...rawData,
      // 确保数字类型字段正确
      status: Number(rawData.status),
    };

    try {
      await (formData.value?.id
        ? updateSelfApp({ body: data as UpdateSelfAppReq })
        : createSelfApp({ body: data as CreateSelfAppReq }));
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
    const data = modalApi.getData<SelfAppInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getSelfAppInfo({ params: { id: data.id } });
      if (!res.info) {
        return;
      }
      formData.value = res.info;
      // 处理数据格式，转换为表单可用的格式
      const formValues = {
        ...formData.value,
        // 确保数字类型字段正确转换
        status: Number(formData.value?.status),
      };

      // 设置到 values
      await formApi.setValues(formValues);
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal :title="getTitle" :width="1000">
    <Form class="mx-4" />
  </Modal>
</template>
