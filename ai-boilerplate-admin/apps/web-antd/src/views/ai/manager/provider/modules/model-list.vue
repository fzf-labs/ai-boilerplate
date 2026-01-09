<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { AiProviderPlatformApi } from '#/api/ai/manager/providerPlatform';
import type { AiProviderModelApi } from '#/api/v1/ai-provider-model';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteAiProviderModel,
  getAiProviderModelList,
} from '#/api/v1/ai-provider-model';
import { $t } from '#/locales';

import { useModelColumns } from './model-data';
import ModelForm from './model-form.vue';

const [ModelFormModal, modelFormModalApi] = useVbenModal({
  connectedComponent: ModelForm,
  destroyOnClose: true,
});

const platformData = ref<AiProviderPlatformApi.AiProviderPlatformInfo>();

const getTitle = computed(() => {
  return platformData.value
    ? `模型管理 - ${platformData.value.name}`
    : '模型管理';
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建模型 */
function onCreate() {
  if (!platformData.value) {
    return;
  }
  modelFormModalApi
    .setData({
      platformId: platformData.value.id,
    })
    .open();
}

/** 编辑模型 */
function onEdit(row: AiProviderModelApi.AiProviderModelInfo) {
  modelFormModalApi.setData(row).open();
}

/** 删除模型 */
async function onDelete(row: AiProviderModelApi.AiProviderModelInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.modelName]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiProviderModel({ id: row.id });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.modelName]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({
  code,
  row,
}: OnActionClickParams<AiProviderModelApi.AiProviderModelInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: useModelColumns(onActionClick),
    height: 600,
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }) => {
          if (!platformData.value?.id) {
            return { list: [], total: 0 };
          }
          return await getAiProviderModelList({
            page: page.currentPage,
            pageSize: page.pageSize,
            platformId: platformData.value.id,
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: true,
    },
  } as VxeTableGridOptions<AiProviderModelApi.AiProviderModelInfo>,
});

const [Modal, modalApi] = useVbenModal({
  fullscreen: true,
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      platformData.value = undefined;
      return;
    }
    // 加载数据
    const data =
      modalApi.getData<AiProviderPlatformApi.AiProviderPlatformInfo>();
    if (!data || !data.id) {
      return;
    }
    platformData.value = data;
    // 刷新表格
    setTimeout(() => {
      onRefresh();
    }, 100);
  },
});
</script>

<template>
  <Modal :title="getTitle">
    <ModelFormModal @success="onRefresh" />
    <div v-if="platformData" class="model-list-container">
      <!-- 平台信息卡片 -->
      <div
        class="mb-4 rounded-lg bg-gradient-to-r from-blue-50 to-indigo-50 p-4"
      >
        <div class="grid grid-cols-2 gap-4 md:grid-cols-4">
          <div>
            <span class="text-sm text-gray-600">平台名称:</span>
            <span class="ml-2 font-semibold text-gray-900">{{
              platformData.name
            }}</span>
          </div>
          <div>
            <span class="text-sm text-gray-600">文档地址:</span>
            <a
              v-if="platformData.docURL"
              :href="platformData.docURL"
              target="_blank"
              class="ml-2 font-semibold text-blue-600 hover:text-blue-800"
            >
              查看文档
            </a>
            <span v-else class="ml-2 font-semibold text-gray-900">未设置</span>
          </div>
        </div>
      </div>

      <!-- 模型列表表格 -->
      <Grid>
        <template #toolbar-tools>
          <Button
            type="primary"
            @click="onCreate"
            v-access:code="['ai:manager:provider:model:create']"
          >
            <Plus class="size-5" />
            {{ $t('ui.actionTitle.create', ['模型']) }}
          </Button>
        </template>
      </Grid>
    </div>
  </Modal>
</template>

<style scoped>
.model-list-container {
  padding: 0 16px 16px;
}

@media (max-width: 768px) {
  .model-list-container {
    padding: 0 8px 8px;
  }
}
</style>
