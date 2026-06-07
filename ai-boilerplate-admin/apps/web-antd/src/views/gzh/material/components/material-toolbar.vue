<script lang="ts" setup>
import { computed, ref, watch } from 'vue';

import {
  CloudUploadOutlined,
  DeleteOutlined,
  DownloadOutlined,
  FilterOutlined,
  ReloadOutlined,
  SearchOutlined,
  SyncOutlined,
} from '@ant-design/icons-vue';
import type { SelectValue } from 'ant-design-vue/es/select';
import type { SegmentedValue } from 'ant-design-vue/es/segmented/src/segmented';
import {
  Button,
  DatePicker,
  Flex,
  Input,
  Segmented,
  Select,
  Space,
} from 'ant-design-vue';

import {
  MaterialType,
  materialGroupModeOptions,
  materialTypeOptions,
  type MaterialGroupMode,
  type MaterialToolbarFilters,
  type MaterialTypeValue,
} from '../helpers'

const props = defineProps<{
  filters?: MaterialToolbarFilters;
  exportLoading?: boolean;
  loading?: boolean;
  tagOptions?: Array<{ label: string; value: string }>;
  selectedCount?: number;
  syncLoading?: boolean;
}>();

const emits = defineEmits<{
  batchDelete: [];
  clearSelection: [];
  export: [];
  groupChange: [groupMode: MaterialGroupMode];
  refresh: [];
  search: [filters: MaterialToolbarFilters];
  sync: [];
  tagChange: [tag?: string];
  upload: [];
}>();

const showAdvancedFilter = ref(false);
const localFilters = ref<MaterialToolbarFilters>({
  keyword: props.filters?.keyword || '',
  type: props.filters?.type,
  dateRange: props.filters?.dateRange,
  groupMode: props.filters?.groupMode || 'none',
  tag: props.filters?.tag,
});

watch(
  () => props.filters,
  (filters) => {
    localFilters.value = {
      keyword: filters?.keyword || '',
      type: filters?.type,
      dateRange: filters?.dateRange,
      groupMode: filters?.groupMode || 'none',
      tag: filters?.tag,
    };
  },
  { deep: true, immediate: true },
);

// 素材类型选项
const materialTypeSelectOptions = [...materialTypeOptions];

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
    type: MaterialType.IMAGE,
    dateRange: undefined,
    groupMode: 'none',
    tag: undefined,
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

const handleExport = () => {
  emits('export');
};

// 批量删除
const handleBatchDelete = () => {
  emits('batchDelete');
};

// 清空选择
const handleClearSelection = () => {
  emits('clearSelection');
};

const handleTagChange = (value: SelectValue) => {
  const nextTag =
    typeof value === 'string' || typeof value === 'number'
      ? String(value)
      : undefined;

  localFilters.value.tag = nextTag;
  emits('tagChange', nextTag);
  handleSearch();
};

const handleGroupChange = (value: SegmentedValue) => {
  const nextGroupMode: MaterialGroupMode =
    value === 'date' || value === 'tag' || value === 'none'
      ? value
      : 'none';

  localFilters.value.groupMode = nextGroupMode;
  emits('groupChange', nextGroupMode);
  handleSearch();
};

// 监听关键词输入
const handleKeywordChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  localFilters.value.keyword = target.value;
};

// 监听类型变化
const handleTypeChange = (value: SelectValue) => {
  const nextType =
    typeof value === 'string' || typeof value === 'number'
      ? (String(value) as MaterialTypeValue)
      : MaterialType.IMAGE;

  localFilters.value.type = nextType;
  handleSearch();
};

// 监听日期范围变化
const handleDateRangeChange = (_dates: any, dateStrings: [string, string]) => {
  localFilters.value.dateRange =
    dateStrings[0] && dateStrings[1] ? dateStrings : undefined;
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
            :options="materialTypeSelectOptions"
            placeholder="选择类型"
            style="width: 120px"
            @change="handleTypeChange"
          />

          <Select
            v-if="tagOptions?.length"
            :allow-clear="true"
            :options="tagOptions"
            :value="localFilters.tag"
            placeholder="标签"
            style="width: 140px"
            @change="handleTagChange"
          />

          <Segmented
            :options="materialGroupModeOptions"
            :value="localFilters.groupMode || 'none'"
            @change="handleGroupChange"
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

          <!-- 导出 -->
          <Button :loading="exportLoading" @click="handleExport">
            <DownloadOutlined />
            导出
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
