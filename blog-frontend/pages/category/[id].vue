<template>
  <LoginNavbar v-if="isLoggedIn"/>
  <Navbar v-else/>
    <div class="container mt-4">
      <div class="row">
        <div class="col-lg-8">
        <!-- Breadcrumb -->
        <nav aria-label="breadcrumb" v-if="category.name">
          <ol class="breadcrumb">
            <li class="breadcrumb-item">
              <NuxtLink to="/user/dashboard" v-if="isLoggedIn">Home</NuxtLink>
              <NuxtLink to="/" v-else>Home</NuxtLink>
            </li>
            <li class="breadcrumb-item">
              <NuxtLink :to="`/category/${category.ID}`">{{ category.name }}</NuxtLink>
            </li>
          </ol>
        </nav>
        <h5 style="color: #1D2B53;">Menampilkan Hasil dari #{{ category.name }}</h5>
        <div class="row">
          <div v-for="article in category.articles" :key="article.ID" class="col-md-6 mb-3">
            <div class="card">
              <img :src="`http://localhost:8080/uploads/${article.file_name}`" class="img-top" alt="img-artikel">
              <div class="card-body">
                <h6 class="text-muted" style="color: #1D2B53;">{{ article.label }}</h6>
                <h5 class="card-title" style="color: #1D2B53;">{{ article.title }}</h5>
                <p class="card-text" style="color: #1D2B53;" v-html="getTruncatedContent(article.content)"></p>
                <p class="text-muted" style="color: #1D2B53;">Dibuat : {{ formatDate(article.CreatedAt) }}</p>
                <NuxtLink :to="`/user/articles/${article.ID}`" class="btn">Baca Selengkapnya   <i class="bi bi-arrow-right-circle"></i></NuxtLink>
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
import { format } from 'date-fns'; // Format tanggal
import { id } from 'date-fns/locale'; // Locale Indonesia
import DOMPurify from 'dompurify';

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});
const category = ref({});
const error = ref('');

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Format tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

const fetchCategory = async (id)=>{
  try{
    const response = await axios.get(`http://localhost:8080/category/${id}`)
    console.log('Respons API:', response.data);
    category.value = response.data.data;
  } catch(err) {
    error.value = err.response?.data?.message || err.message;
    console.error('Error fetching category:', err);
  }
};

// Truncate konten artikel
function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 50).join(" ") + "...";
  return DOMPurify.sanitize(truncated); // Aman untuk dirender
}

// Mengambil ID dari URL
const route = useRoute();
onMounted(() => {
  const categoryId = route.params.id;
  fetchCategory(categoryId);
});
</script>

<style scoped>
body {
  transition: all 0.3s ease;
}

.breadcrumb {
  background-color: transparent;
  padding: 0;
  margin-bottom: 1rem;
}
.breadcrumb-item a {
  text-decoration: none;
  color: #FF004D;
}
.breadcrumb-item.active {
  color: #1D2B53;
}
.btn{
  color: #1D2B53;
  background-color: #FF004D;
  border-color: #1D2B53;
}
.btn:hover{
  color: #FF004D;
  background-color: #1D2B53;
  border-color: #FF004D;
}
</style>