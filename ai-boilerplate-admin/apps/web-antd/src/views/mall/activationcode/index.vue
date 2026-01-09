<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetMallActivationCodeListParams,
  MallActivationCodeInfo,
} from '#/api/v1/mall-activation-code';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteMallActivationCode,
  getMallActivationCodeList,
  updateMallActivationCodeStatus,
} from '#/api/v1/mall-activation-code';
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

/** 创建激活码 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 查看激活码详情 */
function onView(row: MallActivationCodeInfo) {
  detailModalApi.setData(row).open();
}

/** 编辑激活码 */
function onEdit(row: MallActivationCodeInfo) {
  formModalApi.setData(row).open();
}

/** 删除激活码 */
async function onDelete(row: MallActivationCodeInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.code]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMallActivationCode({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.code]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: MallActivationCodeInfo) {
  try {
    await updateMallActivationCodeStatus({
      body: { id: row.id!, status: newStatus },
    });
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
function onActionClick({
  code,
  row,
}: OnActionClickParams<MallActivationCodeInfo>) {
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
          return await getMallActivationCodeList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetMallActivationCodeListParams,
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
  } as VxeTableGridOptions<MallActivationCodeInfo>,
});
</script>

<template>
  <Page
    auto-content-height
    description="激活码管理是商城系统的核心模块，支持激活码的增删改查、状态管理等功能。"
  >
    <FormModal @success="onRefresh" />
    <DetailModalComponent />

    <Grid table-title="激活码列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['mall:activationcode:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['激活码']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
