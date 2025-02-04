<template>
  <div class="container">
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
        <label for="articleAuthor" class="form-label">Author</label>
        <input
          type="text"
          id="articleAuthor"
          class="form-control"
          v-model="newArticle.author"
          required
          placeholder="Masukkan author"
        />
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
      <button type="submit" class="btn btn-primary">Tambah Artikel</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const router = useRouter();
const newArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author: '',
});
const categories = ref([]);
const newArticleId = ref(null); // Untuk menyimpan ID artikel yang baru dibuat

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
      alert('Token tidak ditemukan. Pastikan Anda sudah login.');
      return;
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

    alert('Artikel berhasil dibuat.');
    router.push('/admin/articles');
  } catch (error) {
    console.error('Error submitting article:', error);
    alert('Gagal membuat artikel');
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
      alert('Token tidak ditemukan. Pastikan Anda sudah login.');
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
    alert('Gagal mengunggah gambar');
  }
};

onMounted(fetchCategories);
</script>