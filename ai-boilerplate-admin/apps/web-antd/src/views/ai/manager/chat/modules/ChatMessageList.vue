<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type * as AiChatMessageApi from '#/api/v1/ai-chat-message';
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';

import { TableAction, useVbenVxeGrid } from '#/adapter/vxe-table';
import { getAiChatMessageList } from '#/api/v1/ai-chat-message';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { useGridColumnsMessage, useGridFormSchemaMessage } from '../data';

const userList = ref<SysAdminInfo[]>([]); // 用户列表

const [Grid] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchemaMessage(),
  },
  gridOptions: {
    columns: useGridColumnsMessage(),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          if (!formValues.conversationId) {
            return { total: 0, list: [] };
          }
          return await getAiChatMessageList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as AiChatMessageApi.GetAiChatMessageListParams,
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
  } as VxeTableGridOptions<AiChatMessageApi.AiChatMessageInfo>,
  separator: false,
});

onMounted(async () => {
  // 获得用户列表
  const res = await getSysAdminSelector();
  userList.value = res.list;
});
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="消息列表">
      <template #toolbar-tools>
        <TableAction :actions="[]" />
      </template>
      <template #adminId="{ row }">
        <span>{{
          userList.find((item) => item.id === row.adminId)?.nickname
        }}</span>
      </template>
    </Grid>
  </Page>
</template>
