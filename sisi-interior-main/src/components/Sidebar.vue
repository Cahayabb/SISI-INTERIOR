<template>
  <div class="admin-shell">

    <!-- Sidebar -->
    <aside class="sidebar">

      <!-- Logo -->
        <div class="sidebar__logo">
        <img 
          src="@/assets/images/logo2.png" 
          alt="Logo SISI Interior"
          class="sidebar__logo-img"
          @error="onLogoError"
          ref="logoImg"
        />
        <div class="sidebar__logo-fallback" ref="logoFallback">S</div>
      </div>
      <!-- Navigation -->
      <nav class="sidebar__nav">
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="sidebar__nav-item"
          :class="{ 'sidebar__nav-item--active': isActive(item.to) }"
        >
          <span class="sidebar__nav-icon" v-html="item.icon"></span>
          <span class="sidebar__nav-label">{{ item.label }}</span>
        </RouterLink>
      </nav>

    </aside>

    <!-- Main Area -->
    <div class="admin-main">
      <RouterView />
    </div>

  </div>
</template>

<script setup>
import { useRoute, RouterLink, RouterView } from 'vue-router'

const route = useRoute()
const isActive = (path) => route.path === path || route.path.startsWith(path + '/')

const navItems = [
  {
    label: 'Dashboard',
    to: '/admin/dashboard',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect x="3" y="3" width="7" height="7" rx="1"/><rect x="14" y="3" width="7" height="7" rx="1"/>
      <rect x="3" y="14" width="7" height="7" rx="1"/><rect x="14" y="14" width="7" height="7" rx="1"/>
    </svg>`,
  },
  {
    label: 'Data Proyek',
    to: '/admin/proyek',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
      <polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/>
    </svg>`,
  },
  {
    label: 'Training Model',
    to: '/admin/training',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
    </svg>`,
  },
  {
    label: 'Estimasi Biaya',
    to: '/admin/estimasi',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/>
      <path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/>
    </svg>`,
  },
  {
    label: 'Analitik',
    to: '/admin/analitik',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
    </svg>`,
  },
  {
    label: 'Portofolio',
    to: '/admin/portfolio',
    icon: `<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
      <circle cx="12" cy="7" r="4"/>
      <path d="M12 11v4M10 13h4"/>
    </svg>`,
  },
]
</script>

<style scoped>
/* ── Shell ── */
.admin-shell {
  display: flex;
  min-height: 100vh;
  background: #F4F5F7;
  font-family: 'Montserrat', sans-serif;
}

/* ── Sidebar ── */
.sidebar {
  width: 240px;
  flex-shrink: 0;
  background: #FFFFFF;
  border-right: 1px solid #EBEBEB;
  display: flex;
  flex-direction: column;
  padding: 32px 0 24px;
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  z-index: 100;
  box-shadow: 2px 0 16px rgba(0,0,0,0.04);
}

/* ── Logo ── */
.sidebar__logo {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 0 16px 24px;
  margin-bottom: 16px;
  border-bottom: 1px solid #F0F0F0;
}

.sidebar__logo-img {
  width: 120px;     
  height: auto;     
  object-fit: contain;
  display: block;
}

/* ── Fallback ── */
.sidebar__logo-fallback {
  display: none;
  width: 50px;
  height: 50px;
  background: #1B7A6E;
  color: #fff;
  font-weight: 700;
  font-size: 18px;
  border-radius: 10px;
  align-items: center;
  justify-content: center;
}

/* ── Nav ── */
.sidebar__nav {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 0 16px;
  flex: 1;
}

.sidebar__nav-item {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 13px 16px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #6B7280;
  text-decoration: none;
  transition: background 0.18s ease, color 0.18s ease;
  cursor: pointer;
}

.sidebar__nav-item:hover {
  background: #F0FAF8;
  color: #1B7A6E;
}

.sidebar__nav-item--active {
  background: #1B7A6E;
  color: #FFFFFF;
  font-weight: 700;
  box-shadow: 0 4px 12px rgba(27, 122, 110, 0.25);
}

.sidebar__nav-item--active .sidebar__nav-icon {
  color: #FFFFFF;
}

.sidebar__nav-icon {
  display: flex;
  align-items: center;
  flex-shrink: 0;
  color: inherit;
}

.sidebar__nav-label {
  color: inherit;
}

/* ── Main ── */
.admin-main {
  flex: 1;
  margin-left: 240px;
  min-height: 100vh;
  background: #F4F5F7;
}
</style>