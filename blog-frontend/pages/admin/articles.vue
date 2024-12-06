<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Artikel</h1>

        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <!-- Search by Title -->
          <input 
            type="text" 
            class="form-control me-2" 
            placeholder="Search by Title" 
            v-model="searchQuery"
            @input="filterArticles"
          />
          <!-- Kategori Filter -->
          <select 
            class="form-select me-2" 
            v-model="selectedCategory" 
            @change="filterArticles"
          >
            <option value="">Semua Kategori</option>
            <option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>

        <!-- Daftar Artikel -->
        <div class="row">
          <div
            v-for="article in filteredArticles"
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
                <h5 class="card-title">{{ article.title }}</h5>
                <p class="card-text">{{ article.content }}</p>
                <button class="btn btn-warning" @click="editArticle(article)">
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
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const articles = ref([]);
const categories = ref([]);
const selectedCategory = ref(null); // Kategori yang dipilih
const searchQuery = ref(''); // Pencarian judul artikel

// Mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
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
        categoryId,
        search,
      },
    });
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Fungsi untuk mengambil kategori
const fetchCategories = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category');
    categories.value = response.data;
  } catch (error) {
    console.error('Error fetching categories:', error);
  }
};

// Computed untuk memfilter artikel berdasarkan kategori dan pencarian judul
const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    const matchesCategory = selectedCategory.value ? article.category.id === selectedCategory.value : true;
    const matchesSearch = article.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesCategory && matchesSearch;
  });
});

// Fungsi untuk memfilter artikel saat kategori atau pencarian berubah
const filterArticles = () => {
  fetchArticles(selectedCategory.value, searchQuery.value);
};

onMounted(() => {
  fetchArticles(); // Ambil artikel saat komponen dimuat
  fetchCategories(); // Ambil kategori
});
</script>
