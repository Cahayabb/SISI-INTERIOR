<template>
  <div class="project-page">

    <!-- ── HEADER ── -->
    <div class="page-header">
      <h1 class="page-title">Data Proyek</h1>
      <div class="header-right">
        <div class="search-box">
          <svg class="search-icon" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
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

    <!-- ══════════════════════════════════════════
         VIEW 1 : LIST DATA PROYEK
    ══════════════════════════════════════════ -->
    <template v-if="!showForm">

      <!-- Toolbar -->
      <div class="toolbar">
        <button class="btn-tambah" @click="openForm">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          Tambah Data Proyek
        </button>

        <div class="toolbar-right">
          <div class="ctrl-wrap">
            <label class="ctrl-label">Sort by</label>
            <div class="select-wrap">
              <select v-model="sortBy" class="ctrl-select">
                <option value="tanggal_desc">Terbaru</option>
                <option value="tanggal_asc">Terlama</option>
                <option value="nama_asc">Nama A–Z</option>
                <option value="nama_desc">Nama Z–A</option>
                <option value="luas_desc">Luas Terbesar</option>
                <option value="luas_asc">Luas Terkecil</option>
              </select>
              <svg class="select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>
          <div class="ctrl-wrap">
            <label class="ctrl-label">Filter</label>
            <div class="select-wrap">
              <select v-model="filterJenis" class="ctrl-select">
                <option value="">Semua</option>
                <option value="Renovasi">Renovasi</option>
                <option value="Desain">Desain</option>
                <option value="Custom">Custom</option>
              </select>
              <svg class="select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Table Card -->
      <div class="table-card">
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama Perusahaan</th>
              <th>Jenis Proyek</th>
              <th>Luas Area</th>
              <th>Jenis Pekerjaan</th>
              <th>Spesifikasi Design</th>
              <th>Area Layanan</th>
              <th>Daftar Item</th>
              <th>Spesifikasi Item</th>
              <th>Harga Satuan</th>
              <th>Tanggal</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="displayedProyek.length === 0">
              <td colspan="12" class="empty-row">
                <div class="empty-state">
                  <svg width="36" height="36" viewBox="0 0 24 24" fill="none" stroke="#D0D0D0" stroke-width="1.5">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                    <polyline points="14 2 14 8 20 8"/>
                    <line x1="16" y1="13" x2="8" y2="13"/>
                    <line x1="16" y1="17" x2="8" y2="17"/>
                  </svg>
                  <p>Belum ada data proyek.<br/>Klik <strong>Tambah Data Proyek</strong> untuk mulai.</p>
                </div>
              </td>
            </tr>
            <tr v-for="(item, index) in displayedProyek" :key="item.id">
              <td class="td-no">{{ index + 1 }}</td>
              <td class="td-main">{{ item.nama_perusahaan }}</td>
              <td class="td-cell">{{ item.jenis_proyek }}</td>
              <td class="td-cell td-nowrap">{{ item.luas_area }} m²</td>
              <td class="td-cell">
                <span class="badge" :class="badgeClass(item.jenisPekerjaan)">
                  {{ item.jenisPekerjaan }}
                </span>
              </td>
              <td class="td-cell td-truncate">{{ item.spesifikasiDesign }}</td>
              <td class="td-cell">{{ item.areaLayanan }}</td>
              <td class="td-cell td-truncate">{{ formatDaftarItem(item.daftarItem) }}</td>
              <td class="td-cell td-muted td-truncate">{{ item.spesifikasiItem || '–' }}</td>
              <td class="td-cell td-nowrap">{{ item.harga_satuan ? formatRupiah(item.harga_satuan) : '–' }}</td>
              <td class="td-cell td-nowrap td-muted">{{ formatTanggal(item.tanggal) }}</td>
              <td class="td-aksi">
                <button class="action-btn action-btn--edit" @click="editItem(item)" title="Edit">
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
      </div>

    </template>

    <!-- ══════════════════════════════════════════
         VIEW 2 : FORM TAMBAH / EDIT
    ══════════════════════════════════════════ -->
    <div v-else class="form-card">

      <div class="form-card__heading">
        <button class="back-btn" @click="closeForm">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <h2 class="form-title">{{ editMode ? 'Edit Data Proyek' : 'Tambah Data Proyek' }}</h2>
      </div>

      <div class="form-body">

        <div class="field">
          <label class="field__label">Nama Perusahaan</label>
          <input v-model="form.nama_perusahaan" type="text" class="field__input" placeholder="Masukkan nama perusahaan" />
        </div>

        <div class="field">
          <label class="field__label">Jenis Proyek</label>
          <input v-model="form.jenis_proyek" type="text" class="field__input" placeholder="Tentukan jenis proyek" />
        </div>

        <div class="field">
          <label class="field__label">Luas Area</label>
          <div class="input-suffix-wrap">
            <input v-model="form.luas_area" type="number" class="field__input field__input--suffix" placeholder="Masukkan nilai luas area dalam m²" />
            <span class="input-suffix">m²</span>
          </div>
        </div>

        <div class="field">
          <label class="field__label">Jenis Pekerjaan</label>
          <div class="field__select-wrap">
            <select v-model="form.jenisPekerjaan" class="field__select" :class="{ filled: form.jenisPekerjaan }">
              <option value="" disabled>Pilih jenis pekerjaan yang sesuai</option>
              <option>Renovasi</option>
              <option>Desain</option>
              <option>Custom</option>
            </select>
            <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
          </div>
        </div>

        <div class="field">
          <label class="field__label">Spesifikasi Design</label>
          <div class="field__select-wrap">
            <select v-model="form.spesifikasiDesign" class="field__select" :class="{ filled: form.spesifikasiDesign }">
              <option value="" disabled>Masukkan spesifikasi design yang sesuai</option>
              <option>Japandi</option>
              <option>Industrial</option>
              <option>Scandinavian</option>
              <option>Modern</option>
              <option>Modern Minimalist</option>
              <option>Ekletik / Colorful</option>
              <option>Bohemian</option>
              <option>Klasik / Victorian</option>
              <option>Rustic</option>
              <option>Coastal / Nautical</option>
              <option>Art Deco</option>
              <option>Minimalis</option>
            </select>
            <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
          </div>
        </div>

        <div class="field">
          <label class="field__label">Area Layanan</label>
          <div class="field__select-wrap">
            <select v-model="form.areaLayanan" class="field__select" :class="{ filled: form.areaLayanan }">
              <option value="" disabled>Pilih area layanan</option>
              <option>Jakarta</option>
              <option>Bogor</option>
              <option>Depok</option>
              <option>Tangerang</option>
              <option>Bekasi</option>
              <option>Bandung</option>
              <option>Surabaya</option>
            </select>
            <svg class="field__select-arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
          </div>
        </div>

        <div class="field">
          <label class="field__label">Daftar Item</label>
          <div class="multiselect" @click.stop>
            <div class="multiselect__trigger" @click="toggleMultiSelect">
              <span v-if="form.daftarItem.length === 0" class="multiselect__placeholder">Silahkan pilih beberapa daftar item</span>
              <span v-else class="multiselect__value">{{ form.daftarItem.join(', ') }}</span>
              <svg class="multiselect__arrow" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"
                :style="{ transform: multiSelectOpen ? 'rotate(180deg)' : 'rotate(0deg)' }">
                <polyline points="6 9 12 15 18 9"/>
              </svg>
            </div>
            <div v-if="multiSelectOpen" class="multiselect__dropdown">
              <label v-for="item in daftarItemOptions" :key="item" class="multiselect__option">
                <input type="checkbox" :value="item" v-model="form.daftarItem" class="multiselect__checkbox" />
                <span>{{ item }}</span>
              </label>
            </div>
          </div>
        </div>

        <div class="field">
          <label class="field__label">Spesifikasi Item <span class="field__optional">(Opsional)</span></label>
          <input v-model="form.spesifikasiItem" type="text" class="field__input" placeholder="Opsional" />
        </div>

        <div class="field">
          <label class="field__label">Harga Satuan <span class="field__optional">(Opsional)</span></label>
          <div class="input-prefix-wrap">
            <span class="input-prefix">Rp</span>
            <input
              :value="hargaDisplay"
              @input="onHargaInput"
              type="text"
              inputmode="numeric"
              class="field__input field__input--prefix"
              placeholder="0"
            />
          </div>
        </div>

        <div class="form-actions">
          <button class="btn-cancel" @click="closeForm">Cancel</button>
          <button class="btn-upload" @click="handleUpload">Upload</button>
        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

