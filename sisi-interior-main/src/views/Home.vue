<template>
  <div class="home">

    <!-- ═══════════════════════════════════════════
         HERO SECTION
    ═══════════════════════════════════════════ -->
    <section class="hero">
    <div class="hero__bg">
        <img src="@/assets/images/hero.png" alt="SISI Interior Hero" class="hero__img" />
        <div class="hero__overlay" />
    </div>

    <div class="hero__content">
        <p class="hero__eyebrow">SISI INTERIOR</p>

        <h1 class="hero__title">
        Transforming <span class="hero__title--gold">Ideas</span> into<br />
        Timeless <span class="hero__title--gold">Interiors</span>
        </h1>

        <p class="hero__desc">
        SISI Interior menghadirkan solusi desain interior, kontraktor, dan custom furniture
        untuk menciptakan ruang yang estetik, fungsional,
        dan sesuai dengan kebutuhan Anda.
        </p>

        <a href="https://wa.me/6281290529092" target="_blank" rel="noopener" class="hero__cta">
        <svg class="hero__cta-icon" viewBox="0 0 24 24" fill="currentColor">
            <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347z"/>
            <path d="M12 0C5.373 0 0 5.373 0 12c0 2.123.554 4.116 1.523 5.843L.057 23.571a.75.75 0 0 0 .92.92l5.698-1.458A11.945 11.945 0 0 0 12 24c6.627 0 12-5.373 12-12S18.627 0 12 0zm0 21.75a9.733 9.733 0 0 1-4.964-1.357l-.356-.211-3.685.944.974-3.595-.232-.37A9.733 9.733 0 0 1 2.25 12C2.25 6.615 6.615 2.25 12 2.25S21.75 6.615 21.75 12 17.385 21.75 12 21.75z"/>
        </svg>
        Konsultasi Sekarang
        </a>
    </div>
    </section>

    <!-- ═══════════════════════════════════════════
         ABOUT SECTION
    ═══════════════════════════════════════════ -->
    <section class="about section" id="about">
    <div class="container">

        <!-- Baris atas: Foto besar kiri + Teks kanan -->
        <div class="about__top">
        <div class="about__img-main">
            <img src="@/assets/images/about-1.png" alt="Interior Office" />
        </div>
        <div class="about__text">
            <div class="about__label">
            <span class="about__label-line" />
            <h2 class="about__title">Tentang Kami</h2>
            </div>
            <p class="about__desc">
            SISI Interior adalah perusahaan yang bergerak di bidang desain interior,
            kontraktor, dan custom furniture. Kami berkomitmen menghadirkan solusi
            ruang yang tidak hanya menarik secara visual, tetapi juga nyaman,
            fungsional, dan sesuai dengan kebutuhan setiap klien
            </p>
            <RouterLink to="/about" class="about__btn">Selengkapnya</RouterLink>
        </div>
        </div>

        <!-- Baris bawah: 3 foto sejajar full width -->
        <div class="about__bottom">
        <img src="@/assets/images/about-2.png" alt="Interior 2" />
        <img src="@/assets/images/about-3.png" alt="Interior 3" />
        <img src="@/assets/images/about-4.png" alt="Interior 4" />
        </div>

    </div>
    </section>

    <!-- ═══════════════════════════════════════════
         SERVICES SECTION
    ═══════════════════════════════════════════ -->
    <section class="services section" id="services">
    <div class="container">

        <div class="services__label">
        <span class="services__label-line" />
        <h2 class="services__title">Layanan Kami</h2>
        </div>

        <div class="services__grid">

        <div class="service-item">
            <div class="service-item__icon">
            <!-- Icon Desain Interior: lampu gantung -->
            <svg width="52" height="52" viewBox="0 0 64 64" fill="none" stroke="#C8A135" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="32" y1="4" x2="32" y2="10"/>
                <path d="M20 30 Q20 18 32 18 Q44 18 44 30 L40 42 H24 Z"/>
                <line x1="24" y1="42" x2="40" y2="42"/>
                <line x1="26" y1="46" x2="38" y2="46"/>
                <line x1="28" y1="50" x2="36" y2="50"/>
                <line x1="14" y1="10" x2="20" y2="18"/>
                <line x1="50" y1="10" x2="44" y2="18"/>
            </svg>
            </div>
            <div class="service-item__body">
            <h3 class="service-item__title">Desain Interior</h3>
            <p class="service-item__desc">Kami merancang interior yang estetik, fungsional, dan sesuai dengan karakter serta kebutuhan ruang Anda.</p>
            </div>
        </div>

        <div class="service-item">
            <div class="service-item__icon">
            <!-- Icon Kontraktor: pohon/tanaman -->
            <svg width="52" height="52" viewBox="0 0 64 64" fill="none" stroke="#C8A135" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="32" y1="56" x2="32" y2="30"/>
                <path d="M32 30 Q22 30 18 20 Q14 10 24 8 Q26 4 32 6 Q38 4 40 8 Q50 10 46 20 Q42 30 32 30Z"/>
                <path d="M32 42 Q24 42 20 34"/>
                <path d="M32 42 Q40 42 44 34"/>
                <line x1="24" y1="56" x2="40" y2="56"/>
            </svg>
            </div>
            <div class="service-item__body">
            <h3 class="service-item__title">Kontraktor</h3>
            <p class="service-item__desc">Kami melayani pelaksanaan pembangunan dan renovasi dengan proses kerja yang terarah, rapi, dan efisien.</p>
            </div>
        </div>

        <div class="service-item">
            <div class="service-item__icon">
            <!-- Icon Custom Furniture: sofa/kursi -->
            <svg width="52" height="52" viewBox="0 0 64 64" fill="none" stroke="#C8A135" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="10" y="28" width="44" height="16" rx="4"/>
                <rect x="16" y="20" width="32" height="10" rx="3"/>
                <rect x="10" y="28" width="8" height="20" rx="3"/>
                <rect x="46" y="28" width="8" height="20" rx="3"/>
                <line x1="16" y1="44" x2="16" y2="52"/>
                <line x1="48" y1="44" x2="48" y2="52"/>
            </svg>
            </div>
            <div class="service-item__body">
            <h3 class="service-item__title">Custom Furniture</h3>
            <p class="service-item__desc">Kami menyediakan pembuatan furniture custom yang dirancang sesuai ukuran, konsep, dan kebutuhan ruang Anda.</p>
            </div>
        </div>

        </div>
    </div>
    </section>

    <!-- ═══════════════════════════════════════════
         PORTFOLIO SECTION
    ═══════════════════════════════════════════ -->
    <section class="portfolio section" id="portfolio">
    <div class="portfolio__inner">

        <!-- Kiri: Teks -->
        <div class="portfolio__text">
        <div class="portfolio__label">
            <span class="portfolio__label-line" />
            <h2 class="portfolio__title-gold">Portofolio</h2>
        </div>
        <h3 class="portfolio__heading">Hasil Karya yang Telah Kami Wujudkan</h3>
        <p class="portfolio__desc">
            Berbagai proyek telah kami kerjakan untuk menghadirkan ruang yang nyaman, estetik,
            dan sesuai kebutuhan klien. Lihat beberapa hasil pekerjaan terbaik kami sebagai inspirasi
            untuk proyek Anda.
        </p>
        <RouterLink to="/portfolio" class="portfolio__btn">Lihat Portofolio Lengkap</RouterLink>
        </div>

        <!-- Kanan: Slider -->
        <div class="portfolio__slider-wrap">

        <!-- Arrow kiri -->
        <button class="portfolio__arrow portfolio__arrow--left" @click="prevSlide" aria-label="Prev">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M15 18l-6-6 6-6"/>
            </svg>
        </button>

        <!-- Slides -->
        <div class="portfolio__slider">
            <div
            v-for="(item, i) in portfolioItems"
            :key="i"
            class="portfolio__slide"
            :class="{ 'portfolio__slide--active': i === activeSlide }"
            >
            <img :src="item.img" :alt="item.title" class="portfolio__slide-img" />

            <!-- Info hanya muncul di slide aktif -->
            <div class="portfolio__slide-info" v-if="i === activeSlide">
                <div class="portfolio__slide-num">
                <span class="portfolio__slide-line" />
                Project {{ String(i + 1).padStart(2, '0') }}
                </div>
                <h4 class="portfolio__slide-title">{{ item.title }}</h4>
                <p class="portfolio__slide-desc">{{ item.desc }}</p>
            </div>
            </div>
        </div>

        <!-- Arrow kanan -->
        <button class="portfolio__arrow portfolio__arrow--right" @click="nextSlide" aria-label="Next">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 18l6-6-6-6"/>
            </svg>
        </button>

        <!-- Dots -->
        <div class="portfolio__dots">
            <button
            v-for="(item, i) in portfolioItems"
            :key="i"
            class="portfolio__dot"
            :class="{ 'portfolio__dot--active': i === activeSlide }"
            @click="goToSlide(i)"
            />
        </div>

        </div>
    </div>
    </section>

    <!-- ═══════════════════════════════════════════
         COST ESTIMATION CTA SECTION
    ═══════════════════════════════════════════ -->
    <section class="estimasi section" id="estimasi">
    <div class="container estimasi__grid">

        <!-- Kiri: Foto overlap -->
        <div class="estimasi__images">
        <img src="https://images.unsplash.com/photo-1618221195710-dd6b41faaea6?w=600&q=80" alt="Estimasi 1" class="estimasi__img estimasi__img--top" />
        <img src="https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=600&q=80" alt="Estimasi 2" class="estimasi__img estimasi__img--bottom" />
        </div>

        <!-- Kanan: Teks -->
        <div class="estimasi__text">
        <div class="estimasi__label">
            <span class="estimasi__label-line" />
            <h2 class="estimasi__title-gold">Estimasi Biaya</h2>
        </div>
        <h3 class="estimasi__heading">Dapatkan Gambaran Estimasi Biaya Proyek Anda</h3>
        <p class="estimasi__desc">
            Masukkan kebutuhan dasar proyek Anda untuk memperoleh estimasi awal
            biaya secara cepat. Fitur ini membantu Anda mendapatkan gambaran
            harga sebelum melakukan konsultasi lebih lanjut.
        </p>
        <RouterLink to="/estimasi" class="estimasi__btn">Coba Estimasi Sekarang</RouterLink>
        </div>

    </div>
    </section>

    <!-- ═══════════════════════════════════════════
         TESTIMONIALS SECTION
    ═══════════════════════════════════════════ -->
    <section class="testimoni section" id="testimoni">
    <div class="container">

        <div class="testimoni__inner">

        <!-- Kiri: Foto + Arrow -->
        <div class="testimoni__left">
            <img src="https://images.unsplash.com/photo-1524758631624-e2822e304c36?w=600&q=80" alt="Testimoni" class="testimoni__img" />
            <div class="testimoni__arrows">
            <button class="testimoni__arrow" @click="prevTestimoni" aria-label="Prev">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M15 18l-6-6 6-6"/></svg>
            </button>
            <button class="testimoni__arrow" @click="nextTestimoni" aria-label="Next">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M9 18l6-6-6-6"/></svg>
            </button>
            </div>
        </div>

        <!-- Kanan: Label + Cards -->
        <div class="testimoni__right">
            <div class="testimoni__label">
            <span class="testimoni__label-line" />
            <h2 class="testimoni__title-gold">Testimoni</h2>
            </div>
            <h3 class="testimoni__heading">Apa Kata Klien Kami</h3>

            <div class="testimoni__cards">
            <div
                v-for="(t, i) in visibleTestimoni"
                :key="i"
                class="testimoni__card"
            >
                <div class="testimoni__quote">"</div>
                <p class="testimoni__text">{{ t.text }}</p>
                <div class="testimoni__author">
                <img :src="t.avatar" :alt="t.name" class="testimoni__avatar" />
                <div>
                    <div class="testimoni__name">{{ t.name }}</div>
                    <div class="testimoni__role">{{ t.role }}</div>
                </div>
                </div>
            </div>
            </div>
        </div>

        </div>
    </div>
    </section>

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { RouterLink } from 'vue-router'

