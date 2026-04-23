<template>
  <div class="dashboard">

    <!-- Top Bar -->
    <div class="topbar">
      <div class="topbar__greeting">
        <h1 class="topbar__hello">Hello, <span class="topbar__name">{{ userName }}</span></h1>
        <p class="topbar__sub">Selamat Datang di Dashboard Admin <span class="topbar__brand">SISI Interior</span></p>
      </div>
      <div class="topbar__right">
        <button class="topbar__search">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
          </svg>
        </button>

        <!-- Avatar + Dropdown -->
        <div class="avatar-menu" ref="avatarMenuRef">
          <button class="topbar__avatar" @click="toggleMenu">
            <img v-if="avatarUrl" :src="avatarUrl" alt="Admin" @error="avatarUrl = ''" />
            <span v-else class="avatar-initials">{{ userInitials }}</span>
          </button>

          <Transition name="dropdown">
            <div class="dropdown" v-if="menuOpen">
              <div class="dropdown__profile">
                <div class="dropdown__avatar">
                  <img v-if="avatarUrl" :src="avatarUrl" alt="Admin" @error="avatarUrl = ''" />
                  <span v-else class="avatar-initials avatar-initials--lg">{{ userInitials }}</span>
                </div>
                <div class="dropdown__info">
                  <p class="dropdown__name">{{ fullName }}</p>
                  <p class="dropdown__email">{{ userEmail }}</p>
                </div>
              </div>
              <div class="dropdown__divider"></div>
              <button class="dropdown__item" @click="openProfile">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                  <circle cx="12" cy="7" r="4"/>
                </svg>
                Edit Profile
              </button>
              <button class="dropdown__item dropdown__item--logout" @click="handleLogout">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
                  <polyline points="16 17 21 12 16 7"/>
                  <line x1="21" y1="12" x2="9" y2="12"/>
                </svg>
                Logout
              </button>
            </div>
          </Transition>
        </div>
      </div>
    </div>

    <!-- ══════════════════════════════════════════
         PROFILE EDIT PANEL (inline, not modal)
    ══════════════════════════════════════════ -->
    <Transition name="panel-slide">
      <div class="profile-panel" v-if="showProfile">

        <!-- Panel Header -->
        <div class="panel-header">
          <div class="panel-header__left">
            <div class="panel-header__icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/>
              </svg>
            </div>
            <div>
              <h2 class="panel-header__title">Edit Profile</h2>
              <p class="panel-header__sub">Perubahan nama akan langsung muncul di sapaan dashboard</p>
            </div>
          </div>
          <button class="panel-close" @click="closeProfile">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
            </svg>
          </button>
        </div>

        <!-- Panel Body -->
        <div class="panel-body">

          <!-- Kiri: Avatar Upload -->
          <div class="panel-left">
            <div class="avatar-upload-wrap">
              <div class="avatar-upload">
                <img v-if="form.avatarPreview" :src="form.avatarPreview" alt="Preview" />
                <span v-else class="avatar-initials avatar-initials--xl">{{ formInitials }}</span>
              </div>
              <label class="avatar-upload__btn" title="Ganti foto">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                  <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                  <circle cx="12" cy="13" r="4"/>
                </svg>
                Ganti Foto
                <input type="file" accept="image/*" style="display:none" @change="onAvatarChange" />
              </label>
              <p class="avatar-hint">Format JPG, PNG. Maks 2MB</p>
              <div class="avatar-badge">Administrator</div>
            </div>
          </div>

          <!-- Kanan: Form Fields -->
          <div class="panel-right">
            <div class="form-grid">

              <!-- Nama Depan -->
              <div class="form-group">
                <label class="form-label">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
                  Nama Depan
                </label>
                <input
                  class="form-input"
                  type="text"
                  v-model="form.firstName"
                  placeholder="Masukkan nama depan"
                  @input="onNameInput"
                />
                <span class="form-hint">Akan tampil di: <b>Hello, {{ previewName }}</b></span>
              </div>

              <!-- Nama Belakang -->
              <div class="form-group">
                <label class="form-label">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>
                  Nama Belakang
                </label>
                <input
                  class="form-input"
                  type="text"
                  v-model="form.lastName"
                  placeholder="Masukkan nama belakang"
                />
              </div>

              <!-- Username -->
              <div class="form-group form-group--full">
                <label class="form-label">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-3.92 7.94"/></svg>
                  Username
                </label>
                <div class="input-prefix-wrap">
                  <span class="input-prefix">@</span>
                  <input
                    class="form-input form-input--prefix"
                    type="text"
                    v-model="form.username"
                    placeholder="username"
                  />
                </div>
              </div>

              <!-- Email -->
              <div class="form-group form-group--full">
                <label class="form-label">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/><polyline points="22,6 12,13 2,6"/></svg>
                  Email
                </label>
                <input
                  class="form-input"
                  type="email"
                  v-model="form.email"
                  placeholder="admin@sisi.com"
                />
              </div>

            </div>

            <!-- Action Buttons -->
            <div class="panel-actions">
              <button class="btn-cancel" @click="cancelEdit">Batalkan</button>
              <button class="btn-save" @click="saveProfile" :class="{ 'btn-save--saved': justSaved }">
                <Transition name="btn-swap" mode="out-in">
                  <span v-if="justSaved" key="saved" class="btn-save__content">
                    <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
                    Tersimpan!
                  </span>
                  <span v-else key="save" class="btn-save__content">
                    <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/><polyline points="17 21 17 13 7 13 7 21"/><polyline points="7 3 7 8 15 8"/></svg>
                    Simpan Perubahan
                  </span>
                </Transition>
              </button>
            </div>

          </div>
        </div>
      </div>
    </Transition>

    <!-- Stat Cards -->
    <div class="stat-cards">
      <div class="stat-card" v-for="(card, i) in statCards" :key="i" :style="{ '--delay': i * 0.07 + 's' }">
        <div class="stat-card__icon-wrap">
          <span class="stat-card__icon" v-html="card.icon"></span>
        </div>
        <div class="stat-card__body">
          <template v-if="card.sub">
            <p class="stat-card__label bold">{{ card.label }}</p>
            <p class="stat-card__sub">{{ card.sub }}</p>
          </template>
          <template v-else>
            <p class="stat-card__number">{{ card.number }}</p>
            <p class="stat-card__label">{{ card.label }}</p>
          </template>
        </div>
      </div>
    </div>

    <!-- Analytics Section -->
    <div class="section-header">
      <h2 class="section-title">Ringkasan Analitik</h2>
      <a href="#" class="section-viewall">
        View all
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"/></svg>
      </a>
    </div>

    <div class="chart-card">
      <div class="chart-card__header">
        <h3 class="chart-card__title">Grafik Tren Biaya Proyek</h3>
        <div class="chart-card__filter">
          This Month
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
        </div>
      </div>
      <div class="chart-wrap">
        <svg class="chart-svg" viewBox="0 0 760 270" xmlns="http://www.w3.org/2000/svg">
          <line x1="50" y1="20"  x2="740" y2="20"  stroke="#E5E7EB" stroke-width="1" stroke-dasharray="4,4"/>
          <line x1="50" y1="67"  x2="740" y2="67"  stroke="#E5E7EB" stroke-width="1" stroke-dasharray="4,4"/>
          <line x1="50" y1="114" x2="740" y2="114" stroke="#E5E7EB" stroke-width="1" stroke-dasharray="4,4"/>
          <line x1="50" y1="161" x2="740" y2="161" stroke="#E5E7EB" stroke-width="1" stroke-dasharray="4,4"/>
          <line x1="50" y1="208" x2="740" y2="208" stroke="#E5E7EB" stroke-width="1" stroke-dasharray="4,4"/>
          <line x1="50" y1="235" x2="740" y2="235" stroke="#E5E7EB" stroke-width="1"/>
          <text x="44" y="24"  font-size="10" fill="#9CA3AF" text-anchor="end" font-family="Montserrat,sans-serif">100K</text>
          <text x="44" y="71"  font-size="10" fill="#9CA3AF" text-anchor="end" font-family="Montserrat,sans-serif">80K</text>
          <text x="44" y="118" font-size="10" fill="#9CA3AF" text-anchor="end" font-family="Montserrat,sans-serif">60K</text>
          <text x="44" y="165" font-size="10" fill="#9CA3AF" text-anchor="end" font-family="Montserrat,sans-serif">40K</text>
          <text x="40" y="212" font-size="10" fill="#9CA3AF" text-anchor="end" font-family="Montserrat,sans-serif">20K</text>
          <path d="M75,161 L131,138 L187,161 L243,150 L299,155 L355,138 L411,30 L467,138 L523,114 L579,124 L635,114 L691,128 L740,114 L740,235 L75,235 Z" fill="url(#areaGrad)" />
          <path d="M75,161 L131,138 L187,161 L243,150 L299,155 L355,138 L411,30 L467,138 L523,114 L579,124 L635,114 L691,128 L740,114" fill="none" stroke="#1B7A6E" stroke-width="2.5" stroke-linejoin="round" stroke-linecap="round"/>
          <circle cx="411" cy="30" r="6" fill="#FFFFFF" stroke="#1B7A6E" stroke-width="2.5"/>
          <line x1="411" y1="44" x2="411" y2="235" stroke="#EF4444" stroke-width="1.5" stroke-dasharray="5,4"/>
          <rect x="372" y="8" width="78" height="22" rx="6" fill="#1B7A6E"/>
          <text x="411" y="23" font-size="11" fill="#FFFFFF" text-anchor="middle" font-family="Montserrat,sans-serif" font-weight="700">83,234</text>
          <text x="75"  y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">10</text>
          <text x="131" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">11</text>
          <text x="187" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">12</text>
          <text x="243" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">13</text>
          <text x="299" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">14</text>
          <text x="355" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">15</text>
          <text x="411" y="252" font-size="11" fill="#1B7A6E" text-anchor="middle" font-family="Montserrat,sans-serif" font-weight="700">16</text>
          <text x="467" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">17</text>
          <text x="523" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">18</text>
          <text x="579" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">19</text>
          <text x="635" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">20</text>
          <text x="691" y="252" font-size="10" fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">21</text>
          <text x="728" y="248" font-size="9"  fill="#9CA3AF" text-anchor="middle" font-family="Montserrat,sans-serif">22</text>
          <defs>
            <linearGradient id="areaGrad" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%"   stop-color="#1B7A6E" stop-opacity="0.35"/>
              <stop offset="100%" stop-color="#1B7A6E" stop-opacity="0"/>
            </linearGradient>
          </defs>
        </svg>
      </div>
    </div>

    <!-- Recent Projects -->
    <div class="section-header mt">
      <h2 class="section-title">Data Proyek Terbaru</h2>
      <a href="#" class="section-viewall">
        View all
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"/></svg>
      </a>
    </div>

    <div class="table-card">
      <table class="table">
        <thead>
          <tr>
            <th>No</th>
            <th>Nama Proyek &amp; Klien</th>
            <th>Jenis</th>
            <th>Biaya</th>
            <th>Tanggal</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(row, i) in projects" :key="i">
            <td class="td-no">{{ i + 1 }}</td>
            <td class="td-client">
              <div class="client-row">
                <div class="avatar-table" :style="{ background: row.color }">{{ row.initials }}</div>
                <div>
                  <p class="client-name">{{ row.name }}</p>
                  <p class="client-project">{{ row.project }}</p>
                </div>
              </div>
            </td>
            <td>
              <span class="badge" :class="'badge--' + row.typeKey">{{ row.type }}</span>
            </td>
            <td class="td-biaya">{{ row.biaya }}</td>
            <td class="td-tanggal">{{ row.tanggal }}</td>
            <td class="td-actions">
              <button class="action-btn">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
              </button>
              <button class="action-btn action-btn--del">
                <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="3 6 5 6 21 6"/>
                  <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
                  <path d="M10 11v6"/><path d="M14 11v6"/>
                  <path d="M9 6V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2"/>
                </svg>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

