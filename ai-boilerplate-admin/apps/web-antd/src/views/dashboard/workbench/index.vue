<script lang="ts" setup>
import { preferences } from '@vben/preferences';
import { useUserStore } from '@vben/stores';

import WeatherCard from './weather-card.vue';
import WorkbenchHeader from './workbench-header.vue';

const userStore = useUserStore();
</script>

<template>
  <div class="p-5">
    <WorkbenchHeader
      :avatar="userStore.userInfo?.avatar || preferences.app.defaultAvatar"
    />

    <!-- 天气卡片 -->
    <div class="mt-6">
      <div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
        <WeatherCard />
      </div>
    </div>
  </div>
</template>
