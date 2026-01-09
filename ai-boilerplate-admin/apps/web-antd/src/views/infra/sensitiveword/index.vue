<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SensitiveWordInfo } from '#/api/v1/sensitive-word';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteSensitiveWord,
  getSensitiveWordList,
} from '#/api/v1/sensitive-word';
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

/** 创建敏感词 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑敏感词 */
function onEdit(row: SensitiveWordInfo) {
  formModalApi.setData(row).open();
}

/** 删除敏感词 */
async function onDelete(row: SensitiveWordInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.word]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSensitiveWord({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.word]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SensitiveWordInfo>) {
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
  },
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getSensitiveWordList({
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
  } as VxeTableGridOptions<SensitiveWordInfo>,
});
</script>

<template>
  <Page
    auto-content-height
    description="敏感词管理用于维护系统的敏感词库，支持敏感词的增删改查功能。"
  >
    <FormModal @success="onRefresh" />
    <Grid table-title="敏感词列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:sensitiveword:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['敏感词']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
