<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <!-- Main Content -->
      <main class="p-4 col-md-12 col-md-9 col-lg-10">
        <h1 class="mb-3">Halaman Dashboard - Admin</h1>
        
        <div class="d-flex justify-content-between">
          <div class="card w-50">
            <div class="card-header">
              Pengunjung
            </div>
            <div class="card-body">
              <TrafficChart />
            </div>
          </div>
          <div class="card w-50 ms-2">
            <div class="card-header">
              Penonton Artikel
            </div>
            <div class="card-body">
              <TrafficChart />
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-between mt-2">
          <div class="card w-50">
            <div class="card-header">
              Artikel Baru
            </div>
            <div class="card-body">
              <Chart />
            </div>
          </div>
          <div class="card w-50 ms-2">
            <div class="card-header">
              Penonton Artikel Berdasarkan Kategori
            </div>
            <div class="card-body">
              <Chart />
            </div>
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

