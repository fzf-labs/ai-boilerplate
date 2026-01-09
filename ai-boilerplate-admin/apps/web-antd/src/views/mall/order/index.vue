<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetMallOrderListParams,
  MallOrderInfo,
} from '#/api/v1/mall-order';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMallOrderList } from '#/api/v1/mall-order';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import PaymentModal from './modules/payment.vue';

const [DetailModalComponent, detailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

const [PaymentModalComponent, paymentModalApi] = useVbenModal({
  connectedComponent: PaymentModal,
  destroyOnClose: true,
});

/** 查看订单详情 */
function onView(row: MallOrderInfo) {
  detailModalApi.setData(row).open();
}

/** 查看支付记录 */
function onPayment(row: MallOrderInfo) {
  paymentModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<MallOrderInfo>) {
  switch (code) {
    case 'payment': {
      onPayment(row);
      break;
    }
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
        query: async ({ page }, formValues) => {
          return await getMallOrderList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetMallOrderListParams,
          });
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
  } as VxeTableGridOptions<MallOrderInfo>,
});
</script>

<template>
  <Page
    auto-content-height
    description="订单管理用于查看和管理所有订单信息，包括订单状态、支付记录等。"
  >
    <DetailModalComponent />
    <PaymentModalComponent />

    <Grid table-title="订单列表" />
  </Page>
</template>
