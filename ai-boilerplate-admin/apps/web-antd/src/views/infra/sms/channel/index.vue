<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SmsChannelInfo } from '#/api/v1/sms-channel';

import { Page, useVbenModal } from '@vben/common-ui';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteSmsChannel, getSmsChannelList } from '#/api/v1/sms-channel';
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

/** 创建短信渠道 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑短信渠道 */
function onEdit(row: SmsChannelInfo) {
  formModalApi.setData(row).open();
}

/** 删除短信渠道 */
async function onDelete(row: SmsChannelInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSmsChannel({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.channelName]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SmsChannelInfo>) {
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
          return await getSmsChannelList({
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
  } as VxeTableGridOptions<SmsChannelInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="短信渠道列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:sms-channel:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['短信渠道']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
