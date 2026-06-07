<script lang="ts" setup>
import type { WxGzhMaterialInfo } from '#/api/v1/wx-gzh-material';

import { computed } from 'vue';

import { CopyOutlined } from '@ant-design/icons-vue';
import { Button, Descriptions, Modal, Space, Tag } from 'ant-design-vue';

import {
  getMaterialTypeColor,
  getMaterialTypeLabel,
  MaterialType,
  normalizeMaterialTags,
} from '../helpers'

const props = defineProps<{
  data?: WxGzhMaterialInfo;
  open: boolean;
}>();

const emits = defineEmits<{
  close: [];
}>();

// 计算属性
const materialTypeLabel = computed(() => {
  return getMaterialTypeLabel(props.data?.type);
});

const isImage = computed(() => {
  return props.data?.type === MaterialType.IMAGE;
});

const isVoice = computed(() => {
  return props.data?.type === MaterialType.VOICE;
});

const isVideo = computed(() => {
  return props.data?.type === MaterialType.VIDEO;
});

const materialTags = computed(() => {
  return normalizeMaterialTags(props.data?.tags);
});

// 格式化时间
const formatTime = (time: Date | string | undefined) => {
  if (!time) return '-';
  return new Date(time).toLocaleString();
};

// 复制链接
const handleCopyUrl = async () => {
  if (!props.data?.URL) return;

  try {
    await navigator.clipboard.writeText(props.data.URL);
    // 这里可以添加成功提示
  } catch (error) {
    console.error('复制失败:', error);
  }
};

// 关闭弹窗
const handleClose = () => {
  emits('close');
};
</script>

<template>
  <Modal
    :open="open"
    :title="`素材预览 - ${materialTypeLabel}`"
    width="800px"
    :footer="null"
    @cancel="handleClose"
  >
    <div v-if="data" class="preview-container">
      <!-- 预览区域 -->
      <div class="preview-area">
        <!-- 图片预览 -->
        <div v-if="isImage" class="image-preview">
          <img :src="data.URL" :alt="data.name" referrerPolicy="no-referrer" />
        </div>

        <!-- 音频预览 -->
        <div v-else-if="isVoice" class="voice-preview">
          <div class="voice-icon">🎵</div>
          <div class="voice-name">{{ data.name }}</div>
          <audio controls :src="data.URL" class="audio-player">
            您的浏览器不支持音频播放
          </audio>
        </div>

        <!-- 视频预览 -->
        <div v-else-if="isVideo" class="video-preview">
          <video
            v-if="data.URL"
            controls
            :poster="data.coverURL || ''"
            :src="data.URL"
            class="video-player"
          />
          <div v-else-if="data.coverURL">
            <img
              :src="data.coverURL"
              :alt="data.name"
              referrerPolicy="no-referrer"
            />
          </div>
          <div v-else>🎬</div>
        </div>
      </div>

      <!-- 素材信息 -->
      <div class="material-info">
        <Descriptions title="素材信息" :column="2" bordered>
          <Descriptions.Item label="素材名称">
            {{ data.name }}
          </Descriptions.Item>
          <Descriptions.Item label="素材类型">
            <Tag
              :color="getMaterialTypeColor(data.type)"
            >
              {{ materialTypeLabel }}
            </Tag>
          </Descriptions.Item>
          <Descriptions.Item label="标签" :span="2">
            <Space wrap>
              <Tag v-if="materialTags.length === 0" color="default">
                -
              </Tag>
              <Tag v-for="tag in materialTags" :key="tag" color="processing">
                {{ tag }}
              </Tag>
            </Space>
          </Descriptions.Item>
          <Descriptions.Item label="Media ID">
            <code>{{ data.mediaId || '-' }}</code>
          </Descriptions.Item>
          <Descriptions.Item label="上传时间">
            {{ formatTime(data.updateTime) }}
          </Descriptions.Item>
          <Descriptions.Item label="访问链接" :span="2">
            <div class="url-container" v-if="data.URL">
              <code class="url-text">{{ data.URL }}</code>
              <Space>
                <Button size="small" @click="handleCopyUrl">
                  <CopyOutlined />
                  复制
                </Button>
              </Space>
            </div>
          </Descriptions.Item>
        </Descriptions>
      </div>
    </div>
  </Modal>
</template>

<style scoped lang="scss">
.preview-container {
  .preview-area {
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 200px;
    padding: 24px;
    margin-bottom: 24px;
    text-align: center;
    background-color: #fafafa;
    border-radius: 8px;

    .image-preview {
      max-width: 100%;
      max-height: 400px;

      img {
        max-width: 100%;
        max-height: 400px;
        object-fit: contain;
        border-radius: 4px;
        box-shadow: 0 2px 8px rgb(0 0 0 / 10%);
      }
    }

    .voice-preview {
      text-align: center;

      .voice-icon {
        margin-bottom: 16px;
        font-size: 64px;
      }

      .voice-name {
        margin-bottom: 16px;
        font-size: 16px;
        font-weight: 500;
        color: #333;
      }

      .audio-player {
        width: 100%;
        max-width: 400px;
      }
    }

    .video-preview {
      .video-player {
        max-width: 100%;
        max-height: 400px;
        border-radius: 4px;
        box-shadow: 0 2px 8px rgb(0 0 0 / 10%);
      }
    }
  }

  .material-info {
    .url-container {
      display: flex;
      gap: 8px;
      align-items: center;

      .url-text {
        flex: 1;
        padding: 4px 8px;
        font-size: 12px;
        word-break: break-all;
        background-color: #f5f5f5;
        border-radius: 4px;
      }
    }
  }
}
</style>
