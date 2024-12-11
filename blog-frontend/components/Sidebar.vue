<template>
  <!-- Sidebar -->
  <div class="col-lg-4 mb-3">
    <div class="mb-4">
      <h5 style="color: #1D2B53;">Paling Populer</h5>
      <!-- Loading Indicator -->
      <Loading v-if="loadingArticles" />
      <ul v-for="article in popularArticles" :key="article.id" class="list-group" v-else>
        <li class="list-group-item mb-3" style="color: #1D2B53;">
          {{ article.title }} <br>
          <span class="text-muted" style="color: #1D2B53;">{{ formatDate(article.CreatedAt) }}</span><br>
          <NuxtLink :to="`/user/articles/${article.ID}`" class="btn btn-sm mt-2"><i class="bi bi-eye"></i>  Lihat</NuxtLink>
        </li>
      </ul>
    </div>
    <div>
      <h5>Kategori</h5>
      <Loading v-if="loadingArticles" />
      <ul v-for="category in category" :key="category.id" class="list-group" v-else>
        <li class="category list-group-item mb-2">
          <NuxtLink class="text-decoration-none" style="color: #1D2B53;" to="/">#{{ category.name }}</NuxtLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { format } from 'date-fns'; // Untuk memformat tanggal
import { id } from 'date-fns/locale'; // Locale Indonesia

const category = ref([]);
const articles = ref([]);
const loadingArticles = ref(true);

// Fungsi untuk memformat tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

// Fungsi untuk mendapatkan data artikel
const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles'); // Ganti dengan URL backend Anda
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  } finally {
    loadingArticles.value = false;
  }
};

// Fungsi untuk mendapatkan data kategori
const fetchCategory = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category'); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  } finally {
    loadingArticles.value = false;
  }
};

// Mengambil 3 artikel secara acak
const popularArticles = computed(() => {
  if (articles.value.length <= 3) return articles.value; // Jika artikel kurang dari atau sama dengan 3
  const shuffled = [...articles.value].sort(() => 0.5 - Math.random());
  return shuffled.slice(0, 3); // Ambil 3 artikel pertama setelah diacak
});

// Panggil API saat komponen dimuat
onMounted(() => {
  fetchCategory();
  fetchArticles();
});
</script>

<style scoped>
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

.category:hover {
  background-color: #FF004D;
  color: #1D2B53;
}
</style>