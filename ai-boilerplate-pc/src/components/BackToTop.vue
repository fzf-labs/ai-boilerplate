<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue'

const isVisible = ref(false)
const showAfter = 200

const onScroll = () => {
  isVisible.value = window.scrollY > showAfter
}

const scrollToTop = () => {
  const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches
  window.scrollTo({ top: 0, behavior: prefersReducedMotion ? 'auto' : 'smooth' })
}

onMounted(() => {
  onScroll()
  window.addEventListener('scroll', onScroll, { passive: true })
})

onUnmounted(() => {
  window.removeEventListener('scroll', onScroll)
})
</script>

<template>
  <button
    v-show="isVisible"
    class="back-to-top"
    type="button"
    aria-label="Back to top"
    @click="scrollToTop"
  >
    <span aria-hidden="true">↑</span>
  </button>
</template>

<style scoped>
.back-to-top {
  position: fixed;
  right: 24px;
  bottom: 24px;
  z-index: 1000;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: 0;
  border-radius: 999px;
  background: #111;
  color: #fff;
  font-size: 18px;
  cursor: pointer;
  opacity: 0.85;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  transition: transform 0.2s ease, opacity 0.2s ease;
}

.back-to-top:hover {
  opacity: 1;
  transform: translateY(-2px);
}

.back-to-top:focus-visible {
  outline: 2px solid #2f80ed;
  outline-offset: 3px;
}

@media (max-width: 640px) {
  .back-to-top {
    right: 16px;
    bottom: 16px;
    width: 36px;
    height: 36px;
    font-size: 16px;
  }
}
</style>
