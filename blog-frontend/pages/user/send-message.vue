<template>
  <div class="container mt-5">
    <h1 class="text-center mb-4">Kirim Pesan ke Admin</h1>
    <form @submit.prevent="sendMessage">
      <div class="mb-3">
        <label for="name" class="form-label">Nama</label>
        <input
          type="text"
          id="name"
          class="form-control"
          v-model="formData.name"
          placeholder="Masukkan nama Anda"
          required
        />
      </div>
      <div class="mb-3">
        <label for="message" class="form-label">Pesan</label>
        <textarea
          id="message"
          class="form-control"
          v-model="formData.content"
          placeholder="Tulis pesan Anda"
          rows="5"
          required
        ></textarea>
      </div>
      <button type="submit" class="btn btn-primary w-100">Kirim Pesan</button>
    </form>
  </div>
</template>
  
<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { jwtDecode } from 'jwt-decode';

definePageMeta({
  middleware: 'auth',
});

// Data untuk form
const formData = ref({
  user_id: '',
  name: '',
  content: '',
});
const router = useRouter();

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fungsi untuk mengirim pesan
const sendMessage = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error("Token tidak ditemukan. Harap login terlebih dahulu.");

    // Dekode token untuk mendapatkan user_id
    const decodedToken = jwtDecode(token);
    const userId = decodedToken.user_id; // Sesuaikan dengan struktur payload token Anda

    // Tambahkan sender_id ke formData
    formData.value.user_id = userId;

    // Kirim data menggunakan Axios
    const response = await axios.post('http://localhost:8080/user/send-message',
      formData.value, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Tampilkan respons dari server
    console.log('Response:', response.data);

    // Reset form jika berhasil
    formData.value = {
      name: '',
      content: '',
    };

    alert('Pesan berhasil dikirim!');
    router.push('/user/dashboard');
  } catch (error) {
    console.error('Error sending message:', error);
    alert('Gagal mengirim pesan. Periksa koneksi atau coba lagi.');
  }
};
</script>
