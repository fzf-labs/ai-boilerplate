<script lang="ts" setup>
import type { MallActivationCodeInfo } from '#/api/v1/mall-activation-code';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  batchGenerateMallActivationCode,
  getMallActivationCodeInfo,
  updateMallActivationCode,
} from '#/api/v1/mall-activation-code';
import { $t } from '#/locales';

import { useCreateFormSchema, useEditFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<MallActivationCodeInfo>();
const isEdit = ref(false);
const getTitle = computed(() => {
  return isEdit.value
    ? $t('ui.actionTitle.edit', ['激活码'])
    : $t('ui.actionTitle.create', ['激活码']);
});

// 创建模式表单
const [CreateForm, createFormApi] = useVbenForm({
  layout: 'horizontal',
  schema: useCreateFormSchema(),
  showDefaultActions: false,
});

// 编辑模式表单
const [EditForm, editFormApi] = useVbenForm({
  layout: 'horizontal',
  schema: useEditFormSchema(),
  showDefaultActions: false,
});

// 当前使用的表单 API
const formApi = computed(() => (isEdit.value ? editFormApi : createFormApi));

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const currentFormApi = formApi.value;
    const { valid } = await currentFormApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    // 提交表单
    const rawData = (await currentFormApi.getValues()) as any;

    try {
      if (isEdit.value) {
        // 编辑模式 - 更新激活码
        const data: any = {
          id: rawData.id,
          platform: rawData.platform,
          platformSoldAt: rawData.platformSoldAt,
          platformOrderNo: rawData.platformOrderNo,
          platformBuyerId: rawData.platformBuyerId,
          platformBuyerName: rawData.platformBuyerName,
          remark: rawData.remark,
        };
        await updateMallActivationCode({ body: data });
      } else {
        // 新增模式 - 批量生成激活码
        const data: any = {
          productType: rawData.productType,
          productId: rawData.productId,
          validSt: rawData.validSt,
          validEd: rawData.validEd,
          num: Number(rawData.num) || 1,
          platform: rawData.platform,
          remark: rawData.remark,
        };
        const result = await batchGenerateMallActivationCode({ body: data });
        // 显示批次号信息
        message.success({
          content: `成功生成 ${data.num} 个激活码，批次号：${result.batchNo}`,
          duration: 5,
          key: 'action_process_msg',
        });
      }
      // 关闭并提示
      await modalApi.close();
      emit('success');
      if (isEdit.value) {
        message.success({
          content: $t('ui.actionMessage.operationSuccess'),
          key: 'action_process_msg',
        });
      }
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      // 关闭时重置
      formData.value = undefined;
      isEdit.value = false;
      await createFormApi.resetForm();
      await editFormApi.resetForm();
      return;
    }
    // 加载数据
    const data = modalApi.getData<MallActivationCodeInfo>();
    if (!data || !data.id) {
      // 新增模式
      isEdit.value = false;
      formData.value = undefined;
      await createFormApi.resetForm();
      return;
    }
    // 编辑模式
    isEdit.value = true;
    modalApi.lock();
    try {
      // 加载数据
      const res = await getMallActivationCodeInfo({ params: { id: data.id } });
      if (!res.info) {
        return;
      }
      formData.value = res.info;

      // 处理数据格式，转换为表单可用的格式
      const formValues = {
        id: formData.value.id,
        platform: formData.value.platform,
        platformSoldAt: formData.value.platformSoldAt,
        platformOrderNo: formData.value.platformOrderNo,
        platformBuyerId: formData.value.platformBuyerId,
        platformBuyerName: formData.value.platformBuyerName,
        remark: formData.value.remark,
      };

      // 设置到编辑表单
      await editFormApi.setValues(formValues);
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal :title="getTitle" :width="1000">
    <CreateForm v-if="!isEdit" class="mx-4" />
    <EditForm v-else class="mx-4" />
  </Modal>
</template>
