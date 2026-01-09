<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { MailLogInfo } from '#/api/v1/mail-log';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMailLogList } from '#/api/v1/mail-log';

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

/** 查看邮件日志 */
function onDetail(row: MailLogInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<MailLogInfo>) {
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
          return await getMailLogList({
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
  } as VxeTableGridOptions<MailLogInfo>,
});
</script>
<template>
  <Page auto-content-height>
    <DetailModal @success="onRefresh" />
    <Grid table-title="邮件日志列表">
      <template #toolbar-tools> </template>
    </Grid>
  </Page>
</template>
