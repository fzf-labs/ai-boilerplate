<script lang="ts" setup>
import type {
  CreateWxGzhTagReq,
  UpdateWxGzhTagReq,
  WxGzhTagInfo,
} from '#/api/v1/wx-gzh-tag';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createWxGzhTag,
  getWxGzhTagInfo,
  updateWxGzhTag,
} from '#/api/v1/wx-gzh-tag';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

interface Props {
  appId?: string;
}

const props = defineProps<Props>();
const emit = defineEmits(['success']);

const formData = ref<WxGzhTagInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['标签'])
    : $t('ui.actionTitle.create', ['标签']);
});

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 120,
  },
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
    // 设置 appId
    const appIdValue = props.appId || data.appId;
    if (!appIdValue) {
      message.error('缺少公众号信息');
      modalApi.unlock();
      return;
    }
    try {
      await (formData.value?.id
        ? updateWxGzhTag({
            body: {
              id: formData.value.id,
              name: data.name,
            } as UpdateWxGzhTagReq,
          })
        : createWxGzhTag({
            body: { appId: appIdValue, name: data.name } as CreateWxGzhTagReq,
          }));
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
    const data = modalApi.getData<{
      appId: string;
      row?: WxGzhTagInfo;
    }>();
    if (!data || !data.appId) {
      return;
    }

    // 设置 appId 到表单
    await formApi.setValues({ appId: data.appId });

    // 如果有 row 数据，说明是编辑模式
    if (data.row?.id) {
      modalApi.lock();
      try {
        const res = await getWxGzhTagInfo({ params: { id: data.row.id } });
        formData.value = res.info;
        // 设置到 values
        if (formData.value) {
          await formApi.setValues(formData.value);
        }
      } finally {
        modalApi.unlock();
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
