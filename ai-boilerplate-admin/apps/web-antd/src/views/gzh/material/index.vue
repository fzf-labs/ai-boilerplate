<script lang="ts" setup>
import type {
  GetWxGzhMaterialStatsReply,
  WxGzhMaterialInfo,
} from '#/api/v1/wx-gzh-material';

import { computed, onMounted, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';

import {
  CloudUploadOutlined,
  DeleteOutlined,
  EyeOutlined,
  ReloadOutlined,
  SearchOutlined,
  SyncOutlined,
} from '@ant-design/icons-vue';
import {
  Button,
  Card,
  Col,
  Empty,
  Flex,
  Input,
  message,
  Pagination,
  Row,
  Space,
  Spin,
  Tabs,
  Tooltip,
} from 'ant-design-vue';

import {
  deleteWxGzhMaterial,
  getWxGzhMaterialList,
  getWxGzhMaterialStats,
  syncWxGzhMaterial,
} from '#/api/v1/wx-gzh-material';
import MaterialPreview from '#/views/gzh/material/components/material-preview.vue';
import MaterialUpload from '#/views/gzh/material/components/material-upload.vue';

import AccountSelect from '../components/account-select/index.vue';

// Material type constants
const MaterialType = {
  IMAGE: 'image',
  VOICE: 'voice',
  VIDEO: 'video',
} as const;

type MaterialTypeValue = (typeof MaterialType)[keyof typeof MaterialType];

const MaterialTypeLabels: Record<string, string> = {
  [MaterialType.IMAGE]: '图片',
  [MaterialType.VOICE]: '语音',
  [MaterialType.VIDEO]: '视频',
};

const appId = ref<string | undefined>(undefined);
const loading = ref(false);
const syncLoading = ref(false);
const currentType = ref<MaterialTypeValue>(MaterialType.IMAGE);
const searchKeyword = ref('');
const selectedMaterials = ref<string[]>([]);

// 分页参数
const pagination = ref({
  current: 1,
  pageSize: 20,
  total: 0,
});

// 素材列表
const materialList = ref<WxGzhMaterialInfo[]>([]);
const materialStats = ref<GetWxGzhMaterialStatsReply>();

// 模态框状态
const previewVisible = ref(false);
const uploadVisible = ref(false);
const currentMaterial = ref<WxGzhMaterialInfo>();

// 计算属性
const hasSelectedMaterials = computed(() => selectedMaterials.value.length > 0);

// 素材类型选项
const materialTypeOptions = [
  { label: MaterialTypeLabels[MaterialType.IMAGE], value: MaterialType.IMAGE },
  { label: MaterialTypeLabels[MaterialType.VOICE], value: MaterialType.VOICE },
  { label: MaterialTypeLabels[MaterialType.VIDEO], value: MaterialType.VIDEO },
];

const handleAccountChange = (value: any) => {
  appId.value = value;
  selectedMaterials.value = [];
};

// 加载素材列表
const loadMaterialList = async () => {
  if (!appId.value) return;

  loading.value = true;
  try {
    const res = await getWxGzhMaterialList({
      params: {
        page: pagination.value.current,
        pageSize: pagination.value.pageSize,
        appId: appId.value,
        type: currentType.value,
        name: searchKeyword.value,
      },
    });
    materialList.value = res.list || [];
    pagination.value.total = res.total || 0;
  } catch (error) {
    console.error('加载素材列表失败:', error);
    message.error('加载素材列表失败');
  } finally {
    loading.value = false;
  }
};

// 加载素材统计
const loadMaterialStats = async () => {
  if (!appId.value) return;
  try {
    const res = await getWxGzhMaterialStats({
      params: { appId: appId.value },
    });
    materialStats.value = res;
  } catch (error) {
    console.error('加载素材统计失败:', error);
  }
};

// 切换素材类型
const handleTypeChange = (type: MaterialTypeValue) => {
  currentType.value = type;
  pagination.value.current = 1;
  selectedMaterials.value = [];
  loadMaterialList();
};

// 搜索
const handleSearch = () => {
  pagination.value.current = 1;
  loadMaterialList();
};

// 分页变化
const handlePageChange = (page: number, pageSize: number) => {
  pagination.value.current = page;
  pagination.value.pageSize = pageSize;
  loadMaterialList();
};

// 刷新
const handleRefresh = () => {
  loadMaterialList();
  loadMaterialStats();
};

// 同步微信素材
const handleSync = async () => {
  if (!appId.value) return;

  syncLoading.value = true;
  try {
    await syncWxGzhMaterial({ body: { appId: appId.value } });
    message.success('同步成功');
    loadMaterialList();
    loadMaterialStats();
  } catch (error) {
    console.error('同步失败:', error);
    message.error('同步失败');
  } finally {
    syncLoading.value = false;
  }
};

// 上传素材
const handleUpload = () => {
  if (!appId.value) {
    message.warning('请先选择公众号');
    return;
  }
  uploadVisible.value = true;
};

// 预览素材
const handlePreview = (material: WxGzhMaterialInfo) => {
  currentMaterial.value = material;
  previewVisible.value = true;
};

// 删除单个素材
const handleDelete = async (material: WxGzhMaterialInfo) => {
  if (!material.id) return;
  try {
    await deleteWxGzhMaterial({ body: { ids: [material.id] } });
    message.success('删除成功');
    loadMaterialList();
    loadMaterialStats();
  } catch (error) {
    console.error('删除失败:', error);
    message.error('删除失败');
  }
};

// 批量删除
const handleBatchDelete = async () => {
  if (selectedMaterials.value.length === 0) {
    message.warning('请选择要删除的素材');
    return;
  }

  try {
    await deleteWxGzhMaterial({ body: { ids: selectedMaterials.value } });
    message.success(`成功删除 ${selectedMaterials.value.length} 个素材`);
    selectedMaterials.value = [];
    loadMaterialList();
    loadMaterialStats();
  } catch (error) {
    console.error('批量删除失败:', error);
    message.error('批量删除失败');
  }
};

// 选择素材
const handleSelectMaterial = (materialId: string, checked: boolean) => {
  if (checked) {
    selectedMaterials.value.push(materialId);
  } else {
    const index = selectedMaterials.value.indexOf(materialId);
    if (index !== -1) {
      selectedMaterials.value.splice(index, 1);
    }
  }
};

// 监听账号变化
watch(appId, (newValue) => {
  if (newValue) {
    loadMaterialList();
    loadMaterialStats();
  }
});

onMounted(() => {
  if (appId.value) {
    loadMaterialList();
    loadMaterialStats();
  }
});
</script>

<template>
  <Page>
    <!-- 账号选择 -->
    <div class="gzh-selector">
      <div class="gzh-selector-left">
        <AccountSelect v-model="appId" @change="handleAccountChange" />
      </div>
    </div>

    <!-- 素材管理主体 -->
    <div v-if="appId" class="gzh-material">
      <!-- 统计信息 -->
      <Card class="stats-card" size="small">
        <Row :gutter="16">
          <Col :span="6">
            <div class="stat-item">
              <div class="stat-value">{{ materialStats?.totalCount || 0 }}</div>
              <div class="stat-label">总素材数</div>
            </div>
          </Col>
          <Col :span="6">
            <div class="stat-item">
              <div class="stat-value">{{ materialStats?.imageCount || 0 }}</div>
              <div class="stat-label">图片</div>
            </div>
          </Col>
          <Col :span="6">
            <div class="stat-item">
              <div class="stat-value">{{ materialStats?.voiceCount || 0 }}</div>
              <div class="stat-label">音频</div>
            </div>
          </Col>
          <Col :span="6">
            <div class="stat-item">
              <div class="stat-value">{{ materialStats?.videoCount || 0 }}</div>
              <div class="stat-label">视频</div>
            </div>
          </Col>
        </Row>
      </Card>

      <!-- 操作工具栏 -->
      <Card class="toolbar-card" size="small">
        <Flex justify="space-between" align="center">
          <Space>
            <!-- 搜索 -->
            <Input.Search
              v-model:value="searchKeyword"
              placeholder="搜索素材名称"
              style="width: 200px"
              @search="handleSearch"
            >
              <template #prefix>
                <SearchOutlined />
              </template>
            </Input.Search>

            <!-- 批量操作 -->
            <Button
              v-if="hasSelectedMaterials"
              danger
              type="primary"
              @click="handleBatchDelete"
            >
              <DeleteOutlined />
              批量删除 ({{ selectedMaterials.length }})
            </Button>
          </Space>

          <Space>
            <!-- 刷新 -->
            <Button @click="handleRefresh">
              <ReloadOutlined />
              刷新
            </Button>

            <!-- 同步 -->
            <Button :loading="syncLoading" @click="handleSync">
              <SyncOutlined />
              同步微信
            </Button>

            <!-- 上传 -->
            <Button type="primary" @click="handleUpload">
              <CloudUploadOutlined />
              上传素材
            </Button>
          </Space>
        </Flex>
      </Card>

      <!-- 素材类型标签页 -->
      <Card class="content-card">
        <Tabs
          :active-key="currentType.toString()"
          @change="(key) => handleTypeChange(key as MaterialTypeValue)"
        >
          <Tabs.TabPane
            v-for="option in materialTypeOptions"
            :key="option.value.toString()"
            :tab="option.label"
          >
            <!-- 素材列表 -->
            <Spin :spinning="loading">
              <div v-if="materialList.length === 0" class="empty-container">
                <Empty description="暂无素材" />
              </div>
              <div v-else class="material-grid">
                <div
                  v-for="material in materialList"
                  :key="material.id"
                  class="material-item"
                  :class="{
                    selected:
                      material.id && selectedMaterials.includes(material.id),
                  }"
                >
                  <!-- 选择框 -->
                  <div class="material-checkbox">
                    <input
                      type="checkbox"
                      :checked="
                        !!(
                          material.id && selectedMaterials.includes(material.id)
                        )
                      "
                      @change="
                        (e) =>
                          material.id &&
                          handleSelectMaterial(
                            material.id,
                            (e.target as HTMLInputElement).checked,
                          )
                      "
                    />
                  </div>

                  <!-- 素材预览 -->
                  <div
                    class="material-preview"
                    @click="handlePreview(material)"
                  >
                    <!-- 图片预览 -->
                    <div
                      v-if="material.type === MaterialType.IMAGE"
                      class="image-preview"
                    >
                      <img
                        :src="material.URL"
                        :alt="material.name"
                        referrerPolicy="no-referrer"
                      />
                    </div>

                    <!-- 音频预览 -->
                    <div
                      v-else-if="material.type === MaterialType.VOICE"
                      class="voice-preview"
                    >
                      <div class="voice-icon">🎵</div>
                      <div class="voice-name">{{ material.name }}</div>
                    </div>

                    <!-- 视频预览 -->
                    <div
                      v-else-if="material.type === MaterialType.VIDEO"
                      class="video-preview"
                    >
                      <div class="video-icon" v-if="!material.coverURL">🎬</div>
                      <div class="video-cover" v-if="material.coverURL">
                        <img
                          :src="material.coverURL"
                          alt="视频封面"
                          referrerPolicy="no-referrer"
                        />
                      </div>
                      <div class="video-name">{{ material.name }}</div>
                    </div>
                  </div>

                  <!-- 素材信息 -->
                  <div class="material-info">
                    <div class="material-name" :title="material.name">
                      {{ material.name }}
                    </div>
                    <div class="material-meta">
                      <span class="material-time">
                        {{
                          material.updateTime
                            ? new Date(material.updateTime).toLocaleDateString()
                            : '-'
                        }}
                      </span>
                    </div>
                  </div>

                  <!-- 操作按钮 -->
                  <div class="material-actions">
                    <Tooltip title="预览">
                      <Button
                        size="small"
                        type="text"
                        @click="handlePreview(material)"
                      >
                        <EyeOutlined />
                      </Button>
                    </Tooltip>
                    <Tooltip title="删除">
                      <Button
                        size="small"
                        type="text"
                        danger
                        @click="handleDelete(material)"
                      >
                        <DeleteOutlined />
                      </Button>
                    </Tooltip>
                  </div>
                </div>
              </div>

              <!-- 分页 -->
              <div v-if="materialList.length > 0" class="pagination-container">
                <Pagination
                  v-model:current="pagination.current"
                  v-model:page-size="pagination.pageSize"
                  :total="pagination.total"
                  :show-size-changer="true"
                  :show-quick-jumper="true"
                  :show-total="(total) => `共 ${total} 条`"
                  @change="handlePageChange"
                />
              </div>
            </Spin>
          </Tabs.TabPane>
        </Tabs>
      </Card>
    </div>

    <!-- 未选择账号提示 -->
    <div v-else class="no-account">
      <Empty description="请先选择公众号账号" />
    </div>

    <!-- 预览和上传模态框 -->
    <MaterialUpload
      v-model:open="uploadVisible"
      :data="{ appId, type: currentType }"
      @close="uploadVisible = false"
      @success="handleRefresh"
    />

    <!-- 素材预览模态框 -->
    <MaterialPreview
      :open="previewVisible"
      :data="currentMaterial"
      @close="previewVisible = false"
    />
  </Page>
