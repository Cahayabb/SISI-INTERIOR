<template>
<div class="estimasi-page" :key="`view-${currentView}`">
    <!-- ── HEADER ── -->
    <div class="page-header">
      <h1 class="page-title">Estimasi Biaya Proyek</h1>
      <div class="header-right">
        <div class="search-box">
          <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
          </svg>
          <input type="text" v-model="searchQuery" placeholder="Search" class="search-input" />
        </div>
        <button class="notif-btn">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"/>
            <path d="M13.73 21a2 2 0 0 1-3.46 0"/>
          </svg>
          <span class="notif-dot"></span>
        </button>
        <div class="avatar">
          <img src="" alt="User" @error="onAvatarError" ref="avatarImg" />
          <div class="avatar-fallback" ref="avatarFallback">A</div>
      </div>
    </div>
     </div>

    <!-- ══════════════════════════════════════
         VIEW 1: LIST DATA ESTIMASI TERSIMPAN
    ══════════════════════════════════════ -->
    <template v-if="currentView === 'list'">

      <div class="toolbar">
        <button class="btn-input-variabel" @click="openForm">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          Input Variabel Proyek
        </button>
        <div class="toolbar-right">
          <div class="ctrl-wrap">
            <div class="select-wrap">
              <select v-model="sortBy" class="ctrl-select">
                <option value="date_desc">Sort by: Date Created</option>
                <option value="date_asc">Sort by: Date Terlama</option>
                <option value="nama_asc">Sort by: Nama A–Z</option>
              </select>
              <svg class="select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>
          <button class="btn-filter" @click="showFilter = !showFilter">
            Filter
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"/></svg>
          </button>
          <div v-if="showFilter" class="filter-box">
          <select v-model="filterJenis">
            <option value="">Semua Jenis</option>
            <option>Custom Furniture</option>
            <option>Desain Interior</option>
            <option>Renovasi</option>
          </select>
        </div>
        </div>
      </div>

      <div class="table-card">
        <div class="table-card__title">Data Estimasi Tersimpan</div>
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Jenis Ruangan</th>
              <th>Luas Area</th>
              <th>Harga Proyek</th>
              <th>Spesifikasi Design</th>
              <th>Tanggal</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="filteredList.length === 0">
              <td colspan="7" class="empty-row">
                <div class="empty-state">
                  <svg width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="#D0D0D0" stroke-width="1.5">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                    <polyline points="14 2 14 8 20 8"/>
                  </svg>
                  <p>Belum ada data estimasi.</p>
                </div>
              </td>
            </tr>
            <tr v-for="(item, index) in paginatedList" :key="item.id">
              <td class="td-no">{{ (currentPage - 1) * perPage + index + 1 }}</td>
              <td class="td-main">{{ item.jenis_ruangan}}</td>
              <td class="td-cell">
                <span class="badge" :class="badgeClass(item.jenis_proyek)">{{ item.luas_area }}</span>
              </td>
              <td class="td-cell">{{ item.harga_proyek }}</td>
              <td class="td-cell td-file">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
                {{ item.spesifikasi_design}}
              </td>
              <td class="td-cell td-muted">{{ formatTanggal(item.created_at) }}</td>
              <td class="td-aksi">
                <button class="action-btn action-btn--edit" @click="viewHasil(item)" title="Lihat Hasil">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                    <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                  </svg>
                </button>
                <button class="action-btn action-btn--delete" @click="deleteItem(item.id)" title="Hapus">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2">
                    <polyline points="3 6 5 6 21 6"/>
                    <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
                    <path d="M10 11v6M14 11v6"/>
                    <path d="M9 6V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2"/>
                  </svg>
                </button>
              </td>
            </tr>
          </tbody>
        </table>

        <!-- Pagination -->
        <div class="pagination-bar">
          <span class="pagination-info">
            Showing 
            <strong>
              {{ (currentPage - 1) * perPage + 1 }}
              -
              {{ Math.min(currentPage * perPage, filteredList.length) }}
            </strong>
            from <strong>{{ filteredList.length }}</strong> data
          </span>
          <div class="pagination-pages">
            <button class="page-btn" :disabled="currentPage === 1" @click="currentPage--">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="15 18 9 12 15 6"/></svg>
            </button>
            <button v-for="p in totalPages" :key="p" class="page-btn" :class="{ active: p === currentPage }" @click="currentPage = p">{{ p }}</button>
            <button class="page-btn" :disabled="currentPage === totalPages" @click="currentPage++">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"/></svg>
            </button>
          </div>
        </div>
      </div>

    </template>

    <!-- ══════════════════════════════════════
         VIEW 2: FORM INPUT VARIABEL PROYEK
    ══════════════════════════════════════ -->
    <template v-if="currentView === 'form'">

      <div class="view-heading">
        <button class="back-btn" @click="currentView = 'list'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <h2 class="view-title">Input Estimasi Biaya Proyek</h2>
      </div>

      <div class="form-layout">

        <!-- Form kiri -->
        <div class="form-card">
          <div class="form-grid-2">
            <div class="field">
              <label class="field__label">Nama Perusahaan</label>
              <input v-model="form.nama_perusahaan" type="text" class="field__input" placeholder="" />
            </div>
            <div class="field">
              <label class="field__label">Luas Area (m²)</label>
              <input v-model="form.luas_area" type="text" class="field__input" placeholder="" />
            </div>
          </div>

          <div class="field">
            <label class="field__label">Jenis Proyek</label>
            <div class="field__select-wrap">
              <select v-model="form.jenis_proyek" class="field__select">
                <option value="" disabled>Pilih jenis proyek</option>
                <option>Rumah</option>
                <option>Apartment</option>
                <option>Cafe</option>
                <option>Kantor</option>
                <option>Clinic</option>
                <option>Tenant</option>
                <option>Kost</option>
                <option>SPA</option>
                <option value="lainnya">Lainnya</option>
              </select>


              <input
                v-if="form.jenis_proyek === 'lainnya'"
                v-model="form.jenis_proyek_lainnya"
                type="text"
                class="field__input"
                placeholder="Masukkan jenis proyek lainnya"
              />
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Jenis Pekerjaan</label>
            <select v-model="form.jenis_pekerjaan" class="field__select">
              <option value="" disabled>Pilih jenis pekerjaan</option>
              <option>Design</option>
              <option>Build / Fit Out</option>
              <option>Furniture</option>
              <option>MEP (Mechanical Electrical Plumbing)</option>
              <option>Maintenance</option>
            </select>
          </div>
          
          <div class="field">
            <label class="field__label">Jenis Ruangan</label>

            <div class="multiselect" ref="multiRef">
              
              <!-- Trigger -->
              <div class="multiselect__trigger" @click="toggleMultiSelect">
                <span v-if="form.jenis_ruangan.length === 0" class="multiselect__placeholder">
                  Silahkan pilih beberapa jenis
                </span>

                <span v-else class="multiselect__value">
                  {{ form.jenis_ruangan.join(', ') }}
                </span>

                <svg class="multiselect__arrow" width="16" height="16"
                  :style="{ transform: multiSelectOpen ? 'rotate(180deg)' : 'rotate(0deg)' }">
                  <polyline points="6 9 12 15 18 9"/>
                </svg>
              </div>

              <!-- Dropdown -->
              <div v-if="multiSelectOpen" class="multiselect__dropdown">
                <label 
                  v-for="item in daftarItemOptions" 
                  :key="item" 
                  class="multiselect__option"
                >
                  <input
                    type="checkbox"
                    :value="item"
                    v-model="form.jenis_ruangan"
                  />
                  <span>{{ item }}</span>
                </label>
              </div>

            </div>
          </div>

          <div class="field">
            <label class="field__label">Tingkat Kerumitan</label>
            <div class="field__select-wrap">
              <select v-model="form.tingkat_kerumitan" class="field__select">
                <option value="" disabled>Pilih tingkat kerumitan</option>
                <option>Mudah</option>
                <option>Sedang</option>
                <option>Sulit</option>
              </select>
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Durasi Pengerjaan (hari)</label>
            <input v-model="form.durasi_pengerjaan" type="number" class="field__input" placeholder="" />
          </div>

          <div class="field">
            <label class="field__label">Lokasi Proyek</label>
            <input v-model="form.lokasi_proyek" type="text" class="field__input" placeholder="" />
          </div>

          <div class="field">
            <label class="field__label">Spesifikasi Design</label>
            <div class="field__select-wrap">
              <select v-model="form.spesifikasi_design" class="field__select">
                <option value="" disabled>Pilih spesifikasi design</option>
                <option>Japandi</option>
                <option>Industrial</option>
                <option>Scandinavian</option>
                <option>Modern</option>
                <option>Modern Minimalist</option>
                <option>Ekletik / Colorful</option>
                <option>Bohemian</option>
                <option>Klasik / Victorian</option>
                <option>Rustic</option>
                <option>Minimalis</option>
              </select>
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Material Khusus</label>
            <div class="field__select-wrap">
              <select v-model="form.material_khusus">
                <option value="">Opsional</option>
                <option>Kayu Jati</option>
                <option>Kayu Pinus</option>
                <option>Baja</option>
                <option>Marmer</option>
              </select>

              <input
                v-if="form.material_khusus === ''"
                v-model="form.material_khusus"
                type="text"
                placeholder="Custom material"
              />
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Kebutuhan Tambahan</label>
            <textarea v-model="form.kebutuhan_tambahan" class="field__textarea" placeholder="Additional information" rows="3"></textarea>
          </div>

          <div class="form-footer">
            <button class="btn-reset" @click="resetForm">Reset</button>
            <button class="btn-simpan" @click="prosesEstimasi" :disabled="isLoading">
              <span v-if="isLoading" class="spinner"></span>
              <span v-else>Simpan</span>
            </button>
          </div>
        </div>

        <div class="ringkasan-card">
          <h3 class="ringkasan-title">Ringkasan Input</h3>
          <div class="ringkasan-rows">
            <div class="ringkasan-row">
              <span class="ringkasan-label">Luas Area (m²)</span>
              <span class="ringkasan-value">{{ form.luas_area || '–' }} <span v-if="form.luas_area">x {{ form.luas_area }}m²</span></span>
            </div>
            <div class="ringkasan-row">
              <span class="ringkasan-label">Kerumitan</span>
              <span class="ringkasan-value">{{ form.tingkat_kerumitan || '–' }}</span>
            </div>
            <div class="ringkasan-row">
              <span class="ringkasan-label">Durasi</span>
             <span class="ringkasan-value">{{ form.durasi_pengerjaan ? form.durasi_pengerjaan + ' Hari' : '–' }}</span>
            </div>
          </div>

          <div class="ringkasan-divider"></div>

          <p class="ringkasan-disclaimer">
            Data yang Anda input hanya digunakan untuk proses estimasi biaya proyek interior. Hasil yang ditampilkan merupakan <strong>estimasi awal</strong> dan dapat berubah sesuai detail kebutuhan proyek. Halaman ini bukan halaman pembayaran atau transaksi.
          </p>

          <button class="btn-proses" @click="prosesEstimasi" :disabled="isLoading">
            <span v-if="isLoading" class="spinner"></span>
            <span v-else>Proses Estimasi</span>
          </button>
        </div>

      </div>
    </template>

    <!-- ══════════════════════════════════════
         VIEW 3: HASIL ESTIMASI
    ══════════════════════════════════════ -->
    <template v-if="currentView === 'hasil'">

      <div class="view-heading">
        <button class="back-btn" @click="currentView = 'form'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <h2 class="view-title">Hasil Estimasi Biaya Proyek</h2>
      </div>

      <div class="hasil-layout">

        <!-- Tabel variabel kiri -->
        <div class="hasil-table-card">
          <table class="variabel-table">
            <thead>
              <tr>
                <th>Variabel</th>
                <th>Detail Isi Variabel</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Nama Perusahaan</td>
                <td>{{ activeItem?.nama_perusahaan || form.nama_perusahaan || '–' }}</td>
              </tr>

              <tr>
                <td>Luas Area</td>
                <td>{{ activeItem?.luas_area || form.luas_area || '–' }}</td>
              </tr>

              <tr>
                <td>Jenis Proyek</td>
                <td>{{ activeItem?.jenis_proyek || form.jenis_proyek || '–' }}</td>
              </tr>

              <tr>
                <td>Tingkat Kerumitan</td>
                <td>{{ activeItem?.tingkat_kerumitan || form.tingkat_kerumitan || '–' }}</td>
              </tr>

              <tr>
                <td>Durasi Pengerjaan</td>
                <td>
                  {{
                    activeItem?.durasi_pengerjaan
                      ? activeItem.durasi_pengerjaan + ' Hari'
                      : (form.durasi_pengerjaan
                          ? form.durasi_pengerjaan + ' Hari'
                          : '–')
                  }}
                </td>
              </tr>

              <tr>
                <td>Lokasi Proyek</td>
                <td>{{ activeItem?.lokasi_proyek || form.lokasi_proyek || '–' }}</td>
              </tr>

              <tr>
                <td>Spesifikasi Desain</td>
                <td>{{ activeItem?.spesifikasi_design || form.spesifikasi_design || '–' }}</td>
              </tr>

              <tr>
                <td>Material Khusus</td>
                <td>{{ activeItem?.material_khusus || form.material_khusus || '–' }}</td>
              </tr>

              <tr>
                <td>Kebutuhan Tambahan</td>
                <td>{{ activeItem?.kebutuhan_tambahan || form.kebutuhan_tambahan || '–' }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="hasil-result-card">
          <h3 class="hasil-result-title">Hasil Estimasi</h3>
          <div class="hasil-result-rows">
            <div class="hasil-result-row">
              <span class="hasil-result-label">Minimum</span>
              <span class="hasil-result-value">{{ formatRupiah(hasilEstimasi?.minimum || 0) }}</span>
            </div>
            <div class="hasil-result-row">
              <span class="hasil-result-label">Maksimum</span>
              <span class="hasil-result-value">{{ formatRupiah(hasilEstimasi?.maksimum || 0) }}</span>
            </div>
            <div class="hasil-result-row highlight">
              <span class="hasil-result-label">Rekomendasi</span>
              <span class="hasil-result-value hasil-result-value--gold">{{ formatRupiah(hasilEstimasi?.rekomendasi || 0) }}</span>
            </div>
          </div>
          <button class="btn-simpan-data" @click="simpanData">Simpan Data</button>
        </div>

      </div>
    </template>

  </div>
</template>

<script setup>
import { ref, computed, onMounted, onBeforeUnmount, nextTick, watch } from 'vue'

const multiSelectOpen = ref(false)
const multiRef = ref(null)

const toggleMultiSelect = () => {
  multiSelectOpen.value = !multiSelectOpen.value
}

// close kalau klik luar
const handleClickOutside = (event) => {
  if (multiRef.value && !multiRef.value.contains(event.target)) {
    multiSelectOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})


// ─────────────────────────────────────────────
// API CONFIG 
// ─────────────────────────────────────────────
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8081/api'

// ─────────────────────────────────────────────
// STATE
// ─────────────────────────────────────────────
const currentView  = ref('list')
const searchQuery  = ref('')
const sortBy       = ref('date_desc')
const currentPage  = ref(1)
const perPage      = ref(12)
const isLoading    = ref(false)
const activeItem   = ref(null)
const showFilter = ref(false)
const filterJenis = ref('')

const estimasiList  = ref([])
const hasilEstimasi = ref({ minimum: 0, maksimum: 0, rekomendasi: 0 })

const avatarImg = ref(null)
const avatarFallback = ref(null)

const onAvatarError = async () => {
  await nextTick()
  if (avatarImg.value) avatarImg.value.style.display = 'none'
  if (avatarFallback.value) avatarFallback.value.style.display = 'flex'
}

const emptyForm = () => ({
  nama_perusahaan: '',
  luas_area: '',
  jenis_ruangan: [],
  jenis_pekerjaan: '',
  tingkat_kerumitan: '',
  durasi_pengerjaan: '', //
  lokasi_proyek: '',
  spesifikasi_design: '',
  material_khusus: '',
  kebutuhan_tambahan: '',
})
const daftarItemOptions = [
  //  Rumah
  "Living Room", "Master Room", "Bed Room", "Kitchen",
  "Dining Room", "Bath Room / Toilet", "Service Area",
  "Workspace", "Indoor Area", "Outdoor Area",
  "Backyard Garden", "Side Terrace",

  // Kantor
  "Lobby", "Director Room", "Manager Room", "Staff Room",
  "Office Room", "Meeting Room", "Waiting Room",
  "Consultation Room", "Treatment Room", "Training Room",
  "Cashier", "Canteen Area", "Mushalla",

  //  Bisnis
  "Pet Shop", "Pet Hotel", "Grooming Room", "Pet Clinic",

  //  Khusus
  "Understair Cabinet", "Kitchen Area", "Kitchen Set", "Kontainer Mart",

  //  Furniture
  "Bar Table", "Coffee Table", "Wooden Bench",
  "Wall Panel", "Pendant Lamp", "Backdrop Logo"
]

const form = ref(emptyForm())

// ─────────────────────────────────────────────
// LOAD DATA
// ─────────────────────────────────────────────
const loading = ref(false)
const apiError = ref('')

const fetchEstimasiList = async () => {
  loading.value = true
  apiError.value = ''
  try {
    const res = await fetch(`${API_BASE_URL}/admin/estimasi`)
    if (!res.ok) {
      const errText = await res.text()
      throw new Error(errText || `HTTP ${res.status}`)
    }
    estimasiList.value = await res.json()
  } catch (err) {
    apiError.value = err.message
    console.error('Gagal fetch:', err)
  } finally {
    loading.value = false
  }
}

onMounted(fetchEstimasiList)

// ─────────────────────────────────────────────
// HELPER
// ─────────────────────────────────────────────
const badgeClass = (jenis) => {
  if (!jenis) return ''

  const val = jenis.toLowerCase()

  if (val.includes('custom')) return 'badge--custom'
  if (val.includes('desain')) return 'badge--desain'
  if (val.includes('renovasi')) return 'badge--renovasi'
  if (val.includes('deteksi')) return 'badge--deteksi'
  if (val.includes('identifikasi')) return 'badge--identifikasi'

  return ''
}

const formatRupiah = (angka) => {
  if (!angka) return 'Rp 0'

  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0
  }).format(angka)
}

