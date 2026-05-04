    <template>
  <div class="portfolio-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">Portofolio</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">Portofolio</span>
        </div>
      </div>
    </div>

    <!-- ── SECTION 1: Hero Banner ── -->
    <section class="hero-banner">
      <div class="hero-banner__img hero-banner__img--left">
        <img src="https://images.unsplash.com/photo-1616594039964-ae9021a400a0?w=600&q=80" alt="Interior bedroom" />
      </div>
      <div class="hero-banner__text">
        <p>
          Lihat berbagai hasil karya
          <strong class="hero-banner__brand">SISI Interior</strong>
          yang dirancang untuk menghadirkan ruang yang estetik, nyaman, dan fungsional.
        </p>
      </div>
      <div class="hero-banner__img hero-banner__img--right">
        <img src="https://images.unsplash.com/photo-1555041469-a586c61ea9bc?w=600&q=80" alt="Interior living room" />
      </div>
    </section>

    <!-- ── SECTION 2: Featured Projects ── -->
    <section class="projects section">
      <div class="container">

        <!-- Project 1: Kitchen Set — Image Left -->
        <div class="project project--img-left">
          <div class="project__image">
            <img src="https://images.unsplash.com/photo-1556909114-f6e7ad7d3136?w=700&q=80" alt="Kitchen Set Minimalis Modern" />
          </div>
          <div class="project__info">
            <span class="project__category">Desain Interior</span>
            <h2 class="project__title">KITCHEN SET MINIMALIS MODERN</h2>
            <p class="project__desc">
              Proyek kitchen set dengan konsep minimalis modern yang dirancang untuk
              memaksimalkan fungsi ruang sekaligus menghadirkan tampilan yang rapi dan elegan.
            </p>
            <a href="#" class="btn-contact">Hubungi Kami</a>
          </div>
        </div>

        <!-- Project 2: Bedroom Set — Image Right -->
        <div class="project project--img-right">
          <div class="project__info">
            <span class="project__category">Desain Interior &amp; Custom Furniture</span>
            <h2 class="project__title">BEDROOM SET CUSTOM ELEGAN</h2>
            <p class="project__desc">
              Penataan kamar tidur dengan furniture custom yang disesuaikan dengan ukuran ruang,
              kebutuhan penyimpanan, dan kenyamanan pengguna.
            </p>
            <a href="#" class="btn-contact">Hubungi Kami</a>
          </div>
          <div class="project__image">
            <img src="https://images.unsplash.com/photo-1631049307264-da0ec9d70304?w=700&q=80" alt="Bedroom Set Custom Elegan" />
          </div>
        </div>

        <!-- Project 3: Ruang Tamu — Image Left -->
        <div class="project project--img-left">
          <div class="project__image">
            <img src="https://images.unsplash.com/photo-1600210492493-0946911123ea?w=700&q=80" alt="Interior Ruang Tamu" />
          </div>
          <div class="project__info">
            <span class="project__category">Kontraktor &amp; Desain Interior</span>
            <h2 class="project__title">INTERIOR RUANG TAMU NYAMAN</h2>
            <p class="project__desc">
              Proyek renovasi ruang tamu untuk menciptakan suasana yang lebih nyaman, estetik,
              dan sesuai dengan kebutuhan aktivitas keluarga.
            </p>
            <a href="#" class="btn-contact">Hubungi Kami</a>
          </div>
        </div>

      </div>
    </section>

    <!-- ── SECTION 3: Hasil Karya Lainnya ── -->
    <section class="gallery section">
      <div class="container">
        <p class="gallery__label">PORTFOLIO</p>
        <h2 class="gallery__title">HASIL KARYA LAINNYA</h2>

        <div class="gallery__grid">

          <div 
            class="gallery__item" 
            v-for="(item, index) in portfolios" 
            :key="item_id"
          >
          <img 
            :src="`http://localhost:8081/uploads/${item.gambar}`" 
            :alt="item.judul_proyek" 
          />

            <div class="gallery__overlay">
              <h3>{{ item.judul_proyek }}</h3>
              <p>{{ item.deskripsi}}</p>
            </div>
          </div>

        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const portfolios = ref([])

const getPortfolios = async () => {
  try {
    const res = await fetch("http://localhost:8081/api/portfolios")
    const data = await res.json()

    portfolios.value = data.data
  } catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  getPortfolios()
})
</script>

]

