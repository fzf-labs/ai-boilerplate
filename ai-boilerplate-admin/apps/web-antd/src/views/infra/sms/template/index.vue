<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SmsTemplateInfo } from '#/api/v1/sms-template';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteSmsTemplate, getSmsTemplateList } from '#/api/v1/sms-template';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';
import SendForm from './modules/send-form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [SendModal, sendModalApi] = useVbenModal({
  connectedComponent: SendForm,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建短信模板 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑短信模板 */
function onEdit(row: SmsTemplateInfo) {
  formModalApi.setData(row).open();
}

/** 发送测试短信 */
function onSend(row: SmsTemplateInfo) {
  sendModalApi.setData(row).open();
}

/** 删除短信模板 */
async function onDelete(row: SmsTemplateInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSmsTemplate({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.templateName]),
      key: 'action_process_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SmsTemplateInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'sms-send': {
      onSend(row);
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
          return await getSmsTemplateList({
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
  } as VxeTableGridOptions<SmsTemplateInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <SendModal />
    <Grid table-title="短信模板列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:sms-template:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['短信模板']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
