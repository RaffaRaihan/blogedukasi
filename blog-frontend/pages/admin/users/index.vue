<template>
  <div class="container-fluid">
    <div class="row">
      <AdminSidebar />
      <main class="col-md-9 col-lg-10 p-4">
        <div class="d-flex justify-content-between align-items-center mb-4">
          <h1>Halaman Admin - Kelola Users</h1>
          <NuxtLink to="users/add-user" class="btn btn-outline-primary">
            <i class="bi bi-person-plus"></i> Undang
          </NuxtLink>
        </div>
        <hr>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Cari Berdasarkan Email" v-model="searchQuery" />
          <select class="form-select me-2" v-model="selectedRole" @change="filteredUsers">
            <option value="">Semua</option>
            <option value="admin">Admin</option>
            <option value="author">Author</option>
            <option value="user">User</option>
          </select>
        </div>

        <!-- Loading Indicators -->
        <Loading v-if="loadingUsers" />

        <table class="table table-hover align-middle" v-else>
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Foto</th>
              <th scope="col">Nama</th>
              <th scope="col">Email</th>
              <th scope="col">Role</th>
              <th scope="col">Waktu Pembuatan</th>
              <th scope="col">Update Terakhir</th>
              <th scope="col">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="users.length === 0">
              <td colspan="7" class="text-center">No users found</td>
            </tr>
            <tr v-for="(user, index) in sortedUsers" :key="user.ID">
              <td>{{ index + 1 }}</td>
              <td><img :src="`http://localhost:8080/uploads/${user.foto}`" alt="Profile" class="rounded img img-fluid" style="width: 5rem; height: 5rem;"/></td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td>{{ formatDate(user.CreatedAt) }}</td>
              <td>{{ formatDate(user.UpdatedAt) }}</td>
              <td>
                <NuxtLink :to="`users/edit/${user.ID}`" class="btn btn-outline-warning me-2"><i class="bi bi-pencil"></i>  Edit</NuxtLink>
                <button v-if="user.role !== 'admin'" @click="removeUser(user.ID)" class="btn btn-outline-danger">
                  <i class="bi bi-person-dash"></i> Hapus
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';
import id from 'date-fns/locale/id/index.js'

const searchQuery = ref('');
const selectedRole = ref('');
const loadingUsers = ref(true);

const users = ref([]);
const fetchUsers = async () => {
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];
    
    const response = await axios.get('http://localhost:8080/admin/users', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    
    users.value = response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
  } finally {
    loadingUsers.value = false;
  }
};

const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

const filteredUsers = computed(() => {
  return users.value.filter(user => {
    const matchesRole = selectedRole.value ? user.role === selectedRole.value : true;
    const matchesSearch = user.email && user.email.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesRole && matchesSearch;
  });
});

const sortedUsers = computed(() => {
  return [...filteredUsers.value].sort((a, b) => a.role.localeCompare(b.role));
});

const removeUser = async (id) => {
  if (!confirm('Apakah Anda yakin ingin menghapus pengguna ini?')) return;
  
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

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

onMounted(() => {
  fetchUsers();
});
</script>
