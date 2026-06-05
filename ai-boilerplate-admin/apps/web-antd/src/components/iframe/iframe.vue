<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';

interface IFrameProps {
  /** iframe 的源地址 */
  src: string;
}

const props = defineProps<IFrameProps>();

const loading = ref(true);
const height = ref('');
const frameRef = ref<HTMLIFrameElement | null>(null);

function init() {
  height.value = `${document.documentElement.clientHeight - 94.5}px`;
}

watch(
  () => props.src,
  () => {
    loading.value = true;
  },
  { immediate: true },
);

onMounted(() => {
  init();
});
// 路由级内链由 layouts 的 IFrameView 处理，这里保留给直接 URL 嵌入使用。
</script>

<template>
  <div v-loading="loading" :style="`height:${height}`">
    <iframe
      ref="frameRef"
      :src="props.src"
      style="width: 100%; height: 100%"
      frameborder="no"
      scrolling="auto"
      @load="loading = false"
    ></iframe>
  </div>
</template>