/* ─── Services Data ─── */
const services = [
  {
    title: 'Desain Interior',
    desc: 'Kami merancang interior yang estetik, fungsional, dan sesuai dengan karakter serta kebutuhan ruang Anda.',
    link: '/service#interior',
    icon: `<svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
      <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
      <polyline points="9,22 9,12 15,12 15,22"/>
    </svg>`,
  },
  {
    title: 'Kontraktor',
    desc: 'Kami melayani pelaksanaan pembangunan dan renovasi dengan proses kerja yang terarah, rapi, dan efisien.',
    link: '/service#contractor',
    icon: `<svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
      <rect x="2" y="3" width="20" height="14" rx="2"/>
      <line x1="8" y1="21" x2="16" y2="21"/>
      <line x1="12" y1="17" x2="12" y2="21"/>
    </svg>`,
  },
  {
    title: 'Custom Furniture',
    desc: 'Kami menyediakan pembuatan furniture custom yang dirancang sesuai ukuran, konsep, dan kebutuhan ruang Anda.',
    link: '/service#furniture',
    icon: `<svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 7H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
      <path d="M6 7V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2"/>
      <line x1="6" y1="17" x2="6" y2="21"/>
      <line x1="18" y1="17" x2="18" y2="21"/>
    </svg>`,
  },
]

