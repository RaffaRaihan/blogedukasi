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
        <h5 style="color: #211951;">Menampilkan Hasil dari #{{ category.name }}</h5>
        <div class="row">
          <div v-for="article in category.articles" :key="article.ID" class="col-md-6 mb-3">
            <div class="card">
              <img :src="`http://localhost:8080/uploads/${article.file_name}`" class="img-top" alt="img-artikel">
              <div class="card-body">
                <h6 class="text-muted" style="color: #211951;">{{ article.label }}</h6>
                <h5 class="card-title" style="color: #211951;">{{ article.title }}</h5>
                <p class="card-text" style="color: #211951;" v-html="getTruncatedContent(article.content)"></p>
                <p class="text-muted" style="color: #211951;">Dibuat : {{ formatDate(article.CreatedAt) }}</p>
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
import useFetchCategory from '~/composables/api/useCategoryById';
import { format } from 'date-fns'; // Format tanggal
import { id } from 'date-fns/locale'; // Locale Indonesia
import DOMPurify from 'dompurify';
import useAuth from '~/composables/api/token/useAuth';

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});

const { getTokenFromCookies } = useAuth()

const { category } = useFetchCategory();

// Format tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

// Truncate konten artikel
function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 50).join(" ") + "...";
  return DOMPurify.sanitize(truncated); // Aman untuk dirender
}
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
  color: #15F5BA;
}
.breadcrumb-item.active {
  color: #1D2B53;
}
.btn{
  color: #F0F3FF;
  background-color: #211951;
  border-color: #211951;
}
.btn:hover{
  color: #211951;
  background-color: #F0F3FF;
  border-color: #211951;
}
</style>