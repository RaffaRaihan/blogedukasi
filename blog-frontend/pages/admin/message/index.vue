<template>
  <div class="container-fluid">
    <div class="row">
      <!-- Sidebar -->
      <AdminSidebar />
      <div class="col-md-9 col-lg-10 p-4">
        <h1 class="mb-4">Halaman Admin - Report</h1>
        <!-- Tabel Kategori -->
        <table class="table table-hover align-middle mt-3">
          <thead>
            <tr>
              <th scope="col">No</th>
              <th scope="col">Sender_Id</th>
              <th scope="col">Nama</th>
              <th scope="col">Pesan</th>
              <th scope="col">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(message, index) in message" :key="message.ID">
              <th scope="row">{{ index + 1 }}</th>
              <td>{{ message.sender_id }}</td>
              <td>{{ message.name }}</td>
              <td>{{ message.content }}</td>
              <td>
                <button class="btn btn-outline-primary" data-bs-toggle="modal" data-bs-target="#exampleModal">Balas</button>
              </td>
            </tr>
          </tbody>
        </table>
        <!-- Modal -->
        <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <h1 class="modal-title fs-5" id="exampleModalLabel">Balas Pesan</h1>
                <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
              </div>
              <div class="modal-body">
                <QuillEditor />
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-outline-danger" data-bs-dismiss="modal">Batal</button>
                <button type="button" class="btn btn-outline-primary">Kirim</button>
              </div>
            </div>
          </div>
        </div>
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

const message = ref([]);

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchMessage = async () => {
  try {
    // Ambil token dari cookies
    const token = getTokenFromCookies();

    if (!token) {
      throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');
    }

    console.log('Token yang diambil:', token);  // Debugging token
    const response = await axios.get('http://localhost:8080/admin/messages', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }); // Ganti dengan URL backend Anda
    message.value = response.data.data;
  } catch (error) {
    console.error('Error fetching category:', error);
  }
};

onMounted(() => {
  fetchMessage();
});
</script>