/* ─── Portfolio Data ─── */
const activeSlide = ref(0)
const portfolioItems = [
  {
    title: 'Kitchen Set Minimalis',
    desc: 'Desain interior dapur dengan konsep modern dan fungsional',
    img: 'https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=600&q=80',
  },
  {
    title: 'Ruang Makan Elegan',
    desc: 'Desain ruang makan dengan nuansa hangat dan natural',
    img: 'https://images.unsplash.com/photo-1617806118233-18e1de247200?w=600&q=80',
  },
  {
    title: 'Bedroom Modern',
    desc: 'Kamar tidur minimalis dengan pencahayaan yang nyaman',
    img: 'https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=600&q=80',
  },
  {
    title: 'Living Room Scandinavian',
    desc: 'Ruang keluarga dengan konsep Scandinavian yang bersih',
    img: 'https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=600&q=80',
  },
]

const nextSlide = () => {
  activeSlide.value = (activeSlide.value + 1) % portfolioItems.length
}

const prevSlide = () => {
  activeSlide.value = (activeSlide.value - 1 + portfolioItems.length) % portfolioItems.length
}

const goToSlide = (i) => {
  activeSlide.value = i
}

/* ─── Testimonials Data ─── */
const testimoniList = [
  {
    text: 'Pelayanan dari SISI Interior sangat membantu dari awal konsultasi hingga hasil akhir. Desainnya sesuai kebutuhan dan pengerjaannya juga rapi.',
    name: 'Sophie Carter',
    role: 'Klien Hunian',
    avatar: 'https://i.pravatar.cc/80?img=47',
  },
  {
    text: 'Custom furniture yang dibuat sangat pas dengan ukuran ruangan dan desainnya juga sesuai ekspektasi. Sangat recommended untuk yang ingin hasil rapi dan estetik.',
    name: 'James Bennett',
    role: 'Klien Custom Furniture',
    avatar: 'https://i.pravatar.cc/80?img=12',
  },
  {
    text: 'Tim SISI Interior sangat profesional dan responsif. Hasil renovasi rumah kami melebihi ekspektasi, sangat puas dengan hasilnya.',
    name: 'Rina Kusuma',
    role: 'Klien Renovasi',
    avatar: 'https://i.pravatar.cc/80?img=25',
  },
  {
    text: 'Proses pengerjaan cepat dan hasilnya rapi sekali. Desain kantor kami jadi lebih nyaman dan produktif setelah ditangani SISI Interior.',
    name: 'Budi Santoso',
    role: 'Klien Office',
    avatar: 'https://i.pravatar.cc/80?img=33',
  },
]

