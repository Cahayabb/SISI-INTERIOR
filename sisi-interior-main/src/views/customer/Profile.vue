<template>
  <div class="profile-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">Profile</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">Profile</span>
        </div>
      </div>
    </div>

    <!-- ── Main Content ── -->
    <section class="profile-section">
      <div class="container">

        <!-- Loading State -->
        <div class="state-loading" v-if="isLoading">
          <div class="spinner"></div>
          <p>Memuat data profil...</p>
        </div>

        <template v-else>

          <!-- ── Profile Card ── -->
          <div class="profile-card">

            <!-- Avatar Row -->
            <div class="profile-card__top">
              <div class="avatar-wrap">
                <div class="avatar-circle">
                  <img v-if="avatarPreview" :src="avatarPreview" alt="Foto Profil" @error="avatarPreview = ''" />
                  <span v-else class="avatar-initials">{{ userInitials }}</span>
                </div>
                <!-- Kamera di luar lingkaran -->
                <label v-if="isEditing" class="avatar-camera-btn" title="Ganti foto">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                    <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                    <circle cx="12" cy="13" r="4"/>
                  </svg>
                  Ganti Foto
                  <input type="file" accept="image/*" style="display:none" @change="onAvatarChange" />
                </label>
              </div>

              <div class="profile-card__meta">
                <p class="profile-card__name">{{ form.nama || '—' }}</p>
                <span class="profile-card__badge">Customer</span>
              </div>

              <!-- Edit / Simpan / Batal -->
              <div class="profile-card__actions">
                <template v-if="!isEditing">
                  <button class="btn-edit" @click="startEdit">
                    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                      <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                      <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                    </svg>
                    Edit
                  </button>
                </template>
                <template v-else>
                  <button class="btn-cancel-edit" @click="cancelEdit">Batal</button>
                  <button class="btn-save" @click="handleSimpan" :disabled="isSaving">
                    <span v-if="isSaving" class="btn-spinner"></span>
                    <template v-else>
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                        <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/>
                        <polyline points="17 21 17 13 7 13 7 21"/>
                        <polyline points="7 3 7 8 15 8"/>
                      </svg>
                      Simpan
                    </template>
                  </button>
                </template>
              </div>
            </div>

            <!-- Toast notification -->
            <Transition name="toast">
              <div class="toast" :class="'toast--' + toast.type" v-if="toast.show">
                <svg v-if="toast.type === 'success'" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
                <svg v-else width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
                {{ toast.message }}
              </div>
            </Transition>

            <!-- ── Form Fields ── -->
            <div class="profile-form">
              <div class="form-grid">

                <!-- Nama -->
                <div class="form-group">
                  <label class="form-label">Nama</label>
                  <div class="form-input-wrap">
                    <input
                      class="form-input"
                      :class="{ 'form-input--active': isEditing }"
                      type="text"
                      v-model="form.nama"
                      placeholder="Nama lengkap"
                      :readonly="!isEditing"
                    />
                    <div class="form-input-bar" v-if="isEditing"></div>
                  </div>
                </div>

                <!-- Email -->
                <div class="form-group">
                  <label class="form-label">Email</label>
                  <div class="form-input-wrap">
                    <input
                      class="form-input"
                      :class="{ 'form-input--active': isEditing }"
                      type="email"
                      v-model="form.email"
                      placeholder="Email"
                      :readonly="!isEditing"
                    />
                    <div class="form-input-bar" v-if="isEditing"></div>
                  </div>
                </div>

                <!-- ID Pengguna -->
                <div class="form-group">
                  <label class="form-label">ID Pengguna</label>
                  <div class="form-input-wrap">
                    <input
                      class="form-input form-input--readonly"
                      type="text"
                      :value="form.id"
                      placeholder="ID Pengguna"
                      readonly
                    />
                    <span class="form-input-lock">
                      <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                        <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                        <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
                      </svg>
                    </span>
                  </div>
                </div>

                <!-- Nomor Telepon -->
                <div class="form-group">
                  <label class="form-label">Nomor Telepon</label>
                  <div class="form-input-wrap">
                    <input
                      class="form-input"
                      :class="{ 'form-input--active': isEditing }"
                      type="tel"
                      v-model="form.noTelepon"
                      placeholder="Nomor telepon"
                      :readonly="!isEditing"
                    />
                    <div class="form-input-bar" v-if="isEditing"></div>
                  </div>
                </div>

              </div>
            </div>
          </div>

          <!-- ── Riwayat Estimasi ── -->
          <div class="history-section">
            <div class="history-header">
              <h2 class="history-title">Riwayat perhitungan estimasi</h2>
              <span class="history-count" v-if="riwayat.length">{{ riwayat.length }} proyek</span>
            </div>

            <!-- Loading riwayat -->
            <div class="state-loading state-loading--sm" v-if="isLoadingRiwayat">
              <div class="spinner spinner--sm"></div>
              <p>Memuat riwayat...</p>
            </div>

            <!-- Empty state -->
            <div class="history-empty" v-else-if="!riwayat.length">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#CCCCCC" stroke-width="1.5">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
              </svg>
              <p>Belum ada riwayat estimasi.</p>
            </div>

            <!-- List -->
            <div class="history-list" v-else>
              <div
                class="history-item"
                v-for="(item, i) in riwayat"
                :key="item.id"
                :style="{ '--delay': i * 0.06 + 's' }"
              >
                <div class="history-item__left">
                  <div class="history-item__icon">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                      <polyline points="14 2 14 8 20 8"/>
                    </svg>
                  </div>
                  <div>
                    <p class="history-item__id">{{ item.id }}</p>
                    <p class="history-item__name">{{ item.namaProyek }}</p>
                  </div>
                </div>

                <div class="history-item__right">
                  <span class="history-item__biaya" v-if="item.estimasi">{{ item.estimasi }}</span>
                  <span class="history-item__date">{{ formatDate(item.tanggal) }}</span>
                  <button
                    class="btn-hapus"
                    @click="confirmHapus(item)"
                    :disabled="isDeletingId === item.id"
                  >
                    <span v-if="isDeletingId === item.id" class="btn-spinner btn-spinner--sm btn-spinner--red"></span>
                    <template v-else>Hapus</template>
                  </button>
                </div>
              </div>
            </div>
          </div>

        </template>

      </div>
    </section>

    <!-- ── Confirm Delete Modal ── -->
    <Transition name="modal">
      <div class="modal-overlay" v-if="confirmDelete.show" @click.self="confirmDelete.show = false">
        <div class="modal-box">
          <div class="modal-box__icon">
            <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="#EF4444" stroke-width="2">
              <polyline points="3 6 5 6 21 6"/>
              <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
              <path d="M10 11v6"/><path d="M14 11v6"/>
              <path d="M9 6V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2"/>
            </svg>
          </div>
          <h3 class="modal-box__title">Hapus Riwayat?</h3>
          <p class="modal-box__desc">
            Riwayat estimasi <strong>{{ confirmDelete.item?.namaProyek }}</strong>
            akan dihapus permanen dan tidak dapat dikembalikan.
          </p>
          <div class="modal-box__actions">
            <button class="modal-btn modal-btn--cancel" @click="confirmDelete.show = false">Batal</button>
            <button class="modal-btn modal-btn--delete" @click="handleHapus">Ya, Hapus</button>
          </div>
        </div>
      </div>
    </Transition>

  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { RouterLink } from 'vue-router'

