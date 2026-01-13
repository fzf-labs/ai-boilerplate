<script setup lang="ts">
import type { Recordable } from '@vben/types';
import type { AiMusicApi } from '#/api/ai/music';

import { onMounted, provide, ref } from 'vue';

import { Col, Empty, Row, TabPane, Tabs, message } from 'ant-design-vue';

import audioBar from './audioBar/index.vue';
import songCard from './songCard/index.vue';
import songInfo from './songInfo/index.vue';
import { createMusicRecord, getMusicPageMy } from '#/api/ai/music';

defineOptions({ name: 'AiMusicListIndex' });

const currentType = ref('mine');
// loading 状态
const loading = ref(false);
// 当前音乐
const currentSong = ref({});

const mySongList = ref<AiMusicApi.Music[]>([]);
const squareSongList = ref<AiMusicApi.Music[]>([]);

async function loadMusicList() {
  const { list } = await getMusicPageMy({ page: 1, pageSize: 50 });
  mySongList.value = list.filter((item) => !item.publicStatus);
  squareSongList.value = list.filter((item) => item.publicStatus);
  if (!currentSong.value?.id && list[0]) {
    currentSong.value = list[0];
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
    const isLyric = formData.generateMode === 'lyric' || !!formData.lyric;
    const desc = formData.desc || formData.description || '';
    const title =
      formData.name || (desc ? desc.slice(0, 12) : '') || '未命名';
    const promptParts = [
      isLyric ? formData.style : desc,
      formData.pure ? '纯音乐' : '',
      formData.version ? `版本:${formData.version}` : '',
    ].filter(Boolean);
    await createMusicRecord({
      title,
      description: desc || formData.style,
      lyric: formData.lyric,
      prompt: promptParts.join(' '),
      tags: formData.style,
      generateMode: isLyric ? 2 : 1,
      publicStatus: false,
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
