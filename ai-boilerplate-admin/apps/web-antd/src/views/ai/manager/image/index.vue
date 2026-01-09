<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type { AiImageRecordApi } from '#/api/ai/manager/image';
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { confirm, Page } from '@vben/common-ui';

import { Image, message, Switch } from 'ant-design-vue';

import {
  TABLE_ACTION_ICON,
  TableAction,
  useVbenVxeGrid,
} from '#/adapter/vxe-table';
import {
  deleteAiImageRecord,
  getAiImageRecordList,
  updateAiImageRecord,
} from '#/api/ai/manager/image';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import { AiImageStatusEnum, useGridColumns, useGridFormSchema } from './data';

const userList = ref<SysAdminInfo[]>([]); // 用户列表

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 删除 */
async function handleDelete(row: AiImageRecordApi.AiImageRecordInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.id]),
    key: 'action_key_msg',
  });
  try {
    await deleteAiImageRecord(row.id);
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.id]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 修改是否发布 */
const handleUpdatePublicStatusChange = async (
  row: AiImageRecordApi.AiImageRecordInfo,
) => {
  try {
    // 修改状态的二次确认
    const text = row.publicStatus ? '公开' : '私有';
    await confirm(`确认要"${text}"该图片吗?`).then(async () => {
      await updateAiImageRecord({
        ...row,
        publicStatus: row.publicStatus,
      });
      onRefresh();
    });
  } catch {
    row.publicStatus = !row.publicStatus;
  }
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getAiImageRecordList({
            page: page.currentPage,
            pageSize: page.pageSize,
            ...formValues,
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: true,
      search: true,
    },
  } as VxeTableGridOptions<AiImageRecordApi.AiImageRecordInfo>,
});

onMounted(async () => {
  // 获得下拉数据
  const res = await getSysAdminSelector();
  userList.value = res.list;
});
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="绘画管理列表">
      <template #toolbar-tools>
        <TableAction :actions="[]" />
      </template>
      <template #picURL="{ row }">
        <Image :src="row.picURL" class="h-20 w-20" />
      </template>
      <template #adminId="{ row }">
        <span>
          {{ userList.find((item) => item.id === row.adminId)?.nickname }}
        </span>
      </template>
      <template #publicStatus="{ row }">
        <Switch
          v-model:checked="row.publicStatus"
          @change="handleUpdatePublicStatusChange(row)"
          :disabled="row.status !== AiImageStatusEnum.SUCCESS"
        />
      </template>
      <template #actions="{ row }">
        <TableAction
          :actions="[
            {
              label: $t('common.delete'),
              type: 'link',
              danger: true,
              icon: TABLE_ACTION_ICON.DELETE,
              auth: ['ai:image:delete'],
              popConfirm: {
                title: $t('ui.actionMessage.deleteConfirm', [row.id]),
                confirm: handleDelete.bind(null, row),
              },
            },
          ]"
        />
      </template>
    </Grid>
  </Page>
</template>
