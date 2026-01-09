<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import {
  ChromePicker,
  CompactPicker,
  SketchPicker,
  SliderPicker,
  SwatchesPicker,
  tinycolor,
} from 'vue-color';
import 'vue-color/style.css';

interface ColorPickerProps {
  /** 当前颜色值 */
  value?: string;
  /** 是否启用透明度 */
  enableAlpha?: boolean;
  /** 是否允许清空 */
  allowClear?: boolean;
  /** 默认选择器类型 */
  defaultPicker?: 'chrome' | 'compact' | 'sketch' | 'slider' | 'swatches';
  /** 可用的选择器类型 */
  availablePickers?: Array<
    'chrome' | 'compact' | 'sketch' | 'slider' | 'swatches'
  >;
  /** 紧凑型选择器的颜色 */
  compactColors?: string[];
  /** 色板选择器的颜色组 */
  swatchColors?: string[][];
}

interface ColorPickerEmits {
  (e: 'update:value', value?: string): void;
  (e: 'change', value?: string): void;
}

const props = withDefaults(defineProps<ColorPickerProps>(), {
  value: undefined,
  enableAlpha: true,
  allowClear: false,
  defaultPicker: 'chrome',
  availablePickers: () => ['chrome', 'compact', 'sketch', 'slider', 'swatches'],
  compactColors: () => [
    '#FFFFFF',
    '#F2F2F2',
    '#E6E6E6',
    '#CCCCCC',
    '#B3B3B3',
    '#999999',
    '#808080',
    '#666666',
    '#4D4D4D',
    '#333333',
    '#1A1A1A',
    '#000000',
    '#FF0000',
    '#FF8000',
    '#FFFF00',
    '#80FF00',
    '#00FF00',
    '#00FF80',
    '#00FFFF',
    '#0080FF',
    '#0000FF',
    '#8000FF',
    '#FF00FF',
    '#FF0080',
  ],
  swatchColors: () => [
    ['#FF6B6B', '#4ECDC4', '#45B7D1', '#96CEB4', '#FFEAA7'],
    ['#DDA0DD', '#98D8C8', '#F7DC6F', '#BB8FCE', '#85C1E9'],
    ['#F8C471', '#82E0AA', '#AED6F1', '#F1948A', '#D7BDE2'],
  ],
});

const emit = defineEmits<ColorPickerEmits>();

// 响应式状态
const showPicker = ref(false);
const currentPicker = ref(props.defaultPicker);
const color = ref(props.value || '#1890ff');
const originalColor = ref(props.value);
const dropdownPosition = ref({ top: 0, left: 0 });
const colorDisplayRef = ref<HTMLElement>();

// 计算属性
const displayColor = computed(() => {
  return props.value || 'transparent';
});

const pickers = computed(() => {
  const allPickers = [
    { key: 'chrome', label: '基础' },
    { key: 'compact', label: '紧凑' },
    { key: 'sketch', label: '手绘' },
    { key: 'slider', label: '滑块' },
    { key: 'swatches', label: '色板' },
  ];

  return allPickers.filter((picker) =>
    props.availablePickers.includes(picker.key as any),
  );
});

// 监听外部值变化
watch(
  () => props.value,
  (newValue) => {
    if (newValue !== color.value) {
      color.value = newValue || '#1890ff';
    }
  },
);

// 方法
const togglePicker = () => {
  if (!showPicker.value) {
    originalColor.value = props.value;
    color.value = props.value || '#1890ff';

    // 计算弹出层位置
    if (colorDisplayRef.value) {
      const rect = colorDisplayRef.value.getBoundingClientRect();
      const dropdownWidth = 280; // 弹出层最小宽度
      const dropdownHeight = 300; // 估算弹出层高度

      let top = rect.bottom + 4;
      let left = rect.left;

      // 确保不超出右边界
      if (left + dropdownWidth > window.innerWidth) {
        left = window.innerWidth - dropdownWidth - 10;
      }

      // 确保不超出左边界
      if (left < 10) {
        left = 10;
      }

      // 如果下方空间不够，显示在上方
      if (top + dropdownHeight > window.innerHeight) {
        top = rect.top - dropdownHeight - 4;
      }

      // 确保不超出顶部边界
      if (top < 10) {
        top = 10;
      }
      dropdownPosition.value = { top, left };
    }
  }
  showPicker.value = !showPicker.value;
};

const handleColorChange = (newColor: any) => {
  if (typeof newColor === 'string') {
    color.value = newColor;
  } else if (newColor && typeof newColor === 'object') {
    // 处理tinycolor对象或其他颜色对象
    const colorInstance = tinycolor(newColor);
    color.value = props.enableAlpha
      ? colorInstance.toRgbString()
      : colorInstance.toHexString();
  }
};

const confirmSelection = () => {
  emit('update:value', color.value);
  emit('change', color.value);
  showPicker.value = false;
};

const cancelSelection = () => {
  color.value = originalColor.value || '#1890ff';
  showPicker.value = false;
};

const clearColor = () => {
  color.value = '';
  emit('update:value', undefined);
  emit('change', undefined);
  showPicker.value = false;
};
</script>

