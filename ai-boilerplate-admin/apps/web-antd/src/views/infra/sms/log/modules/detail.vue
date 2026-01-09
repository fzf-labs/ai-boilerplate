<script lang="ts" setup>
import type { SmsLogInfo } from '#/api/v1/sms-log';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { Descriptions } from 'ant-design-vue';

const formData = ref<SmsLogInfo>();

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载数据
    const data = modalApi.getData<SmsLogInfo>();
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
    title="短信日志详情"
    class="sms-log-detail-modal w-full max-w-4xl"
    :show-cancel-button="false"
    :show-confirm-button="false"
  >
    <Descriptions
      bordered
      :column="2"
      size="middle"
      class="mx-4"
      :label-style="{ width: '140px' }"
    >
      <Descriptions.Item label="用户ID">
        {{ formData?.userId }}
      </Descriptions.Item>
      <Descriptions.Item label="手机号">
        {{ formData?.mobile }}
      </Descriptions.Item>
      <Descriptions.Item label="短信渠道">
        {{ formData?.smsChannelId }}
      </Descriptions.Item>
      <Descriptions.Item label="短信模板">
        {{ formData?.smsTemplateId }}
      </Descriptions.Item>
      <Descriptions.Item label="短信内容">
        {{ formData?.smsParamsContent }}
      </Descriptions.Item>
      <Descriptions.Item label="发送状态">
        {{ formData?.sendStatus }}
      </Descriptions.Item>
      <Descriptions.Item label="发送时间">
        {{ formatDateTime(formData?.sendTime || '') }}
      </Descriptions.Item>
      <Descriptions.Item label="接收状态">
        {{ formData?.receiveStatus }}
      </Descriptions.Item>
      <Descriptions.Item label="接收时间">
        {{ formatDateTime(formData?.receiveTime || '') }}
      </Descriptions.Item>
      <Descriptions.Item label="API 发送编码">
        {{ formData?.apiSendCode }}
      </Descriptions.Item>
      <Descriptions.Item label="API 发送消息">
        {{ formData?.apiSendMsg }}
      </Descriptions.Item>
      <Descriptions.Item label="API 接收编码">
        {{ formData?.apiReceiveCode }}
      </Descriptions.Item>
      <Descriptions.Item label="API 接收消息">
        {{ formData?.apiReceiveMsg }}
      </Descriptions.Item>
      <Descriptions.Item label="API 请求 ID">
        {{ formData?.apiRequestId }}
      </Descriptions.Item>
      <Descriptions.Item label="API 序列号">
        {{ formData?.apiSerialNo }}
      </Descriptions.Item>
    </Descriptions>
  </Modal>
</template>
