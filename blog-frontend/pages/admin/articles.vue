<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
        <div class="col-md-9 col-lg-10 p-4">
          <h1 class="mb-4">Halaman Admin - Kelola Artikel</h1>
        
          <!-- Tombol untuk menambah artikel baru -->
          <NuxtLink to="/" class="btn btn-primary mb-4">Tambah Artikel Baru</NuxtLink>
        
          <!-- Daftar artikel dalam bentuk card -->
          <div class="row">
            <!-- Loop untuk menampilkan artikel -->
            <div v-for="article in articles" :key="article.id" class="col-md-4 mb-4">
              <div class="card">
                <img src="/assets/img/windah senyum Roblox.jpg" class="card-img-top" alt="Thumbnail Artikel" />
                <div class="card-body">
                  <h6 class="card-label">{{ article.label }}</h6>
                  <h5 class="card-title">{{ article.title }}</h5>
                  <p class="card-text">{{ article.content }}</p>
                  <NuxtLink class="btn btn-warning">Edit</NuxtLink>
                  <button class="btn btn-danger">Hapus</button>
                </div>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const articles = ref([]);
// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchArticles = async () => {
  try {
      // Ambil token dari cookies
      const token = getTokenFromCookies();
    const response = await axios.get('http://localhost:8080/articles', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }); // Ganti dengan URL backend Anda
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