<template>
  <div class="portfolio-page">

    <!-- ── LIST VIEW ── -->
    <template v-if="!showForm">

      <!-- Toolbar -->
      <div class="toolbar">
        <button class="btn-add" @click="openForm(null)">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
            <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          Input Variabel Proyek
        </button>

        <div class="toolbar__right">
          <div class="sort-wrap">
            <select v-model="sortBy" class="sort-select">
              <option value="date">Sort by: Date Created</option>
              <option value="title">Sort by: Title</option>
            </select>
            <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="sort-chevron">
              <polyline points="6 9 12 15 18 9"/>
            </svg>
          </div>

          <button class="btn-filter">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polygon points="22 3 2 3 10 12.46 10 19 14 21 14 12.46 22 3"/>
            </svg>
            Filter
          </button>
        </div>
      </div>

      <!-- Empty state -->
      <div class="empty-state" v-if="filteredItems.length === 0">
        <svg width="52" height="52" viewBox="0 0 24 24" fill="none" stroke="#D0D0D0" stroke-width="1.2">
          <rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/>
          <rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/>
        </svg>
        <p>Belum ada portofolio.<br/>Klik "+ Input Variabel Proyek" untuk menambah.</p>
      </div>

      <!-- Grid -->
      <div class="portfolio-grid" v-else>
        <div
          class="portfolio-card"
          v-for="item in filteredItems"
          :key="item.id"
        >
          <!-- Gambar + overlay hover -->
          <div class="card-img-wrap">
            <img :src="getImageUrl(item.gambar)" :alt="item.judul_proyek" class="card-img" />

            <div class="card-overlay">
              <button class="overlay-cart">Add to cart</button>
              <div class="overlay-actions">
                <button class="overlay-action">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/>
                    <line x1="8.59" y1="13.51" x2="15.42" y2="17.49"/>
                    <line x1="15.41" y1="6.51" x2="8.59" y2="10.49"/>
                  </svg>
                  Share
                </button>
                <button class="overlay-action">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <polyline points="16 3 21 3 21 8"/><line x1="4" y1="20" x2="21" y2="3"/>
                    <polyline points="21 16 21 21 16 21"/><line x1="15" y1="15" x2="21" y2="21"/>
                  </svg>
                  Compare
                </button>
                <button
                  class="overlay-action"
                  :class="{ 'overlay-action--liked': item.liked }"
                  @click="item.liked = !item.liked"
                >
                  <svg width="13" height="13" viewBox="0 0 24 24" :fill="item.liked ? 'currentColor' : 'none'" stroke="currentColor" stroke-width="2">
                    <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                  </svg>
                  Like
                </button>
              </div>
            </div>
          </div>

          <!-- Info + aksi -->
          <div class="card-body">
            <div class="card-info">
              <h3 class="card-title">{{ item.judul_proyek}}</h3>
              <p class="card-desc">{{ item.deskripsi }}</p>
            </div>
            <div class="card-btns">
              <button class="card-btn card-btn--edit" @click="openForm(item)" title="Edit">
                <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
              </button>
              <button class="card-btn card-btn--delete" @click="deleteItem(item.id)" title="Hapus">
                <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <polyline points="3 6 5 6 21 6"/>
                  <path d="M19 6l-1 14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2L5 6"/>
                  <path d="M10 11v6"/><path d="M14 11v6"/>
                  <path d="M9 6V4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2"/>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

    </template>

    <!-- ── FORM VIEW ── -->
    <template v-else>

      <div class="form-header">
        <button class="form-back" @click="closeForm">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <h2 class="form-header__title">Tambah Data Portofolio</h2>
      </div>

      <div class="form-card">

        <div class="form-field">
          <label class="form-label">Judul</label>
          <input
            v-model="formData.judul_proyek"
            type="text"
            class="form-input"
            :class="{ 'form-input--error': formErrors.title }"
            placeholder="Masukkan judul yang sesuai"
          />
          <span class="form-error" v-if="formErrors.title">{{ formErrors.title }}</span>
        </div>

        <div class="form-field">
          <label class="form-label">Deskripsi</label>
          <input
            v-model="formData.deskripsi"
            type="text"
            class="form-input"
            :class="{ 'form-input--error': formErrors.description }"
            placeholder="Masukkan deskripsi"
          />
          <span class="form-error" v-if="formErrors.description">{{ formErrors.description }}</span>
        </div>

        <div class="form-field">
          <label class="form-label">Gambar Sampul</label>

          <!-- Preview -->
          <div v-if="imagePreview" class="preview-wrap">
            <img :src="imagePreview" alt="Preview" class="preview-img" />
            <button class="preview-remove" @click="removeImage">
              <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>

          <!-- Dropzone -->
          <div
            v-else
            class="dropzone"
            :class="{ 'dropzone--over': isDragging, 'dropzone--error': formErrors.gambar }"
            @dragenter.prevent="isDragging = true"
            @dragleave.prevent="isDragging = false"
            @dragover.prevent
            @drop.prevent="onDrop"
          >
            <svg width="44" height="44" viewBox="0 0 24 24" fill="none" stroke="#C8C8C8" stroke-width="1.3">
              <rect x="3" y="3" width="18" height="18" rx="2"/>
              <circle cx="8.5" cy="8.5" r="1.5"/>
              <polyline points="21 15 16 10 5 21"/>
            </svg>
            <p class="dropzone__title">Drag and Drop Your File Here!</p>
            <p class="dropzone__hint">
              Please upload <strong>JPG, JPEG, PNG or SVG</strong> files<br/>
              A file maximum size should be <strong>5 MB</strong>
            </p>
            <label class="btn-upload-file">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
                <polyline points="16 16 12 12 8 16"/>
                <line x1="12" y1="12" x2="12" y2="21"/>
                <path d="M20.39 18.39A5 5 0 0 0 18 9h-1.26A8 8 0 1 0 3 16.3"/>
              </svg>
              Upload File
              <input
                type="file"
                accept=".jpg,.jpeg,.png,.svg"
                style="display:none"
                ref="fileInput"
                @change="onFileChange"
              />
            </label>
          </div>
          <span class="form-error" v-if="formErrors.gambar">{{ formErrors.gambar }}</span>
        </div>

        <div class="form-actions">
          <button class="btn-cancel" @click="closeForm">Cancel</button>
          <button class="btn-submit" @click="submitForm" :disabled="isSubmitting">
            <span v-if="isSubmitting" class="btn-spinner"></span>
            <span v-else>Upload</span>
          </button>
        </div>

      </div>
    </template>

  </div>
