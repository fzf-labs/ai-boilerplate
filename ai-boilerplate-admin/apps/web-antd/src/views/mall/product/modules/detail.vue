<script lang="ts" setup>
import type { MallProductInfo } from '#/api/v1/mall-product';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { Descriptions, DescriptionsItem } from 'ant-design-vue';

const formData = ref<MallProductInfo>();

const getTitle = computed(() => {
  return `商品详情 - ${formData.value?.productName || ''}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 获取数据
    const data = modalApi.getData<MallProductInfo>();
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

// 状态映射
const statusMap = {
  '-1': '下架',
  '0': '待上架',
  '1': '在售',
  '2': '售罄',
};
</script>

<template>
  <Modal :title="getTitle" class="product-detail-modal w-full max-w-4xl">
    <div v-if="formData" class="mx-4">
      <Descriptions bordered :column="2">
        <DescriptionsItem label="商品名称">
          {{ formData.productName }}
        </DescriptionsItem>
        <DescriptionsItem label="商品类型">
          {{
            productTypeMap[
              formData.productType as keyof typeof productTypeMap
            ] || formData.productType
          }}
        </DescriptionsItem>
        <DescriptionsItem label="原价">
          ¥{{ Number(formData.originalPrice || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="现价">
          ¥{{ Number(formData.currentPrice || 0).toFixed(2) }}
        </DescriptionsItem>
        <DescriptionsItem label="库存数量">
          {{
            formData.stockQuantity === -1 ? '无限' : formData.stockQuantity || 0
          }}
        </DescriptionsItem>
        <DescriptionsItem label="已售数量">
          {{ formData.soldQuantity || 0 }}
        </DescriptionsItem>
        <DescriptionsItem label="排序">
          {{ formData.sort || 0 }}
        </DescriptionsItem>
        <DescriptionsItem label="状态">
          {{
            statusMap[formData.status.toString() as keyof typeof statusMap] ||
            '未知'
          }}
        </DescriptionsItem>
        <DescriptionsItem label="创建时间">
          {{ formData.createdAt }}
        </DescriptionsItem>
        <DescriptionsItem label="更新时间">
          {{ formData.updatedAt }}
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.productDesc"
          label="商品描述"
          :span="2"
        >
          {{ formData.productDesc }}
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.productImages && formData.productImages.length > 0"
          label="商品图片"
          :span="2"
        >
          <div class="flex flex-wrap gap-2">
            <img
              v-for="(image, index) in formData.productImages"
              :key="index"
              :src="image"
              :alt="`商品图片${index + 1}`"
              class="h-20 w-20 rounded border object-cover"
            />
          </div>
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.productDetail && formData.productDetail.length > 0"
          label="商品详情"
          :span="2"
        >
          <div class="space-y-1">
            <div
              v-for="(detail, index) in formData.productDetail"
              :key="index"
              class="text-sm"
            >
              {{ detail }}
            </div>
          </div>
        </DescriptionsItem>
        <DescriptionsItem
          v-if="formData.productConfig"
          label="商品配置"
          :span="2"
        >
          <pre class="rounded bg-gray-50 p-2 text-sm">{{
            JSON.stringify(formData.productConfig, null, 2)
          }}</pre>
        </DescriptionsItem>
      </Descriptions>
    </div>
  </Modal>
</template>
