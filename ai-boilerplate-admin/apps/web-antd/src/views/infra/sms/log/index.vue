<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SmsLogInfo } from '#/api/v1/sms-log';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSmsLogList } from '#/api/v1/sms-log';

import { useGridColumns, useGridFormSchema } from './data';
import Detail from './modules/detail.vue';

const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 查看短信日志详情 */
function onDetail(row: SmsLogInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SmsLogInfo>) {
  switch (code) {
    case 'detail': {
      onDetail(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
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
          return await getSmsLogList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            },
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<SmsLogInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <DetailModal @success="onRefresh" />
    <Grid table-title="短信日志列表" />
  </Page>
</template>
