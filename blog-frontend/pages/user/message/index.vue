<template>
  <LoginNavbar v-if="isLoggedIn"/>
  <Navbar v-else/>
  <div class="container mt-4">
    <h1 class="mb-4">Pesan yang Diterima</h1>
    <table class="table table-hover align-middle mt-3">
      <thead>
        <tr>
          <th scope="col">No</th>
          <th scope="col">Pesan Anda</th>
          <th scope="col">Pesan Admin</th>
          <th scope="col">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <!-- Loop untuk menampilkan kategori -->
        <tr v-for="(messages, index) in messages" :key="messages.ID">
          <th scope="row">{{ index + 1 }}</th>
          <td>{{ messages.content }}</td>
          <td>{{ messages.reply }}</td>
          <td>
            <i 
              v-if="messages.reply" 
              class="bi bi-check-circle-fill text-success"
              title="Sudah dibalas">
            </i>
            <p 
              v-else 
              class="text-muted">
              Belum Dibalas
            </p>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  <Footer />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

definePageMeta({
  middleware: 'auth',
});

// Data pesan
const messages = ref([]);
const isLoggedIn = computed(() => {
  const token = getTokenFromCookies();
  return !!token; // Return true jika token ada, false jika tidak ada
});

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))?.split('=')[1];
  
  return token;
};

const fetchMessage = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    // Decode token untuk mendapatkan userId
    const decoded = jwtDecode(token);
    const userId = decoded.user_id; // Pastikan nama field sesuai dengan isi payload token Anda

    const response = await axios.get(`http://localhost:8080/user/messages/users/${userId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log(response.data)
    messages.value = response.data.data;
  } catch (error) {
    console.error('Error fetching messages:', error);
  }
};

// Ambil data pesan saat komponen dimuat
onMounted(() => {
  fetchMessage();
});
</script>

