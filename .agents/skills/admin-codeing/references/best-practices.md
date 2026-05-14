# Admin 开发最佳实践

## 1. 项目结构规范

### 目录组织

```
views/
├── {module}/              # 模块目录
│   ├── {feature}/        # 功能目录
│   │   ├── index.vue     # 列表页
│   │   ├── data.ts       # 配置文件
│   │   ├── composables/  # 组合式函数
│   │   │   └── use{Feature}.ts
│   │   └── modules/      # 子组件
│   │       ├── form.vue  # 表单
│   │       └── detail.vue # 详情
```

**规则：**
- 一个功能一个目录
- 复杂组件拆分到 `modules/` 目录
- 可复用逻辑提取到 `composables/` 目录
- 配置和数据分离到 `data.ts` 文件

---

## 2. 代码组织

### Composition API 规范

```vue
<script setup lang="ts">
// 1. 导入类型
import type { SystemUserApi } from '#/api/system/user';

// 2. 导入 Vue API
import { onMounted, ref, computed } from 'vue';

// 3. 导入组件
import { Page, useVbenModal } from '@vben/common-ui';
import { Button } from 'ant-design-vue';

// 4. 导入工具函数
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getUserList } from '#/api/system/user';

// 5. 导入本地模块
import { useGridColumns } from './data';
import Form from './modules/form.vue';

// 6. 定义响应式数据
const loading = ref(false);
const data = ref<SystemUserApi.User[]>([]);

// 7. 定义计算属性
const total = computed(() => data.value.length);

// 8. 定义方法
function handleClick() {
  // ...
}

// 9. 生命周期钩子
onMounted(() => {
  // ...
});
</script>
```

### Props 和 Emits 规范

```vue
<script setup lang="ts">
// Props 定义
interface Props {
  id?: string;
  name: string;
  status?: number;
}

const props = withDefaults(defineProps<Props>(), {
  status: 1,
});

// Emits 定义
const emit = defineEmits<{
  success: [data: any];
  close: [];
  update: [id: string, data: any];
}>();

// 使用
emit('success', data);
emit('update', id, data);
</script>
```

---

## 3. TypeScript 最佳实践

### 类型定义

```typescript
// ✅ 推荐：使用 namespace 组织类型
export namespace SystemUserApi {
  export interface User {
    id: string;
    username: string;
    status: number;
  }

  export interface UserQuery {
    username?: string;
    status?: number;
    page?: number;
    pageSize?: number;
  }

  export interface CreateUserReq {
    username: string;
    password: string;
  }
}

// ❌ 不推荐：直接导出多个类型
export interface User { }
export interface UserQuery { }
```

### 类型断言

```typescript
// ✅ 推荐：使用 as
const user = data as SystemUserApi.User;

// ❌ 不推荐：使用 <>
const user = <SystemUserApi.User>data;
```

### 类型守卫

```typescript
function isUser(obj: any): obj is SystemUserApi.User {
  return obj && typeof obj.id === 'string';
}

if (isUser(data)) {
  // data 的类型是 User
  console.log(data.username);
}
```

---

## 4. API 接口规范

### 函数命名

```typescript
// ✅ 推荐：语义化命名
export function getUserList(params: UserQuery) { }
export function getUserInfo(id: string) { }
export function createUser(data: User) { }
export function updateUser(data: User) { }
export function deleteUser(id: string) { }
export function updateUserStatus(id: string, status: number) { }

// ❌ 不推荐：简写
export function list() { }
export function get() { }
export function add() { }
```

### 错误处理

```typescript
// ✅ 推荐：在组件中处理错误
async function handleDelete(row: User) {
  const hide = message.loading('删除中...', 0);
  try {
    await deleteUser(row.id);
    message.success('删除成功');
    refresh();
  } catch (error) {
    message.error(error.message || '删除失败');
  } finally {
    hide();
  }
}

// ❌ 不推荐：在 API 层捕获错误
export async function deleteUser(id: string) {
  try {
    return await requestClient.post('/api/delete', { id });
  } catch (error) {
    // 不要在这里处理
    message.error('删除失败');
  }
}
```

---

## 5. 性能优化

### 列表渲染优化

