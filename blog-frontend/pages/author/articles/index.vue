<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AuthorSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Author - Kelola Artikel</h1>
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
            <option disabled>Pilih Kategori</option>
            <option value="">Semua Kategori</option>
            <option v-for="category in categories" :key="category.id" :value="category.ID">
              {{ category.name }}
            </option>
          </select>
        </div>
        <button class="btn btn-outline-primary mb-4" @click="navigateToAddArticle"><i class="bi bi-plus"></i>Tambah</button>

        <!-- Loading Indicator -->
        <Loading v-if="loadingArticles" />

        <!-- Daftar Artikel -->
        <div class="row" v-else>
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
                <p class="card-text" v-html="getTruncatedContent(article.content)"></p>
                <button class="btn btn-outline-warning" @click="navigateToEditArticle(article.ID)"><i class="bi bi-pencil"></i></button>
                <button @click="removeArticles(article.ID)"  class="btn btn-outline-danger ms-2"><i class="bi bi-trash"></i></button>
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
import { useRouter } from 'vue-router';
import axios from 'axios';
import DOMPurify from "dompurify";

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const router = useRouter();
const articles = ref([]);
const categories = ref([]);
const selectedCategory = ref(null);
const searchQuery = ref('');
const loadingArticles = ref(true)

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  return token;
};

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
  } finally {
    loadingArticles.value = false;
  }
};

const fetchCategories = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category');
    categories.value = response.data;
  } catch (error) {
    console.error('Error fetching categories:', error);
  }
};

const removeArticles = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus artikel ini?')) {
    return;
  }

  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.delete(`http://localhost:8080/admin/articles/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    articles.value = articles.value.filter(articles => articles.ID !== id);
    console.log(`artikel dengan ID ${id} berhasil dihapus.`);
  } catch (error) {
    console.error('Error removing user:', error);
  }
};

const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    const matchesCategory = selectedCategory.value ? article.category.ID === selectedCategory.value : true;
    const matchesSearch = article.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesCategory && matchesSearch;
  });
});

const filterArticles = () => {
  fetchArticles(selectedCategory.value, searchQuery.value);
};

const navigateToAddArticle = () => {
  router.push('/admin/articles/add-articles');
};

const navigateToEditArticle = (id) => {
  router.push(`/admin/articles/edit/${id}`);
};

// Truncate konten artikel
function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 50).join(" ") + "...";
  return DOMPurify.sanitize(truncated); // Aman untuk dirender
}

onMounted(() => {
  fetchArticles();
  fetchCategories();
});
</script>
