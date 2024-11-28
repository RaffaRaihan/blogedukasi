<template>
  <div class="container-fluid">
    <div class="row">
      <AdminSidebar />

      <!-- Main Content -->
      <main class="col-md-9 col-lg-10 p-4">
        <!-- Header -->
        <div class="d-flex justify-content-between align-items-center mb-4">
          <h1>Users</h1>
          <button class="btn btn-primary">+ Invite</button>
        </div>

        <!-- Filters -->
        <div class="d-flex align-items-center mb-4">
          <input type="text" class="form-control me-2" placeholder="Search" />
          <select class="form-select me-2">
            <option>All</option>
            <option>Admin</option>
            <option>User</option>
          </select>
        </div>

        <!-- User Table -->
        <table class="table table-hover align-middle">
          <thead>
            <tr>
              <th scope="col"><input type="checkbox" /></th>
              <th scope="col">Name</th>
              <th scope="col">Email</th>
              <th scope="col">Role</th>
              <th scope="col">CreatAt</th>
              <th scope="col">UpdateAt</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="users.length === 0">
              <td colspan="7" class="text-center">No users found</td>
            </tr>
            <tr v-for="(user, index) in users" :key="user.id">
              <td scope="row">{{ index + 1 }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td>{{ user.CreatedAt }}</td>
              <td>{{ user.UpdatedAt }}</td>
              <td>
                <NuxtLink class="btn btn-warning btn-sm me-2">Edit</NuxtLink>
                <button class="btn btn-danger btn-sm">Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

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
