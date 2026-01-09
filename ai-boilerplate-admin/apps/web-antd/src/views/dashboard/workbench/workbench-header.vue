<script lang="ts" setup>
import { computed } from 'vue';

import { VbenAvatar } from '@vben/common-ui';
import { useUserStore } from '@vben/stores';

interface Props {
  avatar?: string;
}

defineOptions({
  name: 'WorkbenchHeader',
});

withDefaults(defineProps<Props>(), {
  avatar: '',
});

const userStore = useUserStore();

// 动态问候语函数
const getGreeting = computed(() => {
  const hour = new Date().getHours();
  const nickname = userStore.userInfo?.nickname || '朋友';

  const greetings = [
    {
      timeRange: [5, 6] as const,
      messages: [
        `${nickname}，晨光初现！🌅 新的征程即将开始，准备好迎接无限可能了吗？`,
        `早安，${nickname}！☀️ 清晨的第一缕阳光为你而来，今天要做最棒的自己！`,
        `${nickname}，早起的鸟儿有虫吃！🐦 新的一天，新的机遇，让我们一起发光发热！`,
      ],
    },
    {
      timeRange: [7, 8] as const,
      messages: [
        `美好的早晨，${nickname}！🌄 咖啡香气伴随着晨曦，是时候开启高效模式了！`,
        `${nickname}，早安！✨ 每一个清晨都是重新出发的机会，今天的你会更精彩！`,
        `早上好，${nickname}！🎯 新的一天，新的目标，让我们一起创造奇迹！`,
      ],
    },
    {
      timeRange: [9, 10] as const,
      messages: [
        `${nickname}，上午好！💼 工作状态已激活，让我们在专业的道路上勇敢前行！`,
        `活力满满的上午，${nickname}！🚀 是时候展现你的才华和创造力了！`,
        `${nickname}，新的一天正式开始！⚡ 带着满腔热情去征服每一个挑战吧！`,
      ],
    },
    {
      timeRange: [11, 11] as const,
      messages: [
        `${nickname}，上午的黄金时光！⏰ 专注力正值巅峰，是攻克难题的最佳时机！`,
        `临近午间，${nickname}！🎨 创意思维正在高速运转，灵感即将爆发！`,
        `${nickname}，继续保持这份专注！🎯 距离午休还有一点时间，再冲刺一下！`,
      ],
    },
    {
      timeRange: [12, 13] as const,
      messages: [
        `${nickname}，午餐时光！🍽️ 工作辛苦了，记得犒赏一下努力的自己哦！`,
        `中午好，${nickname}！☀️ 暂停一下忙碌的节奏，给大脑和身体充充电！`,
        `${nickname}，该休息一下了！🌼 美食加阳光，下午会有更棒的表现！`,
        `午间休憩，${nickname}！😌 劳逸结合才是持续高效的秘诀！`,
      ],
    },
    {
      timeRange: [14, 15] as const,
      messages: [
        `${nickname}，下午好！🌻 午后的阳光正暖，是时候重新投入战斗了！`,
        `下午的开始，${nickname}！💫 带着午休后的活力，继续创造精彩！`,
        `${nickname}，下午时光！🎪 让我们用饱满的精神状态迎接新的挑战！`,
      ],
    },
    {
      timeRange: [16, 17] as const,
      messages: [
        `${nickname}，下午进行时！⚡ 保持专注，距离今天的胜利越来越近了！`,
        `加油，${nickname}！🏃‍♂️ 下午的冲刺阶段，让我们全力以赴！`,
        `${nickname}，稳住节奏！🎵 每一分努力都在为成功积累能量！`,
      ],
    },
    {
      timeRange: [18, 19] as const,
      messages: [
        `${nickname}，傍晚好！🌇 夕阳西下，今天的努力即将收获满满的成就感！`,
        `晚上好，${nickname}！🌟 华灯初上，你的坚持让这一天格外闪亮！`,
        `${nickname}，傍晚时分！🎭 回顾今天的收获，为明天的精彩做好准备！`,
      ],
    },
    {
      timeRange: [20, 21] as const,
      messages: [
        `${nickname}，夜晚好！🌃 今天的努力值得为自己点个赞，你真的很棒！`,
        `晚间时光，${nickname}！✨ 忙碌了一天，是时候放松一下享受生活了！`,
        `${nickname}，夜幕降临！🌙 今天的每一份付出都是明天成功的基石！`,
      ],
    },
    {
      timeRange: [22, 23] as const,
      messages: [
        `${nickname}，深夜了！🌌 记得早点休息，明天还有更多精彩等着你！`,
        `夜深人静，${nickname}！💤 劳逸结合是智者的选择，好梦！`,
        `${nickname}，该说晚安了！🛌 充足的睡眠是明天高效工作的保证！`,
      ],
    },
    {
      timeRange: [0, 4] as const,
      messages: [
        `${nickname}，夜猫子模式？🦉 虽然夜深了，但你的努力让人敬佩！记得适度休息哦！`,
        `深夜好，${nickname}！🌟 夜空中最亮的星，就是坚持梦想的你！`,
        `${nickname}，凌晨时分！🌙 夜深了要注意身体，明天的太阳还在等你！`,
        `${nickname}，夜已深！😴 再努力也要记得照顾好自己，健康最重要！`,
      ],
    },
  ];

  const currentGreeting = greetings.find((greeting) => {
    const [start, end] = greeting.timeRange;
    return start <= end
      ? hour >= start && hour <= end
      : hour >= start || hour <= end;
  });

  if (currentGreeting?.messages) {
    // 随机选择一条问候语
    const randomIndex = Math.floor(
      Math.random() * currentGreeting.messages.length,
    );
    return currentGreeting.messages[randomIndex];
  }

  // 默认问候语
  const defaultMessages = [
    `你好，${nickname}！让我们一起努力创造美好的一天！✨`,
    `${nickname}，欢迎回来！今天也要元气满满哦！🌈`,
    `嗨，${nickname}！准备好迎接新的挑战了吗？💪`,
    `${nickname}，又见面了！让我们一起创造精彩时刻！🚀`,
  ];

  const randomDefaultIndex = Math.floor(Math.random() * defaultMessages.length);
  return defaultMessages[randomDefaultIndex];
});
</script>
<template>
  <div class="card-box p-4 py-6 lg:flex">
    <VbenAvatar :src="avatar" class="size-20" />
    <div class="flex flex-col justify-center md:ml-6 md:mt-0">
      <h1 class="text-md font-semibold md:text-xl">
        {{ getGreeting }}
      </h1>
      <span v-if="$slots.description" class="text-foreground/80 mt-1">
        <slot name="description"></slot>
      </span>
    </div>
  </div>
</template>
