<template>
  <LoginNavbar />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <!-- Breadcrumb -->
        <nav aria-label="breadcrumb" v-if="articles?.category?.name">
          <ol class="breadcrumb">
            <li class="breadcrumb-item">
              <NuxtLink to="/user/dashboard">Home</NuxtLink>
            </li>
            <li class="breadcrumb-item">
              <NuxtLink :to="`/categories/${articles.category.id}`">{{ articles.category.name }}</NuxtLink>
            </li>
            <li class="breadcrumb-item active" aria-current="page">
              {{ articles.title }}
            </li>
          </ol>
        </nav>
        <!-- Article -->
        <div v-if="articles">
          <h2>{{ articles.title }}</h2>
          <p class="text-muted">By {{ articles.author }} - {{ formatDate(articles.CreatedAt) }}</p>        
          <h4>Apa itu {{ articles.title }}?</h4>
          <p>{{ articles.content }}</p>        
          <img :src="`http://localhost:8080/uploads/${articles.file_name}`" class="card-img-top" alt="...">
          <p>{{ articles.content }}</p>
        </div>
      </div>
      <!-- Sidebar -->
      <Sidebar />
    </div>
  </div>
  <Footer />
</template>

<style scoped>
.breadcrumb {
  background-color: transparent;
  padding: 0;
  margin-bottom: 1rem;
}
.breadcrumb-item a {
  text-decoration: none;
  color: #00A885;
}
.breadcrumb-item.active {
  color: #6c757d;
}
</style>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';
import { id } from 'date-fns/locale';
import { useRoute } from 'vue-router';

definePageMeta({
  middleware: 'auth', // Middleware untuk autentikasi
});

const articles = ref([]);
const error = ref(null);

// Fungsi untuk memformat tanggal
const formatDate = (date) => {
  const formattedDate = new Date(date);
  if (isNaN(formattedDate)) {
    return 'Tanggal tidak valid'; // Tampilkan pesan default jika tanggal tidak valid
  }
  return format(formattedDate, 'dd MMMM yyyy', { locale: id });
};

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchArticle = async (id) => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const response = await axios.get(`http://localhost:8080/user/articles/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    console.log('Respons API:', response.data);
    articles.value = response.data.data;
  } catch (err) {
    error.value = err.response?.data?.message || err.message;
    console.error('Error fetching article:', err);
  }
};

// Mengambil ID dari URL
const route = useRoute();
onMounted(() => {
  const articleId = route.params.id;
  fetchArticle(articleId);
});
</script>
