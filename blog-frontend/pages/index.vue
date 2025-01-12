<template>
  <Navbar />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <!-- Carousel -->
        <div class="card mb-4">
          <div id="carouselExampleSlidesOnly" class="carousel slide" data-bs-ride="carousel" data-bs-interval="3000">
            <div class="carousel-inner">
              <div class="carousel-item active">
                <img src="/assets/img/Shinobu.jpg" class="d-block w-100" alt="...">
              </div>
              <div class="carousel-item">
                <img src="/assets/img/Gojo-lanscape.jpeg" class="d-block w-100" alt="...">
              </div>
              <div class="carousel-item">
                <img src="/assets/img/tanjiro.jpeg" class="d-block w-100" alt="...">
              </div>
              <div class="carousel-item">
                <img src="/assets/img/Tanjiro-2.jpeg" class="d-block w-100" alt="...">
              </div>
            </div>
          </div>
        </div>

        <!-- Loading Indicator -->
         <div class="mt-3 mt-3" v-if="loadingArticles" >
          <Loading />
         </div>
        
        <!-- Content -->
        <div v-else>
          <div class="brutalist-container mb-3">
            <input
              class="form-control brutalist-input smooth-type"
              v-model="searchQuery"
              placeholder="CARI ARTIKEL"
            />
            <label class="brutalist-label text-uppercase">Cari Artikel</label>
          </div>
          <h5 style="color: #1D2B53;">Postingan Terbaru</h5>
          <div class="row">
          <!-- Tampilkan pesan jika tidak ada artikel -->
          <div v-if="paginatedArticles.length === 0" class="text-center">
            <p style="color: #1D2B53;">Tidak ada artikel yang ditemukan.</p>
          </div>
            <div v-else v-for="articlesItem in paginatedArticles" :key="articlesItem.ID" class="col-md-6 mb-3">
              <div class="card">
                <img :src="`http://localhost:8080/uploads/${articlesItem.file_name}`" class="card-img-top" alt="...">
                <div class="card-body">
                  <h6 class="text-muted" style="color: #1D2B53;">{{ articlesItem.label }}</h6>
                  <h5 class="card-title" style="color: #1D2B53;">{{ articlesItem.title }}</h5>
                  <p class="card-text" style="color: #1D2B53;" v-html="getTruncatedContent(articlesItem.content)"></p>
                  <p class="text-muted" style="color: #1D2B53;">Dibuat : {{ formatDate(articlesItem.CreatedAt) }}</p>
                  <NuxtLink :to="`/user/articles/${articlesItem.ID}`" class="btn">Baca Selengkapnya   <i class="bi bi-arrow-right-circle"></i></NuxtLink>
                </div>
              </div>
            </div>
          </div>

          <!-- Pagination -->
          <nav aria-label="Page navigation example">
            <ul class="pagination justify-content-center">
              <li class="page-item" :class="{'disabled': currentPage === 1}">
                <button @click.prevent="changePage(currentPage - 1)" class="page-link">Previous</button>
              </li>
              <li v-for="page in totalPages" :key="page" class="page-item" :class="{'active': currentPage === page}">
                <button @click.prevent="changePage(page)" class="page-link">{{ page }}</button>
              </li>
              <li class="page-item" :class="{'disabled': currentPage === totalPages}">
                <button @click.prevent="changePage(currentPage + 1)" class="page-link">Next</button>
              </li>
            </ul>
          </nav>
        </div>
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
const loadingArticles = ref(true);

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
  } finally {
    loadingArticles.value = false;
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

  // Scroll ke atas setelah ganti halaman
  window.scrollTo({
    top: 0,
    behavior: 'smooth', // Scroll dengan animasi halus
  });
};

// Truncate konten artikel
function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 50).join(" ") + "...";
  return DOMPurify.sanitize(truncated); // Aman untuk dirender
}

// Fetch data saat komponen di-mount
onMounted(() => {
  fetchArticles();
});
</script>

<style scoped>
body {
  transition: all 0.3s ease;
}