const showForm        = ref(false)
const editMode        = ref(false)
const searchQuery     = ref('')
const sortBy          = ref('tanggal_desc')
const filterJenis     = ref('')
const multiSelectOpen = ref(false)
const avatarImg       = ref(null)
const avatarFallback  = ref(null)

const daftarItemOptions = [
  'Living Room', 'Kitchen', 'Bedroom', 'Bathroom',
  'Dining Room', 'Home Office', 'Foyer / Entrance',
  'Balcony', 'Garden', 'Basement', 'Garage',
]

const proyekList = ref([])

const getProjects = async () => {
  try {
    const res = await fetch("http://localhost:8081/api/admin/projects")
    const data = await res.json()
    console.log("RESPONSE:", data)

    // mapping ke format frontend kamu
    proyekList.value = data.map(p => ({
      id: p.id,
      nama_perusahaan: p.nama_perusahaan,
      jenis_proyek: p.jenis_proyek,
      luas_area: p.luas_area,
      jenisPekerjaan: p.jenis_pekerjaan,
      spesifikasiDesign: p.spesifikasi_design,
      areaLayanan: p.area_layanan,
      daftarItem: p.daftar_item ? p.daftar_item.split(',') : [],
      spesifikasiItem: p.spesifikasi_item,
      harga_satuan: p.harga_satuan,
      tanggal: p.tanggal
    }))

  } catch (err) {
    console.error("Gagal ambil data:", err)
  }
}

