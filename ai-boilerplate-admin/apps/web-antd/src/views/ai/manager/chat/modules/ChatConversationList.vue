<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type * as AiChatConversationApi from '#/api/v1/ai-chat-conversation';
import type { GetSysAdminSelectorItem } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';

import { message } from 'ant-design-vue';

import {
  TABLE_ACTION_ICON,
  TableAction,
  useVbenVxeGrid,
} from '#/adapter/vxe-table';
import {
  deleteAiChatConversation,
  getAiChatConversationList,
} from '#/api/v1/ai-chat-conversation';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import {
  useGridColumnsConversation,
  useGridFormSchemaConversation,
} from '../data';

const userList = ref<GetSysAdminSelectorItem[]>([]); // 用户列表

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 删除 */
async function handleDelete(row: AiChatConversationApi.AiChatConversationInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.id]),
    key: 'action_key_msg',
  });
  try {
    if (!row.id) {
      return;
    }
    await deleteAiChatConversation({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.id]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchemaConversation(),
  },
  gridOptions: {
    columns: useGridColumnsConversation(),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getAiChatConversationList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as AiChatConversationApi.GetAiChatConversationListParams,
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
  } as VxeTableGridOptions<AiChatConversationApi.AiChatConversationInfo>,
  separator: false,
});

onMounted(async () => {
  // 获得用户列表
  const res = await getSysAdminSelector({});
  userList.value = res.list || [];
});
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="对话列表">
      <template #toolbar-tools>
        <TableAction :actions="[]" />
      </template>
      <template #adminId="{ row }">
        <span>
          {{ userList.find((item) => item.id === row.adminId)?.nickname }}
        </span>
      </template>
      <template #actions="{ row }">
        <TableAction
          :actions="[
            {
              label: $t('common.delete'),
              type: 'link',
              danger: true,
              icon: TABLE_ACTION_ICON.DELETE,
              auth: ['ai:chat-conversation:delete'],
              popConfirm: {
                title: $t('ui.actionMessage.deleteConfirm', [row.id]),
                confirm: handleDelete.bind(null, row),
              },
            },
          ]"
        />
      </template>
    </Grid>
  </Page>
</template>