</template>

<style scoped lang="scss">
.gzh-selector {
  margin-bottom: 16px;
}

.gzh-material {
  .stats-card {
    margin-bottom: 16px;

    .stat-item {
      text-align: center;

      .stat-value {
        font-size: 24px;
        font-weight: 600;
        line-height: 1;
        color: #1890ff;
      }

      .stat-label {
        margin-top: 4px;
        font-size: 12px;
        color: #666;
      }
    }
  }

  .toolbar-card {
    margin-bottom: 16px;
  }

  .content-card {
    .material-grid {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
      gap: 16px;
      margin-bottom: 24px;

      .material-item {
        position: relative;
        overflow: hidden;
        cursor: pointer;
        border: 1px solid #d9d9d9;
        border-radius: 8px;
        transition: all 0.3s;

        &:hover {
          border-color: #1890ff;
          box-shadow: 0 2px 8px rgb(0 0 0 / 10%);
        }

        &.selected {
          background-color: #f0f8ff;
          border-color: #1890ff;
        }

        .material-checkbox {
          position: absolute;
          top: 8px;
          left: 8px;
          z-index: 2;

          input[type='checkbox'] {
            width: 16px;
            height: 16px;
          }
        }

        .material-preview {
          display: flex;
          align-items: center;
          justify-content: center;
          height: 150px;
          background-color: #fafafa;

          .image-preview {
            width: 100%;
            height: 100%;

            img {
              width: 100%;
              height: 100%;
              object-fit: cover;
            }
          }

          .voice-preview,
          .video-preview {
            text-align: center;

            .voice-icon,
            .video-icon {
              margin-bottom: 8px;
              font-size: 48px;
            }

            .voice-name,
            .video-name {
              padding: 0 8px;
              font-size: 12px;
              color: #666;
              word-break: break-all;
            }
          }

          .thumb-preview {
            width: 100%;
            height: 100%;

            img {
              width: 100%;
              height: 100%;
              object-fit: cover;
            }
          }
        }

        .material-info {
          padding: 12px;

          .material-name {
            margin-bottom: 4px;
            overflow: hidden;
            text-overflow: ellipsis;
            font-size: 14px;
            font-weight: 500;
            color: #333;
            white-space: nowrap;
          }

          .material-meta {
            display: flex;
            justify-content: space-between;
            font-size: 12px;
            color: #999;

            .material-size {
              color: #666;
            }
          }
        }

        .material-actions {
          position: absolute;
          top: 8px;
          right: 8px;
          display: none;
          padding: 4px;
          background-color: rgb(255 255 255 / 90%);
          border-radius: 4px;
        }

        &:hover .material-actions {
          display: flex;
        }
      }
    }

    .pagination-container {
      padding: 16px 0;
      text-align: center;
    }

    .empty-container {
      padding: 40px 0;
      text-align: center;
    }
  }
}

.no-account {
  padding: 40px 0;
  text-align: center;
}
</style>
