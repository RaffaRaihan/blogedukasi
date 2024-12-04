<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Artikel</h1>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input 
            type="text" 
            class="form-control me-2" 
            placeholder="Search by Title" 
            v-model="searchQuery" 
            @input="filterArticles"
          />
          <select 
            class="form-select me-2" 
            v-model="selectedCategory" 
            @change="filterArticlesByCategory"
          >
            <option value="">Semua Kategori</option>
            <option v-for="category in category" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>

        <!-- Tombol untuk menambah artikel baru -->
        <button
          class="btn btn-primary mb-4"
          data-bs-toggle="modal"
          data-bs-target="#articleModal"
          @click="openModal('add')"
        >
          Tambah Artikel Baru
        </button>

        <!-- Daftar artikel dalam bentuk card -->
        <div class="row">
          <div
            v-for="article in articles"
            :key="article.id"
            class="col-md-4 mb-4"
          >
            <div class="card">
              <img
                :src="`http://localhost:8080/uploads/${article.file_name}`"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h6 class="card-label">{{ article.label }}</h6>
                <h5 class="card-title">{{ article.title }}</h5>
                <p class="card-text">{{ article.content }}</p>
                <button
                  class="btn btn-warning"
                  data-bs-toggle="modal"
                  data-bs-target="#articleModal"
                  @click="openModal('edit', article)"
                >
                  Edit
                </button>
                <button class="btn btn-danger ms-2">Hapus</button>
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

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const articles = ref([]);
const category = ref([]);
const selectedCategory = ref(null); // Menyimpan kategori yang dipilih
const searchQuery = ref(''); // Menyimpan query pencarian

// Mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fungsi untuk mengambil artikel
const fetchArticles = async (categoryId = null, search = '') => {
  try {
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.get('http://localhost:8080/articles', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params: {
        categoryId, // Kirimkan kategori sebagai parameter jika ada
        search, // Kirimkan query pencarian jika ada
      },
    });
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Fungsi untuk mengambil kategori
const fetchCategory = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category');
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  }
};

// Fungsi untuk filter artikel berdasarkan kategori
const filterArticlesByCategory = () => {
  fetchArticles(selectedCategory.value, searchQuery.value); // Panggil fetchArticles dengan kategori dan pencarian
};

// Fungsi untuk filter artikel berdasarkan pencarian judul
const filterArticles = () => {
  fetchArticles(selectedCategory.value, searchQuery.value); // Panggil fetchArticles dengan query pencarian
};

onMounted(() => {
  fetchArticles();
  fetchCategory();
});
</script>
