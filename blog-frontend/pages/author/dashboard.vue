<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AuthorSidebar />
      <!-- Main Content -->
      <main class="p-4 col-md-12 col-md-9 col-lg-10">
        <h1 class="mb-5">Workspace Author</h1>
        <hr>

        <div class="row">
          <div class="col-sm-6 mb-3 mb-sm-0">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title">Draft Artikel Anda</h5>
                <div class="d-flex justify-content-between mb-2">
                  <div class="card" style="width: 18rem;">
                    <img src="/assets/img/windah senyum Roblox.jpg" class="card-img-top" alt="...">
                    <div class="card-body">
                      <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    </div>
                  </div>
                  <div class="card" style="width: 18rem; height: 17rem;">
                    <div class="card-body text-center mt-5">
                      <NuxtLink to="/author/articles/add-articles" class="btn btn-outline-primary mt-5"><i class="bi bi-plus"></i> Tambah</NuxtLink>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title">Artikel Yang Berhasil Di Publikasi</h5>
                <div class="d-flex justify-content-between">
                  <div class="card" style="width: 18rem;">
                    <img src="/assets/img/windah senyum Roblox.jpg" class="card-img-top" alt="...">
                    <div class="card-body">
                      <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    </div>
                  </div>
                  <div class="card" style="width: 18rem;">
                    <img src="/assets/img/windah senyum Roblox.jpg" class="card-img-top" alt="...">
                    <div class="card-body">
                      <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'auth',
  requiresAuthor: true,
});
</script>