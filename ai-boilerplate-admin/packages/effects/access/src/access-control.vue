<script lang="ts" setup>
import { computed } from 'vue';

import { useAccess } from './use-access';

interface Props {
  /**
   * Auth values visible to this component.
   * @default []
   */
  codes?: string[];

  /**
   * 通过什么方式来控制组件，如果是 role，则传入角色，如果是 code，则传入权限码
   * @default 'role'
   */
  type?: 'code' | 'role';

  /**
   * 多个权限值的匹配方式
   * @default 'any'
   */
  match?: 'any' | 'all';

  /**
   * 当没有传入任何权限值时，是否直接展示内容
   * @default true
   */
  showWhenEmpty?: boolean;
}

defineOptions({
  name: 'AccessControl',
});

const props = withDefaults(defineProps<Props>(), {
  codes: () => [],
  type: 'role',
  match: 'any',
  showWhenEmpty: true,
});

const { hasAccessByCodes, hasAccessByRoles } = useAccess();

const hasAuth = computed(() => {
  const { codes, match, showWhenEmpty, type } = props;
  if (codes.length === 0) {
    return showWhenEmpty;
  }

  const checker = type === 'role' ? hasAccessByRoles : hasAccessByCodes;
  if (match === 'all') {
    return codes.every(code => checker([code]));
  }

  return checker(codes);
});
</script>

<template>
  <slot v-if="hasAuth"></slot>
</template>
