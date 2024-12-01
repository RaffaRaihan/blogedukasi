<template>
    <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
    <div class="col-md-9 col-lg-10 p-4">
      <h1 class="mb-4">Halaman Admin - Kelola Kategori</h1>
  
      <!-- Tombol untuk menambah kategori baru -->
      <NuxtLink to="/tambah-category" class="btn btn-primary mb-4">Tambah Kategori Baru</NuxtLink>
  
      <!-- Tabel Kategori -->
      <table class="table table-striped">
        <thead>
          <tr>
            <th scope="col">No</th>
            <th scope="col">Nama Kategori</th>
            <th scope="col">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <!-- Loop untuk menampilkan kategori -->
          <tr v-for="(kategori, index) in category" :key="kategori.id">
            <th scope="row">{{ index + 1 }}</th>
            <td>{{ kategori.name }}</td>
            <td>
              <NuxtLink class="btn btn-warning btn-sm me-2">Edit</NuxtLink>
              <button class="btn btn-danger btn-sm">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    </div>
    </div>
  </template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const category = ref([]);
// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchCategory = async () => {
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();
    const response = await axios.get('http://localhost:8080/category', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }); // Ganti dengan URL backend Anda
    category.value = response.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  }
};

// Panggil fetchCategory saat komponen dimuat
onMounted(() => {
  fetchCategory();
});
</script>