<template>
  <div class="estimasi-page">

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
          <button class="btn-filter">
            Filter
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"/></svg>
          </button>
        </div>
      </div>

      <div class="table-card">
        <div class="table-card__title">Data Estimasi Tersimpan</div>
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama Proyek</th>
              <th>Jenis Proyek</th>
              <th>Luas Ruangan</th>
              <th>Biaya</th>
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
            <tr v-for="(item, index) in filteredList" :key="item.id">
              <td class="td-no">{{ index + 1 }}</td>
              <td class="td-main">{{ item.namaProyek }}</td>
              <td class="td-cell">
                <span class="badge" :class="badgeClass(item.jenisProyek)">{{ item.jenisProyek }}</span>
              </td>
              <td class="td-cell">{{ item.luasRuangan }}</td>
              <td class="td-cell td-file">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
                {{ item.biayaFile }}
              </td>
              <td class="td-cell td-muted">{{ formatTanggal(item.tanggal) }}</td>
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
          <span class="pagination-info">Showing <strong>1-{{ Math.min(perPage, filteredList.length) }}</strong> from <strong>{{ filteredList.length }}</strong> data</span>
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
              <label class="field__label">Nama Proyek</label>
              <input v-model="form.namaProyek" type="text" class="field__input" placeholder="" />
            </div>
            <div class="field">
              <label class="field__label">Luas Ruangan (m²)</label>
              <input v-model="form.luasRuangan" type="text" class="field__input" placeholder="" />
            </div>
          </div>

          <div class="field">
            <label class="field__label">Jenis Proyek</label>
            <div class="field__select-wrap">
              <select v-model="form.jenisProyek" class="field__select">
                <option value="" disabled>Pilih jenis proyek</option>
                <option>Custom Furniture</option>
                <option>Desain Interior</option>
                <option>Alat Deteksi</option>
                <option>Alat Identifikasi</option>
                <option>Renovasi</option>
              </select>
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Tingkat Kerumitan</label>
            <div class="field__select-wrap">
              <select v-model="form.tingkatKerumitan" class="field__select">
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
            <input v-model="form.durasiPengerjaan" type="number" class="field__input" placeholder="" />
          </div>

          <div class="field">
            <label class="field__label">Lokasi Proyek</label>
            <input v-model="form.lokasiProyek" type="text" class="field__input" placeholder="" />
          </div>

          <div class="field">
            <label class="field__label">Konsep/Gaya Desain</label>
            <div class="field__select-wrap">
              <select v-model="form.konsepDesain" class="field__select">
                <option value="" disabled>Pilih konsep/gaya desain</option>
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
              <select v-model="form.materialKhusus" class="field__select">
                <option value="" disabled>Pilih material khusus</option>
                <option>Kayu Jati</option>
                <option>Kayu Pinus</option>
                <option>Baja</option>
                <option>Marmer</option>
                <option>Granit</option>
                <option>Kaca Tempered</option>
              </select>
              <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <div class="field">
            <label class="field__label">Kebutuhan Tambahan</label>
            <textarea v-model="form.kebutuhanTambahan" class="field__textarea" placeholder="Additional information" rows="3"></textarea>
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
              <span class="ringkasan-label">Luas Ruangan (m²)</span>
              <span class="ringkasan-value">{{ form.luasRuangan || '–' }} <span v-if="form.luasRuangan">x {{ form.luasRuangan }}m²</span></span>
            </div>
            <div class="ringkasan-row">
              <span class="ringkasan-label">Kerumitan</span>
              <span class="ringkasan-value">{{ form.tingkatKerumitan || '–' }}</span>
            </div>
            <div class="ringkasan-row">
              <span class="ringkasan-label">Durasi</span>
              <span class="ringkasan-value">{{ form.durasiPengerjaan ? form.durasiPengerjaan + ' Hari' : '–' }}</span>
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
              <tr><td>Nama Proyek</td><td>{{ activeItem?.namaProyek || form.namaProyek }}</td></tr>
              <tr><td>Luas Ruangan</td><td>{{ activeItem?.luasRuangan || form.luasRuangan }}</td></tr>
              <tr><td>Jenis Proyek</td><td>{{ activeItem?.jenisProyek || form.jenisProyek }}</td></tr>
              <tr><td>Tingkat Kerumitan</td><td>{{ activeItem?.tingkatKerumitan || form.tingkatKerumitan }}</td></tr>
              <tr><td>Durasi Pengerjaan</td><td>{{ activeItem?.durasiPengerjaan ? activeItem.durasiPengerjaan + ' Hari' : (form.durasiPengerjaan ? form.durasiPengerjaan + ' Hari' : '–') }}</td></tr>
              <tr><td>Lokasi Proyek</td><td>{{ activeItem?.lokasiProyek || form.lokasiProyek }}</td></tr>
              <tr><td>Konsep/Desain</td><td>{{ activeItem?.konsepDesain || form.konsepDesain }}</td></tr>
              <tr><td>Material Khusus</td><td>{{ activeItem?.materialKhusus || form.materialKhusus }}</td></tr>
              <tr><td>Kebutuhan Tambahan</td><td>{{ activeItem?.kebutuhanTambahan || form.kebutuhanTambahan || '–' }}</td></tr>
            </tbody>
          </table>
        </div>

        <div class="hasil-result-card">
          <h3 class="hasil-result-title">Hasil Estimasi</h3>
          <div class="hasil-result-rows">
            <div class="hasil-result-row">
              <span class="hasil-result-label">Minimum</span>
              <span class="hasil-result-value">{{ formatRupiah(hasilEstimasi.minimum) }}</span>
            </div>
            <div class="hasil-result-row">
              <span class="hasil-result-label">Maksimum</span>
              <span class="hasil-result-value">{{ formatRupiah(hasilEstimasi.maksimum) }}</span>
            </div>
            <div class="hasil-result-row highlight">
              <span class="hasil-result-label">Rekomendasi</span>
              <span class="hasil-result-value hasil-result-value--gold">{{ formatRupiah(hasilEstimasi.rekomendasi) }}</span>
            </div>
          </div>
          <button class="btn-simpan-data" @click="simpanData">Simpan Data</button>
        </div>

      </div>
    </template>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

