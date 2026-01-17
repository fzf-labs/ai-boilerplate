<script lang="ts" setup>
import type { VxeTableGridOptions } from '#/adapter/vxe-table';
import type * as AiAudioRecordApi from '#/api/v1/ai-audio-record';
import type { SysAdminInfo } from '#/api/v1/sys-admin';

import { onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';

import { Button, message, Switch } from 'ant-design-vue';

import {
  TABLE_ACTION_ICON,
  TableAction,
  useVbenVxeGrid,
} from '#/adapter/vxe-table';
import {
  deleteAiAudioRecord,
  getAiAudioRecordList,
} from '#/api/v1/ai-audio-record';
import { getSysAdminSelector } from '#/api/v1/sys-admin';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';

const userList = ref<SysAdminInfo[]>([]); // 用户列表

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 删除 */
async function handleDelete(row: AiAudioRecordApi.AiAudioRecordInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.id]),
    key: 'action_key_msg',
  });
  try {
    if (!row.id) {
      return;
    }
    await deleteAiAudioRecord({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.id]),
      key: 'action_key_msg',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

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
          return await getAiAudioRecordList({
            params: {
              page: page.currentPage,
              pageSize: page.pageSize,
              ...formValues,
            } as AiAudioRecordApi.GetAiAudioRecordListParams,
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
  } as VxeTableGridOptions<AiAudioRecordApi.AiAudioRecordInfo>,
});

onMounted(async () => {
  // 获得下拉数据
  const res = await getSysAdminSelector();
  userList.value = res.list;
});
</script>

<template>
  <Page auto-content-height>
    <Grid table-title="音乐管理列表">
      <template #toolbar-tools>
        <TableAction :actions="[]" />
      </template>

      <template #adminId="{ row }">
        <span>
          {{ userList.find((item) => item.id === row.adminId)?.nickname }}
        </span>
      </template>
      <template #content="{ row }">
        <Button
          type="link"
          v-if="row.audioURL?.length > 0"
          :href="row.audioURL"
          target="_blank"
          class="p-0"
        >
          音乐
        </Button>
        <Button
          type="link"
          v-if="row.videoURL?.length > 0"
          :href="row.videoURL"
          target="_blank"
          class="p-0 !pl-1"
        >
          视频
        </Button>
        <Button
          type="link"
          v-if="row.imageURL?.length > 0"
          :href="row.imageURL"
          target="_blank"
          class="p-0 !pl-1"
        >
          封面
        </Button>
      </template>
      <template #publicStatus="{ row }">
        <Switch v-model:checked="row.publicStatus" :disabled="true" />
      </template>
      <template #actions="{ row }">
        <TableAction
          :actions="[
            {
              label: $t('common.delete'),
              type: 'link',
              danger: true,
              icon: TABLE_ACTION_ICON.DELETE,
              auth: ['ai:music:delete'],
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
