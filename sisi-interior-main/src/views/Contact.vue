<template>
  <div class="contact-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">Contact</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">Contact</span>
        </div>
      </div>
    </div>

    <!-- ── Main Content ── -->
    <section class="contact-section section">
      <div class="container contact-grid">

        <!-- ── Kiri: Info Card ── -->
        <div class="info-card">

          <div class="info-card__item">
            <div class="info-card__icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="white" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7zm0 9.5c-1.38 0-2.5-1.12-2.5-2.5S10.62 6.5 12 6.5s2.5 1.12 2.5 2.5S13.38 11.5 12 11.5z"/>
              </svg>
            </div>
            <div class="info-card__body">
              <p class="info-card__label">Alamat</p>
              <p class="info-card__value">Kota Bogor, Jawa Barat, Indonesia. 16163</p>
            </div>
          </div>

          <div class="info-card__item">
            <div class="info-card__icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="white" xmlns="http://www.w3.org/2000/svg">
                <path d="M6.62 10.79a15.05 15.05 0 006.59 6.59l2.2-2.2a1 1 0 011.01-.24c1.12.37 2.33.57 3.58.57a1 1 0 011 1V20a1 1 0 01-1 1C10.61 21 3 13.39 3 4a1 1 0 011-1h3.5a1 1 0 011 1c0 1.25.2 2.45.57 3.58a1 1 0 01-.25 1.01l-2.2 2.2z"/>
              </svg>
            </div>
            <div class="info-card__body">
              <p class="info-card__label">Telepon</p>
              <p class="info-card__value">Mobile:<br />0812-9052-9092</p>
            </div>
          </div>

          <div class="info-card__item">
            <div class="info-card__icon">
              <svg width="22" height="22" viewBox="0 0 24 24" fill="white" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 2a10 10 0 100 20A10 10 0 0012 2zm.5 5v5.25l4.5 2.67-.75 1.23L11 13V7h1.5z"/>
              </svg>
            </div>
            <div class="info-card__body">
              <p class="info-card__label">Jam Kerja</p>
              <p class="info-card__value">
                Monday-Friday:<br />9:00 - 22:00<br /><br />
                Saturday-Sunday:<br />9:00 - 21:00
              </p>
            </div>
          </div>

        </div>

        <!-- ── Kanan: Form ── -->
        <div class="contact-form">

          <div class="form-group">
            <label class="form-label">Nama</label>
            <input
              v-model="form.nama"
              type="text"
              class="form-input"
              placeholder="Abc"
            />
          </div>

          <div class="form-group">
            <label class="form-label">Alamat email</label>
            <input
              v-model="form.email"
              type="email"
              class="form-input"
              placeholder="Abc@def.com"
            />
          </div>

          <div class="form-group">
            <label class="form-label">Subjek</label>
            <input
              v-model="form.subjek"
              type="text"
              class="form-input"
              placeholder="Opsional"
            />
          </div>

          <div class="form-group">
            <label class="form-label">Pesan</label>
            <textarea
              v-model="form.pesan"
              class="form-textarea"
              placeholder="Hi! saya ingin menanyakan sesuatu"
              rows="5"
            ></textarea>
          </div>

          <button class="btn-submit" @click="handleSubmit">Submit</button>

        </div>

      </div>
    </section>

  </div>
</template>

<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'

const form = ref({
  nama: '',
  email: '',
  subjek: '',
  pesan: '',
})

const handleSubmit = () => {
  console.log('Form submitted:', form.value)
  // TODO: Tambahkan logika pengiriman form di sini
}
</script>

<style scoped>
.contact-page {
  padding-top: 72px;
  font-family: 'Montserrat', sans-serif;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 40px;
}

.section {
  padding: 30px 0 80px;
  background: #FFFFFF;
}

/* ── Page Header ── */
.page-header {
  padding: 40px 0 20px;
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
}

.breadcrumb__link {
  color: #888;
  text-decoration: none;
  transition: color 0.2s;
}

.breadcrumb__link:hover { color: #1B7A6E; }
.breadcrumb__sep { color: #aaa; }
.breadcrumb__current { color: #555; }

/* ── Contact Grid ── */
.contact-grid {
  display: grid;
  grid-template-columns: 380px 1fr;
  gap: 60px;
  align-items: start;
}

/* ── Info Card ── */
.info-card {
  background: #EAB834;
  border-radius: 8px;
  padding: 48px 36px;
  display: flex;
  flex-direction: column;
  gap: 40px;
}

.info-card__item {
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.info-card__icon {
  width: 40px;
  height: 40px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info-card__label {
  font-size: 20px;
  font-weight: 700;
  color: #FFFFFF;
  margin: 0 0 6px;
}

.info-card__value {
  font-size: 13.5px;
  line-height: 1.7;
  color: #FFFFFF;
  margin: 0;
  opacity: 0.92;
}

/* ── Contact Form ── */
.contact-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
  padding-top: 8px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 14px;
  font-weight: 600;
  color: #2C2C2C;
}

.form-input,
.form-textarea {
  width: 100%;
  border: 1.5px solid #DDDDDD;
  border-radius: 6px;
  padding: 14px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  transition: border-color 0.2s;
  box-sizing: border-box;
}

.form-input::placeholder,
.form-textarea::placeholder {
  color: #AAAAAA;
}

.form-input:focus,
.form-textarea:focus {
  border-color: #EAB834;
}

.form-textarea {
  resize: vertical;
  min-height: 120px;
}

/* ── Submit Button ── */
.btn-submit {
  width: auto; 
  min-width: 160px; 
  background: #EAB834;
  color: #FFFFFF;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  letter-spacing: 0.5px;
  padding: 14px 28px; 
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
  align-self: flex-start;
}

.btn-submit:hover {
  background: #EAB834;
  transform: translateY(-1px);
}

/* ── Responsive ── */
@media (max-width: 860px) {
  .contact-grid {
    grid-template-columns: 1fr;
    gap: 36px;
  }

  .btn-submit {
    max-width: 100%;
  }
}

@media (max-width: 560px) {
  .container {
    padding: 0 20px;
  }

  .info-card {
    padding: 36px 24px;
  }
}
</style>