<script lang="ts" setup>
import type { SysMenuInfo } from '#/api/v1/sys-menu';
import type { SysTenantInfo } from '#/api/v1/sys-tenant';

import { computed, ref } from 'vue';

import { useVbenModal, VbenTree } from '@vben/common-ui';
import { handleTree } from '@vben/utils';

import { message } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { getSysMenuList } from '#/api/v1/sys-menu';
import {
  createSysTenant,
  getSysTenantInfo,
  updateSysTenant,
} from '#/api/v1/sys-tenant';
import { $t } from '#/locales';

import { useFormSchema } from '../data';

// 扩展 SysTenantInfo 类型以包含 username 和 password 字段(仅用于创建租户)
interface SysTenantFormData extends SysTenantInfo {
  username?: string;
  password?: string;
}

const emit = defineEmits(['success']);
const formData = ref<SysTenantInfo>();
const getTitle = computed(() => {
  return formData.value
    ? $t('ui.actionTitle.edit', ['租户'])
    : $t('ui.actionTitle.create', ['租户']);
});

const menuTree = ref<SysMenuInfo[]>([]); // 菜单树
const menuLoading = ref(false); // 加载菜单列表

const [Form, formApi] = useVbenForm({
  layout: 'horizontal',
  schema: useFormSchema(),
  showDefaultActions: false,
});

const [Modal, modalApi] = useVbenModal({
  async onConfirm() {
    const { valid } = await formApi.validate();
    if (!valid) {
      return;
    }
    modalApi.lock();
    // 提交表单
    const data = (await formApi.getValues()) as SysTenantFormData;
    try {
      formData.value
        ? await updateSysTenant({
            body: {
              id: data.id!,
              name: data.name!,
              remark: data.remark,
              expireTime: data.expireTime,
              menuIds: data.menuIds,
              status: data.status,
            },
          })
        : await createSysTenant({
            body: {
              name: data.name!,
              remark: data.remark,
              expireTime: data.expireTime,
              menuIds: data.menuIds,
              status: data.status,
              username: data.username!,
              password: data.password!,
            },
          });
      // 关闭并提示
      await modalApi.close();
      emit('success');
      message.success({
        content: $t('ui.actionMessage.operationSuccess'),
        key: 'action_process_msg',
      });
    } finally {
      modalApi.lock(false);
    }
  },
  async onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      formData.value = undefined;
      return;
    }
    // 加载菜单列表
    await loadMenuTree();
    // 加载数据
    const data = modalApi.getData<SysTenantInfo>();
    if (!data || !data.id) {
      return;
    }
    modalApi.lock();
    try {
      const res = await getSysTenantInfo({ params: { id: data.id } });
      formData.value = res.info;
      // 设置到 values
      if (formData.value) {
        await formApi.setValues(formData.value);
      }
    } finally {
      modalApi.lock(false);
    }
  },
});

/** 加载菜单树 */
async function loadMenuTree() {
  menuLoading.value = true;
  try {
    const res = await getSysMenuList({ options: {} });
    menuTree.value = handleTree(res.list || []) as SysMenuInfo[];
  } finally {
    menuLoading.value = false;
  }
}
</script>
<template>
  <Modal :title="getTitle" class="w-[40%]">
    <Form class="mx-6">
      <template #menuIds="slotProps">
        <VbenTree
          class="max-h-[400px] overflow-y-auto"
          :loading="menuLoading"
          :tree-data="menuTree"
          multiple
          bordered
          v-bind="slotProps"
          value-field="id"
          label-field="name"
        />
      </template>
    </Form>
  </Modal>
</template>
