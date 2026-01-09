<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { MailAccountInfo } from '#/api/v1/mail-account';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteMailAccount, getMailAccountList } from '#/api/v1/mail-account';
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

/** 创建邮箱账号 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑邮箱账号 */
function onEdit(row: MailAccountInfo) {
  formModalApi.setData(row).open();
}

/** 删除邮箱账号 */
async function onDelete(row: MailAccountInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.mail]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMailAccount({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.mail]),
      key: 'action_process_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<MailAccountInfo>) {
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
          return await getMailAccountList({
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
  } as VxeTableGridOptions<MailAccountInfo>,
});
</script>
<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="邮箱账号列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:mail-account:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['邮箱账号']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
