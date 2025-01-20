<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AuthorSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Kategori</h1>
        <!-- Tombol untuk menambah kategori baru -->
        <button class="btn btn-outline-primary" @click="showModal = true">+ Add Categories</button>

        <!-- Loading Indicator -->
        <Loading v-if="loadingCategory" />

        <!-- Tabel Kategori -->
        <table class="table table-hover align-middle mt-3" v-else>
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Nama Kategori</th>
              <th scope="col">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <!-- Loop untuk menampilkan kategori -->
            <tr v-for="(category, index) in category" :key="category.ID">
              <th scope="row">{{ index + 1 }}</th>
              <td>{{ category.name }}</td>
              <td>
                <button class="btn btn-outline-warning btn-sm me-2" @click="openModal('edit', category)"><i class="bi bi-pencil"></i> Edit</button>
                <button @click="removeCategory(category.ID)" class="btn btn-outline-danger btn-sm"><i class="bi bi-trash"></i> Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <!-- Modal Add Categoris -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showModal }" style="display: block;" v-if="showModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Add Categories</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addCategory">
              <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" v-model="newCategory.name" required />
              </div>
              <button type="submit" class="btn btn-outline-primary">+ Add Categories</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="modal-backdrop fade" :class="{ 'show': showModal }" v-if="showModal" @click="showModal = false"></div>
    <!-- Modal Edit Categoris -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showEditModal }" style="display: block;" v-if="showEditModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit Categories</h5>
            <button type="button" class="btn-close" @click="showEditModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateCategory">
              <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" v-model="editCategory.name" required />
              </div>
              <button type="submit" class="btn btn-primary">Save Changes</button>
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

const category = ref([]);
const showModal = ref(false);
const newCategory = ref({
  name: '',
})
const showEditModal = ref(false);
const editCategory = ref({
  name: '',
});
const loadingCategory = ref(true)

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchCategory = async () => {
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Token yang diambil:', token);  // Debugging token
    const response = await axios.get('http://localhost:8080/category', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  } finally {
    loadingCategory.value = false;
  }
};

// Fungsi untuk menambah admin
const addCategory = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.post(
      'http://localhost:8080/admin/category',
      newCategory.value,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Refresh daftar users dari API
    await fetchCategory(); // Tambahkan admin baru ke daftar
    showModal.value = false; // Tutup modal setelah berhasil
    newCategory.value = { name: ''}; // Reset form
  } catch (error) {
    console.error('Error adding admin:', error);
  }
};

const updateCategory = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    // Kirim permintaan ke backend untuk memperbarui kategori
    await axios.put(
      `http://localhost:8080/admin/category/${editCategory.value.ID}`, 
      { name: editCategory.value.name },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Perbarui daftar kategori
    await fetchCategory()
    // Tutup modal setelah berhasil
    showEditModal.value = false;
    console.log(`Category dengan ID ${editCategory.value.ID} berhasil diperbarui.`);
  } catch (error) {
    console.error('Error updating category:', error);
  }
};

// Fungsi untuk membuka modal edit
const openModal = (type, category) => {
  if (type === 'edit') {
    editCategory.value = { ...category };
    showEditModal.value = true;
  }
};

const removeCategory = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus pengguna ini?')) {
    return;
  }

  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.delete(`http://localhost:8080/admin/category/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    category.value = category.value.filter(category => category.ID !== id);
    console.log(`Categories dengan ID ${id} berhasil dihapus.`);
  } catch (error) {
    console.error('Error removing categories:', error);
  }
};

// Panggil fetchCategory saat komponen dimuat
onMounted(() => {
  fetchCategory();
});
</script>