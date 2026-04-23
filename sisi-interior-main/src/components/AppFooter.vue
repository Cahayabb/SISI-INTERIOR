<template>
  <footer class="footer">
    <div class="container footer__grid">

      <!-- Logo + Deskripsi + Sosmed -->
      <div class="footer__brand">
        <RouterLink to="/" class="footer__logo">
          <img
            src="@/assets/images/logo-sisi.png"
            alt="SISI Interior"
            class="footer__logo-img"
            @error="onLogoError"
            ref="logoImg"
          />
          <div class="footer__logo-fallback" ref="logoFallback" style="display:none">
            <span class="footer__logo-main">SISI</span>
            <span class="footer__logo-sub">INTERIOR</span>
          </div>
        </RouterLink>

        <p class="footer__desc">
          SISI Interior adalah perusahaan yang melayani
          desain interior, kontraktor, dan custom furniture
          untuk hunian.
        </p>

        <div class="footer__socials">
          <!-- Facebook -->
          <a href="#" class="footer__social" aria-label="Facebook">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>
            </svg>
          </a>
          <!-- Twitter -->
          <a href="#" class="footer__social" aria-label="Twitter">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path d="M23 3a10.9 10.9 0 0 1-3.14 1.53A4.48 4.48 0 0 0 22.43.36a9 9 0 0 1-2.88 1.1A4.52 4.52 0 0 0 16.11 0c-2.5 0-4.52 2-4.52 4.5 0 .35.04.7.11 1.03C7.69 5.3 4.07 3.56 1.64.9a4.48 4.48 0 0 0-.61 2.27c0 1.56.8 2.94 2 3.75a4.5 4.5 0 0 1-2.05-.56v.06c0 2.18 1.55 4 3.6 4.42a4.52 4.52 0 0 1-2.04.08 4.53 4.53 0 0 0 4.22 3.13A9.05 9.05 0 0 1 0 19.54 12.77 12.77 0 0 0 6.92 21.5c8.3 0 12.85-6.88 12.85-12.85 0-.2 0-.39-.01-.58A9.17 9.17 0 0 0 23 3z"/>
            </svg>
          </a>
          <!-- Instagram -->
          <a href="#" class="footer__social" aria-label="Instagram">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="2" y="2" width="20" height="20" rx="5" ry="5"/>
              <circle cx="12" cy="12" r="4"/>
              <circle cx="17.5" cy="6.5" r="1" fill="currentColor" stroke="none"/>
            </svg>
          </a>
          <!-- LinkedIn -->
          <a href="#" class="footer__social" aria-label="LinkedIn">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
              <path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"/>
              <rect x="2" y="9" width="4" height="12"/>
              <circle cx="4" cy="4" r="2"/>
            </svg>
          </a>
        </div>
      </div>

      <!-- Layanan -->
      <div class="footer__col">
        <h4 class="footer__col-title">Layanan</h4>
        <ul class="footer__list">
          <li><RouterLink to="/service#interior" class="footer__link">Desain Interior</RouterLink></li>
          <li><RouterLink to="/service#contractor" class="footer__link">Kontraktor</RouterLink></li>
          <li><RouterLink to="/service#furniture" class="footer__link">Custom Furniture</RouterLink></li>
        </ul>
      </div>

      <!-- Navigasi -->
      <div class="footer__col">
        <h4 class="footer__col-title">Navigasi</h4>
        <ul class="footer__list">
          <li><RouterLink to="/" class="footer__link">Home</RouterLink></li>
          <li><RouterLink to="/about" class="footer__link">About Us</RouterLink></li>
          <li><RouterLink to="/service" class="footer__link">Services</RouterLink></li>
          <li><RouterLink to="/portfolio" class="footer__link">Portfolio</RouterLink></li>
          <li><RouterLink to="/contact" class="footer__link">Contact</RouterLink></li>
        </ul>
      </div>

      <!-- Kontak -->
      <div class="footer__col">
        <h4 class="footer__col-title">Kontak</h4>
        <ul class="footer__list">
          <li class="footer__contact-item">0812-9052-9092</li>
          <li class="footer__contact-item">Kota Bogor, Jawa barat,<br />Indonesia</li>
        </ul>
      </div>

    </div>
  </footer>
</template>

<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'

const logoImg      = ref(null)
const logoFallback = ref(null)

const onLogoError = () => {
  if (logoImg.value)      logoImg.value.style.display = 'none'
  if (logoFallback.value) logoFallback.value.style.display = 'flex'
}
</script>

<style scoped>
.footer {
  background: #FFFFFF;
  border-top: 1px solid #EEEEEE;
  padding: 60px 0 40px;
  box-shadow: 0 -4px 24px rgba(0, 0, 0, 0.1);
}

.footer__grid {
  display: grid;
  grid-template-columns: 1.4fr 1fr 1fr 1fr;
  gap: 48px;
  align-items: start;
}

.footer__brand {
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding-left: 70px;
}

.footer__logo {
  display: inline-flex;
  align-items: center;
  text-decoration: none;
}

.footer__logo-img {
  height: 48px;
  width: auto;
  object-fit: contain;
  display: block;
}

.footer__logo-fallback {
  flex-direction: column;
  line-height: 1;
}

.footer__logo-main {
  font-family: 'Playfair Display', 'Georgia', serif;
  font-size: 20px;
  font-weight: 700;
  color: #1B7A6E;
  letter-spacing: 6px;
}

.footer__logo-sub {
  font-family: 'Montserrat', sans-serif;
  font-size: 7px;
  font-weight: 600;
  color: #C8A135;
  letter-spacing: 4px;
  text-transform: uppercase;
  margin-top: 2px;
}

.footer__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  line-height: 1.85;
  color: #888888;
  margin: 0;
  max-width: 220px;
}

/* ── Sosial Media ── */
.footer__socials {
  display: flex;
  gap: 10px;
}

.footer__social {
  width: 36px;
  height: 36px;
  border-radius: 6px;
  background: #1B7A6E;
  color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  text-decoration: none;
  transition: background 0.2s ease, transform 0.2s ease;
}

.footer__social:hover {
  background: #145F56;
  transform: translateY(-2px);
}

/* ── Kolom ── */
.footer__col-title {
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 700;
  color: #2C2C2C;
  margin: 0 0 24px;
}

.footer__list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.footer__link {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  color: #888888;
  text-decoration: none;
  transition: color 0.2s ease;
}

.footer__link:hover {
  color: #1B7A6E;
}

.footer__contact-item {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  color: #888888;
  line-height: 1.7;
}

@media (max-width: 900px) {
  .footer__grid {
    grid-template-columns: 1fr 1fr;
    gap: 36px;
  }
}

@media (max-width: 560px) {
  .footer__grid {
    grid-template-columns: 1fr;
  }
}
</style>