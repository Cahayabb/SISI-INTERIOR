<template>
  <nav
    class="navbar"
    :class="{ 'navbar--scrolled': isScrolled, 'navbar--open': menuOpen }"
  >
    <div class="navbar__container">

      <!-- Logo -->
      <RouterLink to="/" class="navbar__logo" @click="closeMenu">
        <img
          src="@/assets/images/logo-sisi.png"
          alt="SISI Interior"
          class="navbar__logo-img"
          @error="onLogoError"
          ref="logoImg"
        />
        <div class="navbar__logo-fallback" ref="logoFallback" style="display:none">
          <span class="navbar__logo-main">SISI</span>
          <span class="navbar__logo-sub">INTERIOR</span>
        </div>
      </RouterLink>

      <!-- Links Desktop -->
      <ul class="navbar__links">
        <li v-for="link in publicLinks" :key="link.to">
          <RouterLink
            :to="link.to"
            class="navbar__link"
            :class="{ 'navbar__link--active': isActive(link.to) }"
            @click="closeMenu"
          >
            {{ link.label }}
          </RouterLink>
        </li>
        <!-- Estimasi hanya muncul jika customer login -->
        <li v-if="isCustomer">
          <RouterLink
            to="/estimasi"
            class="navbar__link navbar__link--estimasi"
            :class="{ 'navbar__link--active': isActive('/estimasi') }"
            @click="closeMenu"
          >
            Estimasi
          </RouterLink>
        </li>
      </ul>

      <!-- CTA: Login (guest) | Avatar dropdown (logged in) -->
      <RouterLink v-if="!isLoggedIn" to="/login" class="navbar__cta" @click="closeMenu">
        Login
      </RouterLink>

      <!-- Avatar + Dropdown (desktop) -->
      <div v-else class="avatar-menu" ref="avatarMenuRef">
        <button class="navbar__avatar" @click="toggleAvatarMenu" :title="displayName">
          <img v-if="avatarUrl" :src="avatarUrl" alt="User" @error="avatarUrl = ''" />
          <span v-else class="avatar-initials">{{ userInitials }}</span>
        </button>

        <Transition name="dropdown">
          <div class="nav-dropdown" v-if="avatarMenuOpen">
            <!-- Profile header -->
            <div class="nav-dropdown__header">
              <div class="nav-dropdown__avatar">
                <img v-if="avatarUrl" :src="avatarUrl" alt="User" @error="avatarUrl = ''" />
                <span v-else class="avatar-initials avatar-initials--md">{{ userInitials }}</span>
              </div>
              <div class="nav-dropdown__info">
                <p class="nav-dropdown__name">{{ displayName }}</p>
                <p class="nav-dropdown__role">{{ roleLabel }}</p>
              </div>
            </div>

            <div class="nav-dropdown__divider"></div>

            <!-- Lihat Profile -->
            <button class="nav-dropdown__item" @click="goToProfile">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
              Lihat Profile
            </button>

            <!-- Logout -->
            <button class="nav-dropdown__item nav-dropdown__item--logout" @click="handleLogout">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
                <polyline points="16 17 21 12 16 7"/>
                <line x1="21" y1="12" x2="9" y2="12"/>
              </svg>
              Logout
            </button>
          </div>
        </Transition>
      </div>

      <!-- Hamburger -->
      <button
        class="navbar__hamburger"
        @click="toggleMenu"
        :aria-expanded="menuOpen"
        aria-label="Toggle menu"
      >
        <span class="navbar__hamburger-line" />
        <span class="navbar__hamburger-line" />
        <span class="navbar__hamburger-line" />
      </button>

    </div>

    <!-- Mobile Menu -->
    <Transition name="mobile-menu">
      <div class="navbar__mobile" v-if="menuOpen">
        <!-- Mobile: user info strip jika login -->
        <div class="mobile-user" v-if="isLoggedIn">
          <div class="mobile-user__avatar">
            <img v-if="avatarUrl" :src="avatarUrl" alt="User" @error="avatarUrl = ''" />
            <span v-else class="avatar-initials avatar-initials--md">{{ userInitials }}</span>
          </div>
          <div>
            <p class="mobile-user__name">{{ displayName }}</p>
            <p class="mobile-user__role">{{ roleLabel }}</p>
          </div>
        </div>
        <div class="mobile-user__divider" v-if="isLoggedIn"></div>

        <ul class="navbar__mobile-links">
          <li v-for="(link, i) in publicLinks" :key="link.to" :style="{ '--i': i }">
            <RouterLink
              :to="link.to"
              class="navbar__mobile-link"
              :class="{ 'navbar__mobile-link--active': isActive(link.to) }"
              @click="closeMenu"
            >
              {{ link.label }}
            </RouterLink>
          </li>
          <li v-if="isCustomer" :style="{ '--i': publicLinks.length }">
            <RouterLink
              to="/estimasi"
              class="navbar__mobile-link navbar__mobile-link--estimasi"
              :class="{ 'navbar__mobile-link--active': isActive('/estimasi') }"
              @click="closeMenu"
            >
              Estimasi
            </RouterLink>
          </li>
        </ul>

        <!-- Mobile: Login atau Profile + Logout -->
        <template v-if="!isLoggedIn">
          <RouterLink
            to="/login"
            class="navbar__cta navbar__cta--mobile"
            @click="closeMenu"
          >
            Login
          </RouterLink>
        </template>
        <template v-else>
          <div class="mobile-actions">
            <button class="mobile-btn mobile-btn--profile" @click="goToProfile">
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
              Lihat Profile
            </button>
            <button class="mobile-btn mobile-btn--logout" @click="handleLogout">
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
                <polyline points="16 17 21 12 16 7"/>
                <line x1="21" y1="12" x2="9" y2="12"/>
              </svg>
              Logout
            </button>
          </div>
        </template>
      </div>
    </Transition>
  </nav>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter, RouterLink } from 'vue-router'