</template>


<script setup>
import { ref, reactive, computed, onMounted } from 'vue'

const fileInput    = ref(null)
const showForm     = ref(false)
const sortBy       = ref('date')
const isDragging   = ref(false)
const isSubmitting = ref(false)
const imagePreview = ref('')
const editingId    = ref(null)

const formData   = reactive({ judul_proyek: '', deskripsi: '', gambar: null })
const formErrors = reactive({ judul_proyek: '', deskripsi: '', gambar: '' })

// DATA DARI BACKEND
const portfolioItems = ref([])

// ================= GET DATA =================
const getPortfolios = async () => {
  try {
    const res = await fetch("http://localhost:8081/api/portfolios")
    const data = await res.json()
    portfolioItems.value = data.data
  } catch (err) {
    console.error("Gagal ambil data:", err)
  }
}

// ================= SORT =================
const filteredItems = computed(() => {
  let list = [...portfolioItems.value]

  if (sortBy.value === 'date') {
    list.sort((a, b) => b.id - a.id)
  } else {
    list.sort((a, b) =>
      a.judul_proyek.localeCompare(b.judul_proyek)
    )
  }

  return list
})

// ================= FORM =================
const openForm = (item) => {
  if (item) {
    editingId.value      = item.id
    formData.judul_proyek       = item.judul_proyek
    formData.deskripsi = item.deskripsi

    // preview dari backend
    imagePreview.value   = `http://localhost:8081/uploads/${item.gambar}`
  } else {
    editingId.value      = null
    formData.judul_proyek       = ''
    formData.deskripsi = ''
    formData.gambar       = null
    imagePreview.value   = ''
  }

  formErrors.judul_proyek = ''
  formErrors.deskripsi = ''
  formErrors.gambar = ''

  showForm.value = true
}

