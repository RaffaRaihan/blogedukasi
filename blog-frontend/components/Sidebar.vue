<template>
  <!-- Sidebar -->
  <div class="col-lg-4 mb-3">
    <div class="mb-4">
      <h5>Paling Populer</h5>
      <!-- Loading Indicator -->
      <Loading v-if="loadingArticles" />
      <ul v-for="article in popularArticles" :key="article.id" class="list-group" v-else>
        <li class="category-1 list-group-item mb-3" style="color: #211951;">
          {{ article.title }} <br>
          <span class="text-muted" style="color: #211951;">{{ formatDate(article.CreatedAt) }}</span><br>
          <NuxtLink :to="`/user/articles/${article.ID}`" class="btn btn-sm mt-2 mb-2"><i class="bi bi-eye"></i>  Lihat</NuxtLink>
        </li>
      </ul>
    </div>
    <div class="mb-4">
      <h5>Kategori</h5>
      <Loading v-if="loadingArticles" />
      <ul v-for="category in category" :key="category.id" class="list-group" v-else>
        <li class="category list-group-item mb-2">
          <NuxtLink class="text-decoration-none" style="color: #211951;" :to="`/category/${category.ID}`">#{{ category.name }}</NuxtLink>
        </li>
      </ul>
    </div>
    <div class="mb-4">
      <h5>Pelayanan</h5>
      <ul class="list-unstyled d-flex jutify-content-between">
        <li class="service me-2"><NuxtLink class="text-decoration-none" style="color: #211951;" to="/user/send-message">Hubungi Admin</NuxtLink></li>
        <li class="service me-2"><NuxtLink class="text-decoration-none" style="color: #211951;" to="/user/send-message">Laporkan Bug</NuxtLink></li>
        <li class="service me-2"><NuxtLink class="text-decoration-none" style="color: #211951;" to="/user/send-message">Kritik Dan Saran</NuxtLink></li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import useArticles from '@/composables/api/useArticles';
import useCategory from '@/composables/api/useCategory';
import { onMounted, computed } from 'vue';
import { format } from 'date-fns'; // Untuk memformat tanggal
import { id } from 'date-fns/locale'; // Locale Indonesia

// Fungsi untuk memformat tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

// Fungsi untuk mendapatkan data artikel
const { articles, loadingArticles, errorArticles } = useArticles();

const { category, loadingCategory, fetchCategory } = useCategory();

// Mengambil 3 artikel secara acak
const popularArticles = computed(() => {
  if (articles.value.length <= 3) return articles.value; // Jika artikel kurang dari atau sama dengan 3
  const shuffled = [...articles.value].sort(() => 0.5 - Math.random());
  return shuffled.slice(0, 3); // Ambil 3 artikel pertama setelah diacak
});

// Panggil API saat komponen dimuat
onMounted(() => {
  fetchCategory();
  useArticles();
});
</script>

<style scoped>
.btn{
  color: #F0F3FF;
  background-color: #1D2B53;
  border-color: #1D2B53;
}
.btn:hover{
  color: #211951;
  background-color: #F0F3FF;
  transition: 0.2s;
}

.category{
  border-color: #211951;
}

.category-1{
  border-color: #211951;
}
.category:hover {
  background-color: #dce1f7;
  color: #211951;
  transition: 0.3s;
}
.service:hover{
  text-decoration: underline;
  transition: 0.3s;
}
</style>