const activeTestimoni = ref(0)

const visibleTestimoni = computed(() => {
  const total = testimoniList.length
  const first = activeTestimoni.value % total
  const second = (activeTestimoni.value + 1) % total
  return [testimoniList[first], testimoniList[second]]
})

const nextTestimoni = () => {
  activeTestimoni.value = (activeTestimoni.value + 1) % testimoniList.length
}

const prevTestimoni = () => {
  activeTestimoni.value = (activeTestimoni.value - 1 + testimoniList.length) % testimoniList.length
}
</script>

<style scoped>
/* ════════════════════════════════
   GLOBAL TOKENS
════════════════════════════════ */
.home {
  --gold: #C8A951;
  --gold-dark: #b8962e;
  --teal: #1B7A6E;
  --teal-dark: #145F56;
  --teal-light: #e8f4f3;
  --white: #FFFFFF;
  --bg: #FAFAF8;
  --text: #2C2C2C;
  --text-muted: #6B6B6B;
  --radius: 16px;
  --radius-sm: 10px;

  font-family: 'Montserrat', 'Segoe UI', sans-serif;
  color: var(--text);
  background: var(--bg);
  overflow-x: hidden;
}

/* ════════════════════════════════
   UTILITY
════════════════════════════════ */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 32px;
}

.section {
  padding: 100px 0;
}

