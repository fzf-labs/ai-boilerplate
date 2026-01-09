<script lang="ts" setup>
import type { MallOrderInfo } from '#/api/v1/mall-order';
import type { MallPaymentRecordInfo } from '#/api/v1/mall-payment-record';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { Table, TableColumn } from 'ant-design-vue';

import { getMallPaymentRecordListByOrderId } from '#/api/v1/mall-payment-record';

const formData = ref<MallOrderInfo>();
const paymentRecords = ref<MallPaymentRecordInfo[]>([]);
const loading = ref(false);

const getTitle = computed(() => {
  return `支付记录 - ${formData.value?.id || ''}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      paymentRecords.value = [];
      return;
    }
    // 获取数据
    const data = modalApi.getData<MallOrderInfo>();
    if (data && data.id) {
      formData.value = data;
      await loadPaymentRecords(data.id);
    }
  },
});

// 加载支付记录
async function loadPaymentRecords(orderId: string) {
  loading.value = true;
  try {
    const res = await getMallPaymentRecordListByOrderId({
      params: { orderId },
    });
    paymentRecords.value = res.list || [];
  } finally {
    loading.value = false;
  }
}

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
const statusMap = {
  0: { color: 'orange', text: '待支付' },
  1: { color: 'green', text: '支付成功' },
  2: { color: 'red', text: '支付失败' },
  3: { color: 'purple', text: '已退款' },
};

// 记录状态映射
const recordStatusMap = {
  '-1': '无效',
  '1': '正常',
};
</script>

<template>
  <Modal :title="getTitle" :width="1200">
    <div v-if="formData" class="mx-4">
      <Table
        :data-source="paymentRecords"
        :loading="loading"
        :pagination="false"
        row-key="id"
      >
        <TableColumn key="id" title="记录ID" data-index="id" width="200">
          <template #default="{ text }">
            <span class="font-mono text-xs">{{ text }}</span>
          </template>
        </TableColumn>
        <TableColumn
          key="transactionId"
          title="交易流水号"
          data-index="transactionId"
          width="200"
        >
          <template #default="{ text }">
            <span class="font-mono text-xs">{{ text }}</span>
          </template>
        </TableColumn>
        <TableColumn
          key="paymentChannel"
          title="支付渠道"
          data-index="paymentChannel"
          width="100"
          align="center"
        >
          <template #default="{ text }">
            {{ channelMap[text as keyof typeof channelMap] || text }}
          </template>
        </TableColumn>
        <TableColumn
          key="paymentMethod"
          title="支付方式"
          data-index="paymentMethod"
          width="100"
          align="center"
        >
          <template #default="{ text }">
            {{ methodMap[text as keyof typeof methodMap] || text }}
          </template>
        </TableColumn>
        <TableColumn
          key="amount"
          title="支付金额"
          data-index="amount"
          width="120"
          align="right"
        >
          <template #default="{ text }">
            ¥{{ Number(text || 0).toFixed(2) }}
          </template>
        </TableColumn>
        <TableColumn
          key="currency"
          title="币种"
          data-index="currency"
          width="80"
          align="center"
        />
        <TableColumn
          key="paymentStatus"
          title="支付状态"
          data-index="paymentStatus"
          width="100"
          align="center"
        >
          <template #default="{ text }">
            <span
              :class="`inline-block rounded px-2 py-1 text-xs ${
                statusMap[text as keyof typeof statusMap]?.color === 'green'
                  ? 'bg-green-100 text-green-800'
                  : statusMap[text as keyof typeof statusMap]?.color ===
                      'orange'
                    ? 'bg-orange-100 text-orange-800'
                    : statusMap[text as keyof typeof statusMap]?.color === 'red'
                      ? 'bg-red-100 text-red-800'
                      : statusMap[text as keyof typeof statusMap]?.color ===
                          'purple'
                        ? 'bg-purple-100 text-purple-800'
                        : 'bg-gray-100 text-gray-800'
              }`"
            >
              {{ statusMap[text as keyof typeof statusMap]?.text || '未知' }}
            </span>
          </template>
        </TableColumn>
        <TableColumn
          key="thirdPartyOrderNo"
          title="第三方订单号"
          data-index="thirdPartyOrderNo"
          width="180"
        >
          <template #default="{ text }">
            <span class="font-mono text-xs">{{ text || '-' }}</span>
          </template>
        </TableColumn>
        <TableColumn
          key="thirdPartyTransactionId"
          title="第三方交易号"
          data-index="thirdPartyTransactionId"
          width="180"
        >
          <template #default="{ text }">
            <span class="font-mono text-xs">{{ text || '-' }}</span>
          </template>
        </TableColumn>
        <TableColumn
          key="status"
          title="记录状态"
          data-index="status"
          width="100"
          align="center"
        >
          <template #default="{ text }">
            <span
              :class="`inline-block rounded px-2 py-1 text-xs ${
                text === 1
                  ? 'bg-green-100 text-green-800'
                  : 'bg-red-100 text-red-800'
              }`"
            >
              {{
                recordStatusMap[
                  text.toString() as keyof typeof recordStatusMap
                ] || '未知'
              }}
            </span>
          </template>
        </TableColumn>
        <TableColumn
          key="callbackTime"
          title="回调时间"
          data-index="callbackTime"
          width="180"
        >
          <template #default="{ text }">
            {{ text || '-' }}
          </template>
        </TableColumn>
        <TableColumn
          key="createdAt"
          title="创建时间"
          data-index="createdAt"
          width="180"
        />
        <template #expandedRowRender="{ record }">
          <div class="rounded bg-gray-50 p-4">
            <div class="grid grid-cols-2 gap-4 text-sm">
              <div v-if="record.errorCode">
                <span class="font-semibold">错误代码:</span>
                <span class="ml-2 text-red-600">{{ record.errorCode }}</span>
              </div>
              <div v-if="record.errorMessage">
                <span class="font-semibold">错误信息:</span>
                <span class="ml-2 text-red-600">{{ record.errorMessage }}</span>
              </div>
              <div v-if="record.callbackData" class="col-span-2">
                <span class="font-semibold">回调数据:</span>
                <pre
                  class="mt-1 max-h-32 overflow-auto rounded bg-white p-2 text-xs"
                >
                  {{ record.callbackData }}
                </pre>
              </div>
            </div>
          </div>
        </template>
      </Table>
    </div>
  </Modal>
</template>
