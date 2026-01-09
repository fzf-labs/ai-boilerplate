<!-- eslint-disable unicorn/numeric-separators-style -->
<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import AMapLoader from '@amap/amap-jsapi-loader';

// 类型声明
declare global {
  interface Window {
    AMap: any;
  }
}

type AMapInstance = any;

// 组件属性
const props = defineProps({
  // 地图API Key
  mapKey: {
    type: String,
    required: true,
  },
  // 标记数据数组
  markers: {
    type: Array,
    default: () => [],
    validator: (data) => {
      if (!Array.isArray(data)) return false;
      if (data.length === 0) return true;
      // 支持格式：{ id?, position: [lng, lat], icon?, text?, content?, ... }
      return data.every(
        (item) => typeof item === 'object' && Array.isArray(item.position),
      );
    },
  },
  // 是否手动初始化（Modal环境推荐true）
  manualInit: {
    type: Boolean,
    default: false,
  },
  // 容器样式
  containerStyle: {
    type: Object,
    default: () => ({
      width: '100%',
      height: '600px',
    }),
  },
  // 地图配置
  mapConfig: {
    type: Object,
    default: () => ({
      resizeEnable: true,
      zoom: 13,
      center: [116.397428, 39.90923],
    }),
  },
  // 样式配置
  styleConfig: {
    type: Object,
    default: () => ({
      defaultIcon:
        '//a.amap.com/jsapi_demos/static/demo-center/icons/poi-marker-default.png',
      markerOffset: [-13, -30],
    }),
  },
});

// Emits 定义
const emit = defineEmits([
  'mapReady',
  'markerClick',
  'markerAdd',
  'markerRemove',
]);

// 响应式数据
const isMapReady = ref(false);
const markerInstances = ref(new Map());

// 地图相关变量
let map: AMapInstance = null;
let AMap: any = null;

// 生成唯一容器ID
const containerId = `marker-map-${Math.random().toString(36).slice(2, 9)}`;

// 计算属性
const totalMarkers = computed(() => props.markers.length);

// 添加标记
const addMarker = (markerConfig: any) => {
  if (!map || !AMap) return null;

  const {
    id,
    position,
    icon,
    content,
    text,
    offset,
    clickable = true,
    ...otherOptions
  } = markerConfig;

  // 安全创建偏移量
  const safeOffset =
    offset || new AMap.Pixel(...props.styleConfig.markerOffset);

  let marker;

  if (content || text) {
    // 自定义内容标记
    const markerContent = document.createElement('div');

    if (icon) {
      const markerImg = document.createElement('img');
      markerImg.src = icon;
      markerImg.setAttribute('width', '25px');
      markerImg.setAttribute('height', '34px');
      markerImg.className = 'marker-icon';
      markerContent.append(markerImg);
    }

    if (text) {
      const markerSpan = document.createElement('span');
      markerSpan.className = 'marker-text';
      markerSpan.innerHTML = text;
      markerContent.append(markerSpan);
    }

    if (content) {
      if (typeof content === 'string') {
        markerContent.innerHTML = content;
      } else if (content instanceof HTMLElement) {
        markerContent.append(content);
      }
    }

    marker = new AMap.Marker({
      position,
      content: markerContent,
      offset: safeOffset,
      ...otherOptions,
    });
  } else {
    // 默认图标标记
    marker = new AMap.Marker({
      position,
      icon: icon || props.styleConfig.defaultIcon,
      offset: safeOffset,
      ...otherOptions,
    });
  }

  // 添加点击事件
  if (clickable) {
    marker.on('click', (e: any) => {
      emit('markerClick', { marker, config: markerConfig, event: e });
    });
  }

  marker.setMap(map);

  if (id) {
    markerInstances.value.set(id, { marker, config: markerConfig });
  }

  emit('markerAdd', { marker, config: markerConfig });
  return marker;
};

// 清除所有标记
const clearAllMarkers = () => {
  markerInstances.value.forEach((markerData, id) => {
    markerData.marker.setMap(null);
    emit('markerRemove', { id, config: markerData.config });
  });
  markerInstances.value.clear();
};

