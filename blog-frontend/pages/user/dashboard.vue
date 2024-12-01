<template>
    <LoginNavbar />
    <div class="container mt-4">
        <div class="row">
            <div class="col-lg-8">
                <h5>Postingan Terbaru</h5>
                <div class="row">
                  <div v-for="articles in articles" :key="articles.id" class="col-md-6 mb-3">
                    <div class="card">
                      <img src="/assets/img/Inumaki.jpg" class="card-img-top" alt="...">
                      <div class="card-body">
                        <h6 class="text-muted">{{ articles.label }}</h6>
                        <h5 class="card-title">{{ articles.title }}</h5>
                        <p class="card-text">{{ articles.content }}</p>
                        <p class="text-muted">{{ articles.CreatedAt }}</p>
                        <NuxtLink to="/login" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                      </div>
                    </div>
                  </div>
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
