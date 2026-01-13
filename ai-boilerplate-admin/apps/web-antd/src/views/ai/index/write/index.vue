<script lang="ts" setup>
import type { CreateAiWriteRecordReq } from '#/api/v1/ai-write-record';

import { nextTick, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { useUserStore } from '@vben/stores';

import { message } from 'ant-design-vue';

import {
  createAiWriteRecord,
  getAiWriteRecordInfo,
} from '#/api/v1/ai-write-record';

import Left from './components/Left.vue';
import Right from './components/Right.vue';
import { WriteExample } from './components/typing';

const writeResult = ref(''); // 写作结果
const isWriting = ref(false); // 是否正在写作中
const userStore = useUserStore();

/** 停止 stream 生成 */
function stopStream() {
  isWriting.value = false;
}

/** 执行写作 */
const rightRef = ref<InstanceType<typeof Right>>();

type WriteSubmitPayload = Partial<CreateAiWriteRecordReq> & {
  model?: string;
  platform?: string;
};

async function submit(data: WriteSubmitPayload) {
  try {
    writeResult.value = '';
    isWriting.value = true;
    if (!data.modelId || !data.model) {
      message.error('请选择模型');
      return;
    }
    const payload: CreateAiWriteRecordReq = {
      adminId: userStore.userInfo?.userId || '',
      platform: data.platform || '',
      modelId: data.modelId,
      model: data.model,
      prompt: data.prompt || '',
      generatedContent: data.generatedContent,
      originalContent: data.originalContent,
      length: data.length,
      format: data.format,
      tone: data.tone,
      language: data.language,
      type: data.type,
    };
    const res = await createAiWriteRecord({
      body: payload,
    });
    if (res?.id) {
      const info = await getAiWriteRecordInfo({
        params: { id: res.id },
      });
      if (info?.info?.generatedContent) {
        writeResult.value = info.info.generatedContent;
        await nextTick();
        rightRef.value?.scrollToBottom();
      } else if (info?.info?.errorMessage) {
        message.error(info.info.errorMessage);
      } else {
        message.success('已提交写作任务');
      }
    }
  } catch (error: any) {
    message.error(error?.message || '写作失败');
  } finally {
    isWriting.value = false;
  }
}

/** 点击示例触发 */
function handleExampleClick(type: keyof typeof WriteExample) {
  writeResult.value = WriteExample[type].data;
}

/** 点击重置的时候清空写作的结果*/
function reset() {
  writeResult.value = '';
}
</script>

<template>
  <Page auto-content-height>
    <div class="absolute bottom-0 left-0 right-0 top-0 m-4 flex">
      <Left
        :is-writing="isWriting"
        class="mr-4 h-full rounded-lg"
        @submit="submit"
        @reset="reset"
        @example="handleExampleClick"
      />
      <Right
        :is-writing="isWriting"
        @stop-stream="stopStream"
        ref="rightRef"
        class="flex-grow"
        v-model:content="writeResult"
      />
    </div>
  </Page>
</template>
