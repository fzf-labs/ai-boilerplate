<!-- dall3 -->
<script setup lang="ts">
import type { ProviderModelOption } from '../../../utils';
import type { ImageRecordView } from '../typing';

import { ref, watch } from 'vue';

import { confirm } from '@vben/common-ui';
import { useUserStore } from '@vben/stores';

import {
  Button,
  InputNumber,
  message,
  Select,
  Space,
  Textarea,
} from 'ant-design-vue';

import { createAiIndexImageRecord } from '#/api/v1/ai-index-image';

import {
  AiImageStatusEnum,
  AiPlatformEnum,
  ImageHotWords,
  OtherPlatformEnum,
} from '../typing';

// 消息弹窗

// 接收父组件传入的模型列表
const props = defineProps({
  models: {
    type: Array<ProviderModelOption>,
    default: () => [] as ProviderModelOption[],
  },
});
const emits = defineEmits(['onDrawStart', 'onDrawComplete']);

// 定义属性
const drawIn = ref<boolean>(false); // 生成中
const selectHotWord = ref<string>(''); // 选中的热词
// 表单
const prompt = ref<string>(''); // 提示词
const width = ref<number>(512); // 图片宽度
const height = ref<number>(512); // 图片高度
const otherPlatform = ref<string>(AiPlatformEnum.TONG_YI); // 平台
const platformModels = ref<ProviderModelOption[]>([]); // 模型列表
const modelId = ref<string>(); // 选中的模型
const userStore = useUserStore();

/** 选择热词 */
async function handleHotWordClick(hotWord: string) {
  // 情况一：取消选中
  if (selectHotWord.value === hotWord) {
    selectHotWord.value = '';
    return;
  }

  // 情况二：选中
  selectHotWord.value = hotWord; // 选中
  prompt.value = hotWord; // 替换提示词
}

/** 图片生成 */
async function handleGenerateImage() {
  // 二次确认
  await confirm(`确认生成内容?`);
  try {
    // 加载中
    drawIn.value = true;
    // 回调
    emits('onDrawStart', otherPlatform.value);
    // 发送请求
    const selectedModel = platformModels.value.find(
      (item) => item.modelId === modelId.value,
    );
    if (!selectedModel) {
      message.error('请选择模型');
      return;
    }
    await createAiIndexImageRecord({
      body: {
        adminId: userStore.userInfo?.userId || '',
        prompt: prompt.value,
        platform: otherPlatform.value,
        modelId: selectedModel.modelId || '',
        model: selectedModel.modelName || selectedModel.modelId || '',
        width: width.value,
        height: height.value,
        status: AiImageStatusEnum.IN_PROGRESS,
        publicStatus: false,
        options: JSON.stringify({}),
      },
    });
  } finally {
    // 回调
    emits('onDrawComplete', otherPlatform.value);
    // 加载结束
    drawIn.value = false;
  }
}

/** 填充值 */
async function settingValues(detail: ImageRecordView) {
  prompt.value = detail.prompt ?? '';
  selectHotWord.value = ImageHotWords.includes(prompt.value)
    ? prompt.value
    : '';
  width.value = detail.width ?? 512;
  height.value = detail.height ?? 512;
  otherPlatform.value = detail.platform ?? AiPlatformEnum.TONG_YI;
  handlerPlatformChange(otherPlatform.value);
}

/** 平台切换 */
function handlerPlatformChange(platform: any) {
  // 根据选择的平台筛选模型
  platformModels.value = props.models.filter(
    (item: ProviderModelOption) => item.platformName === platform,
  );
  modelId.value =
    platformModels.value.length > 0 && platformModels.value[0]
      ? platformModels.value[0].modelId
      : undefined;
  // 切换平台，默认选择一个模型
}

/** 监听 models 变化 */
watch(
  () => props.models,
  () => {
    handlerPlatformChange(otherPlatform.value);
  },
  { immediate: true, deep: true },
);
/** 暴露组件方法 */
defineExpose({ settingValues });
</script>
<template>
  <div class="prompt">
    <b>画面描述</b>
    <Textarea
      v-model:value="prompt"
      :maxlength="1024"
      :rows="5"
      class="mt-4 w-full"
      placeholder="请输入画面描述"
      show-count
    />
  </div>

  <div class="mt-8 flex flex-col">
    <div>
      <b>随机热词</b>
    </div>
    <Space wrap class="mt-4 flex flex-wrap justify-start">
      <Button
        shape="round"
        class="m-0"
        :type="selectHotWord === hotWord ? 'primary' : 'default'"
        v-for="hotWord in ImageHotWords"
        :key="hotWord"
        @click="handleHotWordClick(hotWord)"
      >
        {{ hotWord }}
      </Button>
    </Space>
  </div>

  <div class="mt-8">
    <div>
      <b>平台</b>
    </div>
    <Space wrap class="mt-4 w-full">
      <Select
        v-model:value="otherPlatform"
        placeholder="请选择平台"
        size="large"
        class="!w-80"
        @change="handlerPlatformChange"
      >
        <Select.Option
          v-for="item in OtherPlatformEnum"
          :key="item.key"
          :value="item.key"
        >
          {{ item.name }}
        </Select.Option>
      </Select>
    </Space>
  </div>

  <div class="mt-8">
    <div>
      <b>模型</b>
    </div>
    <Space wrap class="mt-4 w-full">
      <Select
        v-model:value="modelId"
        placeholder="请选择模型"
        size="large"
        class="!w-80"
      >
        <Select.Option
          v-for="item in platformModels"
          :key="item.id || item.modelId"
          :value="item.modelId"
        >
          {{ item.modelName || item.modelId }}
        </Select.Option>
      </Select>
    </Space>
  </div>

  <div class="mt-8">
    <div>
      <b>图片尺寸</b>
    </div>
    <Space wrap class="mt-4 flex flex-wrap gap-x-5">
      <InputNumber
        v-model:value="width"
        class="w-40"
        placeholder="图片宽度"
        addon-before="宽"
        addon-after="px"
      />
      <InputNumber
        v-model:value="height"
        class="w-40"
        placeholder="图片高度"
        addon-before="高"
        addon-after="px"
      />
    </Space>
  </div>

  <div class="mt-12 flex justify-center">
    <Button
      type="primary"
      size="large"
      shape="round"
      :loading="drawIn"
      :disabled="prompt.length === 0"
      @click="handleGenerateImage"
    >
      {{ drawIn ? '生成中' : '生成内容' }}
    </Button>
  </div>
</template>
