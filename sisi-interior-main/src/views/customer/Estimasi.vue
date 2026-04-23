<template>
  <div class="estimasi-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">Estimasi Biaya Proyek</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">Estimasi</span>
        </div>
        <p class="page-header__desc">
          Lihat berbagai hasil karya <strong class="page-header__brand">SISI Interior</strong>
          yang dirancang untuk menghadirkan ruang yang estetik, nyaman, dan fungsional.
        </p>
      </div>
    </div>

    <!-- ── Main Content ── -->
    <section class="estimasi-section section">
      <div class="container estimasi-grid">
        <div class="estimasi-form">
          <div class="form-row">
            <div class="form-group">
              <label class="form-label">Nama Proyek</label>
              <input
                v-model="form.namaProyek"
                type="text"
                class="form-input"
                placeholder=""
              />
            </div>
            <div class="form-group">
              <label class="form-label">Luas Ruangan (m²)</label>
              <input
                v-model="form.luasRuangan"
                type="number"
                class="form-input"
                placeholder=""
                min="1"
              />
            </div>
          </div>

          <!-- Jenis Proyek -->
          <div class="form-group">
            <label class="form-label">Jenis Proyek</label>
            <div class="form-select-wrap">
              <select v-model="form.jenisProyek" class="form-select">
                <option value="">Pilih Jenis Proyek</option>
                <option value="Custom Furniture">Custom Furniture</option>
                <option value="Desain Interior">Desain Interior</option>
                <option value="Kontraktor">Kontraktor</option>
                <option value="Renovasi">Renovasi</option>
                <option value="Bangunan Baru">Bangunan Baru</option>
              </select>
              <svg class="form-select-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <!-- Tingkat Kerumitan -->
          <div class="form-group">
            <label class="form-label">Tingkat Kerumitan</label>
            <div class="form-select-wrap">
              <select v-model="form.tingkatKerumitan" class="form-select">
                <option value="">Pilih Tingkat Kerumitan</option>
                <option value="Mudah">Mudah</option>
                <option value="Sedang">Sedang</option>
                <option value="Sulit">Sulit</option>
                <option value="Sangat Sulit">Sangat Sulit</option>
              </select>
              <svg class="form-select-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <!-- Durasi Pengerjaan -->
          <div class="form-group">
            <label class="form-label">Durasi Pengerjaan (hari)</label>
            <input
              v-model="form.durasi"
              type="number"
              class="form-input"
              placeholder=""
              min="1"
            />
          </div>

          <!-- Lokasi Proyek -->
          <div class="form-group">
            <label class="form-label">Lokasi Proyek</label>
            <input
              v-model="form.lokasi"
              type="text"
              class="form-input"
              placeholder=""
            />
          </div>

          <!-- Konsep/Gaya Desain -->
          <div class="form-group">
            <label class="form-label">Konsep/Gaya Desain</label>
            <div class="form-select-wrap">
              <select v-model="form.konsep" class="form-select">
                <option value="">Pilih Konsep</option>
                <option value="Japandi">Japandi</option>
                <option value="Minimalis">Minimalis</option>
                <option value="Modern">Modern</option>
                <option value="Klasik">Klasik</option>
                <option value="Industrial">Industrial</option>
                <option value="Skandinavia">Skandinavia</option>
                <option value="Bohemian">Bohemian</option>
                <option value="Tropis">Tropis</option>
              </select>
              <svg class="form-select-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <!-- Material Khusus -->
          <div class="form-group">
            <label class="form-label">Material Khusus</label>
            <div class="form-select-wrap">
              <select v-model="form.material" class="form-select">
                <option value="">Pilih Material</option>
                <option value="Kayu Jati">Kayu Jati</option>
                <option value="Kayu Mahoni">Kayu Mahoni</option>
                <option value="Kayu Jati Muda">Kayu Jati Muda</option>
                <option value="Multiplek">Multiplek</option>
                <option value="MDF">MDF</option>
                <option value="Besi">Besi</option>
                <option value="Kaca">Kaca</option>
                <option value="Marmer">Marmer</option>
              </select>
              <svg class="form-select-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="6 9 12 15 18 9"/></svg>
            </div>
          </div>

          <!-- Kebutuhan Tambahan -->
          <div class="form-group">
            <label class="form-label">Kebutuhan Tambahan</label>
            <textarea
              v-model="form.kebutuhanTambahan"
              class="form-textarea"
              placeholder="Additional information"
              rows="4"
            ></textarea>
          </div>

          <!-- Tombol Reset + Simpan -->
          <div class="form-actions">
            <button class="btn-reset" @click="handleReset">Reset</button>
            <button class="btn-simpan" @click="handleSimpan" :disabled="isSaving">
              <span v-if="isSaving" class="btn-spinner"></span>
              <span v-else>Simpan</span>
            </button>
          </div>

        </div>

        <!-- ── Kanan: Ringkasan + Proses Estimasi ── -->
        <div class="estimasi-sidebar">

          <div class="ringkasan-card">
            <h2 class="ringkasan-card__title">Ringkasan Input</h2>

            <div class="ringkasan-card__rows">
              <div class="ringkasan-row">
                <span class="ringkasan-row__label">Luas Ruangan (m²)</span>
                <span class="ringkasan-row__value">
                  <template v-if="form.luasRuangan">
                    {{ form.luasRuangan }} m² <span class="ringkasan-row__x">x</span> {{ form.luasRuangan }}m²
                  </template>
                  <template v-else>—</template>
                </span>
              </div>

              <div class="ringkasan-row">
                <span class="ringkasan-row__label">Kerumitan</span>
                <span class="ringkasan-row__value">{{ form.tingkatKerumitan || '—' }}</span>
              </div>

              <div class="ringkasan-row">
                <span class="ringkasan-row__label">Durasi</span>
                <span class="ringkasan-row__value">
                  {{ form.durasi ? form.durasi + ' Hari' : '—' }}
                </span>
              </div>
            </div>

            <div class="ringkasan-divider"></div>

            <p class="ringkasan-card__disclaimer">
              Data yang Anda input hanya digunakan untuk proses estimasi
              biaya proyek interior. Hasil yang ditampilkan merupakan
              <strong>estimasi awal</strong> dan dapat berubah sesuai detail kebutuhan proyek.
              Halaman ini bukan halaman pembayaran atau transaksi.
            </p>

            <button
              class="btn-estimasi"
              @click="handleEstimasi"
              :disabled="isEstimating"
            >
              <span v-if="isEstimating" class="btn-spinner btn-spinner--dark"></span>
              <span v-else>Proses Estimasi</span>
            </button>

            <!-- Hasil Estimasi (muncul setelah proses) -->
            <div class="estimasi-result" v-if="estimasiResult">
              <div class="estimasi-result__divider"></div>
              <p class="estimasi-result__label">Estimasi Biaya</p>
              <p class="estimasi-result__value">{{ estimasiResult }}</p>
              <p class="estimasi-result__note">*Estimasi ini bersifat awal dan dapat berubah.</p>
            </div>

          </div>

        </div>

      </div>
    </section>

  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { RouterLink } from 'vue-router'

