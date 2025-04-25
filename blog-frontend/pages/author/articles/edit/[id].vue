<template>
  <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
  <div class="container mt-3 mb-3">
    <h1>Edit Artikel</h1>
    <form @submit.prevent="updateArticle">
      <div class="mb-3">
        <label for="editTitle" class="form-label">Judul Artikel</label>
        <input
          type="text"
          id="editTitle"
          class="form-control"
          v-model="editArticle.title"
          required
        />
      </div>
      <div class="mb-3">
        <label for="editLabel" class="form-label">Label</label>
        <input
          type="text"
          id="editLabel"
          class="form-control"
          v-model="editArticle.label"
          required
        />
      </div>
      <div class="mb-3">
        <label for="editContent" class="form-label">Konten Artikel</label>
        <QuillEditor v-model="editArticle.content" />
      </div>
      <div class="mb-3">
        <label for="editCategory" class="form-label">Kategori</label>
        <select
          id="editCategory"
          class="form-select"
          v-model="editArticle.category_id"
          required
        >
          <option value="" disabled>Pilih kategori</option>
          <option v-for="category in categories" :key="category.ID" :value="category.ID">
            {{ category.name }}
          </option>
        </select>
      </div>
      <!-- Input untuk upload gambar -->
      <div class="mb-3">
        <label for="editImage" class="form-label">Gambar Artikel</label>
        <input
          type="file"
          id="editImage"
          class="form-control"
          @change="handleImageChange"
        />
      </div>
      <button type="submit" class="btn btn-primary me-2">Update Artikel</button>
      <NuxtLink to="/author/articles" class="btn btn-danger">Batal</NuxtLink>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router';  // Pastikan mengimpor useRouter
import QuillEditor from '@/components/QuillEditor.vue';

definePageMeta({
  middleware: 'auth',
  requiresAuthor: true,
});

const editArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author: '',
  file: null, // Tambahkan properti file untuk menyimpan file gambar
});
const categories = ref([]);
const route = useRoute();
const router = useRouter(); // Inisialisasi router untuk navigasi
const articleId = route.params.id; // Assuming the article ID is passed as a URL param
const alertMessage = ref('')
const alertClass = ref('')

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  return token;
};

const fetchArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }
    const response = await axios.get(`http://localhost:8080/author/articles/${articleId}`,
    {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    const articleData = response.data.data;
    editArticle.value = {
      label: articleData.label,
      title: articleData.title,
      content: articleData.content,
      category_id: articleData.category_id,
      author: articleData.author,
      file: null, // Reset file karena ini hanya untuk upload baru
    };
  } catch (error) {
    console.error('Error fetching article:', error);
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

// Fungsi untuk menangani perubahan file gambar
const handleImageChange = (event) => {
  const file = event.target.files[0];
  if (file) {
    editArticle.value.file = file; // Simpan file gambar yang dipilih
  }
};

const updateArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Artikel yang akan di-update:', editArticle.value);

    // Pertama, update artikel
    const updatedArticle = {
      label: editArticle.value.label,
      title: editArticle.value.title,
      content: editArticle.value.content,
      category_id: editArticle.value.category_id,
      author: editArticle.value.author,
    };

    await axios.put(
      `http://localhost:8080/author/articles/${articleId}`,
      updatedArticle,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Jika ada file gambar yang diubah, upload gambar tersebut
    if (editArticle.value.file) {
      await uploadImage();
    }

    alertMessage.value = `Artikel berhasil dibuat.`
    alertClass.value = 'alert alert-success'
    setTimeout(() => {
      window.location.href ='/author/articles';
    }, 1500);
  } catch (error) {
    console.error('Error updating article:', error);
  }
};

// Fungsi untuk mengunggah gambar
const uploadImage = async () => {
  const formData = new FormData();
  formData.append('file', editArticle.value.file);

  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.post(
      `http://localhost:8080/author/articles/${articleId}/uploads`,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(response.data)
  } catch (error) {
    console.error('Error uploading image:', error);
    alertMessage.value = `Gagal mengunggah gambar.`
    alertClass.value = 'alert alert-danger'
  }
};

onMounted(() => {
  fetchCategories();
  fetchArticle();
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