const route  = useRoute()
const router = useRouter()

const isScrolled     = ref(false)
const menuOpen       = ref(false)
const avatarMenuOpen = ref(false)
const logoImg        = ref(null)
const logoFallback   = ref(null)
const avatarMenuRef  = ref(null)
const avatarUrl      = ref('')

const publicLinks = [
  { label: 'Home',       to: '/' },
  { label: 'About',      to: '/about' },
  { label: 'Service',    to: '/service' },
  { label: 'Portofolio', to: '/portfolio' },
  { label: 'Contact',    to: '/contact' },
]

// ── Auth state ──
const isLoggedIn = computed(() => !!localStorage.getItem('token'))
const role       = computed(() => localStorage.getItem('role') || '')
const isCustomer = computed(() => role.value === 'customer' && isLoggedIn.value)
const isAdmin    = computed(() => role.value === 'admin' && isLoggedIn.value)
const roleLabel  = computed(() => isAdmin.value ? 'Administrator' : 'Customer')

// ── User data ──
const userData = computed(() => {
  try { return JSON.parse(localStorage.getItem('user') || '{}') } catch { return {} }
})

const displayName = computed(() => {
  const u = userData.value
  const fn = u.first_name || ''
  const ln = u.last_name  || ''
  return (fn + ' ' + ln).trim() || u.username || 'User'
})

const userInitials = computed(() => {
  const u  = userData.value
  const fn = u.first_name?.[0] || ''
  const ln = u.last_name?.[0]  || ''
  return (fn + ln).toUpperCase() || (u.username?.[0] || 'U').toUpperCase()
})

onMounted(() => {
  const saved = localStorage.getItem('admin_avatar') || localStorage.getItem('user_avatar')
  if (saved) avatarUrl.value = saved

  window.addEventListener('scroll', handleScroll, { passive: true })
  document.addEventListener('click', onClickOutside)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  document.removeEventListener('click', onClickOutside)
})

// ── Dropdown outside click ──
const onClickOutside = (e) => {
  if (avatarMenuRef.value && !avatarMenuRef.value.contains(e.target)) {
    avatarMenuOpen.value = false
  }
}

const toggleAvatarMenu = () => { avatarMenuOpen.value = !avatarMenuOpen.value }

// ── Navigation ──
const goToProfile = () => {
  avatarMenuOpen.value = false
  closeMenu()
  if (isAdmin.value) {
    router.push({ name: 'admin-dashboard' })
  } else {
    router.push({ name: 'profile' })
  }
}

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('role')
  localStorage.removeItem('user')
  avatarMenuOpen.value = false
  closeMenu()
  router.push('/login')
}

const isActive   = (path) => path === '/' ? route.path === '/' : route.path.startsWith(path)
const toggleMenu = () => { menuOpen.value = !menuOpen.value }
const closeMenu  = () => { menuOpen.value = false }

const onLogoError = () => {
  if (logoImg.value)      logoImg.value.style.display = 'none'
  if (logoFallback.value) logoFallback.value.style.display = 'flex'
}

const handleScroll = () => { isScrolled.value = window.scrollY > 40 }
</script>

