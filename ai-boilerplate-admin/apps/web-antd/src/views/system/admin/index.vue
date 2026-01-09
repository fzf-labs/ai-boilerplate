<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysAdminInfo } from '#/api/v1/sys-admin';
import type { SysDeptInfo } from '#/api/v1/sys-dept';

import { ref } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message, Modal } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteSysAdmin,
  getSysAdminList,
  updateSysAdminStatus,
} from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DeptTree from './modules/dept-tree.vue';
import Detail from './modules/detail.vue';
import Form from './modules/form.vue';
import ResetPasswordForm from './modules/reset-password-form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [ResetPasswordModal, resetPasswordModalApi] = useVbenModal({
  connectedComponent: ResetPasswordForm,
  destroyOnClose: true,
});

const [DetailModal, detailModalApi] = useVbenModal({
  connectedComponent: Detail,
  destroyOnClose: true,
  showConfirmButton: false,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}
/** 选择部门 */
const searchDeptId = ref<string | undefined>(undefined);
async function onDeptSelect(dept: SysDeptInfo) {
  searchDeptId.value = dept.id;
  onRefresh();
}

/** 创建用户 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑用户 */
function onEdit(row: SysAdminInfo) {
  formModalApi.setData(row).open();
}

/** 删除用户 */
async function onDelete(row: SysAdminInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.nickname]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSysAdmin({ body: { id: row.id as string } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.username]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 重置密码 */
function onResetPassword(row: SysAdminInfo) {
  resetPasswordModalApi.setData(row).open();
}

/** 查看详情 */
function onDetail(row: SysAdminInfo) {
  detailModalApi.setData(row).open();
}

// TODO @芋艿：后续怎么简化一下 confirm 的实现。
/** 更新用户状态 */
async function onStatusChange(
  newStatus: number,
  row: SysAdminInfo,
): Promise<boolean | undefined> {
  return new Promise((resolve, reject) => {
    // 启用 1 禁用 -1
    let newStatusText = '启用';
    if (newStatus === -1) {
      newStatusText = '禁用';
    }
    Modal.confirm({
      title: '切换状态',
      content: `你要将${row.username}的状态切换为【${newStatusText}】吗？`,
      onCancel() {
        reject(new Error('已取消'));
      },
      onOk() {
        // 更新用户状态
        updateSysAdminStatus({
          body: { id: row.id as string, status: newStatus },
        })
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
function onActionClick({ code, row }: OnActionClickParams<SysAdminInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'detail': {
      onDetail(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'reset-password': {
      onResetPassword(row);
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
          return await getSysAdminList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
              deptId: searchDeptId.value,
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
  } as VxeTableGridOptions<SysAdminInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <ResetPasswordModal @success="onRefresh" />
    <DetailModal />

    <div class="flex h-full">
      <!-- 左侧部门树 -->
      <div class="w-1/6 pr-3">
        <DeptTree @select="onDeptSelect" />
      </div>
      <!-- 右侧用户列表 -->
      <div class="w-5/6">
        <Grid table-title="用户列表">
          <template #toolbar-tools>
            <Button
              type="primary"
              @click="onCreate"
              v-access:code="['system:admin:create']"
            >
              <Plus class="size-5" />
              {{ $t('ui.actionTitle.create', ['用户']) }}
            </Button>
          </template>
        </Grid>
      </div>
    </div>
  </Page>
</template>
