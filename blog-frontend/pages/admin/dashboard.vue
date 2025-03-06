<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <!-- Main Content -->
      <main class="p-4 col-md-12 col-md-9 col-lg-10">
        <h1 class="mb-3">Halaman Dashboard - Admin</h1>
        <div class="card mb-3">
          <div class="card-header text-center">Data User</div>
          <div class="card-body">
            <p class="mb-0">Admin</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar bg-info" style="width: 25%">25%</div>
            </div>
            <p class="mb-0">Author</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar bg-warning" style="width: 15%">15%</div>
            </div>
            <p class="mb-0">User</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar" style="width: 35%">35%</div>
            </div>
            <NuxtLink to="/admin/users" class="btn btn-outline-primary btn-md mt-2">Lihat</NuxtLink>
          </div>
        </div>
        <div class="card mb-3">
          <div class="card-header text-center">Data Artikel</div>
          <div class="card-body">
            <p class="mb-0">Artikel dibuat</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar" style="width: 25%">25%</div>
            </div>
            <p class="mb-0">Artikel yang sesuai</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar bg-success" style="width: 35%">35%</div>
            </div>
            <p class="mb-0">Artikel yang belum sesuai</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar bg-danger" style="width: 15%">15%</div>
            </div>
            <NuxtLink to="/admin/articles" class="btn btn-outline-primary btn-md mt-2">Lihat</NuxtLink>
          </div>
        </div>
        <div class="card mb-3">
          <div class="card-header text-center">Data Kategori</div>
          <div class="card-body">
            <p class="mb-0">Masukan Kategori</p>
            <div class="progress mb-2" role="progressbar" aria-label="Example with label" aria-valuenow="25" aria-valuemin="0" aria-valuemax="100" style="height: 20px;">
              <div class="progress-bar" style="width: 25%">25%</div>
            </div>
            <NuxtLink to="/admin/category" class="btn btn-outline-primary btn-md mt-2">Lihat</NuxtLink>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
main {
  transition: all 0.3s ease;
}
</style>