.section__label {
  display: inline-block;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 4px;
  text-transform: uppercase;
  color: var(--gold);
  margin-bottom: 16px;
  position: relative;
  padding-left: 40px;
}

.section__label::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 28px;
  height: 2px;
  background: var(--gold);
}

.section__label--light {
  color: rgba(255, 255, 255, 0.7);
}
.section__label--light::before {
  background: rgba(255, 255, 255, 0.5);
}

.section__title {
  font-family: 'Playfair Display', 'Georgia', serif;
  font-size: clamp(28px, 4vw, 42px);
  font-weight: 700;
  line-height: 1.2;
  color: var(--text);
  margin: 0 0 24px;
}

.section__title--light {
  color: var(--white);
}

.section__header {
  margin-bottom: 60px;
}

.section__header.center {
  text-align: center;
}

.section__header.center .section__label {
  padding-left: 0;
}

.section__header.center .section__label::before {
  display: none;
}

.text-gold { color: var(--gold); }

/* ════════════════════════════════
   BUTTONS
════════════════════════════════ */
.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 14px 28px;
  border-radius: 8px;
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.3px;
  text-decoration: none;
  transition: all 0.25s ease;
  cursor: pointer;
}

.btn--primary {
  background: var(--teal);
  color: var(--white);
  box-shadow: 0 4px 20px rgba(27, 122, 110, 0.35);
}
.btn--primary:hover {
  background: var(--teal-dark);
  transform: translateY(-2px);
  box-shadow: 0 8px 28px rgba(27, 122, 110, 0.45);
}

.btn--outline {
  background: transparent;
  color: var(--teal);
  border: 2px solid var(--teal);
}
.btn--outline:hover {
  background: var(--teal);
  color: var(--white);
  transform: translateY(-2px);
}

.btn--gold {
  background: var(--gold);
  color: var(--white);
  box-shadow: 0 4px 20px rgba(200, 169, 81, 0.4);
}
.btn--gold:hover {
  background: var(--gold-dark);
  transform: translateY(-2px);
  box-shadow: 0 8px 28px rgba(200, 169, 81, 0.5);
}

/* ════════════════════════════════
   HERO
════════════════════════════════ */
.hero {
  position: relative;
  width: 100%;
  height: 100vh;
  min-height: 600px;
  display: flex;
  align-items: center;
  overflow: hidden;
}

.hero__bg {
  position: absolute;
  inset: 0;
  z-index: 0;
}

.hero__img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  display: block;
}

.hero__overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(
    to right,
    rgba(0, 0, 0, 0.78) 0%,
    rgba(0, 0, 0, 0.60) 40%,
    rgba(0, 0, 0, 0.25) 75%,
    rgba(0, 0, 0, 0.10) 100%
  );
}

.hero__content {
  position: relative;
  z-index: 2;
  padding: 0 80px;
  padding-left: 120px;
  max-width: 760px;

  transform: translateY(30px);
}

.hero__eyebrow {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  font-weight: 700;
  letter-spacing: 3px;
  color: #FFFFFF;
  margin: 0 0 18px;
  text-transform: uppercase;
}

.hero__title {
  font-family: 'Montserrat', sans-serif;
  font-size: clamp(34px, 4.8vw, 56px);
  font-weight: 800;
  line-height: 1.15;
  color: #FFFFFF;
  margin: 0 0 24px;
  letter-spacing: -0.3px;

  white-space: nowrap; 
}

.hero__title--gold {
  color: #C8A135;
}

.hero__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 13.5px;
  font-weight: 400;
  line-height: 1.75;
  color: rgba(255, 255, 255, 0.88);
  margin: 0 0 36px;
  max-width: 520px;
}

.hero__cta {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 14px 28px;
  background: #1B7A6E;
  color: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 600;
  text-decoration: none;
  border-radius: 8px;
  transition: background 0.2s ease, transform 0.2s ease;
}

.hero__cta:hover {
  background: #145F56;
  transform: translateY(-2px);
}

