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
                  <img src="/assets/img/Inumaki.jpg" class="d-block w-100" alt="...">
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

        <h5>Postingan Terbaru</h5>
        <div class="row">
          <div v-for="articlesItem in paginatedArticles" :key="articlesItem.ID" class="col-md-6 mb-3">
            <div class="card">
              <img :src="`http://localhost:8080/uploads/${articlesItem.file_name}`" class="card-img-top" alt="...">
              <div class="card-body">
                <h6 class="text-muted">{{ articlesItem.label }}</h6>
                <h5 class="card-title">{{ articlesItem.title }}</h5>
                <p class="card-text" v-html="articlesItem.content.substring(0, 30)"></p>
                <p class="text-muted">{{ formatDate(articlesItem.CreatedAt) }}</p>
                <NuxtLink :to="`/user/articles/${articlesItem.ID}`" class="btn btn-success btn-sm">Baca Selengkapnya</NuxtLink>
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
      <LoginSidebar />
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
const currentPage = ref(1); // Track the current page
const articlesPerPage = 4;  // Number of articles per page

// Function to fetch articles from the backend
const fetchArticles = async () => {
  try {
    const response = await axios.get('http://localhost:8080/articles');
    articles.value = response.data.sort((a, b) => new Date(b.CreatedAt) - new Date(a.CreatedAt));
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

// Fetch articles when component is mounted
onMounted(() => {
  fetchArticles();
});
</script>
