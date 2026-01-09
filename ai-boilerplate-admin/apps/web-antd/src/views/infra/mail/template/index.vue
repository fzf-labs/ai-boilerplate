<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { MailAccountInfo } from '#/api/v1/mail-account';
import type { MailTemplateInfo } from '#/api/v1/mail-template';

import { onMounted, ref } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getMailAccountSelector } from '#/api/v1/mail-account';
import {
  deleteMailTemplate,
  getMailTemplateList,
} from '#/api/v1/mail-template';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';
import SendForm from './modules/send-form.vue';

const accountList = ref<MailAccountInfo[]>([]);

/** 获取邮箱账号 */
const getAccountMail = (accountId: string) => {
  return accountList.value.find((account) => account.id === accountId)?.mail;
};

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

/** 创建邮件模板 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑邮件模板 */
function onEdit(row: MailTemplateInfo) {
  formModalApi.setData(row).open();
}

/** 发送测试邮件 */
function onSend(row: MailTemplateInfo) {
  sendModalApi.setData(row).open();
}

/** 删除邮件模板 */
async function onDelete(row: MailTemplateInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMailTemplate({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<MailTemplateInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'send': {
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
    columns: useGridColumns(onActionClick, getAccountMail),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getMailTemplateList({
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
  } as VxeTableGridOptions<MailTemplateInfo>,
});

/** 初始化 */
onMounted(async () => {
  const res = await getMailAccountSelector({});
  accountList.value = res.list || [];
});
</script>
<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <SendModal />
    <Grid table-title="邮件模板列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['infra:mail-template:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['邮件模板']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