```vue
<script setup>
import { computed } from 'vue';

// ✅ 推荐：使用计算属性过滤数据
const filteredList = computed(() => {
  return list.value.filter(item => item.status === 1);
});

// ❌ 不推荐：在模板中过滤
</script>

<template>
  <!-- ✅ 推荐 -->
  <div v-for="item in filteredList" :key="item.id">
    {{ item.name }}
  </div>

  <!-- ❌ 不推荐 -->
  <div v-for="item in list.filter(i => i.status === 1)" :key="item.id">
    {{ item.name }}
  </div>
</template>
```

### 组件懒加载

```typescript
// ✅ 推荐：路由懒加载
{
  path: '/user',
  component: () => import('#/views/system/user/index.vue'),
}

// ✅ 推荐：动态导入组件
const FormModal = defineAsyncComponent(
  () => import('./modules/form.vue')
);
```

### 防抖节流

```typescript
import { useDebounceFn, useThrottleFn } from '@vueuse/core';

// 防抖：输入框搜索
const debouncedSearch = useDebounceFn((value: string) => {
  search(value);
}, 300);

// 节流：滚动加载
const throttledScroll = useThrottleFn(() => {
  loadMore();
}, 1000);
```

### 虚拟滚动

```vue
<template>
  <!-- 大列表使用虚拟滚动 -->
  <VirtualList
    :data="largeList"
    :item-height="50"
    :height="600"
  >
    <template #default="{ item }">
      <div>{{ item.name }}</div>
    </template>
  </VirtualList>
</template>
```

---

## 6. 状态管理

### 何时使用 Pinia

**需要使用：**
- 跨页面共享的数据
- 需要持久化的数据
- 复杂的业务逻辑

**不需要使用：**
- 页面内部的临时状态
- 简单的表单数据
- 只在一个组件使用的数据

### Pinia Store 示例

```typescript
import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

export const useUserStore = defineStore('user', () => {
  // State
  const userInfo = ref<User | null>(null);
  const token = ref('');

  // Getters
  const isLogin = computed(() => !!token.value);

  // Actions
  async function login(username: string, password: string) {
    const res = await loginApi({ username, password });
    token.value = res.token;
    userInfo.value = res.userInfo;
  }

  function logout() {
    token.value = '';
    userInfo.value = null;
  }

  return {
    userInfo,
    token,
    isLogin,
    login,
    logout,
  };
}, {
  persist: true,  // 持久化
});
```

---

## 7. 样式规范

### CSS 类名

```vue
<template>
  <!-- ✅ 推荐：使用 Tailwind CSS -->
  <div class="flex items-center justify-between p-4 mb-2">
    <span class="text-lg font-bold">标题</span>
    <Button class="ml-auto">操作</Button>
  </div>

  <!-- ✅ 推荐：自定义类使用 BEM 命名 -->
  <div class="user-card">
    <div class="user-card__header">
      <span class="user-card__title">标题</span>
    </div>
    <div class="user-card__body">内容</div>
  </div>
</template>

<style scoped>
.user-card {
  border: 1px solid #ddd;
}

.user-card__header {
  padding: 10px;
  border-bottom: 1px solid #ddd;
}

.user-card__title {
  font-size: 16px;
  font-weight: bold;
}
</style>
```

### Scoped 样式

```vue
<style scoped>
/* ✅ 推荐：使用 scoped */
.custom-button {
  background: blue;
}

/* ✅ 推荐：深度选择器 */
:deep(.ant-btn) {
  color: red;
}

/* ❌ 不推荐：全局样式 */
</style>

<style>
/* 全局样式会影响其他组件 */
.ant-btn {
  color: red;
}
</style>
```

---

## 8. 错误处理

### 全局错误处理

```typescript
// main.ts
app.config.errorHandler = (err, instance, info) => {
  console.error('Global error:', err, info);
  message.error('系统错误，请稍后重试');
};
```

### 组件错误边界

```vue
<script setup>
import { onErrorCaptured } from 'vue';

onErrorCaptured((err, instance, info) => {
  console.error('Component error:', err, info);
  // 返回 false 阻止错误继续传播
  return false;
});
</script>
```

### API 错误处理

