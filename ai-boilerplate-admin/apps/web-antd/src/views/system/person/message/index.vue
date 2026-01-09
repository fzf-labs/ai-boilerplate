<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysNotifyMessageInfo } from '#/api/v1/sys-notify-message';

import { Page, useVbenModal } from '@vben/common-ui';
import { MdiCheckboxMarkedCircleOutline } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  getSysNotifyMessageMyUnreadList,
  updateSysNotifyMessageAllRead,
  updateSysNotifyMessageRead,
} from '#/api/v1/sys-notify-message';

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

/** 标记一条站内信已读 */
async function onRead(row: SysNotifyMessageInfo) {
  message.loading({
    content: '正在标记已读...',
    duration: 0,
    key: 'action_process_msg',
  });
  // 执行标记已读操作
  await updateSysNotifyMessageRead({ body: { ids: [row.id!] } });
  // 提示成功
  message.success({
    content: '标记已读成功',
    key: 'action_process_msg',
  });
  onRefresh();

  // 打开详情
  onDetail(row);
}

/** 标记选中的站内信为已读 */
async function onMarkRead() {
  const rows = gridApi.grid.getCheckboxRecords();
  if (!rows || rows.length === 0) {
    message.warning({
      content: '请选择需要标记的站内信',
      key: 'action_process_msg',
    });
    return;
  }

  const ids = rows
    .map((row: SysNotifyMessageInfo) => row.id!)
    .filter(Boolean) as string[];
  message.loading({
    content: '正在标记已读...',
    duration: 0,
    key: 'action_process_msg',
  });
  // 执行标记已读操作
  await updateSysNotifyMessageRead({ body: { ids } });
  // 提示成功
  message.success({
    content: '标记已读成功',
    key: 'action_process_msg',
  });
  await gridApi.grid.setAllCheckboxRow(false);
  onRefresh();
}

/** 标记所有站内信为已读 */
async function onMarkAllRead() {
  message.loading({
    content: '正在标记全部已读...',
    duration: 0,
    key: 'action_process_msg',
  });
  // 执行标记已读操作
  await updateSysNotifyMessageAllRead({ body: {} });
  // 提示成功
  message.success({
    content: '全部标记已读成功',
    key: 'action_process_msg',
  });
  await gridApi.grid.setAllCheckboxRow(false);
  onRefresh();
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
    case 'read': {
      onRead(row);
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
          return await getSysNotifyMessageMyUnreadList({
            options: {
              params: {
                page: page.currentPage,
                pageSize: page.pageSize,
                ...formValues,
              },
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
    checkboxConfig: {
      checkMethod: (params: { row: SysNotifyMessageInfo }) => {
        // 只允许选择未读的消息
        return !params.row.readTime || params.row.readTime === '';
      },
      highlight: true,
    },
  } as VxeTableGridOptions<SysNotifyMessageInfo>,
});
</script>
<template>
  <Page auto-content-height>
    <DetailModal @success="onRefresh" />
    <Grid table-title="我的站内信">
      <template #toolbar-tools>
        <Button type="primary" @click="onMarkRead">
          <MdiCheckboxMarkedCircleOutline />
          标记已读
        </Button>
        <Button type="primary" class="ml-2" @click="onMarkAllRead">
          <MdiCheckboxMarkedCircleOutline />
          全部已读
        </Button>
      </template>
    </Grid>
  </Page>
</template>
