# AccountSelect 组件

一个用于选择公众号账号的独立组件，专注于账号筛选功能。

## 功能特性

- 🔄 自动获取账号列表
- 📱 响应式设计
- 🎨 可自定义样式
- 🔧 灵活的配置选项

## 基础用法

```vue
<template>
  <AccountSelect v-model="appId" @change="handleAccountChange" />
</template>

<script setup>
import { ref } from 'vue';
import AccountSelect from '../components/account-select/index.vue';

const appId = ref();

const handleAccountChange = (value) => {
  console.log('选择的账号ID:', value);
};
</script>
```

## Props

| 参数        | 说明         | 类型                  | 默认值           |
| ----------- | ------------ | --------------------- | ---------------- |
| modelValue  | 选中的账号ID | `string \| undefined` | `undefined`      |
| placeholder | 选择框占位符 | `string`              | `'请选择公众号'` |
| width       | 选择框宽度   | `string`              | `'200px'`        |

## Events

| 事件名 | 说明 | 回调参数 |
| --- | --- | --- |
| change | 账号选择变化时触发 | `(value: string \| undefined) => void` |
| update:modelValue | v-model 更新事件 | `(value: string \| undefined) => void` |

## 样式自定义

组件使用了 scoped 样式，主要的 CSS 类名：

- `.gzh-selector-title` - 标题文本
- `.gzh-selector-select` - 选择框容器

注意：组件现在只渲染标题和选择框，需要配合外层容器使用。

## 注意事项

1. 组件会自动获取账号列表，如果没有配置账号会自动跳转到账号管理页面
2. 如果账号列表不为空且没有选中值，会自动选择第一个账号
3. 组件依赖 `#/api/gzh/account` 中的 `getAccountSelector` API
4. 需要在项目中正确配置路由 `MpAccount`

## 依赖

- Vue 3
- Ant Design Vue
- @vben/hooks
- vue-router
