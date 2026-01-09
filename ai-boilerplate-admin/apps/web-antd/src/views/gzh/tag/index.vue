<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { GetWxGzhTagListParams, WxGzhTagInfo } from '#/api/v1/wx-gzh-tag';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus, RotateCw } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteWxGzhTag,
  getWxGzhTagList,
  syncWxGzhTag,
} from '#/api/v1/wx-gzh-tag';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建标签 */
async function onCreate() {
  // 获取当前搜索表单的 appId
  const formValues = await gridApi.formApi.getValues();
  const currentAppId = formValues?.appId;

  formModalApi.setData({ appId: currentAppId }).open();
}

/** 编辑标签 */
function onEdit(row: WxGzhTagInfo) {
  formModalApi.setData({ row }).open();
}

/** 删除标签 */
async function onDelete(row: WxGzhTagInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    key: 'action_key_msg',
  });
  try {
    await deleteWxGzhTag({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 同步标签 */
async function onSync() {
  const formValues = await gridApi.formApi.getValues();
  const currentAppId = formValues?.appId;

  if (!currentAppId) {
    message.warning('请先选择公众号');
    return;
  }

  const hideLoading = message.loading({
    content: '正在同步标签...',
    key: 'sync_key_msg',
  });
  try {
    await syncWxGzhTag({ params: { appId: currentAppId } });
    message.success({
      content: '标签同步成功',
      key: 'sync_key_msg',
    });
    onRefresh();
  } catch {
    message.error({
      content: '标签同步失败',
      key: 'sync_key_msg',
    });
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<WxGzhTagInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
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
          return await getWxGzhTagList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetWxGzhTagListParams,
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
  } as VxeTableGridOptions<WxGzhTagInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="标签列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['gzh:tag:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['标签']) }}
        </Button>
        <Button @click="onSync" v-access:code="['gzh:tag:sync']">
          <RotateCw class="size-5" />
          同步标签
        </Button>
      </template>
    </Grid>
  </Page>
</template>