const isSaving    = ref(false)
const isEstimating = ref(false)
const estimasiResult = ref('')

const form = reactive({
  namaProyek:        '',
  luasRuangan:       '',
  jenisProyek:       '',
  tingkatKerumitan:  '',
  durasi:            '',
  lokasi:            '',
  konsep:            '',
  material:          '',
  kebutuhanTambahan: '',
})

const handleReset = () => {
  Object.keys(form).forEach(k => form[k] = '')
  estimasiResult.value = ''
}

const handleSimpan = async () => {
  isSaving.value = true
  try {
    const token = localStorage.getItem('token')
    const res = await fetch('http://localhost:8080/estimasi/simpan', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        nama_proyek:        form.namaProyek,
        luas_ruangan:       Number(form.luasRuangan),
        jenis_proyek:       form.jenisProyek,
        tingkat_kerumitan:  form.tingkatKerumitan,
        durasi:             Number(form.durasi),
        lokasi:             form.lokasi,
        konsep:             form.konsep,
        material:           form.material,
        kebutuhan_tambahan: form.kebutuhanTambahan,
      }),
    })
    const data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal menyimpan.')
    alert('Data berhasil disimpan!')
  } catch (err) {
    alert(err.message || 'Tidak dapat terhubung ke server.')
  } finally {
    isSaving.value = false
  }
}

