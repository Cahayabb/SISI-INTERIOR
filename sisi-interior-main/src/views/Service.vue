<template>
  <div class="service-page">

    <!-- ── Page Header ── -->
    <div class="page-header">
      <div class="container">
        <h1 class="page-header__title">Service</h1>
        <div class="page-header__breadcrumb">
          <RouterLink to="/" class="breadcrumb__link">Home</RouterLink>
          <span class="breadcrumb__sep">›</span>
          <span class="breadcrumb__current">Service</span>
        </div>
      </div>
    </div>

    <!-- ── SECTION 1: Custom Furniture — Image Left ── -->
    <section class="service-item section">
      <div class="container service-item__grid">

        <div class="service-item__image">
          <img
            src="https://images.unsplash.com/photo-1524758631624-e2822e304c36?w=700&q=80"
            alt="Custom Furniture"
          />
        </div>

        <div class="service-item__text">
          <h2 class="service-item__title">Custom Furniture</h2>
          <p class="service-item__desc">
            SISI INTERIOR merupakan produsen furniture yang memproduksi dan menjual
            berbagai macam furniture yang dapat Anda pesan sesuai dengan
            keinginan Anda, baik secara model, ukuran, dan desain.
          </p>
          <p class="service-item__desc">
            Kami menyediakan berbagai macam kebutuhan furniture Anda mulai dari
            meja makan, kursi makan, wardrobe (lemari pakaian), rak buku, foyer (lobi),
            lemari pemisah ruang, nakas, meja rias, background TV, partisi, lemari
            pajangan, meja kopi, dan lain-lain.
          </p>
          <p class="service-item__desc">
            Anda juga dapat memesan produk kami dengan desain yang telah anda
            buat sendiri, yang nantinya akan kami pelajari untuk diproduksi, setiap
            produk yang kami tawarkan memiliki tingkatan harga yang wajar sesuai
            dengan kualitas produknya serta dikerjakan dengan rapi dan teliti.
          </p>
        </div>

      </div>
    </section>

    <!-- ── SECTION 2: Desain Interior — Image Right ── -->
    <section class="service-item section">
      <div class="container service-item__grid service-item__grid--reverse">

        <div class="service-item__text">
          <h2 class="service-item__title">Desain Interior</h2>
          <p class="service-item__desc">
            SISI INTERIOR memberikan penawaran desain untuk menciptakan ruangan
            seperti foyer (lobi), living room set, bedroom set, kitchen set, dan lain-lain
            agar menarik indah dipandang mata.
          </p>
          <p class="service-item__desc">
            Desain yang kami buat dikerjakan oleh desainer yang dapat menvisualisasi
            hasil kerja yang diinginkan dengan gambar 3D, sehingga pelanggan kami
            akan lebih puas dan meminimalisir adanya salah pengertian.
          </p>
          <p class="service-item__desc">
            Perhitungan harga jual kami juga dibuat transparan sehingga mudah
            dimengerti dan jelas.
          </p>
          <p class="service-item__desc">
            Pengerjaan interior dan furniture kami dibuat dari bahan-bahan pilihan
            yang baik seperti dari kayu mahoni, kayu jati muda, multiplek, MDF, dan jenis
            kayu lainnya sesuai dengan pesanan anda dan di finishing menggunakan
            melamin, HPL, ataupun cat duco.
          </p>
        </div>

        <div class="service-item__image">
          <img
            src="https://images.unsplash.com/photo-1586023492125-27b2c045efd7?w=700&q=80"
            alt="Desain Interior"
          />
        </div>

      </div>
    </section>

    <!-- ── SECTION 3: Kontraktor — Image Left ── -->
    <section class="service-item section">
      <div class="container service-item__grid">

        <div class="service-item__image">
          <img
            src="https://images.unsplash.com/photo-1570129477492-45c003edd2be?w=700&q=80"
            alt="Kontraktor"
          />
        </div>

        <div class="service-item__text">
          <h2 class="service-item__title">Kontraktor</h2>
          <p class="service-item__desc">
            SISI INTERIOR memberikan pelayanan jasa dalam membangun hunian
            maupun tempat kerja sesuai dengan kebutuhan Anda.
          </p>
          <p class="service-item__desc">
            Kami juga membantu dalam menciptakan desain eksterior maupun interior
            untuk keperluan bangunan baru maupun renovasi agar area tersebut
            menjadi lebih baik.
          </p>
          <p class="service-item__desc">
            Segala pekerjaan mulai dari lantai, dinding sampai atap beserta proses
            finishingnya akan kami buat sebaik mungkin karena kepuasan Anda adalah
            prioritas kami.
          </p>
        </div>

      </div>
    </section>

  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>

<style scoped>
.service-page {
  padding-top: 72px; /* tinggi navbar */
  font-family: 'Montserrat', sans-serif;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 40px;
}

.section {
  padding: 60px 0;
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

/* ── Service Item ── */
.service-item__grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 72px;
  align-items: center;
}

/* Reverse layout: text left, image right */
.service-item__grid--reverse {
  /* default order is already text-left image-right via template order */
}

.service-item__image {
  border-radius: 6px;
  overflow: hidden;
  aspect-ratio: 4/3;
}

.service-item__image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.service-item__title {
  font-family: 'Montserrat', sans-serif;
  font-size: 32px;
  font-weight: 700;
  color: #C8A135;
  margin: 0 0 24px;
  font-style: normal;
}

.service-item__desc {
  font-size: 13.5px;
  line-height: 1.9;
  color: #444;
  text-align: justify;
  margin: 0 0 16px;
}

.service-item__desc:last-child {
  margin-bottom: 0;
}

/* ── Responsive ── */
@media (max-width: 860px) {
  .service-item__grid,
  .service-item__grid--reverse {
    grid-template-columns: 1fr;
    gap: 36px;
  }

  /* On mobile, always show image first */
  .service-item__grid--reverse .service-item__text {
    order: 2;
  }

  .service-item__grid--reverse .service-item__image {
    order: 1;
  }
}

@media (max-width: 560px) {
  .container {
    padding: 0 20px;
  }

  .service-item__title {
    font-size: 26px;
  }
}
</style>