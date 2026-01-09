<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { AiProviderPlatformApi } from '#/api/ai/manager/providerPlatform';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteAiProviderPlatform,
  getAiProviderPlatformList,
} from '#/api/ai/manager/providerPlatform';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import ModelList from './modules/model-list.vue';
import PlatformForm from './modules/platform-form.vue';

const [PlatformFormModal, platformFormModalApi] = useVbenModal({
  connectedComponent: PlatformForm,
  destroyOnClose: true,
});

const [ModelListModal, modelListModalApi] = useVbenModal({
  fullscreen: true,
  connectedComponent: ModelList,
  destroyOnClose: true,
  showConfirmButton: false,
  showCancelButton: false,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建服务商平台 */
function onCreate() {
  platformFormModalApi.setData(null).open();
}

/** 编辑服务商平台 */
function onEdit(row: AiProviderPlatformApi.AiProviderPlatformInfo) {
  platformFormModalApi.setData(row).open();
}

/** 删除服务商平台 */
async function onDelete(row: AiProviderPlatformApi.AiProviderPlatformInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiProviderPlatform({ id: row.id });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 管理模型 */
function onManageModels(row: AiProviderPlatformApi.AiProviderPlatformInfo) {
  modelListModalApi.setData(row).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row,
}: OnActionClickParams<AiProviderPlatformApi.AiProviderPlatformInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'manageModels': {
      onManageModels(row);
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
          return await getAiProviderPlatformList({
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
      refresh: true,
      search: true,
    },
  } as VxeTableGridOptions<AiProviderPlatformApi.AiProviderPlatformInfo>,
});
</script>

<template>
  <Page auto-content-height description="AI服务商及模型管理">
    <PlatformFormModal @success="onRefresh" />
    <ModelListModal />
    <Grid>
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['ai:manager:provider:platform:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['服务商']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