const formatTanggal = (tanggal) => {
  if (!tanggal) return '-'

  const date = new Date(tanggal)

  return date.toLocaleDateString('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric'
  })
}

// ─────────────────────────────────────────────
// PROSES ESTIMASI
// ─────────────────────────────────────────────
const prosesEstimasi = async () => {
  if (!form.value.nama_perusahaan) return

  isLoading.value = true

  try {
const payload = {
      nama_perusahaan: form.value.nama_perusahaan,
      luas_area: Number(form.value.luas_area),
      jenis_ruangan: form.value.jenis_ruangan,
      jenis_pekerjaan: form.value.jenis_pekerjaan,
      tingkat_kerumitan: form.value.tingkat_kerumitan,
      durasi_pengerjaan: Number(form.value.durasi_pengerjaan),
      lokasi_proyek: form.value.lokasi_proyek,
      spesifikasi_design: form.value.spesifikasi_design,
      material_khusus: form.value.material_khusus || null,
      kebutuhan_tambahan: form.value.kebutuhan_tambahan,
    }

    console.log("PAYLOAD FIX:", payload)

    const res = await fetch(`${API_BASE_URL}/admin/estimasi`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    })

    const text = await res.text()
    if (!res.ok) throw new Error(text)

    const data = JSON.parse(text)
    const estimasi = data.estimasi_harga || 0

    hasilEstimasi.value = {
      minimum: estimasi * 0.8,
      maksimum: estimasi * 1.2,
      rekomendasi: estimasi
    }

    currentView.value = 'hasil'

  } catch (err) {
    console.error('Estimasi error:', err)
  } finally {
    isLoading.value = false
  }
}
// ─────────────────────────────────────────────
// SIMPAN DATA
// ─────────────────────────────────────────────
const simpanData = async () => {
  try {
    const payload = {
      nama_perusahaan: form.value.nama_perusahaan,
      luas_area: Number(form.value.luas_area),
      jenis_ruangan: form.value.jenis_ruangan,
      jenis_pekerjaan: form.value.jenis_pekerjaan,
      tingkat_kerumitan: form.value.tingkat_kerumitan,
      durasi_pengerjaan: Number(form.value.durasi_pengerjaan),
      lokasi_proyek: form.value.lokasi_proyek,
      spesifikasi_design: form.value.spesifikasi_design,
      material_khusus: form.value.material_khusus,
      kebutuhan_tambahan: form.value.kebutuhan_tambahan,
      harga_proyek: hasilEstimasi.value.rekomendasi
    }

    const res = await fetch(`${API_BASE_URL}/admin/estimasi`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload)
    })

    const text = await res.text()
    if (!res.ok) throw new Error(text)

    await fetchEstimasiList()

    currentView.value = 'list'
    form.value = emptyForm()

  } catch (err) {
    console.error('Gagal simpan:', err)
  }
}

