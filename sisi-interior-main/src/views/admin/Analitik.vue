<template>
  <div class="analitik-page">

    <!-- ── HEADER ── -->
    <div class="page-header">
      <h1 class="page-title">Dashboard Analitik</h1>
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

    <!-- ── ROW 1: STAT CARDS ── -->
    <div class="stats-row">
      <div class="stat-card" v-for="stat in stats" :key="stat.key">
        <p class="stat-label">{{ stat.label }}</p>
        <p class="stat-value" :class="{ 'stat-value--gold': stat.gold }">{{ stat.value }}</p>
        <p class="stat-sub" :class="{ 'stat-sub--link': stat.linkSub }">{{ stat.sub }}</p>
      </div>
    </div>

    <!-- ── ROW 2: GRAFIK + KATEGORI ── -->
    <div class="row-two">

      <!-- Grafik Tren -->
      <div class="chart-card">
        <div class="chart-card__header">
          <h2 class="section-title">Grafik Tren Biaya Proyek</h2>
          <div class="select-wrap">
            <select v-model="chartPeriod" class="period-select">
              <option value="this_month">This Month</option>
              <option value="last_month">Last Month</option>
              <option value="this_year">This Year</option>
            </select>
            <svg class="select-arrow" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
          </div>
        </div>
        <div class="chart-wrap">
          <canvas ref="chartCanvas"></canvas>
        </div>
      </div>

      <!-- Kategori Proyek -->
      <div class="side-card">
        <h2 class="section-title">Kategori Proyek</h2>
        <div class="kategori-list">
          <div class="kategori-row" v-for="k in kategoriProyek" :key="k.nama">
            <span class="kategori-nama">{{ k.nama }}</span>
            <span class="kategori-pct">{{ k.pct }}%</span>
          </div>
        </div>
      </div>

    </div>

    <!-- ── ROW 3: PERFORMA + TINGKAT KERUMITAN ── -->
    <div class="row-two">

      <!-- Performa / Ringkasan Proyek -->
      <div class="chart-card">
        <div class="chart-card__header">
          <h2 class="section-title">Performa / Ringkasan Proyek</h2>
          <button class="btn-see-all">See All</button>
        </div>
        <table class="performa-table">
          <thead>
            <tr>
              <th>Nama Proyek</th>
              <th>Jenis</th>
              <th>Biaya</th>
              <th>Durasi</th>
              <th>Tanggal</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in performaProyek" :key="p.nama">
              <td class="td-nama">{{ p.nama }}</td>
              <td class="td-jenis">{{ p.jenis }}</td>
              <td class="td-biaya">{{ p.biaya }}</td>
              <td class="td-durasi">{{ p.durasi }}</td>
              <td class="td-tanggal">{{ p.tanggal }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Tingkat Kerumitan -->
      <div class="side-card">
        <h2 class="section-title">Tingkat Kerumitan</h2>
        <div class="kerumitan-list">
          <div class="kerumitan-row" v-for="k in tingkatKerumitan" :key="k.label">
            <span class="kerumitan-dot" :style="{ background: k.color }"></span>
            <span class="kerumitan-label">{{ k.label }}</span>
            <span class="kerumitan-pct">{{ k.pct }}%</span>
          </div>
        </div>
      </div>

    </div>

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

// ─────────────────────────────────────────────
// API CONFIG — ganti BASE_URL ke endpoint Golang
// ─────────────────────────────────────────────
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

// TODO: uncomment saat backend Golang siap
// const fetchStats          = () => fetch(`${API_BASE_URL}/analitik/stats`).then(r => r.json())
// const fetchTrenBiaya      = (period) => fetch(`${API_BASE_URL}/analitik/tren?period=${period}`).then(r => r.json())
// const fetchKategoriProyek = () => fetch(`${API_BASE_URL}/analitik/kategori`).then(r => r.json())
// const fetchPerforma       = () => fetch(`${API_BASE_URL}/analitik/performa`).then(r => r.json())
// const fetchKerumitan      = () => fetch(`${API_BASE_URL}/analitik/kerumitan`).then(r => r.json())

// ─────────────────────────────────────────────
// STATE
// ─────────────────────────────────────────────
const searchQuery  = ref('')
const chartPeriod  = ref('this_month')
const chartCanvas  = ref(null)
const avatarImg    = ref(null)
const avatarFallback = ref(null)
let chartInstance  = null

// ─────────────────────────────────────────────
// DUMMY DATA 
// ─────────────────────────────────────────────
const stats = ref([
  { key: 'total',   label: 'Total Proyek',         value: '128',            sub: 'Proyek',               gold: true,  linkSub: true  },
  { key: 'rata',    label: 'Rata-rata Biaya Proyek',value: 'Rp 18.500.000', sub: 'Dalam 1 Bulan terakhir', gold: true,  linkSub: false },
  { key: 'tinggi',  label: 'Biaya Tertinggi',       value: 'Rp 45.000.000', sub: 'Per proyek',            gold: true,  linkSub: false },
  { key: 'rendah',  label: 'Biaya Terendah',        value: 'Rp 6.500.000',  sub: 'Per proyek',            gold: true,  linkSub: false },
])

// Dummy tren biaya — key: period value
const trenData = {
  this_month: {
    labels: ['10','11','12','13','14','15','16','17','18','19','20','21','22','23','24'],
    values: [40000,42000,45000,43000,44000,46000,83234,68000,62000,70000,66000,71000,58000,62000,60000],
    peak: { index: 6, label: '83,234' },
  },
  last_month: {
    labels: ['1','3','5','7','9','11','13','15','17','19','21','23','25','27','29'],
    values: [35000,38000,36000,42000,40000,45000,43000,60000,55000,58000,52000,48000,50000,46000,44000],
    peak: { index: 7, label: '60,000' },
  },
  this_year: {
    labels: ['Jan','Feb','Mar','Apr','Mei','Jun','Jul','Agu','Sep','Okt','Nov','Des'],
    values: [50000,55000,62000,58000,70000,65000,80000,75000,78000,72000,68000,85000],
    peak: { index: 11, label: '85,000' },
  },
}

const kategoriProyek = ref([
  { nama: 'Desain Interior', pct: 96.42 },
  { nama: 'Konstraktor',     pct: 2.76  },
  { nama: 'Custom Furniture',pct: 0.82  },
])

const performaProyek = ref([
  { nama: 'Kitchen Set', jenis: 'Desain Interior', biaya: '12 jt', durasi: 10, tanggal: '10 Maret 2026' },
  { nama: 'Cafe',        jenis: 'Desain Interior', biaya: '8 jt',  durasi: 5,  tanggal: '12 Maret 2026' },
  { nama: 'Bathroom',    jenis: 'Desain Interior', biaya: '10 jt', durasi: 12, tanggal: '1 Maret 2026'  },
])

const tingkatKerumitan = ref([
  { label: 'Rendah', pct: 25, color: '#6DD69A' },
  { label: 'Sedang', pct: 50, color: '#E8C96A' },
  { label: 'Tinggi', pct: 25, color: '#E85C5C' },
])

// ─────────────────────────────────────────────
// CHART.JS
// ─────────────────────────────────────────────
const buildChart = () => {
  if (!chartCanvas.value) return
  if (chartInstance) { chartInstance.destroy(); chartInstance = null }

  const d = trenData[chartPeriod.value]

  // Annotation: vertical dashed line at peak
  const peakIdx = d.peak.index

  // Load Chart.js from CDN — injected once
  if (!window.Chart) {
    console.warn('Chart.js belum ready')
    return
  }

  const Chart = window.Chart
  const ctx = chartCanvas.value.getContext('2d')

  // Gradient fill
  const grad = ctx.createLinearGradient(0, 0, 0, 280)
  grad.addColorStop(0, 'rgba(27,122,110,0.22)')
  grad.addColorStop(1, 'rgba(27,122,110,0)')

  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels: d.labels,
      datasets: [{
        data: d.values,
        borderColor: '#1B7A6E',
        borderWidth: 2,
        pointRadius: 0,
        pointHoverRadius: 4,
        pointHoverBackgroundColor: '#1B7A6E',
        fill: true,
        backgroundColor: grad,
        tension: 0.3,
      }],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
        // inline annotation for peak tooltip bubble
        afterDraw: undefined,

      scales: {
        x: {
          grid: { color: 'rgba(0,0,0,0.04)', drawBorder: false },
          ticks: { font: { size: 11, family: 'Montserrat' }, color: '#AAAAAA' },
          border: { display: false },
        },
        y: {
          min: 0,
          max: 120000,
          grid: { color: 'rgba(0,0,0,0.06)', drawBorder: false },
          ticks: {
            stepSize: 20000,
            font: { size: 11, family: 'Montserrat' },
            color: '#AAAAAA',
            callback: v => v >= 1000 ? (v/1000) + 'K' : v,
          },
          border: { display: false },
        },
      },
    },
    plugins: [{
      id: 'peakAnnotation',
      afterDraw(chart) {
        const { ctx, chartArea } = chart

        const meta = chart.getDatasetMeta(0)
        const point = meta.data[peakIdx]
        if (!point) return

        const x = point.x
        const yVal = point.y
        const yTop = chartArea.top
        const yBot = chartArea.bottom

        // dashed vertical line (red)
        ctx.save()
        ctx.setLineDash([4, 4])
        ctx.strokeStyle = '#E85C5C'
        ctx.lineWidth = 1.5
        ctx.beginPath()
        ctx.moveTo(x, yVal)
        ctx.lineTo(x, yBot)
        ctx.stroke()
        ctx.restore()

        // tooltip bubble
        const label = d.peak.label
        const bw = 66, bh = 24, br = 6
        const bx = x - bw / 2, by = yVal - bh - 8

        ctx.save()
        ctx.fillStyle = '#1B7A6E'
        ctx.beginPath()
        ctx.roundRect(bx, by, bw, bh, br)
        ctx.fill()

        // small triangle
        ctx.beginPath()
        ctx.moveTo(x - 5, by + bh)
        ctx.lineTo(x + 5, by + bh)
        ctx.lineTo(x, by + bh + 6)
        ctx.fill()

        ctx.fillStyle = '#FFFFFF'
        ctx.font = '600 11px Montserrat, sans-serif'
        ctx.textAlign = 'center'
        ctx.textBaseline = 'middle'
        ctx.fillText(label, x, by + bh / 2)
        ctx.restore()
      },
    }],
  })
}

