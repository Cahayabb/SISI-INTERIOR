<template>
  <div class="about-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">About</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">About</span>
        </div>
      </div>
    </div>

    <!-- ── SECTION 1: Foto + Deskripsi ── -->
    <section class="about-intro section">
      <div class="container about-intro__grid">

        <!-- Kiri: Foto -->
        <div class="about-intro__images">
          <div class="about-intro__img-top">
            <img src="https://images.unsplash.com/photo-1600585154340-be6161a56a0c?w=700&q=80" alt="Interior 1" />
          </div>
          <div class="about-intro__img-row">
            <img src="https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=400&q=80" alt="Interior 2" />
            <img src="https://images.unsplash.com/photo-1616486338812-3dadae4b4ace?w=400&q=80" alt="Interior 3" />
          </div>
        </div>

        <!-- Kanan: Logo + Teks -->
        <div class="about-intro__text">
          <!-- Logo fallback teks -->
          <div class="about-intro__logo">
            <img
              src="@/assets/images/logo-sisi.png"
              alt="SISI Interior"
              class="about-intro__logo-img"
              @error="onLogoError"
              ref="logoImg"
            />
            <div class="about-intro__logo-fallback" ref="logoFallback" style="display:none">
              <span class="about-intro__logo-main">SISI</span>
              <span class="about-intro__logo-sub">INTERIOR</span>
            </div>
          </div>

          <p class="about-intro__desc">
            SISI INTERIOR berkedudukan sebagai perusahaan yang mewadahi bidang
            kontraktor, renovasi bangunan, atau pembuatan bangunan baru hingga
            akhirnya berkembang dalam bidang Desain Interior serta Eksterior, Hotel,
            Bungalow, Rumah, Apartemen, Gedung Kantor, Ruko, Restaurant,
            Gudang, serta Pabrik.
          </p>
          <p class="about-intro__desc">
            Memiliki visi untuk menciptakan atmosfer baru dalam penataan dan
            penempatan interior yang tepat dan sesuai pada fungsinya, kami bekerja
            dengan dasar inovasi dan kreatifitas yang segar untuk memberikan warna
            dalam setiap karya sehingga dapat menciptakan produk bersaing di masa
            depan. Dengan memperhatikan keefisienan waktu, kualitas, dan
            anggaran, SISI INTERIOR siap menjadi pendamping demi kenyamanan
            dan keindahan bangunan anda.
          </p>
        </div>

      </div>
    </section>

    <!-- ── SECTION 2: Visi + Foto Grid ── -->
    <section class="visi section">
      <div class="container visi__grid">

        <!-- Kiri: Visi teks -->
        <div class="visi__text">
          <h2 class="visi__title">Visi</h2>
          <p class="visi__desc">
            Menjadikan SISI INTERIOR sebagai perusahaan yang mumpuni,
            dapat dipercaya, dan mampu menciptakan atmosfer baru dalam
            bidang jasa konstruksi serta penataan ruang dan penempatan
            interior yang tepat dan sesuai pada fungsinya.
          </p>
        </div>

        <!-- Kanan: Foto 2x2 -->
        <div class="visi__images">
          <img src="https://images.unsplash.com/photo-1600210492493-0946911123ea?w=400&q=80" alt="Visi 1" />
          <img src="https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=400&q=80" alt="Visi 2" />
          <img src="https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=400&q=80" alt="Visi 3" />
          <img src="https://images.unsplash.com/photo-1497366216548-37526070297c?w=400&q=80" alt="Visi 4" />
        </div>

      </div>
    </section>

    <!-- ── SECTION 3: Misi ── -->
    <section class="misi section">
      <div class="container">
        <h2 class="misi__title">Misi</h2>
        <div class="misi__grid">

          <div class="misi__card">
            <div class="misi__icon">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#1B7A6E" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
                <polyline points="9,22 9,12 15,12 15,22"/>
                <path d="M12 2l2 4-2 1-2-1z"/>
              </svg>
            </div>
            <p class="misi__text">Memberikan hasil pekerjaan yang berkualitas</p>
          </div>

          <div class="misi__card">
            <div class="misi__icon">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#1B7A6E" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
            </div>
            <p class="misi__text">Menjalin hubungan integritas, dan kepercayaan</p>
          </div>

          <div class="misi__card">
            <div class="misi__icon">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#1B7A6E" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"/>
                <circle cx="9" cy="7" r="4"/>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"/>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
              </svg>
            </div>
            <p class="misi__text">Memberi pelayanan terbaik, jujur, profesional, berkualitas serta menjunjung tinggi harkat, derajat, dan persaudaraan</p>
          </div>

          <div class="misi__card">
            <div class="misi__icon">
              <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#1B7A6E" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="20 12 20 22 4 22 4 12"/>
                <rect x="2" y="7" width="20" height="5"/>
                <line x1="12" y1="22" x2="12" y2="7"/>
                <path d="M12 7H7.5a2.5 2.5 0 0 1 0-5C11 2 12 7 12 7z"/>
                <path d="M12 7h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z"/>
              </svg>
            </div>
            <p class="misi__text">Memberi nilai tambah untuk klien</p>
          </div>

        </div>
      </div>
    </section>

    <!-- ── SECTION 4: Mitra Banner ── -->
    <section class="mitra">
      <div class="mitra__content">
        <h2 class="mitra__title">Kami hadir sebagai mitra yang dapat dipercaya</h2>
        <p class="mitra__desc">
          Kami sangat terbuka dalam membangun hubungan kerja secara
          hangat, profesional, dan penuh komitmen demi memenuhi
          kebutuhan dan memberikan kepuasan kepada klien.
        </p>
      </div>
    </section>

  </div>
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
.about-page {
  padding-top: 72px; 
  font-family: 'Montserrat', sans-serif;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 40px;
}

