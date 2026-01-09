<script lang="ts" setup>
import type { MailLogInfo } from '#/api/v1/mail-log';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Descriptions } from 'ant-design-vue';
import DOMPurify from 'dompurify';

const formData = ref<MailLogInfo>();

const sanitizedContent = computed(() => {
  if (!formData.value?.templateContent) return '';
  return DOMPurify.sanitize(formData.value.templateContent);
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<MailLogInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      formData.value = data;
    } finally {
      modalApi.lock(false);
    }
  },
});
</script>

<template>
  <Modal
    title="邮件日志详情"
    class="mail-log-detail-modal w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div class="p-4">
      <Descriptions :column="2" bordered :label-style="{ width: '140px' }">
        <Descriptions.Item label="编号">{{ formData?.id }}</Descriptions.Item>
        <Descriptions.Item label="创建时间">
          {{ formatDateTime(formData?.createdAt || '') }}
        </Descriptions.Item>
        <Descriptions.Item label="收件邮箱">
          {{ formData?.toMail }}
        </Descriptions.Item>
        <Descriptions.Item label="发送邮箱">
          {{ formData?.fromMail }}
        </Descriptions.Item>
        <Descriptions.Item label="模板编号">
          {{ formData?.templateId }}
        </Descriptions.Item>
        <Descriptions.Item label="模板编码">
          {{ formData?.templateCode }}
        </Descriptions.Item>
        <Descriptions.Item label="邮件标题" :span="2">
          {{ formData?.templateTitle }}
        </Descriptions.Item>
        <Descriptions.Item label="邮件内容" :span="2">
          <!-- eslint-disable-next-line vue/no-v-html -->
          <div v-html="sanitizedContent"></div>
        </Descriptions.Item>
        <Descriptions.Item label="发送状态">
          {{ formData?.sendStatus }}
        </Descriptions.Item>
        <Descriptions.Item label="发送时间">
          {{ formatDateTime(formData?.sendTime || '') }}
        </Descriptions.Item>
        <Descriptions.Item label="发送消息编号">
          {{ formData?.sendMessageId }}
        </Descriptions.Item>
        <Descriptions.Item label="发送异常">
          {{ formData?.sendException }}
        </Descriptions.Item>
      </Descriptions>
    </div>
  </Modal>
</template>
