<script lang="ts" setup>
import type { MallPaymentRecordInfo } from '#/api/v1/mall-payment-record';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { Descriptions, DescriptionsItem } from 'ant-design-vue';

const formData = ref<MallPaymentRecordInfo>();

const getTitle = computed(() => {
  return `支付记录详情 - ${formData.value?.id || ''}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 获取数据
    const data = modalApi.getData<MallPaymentRecordInfo>();
    if (data) {
      formData.value = data;
    }
  },
});

// 支付渠道映射
const channelMap = {
  wechat: '微信',
  alipay: '支付宝',
};

// 支付方式映射
const methodMap = {
  mini_program: '小程序',
  h5: 'H5',
  native: '扫码',
  jsapi: 'JS API',
};

// 支付状态映射
const paymentStatusMap = {
  0: '待支付',
  1: '支付成功',
  2: '支付失败',
  3: '已退款',
};

// 记录状态映射
const statusMap = {
  '-1': '无效',
  '1': '正常',
};
</script>

<template>
  <Modal :title="getTitle" class="payment-detail-modal w-full max-w-4xl">
    <div v-if="formData" class="mx-4">
      <Descriptions bordered :column="2">
        <DescriptionsItem label="记录ID" :span="2">
          {{ formData.id }}
        </DescriptionsItem>
        <DescriptionsItem label="订单ID" :span="2">
          {{ formData.orderId }}
        </DescriptionsItem>
        <DescriptionsItem label="交易流水号" :span="2">
          {{ formData.transactionId }}
        </DescriptionsItem>
        <DescriptionsItem label="支付渠道">
          {{
            channelMap[formData.paymentChannel as keyof typeof channelMap] ||
            formData.paymentChannel
          }}
        </DescriptionsItem>
        <DescriptionsItem label="支付方式">
          {{
            methodMap[formData.paymentMethod as keyof typeof methodMap] ||
            formData.paymentMethod
          }}
        </DescriptionsItem>
        <DescriptionsItem label="支付金额">
          ¥{{ Number(formData.amount || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="币种">
          {{ formData.currency }}
        </DescriptionsItem>
        <DescriptionsItem label="支付状态">
          {{
            paymentStatusMap[
              formData.paymentStatus as keyof typeof paymentStatusMap
            ] || '未知'
          }}
        </DescriptionsItem>
        <DescriptionsItem label="记录状态">
          {{
            formData.status !== undefined
              ? statusMap[
                  formData.status.toString() as keyof typeof statusMap
                ] || '未知'
              : '-'
          }}
        </DescriptionsItem>
        <DescriptionsItem label="第三方订单号" :span="2">
          {{ formData.thirdPartyOrderNo || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="第三方交易号" :span="2">
          {{ formData.thirdPartyTransactionId || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="回调时间">
          {{ formData.callbackTime || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="创建时间">
          {{ formData.createdAt }}
        </DescriptionsItem>
        <DescriptionsItem label="更新时间">
          {{ formData.updatedAt }}
        </DescriptionsItem>
        <DescriptionsItem v-if="formData.errorCode" label="错误代码">
          <span class="text-red-600">{{ formData.errorCode }}</span>
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.errorMessage"
          label="错误信息"
          :span="2"
        >
          <span class="text-red-600">{{ formData.errorMessage }}</span>
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.callbackData"
          label="回调数据"
          :span="2"
        >
          <pre class="max-h-64 overflow-auto rounded bg-gray-50 p-3 text-sm">{{
            formData.callbackData
          }}</pre>
        </DescriptionsItem>
      </Descriptions>
    </div>
  </Modal>
</template>
