<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { ConfigDatumInfo } from '#/api/v1/config-data';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteConfigDatum, getConfigDatumList } from '#/api/v1/config-data';
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

/** 创建参数 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑参数 */
function onEdit(row: ConfigDatumInfo) {
  formModalApi.setData(row).open();
}

/** 删除参数 */
async function onDelete(row: ConfigDatumInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteConfigDatum({ body: { id: row.id } });
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
function onActionClick({ code, row }: OnActionClickParams<ConfigDatumInfo>) {
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
          return await getConfigDatumList({
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
  } as VxeTableGridOptions<ConfigDatumInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="参数列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:config:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['参数']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
