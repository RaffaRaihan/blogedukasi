<template>
  <div class="container-fluid">
    <div class="row">
      <AdminSidebar />
      <main class="col-md-9 col-lg-10 p-4">
        <div class="d-flex justify-content-between align-items-center mb-4">
          <h1>Users</h1>
          <button class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#articleModal" @click="openModal('add')">+ Invite</button>
        </div>
        <div class="d-flex align-items-center mb-4">
          <input type="text" class="form-control me-2" placeholder="Search" />
          <select class="form-select me-2">
            <option>All</option>
            <option>Admin</option>
            <option>User</option>
          </select>
        </div>
        <table class="table table-hover align-middle">
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Name</th>
              <th scope="col">Email</th>
              <th scope="col">Role</th>
              <th scope="col">CreatedAt</th>
              <th scope="col">UpdatedAt</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="users.length === 0">
              <td colspan="7" class="text-center">No users found</td>
            </tr>
            <tr v-for="(user, index) in users" :key="user.id">
              <td>{{ index + 1 }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td>{{ formatDate(user.CreatedAt) }}</td>
              <td>{{ formatDate(user.UpdatedAt) }}</td>
              <td>
                <button class="btn btn-warning btn-sm me-2" @click="openModal('edit', user)">Edit</button>
                <button class="btn btn-danger btn-sm">Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </main>
    </div>

    <!-- Modal -->
    <div class="modal fade" id="userModal" tabindex="-1" aria-labelledby="userModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="userModalLabel">{{ modalMode === 'add' ? 'Add User' : 'Edit User' }}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="handleSubmit">
              <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" id="name" v-model="formData.name" class="form-control" required />
              </div>
              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" id="email" v-model="formData.email" class="form-control" required />
              </div>
              <div class="mb-3">
                <label for="role" class="form-label">Role</label>
                <select id="role" v-model="formData.role" class="form-select" required>
                  <option value="Admin">Admin</option>
                  <option value="User">User</option>
                </select>
              </div>
              <button type="submit" class="btn btn-primary w-100">Save</button>
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
import { format } from 'date-fns';  // Import date-fns
import { id } from 'date-fns/locale';  // Import locale for Indonesian

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const users = ref([]);
const modalInstance = ref(null);
const modalMode = ref('add');
const formData = ref({
  id: null,
  name: '',
  email: '',
  role: 'User',
});

// Function to format the date
const formatDate = (date) => {
  return format(new Date(date), 'dd MMMM yyyy', { locale: id });
};

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

// Fungsi untuk membuka modal
const openModal = (mode, user = null) => {
  modalMode.value = mode;
  if (mode === 'edit' && user) {
    formData.value = { ...user };
  } else {
    formData.value = { id: null, name: '', email: '', role: 'User' };
  }
  if (!modalInstance.value) {
    const modalEl = document.getElementById('userModal');
    modalInstance.value = new bootstrap.Modal(modalEl);
  }
  modalInstance.value.show();
};

// Fungsi untuk submit data
const handleSubmit = async () => {
  try {
    if (modalMode.value === 'add') {
      await axios.post('http://localhost:8080/admin/users', formData.value, {
        headers: { Authorization: `Bearer ${getTokenFromCookies()}` },
      });
    } else if (modalMode.value === 'edit') {
      await axios.put(`http://localhost:8080/admin/users/${formData.value.id}`, formData.value, {
        headers: { Authorization: `Bearer ${getTokenFromCookies()}` },
      });
    }
    fetchUsers(); // Refresh data setelah submit
    modalInstance.value.hide();
  } catch (error) {
    console.error('Error submitting data:', error);
  }
};

// Panggil fetchUsers saat komponen dimuat
onMounted(() => {
  fetchUsers();
});
</script>
