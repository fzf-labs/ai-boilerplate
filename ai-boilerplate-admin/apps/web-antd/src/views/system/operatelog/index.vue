<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysOperateLogInfo } from '#/api/v1/sys-operate-log';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysOperateLogList } from '#/api/v1/sys-operate-log';

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

/** 查看操作日志详情 */
function onDetail(row: SysOperateLogInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SysOperateLogInfo>) {
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
          return await getSysOperateLogList({
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
  } as VxeTableGridOptions<SysOperateLogInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <DetailModal @success="onRefresh" />
    <Grid table-title="操作日志列表">
      <template #toolbar-tools> </template>
    </Grid>
  </Page>
</template>