const closeForm = () => {
  showForm.value = false
  imagePreview.value = ''
  formData.gambar = null
}

// ================= VALIDASI =================
const validateForm = () => {
  let valid = true

  formErrors.judul_proyek = ''
  formErrors.deskripsi = ''
  formErrors.gambar = ''

  if (!formData.judul_proyek.trim()) {
    formErrors.judul_proyek = 'Judul wajib diisi.'
    valid = false
  }

  if (!formData.deskripsi.trim()) {
    formErrors.deskripsi = 'Deskripsi wajib diisi.'
    valid = false
  }

  if (!imagePreview.value) {
    formErrors.gambar = 'Gambar wajib diisi.'
    valid = false
  }

  return valid
}

// ================= SUBMIT (POST & PUT) =================
const submitForm = async () => {
  if (!validateForm()) return

console.log("SUBMIT DIKLIK")
isSubmitting.value = true

try {
   const form = new FormData()
    form.append("judul_proyek", formData.judul_proyek)
    form.append("deskripsi", formData.deskripsi)

     // kalau ada file baru
    if (formData.gambar) {
      form.append("gambar", formData.gambar)
    }
      const url = editingId.value
          ? `http://localhost:8081/api/admin/portfolios/${editingId.value}`
          : `http://localhost:8081/api/admin/portfolios`

    const method = editingId.value ? "PUT" : "POST"

    const token = localStorage.getItem("token")
    const res = await fetch(url, { 
      method: method,
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: form
    })
    const result = await res.json()
    console.log("RESPONSE:", result)

    await getPortfolios()
    closeForm()
  

  } catch (err) {
    console.error("Gagal submit:", err)
  }

  isSubmitting.value = false
}

// ================= DELETE =================
const deleteItem = async (id) => {
  if (!confirm("Hapus portofolio ini?")) return

  try {
    await fetch(`http://localhost:8081/api/admin/portfolios/${id}`, {
      method: "DELETE",
      headers: {
        "Authorization": "Bearer admin-token"
      }
    })

    getPortfolios()
  } catch (err) {
    console.error("Gagal hapus:", err)
  }
}

// ================= UPLOAD PREVIEW =================
const processFile = (file) => {
  if (!file) return

  const allowed = ['image/jpeg', 'image/jpg', 'image/png']
  if (!allowed.includes(file.type)) {
    formErrors.gambar = 'Format tidak didukung.'
    return
  }

  if (file.size > 5 * 1024 * 1024) {
    formErrors.gambar = 'Maks 5MB.'
    return
  }

  formErrors.gambar = ''
  formData.gambar = file

  const reader = new FileReader()
  reader.onload = (e) => {
    imagePreview.value = e.target.result
  }
  reader.readAsDataURL(file)
}

const onFileChange = (e) => processFile(e.target.files[0])
const onDrop = (e) => {
  isDragging.value = false
  processFile(e.dataTransfer.files[0])
}

const removeImage = () => {
  imagePreview.value = ''
  formData.gambar = null
  if (fileInput.value) fileInput.value.value = ''
}

const getImageUrl = (gambar) => {
  if (!gambar) return ''
  return `http://localhost:8081/uploads/${gambar}`
}

// ================= LOAD AWAL =================
onMounted(() => {
  getPortfolios()
})
</script>

<style scoped>
.portfolio-page {
  font-family: 'Montserrat', sans-serif;
  padding: 24px 32px;
}

/* ── Toolbar ── */
.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
}

.btn-add {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: #FFFFFF;
  border: 1.5px solid #1B7A6E;
  color: #1B7A6E;
  font-size: 13.5px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}