.hero__cta-icon {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

/* Tablet */
@media (max-width: 768px) {
  .hero__content {
    padding: 0 28px;
    padding-left: 40px; /* tetap agak ke kanan tapi responsif */
  }

  .hero__title {
    font-size: clamp(30px, 6vw, 42px);
  }
}

/* Mobile */
@media (max-width: 480px) {
  .hero__content {
    padding: 0 20px;
  }
}

/* ════════════════════════════════
   ABOUT
════════════════════════════════ */
.about {
  padding: 80px 0;
  background: #FFFFFF;
}
.about__top {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
  margin-bottom: 16px;
}

.about__img-main {
  width: 100%;
  border-radius: 4px;
  overflow: hidden;
  aspect-ratio: 4 / 3;
}

.about__img-main img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.about__text {
  padding-top: 8px;
}

.about__label {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 28px;
}

.about__label-line {
  display: inline-block;
  width: 48px;
  height: 3px;
  background: #C8A135;
  border-radius: 2px;
  flex-shrink: 0;
}

.about__title {
  font-family: 'Montserrat', sans-serif;
  font-size: 32px;
  font-weight: 700;
  color: #C8A135;
  margin: 0;
}

.about__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 15px;
  font-weight: 400;
  line-height: 2;
  color: #2C2C2C;
  text-align: justify;
  margin: 0 0 40px;
}

.about__btn {
  display: inline-block;
  padding: 14px 40px;
  background: #1B7A6E;
  color: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  font-size: 14px;
  font-weight: 600;
  text-decoration: none;
  border-radius: 6px;
  transition: background 0.2s ease, transform 0.2s ease;
}

.about__btn:hover {
  background: #145F56;
  transform: translateY(-2px);
}

.about__bottom {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  gap: 16px;
}

.about__bottom img {
  width: 100%;
  height: 260px;
  object-fit: cover;
  border-radius: 4px;
  display: block;
}

/* Responsive */
@media (max-width: 860px) {
  .about__top {
    grid-template-columns: 1fr;
    gap: 32px;
  }
  .about__bottom {
    grid-template-columns: 1fr;
  }
  .about__bottom img {
    height: 240px;
  }
}

/* ════════════════════════════════
   SERVICES
════════════════════════════════ */
/* ── SERVICES ── */
.services {
  padding: 72px 0;
  background: #FFFFFF;
}

.services__label {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 48px;
}

.services__label-line {
  display: inline-block;
  width: 48px;
  height: 3px;
  background: #C8A135;
  border-radius: 2px;
  flex-shrink: 0;
}

.services__title {
  font-family: 'Montserrat', sans-serif;
  font-size: 32px;
  font-weight: 700;
  color: #C8A135;
  margin: 0;
}

.services__grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
}

.service-item {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.service-item__icon {
  margin-bottom: 16px;
}

.service-item__title {
  font-family: 'Montserrat', sans-serif;
  font-size: 17px;
  font-weight: 700;
  color: #2C2C2C;
  margin: 0 0 12px;
}

.service-item__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 13.5px;
  font-weight: 400;
  line-height: 1.8;
  color: #666666;
  margin: 0;
}

@media (max-width: 768px) {
  .services__grid {
    grid-template-columns: 1fr;
    gap: 36px;
  }
}

/* ════════════════════════════════
   PORTFOLIO
════════════════════════════════ */
/* ── PORTFOLIO ── */
.portfolio {
  padding: 80px 0;
  background: #F5F5EE;
  overflow: hidden;
}

.portfolio__inner {
  display: grid;
  grid-template-columns: 480px 1fr;
  align-items: center;
  min-height: 580px;
}

/* Teks kiri */
.portfolio__text {
  padding: 0 48px 0 80px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.portfolio__label {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
}

.portfolio__label-line {
  display: inline-block;
  width: 40px;
  height: 3px;
  background: #C8A135;
  border-radius: 2px;
  flex-shrink: 0;
}

.portfolio__title-gold {
  font-family: 'Montserrat', sans-serif;
  font-size: 28px;
  font-weight: 700;
  color: #C8A135;
  margin: 0;
}

.portfolio__heading {
  font-family: 'Montserrat', sans-serif;
  font-size: 28px;
  font-weight: 700;
  color: #2C2C2C;
  line-height: 1.3;
  margin: 0 0 24px;
}

.portfolio__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 13.5px;
  line-height: 1.85;
  color: #666666;
  margin: 0 0 40px;
}

