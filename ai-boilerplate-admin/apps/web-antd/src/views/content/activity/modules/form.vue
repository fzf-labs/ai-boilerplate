<script lang="ts" setup>
import type { ActivityInfo } from '#/api/v1/activity';
import type {
  CreateActivityReq,
  UpdateActivityReq,
} from '#/api/v1/activity/types';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  createActivity,
  getActivityInfo,
  updateActivity,
} from '#/api/v1/activity';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<ActivityInfo>();

const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['活动'])
    : $t('ui.actionTitle.create', ['活动']);
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
    try {
      const values = (await formApi.getValues()) as Record<string, any>;

      if (formData.value?.id) {
        const body: UpdateActivityReq = {
          id: values.id,
          title: values.title,
          imageURL: values.imageURL,
          linkURL: values.linkURL,
          linkType: values.linkType,
          sort: values.sort,
          status: values.status,
          startTime: values.startTime,
          endTime: values.endTime,
        };
        await updateActivity({ body });
      } else {
        const body: CreateActivityReq = {
          title: values.title,
          imageURL: values.imageURL,
          linkURL: values.linkURL,
          linkType: values.linkType,
          sort: values.sort,
          status: values.status,
          startTime: values.startTime,
          endTime: values.endTime,
        };
        await createActivity({ body });
      }

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

    const data = modalApi.getData<ActivityInfo>();
    if (!data?.id) {
      await formApi.resetForm();
      await formApi.setValues({ status: 1, sort: 0, linkType: 'internal' });
      return;
    }

    modalApi.lock();
    try {
      const res = await getActivityInfo({ params: { id: data.id } });
      formData.value = res.info;
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
  <Modal :title="getTitle" class="w-2/3">
    <Form class="mx-4" />
  </Modal>
</template>