.btn-add:hover { background: #1B7A6E; color: #FFFFFF; }

.toolbar__right { display: flex; align-items: center; gap: 10px; }

.sort-wrap {
  position: relative;
  display: flex;
  align-items: center;
}

.sort-select {
  appearance: none;
  background: #FFFFFF;
  border: 1.5px solid #E5E5E5;
  border-radius: 8px;
  padding: 9px 36px 9px 14px;
  font-size: 13px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  cursor: pointer;
  outline: none;
  transition: border-color 0.2s;
}
.sort-select:focus { border-color: #1B7A6E; }

.sort-chevron {
  position: absolute;
  right: 11px;
  pointer-events: none;
  color: #777;
}

.btn-filter {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: #FFFFFF;
  border: 1.5px solid #E5E5E5;
  border-radius: 8px;
  padding: 9px 16px;
  font-size: 13px;
  font-family: 'Montserrat', sans-serif;
  color: #555;
  cursor: pointer;
  transition: border-color 0.2s, color 0.2s;
}
.btn-filter:hover { border-color: #1B7A6E; color: #1B7A6E; }

/* ── Empty ── */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 14px;
  padding: 80px 20px;
  color: #BBBBBB;
  font-size: 14px;
  text-align: center;
  line-height: 1.6;
}

/* ── Grid ── */
.portfolio-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

/* ── Card ── */
.portfolio-card {
  background: #FFFFFF;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0,0,0,0.06);
  transition: box-shadow 0.25s, transform 0.25s;
}
.portfolio-card:hover {
  box-shadow: 0 8px 28px rgba(0,0,0,0.12);
  transform: translateY(-2px);
}

.card-img-wrap {
  position: relative;
  aspect-ratio: 1 / 1;
  overflow: hidden;
}

.card-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.4s ease;
}
.portfolio-card:hover .card-img { transform: scale(1.05); }

/* Overlay — sesuai PDF */
.card-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0,0,0,0.46);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 12px;
  opacity: 0;
  transition: opacity 0.25s;
}
.portfolio-card:hover .card-overlay { opacity: 1; }

.overlay-cart {
  background: #FFFFFF;
  border: none;
  color: #C8A135;
  font-size: 13px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  padding: 9px 28px;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}
