<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createMembershipBenefit,
  getMembershipBenefitInfo,
  updateMembershipBenefit,
} from '#/api/v1/membership-benefit';
import { $t } from '#/locales';

import { useBenefitFormSchema } from './benefit-data';

const emit = defineEmits<{
  success: [];
}>();

const formData = ref<any>();

const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['权益'])
    : $t('ui.actionTitle.create', ['权益']);
});

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useBenefitFormSchema(),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();

    try {
      // 提交表单
      const formValues = await formApi.getValues();
      await (formData.value?.id
        ? updateMembershipBenefit(formValues)
        : createMembershipBenefit(formValues));

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
    const data = modalApi.getData<>();
    if (!data) {
      return;
    }

    if (data.id) {
      // 编辑模式：加载详细信息
      modalApi.lock();
      try {
        const res = await getMembershipBenefitInfo(data.id);
        if (!res.info) {
          return;
        }
        formData.value = res.info;
        await formApi.setValues(formData.value);
      } finally {
        modalApi.lock(false);
      }
    } else {
      // 新建模式：设置会员类型
      formData.value = undefined;
      await formApi.setValues({
        membershipType: data.membershipType,
        status: 1, // 默认启用状态
      });
    }
  },
});
</script>

<template>
  <Modal :title="getTitle">
    <Form class="mx-4" />
  </Modal>
</template>
