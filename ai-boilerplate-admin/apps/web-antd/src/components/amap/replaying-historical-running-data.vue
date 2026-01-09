<!-- eslint-disable unicorn/numeric-separators-style -->
<script setup lang="ts">
import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import AMapLoader from '@amap/amap-jsapi-loader';

// 组件属性
const props = defineProps({
  // 地图API Key
  mapKey: {
    type: String,
    required: true,
  },
  // 轨迹数据（支持带时间戳的格式）
  trackData: {
    type: Array,
    default: () => [],
    validator: (data) => {
      if (!Array.isArray(data)) return false;
      if (data.length === 0) return true;
      // 支持两种格式：[lng, lat] 或 { position: [lng, lat], timestamp: '时间' }
      return data.every(
        (item) =>
          (Array.isArray(item) && item.length >= 2) ||
          (typeof item === 'object' && Array.isArray(item.position)),
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
      zoom: 16,
    }),
  },
  // 动画配置
  animationConfig: {
    type: Object,
    default: () => ({
      duration: 5000, // 总动画时长（毫秒）- 增加基础时长
      speed: 1, // 播放速度倍率
      autoRotation: true, // 自动旋转
      showDirection: true, // 显示方向
    }),
  },
  // 样式配置
  styleConfig: {
    type: Object,
    default: () => ({
      markerIcon:
        'https://a.amap.com/jsapi_demos/static/demo-center-v2/car.png',
      trackColor: '#1890ff',
      passedColor: '#52c41a',
      strokeWeight: 6,
    }),
  },
  // 是否显示控制面板
  showControls: {
    type: Boolean,
    default: true,
  },
});

// Emits 定义
const emit = defineEmits([
  'mapReady',
  'animationStart',
  'animationPause',
  'animationResume',
  'animationStop',
  'animationEnd',
  'positionChange',
]);

// 响应式数据
const isPlaying = ref(false);
const isPaused = ref(false);
const currentPosition = ref(0);
const playbackSpeed = ref(1);
const isMapReady = ref(false);
const animationStartTime = ref(0);
const expectedAnimationDuration = ref(0);

// 地图相关变量
let map: any = null;
let marker: any = null;
let passedPolyline: any = null;
let fullPolyline: any = null;
let AMap: any = null;

// 生成唯一容器ID
const containerId = `track-replay-${Math.random().toString(36).slice(2, 9)}`;

// 计算属性
const processedTrackData = computed(() => {
  if (props.trackData.length === 0) return [];

  return props.trackData.map((item: any) => {
    if (Array.isArray(item)) {
      return { position: item, timestamp: null };
    }
    return {
      position: item.position || item,
      timestamp: item.timestamp || item.createdAt || null,
    };
  });
});

const trackPositions = computed(() =>
  processedTrackData.value.map((item) => item.position),
);

const totalPoints = computed(() => trackPositions.value.length);

const currentSpeedText = computed(() => {
  const speedMap: Record<number, string> = {
    0.5: '0.5x',
    1: '1x',
    2: '2x',
    4: '4x',
    8: '8x',
  };
  return speedMap[playbackSpeed.value] || `${playbackSpeed.value}x`;
});

const progressPercentage = computed(() => {
  if (totalPoints.value === 0) return 0;
  return Math.round((currentPosition.value / (totalPoints.value - 1)) * 100);
});

// 动画控制方法
const startAnimation = () => {
  if (!marker || !trackPositions.value || trackPositions.value.length === 0)
    return;

  try {
    // 修改速度计算：使用乘法而非除法，让高速度档位更快
    const baseDuration = props.animationConfig.duration;
    const speedMultiplier =
      playbackSpeed.value >= 1
        ? playbackSpeed.value * 2 // 1x以上档位：2倍、4倍、8倍
        : playbackSpeed.value; // 1x以下档位保持原有计算

    const duration = Math.floor(baseDuration / speedMultiplier);

    if (typeof marker.moveAlong === 'function') {
      marker.moveAlong(trackPositions.value, {
        duration,
        autoRotation: props.animationConfig.autoRotation,
      });

      isPlaying.value = true;
      isPaused.value = false;
      animationStartTime.value = Date.now();
      expectedAnimationDuration.value = duration;

      emit('animationStart', {
        trackData: processedTrackData.value,
        duration,
        speed: playbackSpeed.value,
      });
    }
  } catch (error) {
    console.error('启动动画失败:', error);
  }
};

