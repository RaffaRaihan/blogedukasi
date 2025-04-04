<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Kelola Kategori</h1>
        <hr>
        <!-- Tombol untuk menambah kategori baru -->
        <button class="btn btn-outline-primary" data-bs-toggle="modal" data-bs-target="#exampleModal">+ Tambah Kategori</button>

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
            <tr v-for="(category, index) in category" :key="category.ID">
              <th scope="row">{{ index + 1 }}</th>
              <td>{{ category.name }}</td>
              <td>
                <!-- Set data sebelum menampilkan modal -->
                <button 
                  class="btn btn-outline-warning me-2" 
                  data-bs-toggle="modal" 
                  data-bs-target="#editModal"
                  @click="setEditCategory(category)"
                >
                  <i class="bi bi-pencil"></i> Edit
                </button>
                <button @click="removeCategory(category.ID)" class="btn btn-outline-danger">
                  <i class="bi bi-trash"></i> Hapus
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal Add Category -->
    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Tambah Kategori</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addCategory">
              <div class="mb-3">
                <label for="name" class="form-label">Nama</label>
                <input type="text" class="form-control" id="name" v-model="newCategory.name" required />
              </div>
              <button type="submit" class="btn btn-outline-primary" data-bs-dismiss="modal">Tambah Kategori</button>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Edit Category -->
    <div class="modal fade" id="editModal" tabindex="-1" aria-labelledby="editModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit Kategori</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateCategory">
              <div class="mb-3">
                <label for="edit-name" class="form-label">Nama</label>
                <input type="text" class="form-control" id="edit-name" v-model="editCategory.name" required />
              </div>
              <button type="submit" class="btn btn-primary" data-bs-dismiss="modal">Simpan Perubahan</button>
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
const newCategory = ref({ name: '' });
const editCategory = ref({ ID: null, name: '' });
const loadingCategory = ref(true);

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
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.get('http://localhost:8080/category', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  } finally {
    loadingCategory.value = false;
  }
};

// Set data kategori yang akan diedit
const setEditCategory = (data) => {
  editCategory.value = { ...data }; // Menyalin data agar tidak mengubah data asli langsung
};

// Fungsi untuk menambah kategori
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

    await fetchCategory(); 
    newCategory.value = { name: '' };
  } catch (error) {
    console.error('Error adding category:', error);
  }
};

// Fungsi untuk memperbarui kategori
const updateCategory = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.put(
      `http://localhost:8080/admin/category/${editCategory.value.ID}`, 
      { name: editCategory.value.name },
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    await fetchCategory();
    console.log(`Kategori dengan ID ${editCategory.value.ID} berhasil diperbarui.`);
  } catch (error) {
    console.error('Error updating category:', error);
  }
};

// Fungsi untuk menghapus kategori
const removeCategory = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus kategori ini?')) {
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

    category.value = category.value.filter(cat => cat.ID !== id);
    console.log(`Kategori dengan ID ${id} berhasil dihapus.`);
  } catch (error) {
    console.error('Error removing category:', error);
  }
};

// Panggil fetchCategory saat komponen dimuat
onMounted(() => {
  fetchCategory();
});
</script>
