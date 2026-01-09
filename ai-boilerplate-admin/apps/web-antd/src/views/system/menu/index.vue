<script lang="ts" setup>
import type {
  OnActionClickParams,
  VxeTableGridOptions,
} from '#/adapter/vxe-table';
import type { SysMenuInfo } from '#/api/v1/sys-menu';

import { ref } from 'vue';

import { Page, useVbenModal } from '@vben/common-ui';
import { IconifyIcon, Plus } from '@vben/icons';

import { Button, message } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { deleteSysMenu, getSysMenuList } from '#/api/v1/sys-menu';
import { $t } from '#/locales';
import { SystemMenuTypeEnum } from '#/utils/constants';

import { useGridColumns } from './data';
import Form from './modules/form.vue';

const [FormModal, formModalApi] = useVbenModal({
  connectedComponent: Form,
  destroyOnClose: true,
});

/** 刷新表格 */
function onRefresh() {
  gridApi.query();
}

/** 创建菜单 */
function onCreate() {
  formModalApi.setData({}).open();
}

/** 添加下级菜单 */
function onAppend(row: SysMenuInfo) {
  formModalApi.setData({ pid: row.id }).open();
}

/** 编辑菜单 */
function onEdit(row: SysMenuInfo) {
  formModalApi.setData(row).open();
}

/** 删除菜单 */
async function onDelete(row: SysMenuInfo) {
  const hideLoading = message.loading({
    content: $t('ui.actionMessage.deleting', [row.name]),
    duration: 0,
    key: 'action_process_msg',
  });
  try {
    await deleteSysMenu({ body: { id: row.id! } });
    message.success({
      content: $t('ui.actionMessage.deleteSuccess', [row.name]),
      key: 'action_process_msg',
    });
    onRefresh();
  } catch {
    hideLoading();
  }
}

/** 表格操作按钮的回调函数 */
function onActionClick({ code, row }: OnActionClickParams<SysMenuInfo>) {
  switch (code) {
    case 'append': {
      onAppend(row);
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

/** 切换树形展开/收缩状态 */
const isExpanded = ref(false);
function toggleExpand() {
  isExpanded.value = !isExpanded.value;
  gridApi.grid.setAllTreeExpand(isExpanded.value);
}

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions: {
    columns: useGridColumns(onActionClick),
    height: 'auto',
    keepSource: true,
    pagerConfig: {
      enabled: false,
    },
    proxyConfig: {
      ajax: {
        query: async (_params) => {
          return await getSysMenuList({ options: {} });
        },
      },
    },
    rowConfig: {
      keyField: 'id',
    },
    toolbarConfig: {
      refresh: { code: 'query' },
    },
    treeConfig: {
      parentField: 'pid',
      rowField: 'id',
      transform: true,
      reserve: true,
    },
  } as VxeTableGridOptions,
});
</script>

<template>
  <Page auto-content-height>
    <FormModal @success="onRefresh" />
    <Grid>
      <template #toolbar-tools>
        <Button
          type="primary"
          @click="onCreate"
          v-access:code="['system:menu:create']"
        >
          <Plus class="size-5" />
          {{ $t('ui.actionTitle.create', ['菜单']) }}
        </Button>
        <Button class="ml-2" @click="toggleExpand">
          {{ isExpanded ? '收缩' : '展开' }}
        </Button>
      </template>
      <template #name="{ row }">
        <div class="flex w-full items-center gap-1">
          <div class="size-5 flex-shrink-0">
            <IconifyIcon
              v-if="row.type === SystemMenuTypeEnum.BUTTON"
              icon="carbon:square-outline"
              class="size-full"
            />
            <IconifyIcon
              v-else-if="row.icon"
              :icon="row.icon || 'carbon:circle-dash'"
              class="size-full"
            />
          </div>
          <span class="flex-auto">{{ $t(row.name) }}</span>
          <div class="items-center justify-end"></div>
        </div>
      </template>
    </Grid>
  </Page>
</template>
