<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />

      <!-- Main Content -->
      <main class="col-md-9 col-lg-10 p-4">
        <h2>Workspace admins</h2>
        <button class="btn btn-outline-primary" @click="showModal = true">
          <i class="bi bi-person-plus"></i>  Add Admin
        </button>
        
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search by email" v-model="searchQuery" />
          <select class="form-select me-2" v-model="selectedRole">
            <option value="">All</option>
            <option value="admin">Admin</option>
            <option value="user">User</option>
          </select>
        </div>

        <!-- Admin Cards -->
        <div class="mb-4">
          <div class="row d-flex">
            <div v-for="user in filteredUsers" :key="user.id" class="card me-2 mb-2" style="width: 170px;">
              <div class="card-body text-center">
                <img
                  :src="`http://localhost:8080/uploads/${user.foto}`"
                  alt="Foto Profil"
                  class="rounded-circle mb-3"
                  style="width: 120px; height: 120px; object-fit: cover;"
                />
                <h6 class="card-title">{{ user.name }}</h6>
                <p class="text-muted">{{ user.email }}</p>
                <p class="text-muted">{{ user.role }}</p>
                <button @click="removeUser(user.ID)" class="btn btn-sm btn-danger"><i class="bi bi-person-dash"></i>  Remove</button>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- Modal Add Admin -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showModal }" style="display: block;" v-if="showModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Add Admin</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addAdmin">
              <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" v-model="newAdmin.name" required />
              </div>
              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" class="form-control" id="email" v-model="newAdmin.email" required />
              </div>
              <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" class="form-control" id="password" v-model="newAdmin.password" required />
              </div>
              <div class="mb-3">
                <label for="role" class="form-label">Role</label>
                <select class="form-select" id="role" v-model="newAdmin.role" required>
                  <option value="admin">Admin</option>
                  <option value="user">User</option>
                </select>
              </div>
              <button type="submit" class="btn btn-primary">Add Admin</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="modal-backdrop fade" :class="{ 'show': showModal }" v-if="showModal" @click="showModal = false"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import 'bootstrap-icons/font/bootstrap-icons.css';

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const users = ref([]);
const searchQuery = ref('');
const selectedRole = ref('');
const showModal = ref(false);
const newAdmin = ref({
  name: '',
  email: '',
  password: '',
  role: 'admin',
});

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  
  return token;
};

// Fungsi untuk fetch users dari API
const fetchUsers = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.get('http://localhost:8080/admin/users', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    users.value = response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
  }
};

// Fungsi untuk menambah admin
const addAdmin = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const response = await axios.post(
      'http://localhost:8080/register',
      newAdmin.value,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Refresh daftar users dari API
    await fetchUsers(); // Tambahkan admin baru ke daftar
    showModal.value = false; // Tutup modal setelah berhasil
    newAdmin.value = { name: '', email: '', password: '', role: 'admin' }; // Reset form
  } catch (error) {
    console.error('Error adding admin:', error);
    if (error.response && error.response.data.error === 'Email already exists') {
      errorMessage.value = 'Email sudah digunakan, silakan gunakan email lain.';
    } else {
      errorMessage.value = 'Terjadi kesalahan saat menambahkan admin. Silakan coba lagi.';
    }
  }
};

const removeUser = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus pengguna ini?')) {
    return;
  }

  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.delete(`http://localhost:8080/admin/users/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    users.value = users.value.filter(user => user.ID !== id);
    console.log(`User dengan ID ${id} berhasil dihapus.`);
  } catch (error) {
    console.error('Error removing user:', error);
  }
};

// Filter users berdasarkan role dan email
const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesRole = selectedRole.value ? user.role === selectedRole.value : true;
    const matchesSearch = user.email && user.email.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesRole && matchesSearch;
  });
});

// Panggil fetchUsers saat komponen dimuat
onMounted(() => {
  fetchUsers();
});
</script>
