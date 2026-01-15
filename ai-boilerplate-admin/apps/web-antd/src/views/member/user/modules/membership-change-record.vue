<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import {
  getUserMembershipChangeList,
  type UserMembershipChangeInfo,
} from '#/api/v1/user-membership-change';

const userId = ref('');
const userLabel = ref('');

const sourceTypeMap: Record<string, string> = {
  order: '订单',
  activation_code: '激活码',
  admin: '后台',
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
    field: 'beforeMembershipType',
    title: '变更前会员',
    minWidth: 120,
    align: 'center',
  },
  {
    field: 'afterMembershipType',
    title: '变更后会员',
    minWidth: 120,
    align: 'center',
  },
  {
    field: 'beforeExpiredAt',
    title: '变更前到期',
    minWidth: 180,
    formatter: 'formatDateTime',
  },
  {
    field: 'afterExpiredAt',
    title: '变更后到期',
    minWidth: 180,
    formatter: 'formatDateTime',
  },
  {
    field: 'durationDays',
    title: '时长(天)',
    minWidth: 100,
    align: 'center',
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
  gridOptions: {
    columns,
    height: 460,
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }) => {
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
