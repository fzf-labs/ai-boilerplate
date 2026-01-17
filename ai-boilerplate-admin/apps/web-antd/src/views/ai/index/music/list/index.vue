<script setup lang="ts">
import type { Recordable } from '@vben/types';

import type { ProviderModelOption } from '../../utils';

import type { AiIndexAudioRecordInfo } from '#/api/v1/ai-index-audio';

import { onMounted, provide, ref } from 'vue';

import { useUserStore } from '@vben/stores';

import { Col, Empty, message, Row, TabPane, Tabs } from 'ant-design-vue';

import {
  createAiIndexAudioRecord,
  getAiIndexAudioRecordList,
} from '#/api/v1/ai-index-audio';

import { fetchProviderModels } from '../../utils';
import audioBar from './audioBar/index.vue';
import songCard from './songCard/index.vue';
import songInfo from './songInfo/index.vue';

defineOptions({ name: 'AiMusicListIndex' });

const currentType = ref('mine');
// loading 状态
const loading = ref(false);
// 当前音乐
const currentSong = ref({});
const userStore = useUserStore();
const modelOptions = ref<ProviderModelOption[]>([]);

type MusicRecordView = AiIndexAudioRecordInfo & {
  audioUrl?: string;
  date?: string;
  desc?: string;
  imageUrl?: string;
};

const mySongList = ref<MusicRecordView[]>([]);
const squareSongList = ref<MusicRecordView[]>([]);

function normalizeMusicRecord(record: AiIndexAudioRecordInfo): MusicRecordView {
  return {
    ...record,
    imageUrl: record.imageURL,
    audioUrl: record.audioURL,
    desc: record.description,
    date: record.createdAt,
  };
}

async function loadMusicList() {
  const { list } = await getAiIndexAudioRecordList({
    params: { page: 1, pageSize: 50 },
  });
  const normalized = (list || []).map((item) => normalizeMusicRecord(item));
  mySongList.value = normalized.filter((item) => !item.publicStatus);
  squareSongList.value = normalized.filter((item) => item.publicStatus);
  if (!currentSong.value?.id && normalized[0]) {
    currentSong.value = normalized[0];
  }
}

/*
 *@Description: 调接口生成音乐列表
 *@MethodAuthor: xiaohong
 *@Date: 2024-06-27 17:06:44
 */
async function generateMusic(formData: Recordable<any>) {
  loading.value = true;
  try {
    if (!formData) {
      message.warning('请填写生成信息');
      return;
    }
    const selectedModel = modelOptions.value[0];
    if (!selectedModel) {
      message.warning('未配置音频模型');
      return;
    }
    const isLyric = formData.generateMode === 'lyric' || !!formData.lyric;
    const desc = formData.desc || formData.description || '';
    const title = formData.name || (desc ? desc.slice(0, 12) : '') || '未命名';
    const promptParts = [
      isLyric ? formData.style : desc,
      formData.pure ? '纯音乐' : '',
      formData.version ? `版本:${formData.version}` : '',
    ].filter(Boolean);
    await createAiIndexAudioRecord({
      body: {
        tenantId: userStore.userInfo?.tenantId || '',
        adminId: userStore.userInfo?.userId || '',
        title,
        description: desc || formData.style,
        lyric: formData.lyric,
        prompt: promptParts.join(' '),
        tags: formData.style,
        generateMode: isLyric ? 2 : 1,
        publicStatus: false,
        status: 10,
        platform: selectedModel.platformName || '',
        modelId: selectedModel.modelId || '',
        model: selectedModel.modelName || selectedModel.modelId || '',
      },
    });
    await loadMusicList();
    message.success('已提交音乐任务');
  } catch (error: any) {
    message.error(error?.message || '音乐生成失败');
  } finally {
    loading.value = false;
  }
}

/*
 *@Description: 设置当前播放的音乐
 *@MethodAuthor: xiaohong
 *@Date: 2024-07-19 11:22:33
 */
function setCurrentSong(music: Recordable<any>) {
  currentSong.value = music;
}

defineExpose({
  generateMusic,
});

provide('currentSong', currentSong);

onMounted(async () => {
  try {
    loading.value = true;
    modelOptions.value = await fetchProviderModels('audio');
    await loadMusicList();
  } catch (error: any) {
    message.error(error?.message || '音乐列表加载失败');
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <div class="flex flex-col">
    <div class="flex flex-auto overflow-hidden">
      <Tabs
        v-model:active-key="currentType"
        class="flex-auto px-5"
        tab-position="bottom"
      >
        <!-- 我的创作 -->
        <TabPane key="mine" tab="我的创作" v-loading="loading">
          <Row v-if="mySongList.length > 0" :gutter="12">
            <Col v-for="song in mySongList" :key="song.id" :span="24">
              <songCard :song-info="song" @play="setCurrentSong(song)" />
            </Col>
          </Row>
          <Empty v-else description="暂无音乐" />
        </TabPane>

        <!-- 试听广场 -->
        <TabPane key="square" tab="试听广场" v-loading="loading">
          <Row v-if="squareSongList.length > 0" :gutter="12">
            <Col v-for="song in squareSongList" :key="song.id" :span="24">
              <songCard :song-info="song" @play="setCurrentSong(song)" />
            </Col>
          </Row>
          <Empty v-else description="暂无音乐" />
        </TabPane>
      </Tabs>
      <!-- songInfo -->
      <songInfo class="flex-none" />
    </div>
    <audioBar class="flex-none" />
  </div>
</template>
<style lang="scss" scoped>
:deep(.ant-tabs) {
  .ant-tabs__content {
    padding: 0 7px;
    overflow: auto;
  }
}
</style>
