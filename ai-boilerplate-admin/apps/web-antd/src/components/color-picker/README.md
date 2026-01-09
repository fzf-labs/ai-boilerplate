# ColorPicker 颜色选择器组件

基于 [vue-color](https://www.npmjs.com/package/vue-color) 库实现的多样式颜色选择器组件。

## 功能特性

- 🎨 **多种选择器样式** - 支持 Chrome、Sketch、紧凑型、滑块、色板等 5 种样式
- 🌈 **透明度支持** - 可选择是否启用透明度功能
- 🌙 **暗色主题** - 内置暗色主题支持
- 📱 **响应式设计** - 适配不同屏幕尺寸
- 🔧 **高度可定制** - 支持自定义颜色面板和样式

## 基本使用

```vue
<template>
  <ColorPicker v-model:value="selectedColor" />
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ColorPicker } from '#/components/color-picker';

const selectedColor = ref('#1890ff');
</script>
```

## 组件属性

| 属性名             | 类型           | 默认值      | 说明                   |
| ------------------ | -------------- | ----------- | ---------------------- |
| `value`            | `string`       | `undefined` | 当前选中的颜色值       |
| `enableAlpha`      | `boolean`      | `true`      | 是否启用透明度选择     |
| `allowClear`       | `boolean`      | `false`     | 是否显示清空按钮       |
| `defaultPicker`    | `PickerType`   | `'chrome'`  | 默认显示的选择器类型   |
| `availablePickers` | `PickerType[]` | 所有类型    | 可用的选择器类型列表   |
| `compactColors`    | `string[]`     | 预设颜色    | 紧凑型选择器的颜色列表 |
| `swatchColors`     | `string[][]`   | 预设颜色组  | 色板选择器的颜色组     |

## 组件事件

| 事件名         | 参数               | 说明             |
| -------------- | ------------------ | ---------------- |
| `update:value` | `(value?: string)` | 颜色值变化时触发 |
| `change`       | `(value?: string)` | 颜色值变化时触发 |

## 选择器类型

- `chrome` - Chrome 浏览器风格的颜色选择器
- `compact` - 紧凑型颜色面板
- `sketch` - Sketch 软件风格的选择器
- `slider` - 滑块式颜色选择器
- `swatches` - 色板选择器

## 使用示例

### 仅显示特定选择器

```vue
<template>
  <ColorPicker
    v-model:value="color"
    :available-pickers="['chrome', 'compact']"
    default-picker="compact"
  />
</template>
```

### 禁用透明度

```vue
<template>
  <ColorPicker v-model:value="color" :enable-alpha="false" />
</template>
```

### 自定义颜色面板

```vue
<template>
  <ColorPicker
    v-model:value="color"
    :compact-colors="customColors"
    default-picker="compact"
  />
</template>

<script setup lang="ts">
const customColors = [
  '#FF5733',
  '#33FF57',
  '#3357FF',
  '#FF33F5',
  '#F5FF33',
  '#33F5FF',
  '#FF8C33',
  '#8CFF33',
];
</script>
```

## 样式定制

组件支持通过 CSS 变量或直接覆盖样式类来自定义外观：

```css
.color-picker-wrapper {
  /* 自定义样式 */
}

.color-display {
  width: 60px;
  height: 40px;
}
```

## 暗色主题

组件会自动适配暗色主题，只需在根元素添加 `dark` 类：

```html
<html class="dark">
  <!-- 应用内容 -->
</html>
```
