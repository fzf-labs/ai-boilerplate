<script lang="ts" setup>
import { ref, watch } from 'vue';

import { InboxOutlined } from '@ant-design/icons-vue';
import {
  Button,
  Form,
  Input,
  message,
  Modal,
  Progress,
  Select,
  Upload,
} from 'ant-design-vue';

import { request } from '#/api/request';

type MaterialTypeValue = 'image' | 'video' | 'voice';

interface UploadData {
  appId?: string;
  type?: MaterialTypeValue;
}

const props = defineProps<{
  data?: UploadData;
  open: boolean;
}>();

const emits = defineEmits<{
  close: [];
  success: [];
}>();

// Material type constants
const MaterialType = {
  IMAGE: 'image',
  VOICE: 'voice',
  VIDEO: 'video',
} as const;

const MaterialTypeLabels: Record<string, string> = {
  [MaterialType.IMAGE]: '图片',
  [MaterialType.VOICE]: '语音',
  [MaterialType.VIDEO]: '视频',
};

// Upload material function
async function uploadMaterial(formData: FormData) {
  return request('/admin/v1/wx_gzh_material/upload', {
    method: 'POST',
    data: formData,
  });
}

interface UploadFile {
  file: File;
  name: string;
  title?: string; // 视频标题（必填）
  introduction?: string; // 视频简介（必填）
  progress: number;
  status: 'error' | 'success' | 'uploading' | 'waiting';
  error?: string;
}

const loading = ref(false);
const uploadFiles = ref<UploadFile[]>([]);
const currentType = ref<MaterialTypeValue>(MaterialType.IMAGE);

// 文件类型限制
const fileTypeMap = {
  [MaterialType.IMAGE]: {
    accept: 'image/*',
    maxSize: 10 * 1024 * 1024, // 10MB
    extensions: ['jpg', 'jpeg', 'png', 'gif', 'bmp'],
  },
  [MaterialType.VOICE]: {
    accept: 'audio/*',
    maxSize: 60 * 1024 * 1024, // 60MB
    extensions: ['mp3', 'wma', 'wav', 'amr'],
  },
  [MaterialType.VIDEO]: {
    accept: 'video/*',
    maxSize: 10 * 1024 * 1024, // 10MB
    extensions: ['mp4'],
  },
};

// 素材类型选项
const materialTypeOptions = [
  { label: MaterialTypeLabels[MaterialType.IMAGE], value: MaterialType.IMAGE },
  { label: MaterialTypeLabels[MaterialType.VOICE], value: MaterialType.VOICE },
  { label: MaterialTypeLabels[MaterialType.VIDEO], value: MaterialType.VIDEO },
];

// 验证文件
const validateFile = (file: File, type: MaterialTypeValue): null | string => {
  const config = fileTypeMap[type];

  // 检查文件大小
  if (file.size > config.maxSize) {
    const maxSizeMB = (config.maxSize / (1024 * 1024)).toFixed(1);
    return `文件大小不能超过 ${maxSizeMB}MB`;
  }

  // 检查文件扩展名
  const extension = file.name.split('.').pop()?.toLowerCase();
  if (!extension || !config.extensions.includes(extension)) {
    return `只支持 ${config.extensions.join(', ')} 格式的文件`;
  }

  return null;
};

// 处理文件选择
const handleFileSelect = (file: File) => {
  // 验证文件
  const error = validateFile(file, currentType.value);
  if (error) {
    message.error(error);
    return false;
  }

  // 添加到上传列表
  const uploadFile: UploadFile = {
    file,
    name: file.name,
    progress: 0,
    status: 'waiting',
  };

  uploadFiles.value.push(uploadFile);
  return false; // 阻止自动上传
};

// 删除文件
const handleRemoveFile = (index: number) => {
  uploadFiles.value.splice(index, 1);
};

// 开始上传
const handleUpload = async () => {
  if (uploadFiles.value.length === 0) {
    message.warning('请选择要上传的文件');
    return;
  }

  if (!props.data?.appId) {
    message.error('缺少账号信息');
    return;
  }

  // 验证视频类型的必填字段
  if (currentType.value === MaterialType.VIDEO) {
    const invalidFiles = uploadFiles.value.filter(
      (file) => !file.title || !file.introduction,
    );
    if (invalidFiles.length > 0) {
      message.error('请填写所有视频的标题和简介');
      return;
    }
  }

  loading.value = true;

  try {
    for (let i = 0; i < uploadFiles.value.length; i++) {
      const uploadFile = uploadFiles.value[i];
      if (!uploadFile) continue;

      uploadFile.status = 'uploading';

      try {
        const formData = new FormData();
        formData.append('file', uploadFile.file);
        formData.append('appId', props.data.appId);
        formData.append('type', currentType.value.toString());

        // 根据素材类型使用不同的字段名
        if (currentType.value === MaterialType.VIDEO) {
          formData.append('title', uploadFile.title || '');
          formData.append('introduction', uploadFile.introduction || '');
        }
        // 模拟进度更新
        const progressInterval = setInterval(() => {
          if (uploadFile && uploadFile.progress < 90) {
            uploadFile.progress += 10;
          }
        }, 100);

        await uploadMaterial(formData);

        clearInterval(progressInterval);
        if (uploadFile) {
          uploadFile.progress = 100;
          uploadFile.status = 'success';
        }
      } catch (error) {
        if (uploadFile) {
          uploadFile.status = 'error';
          uploadFile.error =
            error instanceof Error ? error.message : '上传失败';
        }
        console.error('上传失败:', error);
      }
    }

    const successCount = uploadFiles.value.filter(
      (f) => f.status === 'success',
    ).length;
    const errorCount = uploadFiles.value.filter(
      (f) => f.status === 'error',
    ).length;

    if (successCount > 0) {
      message.success(`成功上传 ${successCount} 个文件`);
      emits('success');
    }

    if (errorCount > 0) {
      message.error(`${errorCount} 个文件上传失败`);
    }

    if (successCount === uploadFiles.value.length) {
      handleClose();
    }
  } finally {
    loading.value = false;
  }
};