.section {
  padding: 72px 0;
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

/* ── SECTION 1: Intro ── */
.about-intro__grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: start;
}

.about-intro__img-top {
  width: 100%;
  border-radius: 6px;
  overflow: hidden;
  margin-bottom: 12px;
  aspect-ratio: 4/3;
}

.about-intro__img-top img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.about-intro__img-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.about-intro__img-row img {
  width: 100%;
  aspect-ratio: 1/1;
  object-fit: cover;
  border-radius: 6px;
  display: block;
}

.about-intro__logo {
  margin-bottom: 24px;
}

.about-intro__logo-img {
  height: 52px;
  width: auto;
  object-fit: contain;
  display: block;
}

.about-intro__logo-fallback {
  flex-direction: column;
  line-height: 1;
}

.about-intro__logo-main {
  font-family: 'Playfair Display', serif;
  font-size: 22px;
  font-weight: 700;
  color: #1B7A6E;
  letter-spacing: 6px;
}

.about-intro__logo-sub {
  font-size: 8px;
  font-weight: 600;
  color: #C8A135;
  letter-spacing: 4px;
  text-transform: uppercase;
  margin-top: 2px;
}

.about-intro__desc {
  font-size: 13.5px;
  line-height: 1.9;
  color: #444;
  text-align: justify;
  margin: 0 0 20px;
}

/* ── SECTION 2: Visi ── */
.visi__grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: center;
}

.visi__title {
  font-size: 36px;
  font-weight: 800;
  color: #C8A135;
  margin: 0 0 24px;
}

.visi__desc {
  font-size: 13.5px;
  line-height: 1.9;
  color: #444;
  text-align: justify;
  margin: 0;
}

.visi__images {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.visi__images img {
  width: 100%;
  aspect-ratio: 1/1;
  object-fit: cover;
  border-radius: 6px;
  display: block;
}

/* ── SECTION 3: Misi ── */
.misi {
  background: #FFFFFF;
}

.misi__title {
  font-size: 36px;
  font-weight: 800;
  color: #C8A135;
  text-align: center;
  margin: 0 0 48px;
}

.misi__grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.misi__card {
  background: #FAFAFA;
  border: 1px solid #EEEEEE;
  border-radius: 10px;
  padding: 36px 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 20px;
  transition: box-shadow 0.2s ease, transform 0.2s ease;
}

.misi__card:hover {
  box-shadow: 0 8px 24px rgba(0,0,0,0.08);
  transform: translateY(-4px);
}

.misi__icon {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  background: rgba(27, 122, 110, 0.08);
  display: flex;
  align-items: center;
  justify-content: center;
}

.misi__text {
  font-size: 13px;
  line-height: 1.7;
  color: #1B7A6E;
  font-weight: 500;
  margin: 0;
}

/* ── SECTION 4: Mitra Banner ── */
.mitra {
  background: #F0F7F6;
  padding: 80px 40px;
  text-align: center;
}

.mitra__content {
  max-width: 600px;
  margin: 0 auto;
}

.mitra__title {
  font-size: 32px;
  font-weight: 800;
  color: #1B7A6E;
  line-height: 1.3;
  margin: 0 0 24px;
}

.mitra__desc {
  font-size: 14px;
  line-height: 1.9;
  color: #555;
  margin: 0;
  text-align: center;
}

/* ── Responsive ── */
@media (max-width: 860px) {
  .about-intro__grid,
  .visi__grid {
    grid-template-columns: 1fr;
    gap: 36px;
  }

  .misi__grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 560px) {
  .misi__grid {
    grid-template-columns: 1fr;
  }

  .container {
    padding: 0 20px;
  }
}
</style>