// ─────────────────────────────────────────────
// API CONFIG 
// ─────────────────────────────────────────────
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

// TODO: uncomment saat backend siap
// const fetchEstimasiList  = () => fetch(`${API_BASE_URL}/estimasi`).then(r => r.json())
// const postEstimasi       = (payload) => fetch(`${API_BASE_URL}/estimasi`, { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(payload) }).then(r => r.json())
// const deleteEstimasi     = (id) => fetch(`${API_BASE_URL}/estimasi/${id}`, { method: 'DELETE' })
// const getEstimasiById    = (id) => fetch(`${API_BASE_URL}/estimasi/${id}`).then(r => r.json())

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
const avatarImg    = ref(null)
const avatarFallback = ref(null)

const hasilEstimasi = ref({ minimum: 0, maksimum: 0, rekomendasi: 0 })

const emptyForm = () => ({
  namaProyek: '', luasRuangan: '', jenisProyek: '',
  tingkatKerumitan: '', durasiPengerjaan: '', lokasiProyek: '',
  konsepDesain: '', materialKhusus: '', kebutuhanTambahan: '',
})

const form = ref(emptyForm())

// ─────────────────────────────────────────────
// DUMMY DATA 
// ─────────────────────────────────────────────
const estimasiList = ref([
  { id: 1,  namaProyek: 'ION SCAN 500DT', jenisProyek: 'Alat Deteksi',     luasRuangan: 'Ion scan 500dt.pdf', biayaFile: 'Ion scan 500dt.pdf', tanggal: '2015-05-27', tingkatKerumitan: 'Sedang', durasiPengerjaan: 10, lokasiProyek: 'Jakarta', konsepDesain: 'Modern', materialKhusus: '', kebutuhanTambahan: '' },
  { id: 2,  namaProyek: 'ION SCAN 500DT', jenisProyek: 'Alat Deteksi',     luasRuangan: 'Ion scan 500dt.pdf', biayaFile: 'Ion scan 500dt.pdf', tanggal: '2012-05-19', tingkatKerumitan: 'Mudah',  durasiPengerjaan: 7,  lokasiProyek: 'Bogor',   konsepDesain: 'Japandi', materialKhusus: 'Kayu Jati', kebutuhanTambahan: '' },
  { id: 3,  namaProyek: 'RIGAKU',          jenisProyek: 'Alat Identifikasi',luasRuangan: 'Rigaku.pdf',         biayaFile: 'Rigaku.pdf',         tanggal: '2016-03-04', tingkatKerumitan: 'Sulit',  durasiPengerjaan: 14, lokasiProyek: 'Surabaya',konsepDesain: 'Industrial', materialKhusus: 'Baja', kebutuhanTambahan: '' },
  { id: 12, namaProyek: 'RIGAKU',          jenisProyek: 'Alat Identifikasi',luasRuangan: 'Rigaku.pdf',         biayaFile: 'Rigaku.pdf',         tanggal: '2013-07-27', tingkatKerumitan: 'Sedang', durasiPengerjaan: 10, lokasiProyek: 'Bandung', konsepDesain: 'Minimalis', materialKhusus: '', kebutuhanTambahan: '' },
])

