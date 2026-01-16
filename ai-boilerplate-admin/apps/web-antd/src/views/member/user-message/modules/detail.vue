<script lang="ts" setup>
import type { UserMessageInfo } from '#/api/v1/user-message';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Descriptions, DescriptionsItem, Image } from 'ant-design-vue';

import { getUserMessageInfo } from '#/api/v1/user-message';

const formData = ref<UserMessageInfo>();

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    const data = modalApi.getData<UserMessageInfo>();
    if (!data?.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getUserMessageInfo({ params: { id: data.id } });
      formData.value = res.info;
    } finally {
      modalApi.lock(false);
    }
  },
});

const categoryText = computed(() => {
  switch (formData.value?.category) {
    case 'transaction':
      return '交易信息';
    case 'system':
      return '系统消息';
    case 'service':
      return '客服消息';
    default:
      return formData.value?.category || '-';
  }
});

const audienceTypeText = computed(() => {
  switch (formData.value?.audienceType) {
    case 'all':
      return '全部用户';
    case 'segment':
      return '条件筛选';
    case 'users':
      return '指定用户';
    default:
      return formData.value?.audienceType || '-';
  }
});

const audienceValueText = computed(() => {
  const raw = formData.value?.audienceValue?.trim();
  if (!raw) {
    return '-';
  }
  try {
    const parsed = JSON.parse(raw);
    return JSON.stringify(parsed, null, 2);
  } catch {
    return raw;
  }
});
</script>

<template>
  <Modal
    title="消息详情"
    class="w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <div v-if="formData" class="p-4">
      <Descriptions bordered :column="2">
        <DescriptionsItem label="消息ID">
          {{ formData.id || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="批次ID">
          {{ formData.messageId || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="分类">
          {{ categoryText }}
        </DescriptionsItem>
        <DescriptionsItem label="标题">
          {{ formData.title || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="摘要" :span="2">
          {{ formData.summary || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="封面图" :span="2">
          <Image
            v-if="formData.coverURL"
            :src="formData.coverURL"
            :width="120"
            :height="120"
          />
          <span v-else>-</span>
        </DescriptionsItem>
        <DescriptionsItem label="内容" :span="2">
          <div class="whitespace-pre-wrap">
            {{ formData.content || '-' }}
          </div>
        </DescriptionsItem>
        <DescriptionsItem label="跳转链接" :span="2">
          {{ formData.linkURL || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="投放范围">
          {{ audienceTypeText }}
        </DescriptionsItem>
        <DescriptionsItem label="投放条件" :span="2">
          <pre class="whitespace-pre-wrap text-xs">{{ audienceValueText }}</pre>
        </DescriptionsItem>
        <DescriptionsItem label="用户ID">
          {{ formData.userId || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="创建人">
          {{ formData.adminId || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="发送时间">
          {{ formatDateTime(formData.sentAt || '') || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="阅读时间">
          {{ formatDateTime(formData.readAt || '') || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="创建时间">
          {{ formatDateTime(formData.createdAt || '') || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="更新时间">
          {{ formatDateTime(formData.updatedAt || '') || '-' }}
        </DescriptionsItem>
      </Descriptions>
    </div>
  </Modal>
</template>
