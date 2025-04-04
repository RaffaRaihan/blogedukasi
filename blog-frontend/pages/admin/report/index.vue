<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Laporan</h1>
        <hr>
        <!-- Tabel Kategori -->
        <table class="table table-hover align-middle mt-3">
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">ID Artikel</th>
              <th scope="col">Nama</th>
              <th scope="col">Alasan</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(report, index) in report" :key="report.ID">
              <th scope="row">{{ index + 1 }}</th>
              <td>{{ report.article_id }}</td>
              <td>{{ report.name }}</td>
              <td v-html="report.reason"></td>
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

definePageMeta({
  middleware: 'auth',
  requiresAdmin: true,
});

const report = ref([]); // Data pesan

// Ambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fetch pesan dari API
const fetchReport = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const response = await axios.get('http://localhost:8080/admin/reports', {
      headers: { Authorization: `Bearer ${token}` },
    });
    report.value = response.data.data;
  } catch (error) {
    console.error('Error fetching messages:', error);
  }
};

// Ambil data pesan saat komponen dimuat
onMounted(() => {
  fetchReport();
});
</script>
  