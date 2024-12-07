<template>
  <!-- Sidebar -->
  <div class="col-lg-4 mb-3">
    <div class="mb-4">
      <h5>Paling Populer</h5>
      <ul class="list-group">
        <li v-for="article in popularArticles" :key="article.id" class="list-group-item">
          {{ article.title }} <br>
          <span class="text-muted">{{ formatDate(article.CreatedAt) }}</span><br>
          <NuxtLink to="/login" class="btn btn-success btn-sm mt-2">Lihat</NuxtLink>
        </li>
      </ul>
    </div>
    <div>
      <h5>Kategori</h5>
      <ul v-for="category in category" :key="category.id" class="list-group">
        <li class="list-group-item">
          <NuxtLink class="text-muted" to="/">#{{ category.name }}</NuxtLink>
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
  }
};

// Fungsi untuk mendapatkan data kategori
const fetchCategory = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category'); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
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