const router        = useRouter()
const menuOpen      = ref(false)
const showProfile   = ref(false)
const avatarMenuRef = ref(null)
const avatarUrl     = ref('')
const justSaved     = ref(false)

const user = ref({})

// Form state (copy of user data for editing)
const form = ref({
  firstName: '',
  lastName: '',
  username: '',
  email: '',
  avatarPreview: '',
})

onMounted(() => {
  try {
    user.value = JSON.parse(localStorage.getItem('user') || '{}')
  } catch {
    user.value = {}
  }
  const savedAvatar = localStorage.getItem('admin_avatar')
  if (savedAvatar) avatarUrl.value = savedAvatar
  document.addEventListener('click', onClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', onClickOutside)
})

// ── Computed ──
const userName = computed(() =>
  user.value.first_name || user.value.username || 'Admin'
)

const fullName = computed(() => {
  const fn = user.value.first_name || ''
  const ln = user.value.last_name  || ''
  return (fn + ' ' + ln).trim() || user.value.username || 'Admin'
})

const userEmail = computed(() =>
  user.value.email || (user.value.username ? user.value.username + '@sisi.com' : 'admin@sisi.com')
)

const userInitials = computed(() => {
  const fn = user.value.first_name?.[0] || ''
  const ln = user.value.last_name?.[0]  || ''
  return (fn + ln).toUpperCase() || 'A'
})

// Live preview for form
const previewName = computed(() =>
  form.value.firstName || form.value.username || 'Admin'
)

const formInitials = computed(() => {
  const fn = form.value.firstName?.[0] || ''
  const ln = form.value.lastName?.[0]  || ''
  return (fn + ln).toUpperCase() || 'A'
})

// ── Menu ──
const toggleMenu = () => { menuOpen.value = !menuOpen.value }

const onClickOutside = (e) => {
  if (avatarMenuRef.value && !avatarMenuRef.value.contains(e.target)) {
    menuOpen.value = false
  }
}

const openProfile = () => {
  menuOpen.value = false
  // Populate form with current user data
  form.value = {
    firstName:     user.value.first_name || '',
    lastName:      user.value.last_name  || '',
    username:      user.value.username   || '',
    email:         user.value.email      || (user.value.username ? user.value.username + '@sisi.com' : ''),
    avatarPreview: avatarUrl.value,
  }
  showProfile.value = true
}

const closeProfile = () => {
  showProfile.value = false
}

const cancelEdit = () => {
  showProfile.value = false
}

const saveProfile = () => {
  // Rebuild object fully — ensures Vue reactive ref triggers computed re-evaluation
  const updated = {
    first_name: (form.value.firstName || '').trim(),
    last_name:  (form.value.lastName  || '').trim(),
    username:   (form.value.username  || '').trim(),
    email:      (form.value.email     || '').trim(),
  }
  user.value = updated
  localStorage.setItem('user', JSON.stringify(updated))

  // Save avatar
  if (form.value.avatarPreview) {
    avatarUrl.value = form.value.avatarPreview
    localStorage.setItem('admin_avatar', form.value.avatarPreview)
  }

  justSaved.value = true
  setTimeout(() => {
    justSaved.value   = false
    showProfile.value = false
  }, 1400)
}

const onNameInput = () => {
  // reactive — previewName computed handles it
}

const onAvatarChange = (e) => {
  const file = e.target.files[0]
  if (!file) return
  const reader = new FileReader()
  reader.onload = (ev) => {
    form.value.avatarPreview = ev.target.result
  }
  reader.readAsDataURL(file)
}

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('role')
  localStorage.removeItem('user')
  router.push('/login')
}

