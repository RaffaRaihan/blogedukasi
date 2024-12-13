<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <!-- Main Content -->
      <main class="p-4 col-md-12 col-md-9 col-lg-10">
        <h2>Workspace admins</h2>

        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search by email" v-model="searchQuery" />
          <select class="form-select me-2" v-model="selectedRole">
            <option value="">All</option>
            <option value="admin">Admin</option>
            <option value="user">User</option>
          </select>
        </div>

        <!-- Loading Indicators -->
        <Loading v-if="loadingUsers" />

        <!-- Admin Cards -->
        <div class="mb-4" v-else>
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
const loadingUsers = ref(true);

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  
  return token;
};

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
  } finally {
    loadingUsers.value = false
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
  } catch (error) {
    console.error('Error removing user:', error);
  }
};

const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesRole = selectedRole.value ? user.role === selectedRole.value : true;
    const matchesSearch = user.email && user.email.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesRole && matchesSearch;
  });
});

onMounted(() => {
  fetchUsers();
});
</script>

<style scoped>
main {
  transition: all 0.3s ease;
}
</style>
