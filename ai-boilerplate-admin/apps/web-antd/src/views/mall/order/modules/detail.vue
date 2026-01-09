<script lang="ts" setup>
import type { MallOrderInfo } from '#/api/v1/mall-order';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { Descriptions, DescriptionsItem } from 'ant-design-vue';

const formData = ref<MallOrderInfo>();

const getTitle = computed(() => {
  return `订单详情 - ${formData.value?.id || ''}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 获取数据
    const data = modalApi.getData<MallOrderInfo>();
    if (data) {
      formData.value = data;
    }
  },
});

// 商品类型映射
const productTypeMap = {
  membership: '会员',
  service: '服务',
  goods: '商品',
};

// 支付状态映射
const paymentStatusMap = {
  0: '待支付',
  1: '已支付',
  2: '支付失败',
  3: '已退款',
};

// 订单状态映射
const orderStatusMap = {
  pendingPayment: '待付款',
  pendingDelivery: '待发货',
  pendingReceipt: '待收货',
  completed: '已完成',
  canceled: '已取消',
  refunded: '已退款',
};
</script>

<template>
  <Modal :title="getTitle" class="order-detail-modal w-full max-w-4xl">
    <div v-if="formData" class="mx-4">
      <Descriptions bordered :column="2">
        <DescriptionsItem label="订单ID" :span="2">
          {{ formData.id }}
        </DescriptionsItem>
        <DescriptionsItem label="用户ID">
          {{ formData.userId }}
        </DescriptionsItem>
        <DescriptionsItem label="商品ID">
          {{ formData.productId }}
        </DescriptionsItem>
        <DescriptionsItem label="商品类型">
          {{
            productTypeMap[
              formData.productType as keyof typeof productTypeMap
            ] || formData.productType
          }}
        </DescriptionsItem>
        <DescriptionsItem label="币种">
          {{ formData.currency }}
        </DescriptionsItem>
        <DescriptionsItem label="原价">
          ¥{{ Number(formData.originalAmount || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="优惠金额">
          ¥{{ Number(formData.discountAmount || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="实付金额">
          ¥{{ Number(formData.actualAmount || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="退款金额">
          ¥{{ Number(formData.refundAmount || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="支付方式">
          {{ formData.paymentMethod }}
        </DescriptionsItem>
        <DescriptionsItem label="支付状态">
          {{
            paymentStatusMap[
              formData.paymentStatus as keyof typeof paymentStatusMap
            ] || '未知'
          }}
        </DescriptionsItem>
        <DescriptionsItem label="订单状态">
          {{
            orderStatusMap[formData.status as keyof typeof orderStatusMap] ||
            formData.status
          }}
        </DescriptionsItem>
        <DescriptionsItem label="支付时间">
          {{ formData.paymentTime || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="确认时间">
          {{ formData.deliveryTime || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="过期时间">
          {{ formData.expiredTime || '-' }}
        </DescriptionsItem>
        <DescriptionsItem label="创建时间">
          {{ formData.createdAt }}
        </DescriptionsItem>
        <DescriptionsItem label="更新时间">
          {{ formData.updatedAt }}
        </DescriptionsItem>
        <DescriptionsItem v-if="formData.remark" label="备注" :span="2">
          {{ formData.remark }}
        </DescriptionsItem>
      </Descriptions>
    </div>
  </Modal>
</template>