const handleEstimasi = async () => {
  if (!form.luasRuangan || !form.jenisProyek || !form.tingkatKerumitan) {
    alert('Harap isi minimal Luas Ruangan, Jenis Proyek, dan Tingkat Kerumitan.')
    return
  }

  isEstimating.value = true
  estimasiResult.value = ''

  try {
    const token = localStorage.getItem('token')
    const res = await fetch('http://localhost:8080/estimasi/proses', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        luas_ruangan:      Number(form.luasRuangan),
        jenis_proyek:      form.jenisProyek,
        tingkat_kerumitan: form.tingkatKerumitan,
        durasi:            Number(form.durasi),
        lokasi:            form.lokasi,
        konsep:            form.konsep,
        material:          form.material,
      }),
    })
    const data = await res.json()
    if (!res.ok) throw new Error(data.message || 'Gagal memproses estimasi.')
    // Asumsi backend return { estimasi: 'Rp 45.000.000 - Rp 60.000.000' }
    estimasiResult.value = data.estimasi || data.result || 'Estimasi tidak tersedia.'
  } catch (err) {
    alert(err.message || 'Tidak dapat terhubung ke server.')
  } finally {
    isEstimating.value = false
  }
}
</script>

<style scoped>
.estimasi-page {
  padding-top: 72px;
  font-family: 'Montserrat', sans-serif;
  background: #FFFFFF;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 40px;
}

.section {
  padding: 48px 0 80px;
}

/* ── Page Header ── */
.page-header {
  padding: 40px 0 0;
  background: #FFFFFF;
}

.page-header__title {
  font-size: 36px;
  font-weight: 800;
  color: #2C2C2C;
  margin: 0 0 10px;
}

.page-header__breadcrumb {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #888;
  margin-bottom: 16px;
}

.breadcrumb__link {
  color: #888;
  text-decoration: none;
  transition: color 0.2s;
}