// 初始化地图
const initMap = async () => {
  if (!props.mapKey || map) return;

  const container = document.querySelector(`#${containerId}`);
  if (!container) {
    console.error('标记地图容器未找到:', containerId);
    return;
  }

  // 检查容器尺寸
  const rect = container.getBoundingClientRect();
  if (rect.width === 0 || rect.height === 0) {
    console.warn('标记地图容器尺寸为0，延迟初始化');
    setTimeout(initMap, 200);
    return;
  }

  // 检查容器是否在可见区域内
  if (!(container as HTMLElement).offsetParent) {
    console.warn('标记地图容器不可见，延迟初始化');
    setTimeout(initMap, 200);
    return;
  }

  try {
    AMap = await AMapLoader.load({
      key: props.mapKey,
      version: '2.0',
      plugins: [],
    });

    // 获取地图中心点
    const center =
      props.markers.length > 0
        ? (props.markers[0] as any).position
        : props.mapConfig.center;

    // 创建地图
    map = new AMap.Map(containerId, {
      ...props.mapConfig,
      center,
    });

    // 添加地图加载错误处理
    map.on('error', (error: any) => {
      console.error('标记地图加载错误:', error);
    });

    // 监听地图加载完成事件
    map.on('complete', () => {
      try {
        isMapReady.value = true;
        emit('mapReady', { map, AMap });

        // 添加标记
        nextTick(() => {
          props.markers.forEach((markerConfig) => {
            try {
              addMarker(markerConfig);
            } catch (error) {
              console.error('添加标记失败:', error);
            }
          });

          // 设置地图显示范围以包含所有标记
          if (props.markers.length > 0 && markerInstances.value.size > 0) {
            const allMarkers = [...markerInstances.value.values()].map(
              (item) => item.marker,
            );
            map.setFitView(allMarkers);
          }
        });
      } catch (error) {
        console.error('地图完成事件处理失败:', error);
      }
    });

    console.warn('标记地图初始化成功');
  } catch (error) {
    console.error('标记地图初始化失败:', error);
  }
};

// 监听标记数据变化
watch(
  () => props.markers,
  () => {
    if (isMapReady.value && map) {
      try {
        // 清除旧标记
        clearAllMarkers();

        // 添加新标记
        props.markers.forEach((markerConfig) => {
          try {
            addMarker(markerConfig);
          } catch (error) {
            console.error('添加标记失败:', error);
          }
        });

        // 重新设置地图显示范围
        if (props.markers.length > 0 && markerInstances.value.size > 0) {
          const allMarkers = [...markerInstances.value.values()].map(
            (item) => item.marker,
          );
          map.setFitView(allMarkers);
        }
      } catch (error) {
        console.error('标记数据变化处理失败:', error);
      }
    }
  },
  { deep: true },
);

// 强制初始化地图（用于调试和特殊情况）
const forceInitMap = () => {
  console.warn('强制初始化标记地图...');
  map = null; // 重置地图实例
  isMapReady.value = false;
  initMap();
};

// 暴露方法给父组件
defineExpose({
  initMap,
  forceInitMap,
  clearAllMarkers,
  isMapReady,
  totalMarkers,
  getMap: () => map,
  getAMap: () => AMap,
});

onMounted(() => {
  if (!props.manualInit) {
    setTimeout(() => {
      initMap();
    }, 200);
  }
});

onUnmounted(() => {
  clearAllMarkers();
  map?.destroy();
});
</script>

<template>
  <div class="marker-map-container">
    <div :id="containerId" class="map-container" :style="containerStyle"></div>

    <!-- 无数据提示 -->
    <div v-if="totalMarkers === 0" class="no-data-overlay">
      <div class="no-data-content">
        <div class="no-data-icon">📍</div>
        <div class="no-data-text">暂无标记数据</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.marker-map-container {
  position: relative;
  width: 100%;
  height: 600px;
}

.map-container {
  width: 100%;
  height: 100%;
  overflow: hidden;
  border-radius: 8px;
}

/* 标记样式 */
:deep(.marker-icon) {
  width: 25px;
  height: 34px;
}

:deep(.marker-text) {
  position: absolute;
  top: -20px;
  right: -118px;
  padding: 4px 10px;
  font-size: 12px;
  color: #fff;
  white-space: nowrap;
  background-color: #25a5f7;
  border-radius: 3px;
  box-shadow: 1px 1px 1px rgb(10 10 10 / 20%);
}

/* 无数据提示样式 */
.no-data-overlay {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgb(0 0 0 / 5%);
  border-radius: 8px;
  backdrop-filter: blur(2px);
}

.no-data-content {
  color: #999;
  text-align: center;
}

.no-data-icon {
  margin-bottom: 12px;
  font-size: 48px;
}

.no-data-text {
  font-size: 16px;
  font-weight: 500;
}
</style>
