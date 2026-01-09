<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { GetUserListParams, UserInfo } from '#/api/v1/user';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteUser, getUserList, updateUserStatus } from '#/api/v1/user';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import DetailModal from './modules/detail.vue';
import Form from './modules/form.vue';
import TokenModal from './modules/token.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [DetailModalComponent, detailModalApi] = useVbenModal({
  connectedComponent: DetailModal,
  destroyOnClose: true,
});

const [TokenModalComponent, tokenModalApi] = useVbenModal({
  connectedComponent: TokenModal,
  destroyOnClose: true,
  showConfirmButton: false,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建用户 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 查看用户详情 */
function onView(row: UserInfo) {
  detailModalApi.setData(row).open();
}

/** 编辑用户 */
function onEdit(row: UserInfo) {
  formModalApi.setData(row).open();
}

/** 删除用户 */
async function onDelete(row: UserInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.nickname || row.phone]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteUser({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [
        row.nickname || row.phone,
      ]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: UserInfo) {
  try {
    await updateUserStatus({
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

/** 生成测试Token */
function onGenerateToken(row: UserInfo) {
  tokenModalApi.setData({ id: row.id, nickname: row.nickname }).open();
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<UserInfo>) {
  switch (code) {
    case 'delete': {
      onDelete(row);
      break;
    }
    case 'edit': {
      onEdit(row);
      break;
    }
    case 'testToken': {
      onGenerateToken(row);
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
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick, onStatusChange),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getUserList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetUserListParams,
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
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<UserInfo>,
});
</script>

<template>
  <Page auto-content-height description="用户信息管理">
    <FormModal @success="onRefresh" />
    <DetailModalComponent />
    <TokenModalComponent />

    <Grid table-title="用户列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['member:user:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['用户']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
