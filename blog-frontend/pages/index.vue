<template>
    <Navbar />
    <!-- Content -->
    <div class="container mt-4">
      <div class="row">
        <!-- Main Content -->
        <div class="col-lg-8">
          <div class="card mb-4">
            <div class="card-body">
                <div id="carousel" class="carousel slide">
                  <div class="carousel-inner">
                    <div class="carousel-item active">
                      <img src="../assets/img/wp6118801.jpg" class="d-block w-100" alt="...">
                    </div>
                    <div class="carousel-item">
                      <img src="../assets/img/OIP.jpg" class="d-block w-100" alt="...">
                    </div>
                    <div class="carousel-item">
                      <img src="../assets/img/windah senyum Roblox.jpg" class="d-block w-100" alt="...">
                    </div>
                  </div>
                    <h6 class="text-muted">In Label</h6>
                    <h5 class="card-title">Ini Adalah Judul Artikel</h5>
                    <p class="card-text">By Ridho Satriawan</p>
                  <button class="carousel-control-prev" type="button" data-bs-target="#Carousel" data-bs-slide="prev">
                    <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                    <span class="visually-hidden">Previous</span>
                  </button>
                  <button class="carousel-control-next" type="button" data-bs-target="#Carousel" data-bs-slide="next">
                    <span class="carousel-control-next-icon" aria-hidden="true"></span>
                    <span class="visually-hidden">Next</span>
                  </button>
                </div>
            </div>
          </div>

          <h5>Postingan Terbaru</h5>
          <div class="row">
            <div class="col-md-6 mb-3">
              <div class="card">
                <div class="card-body">
                  <h6 class="text-muted">In Label</h6>
                  <h5 class="card-title">Ini Adalah Judul Artikel</h5>
                  <p class="card-text">Apa itu HTML 7? Buat kamu yang pernah...</p>
                  <p class="text-muted">22 Agustus 2022</p>
                  <NuxtLink to="/detail" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                </div>
              </div>
            </div>
            <div class="col-md-6 mb-3">
              <div class="card">
                <div class="card-body">
                  <h6 class="text-muted">In Label</h6>
                  <h5 class="card-title">Ini Adalah Judul Artikel</h5>
                  <p class="card-text">Apa itu HTML 7? Buat kamu yang pernah...</p>
                  <p class="text-muted">22 Agustus 2022</p>
                  <NuxtLink to="/detail" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                </div>
              </div>
            </div>
            <div class="col-md-6 mb-3">
              <div class="card">
                <div class="card-body">
                  <h6 class="text-muted">In Label</h6>
                  <h5 class="card-title">Ini Adalah Judul Artikel</h5>
                  <p class="card-text">Apa itu HTML 7? Buat kamu yang pernah...</p>
                  <p class="text-muted">22 Agustus 2022</p>
                </div>
              </div>
            </div>
            <!-- Tambah artikel lainnya sesuai kebutuhan -->
          </div>
        </div>

        <!-- Sidebar -->
        <Sidebar />
      </div>
    </div>
    <Footer />
</template>