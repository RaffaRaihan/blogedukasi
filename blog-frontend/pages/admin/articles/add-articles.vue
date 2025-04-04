<template>
  <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
  <div class="container mb-4">
    <h1 class="my-4">Tambah Artikel Baru</h1>
    <form @submit.prevent="submitArticle">
      <div class="mb-3">
        <label for="articleLabel" class="form-label">Label</label>
        <input
          type="text"
          id="articleLabel"
          class="form-control"
          v-model="newArticle.label"
          required
          placeholder="Masukkan label"
        />
      </div>
      <div class="mb-3">
        <label for="articleTitle" class="form-label">Judul Artikel</label>
        <input
          type="text"
          id="articleTitle"
          class="form-control"
          v-model="newArticle.title"
          required
          placeholder="Masukkan judul artikel"
        />
      </div>
      <div class="mb-3">
        <label for="articleContent" class="form-label">Konten Artikel</label>
        <QuillEditor v-model="newArticle.content" />
      </div>
      <div class="mb-3">
        <label for="articleCategory" class="form-label">Kategori</label>
        <select
          id="articleCategory"
          class="form-select"
          v-model="newArticle.category_id"
          required
        >
          <option value="" disabled>Pilih kategori</option>
          <option v-for="category in categories" :key="category.ID" :value="category.ID">
            {{ category.name }}
          </option>
        </select>
      </div>
      <div class="mb-3">
        <label for="articleImage" class="form-label">Gambar Artikel</label>
        <input
          type="file"
          id="articleImage"
          class="form-control"
          ref="fileInput"
        />
      </div>
      <button type="submit" class="btn btn-primary me-2">Tambah Artikel</button>
      <NuxtLink to="/admin/articles" class="btn btn-danger">Batal</NuxtLink>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const router = useRouter();
const newArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author_id: '',
});
const categories = ref([]);
const newArticleId = ref(null); // Untuk menyimpan ID artikel yang baru dibuat
const alertMessage = ref('');
const alertClass = ref('');

const getTokenFromCookies = () => {
  return document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
};

// Fungsi untuk mendekode payload token JWT
const parseJwt = (token) => {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    return JSON.parse(atob(base64));
  } catch (error) {
    console.error('Gagal mendekode token:', error);
    return null;
  }
};

const setUserIdFromToken = () => {
  const token = getTokenFromCookies();
  if (token) {
    const decodedToken = parseJwt(token);
    if (decodedToken?.user_id) {
      newArticle.value.author_id = decodedToken.user_id; // Set ID user sebagai author
    }
  }
};

const fetchCategories = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category');
    categories.value = response.data;
  } catch (error) {
    console.error('Error fetching categories:', error);
  }
};

const submitArticle = async () => {
  try {
    // Ambil token dari cookie atau localStorage
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      setTimeout(() => {
        window.location.href = "/login";
      }, 3000);
      return;
    }

    if (!newArticle.value.author_id) {
      setUserIdFromToken();
    }

    // Kirim artikel pertama
    const response = await axios.post(
      'http://localhost:8080/admin/articles',
      newArticle.value,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Dapatkan ID artikel yang baru dibuat
    const articleId = response.data.data.ID;

    // Simpan ID artikel untuk digunakan saat upload gambar
    if (articleId) {
      newArticleId.value = articleId; // Simpan ID artikel
    }

    // Jika ada file gambar, unggah
    const fileInput = document.getElementById('articleImage');
    const file = fileInput?.files[0];
    if (file) {
      await handleImageUpload(file, newArticleId.value);
    }

    alertMessage.value = `Artikel berhasil dibuat.`;
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.href = "/admin/articles";
    }, 3000);
  } catch (error) {
    console.error('Error submitting article:', error);
    alertMessage.value = `Gagal membuat artikel`;
    alertClass.value = 'alert alert-danger';
    setTimeout(() => {
      window.location.href = "/admin/articles";
    }, 3000);
  }
};

const handleImageUpload = async (file, articleId) => {
  const formData = new FormData();
  formData.append('file', file);

  try {
    // Ambil token dari cookie atau localStorage
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      setTimeout(() => {
        window.location.href = "/login";
      }, 3000);
      return;
    }

    const response = await axios.post(
      `http://localhost:8080/admin/articles/${articleId}/uploads`,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log('File uploaded:', response.data);
  } catch (error) {
    console.error('Error uploading file:', error);
    alertMessage.value = `Gagal menggungah gambar`;
    alertClass.value = 'alert alert-danger';
  }
};

onMounted(() => {
  fetchCategories();
  setUserIdFromToken(); // Ambil ID user saat komponen dimuat
});
</script>

<style scoped>
.alert {
  position: fixed;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 500px;
  z-index: 1050;
  text-align: center;
}
</style>