<script lang="ts" setup>
import type { ArticleInfo } from '#/api/v1/article';
import type {
  CreateArticleReq,
  UpdateArticleReq,
} from '#/api/v1/article/types';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { createArticle, getArticleInfo, updateArticle } from '#/api/v1/article';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

const emit = defineEmits(['success']);
const formData = ref<ArticleInfo>();

const getTitle = computed(() => {
  return formData.value?.id
    ? $t('ui.actionTitle.edit', ['文章'])
    : $t('ui.actionTitle.create', ['文章']);
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
        const body: UpdateArticleReq = {
          id: values.id,
          title: values.title,
          summary: values.summary,
          coverImage: values.coverImage,
          contentMarkdown: values.contentMarkdown,
          isRecommend: values.isRecommend,
          isHot: values.isHot,
        };
        await updateArticle({ body });
      } else {
        const body: CreateArticleReq = {
          title: values.title,
          summary: values.summary,
          coverImage: values.coverImage,
          contentMarkdown: values.contentMarkdown,
          status: values.status,
          publishTime: values.publishTime,
          isRecommend: values.isRecommend,
          isHot: values.isHot,
        };
        await createArticle({ body });
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

    const data = modalApi.getData<ArticleInfo>();
    if (!data?.id) {
      await formApi.resetForm();
      return;
    }

    modalApi.lock();
    try {
      const res = await getArticleInfo({ params: { id: data.id } });
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
