<template>
  <Navbar />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <!-- Carousel -->
        <div class="card mb-4">
          <div id="carouselExampleInterval" class="carousel slide" data-bs-ride="carousel">
            <div class="carousel-indicators">
              <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
              <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
              <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
              <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="3" aria-label="Slide 4"></button>
            </div>
            <div class="carousel-inner">
              <div class="carousel-item active" data-bs-interval="3000">
                <img src="/assets/img/Inumaki.jpg" class="d-block w-100" alt="...">
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
              <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleInterval" data-bs-slide="prev">
                <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                <span class="visually-hidden">Previous</span>
              </button>
              <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleInterval" data-bs-slide="next">
                <span class="carousel-control-next-icon" aria-hidden="true"></span>
                <span class="visually-hidden">Next</span>
              </button>
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
          <h5>Postingan Terbaru</h5>
          <div class="row">
          <!-- Tampilkan pesan jika tidak ada artikel -->
          <div v-if="paginatedArticles.length === 0" class="text-center">
            <p>Tidak ada artikel yang ditemukan.</p>
          </div>
            <div v-else v-for="articlesItem in paginatedArticles" :key="articlesItem.ID" class="col-md-6 mb-3">
              <div class="card shadow bg-body-tertiary rounded">
                <img :src="`http://localhost:8080/uploads/${articlesItem.file_name}`" class="card-img-top" alt="...">
                <div class="card-body">
                  <h5 class="card-title">{{ articlesItem.title }}</h5>
                  <h6 class="text-muted">{{ articlesItem.label }}</h6>
                  <p class="card-text" v-html="getTruncatedContent(articlesItem.content)"></p>
                  <p class="text-muted">Dibuat : {{ formatDate(articlesItem.CreatedAt) }}</p>
                  <NuxtLink :to="`/user/articles/${articlesItem.ID}`" @click="trackArticleView(articlesItem.ID)" class="btn">Baca Selengkapnya   <i class="bi bi-arrow-right-circle"></i></NuxtLink>
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
import useArticles from '@/composables/api/useArticles';
import { ref, computed, onMounted } from 'vue';
import { format } from 'date-fns'; // Format tanggal
import id from 'date-fns/locale/id/index.js'
import DOMPurify from "dompurify";

// State untuk artikel dan pencarian
const currentPage = ref(1); // Halaman saat ini
const articlesPerPage = 4; // Jumlah artikel per halaman

// Format tanggal
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

const { articles, loadingArticles, trackArticleView } = useArticles();

const searchQuery = ref('');

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
  useArticles();
});
</script>

<style scoped>
body {
  transition: all 0.3s ease;
}

.btn{
  color: #F0F3FF;
  background-color: #211951;
  border-color: #211951;
}
.btn:hover{
  color: #211951;
  background-color: #F0F3FF;
}
.brutalist-container {
  position: relative;
}
.brutalist-input {
  width: 100%;
  padding: 15px;
  font-size: 18px;
  font-weight: bold;
  color: #211951;
  background-color: #fff;
  border: 4px solid #211951;
  position: relative;
  overflow: hidden;
  border-radius: 0;
  outline: none;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  box-shadow: 5px 5px 0 #211951;
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
  background: #211951;
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
  background-color: #211951;
  padding: 5px 10px;
  transform: rotate(-1deg);
  z-index: 1;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  margin-top: 1rem;
}
.brutalist-input:focus + .brutalist-label {
  transform: rotate(0deg) scale(1.05);
  background-color: #211951;
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
    border-color: #211951;
  }
  50% {
    border-color: #211951;
  }
}

.pagination .page-item.active .page-link {
  background-color: #211951;
  border-color: #211951;
  color: #F9F6E6;
}
.pagination .page-link {
  color: #211951;
}
.pagination .page-link:hover {
  color: #F9F6E6;
  background-color: #211951;
}
</style>
