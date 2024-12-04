<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />

      <!-- Main Content -->
      <main class="col-md-9 col-lg-10 p-4">
        <h2>Workspace admins</h2>
        <button class="btn btn-outline-primary">+ Add Admin</button>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search" />
          <select class="form-select me-2">
            <option>All</option>
          </select>
        </div>
        <!-- Admin Cards -->
        <div class="mb-4">
          <div class="row d-flex">
            <div v-for="users in users" class="card me-2 mb-2" style="width: 150px;">
              <div class="card-body text-center">
                <h6 class="card-title">{{ users.name }}</h6>
                <p class="text-muted">{{ users.role }}</p>
                <button class="btn btn-sm btn-danger">Remove</button>
              </div>
            </div>
          </div>
        </div>
      </main>
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

const users = ref([]);

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

// Fungsi untuk fetch users dari API
const fetchUsers = async () => {
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Token yang diambil:', token);  // Debugging token

    // Panggil API dengan token
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

// Panggil fetchUsers saat komponen dimuat
onMounted(() => {
  fetchUsers();
});
</script>

  