// ── Static Data ──
const statCards = [
  {
    number: '120',
    label: 'Data proyek',
    icon: `<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>`,
  },
  {
    number: '28',
    label: 'Jumlah estimasi',
    icon: `<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="23 7 16 12 23 17 23 7"/><rect x="1" y="5" width="15" height="14" rx="2" ry="2"/></svg>`,
  },
  {
    label: 'Model Aktif',
    sub: 'Last trained: 12 Mei 2026',
    icon: `<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M23 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>`,
  },
  {
    number: '128',
    label: 'Total User Terdaftar',
    icon: `<svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="8" r="4"/><path d="M6 20v-2a6 6 0 0 1 12 0v2"/></svg>`,
  },
]

const projects = [
  { initials: 'JC', name: 'Jane Cooper',   project: 'Kitchen Set Bogor', type: 'Interior',   typeKey: 'interior',   biaya: 'Rp 18.000.000', tanggal: '5/27/15', color: '#3B5BDB' },
  { initials: 'WW', name: 'Wade Warren',   project: 'Renovasi Cafe A',   type: 'Kontraktor', typeKey: 'kontraktor', biaya: 'Rp 40.000.000', tanggal: '5/19/12', color: '#6B7280' },
  { initials: 'EH', name: 'Esther Howard', project: 'Bedroom Decor',     type: 'Interior',   typeKey: 'interior',   biaya: 'Rp 10.000.000', tanggal: '3/4/16',  color: '#3B5BDB' },
  { initials: 'JW', name: 'Jenny Wilson',  project: 'Custom Tabel',      type: 'Custom',     typeKey: 'custom',     biaya: 'Rp 2.000.000',  tanggal: '3/4/16',  color: '#F59F00' },
  { initials: 'GH', name: 'Guy Hawkins',   project: 'Custom Wardrobe',   type: 'Custom',     typeKey: 'custom',     biaya: 'Rp 7.000.000',  tanggal: '7/27/13', color: '#F59F00' },
]
</script>

