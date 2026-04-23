<template>
  <div class="training-page">

    <!-- ── HEADER ── -->
    <div class="page-header">
      <h1 class="page-title">Training Model</h1>
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
         VIEW 1: MAIN — RINGKASAN + RIWAYAT
    ══════════════════════════════════════ -->
    <template v-if="currentView === 'main'">

      <!-- Row 1: Ringkasan + Train Model -->
      <div class="top-row">

        <!-- Ringkasan Dataset -->
        <div class="ringkasan-card">
          <h2 class="card-title card-title--green">Ringkasan Dataset</h2>
          <div class="ringkasan-body">
            <p class="ringkasan-item">Total Data Proyek: <strong>{{ ringkasan.totalData }}</strong></p>
            <p class="ringkasan-item">Data Siap Training: <strong>{{ ringkasan.dataTraining }}</strong></p>
            <p class="ringkasan-item">Variabel Input: <strong>{{ ringkasan.variabelInput }}</strong></p>
            <p class="ringkasan-item">Update Data Terakhir: <strong>{{ ringkasan.updateTerakhir }}</strong></p>
          </div>
          <div class="ringkasan-footer">
            <button class="btn-detail" @click="currentView = 'evaluasi'">Detail</button>
          </div>
        </div>

        <!-- Train Model -->
        <div class="train-card">
          <h2 class="train-card__title">Train Model</h2>
          <button class="btn-train btn-train--gold" @click="handleLatih" :disabled="isTraining">
            <span v-if="isTraining && trainingAction === 'latih'" class="spinner"></span>
            <span v-else>Latih</span>
          </button>
          <button class="btn-train btn-train--gold" @click="handleLatihUlang" :disabled="isTraining">
            <span v-if="isTraining && trainingAction === 'ulang'" class="spinner"></span>
            <span v-else>Latih Ulang</span>
          </button>
          <button class="btn-reset-txt" @click="handleReset">Reset</button>
        </div>

      </div>

      <!-- Riwayat Training -->
      <div class="section-block">
        <div class="section-header">
          <h2 class="section-title">Riwayat Training</h2>
          <button class="btn-view-all">
            View all
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"/></svg>
          </button>
        </div>

        <div class="table-card">
          <table class="data-table">
            <thead>
              <tr>
                <th style="width:5%">No</th>
                <th style="width:28%">Nama &amp; ID Training</th>
                <th style="width:14%">Status</th>
                <th style="width:30%">Nilai MAE</th>
                <th style="width:14%">Tanggal Training</th>
                <th style="width:9%"></th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, idx) in riwayatList" :key="item.id">
                <td class="td-no">{{ idx + 1 }}</td>
                <td>
                  <p class="td-nama">{{ item.nama }}</p>
                  <p class="td-id">{{ item.trainingId }}</p>
                </td>
                <td>
                  <span class="status-badge" :class="statusClass(item.status)">
                    <span class="status-dot"></span>
                    {{ item.status }}
                  </span>
                </td>
                <td class="td-mae">{{ item.nilaiMAE }}</td>
                <td class="td-tgl td-muted">{{ formatTanggal(item.tanggal) }}</td>
                <td class="td-aksi">
                  <button class="action-btn action-btn--edit" title="Edit">
                    <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2">
                      <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                      <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                    </svg>
                  </button>
                  <button class="action-btn action-btn--delete" title="Hapus" @click="deleteRiwayat(item.id)">
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
      </div>

      <!-- Hasil Training Terbaru -->
      <div class="section-block">
        <h2 class="section-title">Hasil Training Terbaru</h2>

        <div class="hasil-card">
          <h3 class="hasil-card__title">Hasil Training</h3>
          <div class="hasil-card__body">
            <p class="hasil-item">Nilai MAE: <strong>{{ hasilTerbaru.nilaiMAE }}</strong></p>
            <p class="hasil-item">Status Akurasi: <strong>{{ hasilTerbaru.statusAkurasi }}</strong></p>
            <p class="hasil-item">Status Model: <strong>{{ hasilTerbaru.statusModel }}</strong></p>
          </div>
          <div class="hasil-card__footer">
            <button class="btn-evaluasi" @click="currentView = 'evaluasi'">Evaluasi</button>
          </div>
        </div>
      </div>

    </template>

    <!-- ══════════════════════════════════════
         VIEW 2: EVALUASI HASIL TRAINING MODEL
    ══════════════════════════════════════ -->
    <template v-if="currentView === 'evaluasi'">

      <div class="view-heading">
        <button class="back-btn" @click="currentView = 'main'">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <h2 class="view-title">Evaluasi Hasil Traing Model</h2>
      </div>

      <div class="evaluasi-layout">

        <!-- Kiri: Evaluasi detail -->
        <div class="evaluasi-main-card">

          <section class="eval-section">
            <h3 class="eval-section-title">Hasil Evaluasi</h3>
            <p class="eval-item">Nilai MAE: <strong>{{ hasilTerbaru.nilaiMAE }}</strong></p>
            <p class="eval-item">Status Akurasi: <strong>{{ hasilTerbaru.statusAkurasi }}</strong></p>
            <p class="eval-item">Status Model: <strong>{{ hasilTerbaru.statusModel }}</strong></p>
          </section>

          <section class="eval-section">
            <h3 class="eval-section-title">Ringkasan Dataset</h3>
            <p class="eval-item">Total Data Proyek: <strong>{{ ringkasan.totalData }}</strong></p>
            <p class="eval-item">Data Training: <strong>{{ ringkasan.dataTraining }}</strong></p>
            <p class="eval-item">Data Validasi: <strong>{{ ringkasan.dataValidasi }}</strong></p>
            <p class="eval-item">Variabel Input: <strong>{{ ringkasan.variabelInput }}</strong></p>
            <p class="eval-item">Tanggal Training Terakhir: <strong>{{ ringkasan.updateTerakhir }}</strong></p>
          </section>

          <section class="eval-section">
            <h3 class="eval-section-title">Informasi Akurasi</h3>
            <p class="eval-desc">
              Model dievaluasi menggunakan Mean Absolute Error (MAE).
              Semakin kecil nilai MAE, semakin baik akurasi model. Hasil
              evaluasi ini digunakan untuk menilai kelayakan model dalam
              proses estimasi biaya proyek
            </p>
          </section>

          <div class="eval-footer">
            <button class="btn-simpan-eval" @click="simpanEvaluasi">Simpan</button>
          </div>
        </div>

        <!-- Kanan: Train Model -->
        <div class="train-card train-card--eval">
          <h2 class="train-card__title">Train Model</h2>
          <button class="btn-train btn-train--gold" @click="handleGunakan">Gunakan Model</button>
          <button class="btn-train btn-train--gold" @click="handleLatihUlang">Latih Ulang</button>
          <button class="btn-reset-txt" @click="currentView = 'main'">Batal</button>
        </div>

      </div>
    </template>

  </div>
