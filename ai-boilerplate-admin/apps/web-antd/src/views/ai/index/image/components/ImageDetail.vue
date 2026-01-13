<script setup lang="ts">
import type { PropType } from 'vue';

import { computed } from 'vue';

import { formatDate } from '@vben/utils';

import { Image } from 'ant-design-vue';

import {
  AiPlatformEnum,
  Dall3StyleList,
  type ImageRecordView,
  normalizeImageRecord,
  StableDiffusionClipGuidancePresets,
  StableDiffusionSamplers,
  StableDiffusionStylePresets,
} from './typing';

// 图片详细信息
const props = defineProps({
  detail: {
    type: Object as PropType<ImageRecordView | null>,
    default: null,
  },
});
const detail = computed(() =>
  props.detail ? normalizeImageRecord(props.detail) : null,
);
</script>

<template>
  <div class="mb-5 w-full overflow-hidden break-words">
    <div class="mt-2 text-gray-600">
      <Image class="rounded-lg" :src="detail?.picUrl || detail?.picURL" />
    </div>
  </div>

  <!-- 时间 -->
  <div class="mb-5 w-full overflow-hidden break-words">
    <div class="text-lg font-bold">时间</div>
    <div class="mt-2 text-gray-600">
      <div>
        提交时间：{{ formatDate(detail?.createTime, 'yyyy-MM-dd HH:mm:ss') }}
      </div>
      <div>
        生成时间：{{ formatDate(detail?.finishTime, 'yyyy-MM-dd HH:mm:ss') }}
      </div>
    </div>
  </div>

  <!-- 模型 -->
  <div class="mb-5 w-full overflow-hidden break-words">
    <div class="text-lg font-bold">模型</div>
    <div class="mt-2 text-gray-600">
      {{ detail?.model }}({{ detail?.height }}x{{ detail?.width }})
    </div>
  </div>

  <!-- 提示词 -->
  <div class="mb-5 w-full overflow-hidden break-words">
    <div class="text-lg font-bold">提示词</div>
    <div class="mt-2 text-gray-600">
      {{ detail?.prompt }}
    </div>
  </div>

  <!-- 图片地址 -->
  <div class="mb-5 w-full overflow-hidden break-words">
    <div class="text-lg font-bold">图片地址</div>
    <div class="mt-2 text-gray-600">
      {{ detail?.picUrl || detail?.picURL }}
    </div>
  </div>

  <!-- StableDiffusion 专属 -->
  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.sampler
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">采样方法</div>
    <div class="mt-2 text-gray-600">
      {{
        StableDiffusionSamplers.find(
          (item) => item.key === detail?.options?.sampler,
        )?.name
      }}
    </div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.clipGuidancePreset
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">CLIP</div>
    <div class="mt-2 text-gray-600">
      {{
        StableDiffusionClipGuidancePresets.find(
          (item) => item.key === detail?.options?.clipGuidancePreset,
        )?.name
      }}
    </div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.stylePreset
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">风格</div>
    <div class="mt-2 text-gray-600">
      {{
        StableDiffusionStylePresets.find(
          (item) => item.key === detail?.options?.stylePreset,
        )?.name
      }}
    </div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.steps
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">迭代步数</div>
    <div class="mt-2 text-gray-600">{{ detail?.options?.steps }}</div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.scale
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">引导系数</div>
    <div class="mt-2 text-gray-600">{{ detail?.options?.scale }}</div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.STABLE_DIFFUSION &&
      detail?.options?.seed
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">随机因子</div>
    <div class="mt-2 text-gray-600">{{ detail?.options?.seed }}</div>
  </div>

  <!-- Dall3 专属 -->
  <div
    v-if="detail?.platform === AiPlatformEnum.OPENAI && detail?.options?.style"
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">风格选择</div>
    <div class="mt-2 text-gray-600">
      {{
        Dall3StyleList.find((item) => item.key === detail?.options?.style)?.name
      }}
    </div>
  </div>

  <!-- Midjourney 专属 -->
  <div
    v-if="
      detail?.platform === AiPlatformEnum.MIDJOURNEY && detail?.options?.version
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">模型版本</div>
    <div class="mt-2 text-gray-600">{{ detail?.options?.version }}</div>
  </div>

  <div
    v-if="
      detail?.platform === AiPlatformEnum.MIDJOURNEY &&
      detail?.options?.referImageUrl
    "
    class="mb-5 w-full overflow-hidden break-words"
  >
    <div class="text-lg font-bold">参考图</div>
    <div class="mt-2 text-gray-600">
      <Image :src="detail?.options?.referImageUrl" />
    </div>
  </div>
</template>
