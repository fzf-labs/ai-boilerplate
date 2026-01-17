<script lang="ts" setup>
import type { VbenFormSchema } from '#/adapter/form';
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { UserMembershipChangeInfo } from '#/api/v1/user-membership-change';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { formatDateTime } from '@vben/utils';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getUserMembershipChangeList } from '#/api/v1/user-membership-change';

const userId = ref('');
const userLabel = ref('');

const sourceTypeMap: Record<string, string> = {
  order: '订单',
  activation_code: '激活码',
  admin: '后台',
};

const membershipTypeMap: Record<string, string> = {
  normal: '普通会员',
  vip: 'VIP会员',
  svip: 'SVIP会员',
};

const membershipStatusMap: Record<number, string> = {
  1: '正常',
  [-1]: '禁用',
};

const sourceTypeOptions = [
  { label: '全部', value: '' },
  { label: '订单', value: 'order' },
  { label: '激活码', value: 'activation_code' },
  { label: '后台', value: 'admin' },
];

const formSchema: VbenFormSchema[] = [
  {
    fieldName: 'sourceType',
    label: '来源类型',
    component: 'Select',
    componentProps: {
      placeholder: '请选择来源类型',
      allowClear: true,
      options: sourceTypeOptions,
    },
  },
  {
    fieldName: 'sourceId',
    label: '来源ID',
    component: 'Input',
    componentProps: {
      placeholder: '请输入来源ID',
      allowClear: true,
    },
  },
];

const formatMembershipType = (value?: string) => {
  if (!value) {
    return '-';
  }
  return membershipTypeMap[value] || value;
};

const formatMembershipExpiredAt = (value?: string) => {
  if (!value) {
    return '永不过期';
  }
  return formatDateTime(value);
};

const formatAutoRenew = (value?: number) => {
  if (value === 1) {
    return '是';
  }
  if (value === 0) {
    return '否';
  }
  return '-';
};

const formatAutoRenewDays = (value?: number) => {
  if (!value) {
    return '无';
  }
  return `${value}天`;
};

const formatMembershipStatus = (value?: number) => {
  if (value === undefined || value === null) {
    return '-';
  }
  return membershipStatusMap[value] || `${value}`;
};

const title = computed(() => {
  if (!userLabel.value) {
    return '会员变更记录';
  }
  return `会员变更记录 - ${userLabel.value}`;
});

const [Modal, modalApi] = useVbenModal({
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      userId.value = '';
      userLabel.value = '';
      return;
    }
    const data = modalApi.getData<{
      id?: string;
      nickname?: string;
      phone?: string;
    }>();
    userId.value = data?.id || '';
    userLabel.value = data?.nickname || data?.phone || data?.id || '';
    if (userId.value) {
      gridApi.query();
    }
  },
});

const columns: VxeTableGridOptions['columns'] = [
  {
    field: 'sourceType',
    title: '来源类型',
    minWidth: 120,
    formatter: ({ cellValue }) => {
      return sourceTypeMap[cellValue] || cellValue || '-';
    },
  },
  {
    field: 'sourceId',
    title: '来源ID',
    minWidth: 180,
    showOverflow: 'tooltip',
  },
  {
    title: '变更前',
    align: 'center',
    children: [
      {
        field: 'before.membershipType',
        title: '会员类型',
        minWidth: 120,
        align: 'center',
        formatter: ({ cellValue }) => formatMembershipType(cellValue),
      },
      {
        field: 'before.expiredAt',
        title: '到期时间',
        minWidth: 180,
        formatter: ({ cellValue }) => formatMembershipExpiredAt(cellValue),
      },
      {
        field: 'before.autoRenew',
        title: '自动续费',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatAutoRenew(cellValue),
      },
      {
        field: 'before.autoRenewDays',
        title: '续费天数',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatAutoRenewDays(cellValue),
      },
      {
        field: 'before.status',
        title: '状态',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatMembershipStatus(cellValue),
      },
    ],
  },
  {
    title: '变更后',
    align: 'center',
    children: [
      {
        field: 'after.membershipType',
        title: '会员类型',
        minWidth: 120,
        align: 'center',
        formatter: ({ cellValue }) => formatMembershipType(cellValue),
      },
      {
        field: 'after.expiredAt',
        title: '到期时间',
        minWidth: 180,
        formatter: ({ cellValue }) => formatMembershipExpiredAt(cellValue),
      },
      {
        field: 'after.autoRenew',
        title: '自动续费',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatAutoRenew(cellValue),
      },
      {
        field: 'after.autoRenewDays',
        title: '续费天数',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatAutoRenewDays(cellValue),
      },
      {
        field: 'after.status',
        title: '状态',
        minWidth: 100,
        align: 'center',
        formatter: ({ cellValue }) => formatMembershipStatus(cellValue),
      },
    ],
  },
  {
    field: 'remark',
    title: '备注',
    minWidth: 160,
    showOverflow: 'tooltip',
  },
  {
    field: 'createdAt',
    title: '创建时间',
    minWidth: 180,
    formatter: 'formatDateTime',
  },
];

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: formSchema,
  },
  gridOptions: {
    columns,
    height: 460,
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          if (!userId.value) {
            return {
              total: 0,
              list: [],
            };
          }
          return await getUserMembershipChangeList({
            params: {
              userId: userId.value,
              page: page.currentPage,
              pageSize: page.pageSize,
              sourceType: formValues?.sourceType ?? '',
              sourceId: formValues?.sourceId ?? '',
            },
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
  } as VxeTableGridOptions<UserMembershipChangeInfo>,
});

defineExpose({ modalApi });
</script>

<template>
  <Modal :title="title" class="w-full max-w-6xl">
    <Grid table-title="会员变更记录" />
  </Modal>
</template>