// ─────────────────────────────────────────
// STATE
// ─────────────────────────────────────────
const isLoading        = ref(true)
const isLoadingRiwayat = ref(true)
const isSaving         = ref(false)
const isDeletingId     = ref(null)
const isEditing        = ref(false)
const avatarPreview    = ref('')
const avatarFile       = ref(null)   // File object untuk upload ke server

const toast = reactive({ show: false, type: 'success', message: '' })
const confirmDelete = reactive({ show: false, item: null })

// Snapshot data sebelum edit (untuk cancel)
let formSnapshot = {}

// Data form user
const form = reactive({
  id:         '',
  nama:       '',
  email:      '',
  noTelepon:  '',
})

// Riwayat estimasi
const riwayat = ref([])

// ─────────────────────────────────────────
// COMPUTED
// ─────────────────────────────────────────
const userInitials = computed(() => {
  const parts = (form.nama || '').trim().split(' ')
  const a = parts[0]?.[0] || ''
  const b = parts[1]?.[0] || ''
  return (a + b).toUpperCase() || 'U'
})

// ─────────────────────────────────────────
// API HELPERS 
// ─────────────────────────────────────────
const BASE_URL = 'http://localhost:8080'

const getHeaders = () => ({
  'Content-Type': 'application/json',
  'Authorization': `Bearer ${localStorage.getItem('token')}`,
})

