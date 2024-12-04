<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Artikel</h1>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search" />
          <select class="form-select me-2">
            <option v-for="category in category">{{category.name}}</option>
          </select>
        </div>

        <!-- Tombol untuk menambah artikel baru -->
        <button
          class="btn btn-primary mb-4"
          data-bs-toggle="modal"
          data-bs-target="#articleModal"
          @click="openModal('add')"
        >
          Tambah Artikel Baru
        </button>

        <!-- Daftar artikel dalam bentuk card -->
        <div class="row">
          <div
            v-for="article in articles"
            :key="article.id"
            class="col-md-4 mb-4"
          >
            <div class="card">
              <img
                :src="`http://localhost:8080/uploads/${article.file_name}`"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h6 class="card-label">{{ article.label }}</h6>
                <h5 class="card-title">{{ article.title }}</h5>
                <p class="card-text">{{ article.content }}</p>
                <button
                  class="btn btn-warning"
                  data-bs-toggle="modal"
                  data-bs-target="#articleModal"
                  @click="openModal('edit', article)"
                >
                  Edit
                </button>
                <button class="btn btn-danger ms-2">Hapus</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Bootstrap -->
    <div
      class="modal fade"
      id="articleModal"
      tabindex="-1"
      aria-labelledby="articleModalLabel"
      aria-hidden="true"
    >
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="articleModalLabel">{{ modalTitle }}</h5>
            <button
              type="button"
              class="btn-close"
              data-bs-dismiss="modal"
              aria-label="Close"
            ></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="handleSubmit">
              <div class="mb-3">
                <label for="title" class="form-label">Judul</label>
                <input
                  type="text"
                  class="form-control"
                  id="title"
                  v-model="form.title"
                  required
                />
              </div>
              <div class="mb-3">
                <label for="label" class="form-label">Label</label>
                <input
                  type="text"
                  class="form-control"
                  id="label"
                  v-model="form.label"
                  required
                />
              </div>
              <div class="mb-3">
                <label for="content" class="form-label">Konten</label>
                <textarea
                  class="form-control"
                  id="content"
                  rows="4"
                  v-model="form.content"
                  required
                ></textarea>
              </div>
              <div class="mb-3">
                <label for="file" class="form-label">Upload File</label>
                <input
                  type="file"
                  class="form-control"
                  id="file"
                  @change="handleFileChange"
                />
                <small v-if="form.file_name">
                  File saat ini: {{ form.file_name }}
                </small>
              </div>
              <button type="submit" class="btn btn-primary">
                {{ modalAction }}
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const articles = ref([]);
const modalTitle = ref('');
const modalAction = ref('');
const form = ref({ title: '', label: '', content: '', id: null, file: null, file_name: '' });

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

const fetchArticles = async () => {
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Token yang diambil:', token);  // Debugging token
    const response = await axios.get('http://localhost:8080/articles', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

const openModal = (type, article = null) => {
  if (type === 'add') {
    modalTitle.value = 'Tambah Artikel Baru';
    modalAction.value = 'Tambah';
    form.value = { title: '', label: '', content: '', id: null, file: null, file_name: '' };
  } else if (type === 'edit' && article) {
    modalTitle.value = 'Edit Artikel';
    modalAction.value = 'Simpan';
    form.value = { ...article, file: null }; // Tidak menyertakan file asli dalam edit
  }
};

const handleFileChange = (event) => {
  form.value.file = event.target.files[0];
};

const handleSubmit = async () => {
  const token = getTokenFromCookies();

  try {
    let uploadedFileName = form.value.file_name; // Untuk menyimpan nama file yang sudah diunggah

    // 1. Upload file jika ada file baru
    if (form.value.file) {
      const fileData = new FormData();
      fileData.append("file", form.value.file);

      const uploadResponse = await axios.post(
        `http://localhost:8080/admin/articles/${form.value.id || 'temp'}/uploads`,
        fileData,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      uploadedFileName = uploadResponse.data.file_name; // Sesuaikan properti nama file dari respons backend Anda
    }

    // 2. Kirim data artikel
    const articleData = {
      title: form.value.title,
      label: form.value.label,
      content: form.value.content,
      file_name: uploadedFileName, // Masukkan nama file dari hasil upload
    };

    if (form.value.id) {
      // Update artikel
      await axios.put(
        `http://localhost:8080/articles/${form.value.id}`,
        articleData,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
    } else {
      // Tambah artikel baru
      await axios.post(`http://localhost:8080/articles`, articleData, {
        headers: { Authorization: `Bearer ${token}` },
      });
    }

    // Refresh data artikel setelah submit
    fetchArticles();

    // Tutup modal
    const modalElement = document.getElementById("articleModal");
    const modal = bootstrap.Modal.getInstance(modalElement);
    modal.hide();
  } catch (error) {
    console.error("Error submitting form:", error);
  }
};


const category = ref([]);

const fetchCategory = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category', {
  }); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  }
};

onMounted(() => {
  fetchArticles();
  fetchCategory();
});
</script>
