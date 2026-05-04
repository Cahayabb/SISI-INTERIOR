<template>
  <div class="login-page">

    <div class="photo-strip">
      <div
        v-for="(url, i) in photoUrlsTop"
        :key="'t'+i"
        class="photo-strip__img"
        :style="{ backgroundImage: `url('${url}')` }"
      ></div>
    </div>

    <div class="login-body">
      <div class="login-box">

        <h1 class="login-title">MASUK</h1>

        <div class="login-alert" v-if="errorMsg">
          <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <line x1="12" y1="8" x2="12" y2="12"/>
            <line x1="12" y1="16" x2="12.01" y2="16"/>
          </svg>
          {{ errorMsg }}
        </div>

        <div class="login-form">

          <div class="field">
            <label class="field__label">Username</label>
            <input
              v-model="form.username"
              type="text"
              class="field__input"
              :class="{ 'field__input--error': errors.username }"
              placeholder="Masukkan username"
              @keyup.enter="handleLogin"
            />
            <span class="field__error" v-if="errors.username">{{ errors.username }}</span>
          </div>

          <div class="field">
            <label class="field__label">Password</label>
            <div class="field__wrap">
              <input
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                class="field__input"
                :class="{ 'field__input--error': errors.password }"
                placeholder="Minimal 8 karakter"
                @keyup.enter="handleLogin"
              />
              <button class="field__eye" type="button" @click="showPassword = !showPassword">
                <svg v-if="!showPassword" width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                  <circle cx="12" cy="12" r="3"/>
                </svg>
                <svg v-else width="17" height="17" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M17.94 17.94A10.07 10.07 0 0112 20c-7 0-11-8-11-8a18.45 18.45 0 015.06-5.94"/>
                  <path d="M9.9 4.24A9.12 9.12 0 0112 4c7 0 11 8 11 8a18.5 18.5 0 01-2.16 3.19"/>
                  <line x1="1" y1="1" x2="23" y2="23"/>
                </svg>
              </button>
            </div>
            <span class="field__error" v-if="errors.password">{{ errors.password }}</span>
          </div>

          <div class="login-forgot">
            <a href="#" class="login-forgot__link">Lupa password?</a>
          </div>

          <button class="login-btn" @click="handleLogin" :disabled="isLoading">
            <span v-if="isLoading" class="login-btn__spinner"></span>
            <span v-else>Login</span>
          </button>

          <p class="login-register">
            Belum punya akun?
            <RouterLink to="/signup" class="login-register__link">Register</RouterLink>
          </p>

        </div>
      </div>
    </div>

    <div class="photo-strip">
      <div
        v-for="(url, i) in photoUrlsBottom"
        :key="'b'+i"
        class="photo-strip__img"
        :style="{ backgroundImage: `url('${url}')` }"
      ></div>
    </div>

  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { login } from '@/services/auth'


const router       = useRouter()
const isLoading    = ref(false)
const showPassword = ref(false)
const errorMsg     = ref('')

const form   = reactive({ username: '', password: '' })
const errors = reactive({ username: '', password: '' })

const photoUrlsTop = [
  'https://images.unsplash.com/photo-1555041469-a586c61ea9bc?w=800&q=80',
  'https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=800&q=80',
  'https://images.unsplash.com/photo-1567538096630-e0c55bd6374c?w=800&q=80',
  'https://images.unsplash.com/photo-1616486338812-3dadae4b4ace?w=800&q=80',
  'https://images.unsplash.com/photo-1600210492493-0946911123ea?w=800&q=80',
]

const photoUrlsBottom = [
  'https://images.unsplash.com/photo-1618221195710-dd6b41faaea6?w=800&q=80',
  'https://images.unsplash.com/photo-1631679706909-1844bbd07221?w=800&q=80',
  'https://images.unsplash.com/photo-1600585154526-990dced4db0d?w=800&q=80',
  'https://images.unsplash.com/photo-1583847268964-b28dc8f51f92?w=800&q=80',
  'https://images.unsplash.com/photo-1524758631624-e2822e304c36?w=800&q=80',
]

const validate = () => {
  let valid = true
  errors.username = ''
  errors.password = ''

  if (!form.username.trim()) {
    errors.username = 'Username wajib diisi.'
    valid = false
  }

  if (!form.password) {
    errors.password = 'Password wajib diisi.'
    valid = false
  }

  return valid
}