.btn{
  color: #1D2B53;
  background-color: #FF004D;
  border-color: #1D2B53;
}
.btn:hover{
  color: #F9F6E6;
  background-color: #1D2B53;
}
.brutalist-container {
  position: relative;
}
.brutalist-input {
  width: 100%;
  padding: 15px;
  font-size: 18px;
  font-weight: bold;
  color: #1D2B53;
  background-color: #fff;
  border: 4px solid #1D2B53;
  position: relative;
  overflow: hidden;
  border-radius: 0;
  outline: none;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  box-shadow: 5px 5px 0 #1D2B53, 10px 10px 0 #FF004D;
  margin-bottom: 2rem;
  margin-top: 3rem;
}
@keyframes glitch {
  0% {
    transform: translate(0);
  }
  20% {
    transform: translate(-2px, 2px);
  }
  40% {
    transform: translate(-2px, -2px);
  }
  60% {
    transform: translate(2px, 2px);
  }
  80% {
    transform: translate(2px, -2px);
  }
  100% {
    transform: translate(0);
  }
}
.brutalist-input:focus {
  animation: focus-pulse 4s cubic-bezier(0.25, 0.8, 0.25, 1) infinite,
    glitch 0.3s cubic-bezier(0.25, 0.8, 0.25, 1) infinite;
}
.brutalist-input:focus::after {
  content: "";
  position: absolute;
  top: -2px;
  left: -2px;
  right: -2px;
  bottom: -2px;
  background: white;
  z-index: -1;
}
.brutalist-input:focus::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #1D2B53;
  z-index: -2;
  clip-path: inset(0 100% 0 0);
  animation: glitch-slice 4s steps(2, end) infinite;
}
@keyframes glitch-slice {
  0% {
    clip-path: inset(0 100% 0 0);
  }
  10% {
    clip-path: inset(0 5% 0 0);
  }
  20% {
    clip-path: inset(0 80% 0 0);
  }
  30% {
    clip-path: inset(0 10% 0 0);
  }
  40% {
    clip-path: inset(0 50% 0 0);
  }
  50% {
    clip-path: inset(0 30% 0 0);
  }
  60% {
    clip-path: inset(0 70% 0 0);
  }
  70% {
    clip-path: inset(0 15% 0 0);
  }
  80% {
    clip-path: inset(0 90% 0 0);
  }
  90% {
    clip-path: inset(0 5% 0 0);
  }
  100% {
    clip-path: inset(0 100% 0 0);
  }
}
.brutalist-label {
  position: absolute;
  left: -3px;
  top: -35px;
  font-size: 14px;
  font-weight: bold;
  color: #fff;
  background-color: #1D2B53;
  padding: 5px 10px;
  transform: rotate(-1deg);
  z-index: 1;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  margin-top: 1rem;
}
.brutalist-input:focus + .brutalist-label {
  transform: rotate(0deg) scale(1.05);
  background-color: #FF004D;
}
.smooth-type {
  position: relative;
  overflow: hidden;
}
.smooth-type::before {
  content: "";
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background: linear-gradient(90deg, #fff 0%, rgba(255, 255, 255, 0) 100%);
  z-index: 1;
  opacity: 0;
  transition: opacity 0.3s ease;
}
.smooth-type:focus::before {
  opacity: 1;
  animation: type-gradient 2s linear infinite;
}
@keyframes type-gradient {
  0% {
    background-position: 300px 0;
  }
  100% {
    background-position: 0 0;
  }
}
.brutalist-input::placeholder {
  color: #888;
  transition: color 0.3s ease;
}
.brutalist-input:focus::placeholder {
  color: transparent;
}
.brutalist-input:focus {
  animation: focus-pulse 4s cubic-bezier(0.25, 0.8, 0.25, 1) infinite;
}
@keyframes focus-pulse {
  0%,
  100% {
    border-color: #1D2B53;
  }
  50% {
    border-color: #FF004D;
  }
}

.page-link {
  color: #1D2B53;
}
.page-item{
  color: #FF004D;
}
</style>
