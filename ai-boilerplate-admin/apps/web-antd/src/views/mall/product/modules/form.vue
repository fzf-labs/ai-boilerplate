<script lang="ts" setup>
import type {
  CreateMallProductReq,
  MallProductInfo,
  UpdateMallProductReq,
} from '#/api/v1/mall-product';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createMallProduct,
  getMallProductInfo,
  updateMallProduct,
} from '#/api/v1/mall-product';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<MallProductInfo>();
const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['商品'])
    : $t('ui.actionTitle.create', ['商品']);
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
    const data: CreateMallProductReq | UpdateMallProductReq = {
      ...rawData,
      // 确保数字类型字段正确
      originalPrice: Number(rawData.originalPrice),
      currentPrice: Number(rawData.currentPrice),
      stockQuantity: rawData.stockQuantity
        ? Number(rawData.stockQuantity)
        : undefined,
      soldQuantity: rawData.soldQuantity
        ? Number(rawData.soldQuantity)
        : undefined,
      sort: rawData.sort ? Number(rawData.sort) : undefined,
      status: Number(rawData.status),
    };

    try {
      await (formData.value?.id
        ? updateMallProduct({ body: data as UpdateMallProductReq })
        : createMallProduct({ body: data as CreateMallProductReq }));
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
    const data = modalApi.getData<MallProductInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getMallProductInfo({ params: { id: data.id } });
      if (!res.info) {
        return;
      }
      formData.value = res.info;
      // 处理数据格式，转换为表单可用的格式
      const formValues = {
        ...formData.value,
        // 确保数字类型字段正确转换
        originalPrice: Number(formData.value.originalPrice),
        currentPrice: Number(formData.value.currentPrice),
        stockQuantity: formData.value.stockQuantity
          ? Number(formData.value.stockQuantity)
          : undefined,
        soldQuantity: formData.value.soldQuantity
          ? Number(formData.value.soldQuantity)
          : undefined,
        sort: formData.value.sort ? Number(formData.value.sort) : undefined,
        status: Number(formData.value.status),
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