onMounted(() => {
  getProjects()
  document.addEventListener('click', closeMultiSelect)
})

const emptyForm = () => ({
  id: null, nama_perusahaan: '', jenis_proyek: '', luas_area: '',
  jenisPekerjaan: '', spesifikasiDesign: '', areaLayanan: '',
  daftarItem: [], spesifikasiItem: '', harga_satuan: '', tanggal: '',
})

const form = ref(emptyForm())

const hargaDisplay = computed(() => {
  if (!form.value.harga_satuan) return ''
  return Number(form.value.harga_satuan).toLocaleString('id-ID')
})

const onHargaInput = (e) => {
  const raw = e.target.value.replace(/\D/g, '')
  form.value.harga_satuan = raw
  e.target.value = raw ? Number(raw).toLocaleString('id-ID') : ''
}

const displayedProyek = computed(() => {
  let list = [...proyekList.value]
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    list = list.filter(p =>
      p.nama_perusahaan.toLowerCase().includes(q) ||
      p.jenis_proyek.toLowerCase().includes(q) ||
      p.areaLayanan.toLowerCase().includes(q) ||
      p.spesifikasiDesign.toLowerCase().includes(q)
    )
  }
  if (filterJenis.value) list = list.filter(p => p.jenisPekerjaan === filterJenis.value)
  list.sort((a, b) => {
    if (sortBy.value === 'tanggal_desc') return new Date(b.tanggal) - new Date(a.tanggal)
    if (sortBy.value === 'tanggal_asc')  return new Date(a.tanggal) - new Date(b.tanggal)
    if (sortBy.value === 'nama_asc')     return a.nama_perusahaan.localeCompare(b.nama_perusahaan)
    if (sortBy.value === 'nama_desc')    return b.nama_perusahaan.localeCompare(a.nama_perusahaan)
    if (sortBy.value === 'luas_desc')    return b.luas_area - a.luas_area
    if (sortBy.value === 'luas_asc')     return a.luas_area - b.luas_area
    return 0
  })
  return list
})

const badgeClass = (j) => j === 'Renovasi' ? 'badge--renovasi' : j === 'Desain' ? 'badge--desain' : j === 'Custom' ? 'badge--custom' : ''
const formatRupiah = (v) => v ? 'Rp ' + Number(v).toLocaleString('id-ID') : '–'
const formatTanggal = (v) => { if (!v) return '–'; const d = new Date(v); return `${d.getMonth()+1}/${d.getDate()}/${String(d.getFullYear()).slice(2)}` }
const formatDaftarItem = (arr) => { if (!arr || !arr.length) return '–'; const s = arr.join(', '); return s.length > 22 ? s.slice(0,22)+'...' : s }

const toggleMultiSelect = () => { multiSelectOpen.value = !multiSelectOpen.value }
const closeMultiSelect  = () => { multiSelectOpen.value = false }

