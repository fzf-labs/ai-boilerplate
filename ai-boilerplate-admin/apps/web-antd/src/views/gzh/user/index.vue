<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetWxGzhUserListParams,
  WxGzhUserInfo,
} from '#/api/v1/wx-gzh-user';

import { Page, useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getWxGzhUserList } from '#/api/v1/wx-gzh-user';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModalComponent from './modules/detail.vue';

const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: DetailModalComponent,
  destroyOnClose: true,
});

/** 查看用户详情 */
function onDetail(row: WxGzhUserInfo) {
  detailModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<WxGzhUserInfo>) {
  switch (code) {
    case 'detail': {
      onDetail(row);
      break;
    }
  }
}

const [Grid] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
    submitOnChange: true,
    submitButtonOptions: {
      show: false,
    },
  },
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      autoLoad: false,
      ajax: {
        query: async ({ page }, formValues) => {
          return await getWxGzhUserList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetWxGzhUserListParams,
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
  } as VxeTableGridOptions<WxGzhUserInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <DetailModal />

    <Grid table-title="公众号粉丝列表" />
  </Page>
</template>