<template>
  <div class="color-picker-wrapper">
    <!-- 颜色展示框 -->
    <div
      ref="colorDisplayRef"
      class="color-display"
      :style="{ backgroundColor: displayColor }"
      @click="togglePicker"
    >
      <span v-if="!value" class="placeholder">选择颜色</span>
    </div>

    <!-- 颜色选择器弹出层 -->
    <Teleport to="body">
      <div
        v-if="showPicker"
        class="color-picker-dropdown"
        :style="{
          top: `${dropdownPosition.top}px`,
          left: `${dropdownPosition.left}px`,
        }"
      >
        <div class="picker-tabs">
          <button
            v-for="picker in pickers"
            :key="picker.key"
            class="tab-button"
            :class="{ active: currentPicker === picker.key }"
            @click="currentPicker = picker.key as any"
          >
            {{ picker.label }}
          </button>
        </div>

        <div class="picker-container">
          <!-- Chrome风格选择器 -->
          <ChromePicker
            v-if="currentPicker === 'chrome'"
            v-model="color"
            :disable-alpha="!enableAlpha"
            @update:model-value="handleColorChange"
          />

          <!-- 紧凑型选择器 -->
          <CompactPicker
            v-else-if="currentPicker === 'compact'"
            v-model="color"
            :colors="compactColors"
            @update:model-value="handleColorChange"
          />

          <!-- Sketch风格选择器 -->
          <SketchPicker
            v-else-if="currentPicker === 'sketch'"
            v-model="color"
            :disable-alpha="!enableAlpha"
            @update:model-value="handleColorChange"
          />

          <!-- 滑块选择器 -->
          <SliderPicker
            v-else-if="currentPicker === 'slider'"
            v-model="color"
            @update:model-value="handleColorChange"
          />

          <!-- 色板选择器 -->
          <SwatchesPicker
            v-else-if="currentPicker === 'swatches'"
            v-model="color"
            :colors="swatchColors"
            @update:model-value="handleColorChange"
          />
        </div>

        <div class="picker-actions">
          <button class="btn btn-cancel" @click="cancelSelection">取消</button>
          <button class="btn btn-confirm" @click="confirmSelection">
            确定
          </button>
          <button v-if="allowClear" class="btn btn-clear" @click="clearColor">
            清空
          </button>
        </div>
      </div>
    </Teleport>

    <!-- 遮罩层 -->
    <Teleport to="body">
      <div v-if="showPicker" class="picker-overlay" @click="togglePicker"></div>
    </Teleport>
  </div>
</template>

<style scoped>
.color-picker-wrapper {
  position: relative;
  display: inline-block;
}

.color-display {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 80px;
  height: 32px;
  padding: 0 8px;
  cursor: pointer;
  background-color: #f5f5f5;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  transition: border-color 0.3s;
}

/* 当有颜色值时，隐藏默认背景 */
.color-display[style*='background-color']:not([style*='transparent']) {
  background-image: none;
}

.color-display:hover {
  border-color: #1890ff;
}

.placeholder {
  font-size: 12px;
  color: #666;
  white-space: nowrap;
}

.picker-overlay {
  position: fixed;
  inset: 0;
  z-index: 9999;
}

.color-picker-dropdown {
  position: fixed;
  z-index: 10000;
  min-width: 280px;
  background: white;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgb(0 0 0 / 15%);
}

.picker-tabs {
  display: flex;
  padding: 8px 8px 0;
  border-bottom: 1px solid #f0f0f0;
}

.tab-button {
  flex: 1;
  padding: 8px 12px;
  font-size: 12px;
  cursor: pointer;
  background: transparent;
  border: none;
  border-radius: 4px 4px 0 0;
  transition: all 0.3s;
}

.tab-button:hover {
  background: #f5f5f5;
}

.tab-button.active {
  color: white;
  background: #1890ff;
}

.picker-container {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
}

.picker-actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
  padding: 12px 16px;
  border-top: 1px solid #f0f0f0;
}

.btn {
  padding: 4px 12px;
  font-size: 12px;
  cursor: pointer;
  background: white;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  transition: all 0.3s;
}

.btn:hover {
  color: #1890ff;
  border-color: #1890ff;
}

.btn-confirm {
  color: white;
  background: #1890ff;
  border-color: #1890ff;
}

.btn-confirm:hover {
  background: #40a9ff;
  border-color: #40a9ff;
}

.btn-clear {
  color: #ff4d4f;
  border-color: #ff4d4f;
}

.btn-clear:hover {
  color: white;
  background: #ff4d4f;
}

/* 适配暗色主题 */
:global(.dark) .color-picker-dropdown {
  color: white;
  background: #1f1f1f;
  border-color: #434343;
}

:global(.dark) .picker-tabs {
  border-bottom-color: #434343;
}

:global(.dark) .tab-button:hover {
  background: #2f2f2f;
}

:global(.dark) .picker-actions {
  border-top-color: #434343;
}

:global(.dark) .btn {
  color: white;
  background: #1f1f1f;
  border-color: #434343;
}

:global(.dark) .btn:hover {
  color: #1890ff;
  border-color: #1890ff;
}
</style>
