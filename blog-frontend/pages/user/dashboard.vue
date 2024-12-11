<template>
  <LoginNavbar />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <div class="card mb-4">
            <div id="carouselExampleSlidesOnly" class="carousel slide" data-bs-ride="carousel">
              <div class="carousel-indicators">
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="3" aria-label="Slide 4"></button>
              </div>
              <div class="carousel-inner">
                <div class="carousel-item active">
                  <img src="/assets/img/ssstik.io_1733559722482.jpeg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="/assets/img/ssstik.io_1733559730444.jpeg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="/assets/img/ssstik.io_1733559725987.jpeg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="/assets/img/ssstik.io_1733559722482.jpeg" class="d-block w-100" alt="...">
                </div>
              </div>
            </div>
          </div>

          <div class="mb-3">
            <input
              class="form-control"
              v-model="searchQuery"
              placeholder="Cari artikel berdasarkan judul..."
            />
          </div>
          <h5>Postingan Terbaru</h5>
          <div class="row">
            <div v-for="articlesItem in paginatedArticles" :key="articlesItem.ID" class="col-md-6 mb-3">
              <div class="card">
                <img :src="`http://localhost:8080/uploads/${articlesItem.file_name}`" class="card-img-top" alt="...">
                <div class="card-body">
                  <h6 class="text-muted">{{ articlesItem.label }}</h6>
                  <h5 class="card-title">{{ articlesItem.title }}</h5>
                  <p class="card-text" v-html="getTruncatedContent(articlesItem.content)"></p>
                  <p class="text-muted">{{ formatDate(articlesItem.CreatedAt) }}</p>
                  <NuxtLink :to="`/user/articles/${articlesItem.ID}`" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                </div>
              </div>
            </div>
          </div>
          
          <!-- Pagination -->
          <nav aria-label="Page navigation example">
            <ul class="pagination justify-content-center">
              <li class="page-item" :class="{'disabled': currentPage === 1}">
                <a @click.prevent="changePage(currentPage - 1)" class="page-link">Previous</a>
              </li>
              <li v-for="page in totalPages" :key="page" class="page-item" :class="{'active': currentPage === page}">
                <a @click.prevent="changePage(page)" class="page-link">{{ page }}</a>
              </li>
              <li class="page-item" :class="{'disabled': currentPage === totalPages}">
                <a @click.prevent="changePage(currentPage + 1)" class="page-link">Next</a>
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
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import { format } from 'date-fns'; // Format tanggal
import { id } from 'date-fns/locale'; // Locale Indonesia
import DOMPurify from "dompurify";

// State untuk artikel dan pencarian
const articles = ref([]);
const searchQuery = ref('');
const currentPage = ref(1); // Halaman saat ini
const articlesPerPage = 4; // Jumlah artikel per halaman

// Format tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

// Ambil artikel dari backend
const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles');
    articles.value = response.data.sort((a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt));
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Filter artikel berdasarkan pencarian
const filteredArticles = computed(() => {
  if (!searchQuery.value.trim()) return articles.value; // Tampilkan semua artikel jika pencarian kosong
  return articles.value.filter((article) =>
    article.title.toLowerCase().includes(searchQuery.value.toLowerCase())
  );
});

// Pagination
const totalPages = computed(() => {
  return Math.ceil(filteredArticles.value.length / articlesPerPage);
});

const getPaginatedArticles = () => {
  const start = (currentPage.value - 1) * articlesPerPage;
  return filteredArticles.value.slice(start, start + articlesPerPage);
};

const paginatedArticles = computed(() => getPaginatedArticles());

const changePage = (page) => {
  if (page < 1 || page > totalPages.value) return;
  currentPage.value = page;
};

// Truncate konten artikel
function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 20).join(" ") + "...";
  return DOMPurify.sanitize(truncated); // Aman untuk dirender
}

// Fetch data saat komponen di-mount
onMounted(() => {
  fetchArticles();
});
</script>