.portfolio__btn {
  display: inline-block;
  padding: 13px 24px;
  background: #1B7A6E;
  color: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  font-weight: 600;
  text-decoration: none;
  border-radius: 6px;
  transition: background 0.2s ease, transform 0.2s ease;
}

.portfolio__btn:hover {
  background: #145F56;
  transform: translateY(-2px);
}

/* Slider kanan */
.portfolio__slider-wrap {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  overflow: hidden;
}

.portfolio__slider {
  display: flex;
  gap: 14px;
  padding: 16px 0;
  width: 100%;
  overflow: hidden;
}

.portfolio__slide {
  position: relative;
  flex-shrink: 0;
  width: 340px;
  border-radius: 12px;
  overflow: hidden;
  transition: width 0.4s ease;
}

.portfolio__slide--active {
  width: 380px;
}

.portfolio__slide-img {
  width: 100%;
  height: 500px;
  object-fit: cover;
  display: block;
  border-radius: 12px;
}

/* Info card di slide aktif */
.portfolio__slide-info {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.72);
  backdrop-filter: blur(6px);
  padding: 18px 20px 20px;
  border-radius: 0 0 12px 12px;
}

.portfolio__slide-num {
  display: flex;
  align-items: center;
  gap: 10px;
  font-family: 'Montserrat', sans-serif;
  font-size: 11.5px;
  font-weight: 500;
  color: #888;
  margin-bottom: 8px;
}

.portfolio__slide-line {
  display: inline-block;
  width: 24px;
  height: 2px;
  background: #888;
  border-radius: 2px;
}

.portfolio__slide-title {
  font-family: 'Montserrat', sans-serif;
  font-size: 19px;
  font-weight: 700;
  color: #2C2C2C;
  margin: 0 0 6px;
}

.portfolio__slide-desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 12px;
  color: #666;
  margin: 0;
  line-height: 1.6;
}

/* Arrow buttons */
.portfolio__arrow {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 42px;
  height: 42px;
  border-radius: 50%;
  background: #FFFFFF;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 12px rgba(0,0,0,0.15);
  color: #2C2C2C;
  transition: background 0.2s ease;
  z-index: 10;
}

.portfolio__arrow:hover {
  background: #f0f0f0;
}

.portfolio__arrow--left {
  display: none;
}

.portfolio__arrow--right {
  right: 16px;
}

/* Dots */
.portfolio__dots {
  display: flex;
  gap: 8px;
  margin-top: 16px;
  padding-left: 4px;
}

.portfolio__dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #D0D0D0;
  border: none;
  cursor: pointer;
  padding: 0;
  transition: background 0.2s ease, transform 0.2s ease;
}

.portfolio__dot--active {
  background: #1B7A6E;
  transform: scale(1.2);
}

@media (max-width: 900px) {
  .portfolio__inner {
    grid-template-columns: 1fr;
  }
  .portfolio__text {
    padding: 0 24px 40px;
  }
  .portfolio__slide {
    width: 260px;
  }
  .portfolio__slide--active {
    width: 290px;
  }
  .portfolio__slide-img {
    height: 360px;
  }
}

/* ════════════════════════════════
   ESTIMATION
════════════════════════════════ */
.estimasi {
  padding: 80px 0;
  background: #FFFFFF;
}

.estimasi__grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 80px;
  align-items: center;
}

.estimasi__images {
  position: relative;
  height: 480px;
}

.estimasi__img {
  position: absolute;
  object-fit: cover;
  border-radius: 8px;
  display: block;
}

.estimasi__img--top {
  width: 75%;
  height: 55%;
  top: 0;
  right: 0;
  z-index: 2;
}

.estimasi__img--bottom {
  width: 72%;
  height: 58%;
  bottom: 0;
  left: 0;
  z-index: 1;
}

.estimasi__label {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 20px;
}

.estimasi__label-line {
  display: inline-block;
  width: 40px;
  height: 3px;
  background: #C8A135;
  border-radius: 2px;
  flex-shrink: 0;
}

.estimasi__title-gold {
  font-family: 'Montserrat', sans-serif;
  font-size: 26px;
  font-weight: 700;
  color: #C8A135;
  margin: 0;
}