```typescript
// ✅ 推荐：统一的错误提示
async function handleAction() {
  const hide = message.loading('处理中...', 0);
  try {
    await someApi();
    message.success('操作成功');
  } catch (error: any) {
    message.error(error.message || '操作失败');
  } finally {
    hide();
  }
}

// ✅ 推荐：不同错误码的处理
catch (error: any) {
  switch (error.code) {
    case 401:
      message.error('请先登录');
      router.push('/login');
      break;
    case 403:
      message.error('无权限访问');
      break;
    default:
      message.error(error.message || '操作失败');
  }
}
```

---

## 9. 测试

### 组件测试

```typescript
import { mount } from '@vue/test-utils';
import UserList from './index.vue';

describe('UserList', () => {
  it('renders properly', () => {
    const wrapper = mount(UserList);
    expect(wrapper.text()).toContain('用户列表');
  });

  it('calls refresh when button clicked', async () => {
    const wrapper = mount(UserList);
    const button = wrapper.find('[data-test="refresh"]');
    await button.trigger('click');
    // 验证刷新逻辑
  });
});
```

### API 测试

```typescript
import { describe, it, expect, vi } from 'vitest';
import { getUserList } from './index';

describe('User API', () => {
  it('fetches user list', async () => {
    const data = await getUserList({ page: 1 });
    expect(data.list).toBeInstanceOf(Array);
  });
});
```

---

## 10. 代码审查清单

### 功能层面

- [ ] 功能是否完整实现
- [ ] 边界情况是否处理
- [ ] 错误是否正确处理
- [ ] 加载状态是否显示
- [ ] 权限控制是否正确

### 代码质量

- [ ] TypeScript 类型是否完整
- [ ] 命名是否语义化
- [ ] 代码是否有注释（复杂逻辑）
- [ ] 是否有重复代码
- [ ] 是否遵循项目规范

### 性能层面

- [ ] 是否有性能问题
- [ ] 是否使用虚拟滚动（大列表）
- [ ] 是否使用防抖节流
- [ ] 组件是否懒加载
- [ ] 图片是否懒加载

### 用户体验

- [ ] 交互是否流畅
- [ ] 错误提示是否友好
- [ ] 加载状态是否明确
- [ ] 是否支持键盘操作
- [ ] 是否支持国际化

---

## 11. 常见问题

### Q: 表单提交后如何刷新列表？

A: 使用事件通知：

```vue
<!-- 列表页 -->
<FormModal @success="onRefresh" />

<!-- 表单组件 -->
<script setup>
const emit = defineEmits(['success']);

async function onSubmit() {
  await saveData();
  emit('success');
}
</script>
```

### Q: 如何处理复杂的表单验证？

A: 使用自定义验证器：

```typescript
{
  field: 'password',
  rules: [
    { required: true, message: '请输入密码' },
    {
      validator: async (rule, value) => {
        if (value.length < 6) {
          return Promise.reject('密码长度不能小于6位');
        }
        // 可以调用 API 验证
        const exists = await checkPassword(value);
        if (exists) {
          return Promise.reject('密码已存在');
        }
        return Promise.resolve();
      },
    },
  ],
}
```

### Q: 如何优化大量数据的渲染？

A: 使用虚拟滚动或分页：

```vue
<!-- 方案1：虚拟滚动 -->
<VirtualList :data="largeList" />

<!-- 方案2：分页 -->
<Grid :pager-config="{ enabled: true, pageSize: 20 }" />
```

### Q: 如何处理文件上传？

A: 使用 `uploadFile` 工具函数：

```typescript
import { uploadFile } from '#/utils/upload-helper';

async function handleUpload(file: File) {
  try {
    const url = await uploadFile(file);
    formApi.setFieldValue('avatar', url);
  } catch {
    message.error('上传失败');
  }
}
```

---

## 12. 开发工具

### VSCode 插件推荐

- Vue Language Features (Volar)
- TypeScript Vue Plugin (Volar)
- ESLint
- Prettier
- Tailwind CSS IntelliSense
- Path Intellisense

### 浏览器插件

- Vue.js devtools
- React Developer Tools

### 调试技巧

```typescript
// 开发环境调试
if (import.meta.env.DEV) {
  console.log('Debug info:', data);
}

// 使用 debugger
function handleClick() {
  debugger;  // 代码会在这里暂停
  // ...
}
```