// 关闭弹窗
const handleClose = () => {
  uploadFiles.value = [];
  emits('close');
};

// 监听数据变化
watch(
  () => props.data,
  (newData) => {
    if (newData && newData.type) {
      currentType.value = newData.type;
    }
  },
  { immediate: true },
);

// 格式化文件大小
const formatFileSize = (size: number) => {
  if (size < 1024) return `${size} B`;
  if (size < 1024 * 1024) return `${(size / 1024).toFixed(1)} KB`;
  return `${(size / (1024 * 1024)).toFixed(1)} MB`;
};
</script>

<template>
  <Modal
    :open="open"
    title="上传素材"
    width="800px"
    :confirm-loading="loading"
    @ok="handleUpload"
    @cancel="handleClose"
  >
    <div class="upload-container">
      <!-- 素材类型选择 -->
      <Form layout="vertical">
        <Form.Item label="素材类型">
          <Select
            v-model:value="currentType"
            :options="materialTypeOptions"
            style="width: 200px"
          />
        </Form.Item>
      </Form>

      <!-- 文件上传区域 -->
      <div class="upload-area">
        <Upload.Dragger
          :accept="fileTypeMap[currentType].accept"
          :before-upload="handleFileSelect"
          :show-upload-list="false"
          multiple
        >
          <p class="ant-upload-drag-icon">
            <InboxOutlined />
          </p>
          <p class="ant-upload-text">点击或拖拽文件到此区域上传</p>
          <p class="ant-upload-hint">
            支持 {{ fileTypeMap[currentType].extensions.join(', ') }} 格式，
            单个文件不超过
            {{
              (fileTypeMap[currentType].maxSize / (1024 * 1024)).toFixed(1)
            }}MB
          </p>
        </Upload.Dragger>
      </div>

      <!-- 文件列表 -->
      <div v-if="uploadFiles.length > 0" class="file-list">
        <div class="file-list-header">
          <span>待上传文件 ({{ uploadFiles.length }})</span>
        </div>
        <div class="file-list-body">
          <div
            v-for="(uploadFile, index) in uploadFiles"
            :key="index"
            class="file-item"
          >
            <div class="file-info">
              <div class="file-name">
                <span class="file-name-text" :title="uploadFile.name">
                  {{ uploadFile.name }}
                </span>
              </div>
              <div class="file-meta">
                <span>{{ formatFileSize(uploadFile.file.size) }}</span>
                <span>{{ uploadFile.file.type }}</span>
              </div>
              <!-- 视频类型的必填字段 -->
              <div
                v-if="currentType === MaterialType.VIDEO"
                class="video-fields"
              >
                <div class="video-title">
                  <Input
                    v-model:value="uploadFile.title"
                    placeholder="视频标题（必填）"
                    :status="!uploadFile.title ? 'error' : ''"
                  />
                </div>
                <div class="video-introduction">
                  <Input
                    v-model:value="uploadFile.introduction"
                    placeholder="视频简介（必填）"
                    :status="!uploadFile.introduction ? 'error' : ''"
                  />
                </div>
              </div>
            </div>

            <div class="file-status">
              <div v-if="uploadFile.status === 'uploading'" class="uploading">
                <Progress :percent="uploadFile.progress" size="small" />
              </div>
              <div v-else-if="uploadFile.status === 'success'" class="success">
                ✅ 上传成功
              </div>
              <div v-else-if="uploadFile.status === 'error'" class="error">
                ❌ {{ uploadFile.error }}
              </div>
            </div>

            <div class="file-actions">
              <Button size="small" danger @click="handleRemoveFile(index)">
                删除
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Modal>
</template>

<style scoped lang="scss">
.upload-container {
  .upload-area {
    margin-bottom: 24px;
  }

  .file-list {
    .file-list-header {
      padding: 8px 0;
      margin-bottom: 16px;
      font-weight: 500;
      border-bottom: 1px solid #f0f0f0;
    }

    .file-list-body {
      .file-item {
        display: flex;
        align-items: center;
        padding: 12px;
        margin-bottom: 8px;
        border: 1px solid #d9d9d9;
        border-radius: 6px;

        .file-info {
          flex: 1;
          margin-right: 16px;

          .file-name {
            margin-bottom: 8px;

            .file-name-text {
              display: block;
              padding: 4px 8px;
              overflow: hidden;
              text-overflow: ellipsis;
              font-size: 14px;
              color: #333;
              word-break: break-all;
              white-space: nowrap;
              background-color: #f5f5f5;
              border: 1px solid #d9d9d9;
              border-radius: 4px;
            }
          }

          .file-meta {
            margin-bottom: 8px;
            font-size: 12px;
            color: #999;

            span {
              margin-right: 16px;
            }
          }

          .video-fields {
            .video-title,
            .video-introduction {
              margin-bottom: 8px;
            }
          }
        }

        .file-status {
          width: 120px;
          margin-right: 16px;

          .success {
            font-size: 12px;
            color: #52c41a;
          }

          .error {
            font-size: 12px;
            color: #ff4d4f;
          }
        }

        .file-actions {
          width: 60px;
        }
      }
    }
  }
}
</style>
