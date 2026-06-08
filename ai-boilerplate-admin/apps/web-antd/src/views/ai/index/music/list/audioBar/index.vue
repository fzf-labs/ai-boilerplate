<script lang="ts" setup>
import { computed, inject, reactive, ref } from 'vue';

import { IconifyIcon } from '@vben/icons';

import { Image, Slider } from 'ant-design-vue';

defineOptions({ name: 'AiMusicAudioBarIndex' });

const currentSong = inject<any>('currentSong', {});

const audioRef = ref<HTMLAudioElement | null>(null);
type SliderValue = number | [number, number];

// Audio state for UI and playback control.
const audioState = reactive({
  autoplay: true,
  paused: false,
  currentTime: 0,
  duration: 0,
  muted: false,
});
const volumePercent = ref(50);

function formatSeconds(value: number) {
  const safeValue = Number.isFinite(value) ? Math.floor(value) : 0;
  const minutes = Math.floor(safeValue / 60);
  const seconds = safeValue % 60;
  return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
}

const currentTimeLabel = computed(() => formatSeconds(audioState.currentTime));
const durationLabel = computed(() => formatSeconds(audioState.duration));

function toggleStatus(type: string) {
  if (type === 'paused') {
    audioState.paused = !audioState.paused;
    if (audioRef.value) {
      if (audioState.paused) {
        audioRef.value.pause();
      } else {
        audioRef.value.play();
      }
    }
    return;
  }
  if (type === 'muted') {
    audioState.muted = !audioState.muted;
    if (audioRef.value) {
      audioRef.value.muted = audioState.muted;
    }
  }
}

function normalizeSliderValue(value: SliderValue) {
  return Array.isArray(value) ? (value[0] ?? 0) : value;
}

function seekAudio(value: SliderValue) {
  if (!audioRef.value) {
    return;
  }
  const nextValue = normalizeSliderValue(value);
  audioRef.value.currentTime = nextValue;
  audioState.currentTime = nextValue;
}

function updateDuration() {
  if (!audioRef.value) {
    return;
  }
  const duration = audioRef.value.duration;
  audioState.duration = Number.isFinite(duration) ? duration : 0;
}

function updateVolume(value: SliderValue) {
  const nextValue = normalizeSliderValue(value);
  volumePercent.value = nextValue;
  if (audioRef.value) {
    audioRef.value.volume = Math.min(1, Math.max(0, nextValue / 100));
  }
}

// 更新播放位置
function audioTimeUpdate() {
  if (!audioRef.value) {
    return;
  }
  audioState.currentTime = audioRef.value.currentTime || 0;
  updateDuration();
}
</script>

<template>
  <div
    class="b-1 b-l-none h-18 bg-card flex items-center justify-between border border-solid border-rose-100 px-2"
  >
    <!-- 歌曲信息 -->
    <div class="flex gap-2.5">
      <Image
        src="/favicon.svg"
        :width="45"
      />
      <div>
        <div>{{ currentSong.title }}</div>
        <div class="text-xs text-gray-400">{{ currentSong.desc }}</div>
      </div>
    </div>
    <!-- 音频controls -->
    <div class="flex items-center gap-3">
      <IconifyIcon
        icon="majesticons:back-circle"
        class="size-5 cursor-pointer text-gray-300"
      />
      <IconifyIcon
        :icon="
          audioState.paused
            ? 'mdi:arrow-right-drop-circle'
            : 'solar:pause-circle-bold'
        "
        class="size-7 cursor-pointer"
        @click="toggleStatus('paused')"
      />
      <IconifyIcon
        icon="majesticons:next-circle"
        class="size-5 cursor-pointer text-gray-300"
      />
      <div class="flex items-center gap-4">
        <span>{{ currentTimeLabel }}</span>
        <Slider
          v-model:value="audioState.currentTime"
          :max="audioState.duration"
          color="#409eff"
          class="!w-40"
          @change="seekAudio"
        />
        <span>{{ durationLabel }}</span>
      </div>
      <!-- 音频 -->
      <audio
        ref="audioRef"
        controls
        :autoplay="audioState.autoplay"
        :muted="audioState.muted"
        :src="currentSong.audioUrl"
        :volume="volumePercent / 100"
        @loadedmetadata="updateDuration"
        @timeupdate="audioTimeUpdate"
      ></audio>
    </div>
    <div class="flex items-center gap-4">
      <IconifyIcon
        :icon="audioState.muted ? 'tabler:volume-off' : 'tabler:volume'"
        class="size-5 cursor-pointer"
        @click="toggleStatus('muted')"
      />
      <Slider
        v-model:value="volumePercent"
        :max="100"
        color="#409eff"
        class="!w-40"
        @change="updateVolume"
      />
    </div>
  </div>
</template>