// ─────────────────────────────────────────────
// COMPUTED
// ─────────────────────────────────────────────
const filteredList = computed(() => {
  let list = [...estimasiList.value]
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    list = list.filter(p =>
      p.namaProyek.toLowerCase().includes(q) ||
      p.jenisProyek.toLowerCase().includes(q)
    )
  }
  list.sort((a, b) => {
    if (sortBy.value === 'date_desc') return new Date(b.tanggal) - new Date(a.tanggal)
    if (sortBy.value === 'date_asc')  return new Date(a.tanggal) - new Date(b.tanggal)
    if (sortBy.value === 'nama_asc')  return a.namaProyek.localeCompare(b.namaProyek)
    return 0
  })
  return list
})

const totalPages = computed(() => Math.max(1, Math.ceil(filteredList.value.length / perPage.value)))

// ─────────────────────────────────────────────
// METHODS
// ─────────────────────────────────────────────
const openForm = () => {
  form.value = emptyForm()
  activeItem.value = null
  currentView.value = 'form'
}

const resetForm = () => { form.value = emptyForm() }

/**
 * prosesEstimasi — kirim ke backend Golang
 */
const prosesEstimasi = async () => {
  if (!form.value.namaProyek || !form.value.jenisProyek) return

  isLoading.value = true

  try {
    // ── DUMMY LOGIC  ──────────────────────
    await new Promise(r => setTimeout(r, 800)) // simulasi loading
    const luas = parseFloat(form.value.luasRuangan) || 10
    const durasi = parseFloat(form.value.durasiPengerjaan) || 10
    const base = luas * durasi * 500000
    hasilEstimasi.value = {
      minimum:     Math.round(base * 0.5),
      maksimum:    Math.round(base * 2.5),
      rekomendasi: Math.round(base),
    }
    // ── END DUMMY ──────────────────────────────────────────────────

    // ── API CALL ─────────────────────
    // const payload = {
    //   nama_proyek:        form.value.namaProyek,
    //   luas_ruangan:       form.value.luasRuangan,
    //   jenis_proyek:       form.value.jenisProyek,
    //   tingkat_kerumitan:  form.value.tingkatKerumitan,
    //   durasi_pengerjaan:  Number(form.value.durasiPengerjaan),
    //   lokasi_proyek:      form.value.lokasiProyek,
    //   konsep_desain:      form.value.konsepDesain,
    //   material_khusus:    form.value.materialKhusus,
    //   kebutuhan_tambahan: form.value.kebutuhanTambahan,
    // }
    // const res  = await postEstimasi(payload)
    // hasilEstimasi.value = {
    //   minimum:     res.minimum,
    //   maksimum:    res.maksimum,
    //   rekomendasi: res.rekomendasi,
    // }
    // ── END API CALL ───────────────────────────────────────────────

    currentView.value = 'hasil'
  } catch (err) {
    console.error('Estimasi error:', err)
  } finally {
    isLoading.value = false
  }
}

