<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { MallPaymentRecordInfo } from '#/api/v1/mall-payment-record';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMallPaymentRecordListByOrderId } from '#/api/v1/mall-payment-record';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';

const [DetailModalComponent, detailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

/** 查看支付记录详情 */
function onView(row: MallPaymentRecordInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row,
}: OnActionClickParams<MallPaymentRecordInfo>) {
  switch (code) {
    case 'view': {
      onView(row);
      break;
    }
  }
}

const [Grid] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page: _page }, formValues) => {
          // For demo purposes, using existing API if orderId is provided
          if (formValues.orderId) {
            return await getMallPaymentRecordListByOrderId({
              params: { orderId: formValues.orderId },
            });
          }
          // Return empty result for now
          return { list: [], total: 0 };
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    showOverflow: 'tooltip',
    size: 'small',
    stripe: true,
    border: true,
    toolbarConfig: {
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<MallPaymentRecordInfo>,
});
</script>

<template>
  <Page
    auto-content-height
    description="支付记录管理用于根据订单ID查询相关的支付记录信息。"
  >
    <DetailModalComponent />

    <Grid table-title="支付记录列表" />
  </Page>
</template>
