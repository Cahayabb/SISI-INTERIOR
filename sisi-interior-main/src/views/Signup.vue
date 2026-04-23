<template>
  <div class="page-wrapper">

    <!-- Navbar -->
    <nav class="navbar">
      <div class="navbar__brand">
        <div class="navbar__logo-icon">
          <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
            <rect x="2" y="2" width="28" height="28" rx="4" stroke="#1B7A6E" stroke-width="2"/>
            <path d="M8 24V14l8-6 8 6v10" stroke="#1B7A6E" stroke-width="1.8" stroke-linejoin="round"/>
            <rect x="12" y="17" width="8" height="7" rx="1" stroke="#C8A135" stroke-width="1.5"/>
          </svg>
        </div>
        <div class="navbar__brand-text">
          <span class="navbar__brand-main">SISI</span>
          <span class="navbar__brand-sub">INTERIOR</span>
        </div>
      </div>
      <ul class="navbar__links">
        <li><a href="#">Home</a></li>
        <li><a href="#">About</a></li>
        <li><a href="#">Service</a></li>
        <li><a href="#">Portofolio</a></li>
        <li><a href="#">Contact</a></li>
      </ul>
      <RouterLink to="/register" class="navbar__cta">Sign Up</RouterLink>
    </nav>

    <!-- Main Content -->
    <main class="main-content">

      <!-- Left: Form -->
      <div class="form-section">
        <h1 class="form-title">DAFTAR AKUN</h1>

        <!-- Alerts -->
        <div class="alert alert--success" v-if="successMsg">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="20 6 9 17 4 12"/></svg>
          {{ successMsg }}
        </div>
        <div class="alert alert--error" v-if="errorMsg">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
          {{ errorMsg }}
        </div>

        <div class="form">

          <!-- Nama -->
          <div class="form-group">
            <label class="form-label">Nama</label>
            <input
              v-model="form.name"
              type="text"
              class="form-input"
              :class="{ 'form-input--error': errors.name }"
              placeholder="Abc"
            />
            <span class="form-error" v-if="errors.name">{{ errors.name }}</span>
          </div>

          <!-- Username -->
          <div class="form-group">
            <label class="form-label">Username</label>
            <input
              v-model="form.username"
              type="text"
              class="form-input"
              :class="{ 'form-input--error': errors.username }"
              placeholder="john_doe"
            />
            <span class="form-error" v-if="errors.username">{{ errors.username }}</span>
          </div>

          <!-- Alamat Email -->
          <div class="form-group">
            <label class="form-label">Alamat email</label>
            <input
              v-model="form.email"
              type="email"
              class="form-input"
              :class="{ 'form-input--error': errors.email }"
              placeholder="Abc@def.com"
            />
            <span class="form-error" v-if="errors.email">{{ errors.email }}</span>
          </div>

          <!-- No. Hp -->
          <div class="form-group">
            <label class="form-label">No. Hp</label>
            <input
              v-model="form.phone"
              type="tel"
              class="form-input"
              :class="{ 'form-input--error': errors.phone }"
              placeholder="Opsional"
            />
            <span class="form-error" v-if="errors.phone">{{ errors.phone }}</span>
          </div>

          <!-- Password -->
          <div class="form-group">
            <label class="form-label">Password</label>
            <div class="input-wrap">
              <input
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                class="form-input form-input--has-icon"
                :class="{ 'form-input--error': errors.password }"
                placeholder="Masukkan password"
              />
              <button class="eye-btn" type="button" @click="showPassword = !showPassword" tabindex="-1">
                <svg v-if="!showPassword" width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                <svg v-else width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M17.94 17.94A10.07 10.07 0 0112 20c-7 0-11-8-11-8a18.45 18.45 0 015.06-5.94"/><path d="M9.9 4.24A9.12 9.12 0 0112 4c7 0 11 8 11 8a18.5 18.5 0 01-2.16 3.19"/><line x1="1" y1="1" x2="23" y2="23"/></svg>
              </button>
            </div>
            <span class="form-error" v-if="errors.password">{{ errors.password }}</span>
          </div>

          <!-- Submit -->
          <button
            class="btn-submit"
            @click="handleSignUp"
            :disabled="isLoading"
          >
            <span v-if="isLoading" class="spinner"></span>
            <span v-else>Submit</span>
          </button>

        </div>

        <p class="footer-text">
          Sudah punya akun?
          <RouterLink to="/login" class="footer-link">Masuk di sini</RouterLink>
        </p>
      </div>

      <!-- Right: Photo Grid -->
      <div class="photo-grid">
        <!-- foto 1: kiri atas, mulai agak ke bawah -->
        <div class="photo-grid__item photo-grid__item--1">
          <img src="https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=600&q=80" alt="Interior bedroom" />
        </div>
        <!-- foto 2: kanan atas, mulai dari paling atas, lebih tinggi -->
        <div class="photo-grid__item photo-grid__item--2">
          <img src="https://images.unsplash.com/photo-1555041469-a586c61ea9bc?w=600&q=80" alt="Interior dining" />
        </div>
        <!-- foto 3: kiri bawah, lebih lebar -->
        <div class="photo-grid__item photo-grid__item--3">
          <img src="https://images.unsplash.com/photo-1616486338812-3dadae4b4ace?w=600&q=80" alt="Interior minimal" />
        </div>
        <!-- foto 4: kanan bawah, lebih sempit, melampaui batas bawah -->
        <div class="photo-grid__item photo-grid__item--4">
          <img src="https://images.unsplash.com/photo-1505693416388-ac5ce068fe85?w=600&q=80" alt="Interior bedroom 2" />
        </div>
      </div>

    </main>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter, RouterLink } from 'vue-router'