<style scoped>
.portfolio-page {
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

/* ── SECTION 1: Hero Banner ── */
.hero-banner {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  gap: 0;
  max-height: 320px;
  overflow: hidden;
}

.hero-banner__img {
  height: 280px;
  overflow: hidden;
}

.hero-banner__img img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.hero-banner__text {
  padding: 40px 48px;
  text-align: center;
  font-size: 14px;
  line-height: 1.8;
  color: #2C2C2C;
  min-width: 260px;
  max-width: 340px;
}

.hero-banner__brand {
  color: #1B7A6E;
  font-weight: 700;
}

/* ── SECTION 2: Featured Projects ── */
.projects {
  background: #FFFFFF;
}

.project {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
  margin-bottom: 80px;
}

.project:last-child {
  margin-bottom: 0;
}

.project--img-right .project__info {
  order: 1;
}

.project--img-right .project__image {
  order: 2;
}

.project__image {
  border-radius: 6px;
  overflow: hidden;
  aspect-ratio: 4/3;
}

.project__image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.project__category {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
  color: #888;
  display: block;
  margin-bottom: 12px;
}

.project__title {
  font-size: 22px;
  font-weight: 800;
  color: #1B7A6E;
  line-height: 1.3;
  margin: 0 0 16px;
  text-transform: uppercase;
}

.project__desc {
  font-size: 13px;
  line-height: 1.9;
  color: #555;
  text-align: justify;
  margin: 0 0 24px;
}

/* ── Button ── */
.btn-contact {
  display: inline-block;
  background: #C8A135;
  color: #FFFFFF;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.5px;
  padding: 10px 22px;
  border-radius: 4px;
  text-decoration: none;
  transition: background 0.2s ease, transform 0.2s ease;
}

.btn-contact:hover {
  background: #b08e2c;
  transform: translateY(-1px);
}

.btn-contact--sm {
  font-size: 11px;
  padding: 8px 16px;
}

/* ── SECTION 3: Gallery ── */
/* ── SECTION 3: Gallery (NEW CLEAN VERSION) ── */
.gallery {
  background: #FAFAFA;
}

.gallery__label {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 3px;
  color: #888;
  margin: 0 0 8px;
}

.gallery__title {
  font-size: 28px;
  font-weight: 800;
  color: #2C2C2C;
  margin: 0 0 36px;
  text-transform: uppercase;
}

/* GRID SIMPLE (seperti gambar 1) */
.gallery__grid {
  display: grid !important;
  grid-template-columns: repeat(3, 1fr) !important;
  gap: 20px;
}

/* CARD */
.gallery__item {
  position: relative;
  width: 100%;
  height: 280px;
  overflow: hidden;
  border-radius: 14px;
  cursor: pointer;
}

/* IMAGE */
.gallery__item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0,5s ease;
}

/* OVERLAY (DEFAULT HIDDEN) */
.gallery__overlay {
  position: absolute;
  inset: 0;

  background: rgba(0, 0, 0, 0.6);

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  text-align: center;
  padding: 20px;

  opacity: 0;
  transform: translateY(20px);
  transition: all 0.4s ease;
}

/* TEXT */
.gallery__overlay h3 {
  font-size: 18px;
  color: #fff;
  margin-bottom: 10px;
}

.gallery__overlay p {
  font-size: 14px;
  color: #eee;
}

/* HOVER EFFECT 🔥 */
.gallery__item:hover img {
  transform: scale(1.1);
}

.gallery__item:hover .gallery__overlay {
  opacity: 1;
  transform: translateY(0);
}

/* Overlay item */

.gallery__overlay-title {
  font-size: 13px;
  font-weight: 700;
  color: #FFFFFF;
  text-align: center;
  margin: 0;
}

/* ── Responsive ── */
@media (max-width: 992px) {
  .gallery__grid {
    grid-template-columns: repeat(2,1fr);
  }

  .hero-banner__img {
    height: 200px;
  }

  .hero-banner__text {
    max-width: 100%;
    padding: 32px 40px;
  }

  .project {
    grid-template-columns: 1fr;
    gap: 32px;
  }

  .project--img-right .project__info {
    order: 2;
  }
  .project--img-right .project__image {
    order: 1;
  }

  .gallery__grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .gallery__item--wide,
  .gallery__item--third,
  .gallery__item--quarter {
    grid-column: span 1;
  }
}

@media (max-width: 576px) {
  .gallery__grid {
    grid-template-columns: 1fr;
  }


  .gallery__item--wide,
  .gallery__item--third,
  .gallery__item--quarter {
    grid-column: span 1;
  }

  .hero-banner__text {
    padding: 24px 20px;
  }
}
</style>