// ─────────────────────────────────────────────
// DELETE
// ─────────────────────────────────────────────
const deleteItem = async (id) => {
  try {
    const res = await fetch(`${API_BASE_URL}/admin/estimasi/${id}`, {
      method: "DELETE"
    })

    const text = await res.text()
    if (!res.ok) throw new Error(text)

    await fetchEstimasiList()

  } catch (err) {
    console.error('Delete error:', err)
  }
}

// ─────────────────────────────────────────────
// VIEW
// ─────────────────────────────────────────────
const openForm = () => {
  form.value = emptyForm()
  currentView.value = 'form'
}

// ─────────────────────────────────────────────
// COMPUTED (SEARCH + FILTER + SORT)
// ─────────────────────────────────────────────
const filteredList = computed(() => {
  let list = [...estimasiList.value]

// RESET PAGE kalau search / filter berubah
// watch([searchQuery, filterJenis], () => {
//   currentPage.value = 1
// })

  // 🔍 SEARCH
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    list = list.filter(p =>
      (p.nama_perusahaan || '').toLowerCase().includes(q)
    )
  }

  // FILTER JENIS
  if (filterJenis.value) {
    list = list.filter(p =>
      p.jenis_proyek === filterJenis.value
    )
  }

  // SORTING
  if (sortBy.value === 'date_desc') {
    list.sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
  }

  if (sortBy.value === 'date_asc') {
    list.sort((a, b) => new Date(a.created_at) - new Date(b.created_at))
  }

  if (sortBy.value === 'nama_asc') {
    list.sort((a, b) =>
      (a.nama_perusahaan || '').localeCompare(b.nama_perusahaan || '')
    )
  }

  return list
})

  // PAGINATION SLICE DATA
