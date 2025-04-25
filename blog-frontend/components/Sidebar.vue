<template>
  <!-- Sidebar -->
  <div class="col-lg-4 mb-3">
    <div class="mb-4">
      <h5>Paling Populer</h5>
      <ul v-if="popularArticles.length > 0" class="list-group">
        <li v-for="article in popularArticles.slice(0, 3)" :key="article.article_id" class="category list-group-item">
          <strong>{{ article.title }}</strong> <br>
          <span class="text-muted">{{ formatDate(article.created_at) }}</span><br>
          <span class="text-muted"><i class="bi bi-eye"></i> {{ article.views }} views</span><br>
          <NuxtLink :to="`/user/articles/${article.article_id}`" class="btn btn-sm mt-2 mb-2" @click="trackArticleView(article.article_id)">
            Lihat
          </NuxtLink>
        </li>
      </ul>
      <p v-else>Tidak ada artikel populer.</p>
    </div>
    <div class="mb-4">
      <h5>Kategori</h5>
      <ul v-for="category in category" :key="category.ID" class="list-group">
        <li class="category list-group-item mb-2">
          <NuxtLink class="text-decoration-none" style="color: #211951;" :to="`/category/${category.ID}`">#{{ category.name }}</NuxtLink>
        </li>
      </ul>
    </div>
    <div class="mb-4">
      <h5>Pelayanan</h5>
      <ul class="list-unstyled d-flex jutify-content-between">
        <li class="service me-2"><NuxtLink class="text-decoration-none" style="color: #211951;" to="/user/send-message">Hubungi Admin</NuxtLink></li>
        <li class="service me-2"><NuxtLink class="text-decoration-none" style="color: #211951;" to="/user/send-message">Kritik Dan Saran</NuxtLink></li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import useCategory from '@/composables/api/useCategory';
import useArticles from '~/composables/api/useArticles';

const popularArticles = ref([])
const errorArticles = ref(null)

const{ articles, trackArticleView } = useArticles()

// Fungsi untuk memformat tanggal
const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  });
};

const fetchPopularArticles = async () => {
    try {
      const response = await axios.get('http://localhost:8080/popular-articles');
      console.log("Popular Articles Response:", response.data);
      popularArticles.value = response.data // Jika response.data undefined, set array kosong
    } catch (error) {
      console.error("Error fetching popular articles:", error);
      errorArticles.value = error;
    } 
  };

const { category, fetchCategory } = useCategory();
fetchCategory();

onMounted(fetchPopularArticles);
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