<style scoped>
.dashboard {
  padding: 36px 36px 60px;
  font-family: 'Montserrat', sans-serif;
}

/* ── Topbar ── */
.topbar {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 32px;
}
.topbar__hello  { font-size: 20px; font-weight: 600; color: #1A1A1A; margin: 0 0 5px; }
.topbar__name   { color: #1B7A6E; font-weight: 700; }
.topbar__sub    { font-size: 16px; font-weight: 700; color: #1A1A1A; margin: 0; }
.topbar__brand  { color: #C8A135; font-weight: 800; }
.topbar__right  { display: flex; align-items: center; gap: 12px; margin-top: 6px; }

.topbar__search {
  width: 40px; height: 40px;
  border-radius: 10px;
  border: 1.5px solid #E5E7EB;
  background: #FFFFFF;
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: #6B7280;
  transition: border-color 0.2s, color 0.2s;
}
.topbar__search:hover { border-color: #1B7A6E; color: #1B7A6E; }

.avatar-menu { position: relative; }

.topbar__avatar {
  width: 40px; height: 40px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid #E5E7EB;
  cursor: pointer;
  background: #1B7A6E;
  display: flex; align-items: center; justify-content: center;
  padding: 0;
  transition: border-color 0.2s, box-shadow 0.2s;
}
.topbar__avatar:hover { border-color: #1B7A6E; box-shadow: 0 0 0 3px rgba(27,122,110,0.15); }
.topbar__avatar img   { width: 100%; height: 100%; object-fit: cover; display: block; }

.avatar-initials        { font-size: 13px; font-weight: 700; color: #FFFFFF; font-family: 'Montserrat', sans-serif; line-height: 1; }
.avatar-initials--lg    { font-size: 20px; }
.avatar-initials--xl    { font-size: 30px; }

/* ── Dropdown ── */
.dropdown {
  position: absolute;
  top: calc(100% + 10px); right: 0;
  width: 220px;
  background: #FFFFFF;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.14);
  border: 1px solid #F0F0F0;
  overflow: hidden; z-index: 100;
}
.dropdown__profile { display: flex; align-items: center; gap: 12px; padding: 16px; }
.dropdown__avatar  { width: 42px; height: 42px; border-radius: 50%; overflow: hidden; background: #1B7A6E; display: flex; align-items: center; justify-content: center; flex-shrink: 0; border: 2px solid #E5E5E5; }
.dropdown__avatar img { width: 100%; height: 100%; object-fit: cover; }
.dropdown__name  { font-size: 13px; font-weight: 700; color: #1A1A1A; margin: 0 0 2px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; max-width: 130px; }
.dropdown__email { font-size: 11px; color: #9CA3AF; margin: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; max-width: 130px; }
.dropdown__divider { height: 1px; background: #F3F3F3; }
.dropdown__item {
  display: flex; align-items: center; gap: 10px;
  width: 100%; padding: 12px 16px;
  background: none; border: none;
  font-size: 13px; font-weight: 500; font-family: 'Montserrat', sans-serif;
  color: #374151; cursor: pointer; text-align: left;
  transition: background 0.15s, color 0.15s;
}
.dropdown__item:hover          { background: #F5F5F5; color: #1B7A6E; }
.dropdown__item--logout        { color: #EF4444; }
.dropdown__item--logout:hover  { background: #FEF2F2; color: #DC2626; }

.dropdown-enter-active, .dropdown-leave-active { transition: opacity 0.18s ease, transform 0.18s ease; }
.dropdown-enter-from, .dropdown-leave-to       { opacity: 0; transform: translateY(-6px); }

/* ══════════════════════════════════════════
   PROFILE PANEL — Inline Edit Form
══════════════════════════════════════════ */
.profile-panel {
  background: #FFFFFF;
  border-radius: 16px;
  border: 1.5px solid #E8F5F3;
  box-shadow: 0 4px 24px rgba(27, 122, 110, 0.08);
  margin-bottom: 28px;
  overflow: hidden;
}

/* Slide-in transition */
.panel-slide-enter-active { transition: all 0.32s cubic-bezier(0.34, 1.2, 0.64, 1); }
.panel-slide-leave-active { transition: all 0.22s ease; }
.panel-slide-enter-from   { opacity: 0; transform: translateY(-16px); }
.panel-slide-leave-to     { opacity: 0; transform: translateY(-8px); }

/* Panel Header */
.panel-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 24px;
  border-bottom: 1.5px solid #F0F0F0;
  background: linear-gradient(135deg, #F0FAF8 0%, #FFFFFF 100%);
}
.panel-header__left {
  display: flex;
  align-items: center;
  gap: 14px;
}
.panel-header__icon {
  width: 38px; height: 38px;
  border-radius: 10px;
  background: #1B7A6E;
  display: flex; align-items: center; justify-content: center;
  color: #FFFFFF;
  flex-shrink: 0;
}
.panel-header__title { font-size: 15px; font-weight: 700; color: #1A1A1A; margin: 0 0 2px; }
.panel-header__sub   { font-size: 11.5px; color: #9CA3AF; margin: 0; }

.panel-close {
  width: 32px; height: 32px;
  border-radius: 8px;
  border: 1.5px solid #E5E5E5;
  background: #FFFFFF;
  cursor: pointer;
  display: flex; align-items: center; justify-content: center;
  color: #9CA3AF;
  transition: all 0.2s;
  flex-shrink: 0;
}
.panel-close:hover { background: #FEF2F2; color: #EF4444; border-color: #EF4444; }

/* Panel Body: 2-column layout */
.panel-body {
  display: flex;
  gap: 0;
}

/* Left: Avatar */
.panel-left {
  width: 220px;
  flex-shrink: 0;
  padding: 32px 24px;
  border-right: 1.5px solid #F0F0F0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar-upload-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
}

.avatar-upload {
  width: 100px; height: 100px;
  border-radius: 50%;
  background: #1B7A6E;
  display: flex; align-items: center; justify-content: center;
  overflow: hidden;
  border: 3px solid #E5E5E5;
  box-shadow: 0 4px 16px rgba(27,122,110,0.18);
}
.avatar-upload img { width: 100%; height: 100%; object-fit: cover; display: block; }

.avatar-upload__btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 7px 16px;
  background: #1B7A6E;
  color: #FFFFFF;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  border: none;
  transition: background 0.2s, transform 0.15s, box-shadow 0.2s;
  box-shadow: 0 2px 8px rgba(27,122,110,0.2);
}
.avatar-upload__btn:hover { background: #156358; transform: translateY(-1px); box-shadow: 0 4px 12px rgba(27,122,110,0.3); }

.avatar-hint {
  font-size: 11px;
  color: #9CA3AF;
  text-align: center;
  margin: 0;
  line-height: 1.4;
}

.avatar-badge {
  display: inline-block;
  background: #F0FAF8;
  color: #1B7A6E;
  border: 1px solid #B2DDD8;
  font-size: 11px;
  font-weight: 700;
  padding: 4px 12px;
  border-radius: 20px;
  text-transform: uppercase;
  letter-spacing: 0.8px;
}

/* Right: Form */
.panel-right {
  flex: 1;
  padding: 28px 28px 24px;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 18px;
  margin-bottom: 24px;
}

.form-group          { display: flex; flex-direction: column; gap: 6px; }
.form-group--full    { grid-column: 1 / -1; }

.form-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11.5px;
  font-weight: 700;
  color: #6B7280;
  text-transform: uppercase;
  letter-spacing: 0.6px;
}

.form-input {
  height: 42px;
  border: 1.5px solid #E5E7EB;
  border-radius: 10px;
  padding: 0 14px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 500;
  color: #1A1A1A;
  background: #FAFAFA;
  transition: border-color 0.2s, background 0.2s, box-shadow 0.2s;
  outline: none;
}
.form-input:focus {
  border-color: #1B7A6E;
  background: #FFFFFF;
  box-shadow: 0 0 0 3px rgba(27,122,110,0.1);
}
.form-input::placeholder { color: #C4C4C4; }

/* Input with @ prefix */
.input-prefix-wrap {
  position: relative;
  display: flex;
  align-items: center;
}
.input-prefix {
  position: absolute;
  left: 14px;
  font-size: 14px;
  font-weight: 700;
  color: #9CA3AF;
  pointer-events: none;
  font-family: 'Montserrat', sans-serif;
}
.form-input--prefix { padding-left: 28px; }

/* Live preview hint */
.form-hint {
  font-size: 11px;
  color: #9CA3AF;
  margin-top: 2px;
}
.form-hint b {
  color: #1B7A6E;
  font-weight: 700;
}

/* Action buttons */
.panel-actions {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 10px;
  padding-top: 4px;
  border-top: 1.5px solid #F0F0F0;
}

.btn-cancel {
  height: 40px;
  padding: 0 20px;
  border-radius: 10px;
  border: 1.5px solid #E5E7EB;
  background: #FFFFFF;
  font-size: 13px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  color: #6B7280;
  cursor: pointer;
  transition: all 0.15s;
}
.btn-cancel:hover { border-color: #9CA3AF; color: #374151; background: #F9F9F9; }

.btn-save {
  height: 40px;
  padding: 0 22px;
  border-radius: 10px;
  border: none;
  background: #1B7A6E;
  font-size: 13px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  color: #FFFFFF;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s, box-shadow 0.2s;
  box-shadow: 0 2px 8px rgba(27,122,110,0.25);
}
.btn-save:hover         { background: #156358; box-shadow: 0 4px 14px rgba(27,122,110,0.35); transform: translateY(-1px); }
.btn-save--saved        { background: #059669; }
.btn-save--saved:hover  { background: #047857; }

.btn-save__content {
  display: flex;
  align-items: center;
  gap: 7px;
}

/* Button text swap transition */
.btn-swap-enter-active, .btn-swap-leave-active { transition: opacity 0.15s ease, transform 0.15s ease; }
.btn-swap-enter-from  { opacity: 0; transform: scale(0.85); }
.btn-swap-leave-to    { opacity: 0; transform: scale(0.85); }

/* ── Stat Cards ── */
.stat-cards {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 32px;
}
.stat-card {
  background: #FFFFFF;
  border-radius: 14px;
  padding: 20px;
  display: flex; align-items: center; gap: 16px;
  box-shadow: 0 1px 8px rgba(0,0,0,0.06);
  animation: fadeUp 0.4s ease both;
  animation-delay: var(--delay, 0s);
}
@keyframes fadeUp {
  from { opacity: 0; transform: translateY(10px); }
  to   { opacity: 1; transform: translateY(0); }
}
.stat-card__icon-wrap {
  width: 50px; height: 50px;
  border-radius: 50%;
  background: #FFF3CD;
  display: flex; align-items: center; justify-content: center;
  flex-shrink: 0; color: #C8A135;
}
.stat-card__number { font-size: 26px; font-weight: 800; color: #1A1A1A; margin: 0 0 2px; line-height: 1; }
.stat-card__label  { font-size: 12px; color: #6B7280; margin: 0; font-weight: 500; }
.stat-card__label.bold { font-weight: 700; color: #1A1A1A; font-size: 13px; }
.stat-card__sub    { font-size: 11px; color: #9CA3AF; margin: 3px 0 0; }

/* ── Section Header ── */
.section-header { display: flex; align-items: center; justify-content: space-between; margin-bottom: 14px; }
.section-header.mt { margin-top: 28px; }
.section-title  { font-size: 16px; font-weight: 700; color: #1A1A1A; margin: 0; }
.section-viewall { display: flex; align-items: center; gap: 4px; font-size: 13px; font-weight: 600; color: #6B7280; text-decoration: none; transition: color 0.2s; }
.section-viewall:hover { color: #1B7A6E; }

/* ── Chart ── */
.chart-card { background: #FFFFFF; border-radius: 14px; padding: 24px 24px 16px; box-shadow: 0 1px 8px rgba(0,0,0,0.06); }
.chart-card__header { display: flex; align-items: center; justify-content: space-between; margin-bottom: 16px; }
.chart-card__title  { font-size: 15px; font-weight: 700; color: #1A1A1A; margin: 0; }
.chart-card__filter { display: flex; align-items: center; gap: 6px; background: #F0FAF8; border: 1px solid #B2DDD8; color: #1B7A6E; font-size: 13px; font-weight: 600; padding: 7px 14px; border-radius: 8px; cursor: pointer; font-family: 'Montserrat', sans-serif; }
.chart-wrap { width: 100%; }
.chart-svg  { width: 100%; height: 280px; display: block; }

/* ── Table ── */
.table-card { background: #FFFFFF; border-radius: 14px; box-shadow: 0 1px 8px rgba(0,0,0,0.06); overflow: hidden; }
.table      { width: 100%; border-collapse: collapse; font-size: 13.5px; }
.table thead tr { border-bottom: 1.5px solid #F0F0F0; }
.table th { padding: 14px 18px; text-align: left; font-size: 12px; font-weight: 600; color: #9CA3AF; white-space: nowrap; font-family: 'Montserrat', sans-serif; }
.table tbody tr { border-bottom: 1px solid #F7F7F7; transition: background 0.15s; }
.table tbody tr:last-child { border-bottom: none; }
.table tbody tr:hover { background: #FAFAFA; }
.table td { padding: 13px 18px; vertical-align: middle; }
.td-no    { color: #9CA3AF; font-weight: 500; width: 40px; }

.client-row { display: flex; align-items: center; gap: 12px; }
.avatar-table { width: 34px; height: 34px; border-radius: 50%; display: flex; align-items: center; justify-content: center; font-size: 11px; font-weight: 700; color: #FFFFFF; flex-shrink: 0; font-family: 'Montserrat', sans-serif; }
.client-name    { font-size: 13px; font-weight: 600; color: #1A1A1A; margin: 0 0 2px; }
.client-project { font-size: 11.5px; color: #9CA3AF; margin: 0; }

.badge { display: inline-flex; align-items: center; gap: 5px; padding: 4px 11px; border-radius: 20px; font-size: 12px; font-weight: 600; font-family: 'Montserrat', sans-serif; }
.badge::before { content: ''; width: 6px; height: 6px; border-radius: 50%; }
.badge--interior   { background: #ECFDF5; color: #1B7A6E; }
.badge--interior::before   { background: #1B7A6E; }
.badge--kontraktor { background: #F3F4F6; color: #4B5563; }
.badge--kontraktor::before { background: #9CA3AF; }
.badge--custom     { background: #FFFBEB; color: #B45309; }
.badge--custom::before     { background: #F59E0B; }

.td-biaya   { font-weight: 600; color: #1A1A1A; white-space: nowrap; }
.td-tanggal { color: #6B7280; white-space: nowrap; }
.td-actions { display: flex; align-items: center; gap: 8px; }

.action-btn {
  width: 30px; height: 30px; border-radius: 7px;
  border: 1px solid #E5E7EB; background: transparent;
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: #9CA3AF; transition: all 0.15s;
}
.action-btn:hover      { border-color: #1B7A6E; color: #1B7A6E; background: #F0FAF8; }
.action-btn--del:hover { border-color: #EF4444; color: #EF4444; background: #FEF2F2; }
</style>