const paginatedList = computed(() => {
  const start = (currentPage.value - 1) * perPage.value
  const end = start + perPage.value
  return filteredList.value.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(filteredList.value.length / perPage.value) || 1
})

watch([searchQuery, filterJenis], () => {
  currentPage.value = 1
})

</script>

<style scoped>
/* ── Page ── */
.estimasi-page {
  padding: 24px 28px;
  min-height: 100vh;
  background: #F4F5F7;
  font-family: 'Montserrat', sans-serif;
}

/* ── Header ── */
.page-header {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 24px;
}
.page-title { font-size: 22px; font-weight: 700; color: #1A1A2E; }
.header-right { display: flex; align-items: center; gap: 10px; }

.search-box {
  display: flex; align-items: center; gap: 8px;
  background: #FFFFFF; border: 1.5px solid #E8E8E8;
  border-radius: 24px; padding: 8px 16px; width: 220px;
}
.search-icon { color: #C8A135; flex-shrink: 0; }
.search-input {
  border: none; outline: none;
  font-family: 'Montserrat', sans-serif; font-size: 12px;
  color: #333; background: transparent; width: 100%;
}
.search-input::placeholder { color: #AAAAAA; }

.notif-btn {
  position: relative; background: #FFFFFF; border: 1.5px solid #E8E8E8;
  border-radius: 50%; width: 38px; height: 38px;
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: #555;
}
.notif-dot {
  position: absolute; top: 7px; right: 7px;
  width: 6px; height: 6px; background: #C8A135;
  border-radius: 50%; border: 1.5px solid #fff;
}
.avatar {
  width: 38px; height: 38px; border-radius: 50%; overflow: hidden;
  background: #1B7A6E; display: flex; align-items: center; justify-content: center; flex-shrink: 0;
}
.avatar img { width: 100%; height: 100%; object-fit: cover; }
.avatar-fallback { display: none; align-items: center; justify-content: center; font-size: 13px; font-weight: 700; color: #fff; }

/* ── Toolbar ── */
.toolbar {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 16px; gap: 10px;
}
.toolbar-right { display: flex; align-items: center; gap: 10px; }

.btn-input-variabel {
  display: flex; align-items: center; gap: 6px;
  padding: 9px 18px; background: transparent;
  border: 1.5px solid #C8A135; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 12px; font-weight: 600;
  color: #C8A135; cursor: pointer; white-space: nowrap;
  transition: background 0.2s, color 0.2s;
}
.btn-input-variabel:hover { background: #FFF8E1; }

.select-wrap { position: relative; }
.ctrl-select {
  appearance: none; -webkit-appearance: none;
  padding: 8px 32px 8px 12px;
  background: #FFFFFF; border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 11px; font-weight: 600;
  color: #444; outline: none; cursor: pointer; white-space: nowrap;
}
.select-arrow {
  position: absolute; right: 8px; top: 50%;
  transform: translateY(-50%); color: #999; pointer-events: none;
}
.btn-filter {
  display: flex; align-items: center; gap: 6px;
  padding: 8px 14px; background: #FFFFFF;
  border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 11px; font-weight: 600;
  color: #666; cursor: pointer;
}

/* ── Table Card ── */
.table-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  overflow: hidden;
}
.table-card__title {
  font-size: 14px; font-weight: 700; color: #1B7A6E;
  padding: 18px 20px 12px;
}

.data-table {
  width: 100%; border-collapse: collapse; table-layout: fixed;
}
.data-table th:nth-child(1) { width: 5%; }
.data-table th:nth-child(2) { width: 22%; }
.data-table th:nth-child(3) { width: 18%; }
.data-table th:nth-child(4) { width: 16%; }
.data-table th:nth-child(5) { width: 18%; }
.data-table th:nth-child(6) { width: 10%; }
.data-table th:nth-child(7) { width: 11%; }

.data-table thead tr { border-bottom: 1px solid #F0F0F0; }
.data-table th {
  padding: 10px 16px;
  text-align: left; font-size: 11px; font-weight: 600;
  color: #AAAAAA;
}
.data-table tbody tr { border-bottom: 1px solid #F6F6F6; transition: background 0.15s; }
.data-table tbody tr:last-child { border-bottom: none; }
.data-table tbody tr:hover { background: #F8FDFC; }
.data-table td {
  padding: 13px 16px;
  font-size: 12px; color: #444; vertical-align: middle;
  overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
}

.td-no    { color: #BBBBBB; font-size: 12px; }
.td-main  { font-weight: 700; color: #1A1A2E; }
.td-muted { color: #AAAAAA; }
.td-file  { display: flex; align-items: center; gap: 5px; color: #555; }

/* Badge */
.badge {
  display: inline-block; padding: 4px 12px;
  border-radius: 20px; font-size: 10px; font-weight: 700; white-space: nowrap;
}
.badge--deteksi      { background: #E6EEFF; color: #3B5BDB; }
.badge--identifikasi { background: #FFFDE7; color: #C8A135; }
.badge--renovasi     { background: #F0F0F0; color: #777; }
.badge--desain       { background: #E6F9E6; color: #27AE60; }
.badge--custom       { background: #FFF0FB; color: #C026A0; }

/* Aksi buttons */
.td-aksi { text-align: right; white-space: nowrap; padding-right: 12px !important; overflow: visible; }
.action-btn {
  display: inline-flex; align-items: center; justify-content: center;
  width: 28px; height: 28px; border: none; border-radius: 6px;
  cursor: pointer; margin-left: 4px; background: transparent;
  transition: background 0.15s, transform 0.15s; color: #CCCCCC;
}
.action-btn:hover { transform: scale(1.1); }
.action-btn--edit:hover   { color: #1B7A6E; background: #EBF7F5; }
.action-btn--delete:hover { color: #E53E3E; background: #FEF0F0; }

/* Empty */
.empty-row { text-align: center; padding: 48px 0 !important; }
.empty-state { display: flex; flex-direction: column; align-items: center; gap: 12px; }
.empty-state p { font-size: 12px; color: #BBBBBB; }

/* ── Pagination ── */
.pagination-bar {
  display: flex; align-items: center; justify-content: space-between;
  padding: 14px 20px; border-top: 1px solid #F0F0F0;
}
.pagination-info { font-size: 12px; color: #999; }
.pagination-info strong { color: #333; }
.pagination-pages { display: flex; align-items: center; gap: 4px; }
.page-btn {
  min-width: 32px; height: 32px; padding: 0 6px;
  border: 1.5px solid #E8E8E8; border-radius: 8px; background: #fff;
  font-family: 'Montserrat', sans-serif; font-size: 12px; font-weight: 600;
  color: #555; cursor: pointer; display: flex; align-items: center; justify-content: center;
  transition: all 0.15s;
}
.page-btn:hover:not(:disabled):not(.active) { border-color: #1B7A6E; color: #1B7A6E; }
.page-btn.active { background: #1B3A2E; border-color: #1B3A2E; color: #fff; }
.page-btn:disabled { opacity: 0.4; cursor: not-allowed; }

/* ── View heading (form & hasil) ── */
.view-heading {
  display: flex; align-items: center; gap: 10px;
  margin-bottom: 20px;
}
.back-btn {
  display: flex; align-items: center; justify-content: center;
  background: none; border: none; cursor: pointer; color: #1B7A6E;
  padding: 4px; border-radius: 6px; transition: background 0.15s;
}
.back-btn:hover { background: #F0FAF8; }
.view-title { font-size: 16px; font-weight: 700; color: #1B7A6E; }

/* ── Form Layout ── */
.form-layout {
  display: grid; grid-template-columns: 1fr 340px; gap: 20px; align-items: start;
}

.form-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  padding: 28px 28px 24px;
  display: flex; flex-direction: column; gap: 18px;
}
.form-grid-2 { display: grid; grid-template-columns: 1fr 1fr; gap: 16px; }

.field { display: flex; flex-direction: column; gap: 6px; }
.field__label { font-size: 12px; font-weight: 600; color: #333; }

.field__input {
  width: 100%; padding: 11px 14px;
  border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px;
  color: #2C2C2C; background: #FFFFFF; outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}
.field__input::placeholder { color: #CCCCCC; }
.field__input:focus { border-color: #1B7A6E; box-shadow: 0 0 0 3px rgba(27,122,110,0.08); }

.field__textarea {
  width: 100%; padding: 11px 14px;
  border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px;
  color: #2C2C2C; background: #FFFFFF; outline: none; resize: vertical;
  transition: border-color 0.2s;
  box-sizing: border-box;
}
.field__textarea::placeholder { color: #CCCCCC; }
.field__textarea:focus { border-color: #1B7A6E; }

.field__select-wrap { position: relative; }
.field__select {
  width: 100%; padding: 11px 44px 11px 14px;
  border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px;
  color: #2C2C2C; background: #FFFFFF; outline: none;
  appearance: none; -webkit-appearance: none; cursor: pointer;
  transition: border-color 0.2s;
}
.field__select:focus { border-color: #1B7A6E; }
.field__select-arrow { position: absolute; right: 14px; top: 50%; transform: translateY(-50%); color: #999; pointer-events: none; }

.form-footer {
  display: flex; align-items: center; justify-content: space-between;
  margin-top: 4px;
}
.btn-reset {
  background: none; border: none; font-family: 'Montserrat', sans-serif;
  font-size: 13px; font-weight: 500; color: #AAAAAA; cursor: pointer;
}
.btn-reset:hover { color: #555; }
.btn-simpan {
  width: 100%; padding: 13px;
  background: #C8A135; border: none; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 14px; font-weight: 700;
  color: #FFFFFF; cursor: pointer; display: flex; align-items: center; justify-content: center; gap: 8px;
  transition: background 0.2s;
}
.btn-simpan:hover { background: #A8841C; }
.btn-simpan:disabled { opacity: 0.6; cursor: not-allowed; }

/* Ringkasan Card */
.ringkasan-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  padding: 22px 20px;
}
.ringkasan-title { font-size: 14px; font-weight: 700; color: #1A1A2E; margin-bottom: 16px; }
.ringkasan-rows { display: flex; flex-direction: column; gap: 10px; margin-bottom: 16px; }
.ringkasan-row { display: flex; justify-content: space-between; align-items: center; font-size: 12px; }
.ringkasan-label { color: #999; }
.ringkasan-value { font-weight: 600; color: #333; text-align: right; }
.ringkasan-divider { border: none; border-top: 1px solid #F0F0F0; margin: 16px 0; }
.ringkasan-disclaimer {
  font-size: 11px; color: #888; line-height: 1.6; margin-bottom: 18px;
}
.ringkasan-disclaimer strong { color: #333; }
.btn-proses {
  width: 100%; padding: 12px;
  background: #FFFFFF; border: 1.5px solid #333; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  color: #333; cursor: pointer; display: flex; align-items: center; justify-content: center; gap: 8px;
  transition: background 0.2s;
}
.btn-proses:hover { background: #F4F5F7; }
.btn-proses:disabled { opacity: 0.6; cursor: not-allowed; }

/* ── Hasil Layout ── */
.hasil-layout {
  display: grid; grid-template-columns: 1fr 280px; gap: 20px; align-items: start;
}

.hasil-table-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06); overflow: hidden;
}

.variabel-table { width: 100%; border-collapse: collapse; }
.variabel-table thead tr {
  background: #C8A135;
}
.variabel-table th {
  padding: 13px 20px;
  text-align: left; font-size: 12px; font-weight: 600; color: #FFFFFF;
}
.variabel-table tbody tr { border-bottom: 1px solid #F6F6F6; transition: background 0.15s; }
.variabel-table tbody tr:last-child { border-bottom: none; }
.variabel-table tbody tr:hover { background: #FAFAFA; }
.variabel-table td {
  padding: 13px 20px; font-size: 13px; color: #333; vertical-align: middle;
}
.variabel-table td:first-child { color: #555; width: 40%; }

.hasil-result-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  padding: 22px 20px;
}
.hasil-result-title { font-size: 16px; font-weight: 700; color: #1A1A2E; margin-bottom: 20px; text-align: center; }
.hasil-result-rows { display: flex; flex-direction: column; gap: 14px; margin-bottom: 24px; }
.hasil-result-row { display: flex; justify-content: space-between; align-items: center; }
.hasil-result-label { font-size: 13px; font-weight: 600; color: #333; }
.hasil-result-value { font-size: 13px; color: #888; text-align: right; }
.hasil-result-value--gold { color: #C8A135 !important; font-weight: 700; }

.btn-simpan-data {
  width: 100%; padding: 11px;
  background: #FFFFFF; border: 1.5px solid #D0D0D0; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 600;
  color: #333; cursor: pointer;
  transition: background 0.2s, border-color 0.2s;
}
.btn-simpan-data:hover { background: #F4F5F7; border-color: #999; }

.checkbox-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 6px;
  font-size: 12px;
}

.multiselect {
  position: relative;
}

.multiselect__trigger {
  border: 1.5px solid #E8E8E8;
  border-radius: 8px;
  padding: 11px 14px;
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 13px;
  background: #fff;
}

.multiselect__placeholder {
  color: #aaa;
}

.multiselect__value {
  color: #333;
}

.multiselect__dropdown {
  position: absolute;
  top: 110%;
  left: 0;
  width: 100%;
  max-height: 220px;
  overflow-y: auto;
  background: #fff;
  border: 1.5px solid #E8E8E8;
  border-radius: 8px;
  z-index: 10;
  padding: 10px;
}

.multiselect__option {
  display: flex;
  gap: 8px;
  font-size: 12px;
  padding: 6px;
  cursor: pointer;
}
/* ── Spinner ── */
.spinner {
  display: inline-block; width: 14px; height: 14px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff; border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>