</template>

<script setup>
import { ref } from 'vue'

// ─────────────────────────────────────────────
// API CONFIG 
// ─────────────────────────────────────────────
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

// const fetchRingkasan    = () => fetch(`${API_BASE_URL}/training/ringkasan`).then(r => r.json())
// const fetchRiwayat      = () => fetch(`${API_BASE_URL}/training/riwayat`).then(r => r.json())
// const fetchHasilTerbaru = () => fetch(`${API_BASE_URL}/training/hasil-terbaru`).then(r => r.json())
// const postLatih         = () => fetch(`${API_BASE_URL}/training/latih`,       { method: 'POST' }).then(r => r.json())
// const postLatihUlang    = () => fetch(`${API_BASE_URL}/training/latih-ulang', { method: 'POST' }).then(r => r.json())
// const postReset         = () => fetch(`${API_BASE_URL}/training/reset`,       { method: 'POST' })
// const postGunakan       = () => fetch(`${API_BASE_URL}/training/gunakan`,     { method: 'POST' })
// const postSimpanEval    = () => fetch(`${API_BASE_URL}/training/simpan-eval`, { method: 'POST' })
// const deleteTraining    = (id) => fetch(`${API_BASE_URL}/training/${id}`,     { method: 'DELETE' })

// ─────────────────────────────────────────────
// STATE
// ─────────────────────────────────────────────
const currentView   = ref('main')
const searchQuery   = ref('')
const isTraining    = ref(false)
const trainingAction = ref('')
const avatarImg     = ref(null)
const avatarFallback = ref(null)

// ─────────────────────────────────────────────
// DUMMY DATA 
// ─────────────────────────────────────────────
const ringkasan = ref({
  totalData:     128,
  dataTraining:  120,
  dataValidasi:  8,
  variabelInput: 'Luas Ruangan, Tingkat Kerumitan, Durasi',
  updateTerakhir:'12 Mei 2026',
})

