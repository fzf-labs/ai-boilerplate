<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetWxGzhAccountListParams,
  WxGzhAccountInfo,
} from '#/api/v1/wx-gzh-account';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteWxGzhAccount,
  getWxGzhAccountList,
} from '#/api/v1/wx-gzh-account';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Detail from './modules/detail.vue';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建账号 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 查看详情 */
function onDetail(row: WxGzhAccountInfo) {
  detailModalApi.setData(row).open();
}

/** 编辑账号 */
function onEdit(row: WxGzhAccountInfo) {
  formModalApi.setData(row).open();
}

/** 删除账号 */
async function onDelete(row: WxGzhAccountInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    if (!row.id) return;
    await deleteWxGzhAccount({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<WxGzhAccountInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'detail': {
      onDetail(row);
      break;
    }
    case 'edit': {
      onEdit(row);
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
          return await getWxGzhAccountList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetWxGzhAccountListParams,
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
  } as VxeTableGridOptions<WxGzhAccountInfo>,
});
</script>

<template>
  <Page auto-content-height description="公众号账号管理">
    <FormModal @success="onRefresh" />
    <DetailModal />

    <Grid>
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['gzh:account:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['公众号账号']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
