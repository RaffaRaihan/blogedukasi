<template>
  <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
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
        <QuillEditor v-model="formData.content" required/>
        <!-- <textarea
          id="message"
          class="form-control"
          v-model="formData.content"
          placeholder="Tulis pesan Anda"
          rows="5"
          required
        ></textarea> -->
      </div>
      <button type="submit" class="btn btn-primary">Kirim Pesan</button>
      <NuxtLink to="/author/message" class="btn btn-danger ms-2">Batal</NuxtLink>
    </form>
  </div>
</template>
  
<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

definePageMeta({
  middleware: 'auth',
});

const alertMessage = ref('');
const alertClass = ref('');

// Data untuk form
const formData = ref({
  user_id: '',
  name: '',
  content: '',
});

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
    const response = await axios.post('http://localhost:8080/author/send-message',
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

    alertMessage.value = `Pesan berhasil dikirim!`;
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.href = "/author/message";
    }, 3000);
  } catch (error) {
    console.error('Error sending message:', error);
    alertMessage.value = `Gagal mengirim pesan!`;
    alertClass.value = 'alert alert-danger';
  }
};
</script>
<style scoped>
.alert {
  position: fixed;
  top: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 500px;
  z-index: 1050;
  text-align: center;
}
</style>