<template>
    <div class="container-fluid">
      <div class="row">
        <!-- Sidebar -->
        <aside class="col-md-3 col-lg-2 d-flex flex-column bg-light vh-100 p-3">
          <h4 class="mb-4">Company</h4>
          <nav class="nav flex-column">
            <a class="nav-link" href="/admin">General</a>
            <a class="nav-link active text-primary" href="/users">Users</a>
            <a class="nav-link" href="#">Statistic</a>
            <a class="nav-link" href="#">Billing</a>
            <a class="nav-link mt-auto" href="#">Sign out</a>
          </nav>
        </aside>
  
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
                <th scope="col"></th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="users.length === 0">
                <td colspan="7" class="text-center">No users found</td>
              </tr>
              <tr v-for="user in users" :key="user.id">
                <td><input type="checkbox" /></td>
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
                <td>{{ user.role }}</td>
                <td>{{ user.created_at }}</td>
                <td>{{ user.update_at }}</td>
                <td><button class="btn btn-sm btn-outline-secondary">...</button></td>
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

const fetchUsers = async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await axios.get('http://localhost:8080/admin/users', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log('API Response:', response);
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

  