// ─────────────────────────────────────────
// FETCH PROFILE
// GET /user/profile
// Response: { id, nama, email, no_telepon, foto_url? }
// ─────────────────────────────────────────
const fetchProfile = async () => {
  try {
    const res  = await fetch(`${BASE_URL}/user/profile`, { headers: getHeaders() })
    const data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal memuat profil.')

    form.id        = data.id        || ''
    form.nama      = data.nama      || data.name  || ''
    form.email     = data.email     || ''
    form.noTelepon = data.no_telepon || data.phone || ''

    if (data.foto_url) {
      avatarPreview.value = data.foto_url
    } else {
      // Fallback ke localStorage jika ada
      const local = localStorage.getItem('user_avatar')
      if (local) avatarPreview.value = local
    }

    // Sync ke localStorage untuk dipakai Navbar
    localStorage.setItem('user', JSON.stringify({
      first_name: form.nama.split(' ')[0] || '',
      last_name:  form.nama.split(' ').slice(1).join(' ') || '',
      email:      form.email,
      username:   data.username || '',
    }))
  } catch (err) {
    // Fallback: baca dari localStorage jika server belum ready
    const local = JSON.parse(localStorage.getItem('user') || '{}')
    form.nama      = (local.first_name + ' ' + (local.last_name || '')).trim() || ''
    form.email     = local.email    || ''
    form.id        = local.id       || local.username || ''
    form.noTelepon = local.phone    || ''
    const savedAvatar = localStorage.getItem('user_avatar')
    if (savedAvatar) avatarPreview.value = savedAvatar
  } finally {
    isLoading.value = false
  }
}

// ─────────────────────────────────────────
// FETCH RIWAYAT ESTIMASI
// GET /estimasi/riwayat
// Response: [{ id, nama_proyek, estimasi, tanggal }]
// ─────────────────────────────────────────
const fetchRiwayat = async () => {
  isLoadingRiwayat.value = true
  try {
    const res  = await fetch(`${BASE_URL}/estimasi/riwayat`, { headers: getHeaders() })
    const data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal memuat riwayat.')

    riwayat.value = (data || []).map(item => ({
      id:         item.id,
      namaProyek: item.nama_proyek || item.namaProyek || '',
      estimasi:   item.estimasi    || '',
      tanggal:    item.tanggal     || item.created_at || '',
    }))
  } catch {
    // Tampilkan data dummy saat dev/offline
    riwayat.value = [
      { id: '1222873717', namaProyek: 'Kitchen set project', estimasi: 'Rp 18.000.000', tanggal: '2025-05-12' },
      { id: '1222873718', namaProyek: 'Library project',     estimasi: 'Rp 32.000.000', tanggal: '2025-06-03' },
    ]
  } finally {
    isLoadingRiwayat.value = false
  }
}

// ─────────────────────────────────────────
// EDIT PROFILE
// ─────────────────────────────────────────
const startEdit = () => {
  formSnapshot = { ...form }
  isEditing.value = true
}

