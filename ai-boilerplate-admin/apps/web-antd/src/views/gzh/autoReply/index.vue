<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { WxGzhAutoReplyInfo } from '#/api/v1/wx-gzh-auto-reply';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteWxGzhAutoReply,
  getWxGzhAutoReplyList,
} from '#/api/v1/wx-gzh-auto-reply';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DetailModalComponent, DetailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建自动回复 */
async function handleCreate() {
  try {
    // 获取当前搜索表单的 appId
    const formValues = await gridApi.formApi.getValues();
    const currentAppId = formValues?.appId;

    formModalApi.setData({ appId: currentAppId }).open();
  } catch (error) {
    console.error('Failed to open create form:', error);
    message.error('打开创建表单失败');
  }
}

/** 查看自动回复 */
function handleView(row: WxGzhAutoReplyInfo) {
  DetailModalApi.setData(row).open();
}

/** 编辑自动回复 */
function handleEdit(row: WxGzhAutoReplyInfo) {
  formModalApi.setData(row).open();
}

/** 删除自动回复 */
async function handleDelete(row: WxGzhAutoReplyInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', ['自动回复']),
    key: 'action_key_msg',
  });
  try {
    await deleteWxGzhAutoReply({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', ['自动回复']),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<WxGzhAutoReplyInfo>) {
  switch (code) {
    case 'delete': {
      handleDelete(row);
      break;
    }
    case 'edit': {
      handleEdit(row);
      break;
    }
    case 'view': {
      handleView(row);
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
          return await getWxGzhAutoReplyList({
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
      refresh: true,
      search: true,
    },
  } as VxeTableGridOptions<WxGzhAutoReplyInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <DetailModalComponent />
    <Grid table-title="自动回复管理">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="handleCreate"
          v-access:code="['gzh:auto-reply:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['自动回复']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
