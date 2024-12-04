<template>
    <LoginNavbar />
    <div class="container mt-4">
        <div class="row">
            <div class="col-lg-8">
              <h5>Postingan Terbaru</h5>
              <div class="row">
                <div v-for="articles in articles.slice(0, 4)" :key="articles.id" class="col-md-6 mb-3">
                  <div class="card">
                    <img :src="`http://localhost:8080/uploads/${articles.file_name}`" class="card-img-top" alt="...">
                    <div class="card-body">
                      <h6 class="text-muted">{{ articles.label }}</h6>
                      <h5 class="card-title">{{ articles.title }}</h5>
                      <p class="card-text">{{ articles.content }}</p>
                      <p class="text-muted">{{ formatDate(articles.CreatedAt) }}</p>
                      <NuxtLink :to="`/user/articles/${articles.ID}`" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                    </div>
                  </div>
                </div>
              </div>
              <!-- Paginations -->
              <nav aria-label="Page navigation example">
              <ul class="pagination justify-content-center">
                <li class="page-item disabled">
                  <a class="page-link">Previous</a>
                </li>
                <li class="page-item"><a class="page-link active" href="#">1</a></li>
                <li class="page-item"><a class="page-link" href="#">2</a></li>
                <li class="page-item"><a class="page-link" href="#">3</a></li>
                <li class="page-item">
                  <a class="page-link" href="#">Next</a>
                </li>
              </ul>
              </nav>
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
import { format } from 'date-fns';  // Import date-fns
import { id } from 'date-fns/locale';  // Import locale for Indonesian

// Function to format the date
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

definePageMeta({
  middleware: 'auth',
});

const articles = ref([]);

const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles'); // Ganti dengan URL backend Anda
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Panggil fetchArticles saat komponen dimuat
onMounted(() => {
  fetchArticles();
});
</script>