.estimasi__heading {
  font-family: 'Montserrat', sans-serif;
  font-size: 26px;
  font-weight: 700;
  color: #2C2C2C;
  line-height: 1.3;
  margin: 0 0 20px;
}

.estimasi__desc {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  line-height: 1.85;
  color: #666666;
  margin: 0 0 36px;
}

.estimasi__btn {
  display: inline-block;
  padding: 13px 24px;
  background: #1B7A6E;
  color: #FFFFFF;
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  font-weight: 600;
  text-decoration: none;
  border-radius: 6px;
  transition: background 0.2s ease, transform 0.2s ease;
}

.estimasi__btn:hover {
  background: #145F56;
  transform: translateY(-2px);
}

/* ════════════════════════════════
   TESTIMONIALS
════════════════════════════════ */
.testimoni {
  padding: 80px 0;
  background: #FFFFFF;
}

.testimoni__inner {
  display: grid;
  grid-template-columns: 340px 1fr;
  gap: 48px;
  align-items: start;
}

/* Kiri */
.testimoni__left {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.testimoni__img {
  width: 100%;
  height: 360px;
  object-fit: cover;
  border-radius: 8px;
  display: block;
}

.testimoni__arrows {
  display: flex;
  gap: 12px;
}

.testimoni__arrow {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 1.5px solid #D0D0D0;
  background: #FFFFFF;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #2C2C2C;
  transition: border-color 0.2s ease, background 0.2s ease;
}

.testimoni__arrow:hover {
  border-color: #1B7A6E;
  color: #1B7A6E;
}

/* Kanan */
.testimoni__right {
  padding-top: 8px;
}

.testimoni__label {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 12px;
}

.testimoni__label-line {
  display: inline-block;
  width: 40px;
  height: 3px;
  background: #C8A135;
  border-radius: 2px;
  flex-shrink: 0;
}

.testimoni__title-gold {
  font-family: 'Montserrat', sans-serif;
  font-size: 26px;
  font-weight: 700;
  color: #C8A135;
  margin: 0;
}

.testimoni__heading {
  font-family: 'Montserrat', sans-serif;
  font-size: 26px;
  font-weight: 700;
  color: #2C2C2C;
  margin: 0 0 28px;
}

.testimoni__cards {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.testimoni__card {
  background: #1B7A6E;
  border-radius: 10px;
  padding: 28px 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.testimoni__quote {
  font-size: 48px;
  font-weight: 700;
  color: rgba(255, 255, 255, 0.5);
  line-height: 1;
  font-family: Georgia, serif;
}

.testimoni__text {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  line-height: 1.8;
  color: #FFFFFF;
  margin: 0;
  flex: 1;
}

.testimoni__author {
  display: flex;
  align-items: center;
  gap: 12px;
}

.testimoni__avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.testimoni__name {
  font-family: 'Montserrat', sans-serif;
  font-size: 13px;
  font-weight: 700;
  color: #FFFFFF;
}

.testimoni__role {
  font-family: 'Montserrat', sans-serif;
  font-size: 11px;
  color: rgba(255, 255, 255, 0.65);
  margin-top: 2px;
}

@media (max-width: 860px) {
  .estimasi__grid,
  .testimoni__inner {
    grid-template-columns: 1fr;
  }
  .estimasi__images {
    height: 320px;
  }
  .testimoni__cards {
    grid-template-columns: 1fr;
  }
}

/* ════════════════════════════════
   RESPONSIVE
════════════════════════════════ */
@media (max-width: 1024px) {
  .about__grid {
    gap: 48px;
  }

  .estimation__content {
    gap: 48px;
  }
}

@media (max-width: 860px) {
  .container { padding: 0 20px; }
  .section { padding: 72px 0; }

  .about__grid,
  .estimation__content {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .services__grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .portfolio__grid {
    grid-template-columns: 1fr;
  }

  .portfolio-card--featured {
    grid-row: span 1;
  }

  .portfolio-card--featured .portfolio-card__img-wrap {
    height: 280px;
  }

  .testimonials__grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .about__img-accent {
    right: 12px;
  }

  .estimation__card {
    max-width: 280px;
  }
}

@media (max-width: 480px) {
  .hero__title--outline {
    -webkit-text-stroke-width: 1.5px;
  }

  .about__img-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .about__img-grid img:last-child {
    display: none;
  }
}
</style>