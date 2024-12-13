<template>
  <LoginNavbar v-if="isLoggedIn"/>
  <Navbar v-else />
  <div class="container mt-4">
    <div class="row">
      <div class="col-lg-8">
        <!-- Breadcrumb -->
        <nav aria-label="breadcrumb" v-if="articles?.category?.name">
          <ol class="breadcrumb">
            <li class="breadcrumb-item">
              <NuxtLink to="/user/dashboard" v-if="isLoggedIn">Home</NuxtLink>
              <NuxtLink to="/" v-else>Home</NuxtLink>
            </li>
            <li class="breadcrumb-item">
              <NuxtLink :to="`/category/${articles.category.ID}`">{{ articles.category.name }}</NuxtLink>
            </li>
            <li class="breadcrumb-item active" aria-current="page">
              {{ articles.title }}
            </li>
          </ol>
        </nav>
        <!-- Article -->
        <div v-if="articles">
          <h2>{{ articles.title }}</h2>
          <p class="text-muted">By {{ articles.author }} - {{ formatDate(articles.CreatedAt) }}</p>        
          <h4>Apa itu {{ articles.title }}?</h4>
          <p v-html="articles.content"></p><br>     
          <img :src="`http://localhost:8080/uploads/${articles.file_name}`" class="card-img-top" alt="...">
          <hr class="mt-5">

          <!-- Comments Section -->
          <div class="mt-5">
            <h4>Komentar</h4>
            <ul class="list-group" v-if="comments.length > 0">
              <li class="list-group-item" v-for="comment in comments" :key="comment.id">
                <strong>{{ comment.user.name }}</strong>: {{ comment.content }}
              </li>
            </ul>
            <p v-else>Belum ada komentar.</p>

            <!-- Add Comment Form -->
            <div class="mt-3 mb-3">
              <h5>Tambah Komentar</h5>
              <form @submit.prevent="submitComment(articles.ID)">
                <div>
                  <textarea v-model="newComment" class="form-control" placeholder="Tulis komentar..."></textarea>
                  <!-- Pastikan articleId dikirim dengan benar -->
                  <button type="submit" class="btn btn-outline-primary mt-3">Comment</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
      <!-- Sidebar -->
      <Sidebar />
    </div>
  </div>
  <Footer />
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';
import { format } from 'date-fns';
import { id } from 'date-fns/locale';
import { useRoute } from 'vue-router';

const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});
const articles = ref([]);
const comments = ref([]);
const newComment = ref('');
const user_id = ref('');
const error = ref(null);

const formatDate = (date) => {
  const formattedDate = new Date(date);
  if (isNaN(formattedDate)) {
    return 'Tanggal tidak valid';
  }
  return format(formattedDate, 'dd MMMM yyyy', { locale: id });
};

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fungsi untuk mendekode token JWT
const decodeToken = () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const decoded = jwtDecode(token); // Dekode JWT
    user_id.value = decoded.user_id; // Ambil email dari payload
    console.log(decoded)
  } catch (error) {
    console.error('Error decoding token:', error);
  }
};

const fetchArticle = async (id) => {
  try {
    const response = await axios.get(`http://localhost:8080/articles/${id}`);
    articles.value = response.data.data;
  } catch (err) {
    error.value = err.response?.data?.message || err.message;
  }
};

const fetchComments = async (articleId) => {
  try {
    const response = await axios.get(`http://localhost:8080/articles/${articleId}/comments`);
    comments.value = response.data;
  } catch (err) {
    console.error('Error fetching comments:', err);
  }
};

const submitComment = async (articleId) => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error("Token tidak ditemukan. Harap login terlebih dahulu.");

    const response = await axios.post(
      `http://localhost:8080/user/articles/${articleId}/comments`,
      {
        content: newComment.value, // Mengambil isi komentar
        article_id: articleId,
        user_id: user_id.value, // Pastikan mengambil nilai user_id yang benar
      },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Tambahkan komentar baru ke daftar komentar
    comments.value.push(response.data);

    // Reset form komentar
    newComment.value = "";

    alert("Komentar berhasil ditambahkan!");
  } catch (err) {
    console.error("Error submitting comment:", err);
    const token = getTokenFromCookies();

    if (!token) {
        alert("Harap login terlebih dahulu.");
        window.location.href = "/login"; // Ganti "/login" dengan path ke halaman login Anda.
    } else {
        alert("Gagal menambahkan komentar.");
    }
  }
};

const route = useRoute();
onMounted(() => {
  const articleId = route.params.id;
  fetchArticle(articleId);
  fetchComments(articleId);
  decodeToken();
});
</script>

<style scoped>
.breadcrumb {
  background-color: transparent;
  padding: 0;
  margin-bottom: 1rem;
}
.breadcrumb-item a {
  text-decoration: none;
  color: #FF004D;
}
.breadcrumb-item.active {
  color: #1D2B53;
}
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
</style>
