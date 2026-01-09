<script lang="ts" setup>
import { onMounted, ref } from 'vue';

import {
  CloudOutlined,
  EnvironmentOutlined,
  EyeOutlined,
  LoadingOutlined,
} from '@ant-design/icons-vue';
import { Card, Spin } from 'ant-design-vue';

// 高德地图IP定位响应类型
interface AmapIpResponse {
  status: string;
  info: string;
  infocode: string;
  province: string;
  city: string;
  adcode: string;
  rectangle: string;
}

// 高德地图天气预报单日数据
interface AmapWeatherCast {
  /** 日期 */
  date: string;
  /** 星期几 */
  week: string;
  /** 白天天气现象 */
  dayweather: string;
  /** 夜间天气现象 */
  nightweather: string;
  /** 白天温度 */
  daytemp: string;
  /** 夜间温度 */
  nighttemp: string;
  /** 白天风向 */
  daywind: string;
  /** 夜间风向 */
  nightwind: string;
  /** 白天风力 */
  daypower: string;
  /** 夜间风力 */
  nightpower: string;
  /** 白天温度浮点数 */
  daytemp_float: string;
  /** 夜间温度浮点数 */
  nighttemp_float: string;
}

// 高德地图天气预报数据
interface AmapWeatherForecast {
  /** 城市名 */
  city: string;
  /** 区域编码 */
  adcode: string;
  /** 省份名 */
  province: string;
  /** 数据发布时间 */
  reporttime: string;
  /** 预报数据数组 */
  casts: AmapWeatherCast[];
}

// 高德地图天气预报响应类型
interface AmapWeatherResponse {
  status: string;
  count: string;
  info: string;
  infocode: string;
  forecasts: AmapWeatherForecast[];
}

interface Props {
  /** 城市区域编码，如果不提供则使用IP定位 */
  adcode?: string;
}

defineOptions({
  name: 'WeatherCard',
});

const props = withDefaults(defineProps<Props>(), {
  adcode: undefined,
});

// 高德地图API Key - 从环境变量中获取
const AMAP_API_KEY = import.meta.env.VITE_AMAP_WEB_KEY;

const weatherData = ref<AmapWeatherForecast | null>(null);
const todayWeather = ref<AmapWeatherCast | null>(null);
const locationData = ref<AmapIpResponse | null>(null);
const loading = ref(true);
const error = ref('');
const lastUpdateTime = ref<string>('');

// 请求超时控制
const createTimeoutPromise = (timeout: number) => {
  return new Promise((_, reject) => {
    setTimeout(() => reject(new Error('请求超时')), timeout);
  });
};

// 带超时的fetch请求
const fetchWithTimeout = async (
  url: string,
  options: RequestInit = {},
  timeout = 10_000,
) => {
  return Promise.race([
    fetch(url, options),
    createTimeoutPromise(timeout),
  ]) as Promise<Response>;
};

// 获取IP定位信息
const fetchLocation = async (): Promise<AmapIpResponse | null> => {
  try {
    const url = `https://restapi.amap.com/v3/ip?key=${AMAP_API_KEY}`;

    const response = await fetchWithTimeout(
      url,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      },
      8000,
    );

    if (!response.ok) {
      throw new Error(`获取定位失败: ${response.status}`);
    }

    const locationResponse: AmapIpResponse = await response.json();

    if (locationResponse.status !== '1') {
      throw new Error(`定位API返回错误: ${locationResponse.info}`);
    }

    return locationResponse;
  } catch (error_) {
    console.error('获取定位信息错误:', error_);
    throw error_;
  }
};

// 获取天气预报数据 - 使用高德地图天气API
const fetchWeatherByAdcode = async (
  adcode: string,
): Promise<AmapWeatherForecast | null> => {
  try {
    const url = `https://restapi.amap.com/v3/weather/weatherInfo?key=${AMAP_API_KEY}&city=${adcode}&extensions=all`;

    const response = await fetchWithTimeout(
      url,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      },
      8000,
    );

    if (!response.ok) {
      throw new Error(`获取天气失败: ${response.status}`);
    }

    const weatherResponse: AmapWeatherResponse = await response.json();

    if (weatherResponse.status !== '1') {
      throw new Error(`天气API返回错误: ${weatherResponse.info}`);
    }

    if (!weatherResponse.forecasts || weatherResponse.forecasts.length === 0) {
      throw new Error('天气预报数据为空');
    }

    return weatherResponse.forecasts[0] || null;
  } catch (error_) {
    console.error('获取天气数据错误:', error_);
    throw error_;
  }
};

