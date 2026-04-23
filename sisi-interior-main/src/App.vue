<template>
  <AppNavbar v-if="!isAdmin && !isAuthPage" :key="route.fullPath" />
  <main :class="{ 'main--with-navbar': !isAdmin && !isAuthPage }">
    <RouterView />
  </main>
  <AppFooter v-if="!isAdmin && !isAuthPage" />
</template>

<script setup>
import { computed } from 'vue'
import { RouterView, useRoute } from 'vue-router'
import AppNavbar from '@/components/AppNavbar.vue'
import AppFooter from '@/components/AppFooter.vue'

const route      = useRoute()
const isAdmin    = computed(() => route.path.startsWith('/admin'))
const isAuthPage = computed(() => ['/login', '/signup'].includes(route.path))
</script>