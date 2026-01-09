<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeGridListeners,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { DictTypeInfo } from '#/api/v1/dict-type';

import { useVbenModal } from '@vben/common-ui';
import { Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteDictType, getDictTypeList } from '#/api/v1/dict-type';
import { $t } from '#/locales';

import { useTypeGridColumns, useTypeGridFormSchema } from '../data';
import TypeForm from './type-form.vue';

const emit = defineEmits(['select']);

const [TypeFormModal, typeFormModalApi] = useVbenModal({
  connectedComponent: TypeForm,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建字典类型 */
function onCreate() {
  typeFormModalApi.setData(null).open();
}

/** 编辑字典类型 */
function onEdit(row: any) {
  typeFormModalApi.setData(row).open();
}

/** 删除字典类型 */
async function onDelete(row: DictTypeInfo) {
  const hideLoading = message.loading({
    content: $t('common.processing'),
    duration: 0,
    key: 'process_message',
  });
  try {
    await deleteDictType({ body: { id: row.id } });
    message.success({
      content: $t('common.operationSuccess'),
      key: 'process_message',
    });
    onRefresh();
  } finally {
    hideLoading();
  }
}

/** 表格操作按钮回调 */
function onActionClick({ code, row }: OnActionClickParams<DictTypeInfo>) {
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

/** 表格事件 */
const gridEvents: VxeGridListeners<DictTypeInfo> = {
  cellClick: ({ row }) => {
    emit('select', row.type);
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions: {
    schema: useTypeGridFormSchema(),
  },
  gridOptions: {
    columns: useTypeGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    proxyConfig: {
      ajax: {
        query: async ({ page }, formValues) => {
          return await getDictTypeList({
            page: page.currentPage,
            pageSize: page.pageSize,
            ...formValues,
          });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
      isCurrent: true,
    },
    toolbarConfig: {
      refresh: { code: 'query' },
      search: true,
    },
  } as VxeTableGridOptions<DictTypeInfo>,
  gridEvents,
});
</script>

<template>
  <div class="h-full">
    <TypeFormModal @success="onRefresh" />

    <Grid table-title="字典类型列表">
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['system:dict:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['字典类型']) }}
        </Button>
      </template>
    </Grid>
  </div>
</template>
