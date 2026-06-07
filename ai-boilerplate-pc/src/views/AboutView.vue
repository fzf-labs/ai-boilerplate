<template>
  <section class="about">
    <p class="eyebrow">AI Boilerplate PC</p>
    <h1>Browser workspace for product operations.</h1>
    <p class="lead">
      This package is the desktop web shell for the shared product stack, ready for dashboards,
      admin tools, and support workflows.
    </p>
  </section>
</template>

<style scoped>
.about {
  box-sizing: border-box;
  min-height: 100vh;
  padding: 64px clamp(20px, 5vw, 72px);
  background: #f4f7f9;
  color: #17202a;
}

.eyebrow {
  margin: 0 0 12px;
  color: #1f6f78;
  font-size: 0.84rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0;
}

h1 {
  margin: 0;
  max-width: 680px;
  font-size: clamp(2rem, 4vw, 3.6rem);
  line-height: 1.05;
}

.lead {
  margin: 20px 0 0;
  max-width: 640px;
  color: #536170;
  font-size: 1.02rem;
  line-height: 1.7;
}
</style>
