<template>
    <Navbar />
    <!-- Content -->
    <div class="container mt-4">
      <div class="row">
        <!-- Main Content -->
        <div class="col-lg-8">
          <div class="card mb-4">
            <div id="carouselExampleSlidesOnly" class="carousel slide" data-bs-ride="carousel">
              <div class="carousel-inner">
                <div class="carousel-item active">
                  <img src="../assets/img/OIP.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/wp6118801.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/windah senyum Roblox.jpg" class="d-block w-100" alt="...">
                </div>
              </div>
            </div>
          </div>

          <h5>Postingan Terbaru</h5>
          <div class="row">
            <div v-for="articles in articles" :key="articles.id" class="col-md-6 mb-3">
              <div class="card">
                <div class="card-body">
                  <img src="../assets/img/windah senyum Roblox.jpg" class="img-fluid" alt="...">
                  <h6 class="text-muted">{{ articles.label }}</h6>
                  <h5 class="card-title">{{ articles.title }}</h5>
                  <p class="card-text">{{ articles.content }}</p>
                  <p class="text-muted">{{ articles.CreatedAt }}</p>
                  <NuxtLink to="/login" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
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

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const articles = ref([]);

const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles'); // Ganti dengan URL backend Anda
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Panggil fetchArticles saat komponen dimuat
onMounted(() => {
  fetchArticles();
});
</script>
