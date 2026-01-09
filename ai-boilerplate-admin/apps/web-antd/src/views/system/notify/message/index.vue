<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysNotifyMessageInfo } from '#/api/v1/sys-notify-message';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysNotifyMessageList } from '#/api/v1/sys-notify-message';

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

/** 查看站内信详情 */
function onDetail(row: SysNotifyMessageInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row,
}: OnActionClickParams<SysNotifyMessageInfo>) {
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
          return await getSysNotifyMessageList({
            page: page.currentPage,
            pageSize: page.pageSize,
            ...formValues,
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
  } as VxeTableGridOptions<SysNotifyMessageInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <DetailModal @success="onRefresh" />
    <Grid table-title="通知信息列表" />
  </Page>
</template>
