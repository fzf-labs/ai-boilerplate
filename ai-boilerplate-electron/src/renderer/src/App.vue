<script setup lang="ts">
import Versions from './components/Versions.vue'

const modules = [
  'Admin console',
  'Mobile app',
  'Backend API',
  'Desktop shell'
]
</script>

<template>
  <main class="desktop-shell">
    <section class="intro">
      <p class="eyebrow">AI Boilerplate Electron</p>
      <h1>Desktop workspace for product operations.</h1>
      <p class="summary">
        Use this shell as the Electron entry point for the same backend, admin, and mobile
        workflows that are included in the repository.
      </p>
    </section>

    <section class="module-list" aria-label="Included modules">
      <div v-for="module in modules" :key="module" class="module-item">
        <span class="status-dot"></span>
        <span>{{ module }}</span>
      </div>
    </section>

    <Versions />
  </main>
</template>
