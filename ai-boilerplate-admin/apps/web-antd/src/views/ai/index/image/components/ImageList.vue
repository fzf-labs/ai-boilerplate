<script setup lang="ts">
import type { ImageMidjourneyButton, ImageRecordView } from './typing';

import { onMounted, onUnmounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';

import { confirm, useVbenDrawer } from '@vben/common-ui';
import { downloadFileFromImageUrl } from '@vben/utils';

import { useDebounceFn } from '@vueuse/core';
import { Button, Card, message, Pagination } from 'ant-design-vue';

import {
  deleteAiIndexImageRecord,
  getAiIndexImageRecordList,
} from '#/api/v1/ai-index-image';

import ImageCard from './ImageCard.vue';
import ImageDetail from './ImageDetail.vue';
import { AiImageStatusEnum, normalizeImageRecord } from './typing';

// 暴露组件方法

const emits = defineEmits(['onRegeneration']);
const router = useRouter(); // 路由
const [Drawer, drawerApi] = useVbenDrawer({
  title: '图片详情',
  footer: false,
});
// 图片分页相关的参数
const queryParams = reactive({
  page: 1,
  pageSize: 10,
});
const pageTotal = ref<number>(0); // page size
const imageList = ref<ImageRecordView[]>([]); // image 列表
const imageListRef = ref<any>(); // ref
// 图片轮询相关的参数（正在生成中的）
const inProgressImageMap = ref<{}>({}); // 监听的 image 映射，一般是生成中（需要轮询），key 为 image 编号，value 为 image
const inProgressTimer = ref<any>(); // 生成中的 image 定时器，轮询生成进展
const showImageDetail = ref<ImageRecordView | null>(null); // 图片详情的图片信息

/** 处理查看绘图作品 */
function handleViewPublic() {
  router.push({
    name: 'AiImageSquare',
  });
}

/** 查看图片的详情  */
async function handleDetailOpen() {
  drawerApi.open();
}
/** 获得 image 图片列表 */
async function getImageList(isSilent = false) {
  const loading = isSilent
    ? null
    : message.loading({
        content: `加载中...`,
      });
  try {
    // 1. 加载图片列表

    const { list, total } = await getAiIndexImageRecordList({
      params: {
        page: queryParams.page,
        pageSize: queryParams.pageSize,
      },
    });
    imageList.value = (list || []).map((item) => normalizeImageRecord(item));
    pageTotal.value = total || 0;

    // 2. 计算需要轮询的图片
    const newWatImages: any = {};
    imageList.value.forEach((item: any) => {
      if (item.status === AiImageStatusEnum.IN_PROGRESS) {
        newWatImages[item.id] = item;
      }
    });
    inProgressImageMap.value = newWatImages;
  } finally {
    // 关闭正在“加载中”的 Loading
    if (loading) {
      loading();
    }
  }
}
const debounceGetImageList = useDebounceFn(getImageList, 80);
const handlePaginationChange = () => {
  debounceGetImageList();
};
/** 轮询生成中的 image 列表 */
async function refreshWatchImages() {
  if (Object.keys(inProgressImageMap.value).length === 0) {
    return;
  }
  await getImageList(true);
}

/** 图片的点击事件 */
async function handleImageButtonClick(
  type: string,
  imageDetail: ImageRecordView,
) {
  // 详情
  if (type === 'more') {
    showImageDetail.value = imageDetail;
    await handleDetailOpen();
    return;
  }
  // 删除
  if (type === 'delete') {
    await confirm(`是否删除照片?`);
    if (imageDetail.id) {
      await deleteAiIndexImageRecord({
        body: {
          id: imageDetail.id,
        },
      });
    }
    await getImageList();
    message.success('删除成功!');
    return;
  }
  // 下载
  if (type === 'download') {
    const source = imageDetail.picUrl || imageDetail.picURL;
    if (!source) {
      message.error('图片地址为空');
      return;
    }
    await downloadFileFromImageUrl({
      fileName: imageDetail.model || imageDetail.modelId || 'image',
      source,
    });
    return;
  }
  // 重新生成
  if (type === 'regeneration') {
    await emits('onRegeneration', imageDetail);
  }
}

/** 处理 Midjourney 按钮点击事件  */
async function handleImageMidjourneyButtonClick(
  button: ImageMidjourneyButton,
  imageDetail: ImageRecordView,
) {
  if (!button || !imageDetail?.id) {
    return;
  }
  message.info('当前接口暂不支持该操作');
}

defineExpose({ getImageList }); /** 组件挂在的时候 */
onMounted(async () => {
  // 获取 image 列表
  await getImageList();
  // 自动刷新 image 列表
  inProgressTimer.value = setInterval(async () => {
    await refreshWatchImages();
  }, 1000 * 3);
});

/** 组件取消挂在的时候 */
onUnmounted(async () => {
  if (inProgressTimer.value) {
    clearInterval(inProgressTimer.value);
  }
});
</script>
<template>
  <Drawer class="w-2/5">
    <ImageDetail :detail="showImageDetail" />
  </Drawer>
  <Card
    class="flex h-full w-full flex-col"
    :body-style="{
      margin: 0,
      padding: 0,
      height: '100%',
      position: 'relative',
      display: 'flex',
      flexDirection: 'column',
    }"
  >
    <template #title>
      绘画任务
      <Button @click="handleViewPublic">绘画作品</Button>
    </template>

    <div
      class="flex flex-1 flex-wrap content-start overflow-y-auto p-5 pb-28 pt-5"
      ref="imageListRef"
    >
      <ImageCard
        v-for="image in imageList"
        :key="image.id"
        :detail="image"
        @on-btn-click="handleImageButtonClick"
        @on-mj-btn-click="handleImageMidjourneyButtonClick"
        class="mb-5 mr-5"
      />
    </div>

    <div
      class="bg-card sticky bottom-0 z-50 flex h-16 items-center justify-center shadow-sm"
    >
      <Pagination
        :total="pageTotal"
        :show-total="(total) => `共 ${total} 条`"
        show-quick-jumper
        show-size-changer
        v-model:current="queryParams.page"
        v-model:page-size="queryParams.pageSize"
        @change="handlePaginationChange"
        @show-size-change="handlePaginationChange"
      />
    </div>
  </Card>
</template>
