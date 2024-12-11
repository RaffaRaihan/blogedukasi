<template>
  <div class="container-fluid">
    <div class="row">
      <AdminSidebar />
      <main class="col-md-9 col-lg-10 p-4">
        <div class="d-flex justify-content-between align-items-center mb-4">
          <h1 class="fw-bold">Users</h1>
          <router-link to="users/add-user" class="btn btn-outline-primary">
            <i class="bi bi-person-plus"></i> Invite
          </router-link>
        </div>
        <!-- Filters -->
        <div class="d-flex align-items-center mb-4 mt-4">
          <input type="text" class="form-control me-2" placeholder="Search by email" v-model="searchQuery" />
          <select class="form-select me-2" v-model="selectedRole" @change="filteredUsers">
            <option value="">All</option>
            <option value="admin">Admin</option>
            <option value="user">User</option>
          </select>
        </div>

        <!-- Loading Indicators -->
        <Loading v-if="loadingUsers" />

        <table class="table table-hover align-middle" v-else>
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Name</th>
              <th scope="col">Email</th>
              <th scope="col">Role</th>
              <th scope="col">Waktu Pembuatan</th>
              <th scope="col">Update Terakhir</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="users.length === 0">
              <td colspan="7" class="text-center">No users found</td>
            </tr>
            <tr v-for="(user, index) in filteredUsers" :key="user.ID">
              <td>{{ index + 1 }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td>{{ formatDate(user.CreatedAt) }}</td>
              <td>{{ formatDate(user.UpdatedAt) }}</td>
              <td>
                <router-link :to="`users/edit/${user.ID}`" class="btn btn-outline-warning btn-sm me-2"><i class="bi bi-pencil"></i></router-link>
                <button @click="removeUser(user.ID)" class="btn btn-sm btn-outline-danger"><i class="bi bi-trash"></i></button>
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
import { useRouter } from 'vue-router';
import axios from 'axios';
import { format } from 'date-fns';
import { id } from 'date-fns/locale';

const users = ref([]);
const searchQuery = ref('');
const selectedRole = ref('');
const loadingUsers = ref(true);

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
