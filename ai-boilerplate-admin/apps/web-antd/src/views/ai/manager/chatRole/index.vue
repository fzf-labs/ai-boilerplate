<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { AiChatRoleApi } from '#/api/ai/manager/chatRole';

import { Page, useVbenModal } from '@vben/common-ui';

import { Image, message } from 'ant-design-vue';

import {
  TABLE_ACTION_ICON,
  TableAction,
  useVbenVxeGrid,
} from '#/adapter/vxe-table';
import { deleteAiChatRole, getAiChatRoleList } from '#/api/ai/manager/chatRole';
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

/** 创建 */
function handleCreate() {
  formModalApi.setData(null).open();
}

/** 编辑 */
function handleEdit(row: AiChatRoleApi.AiChatRoleInfo) {
  formModalApi.setData(row).open();
}

/** 删除 */
async function handleDelete(row: AiChatRoleApi.AiChatRoleInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiChatRole({ id: row.id });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getAiChatRoleList({
            page: page.currentPage,
            pageSize: page.pageSize,
            ...formValues,
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
  } as VxeTableGridOptions<AiChatRoleApi.AiChatRoleInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="AI聊天角色列表">
      <template #toolbar-tools>
        <TableAction
          :actions="[
            {
              label: $t('ui.actionTitle.create', ['AI聊天角色']),
              type: 'primary',
              icon: TABLE_ACTION_ICON.ADD,
              auth: ['ai:chat-role:create'],
              onClick: handleCreate,
            },
          ]"
        />
      </template>
      <template #avatar="{ row }">
        <Image :src="row.avatar" class="h-8 w-8" />
      </template>
      <template #actions="{ row }">
        <TableAction
          :actions="[
            {
              label: $t('common.edit'),
              type: 'link',
              icon: TABLE_ACTION_ICON.EDIT,
              auth: ['ai:chat-role:update'],
              onClick: handleEdit.bind(null, row),
            },
            {
              label: $t('common.delete'),
              type: 'link',
              danger: true,
              icon: TABLE_ACTION_ICON.DELETE,
              auth: ['ai:chat-role:delete'],
              popConfirm: {
                title: $t('ui.actionMessage.deleteConfirm', [row.name]),
                confirm: handleDelete.bind(null, row),
              },
            },
          ]"
        />
      </template>
    </Grid>
  </Page>
</template>