// ─────────────────────────────────────────────
// LOAD DATA — untuk fetch dari Golang
// ─────────────────────────────────────────────
const loadData = async () => {
  // ── DUMMY: data sudah di-set di ref di atas ─────────────────────
  // Tidak ada fetch, langsung render chart

  // ── API (uncomment saat backend siap) ──────────────────────────
  // try {
  //   const [statsRes, kategoriRes, performaRes, kerumitanRes] = await Promise.all([
  //     fetchStats(),
  //     fetchKategoriProyek(),
  //     fetchPerforma(),
  //     fetchKerumitan(),
  //   ])
  //   stats.value          = statsRes
  //   kategoriProyek.value = kategoriRes
  //   performaProyek.value = performaRes
  //   tingkatKerumitan.value = kerumitanRes
  // } catch (err) {
  //   console.error('Analitik fetch error:', err)
  // }
}

const loadTren = async () => {
  // ── DUMMY: data sudah ada di trenData object ─────────────────────
  buildChart()

  // ── API (uncomment saat backend siap) ──────────────────────────
  // try {
  //   const res = await fetchTrenBiaya(chartPeriod.value)
  //   // update trenData[chartPeriod.value] dengan res.labels, res.values, res.peak
  //   buildChart()
  // } catch (err) {
  //   console.error('Tren fetch error:', err)
  // }
}

