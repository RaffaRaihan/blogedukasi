<template>
    <Navbar />
    <!-- Content -->
    <div class="container mt-4">
      <div class="search-box mb-4">
        <form action="#" method="GET" class="d-flex">
          <input class="form-control me-2" type="search" placeholder="Cari artikel..." name="query" aria-label="Search">
          <button class="btn btn-success" type="submit">Cari</button>
        </form>
      </div>

      <h4>Hasil Pencarian untuk: "HTML"</h4>
      <p class="text-muted">Ditemukan 3 hasil</p>

      <div class="row">
        <!-- Article 1 -->
        <div class="col-12 mb-4">
          <div class="card">
            <div class="card-body">
              <h5 class="article-title">Apa Itu HTML?</h5>
              <p class="article-meta">By Ridho Satriawan - 22 Agustus 2022</p>
              <p>HTML adalah singkatan dari HyperText Markup Language, yang digunakan untuk membuat struktur halaman website...</p>
              <a href="#" class="btn btn-success btn-sm">Baca Selengkapnya</a>
            </div>
          </div>
        </div>

        <!-- Article 2 -->
        <div class="col-12 mb-4">
          <div class="card">
            <div class="card-body">
              <h5 class="article-title">Panduan Dasar HTML untuk Pemula</h5>
              <p class="article-meta">By Jane Doe - 15 Juli 2022</p>
              <p>Panduan ini dirancang untuk membantu pemula memahami dasar-dasar HTML dan cara penggunaannya...</p>
              <a href="#" class="btn btn-success btn-sm">Baca Selengkapnya</a>
            </div>
          </div>
        </div>

        <!-- Article 3 -->
        <div class="col-12 mb-4">
          <div class="card">
            <div class="card-body">
              <h5 class="article-title">Sejarah dan Perkembangan HTML</h5>
              <p class="article-meta">By John Smith - 10 Juni 2022</p>
              <p>HTML telah berkembang sejak pertama kali diperkenalkan oleh Tim Berners-Lee pada tahun 1989...</p>
              <a href="#" class="btn btn-success btn-sm">Baca Selengkapnya</a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Footer />
</template>

<style scoped>
.search-box {
  background-color: #f8f9fa;
  padding: 15px;
  border-radius: 5px;
}
.article-title {
  font-size: 1.25rem;
  font-weight: bold;
  margin-bottom: 0.5rem;
}
.article-meta {
  font-size: 0.9rem;
  color: #6c757d;
}
</style>