<style scoped>
.navbar {
  --gold:           #C8A135;
  --teal:           #1B7A6E;
  --text-nav:       #2C2C2C;
  --text-nav-muted: #555555;
  --white:          #FFFFFF;
  --nav-height:     72px;

  position: fixed;
  top: 0; left: 0; right: 0;
  z-index: 1000;
  background: var(--white);
  box-shadow: 0 2px 12px rgba(0,0,0,0.07);
  transition: box-shadow 0.3s ease;
}

.navbar--scrolled {
  box-shadow: 0 4px 24px rgba(0,0,0,0.12);
}

/* ── Container ── */
.navbar__container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 40px;
  height: var(--nav-height);
  display: flex;
  align-items: center;
}

/* ── Logo ── */
.navbar__logo {
  display: flex;
  align-items: center;
  text-decoration: none;
  flex-shrink: 0;
  margin-right: auto;
}

.navbar__logo-img {
  height: 40px;
  width: auto;
  object-fit: contain;
  display: block;
}

.navbar__logo-fallback {
  flex-direction: column;
  line-height: 1;
}

.navbar__logo-main {
  font-family: 'Playfair Display', 'Georgia', serif;
  font-size: 22px;
  font-weight: 700;
  color: var(--teal);
  letter-spacing: 6px;
}

.navbar__logo-sub {
  font-family: 'Montserrat', sans-serif;
  font-size: 8px;
  font-weight: 600;
  color: var(--gold);
  letter-spacing: 4px;
  text-transform: uppercase;
  margin-top: 3px;
}

/* ── Nav Links ── */
.navbar__links {
  display: flex;
  align-items: center;
  gap: 4px;
  list-style: none;
  margin: 0;
  padding: 0;
}

.navbar__link {
  display: inline-block;
  padding: 8px 18px;
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: var(--text-nav-muted);
  text-decoration: none;
  letter-spacing: 0.1px;
  transition: color 0.2s ease;
  white-space: nowrap;
}

.navbar__link:hover   { color: var(--teal); }
.navbar__link--active { color: var(--teal); font-weight: 700; }

