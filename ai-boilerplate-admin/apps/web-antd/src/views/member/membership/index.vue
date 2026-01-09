<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetMembershipListParams,
  MembershipInfo,
} from '#/api/v1/membership';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteMemberMembership,
  getMemberMembershipList,
  updateMemberMembershipStatus,
} from '#/api/v1/membership';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import BenefitModal from './modules/benefit-modal.vue';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

const [BenefitManageModal, benefitManageModalApi] = useVbenModal({
  fullscreen: true,
  connectedComponent: BenefitModal,
  destroyOnClose: true,
  showConfirmButton: false,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建会员类型 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑会员类型 */
function onEdit(row: MembershipInfo) {
  formModalApi.setData(row).open();
}

/** 权益管理 */
function onBenefitManage(row: MembershipInfo) {
  benefitManageModalApi.setData(row).open();
}

/** 删除会员类型 */
async function onDelete(row: MembershipInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMemberMembership({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(newStatus: number, row: MembershipInfo) {
  try {
    await updateMemberMembershipStatus({
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
function onActionClick({ code, row }: OnActionClickParams<MembershipInfo>) {
  switch (code) {
    case 'benefit': {
      onBenefitManage(row);
      break;
    }
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
          return await getMemberMembershipList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetMembershipListParams,
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
  } as VxeTableGridOptions<MembershipInfo>,
});
</script>

<template>
  <Page auto-content-height description="会员类型信息管理">
    <FormModal @success="onRefresh" />
    <BenefitManageModal />

    <Grid table-title="会员类型列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['member:membership:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['会员类型']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
