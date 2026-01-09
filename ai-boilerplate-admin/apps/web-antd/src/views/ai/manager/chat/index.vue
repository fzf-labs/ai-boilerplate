<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { AiChatMessageApi } from '#/api/ai/manager/chatMessage';
import type { AiChatConversationApi } from '#/api/v1/ai-chat-conversation';
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';

import { message, Modal } from 'ant-design-vue';

import {
  TABLE_ACTION_ICON,
  TableAction,
  useVbenVxeGrid,
} from '#/adapter/vxe-table';
import {
  deleteAiChatMessage,
  getAiChatMessageList,
} from '#/api/ai/manager/chatMessage';
import {
  deleteAiChatConversation,
  getAiChatConversationList,
} from '#/api/v1/ai-chat-conversation';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import {
  useGridColumnsConversation,
  useGridColumnsMessage,
  useGridFormSchemaMessage,
} from './data';

const userList = ref<SysAdminInfo[]>([]); // 用户列表
const messageModalVisible = ref(false); // 消息弹窗显示状态
const selectedConversation =
  ref<AiChatConversationApi.AiChatConversationInfo | null>(null); // 当前选中的对话

/** 刷新对话列表 */
function onRefreshConversation() {
  conversationGridApi.query();
}

/** 刷新消息列表 */
function onRefreshMessage() {
  messageGridApi.query();
}

/** 删除对话 */
async function handleDeleteConversation(
  row: AiChatConversationApi.AiChatConversationInfo,
) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.id]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiChatConversation(row.id);
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.id]),
      key: 'action_key_msg',
    });
    onRefreshConversation();
  } finally {
    hideLoading();
  }
}

/** 删除消息 */
async function handleDeleteMessage(row: AiChatMessageApi.AiChatMessageInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.id]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiChatMessage(row.id);
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.id]),
      key: 'action_key_msg',
    });
    onRefreshMessage();
  } finally {
    hideLoading();
  }
}

/** 查看对话消息 */
function handleViewMessages(row: AiChatConversationApi.AiChatConversationInfo) {
  selectedConversation.value = row;
  messageModalVisible.value = true;
  // 延迟加载消息列表，确保弹窗已渲染
  setTimeout(() => {
    messageGridApi.query();
  }, 100);
}

/** 关闭消息弹窗 */
function handleCloseModal() {
  messageModalVisible.value = false;
  selectedConversation.value = null;
}

// 对话列表
const [ConversationGrid, conversationGridApi] = useVbenVxeGrid({
  formOptions: {
    schema: [],
  },
  gridOptions: {
    columns: useGridColumnsConversation(),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getAiChatConversationList({
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
  } as VxeTableGridOptions<AiChatConversationApi.AiChatConversationInfo>,
  separator: false,
});

// 消息列表（用于弹窗）
const [MessageGrid, messageGridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchemaMessage(),
  },
  gridOptions: {
    columns: useGridColumnsMessage(),
    height: 600,
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          if (!selectedConversation.value) {
            return { total: 0, list: [] };
          }
          return await getAiChatMessageList({
            page: page.currentPage,
            pageSize: page.pageSize,
            conversationId: selectedConversation.value.id,
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
    <ConversationGrid table-title="对话列表">
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
              label: '消息详情',
              type: 'link',
              icon: TABLE_ACTION_ICON.VIEW,
              onClick: handleViewMessages.bind(null, row),
            },
            {
              label: $t('common.delete'),
              type: 'link',
              danger: true,
              icon: TABLE_ACTION_ICON.DELETE,
              auth: ['ai:chat-conversation:delete'],
              popConfirm: {
                title: $t('ui.actionMessage.deleteConfirm', [row.id]),
                confirm: handleDeleteConversation.bind(null, row),
              },
            },
          ]"
        />
      </template>
    </ConversationGrid>

    <!-- 消息详情弹窗 -->
    <Modal
      v-model:open="messageModalVisible"
      :title="`消息详情 - ${selectedConversation?.title || ''}`"
      width="90%"
      :footer="null"
      :destroy-on-close="true"
      @cancel="handleCloseModal"
    >
      <div class="message-modal-content">
        <MessageGrid>
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
                  auth: ['ai:chat-message:delete'],
                  popConfirm: {
                    title: $t('ui.actionMessage.deleteConfirm', [row.id]),
                    confirm: handleDeleteMessage.bind(null, row),
                  },
                },
              ]"
            />
          </template>
        </MessageGrid>
      </div>
    </Modal>
  </Page>
</template>

<style scoped>
.message-modal-content {
  padding: 12px 0;
}
</style>