const hasilTerbaru = ref({
  nilaiMAE:     'Rp 1.250.000',
  statusAkurasi:'Baik',
  statusModel:  'Siap Digunakan',
})

const riwayatList = ref([
  { id: 1, nama: 'Training-01', trainingId: '123456789', status: 'Active',  nilaiMAE: 'jessica.hanson@example.com',  tanggal: '2015-05-27' },
  { id: 2, nama: 'Training-02', trainingId: '123456789', status: 'Offline', nilaiMAE: 'willie.jennings@example.com', tanggal: '2012-05-19' },
  { id: 3, nama: 'Training-03', trainingId: '123456789', status: 'Offline', nilaiMAE: 'd.chambers@example.com',      tanggal: '2016-03-04' },
  { id: 4, nama: 'Training-04', trainingId: '123456789', status: 'Deny',   nilaiMAE: 'willie.jennings@example.com', tanggal: '2016-03-04' },
  { id: 5, nama: 'Training-05', trainingId: '123456789', status: 'Offline', nilaiMAE: 'willie.jennings@example.com', tanggal: '2013-07-27' },
])

// ─────────────────────────────────────────────
// ACTIONS — siapkan untuk integrasi ML backend
// ─────────────────────────────────────────────

/**
 * handleLatih — kirim request training model baru ke Golang/ML service
 */
const handleLatih = async () => {
  trainingAction.value = 'latih'
  isTraining.value = true
  try {
    // ── DUMMY ──────────────────────────────────────────────────────
    await new Promise(r => setTimeout(r, 1500))
    const newEntry = {
      id: Date.now(),
      nama: `Training-0${riwayatList.value.length + 1}`,
      trainingId: String(Math.floor(Math.random() * 900000000) + 100000000),
      status: 'Active',
      nilaiMAE: 'new.training@example.com',
      tanggal: new Date().toISOString().split('T')[0],
    }
    // Set semua yang sebelumnya Active ke Offline
    riwayatList.value = riwayatList.value.map(r => ({ ...r, status: r.status === 'Active' ? 'Offline' : r.status }))
    riwayatList.value.unshift(newEntry)
    // ── END DUMMY ──────────────────────────────────────────────────

    // ── API  ─────────────────────────
    // const res = await postLatih()
    // await loadRiwayat()
    // hasilTerbaru.value = res.hasil
    // ── END API ────────────────────────────────────────────────────
  } catch (err) {
    console.error('Latih error:', err)
  } finally {
    isTraining.value = false
    trainingAction.value = ''
  }
}

/**
 * handleLatihUlang — retrain model dari awal
 */
const handleLatihUlang = async () => {
  trainingAction.value = 'ulang'
  isTraining.value = true
  try {
    // ── DUMMY ──────────────────────────────────────────────────────
    await new Promise(r => setTimeout(r, 1500))
    // ── END DUMMY ──────────────────────────────────────────────────

    // ── API  ─────────────────────────
    // await postLatihUlang()
    // await loadRiwayat()
    // ── END API ────────────────────────────────────────────────────
  } catch (err) {
    console.error('Latih ulang error:', err)
  } finally {
    isTraining.value = false
    trainingAction.value = ''
  }
}

/**
 * handleReset — reset model ke state default
 */
const handleReset = async () => {
  // ── DUMMY ──────────────────────────────────────────────────────
  riwayatList.value = riwayatList.value.map(r => ({ ...r, status: 'Offline' }))
  // ── END DUMMY ──────────────────────────────────────────────────

  // ── API  ─────────────────────────
  // await postReset()
  // await loadRiwayat()
  // ── END API ────────────────────────────────────────────────────
}

/**
 * handleGunakan — aktifkan model yang sudah dievaluasi
 */
const handleGunakan = async () => {
  // ── DUMMY ──────────────────────────────────────────────────────
  alert('Model diaktifkan!')
  // ── END DUMMY ──────────────────────────────────────────────────

  // ── API  ─────────────────────────
  // await postGunakan()
  // ── END API ────────────────────────────────────────────────────
}

/**
 * simpanEvaluasi — simpan hasil evaluasi ke backend
 */
const simpanEvaluasi = async () => {
  // ── DUMMY ──────────────────────────────────────────────────────
  currentView.value = 'main'
  // ── END DUMMY ──────────────────────────────────────────────────

  // ── API  ─────────────────────────
  // await postSimpanEval()
  // currentView.value = 'main'
  // ── END API ────────────────────────────────────────────────────
}

