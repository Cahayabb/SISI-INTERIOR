import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // ── Public ──
    { path: '/',          name: 'home',      component: Home },
    { path: '/about',     name: 'about',     component: () => import('../views/About.vue') },
    { path: '/service',   name: 'service',   component: () => import('../views/Service.vue') },
    { path: '/portfolio', name: 'portfolio', component: () => import('../views/Portfolio.vue') },
    { path: '/contact',   name: 'contact',   component: () => import('../views/Contact.vue') },

    // ── Auth ──
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue'),
      meta: { guestOnly: true },
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import('../views/Signup.vue'),
      meta: { guestOnly: true },
    },

    // ── users
    {
      path: '/estimasi',
      name: 'estimasi',
      component: () => import('../views/Estimasi.vue'),
      meta: { requiresAuth: true, role: 'users'},
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../views/Profile.vue'),
      meta: { requiresAuth: true, role: 'users'},
    },

    {
      path: '/admin',
      component: () => import('../components/Sidebar.vue'),
      redirect: '/admin/dashboard',
      meta: { requiresAuth: true, role: 'admin' },
      children: [
        {
          path: 'dashboard',
          name: 'admin-dashboard',
          component: () => import('../views/admin/Dashboard.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
        {
          path: 'portfolio',
          name: 'admin-portfolio',
          component: () => import('../views/admin/Portfolio.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
        {
          path: 'proyek',
          name: 'admin-proyek',
          component: () => import('../views/admin/Project.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
        {
          path: 'training',
          name: 'admin-training',
          component: () => import('../views/admin/Training.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
        {
          path: 'estimasi',
          name: 'admin-estimasi',
          component: () => import('../views/admin/EstimasiBiaya.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
        {
          path: 'analitik',
          name: 'admin-analitik',
          component: () => import('../views/admin/Analitik.vue'),
          meta: { requiresAuth: true, role: 'admin' },
        },
      ],
    },

    // ── 404 ──
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('../views/NotFound.vue'),
    },
  ],

  scrollBehavior() {
    return { top: 0, behavior: 'smooth' }
  },
})

// ── Route Guard ──
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  const role  = localStorage.getItem('role')

  if (to.meta.requiresAuth) {
    if (!token) {
      return next({ name: 'login', query: { redirect: to.fullPath } })
    }
    if (to.meta.role && to.meta.role !== role) {
      return role === 'admin'
        ? next({ name: 'admin-dashboard' })
        : next({ name: 'estimasi' })
    }
  }

  if (to.meta.guestOnly && token) {
    return role === 'admin'
      ? next({ name: 'admin-dashboard' })
      : next({ name: 'estimasi' })
  }

  next()
})

export default router