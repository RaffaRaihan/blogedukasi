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
      <div class="mb-3">
        <label class="form-label">Status</label>
        <select class="form-select" v-model="editArticle.status">
          <option value="Sesuai">Sesuai</option>
          <option value="Belum Sesuai">Belum Sesuai</option>
        </select>
      </div>
      <div class="mb-3" v-if="showNoteField">
        <label for="editNote" class="form-label">Catatan</label>
        <QuillEditor v-model="editArticle.note" />
      </div>
      <!-- Input dan preview gambar artikel -->
      <div class="mb-3">
        <label for="editImage" class="form-label">Gambar Artikel</label>
        <input
          type="file"
          id="editImage"
          class="form-control"
          @change="handleImageChange"
        />
        <div v-if="imagePreview" class="mt-2">
          <img :src="imagePreview" alt="Preview Gambar" class="img-thumbnail" style="max-height: 200px;" />
        </div>
      </div>
      <button type="submit" class="btn btn-primary me-2">Update Artikel</button>
      <NuxtLink to="/admin/articles" class="btn btn-danger">Batal</NuxtLink>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import QuillEditor from '@/components/QuillEditor.vue';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const editArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author: '',
  status: '',
  note: '',
  file_name: null,
});

const categories = ref([]);
const route = useRoute();
const articleId = route.params.id;
const alertMessage = ref('');
const alertClass = ref('');
const imagePreview = ref('');

// Menentukan apakah field catatan ditampilkan
const showNoteField = computed(() => editArticle.value.status !== 'Sesuai');

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  return token;
};

const fetchArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const response = await axios.get(`http://localhost:8080/admin/articles/${articleId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    const articleData = response.data.data;
    editArticle.value = {
      label: articleData.label,
      title: articleData.title,
      content: articleData.content,
      category_id: articleData.category_id,
      author: articleData.author,
      status: articleData.status,
      note: articleData.note,
      file_name: null,
    };

    // Set preview gambar jika ada
    imagePreview.value = articleData.file_name || `http://localhost:8080/uploads/${articleData.file_name}`;
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

const handleImageChange = (event) => {
  const file = event.target.files[0];
  if (file) {
    editArticle.value.file_name = file;
    imagePreview.value = URL.createObjectURL(file);
  }
};

const updateArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      setTimeout(() => {
        window.location.href = "/login";
      }, 3000);
      return;
    }

    const updatedArticle = {
      label: editArticle.value.label,
      title: editArticle.value.title,
      content: editArticle.value.content,
      category_id: editArticle.value.category_id,
      author: editArticle.value.author,
      status: editArticle.value.status,
      note: editArticle.value.note,
    };

    await axios.put(
      `http://localhost:8080/admin/articles/${articleId}`,
      updatedArticle,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    if (editArticle.value.file) {
      await uploadImage();
    }

    alertMessage.value = `Artikel berhasil diubah.`;
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.href = "/admin/articles";
    }, 3000);
  } catch (error) {
    console.error('Error updating article:', error);
    alertMessage.value = `Gagal mengubah artikel`;
    alertClass.value = 'alert alert-danger';
    setTimeout(() => {
      window.location.href = "/admin/articles";
    }, 3000);
  }
};

const uploadImage = async () => {
  const formData = new FormData();
  formData.append('file', editArticle.value.file);

  try {
    const token = getTokenFromCookies();
    if (!token) {
      alertMessage.value = `Token tidak ditemukan. Pastikan Anda sudah login!!`;
      alertClass.value = 'alert alert-danger';
      setTimeout(() => {
        window.location.href = "/login";
      }, 3000);
      return;
    }

    await axios.post(
      `http://localhost:8080/admin/articles/${articleId}/uploads`,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: `Bearer ${token}`,
        },
      }
    );
  } catch (error) {
    console.error('Error uploading image:', error);
    alertMessage.value = `Gagal menggungah gambar`;
    alertClass.value = 'alert alert-danger';
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
