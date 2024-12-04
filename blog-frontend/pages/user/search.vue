<template>
  <LoginNavbar />
  <div class="container mt-4">
    <h4>Hasil Pencarian untuk: "{{ route.query.query || 'Semua Artikel' }}"</h4>
    <p class="text-muted">Ditemukan {{ meta.total_data || 0 }} hasil</p>

    <!-- Daftar Artikel -->
    <div class="row" v-if="articles.length">
      <div class="col-12 mb-4" v-for="article in articles" :key="article.id">
        <div class="card">
          <img :src="`http://localhost:8080/uploads/${article.file_name}`" class="card-img-top img-fluid" alt="..."/>
          <div class="card-body">
            <h5 class="article-title">{{ article.title }}</h5>
            <p class="article-meta">By {{ article.author || 'Unknown' }} - {{ formatDate(article.CreatedAt) }}</p>
            <p>{{ article.content.substring(0, 100) }}...</p>
            <a :href="'/user/articles/' + article.ID" class="btn btn-success btn-sm">Baca Selengkapnya</a>
          </div>
        </div>
      </div>
    </div>
    <p v-else>Tidak ada artikel ditemukan.</p>

    <!-- Pagination -->
    <div class="d-flex justify-content-between align-items-center mt-4">
      <button :disabled="page <= 1" class="btn btn-outline-primary" @click="prevPage">Previous</button>
      <span>Halaman {{ page }} dari {{ Math.ceil(meta.total_data / limit) || 1 }}</span>
      <button :disabled="page >= Math.ceil(meta.total_data / limit)" class="btn btn-outline-primary" @click="nextPage">Next</button>
    </div>
  </div>
  <Footer />
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

definePageMeta({
  middleware: 'auth',
});

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

// Route dan Router
const route = useRoute()
const router = useRouter()

// State
const articles = ref([])
const meta = ref({})
const page = ref(Number(route.query.page) || 1)
const limit = ref(10)
const loading = ref(false)

// Fungsi API untuk mencari artikel
const searchArticles = async () => {
  loading.value = true
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Token yang diambil:', token);  // Debugging token
    const response = await axios.get('http://localhost:8080/user/articles/search', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params: {
        query: route.query.query || '',
        category_id: route.query.category_id || '',
        page: page.value,
        limit: limit.value,
      },
    })
    articles.value = response.data.data
    meta.value = response.data.meta
  } catch (error) {
    console.error('Failed to fetch articles:', error)
  } finally {
    loading.value = false
  }
}

// Fungsi format tanggal
const formatDate = (date) => {
  const options = { year: 'numeric', month: 'long', day: 'numeric' }
  return new Date(date).toLocaleDateString('id-ID', options)
}

// Navigasi ke halaman sebelumnya
const prevPage = () => {
  if (page.value > 1) {
    page.value--
    updateQuery()
  }
}

// Navigasi ke halaman berikutnya
const nextPage = () => {
  if (page.value < Math.ceil(meta.value.total_data / limit.value)) {
    page.value++
    updateQuery()
  }
}

// Perbarui query parameter di URL
const updateQuery = () => {
  router.push({
    path: '/user/search',
    query: {
      ...route.query,
      page: page.value,
    },
  })
}

// Panggil API saat parameter query berubah
watch(() => route.query, () => {
  page.value = Number(route.query.page) || 1
  searchArticles()
}, { immediate: true })
</script>

<style scoped>
.article-title {
  font-size: 1.25rem;
  font-weight: bold;
}
.article-meta {
  font-size: 0.9rem;
  color: #6c757d;
}
</style>
