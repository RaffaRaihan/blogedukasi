<template>
  <div class="container-fluid">
    <div class="row">
      <AdminSidebar />
      <main class="col-md-9 col-lg-10 p-4">
        <div class="d-flex justify-content-between align-items-center mb-4">
          <h1>Users</h1>
          <button class="btn btn-outline-primary" data-bs-toggle="modal" data-bs-target="#articleModal" @click="showModal = true">+ Invite</button>
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
        <table class="table table-hover align-middle table-striped">
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
            <tr v-for="(user, index) in filteredUsers" :key="user.id">
              <td>{{ index + 1 }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td>{{ formatDate(user.CreatedAt) }}</td>
              <td>{{ formatDate(user.UpdatedAt) }}</td>
              <td>
                <button class="btn btn-outline-warning btn-sm me-2" @click="openModal('edit', user)">Edit</button>
                <button @click="removeUser(user.ID)" class="btn btn-sm btn-outline-danger">Hapus</button>
              </td>
            </tr>
          </tbody>
        </table>
      </main>
    </div>

    <!-- Modal Add Admin -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showModal }" style="display: block;" v-if="showModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Add Users</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="addUsers">
              <div class="mb-3">
                <label for="name" class="form-label">Name</label>
                <input type="text" class="form-control" id="name" v-model="newUsers.name" required />
              </div>
              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" class="form-control" id="email" v-model="newUsers.email" required />
              </div>
              <div class="mb-3">
                <label for="password" class="form-label">Password</label>
                <input type="password" class="form-control" id="password" v-model="newUsers.password" required />
              </div>
              <div class="mb-3">
                <label for="role" class="form-label">Role</label>
                <select class="form-select" id="role" v-model="newUsers.role" required>
                  <option value="user">User</option>
                  <option value="admin">Admin</option>
                </select>
              </div>
              <button type="submit" class="btn btn-primary">Add Users</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <div class="modal-backdrop fade" :class="{ 'show': showModal }" v-if="showModal" @click="showModal = false"></div>
    <!-- Modal Edit -->
    <div class="modal fade" tabindex="-1" :class="{ 'show': showEditModal }" style="display: block;" v-if="showEditModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit User</h5>
            <button type="button" class="btn-close" @click="showEditModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateUser">
              <div class="mb-3">
                <label for="editName" class="form-label">Name</label>
                <input type="text" class="form-control" id="editName" v-model="editUser.name" required />
              </div>
              <div class="mb-3">
                <label for="editEmail" class="form-label">Email</label>
                <input type="email" class="form-control" id="editEmail" v-model="editUser.email" required />
              </div>
              <div class="mb-3">
                <label for="editRole" class="form-label">Role</label>
                <select class="form-select" id="editRole" v-model="editUser.role" required>
                  <option value="user">User</option>
                  <option value="admin">Admin</option>
                </select>
              </div>
              <div class="mb-3">
                <label for="editFoto" class="form-label">Profile Photo</label>
                <input type="file" class="form-control" id="editFoto" @change="handleFileUpload" />
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
import { format } from 'date-fns';  // Import date-fns
import { id } from 'date-fns/locale';  // Import locale for Indonesian

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const users = ref([]);
const searchQuery = ref('');
const selectedRole = ref('');
const showModal = ref(false);
const newUsers = ref({
  name: '',
  email: '',
  password: '',
  role: 'user',
});
const showEditModal = ref(false);
const editUser = ref({
  id: null,
  name: '',
  email: '',
  role: '',
});
const selectedFile = ref(null);

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

// Fungsi untuk menambah admin
const addUsers = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    await axios.post(
      'http://localhost:8080/register',
      newUsers.value,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Refresh daftar users dari API
    await fetchUsers(); // Tambahkan admin baru ke daftar
    showModal.value = false; // Tutup modal setelah berhasil
    newUsers.value = { name: '', email: '', password: '', role: 'user' }; // Reset form
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

// Fungsi untuk menangani file yang diunggah
const handleFileUpload = (event) => {
  selectedFile.value = event.target.files[0]; // Ambil file pertama dari input
};

// Fungsi untuk membuka modal edit
const openModal = (type, user) => {
  if (type === 'edit') {
    editUser.value = { ...user };
    showEditModal.value = true;
  }
};

const updateUser = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    const formData = new FormData();
    formData.append('name', editUser.value.name);
    formData.append('email', editUser.value.email);
    formData.append('role', editUser.value.role);

    // Cek apakah ada foto yang diunggah
    if (selectedFile.value) {
      formData.append('foto', selectedFile.value); // Menambahkan file foto ke formData
    }

    // Kirim ke backend (untuk update profil pengguna)
    await axios.put(
      `http://localhost:8080/admin/users/${editUser.value.ID}/foto`,
      formData,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'multipart/form-data',
        },
      }
    );

    // Segarkan daftar pengguna
    await fetchUsers();
    showEditModal.value = false; // Tutup modal setelah berhasil
    console.log('User berhasil diperbarui.');
  } catch (error) {
    console.error('Error updating user:', error);
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
