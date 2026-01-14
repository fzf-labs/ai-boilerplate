<script lang="ts" setup>
import type * as AiIndexChatApi from '#/api/v1/ai-index-chat';

import { computed, nextTick, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

import { alert, confirm, Page, useVbenModal } from '@vben/common-ui';
import { IconifyIcon } from '@vben/icons';

import { Button, Layout, message, Switch } from 'ant-design-vue';

import {
  getAiIndexChatConversationItem,
  getAiIndexChatMessageList,
} from '#/api/v1/ai-index-chat';

import ConversationList from './components/conversation/ConversationList.vue';
import KnowledgeSettingForm from './components/conversation/KnowledgeSettingForm.vue';
import McpSettingForm from './components/conversation/McpSettingForm.vue';
import ModelSettingForm from './components/conversation/ModelSettingForm.vue';
import PromptSettingForm from './components/conversation/PromptSettingForm.vue';
import MessageList from './components/message/MessageList.vue';
import MessageListEmpty from './components/message/MessageListEmpty.vue';
import MessageLoading from './components/message/MessageLoading.vue';
import MessageNewConversation from './components/message/MessageNewConversation.vue';
/** AI 聊天对话 列表 */
defineOptions({ name: 'AiChat' });

const route = useRoute(); // 路由

// 四个独立的设置弹窗
const [PromptModal, promptModalApi] = useVbenModal({
  connectedComponent: PromptSettingForm,
  destroyOnClose: true,
});

const [ModelModal, modelModalApi] = useVbenModal({
  connectedComponent: ModelSettingForm,
  destroyOnClose: true,
});

const [KnowledgeModal, knowledgeModalApi] = useVbenModal({
  connectedComponent: KnowledgeSettingForm,
  destroyOnClose: true,
});

const [McpModal, mcpModalApi] = useVbenModal({
  connectedComponent: McpSettingForm,
  destroyOnClose: true,
});
// 聊天对话
const conversationListRef = ref();
const activeConversationId = ref<null | string>(null); // 选中的对话编号
const activeConversation =
  ref<AiIndexChatApi.AiIndexChatConversationItem | null>(null); // 选中的 Conversation
const conversationInProgress = ref(false); // 对话是否正在进行中。目前只有【发送】消息时，会更新为 true，避免切换对话、删除对话等操作

// 消息列表
const messageRef = ref();
const activeMessageList = ref<AiIndexChatApi.AiIndexChatMessageItem[]>([]); // 选中对话的消息列表
const activeMessageListLoading = ref<boolean>(false); // activeMessageList 是否正在加载中
const activeMessageListLoadingTimer = ref<any>(); // activeMessageListLoading Timer 定时器。如果加载速度很快，就不进入加载中
// 发送消息输入框
const isComposing = ref(false); // 判断用户是否在输入
const inputTimeout = ref<any>(); // 处理输入中回车的定时器
const prompt = ref<string>(); // prompt
const enableContext = ref<boolean>(true); // 是否开启上下文

// =========== 【聊天对话】相关 ===========

/** 获取对话信息 */
async function getConversation(id: null | string) {
  if (!id) {
    return;
  }
  const res = await getAiIndexChatConversationItem({
    params: { id },
  });
  if (!res?.info) {
    return;
  }
  activeConversation.value = res.info;
  activeConversationId.value = res.info.id;
}

/**
 * 点击某个对话
 *
 * @param conversation 选中的对话
 * @return 是否切换成功
 */
async function handleConversationClick(
  conversation: AiIndexChatApi.AiIndexChatConversationItem,
) {
  if (!conversation.id) {
    return false;
  }
  // 对话进行中，不允许切换
  if (conversationInProgress.value) {
    alert('对话中，不允许切换!');
    return false;
  }

  // 更新选中的对话 id
  activeConversationId.value = conversation.id;
  activeConversation.value = conversation;
  // 刷新 message 列表
  await getMessageList();
  // 滚动底部
  scrollToBottom(true);
  // 清空输入框
  prompt.value = '';
  return true;
}

/** 删除某个对话*/
async function handlerConversationDelete(
  delConversation: AiIndexChatApi.AiIndexChatConversationItem,
) {
  // 删除的对话如果是当前选中的，那么就重置
  if (activeConversationId.value === delConversation.id) {
    await handleConversationClear();
  }
}

/** 清空选中的对话 */
async function handleConversationClear() {
  // 对话进行中，不允许切换
  if (conversationInProgress.value) {
    alert('对话中，不允许切换!');
    return false;
  }
  activeConversationId.value = null;
  activeConversation.value = null;
  activeMessageList.value = [];
}

// 打开提示词设置
async function openPromptSetting() {
  promptModalApi.setData({ id: activeConversationId.value }).open();
}

// 打开模型设置
async function openModelSetting() {
  modelModalApi.setData({ id: activeConversationId.value }).open();
}

// 打开知识库设置
async function openKnowledgeSetting() {
  knowledgeModalApi.setData({ id: activeConversationId.value }).open();
}

// 打开 MCP 设置
async function openMcpSetting() {
  mcpModalApi.setData({ id: activeConversationId.value }).open();
}

// 处理设置更新成功
async function handleConversationUpdateSuccess() {
  // 对话更新成功，刷新最新信息
  await getConversation(activeConversationId.value);
}

/** 处理聊天对话的创建成功 */
async function handleConversationCreate() {
  // 创建对话
  await conversationListRef.value.createConversation();
}
/** 处理聊天对话的创建成功 */
async function handleConversationCreateSuccess() {
  // 创建新的对话，清空输入框
  prompt.value = '';
}

// =========== 【消息列表】相关 ===========

/** 获取消息 message 列表 */
async function getMessageList() {
  try {
    if (activeConversationId.value === null) {
      return;
    }
    // Timer 定时器，如果加载速度很快，就不进入加载中
    activeMessageListLoadingTimer.value = setTimeout(() => {
      activeMessageListLoading.value = true;
    }, 60);

    // 获取消息列表
    const res = await getAiIndexChatMessageList({
      params: {
        page: 1,
        pageSize: 200,
        conversationId: activeConversationId.value,
      } as any,
    });
    activeMessageList.value = res.list || [];

    // 滚动到最下面
    await nextTick();
    await scrollToBottom();
  } finally {
    // time 定时器，如果加载速度很快，就不进入加载中
    if (activeMessageListLoadingTimer.value) {
      clearTimeout(activeMessageListLoadingTimer.value);
    }
    // 加载结束
    activeMessageListLoading.value = false;
  }
}

/**
 * 消息列表
 *
 * 和 {@link #getMessageList()} 的差异是，把 systemMessage 考虑进去
 */
const messageList = computed(() => {
  if (activeMessageList.value.length > 0) {
    return activeMessageList.value;
  }
  // 没有消息时，如果有 systemMessage 则展示它
  if (activeConversation.value?.promptSetting?.prompt) {
    return [
      {
        id: 'system',
        type: 'system',
        content: activeConversation.value.promptSetting.prompt,
        createdAt: activeConversation.value.createdAt,
      },
    ];
  }
  return [];
});

/** 处理删除 message 消息 */
function handleMessageDelete(message?: AiIndexChatApi.AiIndexChatMessageItem) {
  if (conversationInProgress.value) {
    alert('回答中，不能删除!');
    return;
  }
  if (message?.id) {
    activeMessageList.value = activeMessageList.value.filter(
      (item) => item.id !== message.id,
    );
    return;
  }
  getMessageList();
}

/** 处理 message 清空 */
async function handlerMessageClear() {
  if (!activeConversationId.value) {
    return;
  }
  try {
    // 确认提示
    await confirm('确认清空对话消息？');
    // 刷新 message 列表
    activeMessageList.value = [];
  } catch {}
}

/** 回到 message 列表的顶部 */
function handleGoTopMessage() {
  messageRef.value.handlerGoTop();
}

// =========== 【发送消息】相关 ===========
/** 处理来自 keydown 的发送消息 */
async function handleSendByKeydown(event: any) {
  // 判断用户是否在输入
  if (isComposing.value) {
    return;
  }
  // 进行中不允许发送
  if (conversationInProgress.value) {
    return;
  }
  const content = prompt.value?.trim() as string;
  if (event.key === 'Enter') {
    if (event.shiftKey) {
      // 插入换行
      prompt.value += '\r\n';
      event.preventDefault(); // 防止默认的换行行为
    } else {
      // 发送消息
      await doSendMessage(content);
      event.preventDefault(); // 防止默认的提交行为
    }
  }
}

/** 处理来自【发送】按钮的发送消息 */
function handleSendByButton() {
  doSendMessage(prompt.value?.trim() as string);
}

/** 处理 prompt 输入变化 */
function handlePromptInput(event: any) {
  // 非输入法 输入设置为 true
  if (!isComposing.value) {
    // 回车 event data 是 null
    if (event.data === null || event.data === 'null') {
      return;
    }
    isComposing.value = true;
  }
  // 清理定时器
  if (inputTimeout.value) {
    clearTimeout(inputTimeout.value);
  }
  // 重置定时器
  inputTimeout.value = setTimeout(() => {
    isComposing.value = false;
  }, 400);
}

function onCompositionstart() {
  isComposing.value = true;
}

function onCompositionend() {
  // console.log('输入结束...')
  setTimeout(() => {
    isComposing.value = false;
  }, 200);
}

/** 真正执行【发送】消息操作 */
async function doSendMessage(content: string) {
  // 校验
  if (content.length === 0) {
    message.error('发送失败，原因：内容为空！');
    return;
  }
  if (activeConversationId.value === null) {
    message.error('还没创建对话，不能发送!');
    return;
  }
  // 清空输入框
  prompt.value = '';
  message.warning('当前接口暂不支持发送消息');
}

/** 停止 stream 流式调用 */
async function stopStream() {
  // 设置为 false
  conversationInProgress.value = false;
}

/** 编辑 message：设置为 prompt，可以再次编辑 */
function handleMessageEdit(message: AiIndexChatApi.AiIndexChatMessageItem) {
  prompt.value = message.content;
}

/** 刷新 message：基于指定消息，再次发起对话 */
function handleMessageRefresh(message: AiIndexChatApi.AiIndexChatMessageItem) {
  doSendMessage(message.content);
}

/** 滚动到 message 底部 */
async function scrollToBottom(isIgnore?: boolean) {
  await nextTick();
  if (messageRef.value) {
    messageRef.value.scrollToBottom(isIgnore);
  }
}

/** 初始化 */
onMounted(async () => {
  // 如果有 conversationId 参数，则默认选中
  if (route.query.conversationId) {
    const id = String(route.query.conversationId);
    activeConversationId.value = id;
    await getConversation(id);
  }

  // 获取列表数据
  activeMessageListLoading.value = true;
  await getMessageList();
});
</script>

<template>
  <Page auto-content-height>
    <Layout class="absolute left-0 top-0 m-4 h-full w-full flex-1">
      <!-- 左侧：对话列表 -->
      <ConversationList
        class="!bg-card"
        :active-id="activeConversationId as any"
        ref="conversationListRef"
        @on-conversation-create="handleConversationCreateSuccess"
        @on-conversation-click="handleConversationClick"
        @on-conversation-clear="handleConversationClear"
        @on-conversation-delete="handlerConversationDelete"
      />

      <!-- 右侧：详情部分 -->
      <Layout class="bg-card mx-4">
        <Layout.Header
          class="!bg-card border-border flex items-center justify-between border-b"
        >
          <div class="text-lg font-bold">
            {{ activeConversation?.title ? activeConversation?.title : '对话' }}
            <span v-if="activeMessageList.length > 0">
              ({{ activeMessageList.length }})
            </span>
          </div>

          <div class="flex justify-end gap-2" v-if="activeConversation">
            <!-- 提示词设置 -->
            <Button
              type="primary"
              ghost
              class="px-3"
              size="small"
              @click="openPromptSetting"
            >
              <IconifyIcon icon="lucide:message-square" class="mr-1 size-4" />
              <span>提示词</span>
            </Button>
            <!-- 模型设置 -->
            <Button
              type="primary"
              ghost
              class="px-3"
              size="small"
              @click="openModelSetting"
            >
              <IconifyIcon icon="lucide:brain" class="mr-1 size-4" />
              <span>模型</span>
            </Button>
            <!-- 知识库设置 -->
            <Button
              type="primary"
              ghost
              class="px-3"
              size="small"
              @click="openKnowledgeSetting"
            >
              <IconifyIcon icon="lucide:database" class="mr-1 size-4" />
              <span>知识库</span>
            </Button>
            <!-- MCP 设置 -->
            <Button
              type="primary"
              ghost
              class="px-3"
              size="small"
              @click="openMcpSetting"
            >
              <IconifyIcon icon="lucide:plug" class="mr-1 size-4" />
              <span>MCP</span>
            </Button>
            <!-- 清空消息 -->
            <Button size="small" class="px-2" @click="handlerMessageClear">
              <IconifyIcon icon="lucide:trash-2" color="#787878" />
            </Button>
            <!-- 导出 -->
            <Button size="small" class="px-2">
              <IconifyIcon icon="lucide:download" color="#787878" />
            </Button>
            <!-- 回到顶部 -->
            <Button size="small" class="px-2" @click="handleGoTopMessage">
              <IconifyIcon icon="lucide:arrow-up" color="#787878" />
            </Button>
          </div>
        </Layout.Header>

        <Layout.Content class="relative m-0 h-full w-full p-0">
          <div class="absolute inset-0 m-0 overflow-y-hidden p-0">
            <MessageLoading v-if="activeMessageListLoading" />
            <MessageNewConversation
              v-if="!activeConversation"
              @on-new-conversation="handleConversationCreate"
            />
            <MessageListEmpty
              v-if="
                !activeMessageListLoading &&
                messageList.length === 0 &&
                activeConversation
              "
              @on-prompt="doSendMessage"
            />
            <MessageList
              v-if="!activeMessageListLoading && messageList.length > 0"
              ref="messageRef"
              :conversation="activeConversation as any"
              :list="messageList as any"
              @on-delete-success="handleMessageDelete"
              @on-edit="handleMessageEdit"
              @on-refresh="handleMessageRefresh"
            />
          </div>
        </Layout.Content>

        <Layout.Footer class="!bg-card m-0 flex flex-col p-0">
          <form
            class="border-border my-5 mb-5 mt-2 flex flex-col rounded-xl border px-2 py-2.5"
          >
            <textarea
              class="box-border h-24 resize-none overflow-auto border-none px-0 py-1 focus:outline-none"
              v-model="prompt"
              @keydown="handleSendByKeydown"
              @input="handlePromptInput"
              @compositionstart="onCompositionstart"
              @compositionend="onCompositionend"
              placeholder="问我任何问题...（Shift+Enter 换行，按下 Enter 发送）"
            ></textarea>
            <div class="flex justify-between pb-0 pt-1">
              <div class="flex items-center">
                <Switch v-model:checked="enableContext" />
                <span class="ml-1 text-sm text-gray-400">上下文</span>
              </div>
              <Button
                type="primary"
                @click="handleSendByButton"
                :loading="conversationInProgress"
                v-if="conversationInProgress === false"
              >
                <IconifyIcon
                  :icon="
                    conversationInProgress
                      ? 'lucide:loader'
                      : 'lucide:send-horizontal'
                  "
                />
                {{ conversationInProgress ? '进行中' : '发送' }}
              </Button>
              <Button
                type="primary"
                danger
                @click="stopStream()"
                v-if="conversationInProgress === true"
              >
                <IconifyIcon icon="lucide:circle-stop" />
                停止
              </Button>
            </div>
          </form>
        </Layout.Footer>
      </Layout>
    </Layout>
    <!-- 四个独立的设置弹窗 -->
    <PromptModal @success="handleConversationUpdateSuccess" />
    <ModelModal @success="handleConversationUpdateSuccess" />
    <KnowledgeModal @success="handleConversationUpdateSuccess" />
    <McpModal @success="handleConversationUpdateSuccess" />
  </Page>
</template>
