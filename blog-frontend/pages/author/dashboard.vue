<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AuthorSidebar />
      <!-- Main Content -->
      <main class="p-4 col-md-12 col-md-9 col-lg-10">
        <h1>Workspace Author</h1>
        <hr>

        <div class="row">
          <div class="col-sm-6 mb-3 mb-sm-0">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title">Artikel Anda Yang Belum Di Notice</h5>
                <div v-if="filteredArticlesNoticed.length === 0">Tidak ada artikel yang belum di-notice.</div>
                <div v-for="article in filteredArticlesNoticed" :key="article.id" class="d-flex justify-content-between mb-2">
                  <div class="card">
                    <img
                      :src="`http://localhost:8080/uploads/${article.file_name}`"
                      class="card-img-top"
                      alt="..."
                    />
                    <div class="card-body">
                      <h5 class="card-title">{{ article.title }}</h5>
                      <p class="card-text" v-html="getTruncatedContent(article.content)"></p>
                      <h5 class="text-muted">*Catatan dari admin</h5>
                      <div class="card">
                        <div class="card-body">
                          <p class="card-text" v-html="article.note"></p>
                        </div>
                        <button class="btn btn-outline-warning mb-2 mx-2" @click="navigateToEditArticle(article.ID)">
                          <i class="bi bi-pencil"></i>  Perbaiki
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-6">
            <div class="card">
              <div class="card-body">
                <h5 class="card-title">Artikel Yang Berhasil Di Publikasi</h5>
                <div v-if="filteredArticlesPublished.length === 0">Tidak ada artikel yang berhasil dipublikasi.</div>
                <div v-for="article in filteredArticlesPublished" :key="article.id" class="d-flex justify-content-between mb-2">
                  <div class="card">
                    <img
                      :src="`http://localhost:8080/uploads/${article.file_name}`"
                      class="card-img-top"
                      alt="..."
                    />
                    <div class="card-body">
                      <h5 class="card-title">{{ article.title }}</h5>
                      <p class="card-text" v-html="getTruncatedContent(article.content)"></p>
                      <NuxtLink :to="`/author/articles/${article.ID}`" class="btn btn-outline-info"><i class="bi bi-eye"></i> Lihat</NuxtLink>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import DOMPurify from "dompurify";

definePageMeta({
  middleware: 'auth',
  requiresAuthor: true,
});

const articles = ref([]);
const loadingArticles = ref(true);

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  return token;
};

const decodeToken = (token) => {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    return payload;
  } catch (error) {
    console.error('Error decoding token:', error);
    return null;
  }
};

const fetchArticles = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const decoded = decodeToken(token);
    if (!decoded || !decoded.user_id) {
      throw new Error('User ID tidak ditemukan dalam token.');
    }

    const response = await axios.get(`http://localhost:8080/author/${decoded.user_id}/articles`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    articles.value = Array.isArray(response.data.data) ? response.data.data : [];
  } catch (error) {
    console.error('Error fetching articles:', error);
    articles.value = [];
  } finally {
    loadingArticles.value = false;
  }
};

const filteredArticlesNoticed = computed(() => {
  return articles.value.filter(article => article.status === "Belum Sesuai");
});

const filteredArticlesPublished = computed(() => {
  return articles.value.filter(article => article.status === "Sesuai");
});

function getTruncatedContent(content) {
  if (!content) return "";
  const truncated = content.split(" ").slice(0, 50).join(" ") + "...";
  return DOMPurify.sanitize(truncated);
}

const router = useRouter()

const navigateToEditArticle = (id) => {
  router.push(`/author/articles/edit/${id}`);
};

onMounted(() => {
  fetchArticles();
});
</script>

<style scoped>
.ql-editor p:has(img) {
  display: flex;
  justify-content: center;
  margin: 0;
  padding: 0;
}
.ql-editor p{
  background-color: transparent;
}
</style>