const cancelEdit = () => {
  Object.assign(form, formSnapshot)
  avatarFile.value    = null
  // Kembalikan preview ke state sebelum edit
  const saved = localStorage.getItem('user_avatar')
  avatarPreview.value = saved || ''
  isEditing.value = false
}

const onAvatarChange = (e) => {
  const file = e.target.files[0]
  if (!file) return
  avatarFile.value = file
  const reader = new FileReader()
  reader.onload = (ev) => { avatarPreview.value = ev.target.result }
  reader.readAsDataURL(file)
}

// ─────────────────────────────────────────
// SIMPAN PROFILE
// PUT /user/profile
// Body: { nama, email, no_telepon } + optional multipart foto
// Response: { message, user: {...} }
// ─────────────────────────────────────────
const handleSimpan = async () => {
  if (!form.nama.trim()) {
    showToast('error', 'Nama tidak boleh kosong.')
    return
  }

  isSaving.value = true
  try {
    let res, data

    if (avatarFile.value) {
      // Kirim sebagai multipart/form-data jika ada foto
      const fd = new FormData()
      fd.append('nama',        form.nama)
      fd.append('email',       form.email)
      fd.append('no_telepon',  form.noTelepon)
      fd.append('foto',        avatarFile.value)

      res  = await fetch(`${BASE_URL}/user/profile`, {
        method:  'PUT',
        headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` },
        body:    fd,
      })
    } else {
      // JSON biasa tanpa foto
      res  = await fetch(`${BASE_URL}/user/profile`, {
        method:  'PUT',
        headers: getHeaders(),
        body:    JSON.stringify({
          nama:        form.nama,
          email:       form.email,
          no_telepon:  form.noTelepon,
        }),
      })
    }

    data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal menyimpan profil.')

    // Update localStorage Navbar
    const namaParts = form.nama.trim().split(' ')
    localStorage.setItem('user', JSON.stringify({
      first_name: namaParts[0] || '',
      last_name:  namaParts.slice(1).join(' ') || '',
      email:      form.email,
    }))

    // Simpan avatar lokal jika ada
    if (avatarFile.value && avatarPreview.value) {
      localStorage.setItem('user_avatar', avatarPreview.value)
    }
    // Update foto_url dari response server
    if (data.user?.foto_url) {
      avatarPreview.value = data.user.foto_url
    }

    avatarFile.value = null
    isEditing.value  = false
    showToast('success', 'Profil berhasil disimpan!')
  } catch (err) {
    showToast('error', err.message || 'Tidak dapat terhubung ke server.')
  } finally {
    isSaving.value = false
  }
}

// ─────────────────────────────────────────
// HAPUS RIWAYAT ESTIMASI
// DELETE /estimasi/riwayat/:id
// Response: { message }
// ─────────────────────────────────────────
const confirmHapus = (item) => {
  confirmDelete.item = item
  confirmDelete.show = true
}

const handleHapus = async () => {
  const item = confirmDelete.item
  if (!item) return
  confirmDelete.show = false
  isDeletingId.value = item.id

  try {
    const res  = await fetch(`${BASE_URL}/estimasi/riwayat/${item.id}`, {
      method:  'DELETE',
      headers: getHeaders(),
    })
    const data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal menghapus.')

    riwayat.value = riwayat.value.filter(r => r.id !== item.id)
    showToast('success', 'Riwayat estimasi berhasil dihapus.')
  } catch (err) {
    showToast('error', err.message || 'Gagal menghapus riwayat.')
  } finally {
    isDeletingId.value = null
  }
}

// ─────────────────────────────────────────
// HELPERS
// ─────────────────────────────────────────
let toastTimer = null
const showToast = (type, message) => {
  toast.type    = type
  toast.message = message
  toast.show    = true
  clearTimeout(toastTimer)
  toastTimer = setTimeout(() => { toast.show = false }, 3500)
}

const formatDate = (dateStr) => {
  if (!dateStr) return '—'
  try {
    return new Date(dateStr).toLocaleDateString('id-ID', {
      day: 'numeric', month: 'long', year: 'numeric'
    })
  } catch { return dateStr }
}

// ─────────────────────────────────────────
// LIFECYCLE
// ─────────────────────────────────────────
onMounted(() => {
  fetchProfile()
  fetchRiwayat()
})
</script>

<style scoped>
.profile-page {
  padding-top: 72px;
  font-family: 'Montserrat', sans-serif;
  background: #FAFAFA;
  min-height: 100vh;
}

.container {
  max-width: 900px;
  margin: 0 auto;
  padding: 0 40px;
}

/* ── Page Header ── */
.page-header {
  padding: 36px 0 0;
  background: #FFFFFF;
  border-bottom: 1px solid #F0F0F0;
}

.page-header__title {
  font-size: 32px;
  font-weight: 800;
  color: #2C2C2C;
  margin: 0 0 8px;
}

.page-header__breadcrumb {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #888;
  padding-bottom: 24px;
}

.breadcrumb__link { color: #888; text-decoration: none; transition: color 0.2s; }
.breadcrumb__link:hover { color: #1B7A6E; }
.breadcrumb__sep  { color: #bbb; }
.breadcrumb__current { color: #555; }

/* ── Section ── */
.profile-section {
  padding: 40px 0 80px;
}

/* ── Loading ── */
.state-loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
  padding: 60px 0;
  color: #9CA3AF;
  font-size: 13px;
}

.state-loading--sm { padding: 24px 0; }

.spinner {
  width: 36px; height: 36px;
  border: 3px solid #E5E7EB;
  border-top-color: #1B7A6E;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

.spinner--sm { width: 22px; height: 22px; border-width: 2px; }

@keyframes spin { to { transform: rotate(360deg); } }

/* ── Profile Card ── */
.profile-card {
  background: #FFFFFF;
  border-radius: 14px;
  border: 1.5px solid #EEEEEE;
  padding: 32px 36px 36px;
  margin-bottom: 28px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.05);
  position: relative;
}

/* ── Top Row: Avatar + Name + Buttons ── */
.profile-card__top {
  display: flex;
  align-items: flex-end;
  gap: 20px;
  margin-bottom: 36px;
  flex-wrap: wrap;
}

/* Avatar */
.avatar-wrap {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  flex-shrink: 0;
}

.avatar-circle {
  width: 88px;
  height: 88px;
  border-radius: 50%;
  background: #1B7A6E;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 3px solid #E0F0EE;
  box-shadow: 0 4px 16px rgba(27,122,110,0.15);
}

.avatar-circle img {
  width: 100%; height: 100%;
  object-fit: cover; display: block;
}

.avatar-initials {
  font-size: 28px;
  font-weight: 700;
  color: #FFFFFF;
  line-height: 1;
  user-select: none;
}

/* Camera button — outside circle */
.avatar-camera-btn {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 6px 12px;
  background: #1B7A6E;
  color: #FFFFFF;
  border-radius: 20px;
  font-size: 11.5px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s;
  white-space: nowrap;
  box-shadow: 0 2px 8px rgba(27,122,110,0.2);
}

.avatar-camera-btn:hover { background: #156358; transform: translateY(-1px); }

/* Name + badge */
.profile-card__meta {
  flex: 1;
}

.profile-card__name {
  font-size: 22px;
  font-weight: 800;
  color: #1A1A1A;
  margin: 0 0 8px;
  line-height: 1.2;
}

.profile-card__badge {
  display: inline-block;
  background: #F0FAF8;
  color: #1B7A6E;
  border: 1px solid #B2DDD8;
  font-size: 11px;
  font-weight: 700;
  padding: 3px 12px;
  border-radius: 20px;
  text-transform: uppercase;
  letter-spacing: 0.8px;
}

/* Action buttons */
.profile-card__actions {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-left: auto;
  flex-shrink: 0;
}

.btn-edit {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  padding: 11px 24px;
  background: #C8A135;
  color: #FFFFFF;
  border: none;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s, box-shadow 0.2s;
  box-shadow: 0 2px 8px rgba(200,161,53,0.25);
}

.btn-edit:hover { background: #b08e2c; transform: translateY(-1px); box-shadow: 0 4px 14px rgba(200,161,53,0.35); }

.btn-cancel-edit {
  padding: 10px 20px;
  background: #FFFFFF;
  border: 1.5px solid #E5E7EB;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  color: #6B7280;
  cursor: pointer;
  transition: all 0.15s;
}

.btn-cancel-edit:hover { border-color: #9CA3AF; color: #374151; }

.btn-save {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  padding: 10px 22px;
  background: #1B7A6E;
  color: #FFFFFF;
  border: none;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s;
  box-shadow: 0 2px 8px rgba(27,122,110,0.25);
}

.btn-save:hover:not(:disabled) { background: #156358; transform: translateY(-1px); }
.btn-save:disabled { opacity: 0.65; cursor: not-allowed; }

/* ── Toast ── */
.toast {
  position: absolute;
  top: 20px; right: 20px;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 11px 18px;
  border-radius: 10px;
  font-size: 13px;
  font-weight: 600;
  z-index: 10;
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
}

.toast--success { background: #ECFDF5; color: #1B7A6E; border: 1px solid #A7F3D0; }
.toast--error   { background: #FEF2F2; color: #EF4444; border: 1px solid #FECACA; }

.toast-enter-active, .toast-leave-active { transition: opacity 0.25s, transform 0.25s; }
.toast-enter-from, .toast-leave-to       { opacity: 0; transform: translateY(-8px); }

/* ── Form Grid ── */
.profile-form {
  border-top: 1.5px solid #F0F0F0;
  padding-top: 28px;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px 32px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 12px;
  font-weight: 700;
  color: #9CA3AF;
  text-transform: uppercase;
  letter-spacing: 0.7px;
}

/* Input wrapper for underline bar */
.form-input-wrap {
  position: relative;
}

.form-input {
  width: 100%;
  border: 1.5px solid #E5E7EB;
  border-radius: 8px;
  padding: 13px 16px;
  font-size: 14px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 500;
  color: #1A1A1A;
  background: #F9F9F9;
  outline: none;
  transition: border-color 0.2s, background 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}

.form-input::placeholder { color: #C4C4C4; }

.form-input--active {
  background: #FFFFFF;
  border-color: #C8A135;
  box-shadow: 0 0 0 3px rgba(200,161,53,0.12);
}

.form-input--active:focus {
  border-color: #C8A135;
}

/* Readonly lock */
.form-input--readonly {
  background: #F3F4F6;
  color: #9CA3AF;
  cursor: not-allowed;
  padding-right: 40px;
}

.form-input-lock {
  position: absolute;
  right: 14px;
  top: 50%;
  transform: translateY(-50%);
  color: #C4C4C4;
  pointer-events: none;
}

/* ── Riwayat Estimasi ── */
.history-section {
  background: #FFFFFF;
  border-radius: 14px;
  border: 1.5px solid #EEEEEE;
  padding: 28px 36px 32px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.05);
}

.history-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.history-title {
  font-size: 16px;
  font-weight: 800;
  color: #1A1A1A;
  margin: 0;
}

.history-count {
  font-size: 12px;
  font-weight: 600;
  color: #9CA3AF;
  background: #F3F4F6;
  padding: 3px 10px;
  border-radius: 20px;
}

/* Empty */
.history-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  padding: 40px 0;
  color: #C4C4C4;
  font-size: 13px;
}

/* History list */
.history-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.history-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  background: #F8F8F8;
  border-radius: 10px;
  padding: 16px 20px;
  border: 1px solid #EEEEEE;
  animation: fadeUp 0.35s ease both;
  animation-delay: var(--delay, 0s);
  transition: background 0.15s, box-shadow 0.15s;
}

.history-item:hover {
  background: #F0FAF8;
  border-color: #C2E5E0;
  box-shadow: 0 2px 8px rgba(27,122,110,0.07);
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(8px); }
  to   { opacity: 1; transform: translateY(0); }
}

.history-item__left {
  display: flex;
  align-items: center;
  gap: 14px;
  flex: 1;
  min-width: 0;
}

.history-item__icon {
  width: 36px; height: 36px;
  border-radius: 8px;
  background: #EFF6FF;
  display: flex; align-items: center; justify-content: center;
  color: #3B82F6;
  flex-shrink: 0;
}

.history-item__id {
  font-size: 11px;
  font-weight: 700;
  color: #9CA3AF;
  margin: 0 0 3px;
  letter-spacing: 0.3px;
}

.history-item__name {
  font-size: 14px;
  font-weight: 600;
  color: #1A1A1A;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.history-item__right {
  display: flex;
  align-items: center;
  gap: 16px;
  flex-shrink: 0;
}

.history-item__biaya {
  font-size: 13px;
  font-weight: 700;
  color: #1B7A6E;
  white-space: nowrap;
}

.history-item__date {
  font-size: 12px;
  color: #9CA3AF;
  white-space: nowrap;
}

/* Hapus button */
.btn-hapus {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 68px;
  padding: 7px 16px;
  background: #EF4444;
  color: #FFFFFF;
  border: none;
  border-radius: 6px;
  font-size: 12.5px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s;
}

.btn-hapus:hover:not(:disabled) { background: #DC2626; transform: translateY(-1px); }
.btn-hapus:disabled { opacity: 0.6; cursor: not-allowed; }

/* ── Spinner ── */
.btn-spinner {
  width: 16px; height: 16px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
  display: inline-block;
}

.btn-spinner--sm  { width: 13px; height: 13px; border-width: 2px; }
.btn-spinner--red { border-color: rgba(255,255,255,0.4); border-top-color: #fff; }

/* ── Confirm Modal ── */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 300;
  padding: 20px;
}

.modal-box {
  background: #FFFFFF;
  border-radius: 14px;
  padding: 36px 32px;
  width: 100%;
  max-width: 380px;
  box-shadow: 0 20px 60px rgba(0,0,0,0.18);
  text-align: center;
}

.modal-box__icon {
  width: 60px; height: 60px;
  background: #FEF2F2;
  border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
  margin: 0 auto 16px;
}

.modal-box__title {
  font-size: 18px;
  font-weight: 800;
  color: #1A1A1A;
  margin: 0 0 10px;
}

.modal-box__desc {
  font-size: 13.5px;
  color: #555;
  line-height: 1.7;
  margin: 0 0 24px;
}

.modal-box__actions {
  display: flex;
  gap: 10px;
}

.modal-btn {
  flex: 1;
  padding: 13px;
  border-radius: 8px;
  font-size: 13.5px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  border: none;
  transition: all 0.15s;
}

.modal-btn--cancel { background: #F3F4F6; color: #374151; }
.modal-btn--cancel:hover { background: #E5E7EB; }

.modal-btn--delete { background: #EF4444; color: #FFFFFF; }
.modal-btn--delete:hover { background: #DC2626; }

.modal-enter-active, .modal-leave-active { transition: opacity 0.22s, transform 0.22s; }
.modal-enter-from, .modal-leave-to       { opacity: 0; transform: scale(0.95); }

/* ── Responsive ── */
@media (max-width: 640px) {
  .container { padding: 0 20px; }
  .profile-card, .history-section { padding: 24px 20px; border-radius: 10px; }
  .profile-card__top { flex-direction: column; align-items: flex-start; }
  .profile-card__actions { margin-left: 0; width: 100%; justify-content: flex-end; }
  .form-grid { grid-template-columns: 1fr; }
  .history-item { flex-direction: column; align-items: flex-start; }
  .history-item__right { width: 100%; justify-content: space-between; }
}
</style>