.navbar__link--estimasi         { color: var(--gold) !important; font-weight: 700; }
.navbar__link--estimasi:hover   { color: #b08e2c !important; }
.navbar__link--estimasi.navbar__link--active { color: #b08e2c !important; }

/* ── CTA Button (Login) ── */
.navbar__cta {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  margin-left: 24px;
  padding: 11px 32px;
  background: var(--gold);
  color: var(--white);
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 700;
  letter-spacing: 0.2px;
  text-decoration: none;
  border-radius: 8px;
  white-space: nowrap;
  flex-shrink: 0;
  border: none;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
}

.navbar__cta:hover {
  background: #b08e2c;
  transform: translateY(-1px);
}

/* ── Avatar Menu (wrapper) ── */
.avatar-menu {
  position: relative;
  margin-left: 20px;
  flex-shrink: 0;
}

/* ── Avatar Button ── */
.navbar__avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
  border: 2.5px solid #E5E7EB;
  cursor: pointer;
  background: var(--teal);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.navbar__avatar:hover {
  border-color: var(--teal);
  box-shadow: 0 0 0 3px rgba(27, 122, 110, 0.18);
}

.navbar__avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

/* ── Avatar Initials ── */
.avatar-initials {
  font-size: 13px;
  font-weight: 700;
  color: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  line-height: 1;
  user-select: none;
}

.avatar-initials--md { font-size: 15px; }

/* ── Dropdown ── */
.nav-dropdown {
  position: absolute;
  top: calc(100% + 12px);
  right: 0;
  width: 210px;
  background: #FFFFFF;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.13);
  border: 1px solid #EFEFEF;
  overflow: hidden;
  z-index: 200;
}

/* Dropdown header */
.nav-dropdown__header {
  display: flex;
  align-items: center;
  gap: 11px;
  padding: 14px 16px;
  background: #F8FFFE;
}

.nav-dropdown__avatar {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  overflow: hidden;
  background: var(--teal);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border: 2px solid #E0F0EE;
}

.nav-dropdown__avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.nav-dropdown__name {
  font-size: 13px;
  font-weight: 700;
  color: #1A1A1A;
  margin: 0 0 2px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 120px;
}

.nav-dropdown__role {
  font-size: 11px;
  color: var(--teal);
  font-weight: 600;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.nav-dropdown__divider {
  height: 1px;
  background: #F0F0F0;
}

/* Dropdown items */
.nav-dropdown__item {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  padding: 12px 16px;
  background: none;
  border: none;
  font-size: 13px;
  font-weight: 500;
  font-family: 'Montserrat', sans-serif;
  color: #374151;
  cursor: pointer;
  text-align: left;
  transition: background 0.15s, color 0.15s;
}

.nav-dropdown__item:hover {
  background: #F5F5F5;
  color: var(--teal);
}

.nav-dropdown__item--logout       { color: #EF4444; }
.nav-dropdown__item--logout:hover { background: #FEF2F2; color: #DC2626; }

/* Dropdown transition */
.dropdown-enter-active,
.dropdown-leave-active { transition: opacity 0.18s ease, transform 0.18s ease; }
.dropdown-enter-from,
.dropdown-leave-to     { opacity: 0; transform: translateY(-6px); }

/* ── Hamburger ── */
.navbar__hamburger {
  display: none;
  flex-direction: column;
  justify-content: center;
  gap: 5px;
  width: 36px;
  height: 36px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  margin-left: auto;
}

.navbar__hamburger-line {
  display: block;
  height: 2px;
  width: 100%;
  background: var(--text-nav);
  border-radius: 2px;
  transition: transform 0.3s ease, opacity 0.3s ease;
  transform-origin: center;
}

.navbar--open .navbar__hamburger-line:nth-child(1) { transform: translateY(7px) rotate(45deg); }
.navbar--open .navbar__hamburger-line:nth-child(2) { opacity: 0; transform: scaleX(0); }
.navbar--open .navbar__hamburger-line:nth-child(3) { transform: translateY(-7px) rotate(-45deg); }

/* ── Mobile Menu ── */
.navbar__mobile {
  background: var(--white);
  padding: 20px 24px 28px;
  border-top: 1px solid rgba(0,0,0,0.07);
  box-shadow: 0 12px 32px rgba(0,0,0,0.1);
}

/* Mobile user strip */
.mobile-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 16px;
}

.mobile-user__avatar {
  width: 42px;
  height: 42px;
  border-radius: 50%;
  overflow: hidden;
  background: var(--teal);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  border: 2px solid #E0F0EE;
}

.mobile-user__avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.mobile-user__name {
  font-size: 14px;
  font-weight: 700;
  color: #1A1A1A;
  margin: 0 0 2px;
}

.mobile-user__role {
  font-size: 11px;
  color: var(--teal);
  font-weight: 600;
  margin: 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.mobile-user__divider {
  height: 1px;
  background: #EEEEEE;
  margin-bottom: 12px;
}

.navbar__mobile-links {
  list-style: none;
  margin: 0 0 20px;
  padding: 0;
}

.navbar__mobile-link {
  display: block;
  padding: 13px 0;
  font-family: 'Montserrat', sans-serif;
  font-size: 15px;
  font-weight: 500;
  color: var(--text-nav-muted);
  text-decoration: none;
  border-bottom: 1px solid rgba(0,0,0,0.06);
  transition: color 0.2s ease, padding-left 0.2s ease;
  animation: slideIn 0.3s ease both;
  animation-delay: calc(var(--i) * 0.06s);
}

.navbar__mobile-link:hover,
.navbar__mobile-link--active { color: var(--teal); padding-left: 8px; }
.navbar__mobile-link--active  { font-weight: 700; }
.navbar__mobile-link--estimasi { color: var(--gold) !important; font-weight: 700; }

/* Mobile actions */
.mobile-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.mobile-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 13px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  transition: all 0.2s;
  border: none;
}

.mobile-btn--profile {
  background: #F0FAF8;
  color: var(--teal);
  border: 1.5px solid #B2DDD8;
}

.mobile-btn--profile:hover {
  background: #D8F0EC;
}

.mobile-btn--logout {
  background: #FEF2F2;
  color: #EF4444;
  border: 1.5px solid #FECACA;
}

.mobile-btn--logout:hover {
  background: #FEE2E2;
  color: #DC2626;
}

.navbar__cta--mobile {
  margin-left: 0;
  width: 100%;
  justify-content: center;
}

/* ── Transitions ── */
.mobile-menu-enter-active,
.mobile-menu-leave-active { transition: opacity 0.25s ease, transform 0.3s ease; }
.mobile-menu-enter-from,
.mobile-menu-leave-to     { opacity: 0; transform: translateY(-10px); }

@keyframes slideIn {
  from { opacity: 0; transform: translateX(-14px); }
  to   { opacity: 1; transform: translateX(0); }
}

/* ── Responsive ── */
@media (max-width: 900px) {
  .navbar__links,
  .navbar__cta:not(.navbar__cta--mobile),
  .avatar-menu {
    display: none;
  }

  .navbar__hamburger {
    display: flex;
  }

  .navbar__container {
    padding: 0 20px;
  }
}
</style>