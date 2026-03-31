<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
/* ── Reset & Base ─────────────────────────────────── */
*, *::before, *::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html {
  font-size: 16px;
  scroll-behavior: smooth;
}

body {
  font-family: 'Segoe UI', system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
  background: #f0f4ff;
  color: #1e293b;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  min-height: 100vh;
}

#app {
  min-height: 100vh;
}

/* ── Leaflet harita container fix ─────────────────── */
/* Leaflet z-index'leri diğer elemanlarla çakışmasın */
.leaflet-container {
  font-family: 'Segoe UI', system-ui, sans-serif !important;
  z-index: 1;
}

.leaflet-control-zoom {
  border: none !important;
  box-shadow: 0 2px 12px rgba(0,0,0,0.12) !important;
}

.leaflet-control-zoom a {
  border-radius: 8px !important;
  border: none !important;
  color: #1e3a5f !important;
  font-weight: 700 !important;
}

.leaflet-control-zoom-in,
.leaflet-control-zoom-out {
  border-radius: 8px !important;
}

/* ── Tooltip stili (Leaflet) ─────────────────────── */
.city-tooltip {
  background: #1e3a5f !important;
  color: white !important;
  border: none !important;
  border-radius: 8px !important;
  font-weight: 600 !important;
  font-size: 12px !important;
  padding: 4px 10px !important;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2) !important;
  white-space: nowrap !important;
}

.city-tooltip::before {
  display: none !important;
}

/* ── Scrollbar ───────────────────────────────────── */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f5f9;
}

::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 99px;
}

::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}

/* ── Focus visible ───────────────────────────────── */
:focus-visible {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
  border-radius: 4px;
}

/* ── Animasyon yardımcıları ──────────────────────── */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}

.fade-in {
  animation: fadeIn 0.3s ease both;
}
</style>