const pauseAnimation = () => {
  if (marker && typeof marker.pauseMove === 'function') {
    try {
      marker.pauseMove();
      isPaused.value = true;
      emit('animationPause');
    } catch (error) {
      console.error('暂停动画失败:', error);
    }
  }
};

const resumeAnimation = () => {
  if (marker && typeof marker.resumeMove === 'function') {
    try {
      marker.resumeMove();
      isPaused.value = false;
      emit('animationResume');
    } catch (error) {
      console.error('恢复动画失败:', error);
    }
  }
};

const stopAnimation = () => {
  if (marker) {
    try {
      if (typeof marker.stopMove === 'function') {
        marker.stopMove();
      }
      isPlaying.value = false;
      isPaused.value = false;
      currentPosition.value = 0;
      animationStartTime.value = 0;
      expectedAnimationDuration.value = 0;

      // 重置标记点位置到起点
      if (
        trackPositions.value.length > 0 &&
        typeof marker.setPosition === 'function'
      ) {
        marker.setPosition(trackPositions.value[0]);
      }

      // 清空已走过的轨迹
      if (passedPolyline && typeof passedPolyline.setPath === 'function') {
        passedPolyline.setPath([]);
      }

      emit('animationStop');
    } catch (error) {
      console.error('停止动画失败:', error);
    }
  }
};

// 速度控制
const changeSpeed = (speed: number) => {
  playbackSpeed.value = speed;

  // 如果正在播放，重新开始以应用新速度
  if (isPlaying.value && !isPaused.value) {
    const wasPlaying = true;
    stopAnimation();
    if (wasPlaying) {
      nextTick(() => {
        startAnimation();
      });
    }
  }
};

// 跳转到指定位置
const seekToPosition = (position: number) => {
  if (!marker || trackPositions.value.length === 0) return;

  const targetIndex = Math.max(
    0,
    Math.min(position, trackPositions.value.length - 1),
  );
  currentPosition.value = targetIndex;

  marker.setPosition(trackPositions.value[targetIndex]);

  // 更新已走过的轨迹
  if (passedPolyline) {
    const passedPath = trackPositions.value.slice(0, targetIndex + 1);
    passedPolyline.setPath(passedPath);
  }

  emit('positionChange', {
    position: targetIndex,
    coordinate: trackPositions.value[targetIndex],
    data: processedTrackData.value[targetIndex],
  });
};