onUnmounted(() => document.removeEventListener('click', closeMultiSelect))

const openForm  = () => { form.value = emptyForm(); editMode.value = false; showForm.value = true }
const closeForm = () => { showForm.value = false; editMode.value = false; form.value = emptyForm(); multiSelectOpen.value = false }

const handleUpload = async () => {
  if (!form.value.nama_perusahaan) return

  try {
    const payload = {
      nama_perusahaan: form.value.nama_perusahaan,
      jenis_proyek: form.value.jenis_proyek,
      luas_area: Number(form.value.luas_area),
      jenis_pekerjaan: form.value.jenisPekerjaan,
      spesifikasi_design: form.value.spesifikasiDesign,
      area_layanan: form.value.areaLayanan,
      daftar_item: form.value.daftarItem.join(','), // array → string
      spesifikasi_item: form.value.spesifikasiItem,
      harga_satuan: Number(form.value.harga_satuan),
      tanggal: new Date().toISOString()
    }

    const res = await fetch("http://localhost:8081/api/admin/projects", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    })

    if (!res.ok) {
      const errData = await res.json()
      console.error("Server error:", errData)
      return
    }

    //refresh dari database
    await getProjects()

    closeForm()

  } catch (err) {
    console.error("Gagal upload:", err)
  }
}

const editItem   = (item) => { form.value = { ...item, daftarItem: [...(item.daftarItem || [])] }; editMode.value = true; showForm.value = true }
const deleteItem = async (id) => {
  try {
    await fetch(`http://localhost:8081/api/admin/projects/${id}`, {
      method: "DELETE"
    })

    await getProjects()

  } catch (err) {
    console.error("Gagal delete:", err)
  }
}
const onAvatarError = () => { if (avatarImg.value) avatarImg.value.style.display = 'none'; if (avatarFallback.value) avatarFallback.value.style.display = 'flex' }
</script>


<style scoped>
/* ── Page ── */
.project-page {
  padding: 24px 28px;
  min-height: 100vh;
  background: #F4F5F7;
  font-family: 'Montserrat', sans-serif;
}

