<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { UserMessageInfo } from '#/api/v1/user-message';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getUserMessageList } from '#/api/v1/user-message';

import { useGridColumns, useGridFormSchema } from './data';
import Detail from './modules/detail.vue';
import SendForm from './modules/send-form.vue';

const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,
});

const [SendModal, sendModalApi] = useVbenModal({
  connectedComponent: SendForm,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 发送消息 */
function onSend() {
  sendModalApi.open();
}

/** 查看详情 */
function onDetail(row: UserMessageInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<UserMessageInfo>) {
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
          return await getUserMessageList({
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
  } as VxeTableGridOptions<UserMessageInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <SendModal @success="onRefresh" />
    <DetailModal />
    <Grid table-title="用户消息">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onSend"
          v-access:code="['member:user-message:send']"
        >
          <Plus class="size-5" />
          发送消息
        </Button>
      </template>
    </Grid>
  </Page>
</template>
