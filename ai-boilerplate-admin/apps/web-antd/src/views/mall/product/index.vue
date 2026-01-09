<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetMallProductListParams,
  MallProductInfo,
} from '#/api/v1/mall-product';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteMallProduct,
  getMallProductList,
  updateMallProductStatus,
} from '#/api/v1/mall-product';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DetailModalComponent, detailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建商品 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 查看商品详情 */
function onView(row: MallProductInfo) {
  detailModalApi.setData(row).open();
}

/** 编辑商品 */
function onEdit(row: MallProductInfo) {
  formModalApi.setData(row).open();
}

/** 删除商品 */
async function onDelete(row: MallProductInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.productName]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMallProduct({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.productName]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: MallProductInfo) {
  try {
    await updateMallProductStatus({ body: { id: row.id!, status: newStatus } });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
    return true;
  } catch {
    return false;
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<MallProductInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'view': {
      onView(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick, onStatusChange),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getMallProductList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetMallProductListParams,
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
  } as VxeTableGridOptions<MallProductInfo>,
});
</script>

<template>
  <Page
    auto-content-height
    description="商品管理是商城系统的核心模块，支持商品的增删改查、状态管理等功能。"
  >
    <FormModal @success="onRefresh" />
    <DetailModalComponent />

    <Grid table-title="商品列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['mall:product:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['商品']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