/* ── Header ── */
.page-header {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 20px;
}
.page-title { font-size: 22px; font-weight: 700; color: #1A1A2E; }
.header-right { display: flex; align-items: center; gap: 10px; }

.search-box {
  display: flex; align-items: center; gap: 8px;
  background: #FFFFFF; border: 1.5px solid #E8E8E8;
  border-radius: 9px; padding: 8px 14px; width: 200px;
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
  border-radius: 9px; width: 36px; height: 36px;
  display: flex; align-items: center; justify-content: center;
  cursor: pointer; color: #555;
}
.notif-dot {
  position: absolute; top: 7px; right: 7px;
  width: 6px; height: 6px; background: #C8A135;
  border-radius: 50%; border: 1.5px solid #fff;
}
.avatar {
  width: 36px; height: 36px; border-radius: 50%; overflow: hidden;
  background: #1B7A6E; display: flex; align-items: center; justify-content: center; flex-shrink: 0;
}
.avatar img { width: 100%; height: 100%; object-fit: cover; }
.avatar-fallback { display: none; align-items: center; justify-content: center; font-size: 13px; font-weight: 700; color: #fff; }

/* ── Toolbar ── */
.toolbar {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 12px; gap: 10px;
}
.toolbar-right { display: flex; align-items: center; gap: 10px; }

.btn-tambah {
  display: flex; align-items: center; gap: 6px;
  padding: 9px 18px; background: #1B7A6E; border: none; border-radius: 9px;
  font-family: 'Montserrat', sans-serif; font-size: 12px; font-weight: 700;
  color: #FFFFFF; cursor: pointer; white-space: nowrap;
  box-shadow: 0 3px 10px rgba(27,122,110,0.25);
  transition: background 0.2s, transform 0.15s;
}
.btn-tambah:hover { background: #166058; transform: translateY(-1px); }

.ctrl-wrap { display: flex; align-items: center; gap: 6px; }
.ctrl-label { font-size: 11px; font-weight: 600; color: #999; white-space: nowrap; }

.select-wrap { position: relative; }
.ctrl-select {
  appearance: none; -webkit-appearance: none;
  padding: 7px 28px 7px 10px;
  background: #FFFFFF; border: 1.5px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 11px; font-weight: 600;
  color: #444; outline: none; cursor: pointer; transition: border-color 0.2s;
}
.ctrl-select:focus { border-color: #1B7A6E; }
.select-arrow {
  position: absolute; right: 8px; top: 50%;
  transform: translateY(-50%); color: #999; pointer-events: none;
}

/* ── Table ── */
.table-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  overflow: hidden;         
  width: 100%;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  table-layout: fixed;      
}

.data-table th:nth-child(1)  { width: 4%; }   /* No */
.data-table th:nth-child(2)  { width: 11%; }  /* Nama Perusahaan */
.data-table th:nth-child(3)  { width: 9%; }   /* Jenis Proyek */
.data-table th:nth-child(4)  { width: 7%; }   /* Luas Area */
.data-table th:nth-child(5)  { width: 9%; }   /* Jenis Pekerjaan */
.data-table th:nth-child(6)  { width: 10%; }  /* Spesifikasi Design */
.data-table th:nth-child(7)  { width: 7%; }   /* Area Layanan */
.data-table th:nth-child(8)  { width: 11%; }  /* Daftar Item */
.data-table th:nth-child(9)  { width: 7%; }  /* Spesifikasi Item */
.data-table th:nth-child(10) { width: 10%; }  /* Harga Satuan */
.data-table th:nth-child(11) { width: 7%; }   /* Tanggal */
.data-table th:nth-child(12) { width: 9%; }   /* Aksi */
.data-table thead tr {
  background: #FFFFFF; border-bottom: 1.5px solid #EBEBEB;
}
.data-table th {
  padding: 11px 8px;
  text-align: left; font-size: 10px; font-weight: 700;
  color: #AAAAAA; white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
}

.data-table tbody tr {
  border-bottom: 1px solid #F3F3F3; transition: background 0.15s;
}
.data-table tbody tr:last-child { border-bottom: none; }
.data-table tbody tr:hover { background: #F8FDFC; }

.data-table td {
  padding: 11px 8px;
  font-size: 11px; color: #444; vertical-align: middle;
  overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
}

.td-no    { color: #BBBBBB; font-weight: 500; padding-left: 16px !important; }.td-main  { font-weight: 700; color: #1A1A2E; }
.td-cell  { }
.td-muted { color: #AAAAAA; }
.td-nowrap { white-space: nowrap; }
.td-truncate { overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
.td-aksi { text-align: left; white-space: nowrap; padding: 11px 8px 11px 0 !important; overflow: visible; }/* Badge */
.badge {
  display: inline-block; padding: 3px 9px;
  border-radius: 20px; font-size: 10px; font-weight: 700; white-space: nowrap;
}
.badge--renovasi { background: #F0F0F0; color: #777; }
.badge--desain   { background: #E6F9E6; color: #27AE60; }
.badge--custom   { background: #FFF8E1; color: #E6A817; }

/* Action buttons */
.action-btn {
  display: inline-flex; align-items: center; justify-content: center;
  width: 26px; height: 26px; border: none; border-radius: 6px;
  cursor: pointer; margin-left: 2px; background: transparent;
  transition: background 0.15s, transform 0.15s; color: #CCCCCC;
}
.action-btn:hover { transform: scale(1.1); }
.action-btn--edit:hover   { color: #1B7A6E; background: #EBF7F5; }
.action-btn--delete:hover { color: #E53E3E; background: #FEF0F0; }

/* Empty */
.empty-row { text-align: center; padding: 48px 0 !important; }
.empty-state { display: flex; flex-direction: column; align-items: center; gap: 12px; }
.empty-state p { font-size: 12px; color: #BBBBBB; line-height: 1.6; text-align: center; }
.empty-state strong { color: #1B7A6E; }

/* ── Form Card ── */
.form-card {
  background: #FFFFFF; border-radius: 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06); padding: 24px 28px 28px;
}
.form-card__heading { display: flex; align-items: center; gap: 10px; margin-bottom: 24px; }
.back-btn {
  display: flex; align-items: center; justify-content: center;
  background: none; border: none; cursor: pointer; color: #1B7A6E;
  padding: 4px; border-radius: 6px; transition: background 0.15s;
}
.back-btn:hover { background: #F0FAF8; }
.form-title { font-size: 16px; font-weight: 700; color: #1B7A6E; }

.form-body { display: flex; flex-direction: column; gap: 18px; }
.field { display: flex; flex-direction: column; gap: 6px; }
.field__label { font-size: 12px; font-weight: 600; color: #1B7A6E; }
.field__optional { font-weight: 400; color: #AAAAAA; font-size: 11px; }

.field__input {
  width: 100%; padding: 12px 16px;
  border: 1.5px solid #E8E8E8; border-radius: 9px;
  font-family: 'Montserrat', sans-serif; font-size: 13px;
  color: #2C2C2C; background: #FAFAFA; outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}
.field__input::placeholder { color: #BBBBBB; }
.field__input:focus { border-color: #1B7A6E; background: #FFFFFF; box-shadow: 0 0 0 3px rgba(27,122,110,0.08); }

.input-suffix-wrap { position: relative; }
.field__input--suffix { padding-right: 44px; }
.input-suffix { position: absolute; right: 14px; top: 50%; transform: translateY(-50%); font-size: 12px; font-weight: 600; color: #999; pointer-events: none; }

.input-prefix-wrap { position: relative; display: flex; align-items: center; }
.input-prefix { position: absolute; left: 16px; font-size: 13px; font-weight: 600; color: #2C2C2C; pointer-events: none; z-index: 1; }
.field__input--prefix { padding-left: 40px; }

.field__select-wrap { position: relative; }
.field__select {
  width: 100%; padding: 12px 44px 12px 16px;
  border: 1.5px solid #E8E8E8; border-radius: 9px;
  font-family: 'Montserrat', sans-serif; font-size: 13px;
  color: #BBBBBB; background: #FAFAFA; outline: none;
  appearance: none; -webkit-appearance: none; cursor: pointer;
  transition: border-color 0.2s, box-shadow 0.2s;
}
.field__select.filled { color: #2C2C2C; }
.field__select:focus { border-color: #1B7A6E; background: #FFFFFF; box-shadow: 0 0 0 3px rgba(27,122,110,0.08); }
.field__select-arrow { position: absolute; right: 14px; top: 50%; transform: translateY(-50%); color: #999; pointer-events: none; }

/* Multi Select */
.multiselect { position: relative; }
.multiselect__trigger {
  display: flex; align-items: center; justify-content: space-between;
  padding: 12px 16px; border: 1.5px solid #E8E8E8; border-radius: 9px;
  background: #FAFAFA; cursor: pointer; min-height: 46px;
  transition: border-color 0.2s;
}
.multiselect__trigger:hover { border-color: #1B7A6E; }
.multiselect__placeholder { font-size: 13px; color: #BBBBBB; }
.multiselect__value { font-size: 13px; color: #2C2C2C; font-weight: 500; }
.multiselect__arrow { flex-shrink: 0; color: #999; transition: transform 0.2s ease; }
.multiselect__dropdown {
  position: absolute; top: calc(100% + 5px); left: 0; right: 0;
  background: #FFFFFF; border: 1.5px solid #E8E8E8; border-radius: 9px;
  box-shadow: 0 8px 20px rgba(0,0,0,0.1); z-index: 200;
  max-height: 200px; overflow-y: auto; padding: 4px 0;
}
.multiselect__option {
  display: flex; align-items: center; gap: 10px;
  padding: 9px 16px; cursor: pointer; font-size: 13px; color: #333;
  transition: background 0.15s;
}
.multiselect__option:hover { background: #F0FAF8; }
.multiselect__checkbox { width: 15px; height: 15px; accent-color: #1B7A6E; cursor: pointer; flex-shrink: 0; }

/* Form Actions */
.form-actions { display: flex; justify-content: flex-end; align-items: center; gap: 14px; margin-top: 6px; }
.btn-cancel {
  padding: 10px 28px; background: transparent; border: none;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 600;
  color: #888; cursor: pointer; border-radius: 9px; transition: color 0.2s, background 0.2s;
}
.btn-cancel:hover { color: #555; background: #F4F5F7; }
.btn-upload {
  padding: 11px 36px; background: #1B7A6E; border: none; border-radius: 9px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  color: #FFFFFF; cursor: pointer;
  box-shadow: 0 4px 12px rgba(27,122,110,0.28);
  transition: background 0.2s, transform 0.15s;
}
.btn-upload:hover { background: #166058; transform: translateY(-1px); }
.btn-upload:active { transform: translateY(0); }
</style>