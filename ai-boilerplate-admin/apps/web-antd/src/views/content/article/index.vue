<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { ArticleInfo } from '#/api/v1/article';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteArticle,
  getArticleList,
  updateArticleStatus,
} from '#/api/v1/article';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

function onRefresh() {
  gridApi.query();
}

function onCreate() {
  formModalApi.setData(null).open();
}

function onEdit(row: ArticleInfo) {
  formModalApi.setData(row).open();
}

async function onDelete(row: ArticleInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.title]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteArticle({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.title]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

async function onPublish(row: ArticleInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.processing', ['发布']),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await updateArticleStatus({ body: { id: row.id!, status: 1 } });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

async function onUnpublish(row: ArticleInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.processing', ['下线']),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await updateArticleStatus({ body: { id: row.id!, status: -1 } });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

function onActionClick({ code, row }: OnActionClickParams<ArticleInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'publish': {
      onPublish(row);
      break;
    }
    case 'unpublish': {
      onUnpublish(row);
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
          return await getArticleList({
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
  } as VxeTableGridOptions<ArticleInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="文章管理">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['content:article:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['文章']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