// ─────────────────────────────────────────────
// LIFECYCLE
// ─────────────────────────────────────────────
onMounted(() => {
  loadData()

  // Inject Chart.js CDN once
  if (!window.Chart) {
    const script = document.createElement('script')
    script.src = 'https://cdnjs.cloudflare.com/ajax/libs/Chart.js/4.4.1/chart.umd.min.js'
    script.onload = () => buildChart()
    document.head.appendChild(script)
  } else {
    buildChart()
  }
})

onUnmounted(() => { 
  if (chartInstance) {
    chartInstance.destroy()
    chartInstance = null
  }
})

watch(chartPeriod, loadTren)

// ─────────────────────────────────────────────
// HELPERS
// ─────────────────────────────────────────────
const onAvatarError = () => {
  if (avatarImg.value) avatarImg.value.style.display = 'none'
  if (avatarFallback.value) avatarFallback.value.style.display = 'flex'
}
</script>

<style scoped>
.analitik-page {
  padding: 24px 28px;
  min-height: 100vh;
  background: #F4F5F7;
  font-family: 'Montserrat', sans-serif;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* ── Header ── */
.page-header {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 4px;
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

/* ── Stat Cards Row ── */
.stats-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}
.stat-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
  padding: 20px 22px 18px;
}
.stat-label {
  font-size: 12px; font-weight: 500; color: #999;
  margin-bottom: 8px;
}
.stat-value {
  font-size: 22px; font-weight: 700; color: #1A1A2E;
  margin-bottom: 4px; line-height: 1.2;
}
.stat-value--gold { color: #C8A135; }
.stat-sub {
  font-size: 11px; color: #999;
}
.stat-sub--link { color: #C8A135; font-weight: 600; cursor: pointer; }
.stat-sub--link:hover { text-decoration: underline; }

/* ── Row layout (chart + side) ── */
.row-two {
  display: grid;
  grid-template-columns: 1fr 260px;
  gap: 16px;
  align-items: start;
}

/* ── Chart Card ── */
.chart-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
  padding: 22px 22px 18px;
}
.chart-card__header {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 20px;
}
.section-title {
  font-size: 15px; font-weight: 700; color: #1A1A2E;
}
.chart-wrap {
  position: relative; height: 220px;
}

.select-wrap { position: relative; }
.period-select {
  appearance: none; -webkit-appearance: none;
  padding: 7px 28px 7px 12px;
  background: #F0FAF8; border: 1.5px solid #C8E8E4; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 11px; font-weight: 600;
  color: #1B7A6E; outline: none; cursor: pointer;
}
.select-arrow {
  position: absolute; right: 8px; top: 50%;
  transform: translateY(-50%); color: #1B7A6E; pointer-events: none;
}

/* ── Side Card ── */
.side-card {
  background: #FFFFFF;
  border-radius: 14px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
  padding: 22px 20px;
}

/* Kategori Proyek */
.kategori-list { display: flex; flex-direction: column; gap: 20px; margin-top: 18px; }
.kategori-row { display: flex; align-items: center; justify-content: space-between; }
.kategori-nama { font-size: 13px; color: #444; }
.kategori-pct  { font-size: 14px; font-weight: 700; color: #1B7A6E; }

/* Performa table */
.performa-table { width: 100%; border-collapse: collapse; }
.performa-table th {
  padding: 8px 0; font-size: 11px; font-weight: 500;
  color: #AAAAAA; text-align: left; border-bottom: 1px solid #F0F0F0;
}
.performa-table td {
  padding: 14px 0; font-size: 12px; color: #444;
  border-bottom: 1px solid #F6F6F6; vertical-align: middle;
}
.performa-table tbody tr:last-child td { border-bottom: none; }
.td-nama   { font-weight: 700; color: #1A1A2E; }
.td-jenis  { font-weight: 600; color: #444; }
.td-biaya  { font-weight: 600; }
.td-durasi { text-align: center; }
.td-tanggal { color: #888; font-size: 11px; white-space: nowrap; }

.btn-see-all {
  padding: 7px 16px;
  background: #F4F5F7; border: 1px solid #E8E8E8; border-radius: 8px;
  font-family: 'Montserrat', sans-serif; font-size: 11px; font-weight: 600;
  color: #555; cursor: pointer; transition: background 0.2s;
}
.btn-see-all:hover { background: #EAEAEA; }

/* Tingkat Kerumitan */
.kerumitan-list { display: flex; flex-direction: column; gap: 20px; margin-top: 18px; }
.kerumitan-row { display: flex; align-items: center; gap: 12px; }
.kerumitan-dot { width: 20px; height: 20px; border-radius: 50%; flex-shrink: 0; }
.kerumitan-label { flex: 1; font-size: 13px; color: #444; }
.kerumitan-pct { font-size: 14px; font-weight: 700; color: #1B7A6E; }
</style>