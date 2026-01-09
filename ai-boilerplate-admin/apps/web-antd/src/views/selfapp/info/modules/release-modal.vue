<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SelfAppInfo } from '#/api/v1/self-app';
import type {
  GetSelfAppReleaseListParams,
  SelfAppReleaseInfo,
} from '#/api/v1/self-app-release';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteSelfAppRelease,
  getSelfAppReleaseList,
  updateSelfAppReleaseStatus,
} from '#/api/v1/self-app-release';
import { $t } from '#/locales';

import {
  useReleaseGridColumns,
  useReleaseGridFormSchema,
} from './release-data';
import ReleaseDetailModal from './release-detail.vue';
import ReleaseForm from './release-form.vue';

const appInfo = ref<SelfAppInfo>();

const getTitle = computed(() => {
  return appInfo.value?.name
    ? `${appInfo.value.name} - 版本发布管理`
    : '版本发布管理';
});

const [ReleaseFormModal, releaseFormModalApi] = useVbenModal({
  connectedComponent: ReleaseForm,
  destroyOnClose: true,
});

const [ReleaseDetailModalComponent, releaseDetailModalApi] = useVbenModal({
  connectedComponent: ReleaseDetailModal,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建版本发布 */
function onCreate() {
  if (!appInfo.value) {
    message.error('请先选择应用');
    return;
  }
  releaseFormModalApi
    .setData({
      appInfo: appInfo.value,
      releaseInfo: null,
    })
    .open();
}

/** 查看版本详情 */
function onView(row: SelfAppReleaseInfo) {
  releaseDetailModalApi.setData(row).open();
}

/** 编辑版本发布 */
function onEdit(row: SelfAppReleaseInfo) {
  if (!appInfo.value) {
    message.error('应用信息不存在');
    return;
  }
  releaseFormModalApi
    .setData({
      appInfo: appInfo.value,
      releaseInfo: row,
    })
    .open();
}

/** 删除版本发布 */
async function onDelete(row: SelfAppReleaseInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [
      `${row.packageName} v${row.version}`,
    ]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSelfAppRelease({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [
        `${row.packageName} v${row.version}`,
      ]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: SelfAppReleaseInfo) {
  try {
    await updateSelfAppReleaseStatus({
      body: {
        id: row.id!,
        status: newStatus,
      },
    });
    message.success({
      content: $t('ui.actionMessage.operationSuccess'),
      key: 'action_process_msg',
    });
    onRefresh();
    return true;
  } catch {
    return false;
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SelfAppReleaseInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'view': {
      onView(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useReleaseGridFormSchema(),
  },
  gridOptions: {
    columns: useReleaseGridColumns(onActionClick, onStatusChange),
    height: '800',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          if (!appInfo.value) return { list: [], total: 0 };

          return await getSelfAppReleaseList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              packageName: appInfo.value.packageName!,
              ...formValues,
            } as GetSelfAppReleaseListParams,
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    showOverflow: 'tooltip',
    size: 'small',
    stripe: true,
    border: true,
    toolbarConfig: {
      refresh: true,
      'refresh-options': { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<SelfAppReleaseInfo>,
});

const [Modal, modalApi] = useVbenModal({
  onOpenChange: (isOpen: boolean) => {
    if (isOpen) {
      // 模态框打开时初始化数据
      const data = modalApi.getData() as SelfAppInfo;
      appInfo.value = data;
      // 刷新表格数据
      setTimeout(() => {
        onRefresh();
      }, 100);
    }
  },
});

defineExpose({ modalApi });
</script>

<template>
  <Modal
    :title="getTitle"
    class="vben-modal"
    :width="1200"
    :footer="false"
    :destroy-on-close="true"
  >
    <ReleaseFormModal @success="onRefresh" />
    <ReleaseDetailModalComponent />

    <div class="p-4">
      <Grid>
        <template #toolbar-tools>
          <Button
            type="primary"
            @click="onCreate"
            v-access:code="['self_app_release:create']"
          >
            <Plus class="size-4" />
            {{ $t('ui.actionTitle.create', ['版本发布']) }}
          </Button>
        </template>
      </Grid>
    </div>
  </Modal>
</template>

<style scoped>
:deep(.vben-modal .ant-modal-body) {
  padding: 0;
}
</style>