// 初始化地图
const initMap = async () => {
  if (!props.mapKey || map) return;

  const container = document.querySelector(`#${containerId}`);
  if (!container) {
    console.error('轨迹回放容器未找到:', containerId);
    return;
  }

  // 检查容器尺寸
  const rect = container.getBoundingClientRect();
  if (rect.width === 0 || rect.height === 0) {
    console.warn('轨迹回放容器尺寸为0，延迟初始化');
    setTimeout(initMap, 200);
    return;
  }

  // 检查容器是否在可见区域内
  if (!(container as HTMLElement).offsetParent) {
    console.warn('轨迹回放容器不可见，延迟初始化');
    setTimeout(initMap, 200);
    return;
  }

  try {
    AMap = await AMapLoader.load({
      key: props.mapKey,
      version: '2.0',
      plugins: ['AMap.MoveAnimation'], // 加载动画插件
    });

    // 获取地图中心点
    const center =
      trackPositions.value.length > 0
        ? trackPositions.value[0]
        : [116.397428, 39.90923];

    // 创建地图
    map = new AMap.Map(containerId, {
      ...props.mapConfig,
      center,
    });

    // 添加地图加载错误处理
    map.on('error', (error: any) => {
      console.error('轨迹回放地图加载错误:', error);
    });

    // 创建标记点
    marker = new AMap.Marker({
      map,
      position: center,
      icon: props.styleConfig.markerIcon,
      offset: AMap ? new AMap.Pixel(-13, -26) : undefined,
    });

    // 绘制完整轨迹线
    if (trackPositions.value.length > 1) {
      fullPolyline = new AMap.Polyline({
        map,
        path: trackPositions.value,
        showDir: props.animationConfig.showDirection,
        strokeColor: props.styleConfig.trackColor,
        strokeWeight: props.styleConfig.strokeWeight,
        strokeOpacity: 0.6,
        strokeStyle: 'dashed',
      });
    }

    // 创建已走过的轨迹
    passedPolyline = new AMap.Polyline({
      map,
      strokeColor: props.styleConfig.passedColor,
      strokeWeight: props.styleConfig.strokeWeight,
      strokeOpacity: 0.9,
    });

    // 监听标记点移动事件
    marker.on('moving', (e: any) => {
      try {
        const passedPath = e.passedPath || [];
        if (passedPolyline && typeof passedPolyline.setPath === 'function') {
          passedPolyline.setPath(passedPath);
        }

        // 修复位置跟踪逻辑：根据已走过路径的长度来计算位置
        const pathLength = passedPath.length;
        if (pathLength > 0) {
          // 确保位置不超出轨迹点数量
          const newPosition = Math.min(
            pathLength - 1,
            trackPositions.value.length - 1,
          );
          currentPosition.value = newPosition;

          // 发送位置变化事件
          emit('positionChange', {
            position: currentPosition.value,
            coordinate: passedPath[pathLength - 1],
            data: processedTrackData.value[currentPosition.value],
          });
        }

        // 地图跟随
        if (
          map &&
          typeof map.setCenter === 'function' &&
          e.target &&
          typeof e.target.getPosition === 'function'
        ) {
          map.setCenter(e.target.getPosition(), true);
        }
      } catch (error) {
        console.error('标记点移动事件处理失败:', error);
      }
    });

    // 监听动画结束事件
    marker.on('moveend', () => {
      // 多重条件判断确保是真正的动画结束
      const isAtLastPosition =
        currentPosition.value >= trackPositions.value.length - 1;
      const hasAnimationStarted = animationStartTime.value > 0;
      const timeElapsed = Date.now() - animationStartTime.value;
      const isTimeComplete =
        timeElapsed >= expectedAnimationDuration.value * 0.8; // 允许20%误差

      if (
        isPlaying.value &&
        isAtLastPosition &&
        hasAnimationStarted &&
        isTimeComplete
      ) {
        isPlaying.value = false;
        isPaused.value = false;
        currentPosition.value = trackPositions.value.length - 1;
        animationStartTime.value = 0;
        expectedAnimationDuration.value = 0;

        emit('animationEnd', {
          finalPosition: currentPosition.value,
          totalPoints: totalPoints.value,
        });
      }
    });

    // 设置地图显示范围
    if (trackPositions.value.length > 0) {
      map.setFitView(
        trackPositions.value.length > 1 ? [fullPolyline] : [marker],
      );
    }

    isMapReady.value = true;
    emit('mapReady', { map, AMap, marker });
  } catch (error: any) {
    console.error('轨迹回放地图初始化失败:', error);
    if (error.message && error.message.includes('transform')) {
      console.error('TrackReplay transform 错误详情:', {
        error: error.message,
        stack: error.stack,
        mapInstance: !!map,
        markerInstance: !!marker,
        AMapInstance: !!AMap,
        trackPositions: trackPositions.value,
        timestamp: new Date().toISOString(),
      });
    }
  }
};

// 监听轨迹数据变化
watch(
  () => props.trackData,
  () => {
    if (isMapReady.value && map) {
      try {
        // 重新绘制轨迹
        if (fullPolyline && typeof fullPolyline.setPath === 'function') {
          fullPolyline.setPath(trackPositions.value);
        }
        if (passedPolyline && typeof passedPolyline.setPath === 'function') {
          passedPolyline.setPath([]);
        }
        if (
          marker &&
          typeof marker.setPosition === 'function' &&
          trackPositions.value.length > 0
        ) {
          marker.setPosition(trackPositions.value[0]);
        }
        currentPosition.value = 0;

        // 重新设置显示范围
        if (
          trackPositions.value.length > 0 &&
          typeof map.setFitView === 'function'
        ) {
          const fitViewObjects = [];
          if (trackPositions.value.length > 1 && fullPolyline) {
            fitViewObjects.push(fullPolyline);
          } else if (marker) {
            fitViewObjects.push(marker);
          }
          if (fitViewObjects.length > 0) {
            map.setFitView(fitViewObjects);
          }
        }
      } catch (error) {
        console.error('轨迹数据变化处理失败:', error);
      }
    }
  },
  { deep: true },
);

// 强制初始化地图（用于调试和特殊情况）
const forceInitMap = () => {
  console.warn('强制初始化轨迹回放地图...');
  map = null; // 重置地图实例
  isMapReady.value = false;
  initMap();
};