.overlay-cart:hover { background: #C8A135; color: #FFFFFF; }

.overlay-actions {
  display: flex;
  gap: 14px;
}

.overlay-action {
  display: flex;
  align-items: center;
  gap: 4px;
  background: none;
  border: none;
  color: #FFFFFF;
  font-size: 12px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  cursor: pointer;
  padding: 0;
  transition: color 0.2s;
}
.overlay-action:hover       { color: #C8A135; }
.overlay-action--liked      { color: #EF4444; }

/* Card body */
.card-body {
  padding: 12px 14px;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 8px;
}

.card-title {
  font-size: 14px;
  font-weight: 700;
  color: #1A1A1A;
  margin: 0 0 3px;
}
.card-desc {
  font-size: 12px;
  color: #888;
  margin: 0;
}

.card-btns {
  display: flex;
  gap: 5px;
  flex-shrink: 0;
}

.card-btn {
  width: 28px; height: 28px;
  border-radius: 6px;
  border: 1.5px solid #E5E5E5;
  background: #FFFFFF;
  cursor: pointer;
  display: flex; align-items: center; justify-content: center;
  color: #666;
  transition: background 0.2s, border-color 0.2s, color 0.2s;
}
.card-btn--edit:hover   { background: #EBF5F3; border-color: #1B7A6E; color: #1B7A6E; }
.card-btn--delete:hover { background: #FEF2F2; border-color: #EF4444; color: #EF4444; }

/* ── Form header ── */
.form-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 18px;
}

.form-back {
  width: 32px; height: 32px;
  border-radius: 8px;
  border: 1.5px solid #E5E5E5;
  background: #FFFFFF;
  cursor: pointer;
  display: flex; align-items: center; justify-content: center;
  color: #555;
  transition: background 0.2s, color 0.2s, border-color 0.2s;
}
.form-back:hover { background: #F0F9F7; color: #1B7A6E; border-color: #1B7A6E; }

.form-header__title {
  font-size: 16px;
  font-weight: 700;
  color: #1B7A6E;
  margin: 0;
}

/* ── Form card ── */
.form-card {
  background: #FFFFFF;
  border-radius: 12px;
  padding: 32px 36px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  max-width: 760px;
}

.form-field {
  display: flex;
  flex-direction: column;
  gap: 7px;
  margin-bottom: 20px;
}

.form-label {
  font-size: 13.5px;
  font-weight: 600;
  color: #1B7A6E;
}

.form-input {
  width: 100%;
  border: 1.5px solid #E8E8E8;
  border-radius: 8px;
  padding: 13px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}
.form-input:focus        { border-color: #1B7A6E; box-shadow: 0 0 0 3px rgba(27,122,110,0.07); }
.form-input--error       { border-color: #EF4444 !important; }
.form-input::placeholder { color: #C0C0C0; }

.form-error { font-size: 11.5px; color: #EF4444; }

/* ── Dropzone ── */
.dropzone {
  border: 2px dashed #DDDDDD;
  border-radius: 10px;
  background: #FAFAFA;
  padding: 44px 20px 36px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  text-align: center;
  cursor: default;
  transition: border-color 0.2s, background 0.2s;
}
.dropzone--over  { border-color: #1B7A6E; background: #F0FDF9; }
.dropzone--error { border-color: #EF4444; }

.dropzone__title {
  font-size: 15px;
  font-weight: 700;
  color: #2C2C2C;
  margin: 0;
}
.dropzone__hint {
  font-size: 12.5px;
  color: #999;
  line-height: 1.7;
  margin: 0;
}

.btn-upload-file {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  padding: 10px 24px;
  border-radius: 20px;
  cursor: pointer;
  margin-top: 4px;
  transition: background 0.2s;
}
.btn-upload-file:hover { background: #b08e2c; }

/* ── Preview ── */
.preview-wrap {
  position: relative;
  border-radius: 10px;
  overflow: hidden;
  border: 1.5px solid #E5E5E5;
}
.preview-img {
  width: 100%;
  max-height: 260px;
  object-fit: cover;
  display: block;
}
.preview-remove {
  position: absolute;
  top: 10px; right: 10px;
  width: 28px; height: 28px;
  border-radius: 50%;
  background: rgba(0,0,0,0.5);
  border: none;
  cursor: pointer;
  color: #fff;
  display: flex; align-items: center; justify-content: center;
  transition: background 0.2s;
}
.preview-remove:hover { background: rgba(0,0,0,0.75); }

/* ── Form actions ── */
.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 8px;
}

.btn-cancel {
  padding: 11px 28px;
  background: #FFFFFF;
  border: 1.5px solid #DDDDDD;
  border-radius: 24px;
  font-size: 13.5px;
  font-weight: 600;
  font-family: 'Montserrat', sans-serif;
  color: #555;
  cursor: pointer;
  transition: border-color 0.2s, color 0.2s;
}
.btn-cancel:hover { border-color: #999; color: #333; }

.btn-submit {
  padding: 11px 36px;
  background: #C8A135;
  border: none;
  border-radius: 24px;
  font-size: 13.5px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  color: #FFFFFF;
  cursor: pointer;
  transition: background 0.2s;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  min-width: 100px;
  justify-content: center;
}
.btn-submit:hover:not(:disabled) { background: #b08e2c; }
.btn-submit:disabled              { opacity: 0.7; cursor: not-allowed; }

.btn-spinner {
  width: 15px; height: 15px;
  border: 2px solid rgba(255,255,255,0.35);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

/* ── Responsive ── */
@media (max-width: 1100px) { .portfolio-grid { grid-template-columns: repeat(3, 1fr); } }
@media (max-width: 800px)  { .portfolio-grid { grid-template-columns: repeat(2, 1fr); } }
@media (max-width: 500px)  {
  .portfolio-grid { grid-template-columns: 1fr; }
  .form-card      { padding: 24px 20px; }
}
</style>