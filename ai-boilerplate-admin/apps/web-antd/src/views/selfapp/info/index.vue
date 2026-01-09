<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { GetSelfAppListParams, SelfAppInfo } from '#/api/v1/self-app';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteSelfApp,
  getSelfAppList,
  updateSelfAppStatus,
} from '#/api/v1/self-app';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import Form from './modules/form.vue';
import ReleaseModal from './modules/release-modal.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DetailModalComponent, detailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

const [ReleaseModalComponent, releaseModalApi] = useVbenModal({
  fullscreen: true,
  connectedComponent: ReleaseModal,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建自应用 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 查看自应用详情 */
function onView(row: SelfAppInfo) {
  detailModalApi.setData(row).open();
}

/** 编辑自应用 */
function onEdit(row: SelfAppInfo) {
  formModalApi.setData(row).open();
}

/** 版本管理 */
function onVersions(row: SelfAppInfo) {
  releaseModalApi.setData(row).open();
}

/** 删除自应用 */
async function onDelete(row: SelfAppInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSelfApp({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: SelfAppInfo) {
  try {
    await updateSelfAppStatus({
      body: {
        id: row.id!,
        status: newStatus,
      },
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
function onActionClick({ code, row }: OnActionClickParams<SelfAppInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'versions': {
      onVersions(row);
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
          return await getSelfAppList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetSelfAppListParams,
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
      refresh: true,
      'refresh-options': { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<SelfAppInfo>,
});
</script>

<template>
  <Page auto-content-height description="自应用信息管理">
    <FormModal @success="onRefresh" />
    <DetailModalComponent />
    <ReleaseModalComponent />
    <Grid>
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['self_app:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['自应用']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
