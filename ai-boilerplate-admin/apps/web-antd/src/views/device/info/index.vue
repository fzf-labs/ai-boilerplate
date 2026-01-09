<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { DeviceInfo } from '#/api/v1/device';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteDevice,
  getDeviceList,
  updateDeviceStatus,
} from '#/api/v1/device';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DeviceDetailModal, deviceDetailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 注册设备 */
function onRegister() {
  formModalApi.setData(null).open();
}

/** 查看设备详情 */
function onDetail(row: DeviceInfo) {
  deviceDetailModalApi.setData(row).open();
}

/** 删除设备 */
async function onDelete(row: DeviceInfo) {
  if (!row.sn) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteDevice({ body: { sn: row.sn } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 更新设备状态 */
async function onStatusChange(
  newStatus: number,
  row: DeviceInfo,
): Promise<boolean | undefined> {
  return new Promise((resolve, reject) => {
    // 启用 1 禁用 -1
    let newStatusText = '启用';
    if (newStatus === -1) {
      newStatusText = '禁用';
    }
    Modal.confirm({
      title: '切换状态',
      content: `你要将设备 ${row.name} 的状态切换为【${newStatusText}】吗？`,
      onCancel() {
        reject(new Error('已取消'));
      },
      onOk() {
        // 更新设备状态
        if (!row.sn) {
          reject(new Error('设备SN不存在'));
          return;
        }
        updateDeviceStatus({ body: { sn: row.sn, status: newStatus } })
          .then(() => {
            // 提示并返回成功
            message.success({
              content: $t('ui.actionMessage.operationSuccess'),
              key: 'action_process_msg',
            });
            resolve(true);
          })
          .catch((error) => {
            reject(error);
          });
      },
    });
  });
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<DeviceInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'detail': {
      onDetail(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick, onStatusChange),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getDeviceList({
            body: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            },
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<DeviceInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <DeviceDetailModal />

    <Grid table-title="设备列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onRegister"
          v-access:code="['kid:device:register']"
        >
          <Plus class="size-5" />
          注册设备
        </Button>
      </template>
    </Grid>
  </Page>
</template>