/**
 * deleteRiwayat — hapus entri riwayat training
 */
const deleteRiwayat = async (id) => {
  // ── DUMMY ──────────────────────────────────────────────────────
  riwayatList.value = riwayatList.value.filter(r => r.id !== id)
  // ── END DUMMY ──────────────────────────────────────────────────

  // ── API ─────────────────────────
  // await deleteTraining(id)
  // await loadRiwayat()
  // ── END API ────────────────────────────────────────────────────
}

// ─────────────────────────────────────────────
// LOAD DATA
// ─────────────────────────────────────────────
// panggil saat component mount untuk fetch dari backend
// const loadRingkasan    = async () => { ringkasan.value    = await fetchRingkasan()    }
// const loadRiwayat      = async () => { riwayatList.value  = await fetchRiwayat()      }
// const loadHasilTerbaru = async () => { hasilTerbaru.value = await fetchHasilTerbaru() }
// onMounted(() => Promise.all([loadRingkasan(), loadRiwayat(), loadHasilTerbaru()]))

// ─────────────────────────────────────────────
// HELPERS
// ─────────────────────────────────────────────
const statusClass = (s) => {
  if (s === 'Active')  return 'status--active'
  if (s === 'Deny')    return 'status--deny'
  return 'status--offline'
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
.training-page {
  padding: 24px 28px;
  min-height: 100vh;
  background: #F4F5F7;
  font-family: 'Montserrat', sans-serif;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* ── Header ── */
.page-header {
  display: flex; align-items: center; justify-content: space-between;
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
  background: #1B7A6E; display: flex; align-items: center; justify-content: center;
}
.avatar img { width: 100%; height: 100%; object-fit: cover; }
.avatar-fallback { display: none; align-items: center; justify-content: center; font-size: 13px; font-weight: 700; color: #fff; }

/* ── Top Row ── */
.top-row {
  display: grid;
  grid-template-columns: 1fr 240px;
  gap: 16px;
  align-items: start;
}

/* ── Ringkasan Card ── */
.ringkasan-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.06);
  padding: 24px 26px 20px;
  display: flex; flex-direction: column;
}
.card-title { font-size: 16px; font-weight: 700; margin-bottom: 18px; }
.card-title--green { color: #1B7A6E; }

.ringkasan-body { display: flex; flex-direction: column; gap: 10px; flex: 1; }
.ringkasan-item { font-size: 14px; color: #333; margin: 0; }
.ringkasan-item strong { font-weight: 600; }

.ringkasan-footer { display: flex; justify-content: flex-end; margin-top: 22px; }
.btn-detail {
  padding: 10px 28px;
  background: #C8A135; border: none; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  color: #FFFFFF; cursor: pointer; transition: background 0.2s;
}
.btn-detail:hover { background: #A8841C; }

/* ── Train Model Card ── */
.train-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.06);
  padding: 22px 18px;
  display: flex; flex-direction: column; align-items: center; gap: 12px;
}
.train-card--eval {
  align-self: start;
}
.train-card__title {
  font-size: 15px; font-weight: 700; color: #1A1A2E;
  margin-bottom: 4px;
}
.btn-train {
  width: 100%; padding: 13px;
  border: none; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  cursor: pointer; display: flex; align-items: center; justify-content: center; gap: 8px;
  transition: background 0.2s, transform 0.15s;
}
.btn-train--gold { background: #C8A135; color: #FFFFFF; }
.btn-train--gold:hover { background: #A8841C; transform: translateY(-1px); }
.btn-train:disabled { opacity: 0.6; cursor: not-allowed; transform: none; }

.btn-reset-txt {
  background: none; border: none;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 500;
  color: #AAAAAA; cursor: pointer; padding: 4px 0;
  transition: color 0.2s;
}
.btn-reset-txt:hover { color: #555; }

/* ── Section Block ── */
.section-block { display: flex; flex-direction: column; gap: 14px; }
.section-header { display: flex; align-items: center; justify-content: space-between; }
.section-title { font-size: 16px; font-weight: 700; color: #1A1A2E; }

.btn-view-all {
  display: flex; align-items: center; gap: 4px;
  background: none; border: none;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 600;
  color: #333; cursor: pointer; transition: color 0.2s;
}
.btn-view-all:hover { color: #1B7A6E; }

/* ── Riwayat Table ── */
.table-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.06);
  overflow: hidden;
}
.data-table { width: 100%; border-collapse: collapse; table-layout: fixed; }
.data-table thead tr { border-bottom: 1.5px solid #F0F0F0; }
.data-table th {
  padding: 12px 16px;
  text-align: left; font-size: 11px; font-weight: 600; color: #AAAAAA;
}
.data-table tbody tr { border-bottom: 1px solid #F6F6F6; transition: background 0.15s; }
.data-table tbody tr:last-child { border-bottom: none; }
.data-table tbody tr:hover { background: #F8FDFC; }
.data-table td { padding: 14px 16px; font-size: 12px; color: #444; vertical-align: middle; }

.td-no  { color: #BBBBBB; }
.td-nama { font-weight: 700; color: #1A1A2E; font-size: 13px; margin: 0 0 2px; }
.td-id   { font-size: 11px; color: #AAAAAA; margin: 0; }
.td-mae  { color: #555; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
.td-tgl  { white-space: nowrap; }
.td-muted { color: #AAAAAA; }
.td-aksi { text-align: right; white-space: nowrap; overflow: visible; padding-right: 12px !important; }

/* Status badge */
.status-badge {
  display: inline-flex; align-items: center; gap: 5px;
  padding: 4px 10px; border-radius: 20px;
  font-size: 11px; font-weight: 600; white-space: nowrap;
}
.status-dot { width: 6px; height: 6px; border-radius: 50%; background: currentColor; flex-shrink: 0; }
.status--active  { background: #E8F9EF; color: #27AE60; }
.status--offline { background: #F2F2F2; color: #888888; }
.status--deny    { background: #FEF0F0; color: #E53E3E; }

/* Action buttons */
.action-btn {
  display: inline-flex; align-items: center; justify-content: center;
  width: 28px; height: 28px; border: none; border-radius: 6px;
  cursor: pointer; margin-left: 4px; background: transparent;
  transition: background 0.15s, transform 0.15s; color: #CCCCCC;
}
.action-btn:hover { transform: scale(1.1); }
.action-btn--edit:hover   { color: #1B7A6E; background: #EBF7F5; }
.action-btn--delete:hover { color: #E53E3E; background: #FEF0F0; }

/* ── Hasil Training Terbaru ── */
.hasil-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.06);
  padding: 24px 26px 20px;
}
.hasil-card__title {
  font-size: 15px; font-weight: 700; color: #1A1A2E; margin-bottom: 16px;
}
.hasil-card__body { display: flex; flex-direction: column; gap: 10px; margin-bottom: 24px; }
.hasil-item { font-size: 14px; color: #333; margin: 0; }
.hasil-item strong { font-weight: 600; }
.hasil-card__footer { display: flex; justify-content: flex-end; }

.btn-evaluasi {
  padding: 11px 32px;
  background: #C8A135; border: none; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  color: #FFFFFF; cursor: pointer; transition: background 0.2s;
}
.btn-evaluasi:hover { background: #A8841C; }

/* ── Evaluasi View ── */
.view-heading {
  display: flex; align-items: center; gap: 10px;
}
.back-btn {
  display: flex; align-items: center; justify-content: center;
  background: none; border: none; cursor: pointer; color: #1B7A6E;
  padding: 4px; border-radius: 6px; transition: background 0.15s;
}
.back-btn:hover { background: #F0FAF8; }
.view-title { font-size: 16px; font-weight: 700; color: #1B7A6E; }

.evaluasi-layout {
  display: grid;
  grid-template-columns: 1fr 240px;
  gap: 16px;
  align-items: start;
}

.evaluasi-main-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.06);
  padding: 28px 30px 24px;
  display: flex; flex-direction: column; gap: 28px;
}

.eval-section { display: flex; flex-direction: column; gap: 10px; }
.eval-section-title { font-size: 15px; font-weight: 700; color: #1A1A2E; margin-bottom: 2px; }
.eval-item { font-size: 14px; color: #333; margin: 0; }
.eval-item strong { font-weight: 600; }
.eval-desc { font-size: 14px; color: #555; line-height: 1.7; margin: 0; }

.eval-footer { display: flex; justify-content: flex-end; margin-top: 4px; }
.btn-simpan-eval {
  padding: 11px 36px;
  background: #C8A135; border: none; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 13px; font-weight: 700;
  color: #FFFFFF; cursor: pointer; transition: background 0.2s;
}
.btn-simpan-eval:hover { background: #A8841C; }

/* ── Spinner ── */
.spinner {
  display: inline-block; width: 14px; height: 14px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff; border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>