<script lang="ts" setup>
import type {
  CreateMembershipBenefitTypeReq,
  MembershipBenefitTypeInfo,
  UpdateMembershipBenefitTypeReq,
} from '#/api/v1/membership-benefit-type';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { useVbenForm } from '#/adapter/form';
import {
  createMembershipBenefitType,
  updateMembershipBenefitType,
} from '#/api/v1/membership-benefit-type';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits<{ success: [] }>();

const formData = ref<MembershipBenefitTypeInfo>();
const getIsUpdate = computed(() => !!formData.value?.id);
const getModalTitle = computed(() => {
  return getIsUpdate.value
    ? $t('ui.actionTitle.edit', ['权益类型'])
    : $t('ui.actionTitle.create', ['权益类型']);
});

const [Form, formApi] = useVbenForm({
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
    formItemClass: 'col-span-2',
    labelWidth: 100,
  },
  handleSubmit: onSubmit,
  layout: 'horizontal',
  schema: useFormSchema(),
  showDefaultActions: false,
  wrapperClass: 'grid-cols-2',
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (valid) {
      modalApi.lock();
      await formApi.submitForm();
      modalApi.close();
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    const data = modalApi.getData<MembershipBenefitTypeInfo>();
    if (!data) {
      await formApi.resetForm();
      return;
    }
    formData.value = data;
    await formApi.setValues({
      ...data,
    });
  },
});

async function onSubmit(values: Record<string, any>) {
  const baseValues = values as CreateMembershipBenefitTypeReq;
  if (getIsUpdate.value) {
    const id = formData.value?.id;
    if (!id) {
      return;
    }
    const body: UpdateMembershipBenefitTypeReq = { ...baseValues, id };
    await updateMembershipBenefitType({ body });
  } else {
    await createMembershipBenefitType({ body: baseValues });
  }
  emit('success');
}
</script>

<template>
  <Modal :title="getModalTitle">
    <Form class="mx-4" />
  </Modal>
</template>