// 暴露方法给父组件
defineExpose({
  initMap,
  forceInitMap,
  startAnimation,
  pauseAnimation,
  resumeAnimation,
  stopAnimation,
  changeSpeed,
  seekToPosition,
  isPlaying,
  isPaused,
  isMapReady,
  currentPosition,
  totalPoints,
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
  if (marker) {
    marker.stopMove();
  }
  map?.destroy();
});

// 状态显示方法
const getStatusText = () => {
  if (!isMapReady.value) return '加载中';
  if (trackPositions.value.length === 0) return '无轨迹数据';
  if (isPlaying.value && !isPaused.value) return '播放中';
  if (isPaused.value) return '已暂停';
  return '就绪';
};

const getStatusClass = () => {
  if (!isMapReady.value) return 'loading';
  if (trackPositions.value.length === 0) return 'no-data';
  if (isPlaying.value && !isPaused.value) return 'playing';
  if (isPaused.value) return 'paused';
  return 'idle';
};
</script>

<template>
  <div class="track-replay-container">
    <div :id="containerId" class="map-container" :style="containerStyle"></div>

    <!-- 控制面板 -->
    <div v-if="showControls" class="control-panel">
      <div class="panel-header">
        <h4>轨迹回放</h4>
        <div class="status-info">
          <span class="status-text" :class="getStatusClass()">
            {{ getStatusText() }}
          </span>
        </div>
      </div>

      <!-- 播放控制 -->
      <div class="control-section">
        <div class="control-buttons">
          <button
            :disabled="
              !isMapReady ||
              trackPositions.length === 0 ||
              (isPlaying && !isPaused)
            "
            class="control-btn start-btn"
            @click="startAnimation"
          >
            <span class="btn-icon">▶</span>
            开始
          </button>
          <button
            :disabled="!isPlaying || isPaused"
            class="control-btn pause-btn"
            @click="pauseAnimation"
          >
            <span class="btn-icon">⏸</span>
            暂停
          </button>
          <button
            :disabled="!isPaused"
            class="control-btn resume-btn"
            @click="resumeAnimation"
          >
            <span class="btn-icon">▶</span>
            继续
          </button>
          <button
            :disabled="!isPlaying && !isPaused"
            class="control-btn stop-btn"
            @click="stopAnimation"
          >
            <span class="btn-icon">⏹</span>
            停止
          </button>
        </div>
      </div>

      <!-- 速度控制 -->
      <div class="control-section">
        <div class="speed-control">
          <label class="control-label">播放速度: {{ currentSpeedText }}</label>
          <div class="speed-buttons">
            <button
              v-for="speed in [0.5, 1, 2, 4, 8]"
              :key="speed"
              class="speed-btn"
              :class="[{ active: playbackSpeed === speed }]"
              @click="changeSpeed(speed)"
            >
              {{ speed }}x
            </button>
          </div>
        </div>
      </div>

      <!-- 进度控制 -->
      <div v-if="totalPoints > 0" class="control-section">
        <div class="progress-control">
          <label class="control-label">
            进度: {{ currentPosition + 1 }} / {{ totalPoints }} ({{
              progressPercentage
            }}%)
          </label>
          <div class="progress-bar">
            <input
              type="range"
              :min="0"
              :max="Math.max(0, totalPoints - 1)"
              :value="currentPosition"
              class="progress-slider"
              @input="
                seekToPosition(
                  Number(($event.target as HTMLInputElement)?.value || 0),
                )
              "
            />
          </div>
        </div>
      </div>

      <!-- 当前位置信息 -->
      <div v-if="processedTrackData[currentPosition]" class="control-section">
        <div class="position-info">
          <div class="info-item">
            <span class="info-label">经度:</span>
            <span class="info-value">{{
              trackPositions[currentPosition]?.[0]?.toFixed(6)
            }}</span>
          </div>
          <div class="info-item">
            <span class="info-label">纬度:</span>
            <span class="info-value">{{
              trackPositions[currentPosition]?.[1]?.toFixed(6)
            }}</span>
          </div>
          <div
            v-if="processedTrackData[currentPosition]?.timestamp"
            class="info-item"
          >
            <span class="info-label">时间:</span>
            <span class="info-value">{{
              processedTrackData[currentPosition]?.timestamp
            }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 无数据提示 -->
    <div v-if="trackPositions.length === 0" class="no-data-overlay">
      <div class="no-data-content">
        <div class="no-data-icon">📍</div>
        <div class="no-data-text">暂无轨迹数据</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.track-replay-container {
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

.control-panel {
  position: absolute;
  top: 15px;
  right: 15px;
  min-width: 280px;
  max-width: 320px;
  padding: 16px;
  background: rgb(255 255 255 / 96%);
  border: 1px solid rgb(255 255 255 / 20%);
  border-radius: 8px;
  box-shadow: 0 4px 24px rgb(0 0 0 / 15%);
  backdrop-filter: blur(8px);
}

.panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 12px;
  margin-bottom: 16px;
  border-bottom: 1px solid #e8e8e8;
}

.panel-header h4 {
  margin: 0;
  font-size: 15px;
  font-weight: 600;
  color: #333;
}

.status-info {
  display: flex;
  align-items: center;
}

.status-text {
  padding: 4px 8px;
  font-size: 12px;
  font-weight: 600;
  background: #f0f0f0;
  border-radius: 12px;
}

.status-text.loading {
  color: #1890ff;
  background: #e6f7ff;
}

.status-text.no-data {
  color: #999;
  background: #f5f5f5;
}

.status-text.idle {
  color: #666;
  background: #f0f0f0;
}

.status-text.playing {
  color: #52c41a;
  background: #f6ffed;
}

.status-text.paused {
  color: #faad14;
  background: #fffbe6;
}

.control-section {
  margin-bottom: 16px;
}

.control-section:last-child {
  margin-bottom: 0;
}

.control-buttons {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}

.control-btn {
  display: flex;
  gap: 4px;
  align-items: center;
  justify-content: center;
  min-height: 36px;
  padding: 8px 12px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  border: none;
  border-radius: 6px;
  transition: all 0.2s ease;
}

.control-btn:disabled {
  cursor: not-allowed;
  opacity: 0.4;
}

.btn-icon {
  font-size: 10px;
}

.start-btn {
  color: white;
  background: #52c41a;
}

.start-btn:not(:disabled):hover {
  background: #389e0d;
}

.pause-btn {
  color: white;
  background: #faad14;
}

.pause-btn:not(:disabled):hover {
  background: #d48806;
}

.resume-btn {
  color: white;
  background: #1890ff;
}

.resume-btn:not(:disabled):hover {
  background: #096dd9;
}

.stop-btn {
  color: white;
  background: #ff4d4f;
}

.stop-btn:not(:disabled):hover {
  background: #cf1322;
}

.control-label {
  display: block;
  margin-bottom: 8px;
  font-size: 12px;
  font-weight: 500;
  color: #666;
}

.speed-control,
.progress-control {
  width: 100%;
}

.speed-buttons {
  display: flex;
  gap: 4px;
}

.speed-btn {
  flex: 1;
  padding: 6px;
  font-size: 11px;
  cursor: pointer;
  background: white;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.speed-btn:hover {
  color: #1890ff;
  border-color: #1890ff;
}

.speed-btn.active {
  color: white;
  background: #1890ff;
  border-color: #1890ff;
}

.progress-bar {
  margin-top: 8px;
}

.progress-slider {
  width: 100%;
  height: 6px;
  appearance: none;
  cursor: pointer;
  outline: none;
  background: #f0f0f0;
  border-radius: 3px;
}

.progress-slider::-webkit-slider-thumb {
  width: 16px;
  height: 16px;
  appearance: none;
  cursor: pointer;
  background: #1890ff;
  border: 2px solid white;
  border-radius: 50%;
  box-shadow: 0 2px 6px rgb(0 0 0 / 15%);
}

.progress-slider::-moz-range-thumb {
  width: 16px;
  height: 16px;
  cursor: pointer;
  background: #1890ff;
  border: 2px solid white;
  border-radius: 50%;
  box-shadow: 0 2px 6px rgb(0 0 0 / 15%);
}

.position-info {
  padding: 12px;
  background: #fafafa;
  border-radius: 6px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 6px;
}

.info-item:last-child {
  margin-bottom: 0;
}

.info-label {
  font-size: 11px;
  font-weight: 500;
  color: #666;
}

.info-value {
  font-family: monospace;
  font-size: 11px;
  color: #333;
}

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
