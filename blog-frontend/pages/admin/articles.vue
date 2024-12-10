<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Artikel</h1>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <!-- Search by Title -->
          <input 
            type="text" 
            class="form-control me-2" 
            placeholder="Search by Title" 
            v-model="searchQuery"
            @input="filterArticles"
          />
          <!-- Kategori Filter -->
          <select 
            class="form-select me-2" 
            v-model="selectedCategory" 
            @change="filterArticles"
          >
            <option value="">Semua Kategori</option>
            <option v-for="category in categories" :key="category.id" :value="category.ID">
              {{ category.name }}
            </option>
          </select>
        </div>
        <button class="btn btn-outline-primary mb-4" @click="showModal = true">Tambah Artikel</button>

        <!-- Daftar Artikel -->
        <div class="row">
          <div
            v-for="article in filteredArticles"
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
                <h5 class="card-title">{{ article.title }}</h5>
                <p class="card-text">{{ article.content.substring(0, 100)}}...</p>
                <p class="text-muted">By {{ article.author }}</p>
                <button class="btn btn-outline-warning" @click="openModal('edit', article)">
                  Edit
                </button>
                <button @click="removeArticles(article.ID)"  class="btn btn-outline-danger ms-2">Hapus</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal Tambah Artikel -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showModal }" style="display: block;" v-if="showModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Tambah Artikel Baru</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="submitArticle">
              <!-- Label -->
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
              <!-- Judul Artikel -->
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
              <!-- Konten Artikel -->
              <div class="mb-3">
                <label for="articleContent" class="form-label">Konten Artikel</label>
                <QuillEditor v-model="newArticle.content"/>
              </div>
              <!-- Kategori Artikel -->
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
              <!-- Author -->
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
              <button type="submit" class="btn btn-primary">Tambah Artikel</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="modal-backdrop fade" :class="{ 'show': showModal }" v-if="showModal" @click="showModal = false"></div>

    <!-- Modal Edit Artikel -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showEditModal }" style="display: block;" v-if="showEditModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit Artikel</h5>
            <button type="button" class="btn-close" @click="showEditModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateArticle">
              <!-- Label -->
              <div class="mb-3">
                <label for="editLabel" class="form-label">Label</label>
                <input
                  type="text"
                  id="editLabel"
                  class="form-control"
                  v-model="editArticle.label"
                  required
                  placeholder="Masukkan label"
                />
              </div>
              <!-- Judul Artikel -->
              <div class="mb-3">
                <label for="editTitle" class="form-label">Judul Artikel</label>
                <input
                  type="text"
                  id="editTitle"
                  class="form-control"
                  v-model="editArticle.title"
                  required
                  placeholder="Masukkan judul artikel"
                />
              </div>
              <!-- Konten Artikel -->
              <div class="mb-3">
                <label for="editContent" class="form-label">Konten Artikel</label>
                <QuillEditor v-model="editArticle.content" />
              </div>
              <!-- Kategori Artikel -->
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
              <!-- Author -->
              <div class="mb-3">
                <label for="editAuthor" class="form-label">Author</label>
                <input
                  type="text"
                  id="editAuthor"
                  class="form-control"
                  v-model="editArticle.author"
                  required
                  placeholder="Masukkan author"
                />
              </div>
              <!-- Upload File -->
              <div class="mb-3">
                <label for="editFoto" class="form-label">Photo</label>
                <input type="file" class="form-control" id="editFoto" @change="handleFileUpload" />
              </div>
              <button type="submit" class="btn btn-primary">Update Artikel</button>
            </form>
          </div>
        </div>
      </div>
    </div>

    <div class="modal-backdrop fade" :class="{ 'show': showEditModal }" v-if="editModal" @click="showEditModal = false"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const articles = ref([]);
const categories = ref([]);
const selectedCategory = ref(null); // Kategori yang dipilih
const searchQuery = ref(''); // Pencarian judul artikel
const newArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author: '',
});
const showModal = ref(false);
const errorMessage = ref('');
const showEditModal = ref(false);
const editArticle = ref({
  label: '',
  title: '',
  content: '',
  category_id: null,
  author: '',
});
const selectedFile = ref(null);

// Mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  return token;
};

// Fungsi untuk mengambil artikel
const fetchArticles = async (categoryId = null, search = '') => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.get('http://localhost:8080/articles', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params: {
        categoryId,
        search,
      },
    });
    articles.value = response.data;
  } catch (error) {
    console.error('Error fetching articles:', error);
  }
};

// Fungsi untuk mengambil kategori
const fetchCategories = async () => {
  try {
    const response = await axios.get('http://localhost:8080/category');
    categories.value = response.data;
  } catch (error) {
    console.error('Error fetching categories:', error);
  }
};

const removeArticles = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus artikel ini?')) {
    return;
  }

  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.delete(`http://localhost:8080/admin/articles/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    articles.value = articles.value.filter(articles => articles.ID !== id);
    console.log(`artikel dengan ID ${id} berhasil dihapus.`);
  } catch (error) {
    console.error('Error removing user:', error);
  }
};

// Fungsi untuk menambah artikel
const submitArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.post(
      'http://localhost:8080/admin/articles',
      newArticle.value,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Refresh daftar users dari API
    await fetchArticles(); // Tambahkan admin baru ke daftar
    showModal.value = false; // Tutup modal setelah berhasil
    newArticle.value = { label: '', title: '', content: '', categoryID: null, author: '' }; // Reset form
  } catch (error) {
    console.error('Error adding articles:', error);
    errorMessage.value = 'Terjadi kesalahan saat menambahkan artikel. Silakan coba lagi.';
  }
};

// Fungsi untuk membuka modal edit
const openModal = (type, article) => {
  if (type === 'edit') {
    editArticle.value = { ...article };
    showEditModal.value = true;
  }
};

// Fungsi untuk menangani file yang diunggah
const handleFileUpload = (event) => {
  selectedFile.value = event.target.files[0]; // Ambil file pertama dari input
};

const updateArticle = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const formData = new FormData();
    formData.append('label', editArticle.value.label);
    formData.append('title', editArticle.value.title);
    formData.append('content', editArticle.value.content);
    formData.append('category_id', editArticle.value.category_id);
    formData.append('author', editArticle.value.author);

    // Cek apakah ada foto yang diunggah
    if (selectedFile.value) {
      formData.append('file', selectedFile.value); // Menambahkan file foto ke formData
    }

    // Kirim ke backend (untuk update profil pengguna)
    await axios.put(
      `http://localhost:8080/admin/articles/${editArticle.value.ID}/uploads`,
      formData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      }
    );

    // Segarkan daftar artikel
    await fetchArticles();
    showEditModal.value = false; // Tutup modal setelah berhasil
    console.log('Artikel berhasil diperbarui.');
  } catch (error) {
    console.error('Error updating artikel:', error);
  }
};

// Computed untuk memfilter artikel berdasarkan kategori dan pencarian judul
const filteredArticles = computed(() => {
  return articles.value.filter(article => {
    const matchesCategory = selectedCategory.value ? article.category.ID === selectedCategory.value : true;
    const matchesSearch = article.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesCategory && matchesSearch;
  });
});

// Fungsi untuk memfilter artikel saat kategori atau pencarian berubah
const filterArticles = () => {
  fetchArticles(selectedCategory.value, searchQuery.value);
};

onMounted(() => {
  fetchArticles(); // Ambil artikel saat komponen dimuat
  fetchCategories(); // Ambil kategori
});
</script>