const router      = useRouter()
const isLoading   = ref(false)
const showPassword = ref(false)
const errorMsg    = ref('')
const successMsg  = ref('')

const form = reactive({
  name: '',
  username: '',
  email: '',
  phone: '',
  password: '',
})

const errors = reactive({
  name: '',
  username: '',
  email: '',
  phone: '',
  password: '',
})

const validate = () => {
  let valid = true
  Object.keys(errors).forEach(k => errors[k] = '')

  if (!form.name.trim()) {
    errors.name = 'Nama wajib diisi.'; valid = false
  }

  if (!form.username.trim()) {
    errors.username = 'Username wajib diisi.'; valid = false
  } else if (!/^[a-zA-Z0-9_]{3,20}$/.test(form.username)) {
    errors.username = 'Username 3-20 karakter, hanya huruf, angka, dan _.'; valid = false
  }

  if (!form.email) {
    errors.email = 'Email wajib diisi.'; valid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email)) {
    errors.email = 'Format email tidak valid.'; valid = false
  }

  if (form.phone && !/^[0-9+\-\s]{8,15}$/.test(form.phone)) {
    errors.phone = 'Format nomor telepon tidak valid.'; valid = false
  }

  if (!form.password) {
    errors.password = 'Password wajib diisi.'; valid = false
  } else if (form.password.length < 8) {
    errors.password = 'Password minimal 8 karakter.'; valid = false
  }

  return valid
}

