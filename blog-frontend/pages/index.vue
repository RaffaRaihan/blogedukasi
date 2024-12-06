<template>
    <Navbar />
    <!-- Content -->
    <div class="container mt-4">
      <div class="row">
        <!-- Main Content -->
        <div class="col-lg-8">
          <div class="card mb-4">
            <div id="carouselExampleCaptions" class="carousel slide">
              <div class="carousel-indicators">
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="3" aria-label="Slide 4"></button>
                <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="4" aria-label="Slide 5"></button>
              </div>
              <div class="carousel-inner">
                <div class="carousel-item active">
                  <img src="../assets/img/Inumaki.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/𓆩⛧𓆪.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/wallpaper.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/20487358.jpg" class="d-block w-100" alt="...">
                </div>
                <div class="carousel-item">
                  <img src="../assets/img/de2158bba1bf3c5b60bb234b16b7545b.jpg" class="d-block w-100" alt="...">
                </div>
              </div>
              <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
                <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                <span class="visually-hidden">Previous</span>
              </button>
              <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
                <span class="carousel-control-next-icon" aria-hidden="true"></span>
                <span class="visually-hidden">Next</span>
              </button>
            </div>
          </div>

          <h5>Postingan Terbaru</h5>
          <div class="row">
            <div v-for="articles in paginatedArticles" :key="articles.id" class="col-md-6 mb-3">
              <div class="card">
                <img :src="`http://localhost:8080/uploads/${articles.file_name}`" class="card-img-top" alt="...">
                <div class="card-body">
                  <h6 class="text-muted">{{ articles.label }}</h6>
                  <h5 class="card-title">{{ articles.title }}</h5>
                  <p class="card-text">{{ articles.content }}</p>
                  <p class="text-muted">{{ formatDate(articles.CreatedAt) }}</p>
                  <NuxtLink to="/login" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
                </div>
              </div>
            </div>
          </div>
        <!-- Paginations -->
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
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';  // Import date-fns
import { id } from 'date-fns/locale';  // Import locale for Indonesian

// Function to format the date
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

const articles = ref([]);
const currentPage = ref(1); // Track the current page
const articlesPerPage = 4;  // Number of articles per page

const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles'); // Ganti dengan URL backend Anda
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Calculate the paginated articles for the current page
const getPaginatedArticles = () => {
  const start = (currentPage.value - 1) * articlesPerPage;
  return articles.value.slice(start, start + articlesPerPage);
};

// Function to change the current page
const changePage = (page) => {
  if (page < 1 || page > totalPages.value) return;
  currentPage.value = page;
};

// Calculate total pages
const totalPages = computed(() => {
  return Math.ceil(articles.value.length / articlesPerPage);
});

// Watch for changes in articles to update the paginated articles
const paginatedArticles = computed(() => getPaginatedArticles());

// Panggil fetchArticles saat komponen dimuat
onMounted(() => {
  fetchArticles();
});
</script>