/**
 * simpanData — simpan hasil estimasi ke list
 */
const simpanData = async () => {
  // ── DUMMY ──────────────────────────────────────────────────────
  const newItem = {
    id: Date.now(),
    namaProyek:        form.value.namaProyek,
    jenisProyek:       form.value.jenisProyek,
    luasRuangan:       form.value.luasRuangan,
    biayaFile:         `${form.value.namaProyek}.pdf`,
    tanggal:           new Date().toISOString().split('T')[0],
    tingkatKerumitan:  form.value.tingkatKerumitan,
    durasiPengerjaan:  form.value.durasiPengerjaan,
    lokasiProyek:      form.value.lokasiProyek,
    konsepDesain:      form.value.konsepDesain,
    materialKhusus:    form.value.materialKhusus,
    kebutuhanTambahan: form.value.kebutuhanTambahan,
  }
  estimasiList.value.push(newItem)
  // ── END DUMMY ──────────────────────────────────────────────────

  // ── API CALL ─────────────────────
  // await postEstimasi({ ...payload, hasil: hasilEstimasi.value })
  // await fetchEstimasiList().then(data => { estimasiList.value = data })
  // ── END API CALL ───────────────────────────────────────────────

  currentView.value = 'list'
  form.value = emptyForm()
  activeItem.value = null
}

const viewHasil = (item) => {
  activeItem.value = item
  // ── DUMMY hasil ─────────────────────────────────────────────────
  hasilEstimasi.value = { minimum: 50000, maksimum: 250000, rekomendasi: 100000 }
  // ── API  ──────────────────────────
  // getEstimasiById(item.id).then(res => { hasilEstimasi.value = res.hasil })
  currentView.value = 'hasil'
}

/**
 * deleteItem — hapus dari list
 */
const deleteItem = (id) => {
  estimasiList.value = estimasiList.value.filter(p => p.id !== id)
  // ── API  ──────────────────────────
  // deleteEstimasi(id)
}

// ─────────────────────────────────────────────
// HELPERS
// ─────────────────────────────────────────────
const badgeClass = (j) => {
  if (j === 'Alat Deteksi')      return 'badge--deteksi'
  if (j === 'Alat Identifikasi') return 'badge--identifikasi'
  if (j === 'Renovasi')          return 'badge--renovasi'
  if (j === 'Desain Interior')   return 'badge--desain'
  if (j === 'Custom Furniture')  return 'badge--custom'
  return ''
}

const formatRupiah = (v) => {
  if (!v && v !== 0) return '–'
  return 'Rs. ' + Number(v).toLocaleString('en-US', { minimumFractionDigits: 2 })
}

const formatTanggal = (v) => {
  if (!v) return '–'
  const d = new Date(v)
  return `${d.getMonth()+1}/${d.getDate()}/${String(d.getFullYear()).slice(2)}`
}

const onAvatarError = () => {
  if (avatarImg.value) avatarImg.value.style.display = 'none'
  if (avatarFallback.value) avatarFallback.value.style.display = 'flex'
}
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

/* ── Spinner ── */
.spinner {
  display: inline-block; width: 14px; height: 14px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff; border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>