.breadcrumb__link:hover { color: #1B7A6E; }
.breadcrumb__sep        { color: #aaa; }
.breadcrumb__current    { color: #555; }

.page-header__desc {
  font-size: 14px;
  line-height: 1.8;
  color: #444;
  margin: 0 0 8px;
}

.page-header__brand {
  color: #1B7A6E;
  font-weight: 700;
}

/* ── Layout Grid ── */
.estimasi-grid {
  display: grid;
  grid-template-columns: 1fr 380px;
  gap: 60px;
  align-items: start;
}

/* ── Form ── */
.estimasi-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 13px;
  font-weight: 600;
  color: #2C2C2C;
}

.form-input {
  width: 100%;
  border: 1.5px solid #DDDDDD;
  border-radius: 6px;
  padding: 13px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-input:focus { border-color: #C8A135; }
.form-input::placeholder { color: #BBBBBB; }

/* Remove number input arrows */
.form-input[type='number']::-webkit-inner-spin-button,
.form-input[type='number']::-webkit-outer-spin-button { -webkit-appearance: none; }
.form-input[type='number'] { -moz-appearance: textfield; }

/* Select */
.form-select-wrap {
  position: relative;
}

.form-select {
  width: 100%;
  border: 1.5px solid #DDDDDD;
  border-radius: 6px;
  padding: 13px 40px 13px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  appearance: none;
  cursor: pointer;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-select:focus { border-color: #C8A135; }

.form-select-icon {
  position: absolute;
  right: 14px;
  top: 50%;
  transform: translateY(-50%);
  color: #888;
  pointer-events: none;
}

/* Textarea */
.form-textarea {
  width: 100%;
  border: 1.5px solid #DDDDDD;
  border-radius: 6px;
  padding: 13px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  resize: vertical;
  min-height: 100px;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-textarea:focus { border-color: #C8A135; }
.form-textarea::placeholder { color: #BBBBBB; }

/* ── Form Actions ── */
.form-actions {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin-top: 4px;
}

.btn-reset {
  background: none;
  border: none;
  font-size: 13.5px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  color: #888;
  cursor: pointer;
  padding: 4px 0;
  transition: color 0.2s;
  align-self: flex-end;
}

.btn-reset:hover { color: #444; }

.btn-simpan {
  flex: 1;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  padding: 16px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn-simpan:hover:not(:disabled) {
  background: #b08e2c;
  transform: translateY(-1px);
}

.btn-simpan:disabled { opacity: 0.7; cursor: not-allowed; }

/* ── Sidebar Ringkasan ── */
.estimasi-sidebar {
  position: sticky;
  top: 96px;
}

.ringkasan-card {
  border: 1.5px solid #E5E5E5;
  border-radius: 10px;
  padding: 28px 28px 32px;
  background: #FFFFFF;
}

.ringkasan-card__title {
  font-size: 18px;
  font-weight: 800;
  color: #2C2C2C;
  margin: 0 0 20px;
}

.ringkasan-card__rows {
  display: flex;
  flex-direction: column;
  gap: 14px;
  margin-bottom: 20px;
}

.ringkasan-row {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 12px;
}

.ringkasan-row__label {
  font-size: 12.5px;
  color: #888;
  flex-shrink: 0;
}

.ringkasan-row__value {
  font-size: 12.5px;
  font-weight: 700;
  color: #2C2C2C;
  text-align: right;
}

.ringkasan-row__x {
  font-size: 11px;
  color: #aaa;
  margin: 0 2px;
}

.ringkasan-divider {
  height: 1px;
  background: #EEEEEE;
  margin: 20px 0;
}

.ringkasan-card__disclaimer {
  font-size: 12px;
  line-height: 1.8;
  color: #666;
  text-align: justify;
  margin: 0 0 24px;
}

/* ── Proses Estimasi Button ── */
.btn-estimasi {
  width: 100%;
  background: #FFFFFF;
  color: #2C2C2C;
  font-size: 14px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  padding: 15px;
  border: 1.5px solid #CCCCCC;
  border-radius: 6px;
  cursor: pointer;
  transition: border-color 0.2s, background 0.2s, transform 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn-estimasi:hover:not(:disabled) {
  border-color: #1B7A6E;
  color: #1B7A6E;
  transform: translateY(-1px);
}

.btn-estimasi:disabled { opacity: 0.6; cursor: not-allowed; }

/* ── Hasil Estimasi ── */
.estimasi-result {
  margin-top: 4px;
}

.estimasi-result__divider {
  height: 1px;
  background: #EEEEEE;
  margin: 20px 0;
}

.estimasi-result__label {
  font-size: 12px;
  color: #888;
  margin: 0 0 6px;
}

.estimasi-result__value {
  font-size: 20px;
  font-weight: 800;
  color: #1B7A6E;
  margin: 0 0 6px;
}

.estimasi-result__note {
  font-size: 11px;
  color: #aaa;
  margin: 0;
}

/* ── Spinner ── */
.btn-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
  display: inline-block;
}

.btn-spinner--dark {
  border-color: rgba(0,0,0,0.15);
  border-top-color: #2C2C2C;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ── Responsive ── */
@media (max-width: 960px) {
  .estimasi-grid {
    grid-template-columns: 1fr;
    gap: 36px;
  }

  .estimasi-sidebar {
    position: static;
    order: -1;
  }
}

@media (max-width: 560px) {
  .container { padding: 0 20px; }
  .form-row  { grid-template-columns: 1fr; }
  .form-actions { flex-direction: column-reverse; align-items: stretch; }
  .btn-reset { text-align: center; }
}
</style>