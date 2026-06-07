<script setup lang="ts">
const modules = [
  { name: "Backend", status: "Go API service" },
  { name: "Admin", status: "Operations console" },
  { name: "Mobile", status: "Uni-app client" },
  { name: "Desktop", status: "Tauri shell" },
];
</script>

<template>
  <main class="shell">
    <section class="hero">
      <p class="eyebrow">AI Boilerplate Desktop</p>
      <h1>Product starter workspace</h1>
      <p class="summary">
        A compact desktop entry point for checking the product modules that ship with this template.
      </p>
    </section>

    <section class="module-grid" aria-label="Template modules">
      <article v-for="module in modules" :key="module.name" class="module-card">
        <span class="module-dot"></span>
        <div>
          <h2>{{ module.name }}</h2>
          <p>{{ module.status }}</p>
        </div>
      </article>
    </section>
  </main>
</template>

<style>
:root {
  color: #17202a;
  background: #f4f7f9;
  font-family:
    Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
  font-synthesis: none;
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

body {
  margin: 0;
  min-width: 320px;
  min-height: 100vh;
}

button,
input {
  font: inherit;
}

#app {
  min-height: 100vh;
}
</style>

<style scoped>
.shell {
  box-sizing: border-box;
  min-height: 100vh;
  padding: 56px clamp(20px, 6vw, 72px);
}

.hero {
  max-width: 760px;
}

.eyebrow {
  margin: 0 0 12px;
  color: #1f6f78;
  font-size: 0.84rem;
  font-weight: 700;
  letter-spacing: 0;
  text-transform: uppercase;
}

h1 {
  margin: 0;
  max-width: 720px;
  color: #111827;
  font-size: clamp(2.4rem, 5vw, 4.6rem);
  line-height: 1.02;
}

.summary {
  margin: 20px 0 0;
  max-width: 620px;
  color: #536170;
  font-size: 1.05rem;
  line-height: 1.7;
}

.module-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
  margin-top: 48px;
}

.module-card {
  display: flex;
  min-height: 112px;
  gap: 14px;
  align-items: flex-start;
  box-sizing: border-box;
  padding: 18px;
  border: 1px solid #d7e0e7;
  border-radius: 8px;
  background: #ffffff;
}

.module-dot {
  flex: 0 0 auto;
  width: 10px;
  height: 10px;
  margin-top: 8px;
  border-radius: 999px;
  background: #2f8f83;
}

h2 {
  margin: 0;
  color: #17202a;
  font-size: 1rem;
}

.module-card p {
  margin: 8px 0 0;
  color: #647281;
  line-height: 1.5;
}

@media (max-width: 860px) {
  .module-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 560px) {
  .shell {
    padding: 36px 18px;
  }

  .module-grid {
    grid-template-columns: 1fr;
  }
}
</style>