const handleSignUp = async () => {
  errorMsg.value  = ''
  successMsg.value = ''
  if (!validate()) return

  isLoading.value = true
  try {
    const res = await fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name:     form.name,
        username: form.username,
        email:    form.email,
        phone:    form.phone,
        password: form.password,
        role:     'customer',
      }),
    })

    const data = await res.json()

    if (!res.ok) {
      errorMsg.value = data.message || 'Pendaftaran gagal. Coba lagi.'
      return
    }

    successMsg.value = 'Akun berhasil dibuat! Mengarahkan ke halaman login...'
    setTimeout(() => router.push('/login'), 2000)

  } catch (err) {
    errorMsg.value = 'Tidak dapat terhubung ke server. Coba lagi nanti.'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
/* ── Base ── */
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

.page-wrapper {
  min-height: 100vh;
  background: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  display: flex;
  flex-direction: column;
}

/* ── Navbar ── */
.navbar {
  display: flex;
  align-items: center;
  gap: 40px;
  padding: 0 60px;
  height: 70px;
  border-bottom: 1px solid #EEEEEE;
  background: #FFFFFF;
}

.navbar__brand {
  display: flex;
  align-items: center;
  gap: 10px;
  text-decoration: none;
  flex-shrink: 0;
}

.navbar__brand-text {
  display: flex;
  flex-direction: column;
  line-height: 1;
}

.navbar__brand-main {
  font-family: 'Playfair Display', serif;
  font-size: 18px;
  font-weight: 700;
  color: #1B7A6E;
  letter-spacing: 5px;
}

.navbar__brand-sub {
  font-size: 7px;
  font-weight: 600;
  color: #C8A135;
  letter-spacing: 3px;
  text-transform: uppercase;
}

.navbar__links {
  display: flex;
  list-style: none;
  gap: 32px;
  flex: 1;
  justify-content: center;
}

.navbar__links a {
  font-size: 13.5px;
  font-weight: 500;
  color: #444444;
  text-decoration: none;
  transition: color 0.2s;
}

.navbar__links a:hover {
  color: #1B7A6E;
}

.navbar__cta {
  flex-shrink: 0;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  padding: 10px 26px;
  border-radius: 6px;
  text-decoration: none;
  transition: background 0.2s;
}

.navbar__cta:hover {
  background: #b08e2c;
}

/* ── Main ── */
.main-content {
  display: flex;
  flex: 1;
  padding: 60px 0 60px 60px;
  gap: 40px;
  overflow: hidden;
  align-items: flex-start;
}

/* ── Form Section ── */
.form-section {
  flex: 0 0 600px;
  max-width: 600px;
}

.form-title {
  font-family: 'Montserrat', sans-serif;
  font-size: 36px;
  font-weight: 800;
  color: #1B7A6E;
  letter-spacing: 2px;
  margin-bottom: 32px;
}

/* ── Alerts ── */
.alert {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  padding: 11px 16px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.alert--error {
  background: #FEF2F2;
  border: 1px solid #FECACA;
  color: #DC2626;
}

.alert--success {
  background: #F0FDF4;
  border: 1px solid #86EFAC;
  color: #16A34A;
}

/* ── Form ── */
.form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 7px;
}

.form-label {
  font-size: 14px;
  font-weight: 600;
  color: #2C2C2C;
}

.input-wrap {
  position: relative;
}

.form-input {
  width: 100%;
  border: 1.5px solid #DDDDDD;
  border-radius: 10px;
  padding: 14px 18px;
  font-size: 14px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.form-input::placeholder {
  color: #BBBBBB;
}

.form-input:focus {
  border-color: #C8A135;
  box-shadow: 0 0 0 3px rgba(200, 161, 53, 0.1);
}

.form-input--has-icon {
  padding-right: 46px;
}

.form-input--error {
  border-color: #DC2626 !important;
}

.form-error {
  font-size: 11.5px;
  color: #DC2626;
}

.eye-btn {
  position: absolute;
  right: 14px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  color: #AAAAAA;
  display: flex;
  align-items: center;
  padding: 0;
  transition: color 0.2s;
}

.eye-btn:hover {
  color: #555;
}

/* ── Submit ── */
.btn-submit {
  margin-top: 6px;
  width: 220px;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  padding: 14px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background 0.2s, transform 0.2s;
}

.btn-submit:hover:not(:disabled) {
  background: #b08e2c;
  transform: translateY(-1px);
}

.btn-submit:disabled {
  opacity: 0.65;
  cursor: not-allowed;
}

.spinner {
  width: 17px;
  height: 17px;
  border: 2px solid rgba(255,255,255,0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
  display: inline-block;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ── Footer ── */
.footer-text {
  margin-top: 20px;
  font-size: 13px;
  color: #888;
}

.footer-link {
  color: #1B7A6E;
  font-weight: 700;
  text-decoration: none;
}

.footer-link:hover {
  text-decoration: underline;
}

/* ── Photo Grid ── */
.photo-grid {
  flex: 1;
  position: relative;
  min-height: 660px;
  align-self: stretch;
  overflow: visible;
}

.photo-grid__item {
  position: absolute;
  overflow: hidden;
  border-radius: 6px;
  opacity: 0.55;
  transition: opacity 0.3s ease;
}

.photo-grid__item:hover {
  opacity: 0.75;
}

.photo-grid__item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  transition: transform 0.4s ease;
  filter: saturate(0.5) brightness(1.1);
}

.photo-grid__item:hover img {
  transform: scale(1.04);
}

/* Foto 1: kiri atas, mulai agak ke bawah */
.photo-grid__item--1 {
  top: 70px;
  left: 0;
  width: 52%;
  height: 270px;
}

/* Foto 2: kanan atas, dari paling atas, terpotong di kanan */
.photo-grid__item--2 {
  top: 0;
  left: calc(52% + 10px);
  right: -60px;
  height: 330px;
}

/* Foto 3: kiri bawah */
.photo-grid__item--3 {
  top: 350px;
  left: 0;
  width: 58%;
  height: 260px;
}

/* Foto 4: kanan bawah, terpotong kanan, melampaui batas bawah */
.photo-grid__item--4 {
  top: 350px;
  left: calc(52% + 10px);
  right: -60px;
  height: 280px;
}

/* ── Responsive ── */
@media (max-width: 900px) {
  .main-content {
    flex-direction: column;
    padding: 40px 24px;
  }

  .form-section {
    flex: none;
    max-width: 100%;
    width: 100%;
  }

  .photo-grid {
    min-height: 380px;
    width: 100%;
  }

  .photo-grid__item--1 { top: 50px; width: 55%; height: 180px; }
  .photo-grid__item--2 { top: 0; right: 0; width: 46%; height: 220px; }
  .photo-grid__item--3 { top: 240px; width: 60%; height: 160px; }
  .photo-grid__item--4 { top: 280px; right: 0; width: 40%; height: 180px; }

  .navbar {
    padding: 0 24px;
    gap: 20px;
  }

  .navbar__links {
    display: none;
  }
}
</style>