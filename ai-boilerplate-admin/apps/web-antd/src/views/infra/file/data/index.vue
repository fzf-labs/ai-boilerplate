<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { FileDatumInfo } from '#/api/v1/file-data';

import { Page, useVbenModal } from '@vben/common-ui';

import { CloudUploadOutlined } from '@ant-design/icons-vue';
import { useClipboard } from '@vueuse/core';
import { Button, Image, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteFileDatum, getFileDatumList } from '#/api/v1/file-data';
import { $t } from '#/locales';

import { useGridColumns, useGridFormSchema } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  showConfirmButton: false, // 不显示确认按钮
  connectedComponent: Form, // 连接表单组件
  destroyOnClose: true, // 关闭时销毁组件
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 上传文件 */
function onUpload() {
  formModalApi.setData(null).open();
}

/** 复制链接到剪贴板 */
const { copy } = useClipboard({ legacy: true });
async function onCopyUrl(row: FileDatumInfo) {
  if (!row.url) {
    message.error('文件 URL 为空');
    return;
  }

  try {
    await copy(row.url);
    message.success('复制成功');
  } catch {
    message.error('复制失败');
  }
}

/** 打开 URL */
function openUrl(url?: string) {
  if (url) {
    window.open(url, '_blank');
  }
}

/** 删除文件 */
async function onDelete(row: FileDatumInfo) {
  if (!row.id) return;
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name || row.path]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteFileDatum({ body: { id: row.id } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name || row.path]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<FileDatumInfo>) {
  switch (code) {
    case 'copyUrl': {
      onCopyUrl(row);
      break;
    }
    case 'delete': {
      onDelete(row);
      break;
    }
  }
}

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useGridFormSchema(),
  },
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getFileDatumList({
            params: {
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
  } as VxeTableGridOptions<FileDatumInfo>,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid table-title="文件列表">
      <template #toolbar-tools>
        <Button type="primary" @click="onUpload">
          <CloudUploadOutlined class="size-5" />
          上传文件
        </Button>
      </template>
      <template #file-content="{ row }">
        <Image v-if="row.ext && row.ext.includes('image')" :src="row.url" />
        <Button
          v-else-if="row.ext && row.ext.includes('pdf')"
          type="link"
          @click="() => openUrl(row.url)"
        >
          预览
        </Button>
        <Button v-else type="link" @click="() => openUrl(row.url)">
          下载
        </Button>
      </template>
    </Grid>
  </Page>
</template>
