<template>
  <div class="placeholder-page">
    <h1>404</h1>
    <p>Halaman tidak ditemukan.</p>
    <RouterLink to="/">Kembali ke Home</RouterLink>
  </div>
</template>
<script setup>
import { RouterLink } from 'vue-router'
</script>
<style scoped>
.placeholder-page { min-height:100vh; display:flex; flex-direction:column; align-items:center; justify-content:center; gap:16px; font-family:'Montserrat',sans-serif; background:#F5F5F0; }
h1 { font-size:72px; font-weight:800; color:#1B7A6E; margin:0; } p { color:#666; }
a  { color:#C8A135; font-weight:700; }
</style>