// 综合获取定位和天气数据
const fetchWeather = async () => {
  try {
    loading.value = true;
    error.value = '';

    // 检查API Key
    if (!AMAP_API_KEY) {
      throw new Error('请在环境变量中配置 VITE_AMAP_WEB_KEY');
    }

    // 优先使用传入的adcode，否则通过IP定位获取
    let adcode: string | undefined = props.adcode;

    if (!adcode) {
      // 通过IP定位获取城市编码
      const locationResult = await fetchLocation();
      if (locationResult) {
        locationData.value = locationResult;
        adcode = locationResult.adcode;
      }
    }
    if (adcode === undefined || adcode === '' || adcode.length === 0) {
      throw new Error('无法获取有效的城市信息');
    }
    // 获取天气预报数据
    const weather = await fetchWeatherByAdcode(adcode);
    if (weather) {
      weatherData.value = weather;
      // 获取今天的天气数据（第一条数据）
      todayWeather.value =
        weather.casts && weather.casts.length > 0
          ? weather.casts[0] || null
          : null;
      lastUpdateTime.value = new Date().toLocaleString();
    }
  } catch (error_) {
    console.error('获取天气数据错误:', error_);
    error.value = `获取天气数据失败: ${
      error_ instanceof Error ? error_.message : '未知错误'
    }`;
  } finally {
    loading.value = false;
  }
};

// 格式化更新时间 - 高德API返回的是标准日期时间格式
const formatUpdateTime = (timeStr: string) => {
  if (!timeStr) return '';

  try {
    // 高德API返回格式如: "2024-01-15 14:30:00"
    const date = new Date(timeStr);
    if (Number.isNaN(date.getTime())) {
      return timeStr; // 如果解析失败，直接返回原字符串
    }

    return date.toLocaleString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    });
  } catch {
    return timeStr;
  }
};

// 根据天气状况返回对应颜色
const getWeatherColor = (weather: string) => {
  const weatherColorMap: Record<string, string> = {
    晴: '#FFD700',
    多云: '#87CEEB',
    阴: '#708090',
    小雨: '#4682B4',
    中雨: '#1E90FF',
    大雨: '#0000FF',
    雷阵雨: '#8B008B',
    雪: '#F0F8FF',
    雾: '#D3D3D3',
    霾: '#A0522D',
  };

  for (const [key, color] of Object.entries(weatherColorMap)) {
    if (weather.includes(key)) {
      return color;
    }
  }

  return '#3B82F6'; // 默认蓝色
};

// 获取位置显示文本
const getLocationDisplay = () => {
  if (weatherData.value?.city && weatherData.value?.province) {
    // 优先显示天气数据中的位置信息
    return `${weatherData.value.province}·${weatherData.value.city}`;
  }

  if (locationData.value) {
    // 使用定位数据中的位置信息
    return `${locationData.value.province}·${locationData.value.city}`;
  }

  return '未知位置';
};

// 判断是否为白天（简单逻辑：6-18点为白天）
const isDaytime = () => {
  const hour = new Date().getHours();
  return hour >= 6 && hour < 18;
};

// 获取当前应该显示的天气信息（白天或夜间）
const getCurrentWeatherInfo = () => {
  if (!todayWeather.value) return null;

  const isDay = isDaytime();
  return {
    weather: isDay
      ? todayWeather.value.dayweather
      : todayWeather.value.nightweather,
    temp: isDay ? todayWeather.value.daytemp : todayWeather.value.nighttemp,
    wind: isDay ? todayWeather.value.daywind : todayWeather.value.nightwind,
    power: isDay ? todayWeather.value.daypower : todayWeather.value.nightpower,
    tempFloat: isDay
      ? todayWeather.value.daytemp_float
      : todayWeather.value.nighttemp_float,
  };
};

// 手动刷新天气数据
const refreshWeather = async () => {
  if (loading.value) return; // 防止重复请求
  await fetchWeather();
};

