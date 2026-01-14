<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type {
  GetMembershipBenefitTypeListParams,
  MembershipBenefitTypeInfo,
} from '#/api/v1/membership-benefit-type';

import { Page, useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  deleteMembershipBenefitType,
  getMembershipBenefitTypeList,
  updateMembershipBenefitTypeStatus,
} from '#/api/v1/membership-benefit-type';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建权益类型 */
function onCreate() {
  formModalApi.setData(null).open();
}

/** 编辑权益类型 */
function onEdit(row: MembershipBenefitTypeInfo) {
  formModalApi.setData(row).open();
}

/** 删除权益类型 */
async function onDelete(row: MembershipBenefitTypeInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.benefitName]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteMembershipBenefitType({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.benefitName]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 状态变更 */
async function onStatusChange(
  newStatus: number,
  row: MembershipBenefitTypeInfo,
) {
  try {
    await updateMembershipBenefitTypeStatus({
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
function onActionClick({
  code,
  row,
}: OnActionClickParams<MembershipBenefitTypeInfo>) {
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
          return await getMembershipBenefitTypeList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as GetMembershipBenefitTypeListParams,
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
  } as VxeTableGridOptions<MembershipBenefitTypeInfo>,
});
</script>

<template>
  <Page auto-content-height description="管理会员权益类型，如：定位、截图、敲一敲等">
    <FormModal @success="onRefresh" />

    <Grid table-title="权益类型列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['member:membership-benefit-type:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['权益类型']) }}
        </Button>
      </template>
    </Grid>
  </Page>
</template>
