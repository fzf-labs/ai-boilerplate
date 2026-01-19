<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { ActivityInfo } from '#/api/v1/activity';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteActivity,
  getActivityList,
  updateActivityStatus,
} from '#/api/v1/activity';
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

function onEdit(row: ActivityInfo) {
  formModalApi.setData(row).open();
}

async function onDelete(row: ActivityInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.title]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteActivity({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.title]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

async function onEnable(row: ActivityInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.processing', ['启用']),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await updateActivityStatus({ body: { id: row.id!, status: 1 } });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

async function onDisable(row: ActivityInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.processing', ['禁用']),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await updateActivityStatus({ body: { id: row.id!, status: -1 } });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

function onActionClick({ code, row }: OnActionClickParams<ActivityInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'disable': {
      onDisable(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'enable': {
      onEnable(row);
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
          return await getActivityList({
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
  } as VxeTableGridOptions<ActivityInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="活动管理">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['content:activity:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['活动']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
