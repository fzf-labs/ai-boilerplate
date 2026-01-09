<script lang="ts" setup>
import { computed, ref } from 'vue';

import {
  CloudUploadOutlined,
  DeleteOutlined,
  FilterOutlined,
  ReloadOutlined,
  SearchOutlined,
  SyncOutlined,
} from '@ant-design/icons-vue';
import { Button, DatePicker, Flex, Input, Select, Space } from 'ant-design-vue';

type MaterialTypeValue = 'image' | 'video' | 'voice';

interface FilterOptions {
  keyword: string;
  type?: MaterialTypeValue;
  dateRange?: [string, string];
  sizeRange?: [number, number];
}

const props = defineProps<{
  filters?: FilterOptions;
  loading?: boolean;
  selectedCount?: number;
  syncLoading?: boolean;
}>();

const emits = defineEmits<{
  batchDelete: [];
  clearSelection: [];
  refresh: [];
  search: [filters: FilterOptions];
  sync: [];
  upload: [];
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

const showAdvancedFilter = ref(false);
const localFilters = ref<FilterOptions>({
  keyword: props.filters?.keyword || '',
  type: props.filters?.type,
  dateRange: props.filters?.dateRange,
  sizeRange: props.filters?.sizeRange,
});

// 素材类型选项
const materialTypeOptions = [
  { label: '全部类型', value: undefined },
  { label: MaterialTypeLabels[MaterialType.IMAGE], value: MaterialType.IMAGE },
  { label: MaterialTypeLabels[MaterialType.VOICE], value: MaterialType.VOICE },
  { label: MaterialTypeLabels[MaterialType.VIDEO], value: MaterialType.VIDEO },
];

// 文件大小选项
const sizeRangeOptions = [
  { label: '全部大小', value: 'all' },
  { label: '小于 1MB', value: 'small' },
  { label: '1MB - 5MB', value: 'medium' },
  { label: '5MB - 10MB', value: 'large' },
  { label: '大于 10MB', value: 'xlarge' },
];

// 大小范围映射
const sizeRangeMap: Record<string, [number, number] | undefined> = {
  all: undefined,
  small: [0, 1024 * 1024],
  medium: [1024 * 1024, 5 * 1024 * 1024],
  large: [5 * 1024 * 1024, 10 * 1024 * 1024],
  xlarge: [10 * 1024 * 1024, Number.MAX_SAFE_INTEGER],
};

// 计算属性
const hasSelectedItems = computed(() => (props.selectedCount || 0) > 0);

// 搜索
const handleSearch = () => {
  emits('search', { ...localFilters.value });
};

// 重置筛选
const handleResetFilters = () => {
  localFilters.value = {
    keyword: '',
    type: undefined,
    dateRange: undefined,
    sizeRange: undefined,
  };
  handleSearch();
};

// 切换高级筛选
const toggleAdvancedFilter = () => {
  showAdvancedFilter.value = !showAdvancedFilter.value;
};

// 刷新
const handleRefresh = () => {
  emits('refresh');
};

// 同步
const handleSync = () => {
  emits('sync');
};

// 上传
const handleUpload = () => {
  emits('upload');
};

// 批量删除
const handleBatchDelete = () => {
  emits('batchDelete');
};

// 清空选择
const handleClearSelection = () => {
  emits('clearSelection');
};

// 监听关键词输入
const handleKeywordChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  localFilters.value.keyword = target.value;
};

// 监听类型变化
const handleTypeChange = (value: any) => {
  localFilters.value.type = value;
  handleSearch();
};

// 监听日期范围变化
const handleDateRangeChange = (_dates: any, dateStrings: [string, string]) => {
  localFilters.value.dateRange =
    dateStrings[0] && dateStrings[1] ? dateStrings : undefined;
  handleSearch();
};

// 监听大小范围变化
const handleSizeRangeChange = (value: any) => {
  localFilters.value.sizeRange = value
    ? sizeRangeMap[value as string]
    : undefined;
  handleSearch();
};
</script>

<template>
  <div class="material-toolbar">
    <!-- 主工具栏 -->
    <div class="main-toolbar">
      <Flex justify="space-between" align="center">
        <!-- 左侧：搜索和筛选 -->
        <Space>
          <!-- 关键词搜索 -->
          <Input.Search
            v-model:value="localFilters.keyword"
            placeholder="搜索素材名称"
            style="width: 240px"
            @search="handleSearch"
            @change="handleKeywordChange"
          >
            <template #prefix>
              <SearchOutlined />
            </template>
          </Input.Search>

          <!-- 类型筛选 -->
          <Select
            v-model:value="localFilters.type"
            :options="materialTypeOptions"
            placeholder="选择类型"
            style="width: 120px"
            @change="handleTypeChange"
          />

          <!-- 高级筛选按钮 -->
          <Button @click="toggleAdvancedFilter">
            <FilterOutlined />
            高级筛选
          </Button>

          <!-- 批量操作 -->
          <div v-if="hasSelectedItems" class="batch-actions">
            <Button danger type="primary" @click="handleBatchDelete">
              <DeleteOutlined />
              批量删除 ({{ selectedCount }})
            </Button>
            <Button @click="handleClearSelection"> 清空选择 </Button>
          </div>
        </Space>

        <!-- 右侧：操作按钮 -->
        <Space>
          <!-- 刷新 -->
          <Button :loading="loading" @click="handleRefresh">
            <ReloadOutlined />
            刷新
          </Button>

          <!-- 同步微信 -->
          <Button :loading="syncLoading" @click="handleSync">
            <SyncOutlined />
            同步微信
          </Button>

          <!-- 上传素材 -->
          <Button type="primary" @click="handleUpload">
            <CloudUploadOutlined />
            上传素材
          </Button>
        </Space>
      </Flex>
    </div>

    <!-- 高级筛选面板 -->
    <div v-if="showAdvancedFilter" class="advanced-filter">
      <div class="filter-content">
        <Space wrap>
          <!-- 日期范围 -->
          <div class="filter-item">
            <label>创建时间：</label>
            <DatePicker.RangePicker
              v-model:value="localFilters.dateRange"
              format="YYYY-MM-DD"
              @change="handleDateRangeChange"
            />
          </div>

          <!-- 文件大小 -->
          <div class="filter-item">
            <label>文件大小：</label>
            <Select
              v-model:value="localFilters.sizeRange"
              :options="sizeRangeOptions"
              placeholder="选择大小范围"
              style="width: 150px"
              @change="handleSizeRangeChange"
            />
          </div>

          <!-- 操作按钮 -->
          <div class="filter-actions">
            <Button type="primary" @click="handleSearch"> 应用筛选 </Button>
            <Button @click="handleResetFilters"> 重置 </Button>
          </div>
        </Space>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.material-toolbar {
  .main-toolbar {
    padding: 16px;
    margin-bottom: 16px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 1px 3px rgb(0 0 0 / 10%);

    .batch-actions {
      display: flex;
      gap: 8px;
      padding-left: 16px;
      border-left: 1px solid #d9d9d9;
    }
  }

  .advanced-filter {
    margin-bottom: 16px;
    background-color: #fafafa;
    border: 1px solid #d9d9d9;
    border-radius: 8px;

    .filter-content {
      padding: 16px;

      .filter-item {
        display: flex;
        gap: 8px;
        align-items: center;

        label {
          font-weight: 500;
          white-space: nowrap;
        }
      }

      .filter-actions {
        display: flex;
        gap: 8px;
      }
    }
  }
}
</style>