const handleLogin = async () => {
  console.log("LOGIN DIKLIK")

  errorMsg.value = ''
  if (!validate()) return

  isLoading.value = true

  try {
    const res = await login({
      username: form.username,
      password: form.password,
    })

    console.log('Response:', res.data)


 const user = res.data?.user

    if (!user) {
      throw new Error("User tidak ditemukan di response")
    }

    // simpan session
    localStorage.setItem('token', '123')
    localStorage.setItem('role', user.role)
    localStorage.setItem('username', user.username)

    console.log("TOKEN SET")
    console.log("ROLE:", user.role)
    console.log("FULL RESPONSE:", JSON.stringify(res.data, null, 2))

    // redirect berdasarkan role
    if (user.role === 'admin') {
      router.push('/admin/dashboard')
    } else {
      router.push('/')
    }

  } catch (error) {
    console.error("ERROR LOGIN:", error)
    errorMsg.value = error.response?.data?.error || 'Login gagal'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-page {
  height: 100vh;
  padding-top: 72px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
}
.photo-strip {
  height: 160px;
  min-height: 160px;   
  max-height: 160px;   
  flex-shrink: 0;
  display: flex;
  gap: 8px;
  padding: 8px;
  box-sizing: border-box;
  overflow: hidden;
}

.photo-strip__img {
  flex: 1;                       
  min-width: 0;
  height: 100%;                   
  background-size: cover;
  background-position: center 40%;
  background-repeat: no-repeat;
  border-radius: 8px;
  opacity: 0.38;
  transition: opacity 0.25s;
}

.photo-strip__img:hover { opacity: 0.56; }

/* ─────────────────────────────────────────────
   FORM AREA 
   ───────────────────────────────────────────── */
.login-body {
  flex: 1;
  min-height: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-box {
  width: 100%;
  max-width: 400px;
  padding: 0 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

/* ── Title ── */
.login-title {
  font-size: 30px;
  font-weight: 900;
  color: #1B7A6E;
  letter-spacing: 6px;
  text-align: center;
  margin: 0 0 24px;
}

/* ── Alert ── */
.login-alert {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #FEF2F2;
  border: 1px solid #FECACA;
  color: #DC2626;
  font-size: 13px;
  padding: 10px 14px;
  border-radius: 8px;
  margin-bottom: 14px;
  width: 100%;
}

/* ── Form ── */
.login-form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.field__label {
  font-size: 13px;
  font-weight: 600;
  color: #2C2C2C;
}

.field__wrap { position: relative; }

.field__input {
  width: 100%;
  border: 1.5px solid #D1D5DB;
  border-radius: 8px;
  padding: 12px 16px;
  font-size: 13.5px;
  font-family: 'Montserrat', sans-serif;
  color: #333;
  background: #FFFFFF;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
}

.field__input:focus {
  border-color: #1B7A6E;
  box-shadow: 0 0 0 3px rgba(27,122,110,0.08);
}

.field__input--error       { border-color: #DC2626 !important; }
.field__input::placeholder { color: #B0B0B0; }
.field__wrap .field__input { padding-right: 46px; }

.field__eye {
  position: absolute; right: 13px; top: 50%;
  transform: translateY(-50%);
  background: none; border: none; cursor: pointer;
  color: #9CA3AF; padding: 0;
  display: flex; align-items: center; transition: color 0.2s;
}
.field__eye:hover { color: #4B5563; }

.field__error { font-size: 11px; color: #DC2626; }

/* ── Lupa password ── */
.login-forgot { text-align: right; }

.login-forgot__link {
  font-size: 12px; color: #1B7A6E;
  text-decoration: none; font-weight: 500;
}
.login-forgot__link:hover { text-decoration: underline; }

/* ── Submit ── */
.login-btn {
  width: 100%;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
  letter-spacing: 1px;
  padding: 14px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.2s, transform 0.15s;
  display: flex; align-items: center; justify-content: center;
}
.login-btn:hover:not(:disabled) { background: #b08e2c; transform: translateY(-1px); }
.login-btn:disabled              { opacity: 0.7; cursor: not-allowed; }

.login-btn__spinner {
  width: 17px; height: 17px;
  border: 2px solid rgba(255,255,255,0.35);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }

/* ── Register ── */
.login-register {
  text-align: center; font-size: 13px;
  color: #6B7280; margin: 0;
}
.login-register__link {
  color: #1B7A6E; font-weight: 700; text-decoration: none;
}
.login-register__link:hover { text-decoration: underline; }

/* ── Responsive ── */
@media (max-width: 600px) {
  .photo-strip           { height: 110px; min-height: 110px; max-height: 110px; }
  .login-title           { font-size: 24px; }
}
</style>