onMounted(() => {
  fetchWeather();
});
</script>

<template>
  <Card class="w-full">
    <template #title>
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <CloudOutlined class="mr-2" />
          实时天气
        </div>
        <div class="flex items-center space-x-2">
          <span v-if="lastUpdateTime" class="text-xs text-gray-400">
            {{ lastUpdateTime }}
          </span>
          <button
            :disabled="loading"
            class="text-blue-500 hover:text-blue-700 disabled:cursor-not-allowed disabled:opacity-50"
            @click="refreshWeather"
          >
            <LoadingOutlined v-if="loading" class="animate-spin" />
            <CloudOutlined v-else />
          </button>
        </div>
      </div>
    </template>

    <Spin :spinning="loading" :indicator="LoadingOutlined">
      <div
        v-if="error"
        class="flex items-center justify-center py-8 text-red-500"
      >
        <EyeOutlined class="mr-2" />
        {{ error }}
      </div>

      <div v-else-if="weatherData && todayWeather" class="space-y-4">
        <!-- 主要天气信息 -->
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <CloudOutlined
              class="text-3xl"
              :style="{
                color: getWeatherColor(getCurrentWeatherInfo()?.weather || ''),
              }"
            />
            <div>
              <div class="text-2xl font-bold">
                {{ getCurrentWeatherInfo()?.temp }}°C
              </div>
              <div class="text-sm text-gray-500">
                {{ getCurrentWeatherInfo()?.weather }}
              </div>
              <div class="text-xs text-gray-400">
                {{ isDaytime() ? '白天' : '夜间' }}
              </div>
            </div>
          </div>
          <div class="text-right">
            <div class="flex items-center text-sm text-gray-500">
              <EnvironmentOutlined class="mr-1" />
              {{ getLocationDisplay() }}
            </div>
            <div class="text-xs text-gray-400">
              {{ formatUpdateTime(weatherData.reporttime) }}
            </div>
          </div>
        </div>

        <!-- 详细信息 -->
        <div class="grid grid-cols-2 gap-4 pt-2">
          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-500">日期</span>
            <span class="text-sm font-medium">{{ todayWeather.date }}</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-500">星期</span>
            <span class="text-sm font-medium">星期{{ todayWeather.week }}</span>
          </div>

          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-500">风向</span>
            <span class="text-sm font-medium">
              {{ getCurrentWeatherInfo()?.wind }}
            </span>
          </div>

          <div class="flex items-center justify-between">
            <span class="text-sm text-gray-500">风力</span>
            <span class="text-sm font-medium">
              {{ getCurrentWeatherInfo()?.power }}级
            </span>
          </div>
        </div>

        <!-- 温度范围显示 -->
        <div class="rounded-md bg-gradient-to-r from-blue-50 to-orange-50 p-3">
          <div class="flex items-center justify-between">
            <div class="text-center">
              <div class="text-xs text-gray-500">白天</div>
              <div class="font-medium text-orange-600">
                {{ todayWeather.daytemp }}°C
              </div>
              <div class="text-xs text-gray-400">
                {{ todayWeather.dayweather }}
              </div>
            </div>
            <div class="text-center">
              <div class="text-xs text-gray-500">夜间</div>
              <div class="font-medium text-blue-600">
                {{ todayWeather.nighttemp }}°C
              </div>
              <div class="text-xs text-gray-400">
                {{ todayWeather.nightweather }}
              </div>
            </div>
          </div>
        </div>

        <!-- 未来几天预报预览 -->
        <div
          v-if="weatherData.casts && weatherData.casts.length > 1"
          class="space-y-2"
        >
          <div class="text-sm font-medium text-gray-700">未来几天</div>
          <div class="grid grid-cols-3 gap-2">
            <div
              v-for="cast in weatherData.casts.slice(1, 4)"
              :key="cast.date"
              class="rounded-md bg-gray-50 p-2 text-center"
            >
              <div class="text-xs text-gray-500">{{ cast.date.slice(5) }}</div>
              <div class="text-xs font-medium">{{ cast.dayweather }}</div>
              <div class="text-xs text-gray-600">
                {{ cast.daytemp }}°/{{ cast.nighttemp }}°
              </div>
            </div>
          </div>
        </div>
      </div>
    </Spin>
  </Card>
</template>
