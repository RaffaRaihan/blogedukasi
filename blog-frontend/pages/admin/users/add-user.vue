<template>
  <div class="container mt-5">
    <h2>Add New User</h2>
    <form @submit.prevent="addUser">
      <div class="mb-3">
        <label for="name" class="form-label">Name</label>
        <input type="text" class="form-control" id="name" v-model="newUser.name" required />
      </div>
      <div class="mb-3">
        <label for="email" class="form-label">Email</label>
        <input type="email" class="form-control" id="email" v-model="newUser.email" required />
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">Password</label>
        <input type="password" class="form-control" id="password" v-model="newUser.password" required />
      </div>
      <div class="mb-3">
        <label for="role" class="form-label">Role</label>
        <select class="form-select" id="role" v-model="newUser.role" required>
          <option value="user">User</option>
          <option value="admin">Admin</option>
        </select>
      </div>
      <div class="mb-3">
        <label for="profileImage" class="form-label">Foto Profile</label>
        <input
          type="file"
          id="profileImage"
          class="form-control"
          ref="fileInput"
        />
      </div>
      <button type="submit" class="btn btn-primary">Add User</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

const newUser = ref({
  name: '',
  email: '',
  password: '',
  role: 'user',
});

const router = useRouter();
const newProfileId = ref(null); // Untuk menyimpan ID artikel yang baru dibuat

const addUser = async () => {
  try {
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    const response = await axios.post('http://localhost:8080/register', 
    newUser.value, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log(response.data)

    // Dapatkan ID artikel yang baru dibuat
    const userId = response.data.data.ID;

    // Simpan ID artikel untuk digunakan saat upload gambar
    if (userId) {
      newProfileId.value = userId; // Simpan ID artikel
    }

    // Jika ada file gambar, unggah
    const fileInput = document.getElementById('profileImage');
    const file = fileInput?.files[0];
    if (file) {
      await handleImageUpload(file, newProfileId.value);
    }
    
    if (role == "admin"){
      alert('Berhail menambah admin');  
    } else {
      alert('Berhail menambah user');
    }
    router.push('/admin/users'); // Kembali ke halaman utama setelah berhasil
  } catch (error) {
    console.error('Error adding user:', error);
  }
};

const handleImageUpload = async (file, userId) => {
  const formData = new FormData();
  formData.append('foto', file);

  try {
    // Ambil token dari cookie atau localStorage
    const token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1];

    if (!token) {
      alert('Token tidak ditemukan. Pastikan Anda sudah login.');
      return;
    }

    const response = await axios.post(
      `http://localhost:8080/admin/users/${userId}/foto`,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log('File uploaded:', response.data);
    alert('Berhasil Menambah Foto Profile');
  } catch (error) {
    console.error('Error uploading file:', error);
    alert('Gagal mengunggah gambar');
  }
};
</script>
  