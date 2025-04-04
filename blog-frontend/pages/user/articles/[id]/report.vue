<template>
  <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
  <div class="container mt-5">
    <h1 class="text-center mb-4">Laporkan Artikel</h1>
    <form @submit.prevent="ReportArticles">
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
      <p class="text-muted">*tolong cantumkan judul nya yaa ><</p>
      <div class="mb-3">
        <label for="message" class="form-label">Alasan</label>
        <QuillEditor v-model="formData.reason" required/>
      </div>
      <button type="submit" class="btn btn-primary">Kirim Pesan</button>
      <button type="button" class="btn btn-danger ms-2" @click="cancelReport">Batal</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router';
import useArticlesById from '~/composables/api/useArticlesById';

definePageMeta({
  middleware: 'auth',
});

const route = useRoute();
const router = useRouter();
const alertMessage = ref('');
const alertClass = ref('');

// Data untuk form
const formData = ref({
  article_id: '',
  name: '',
  reason: '',
});

// Ambil article_id dari parameter URL
onMounted(() => {
  formData.value.article_id = route.params.id;
});

const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  return token;
};

// Fungsi untuk mengirim pesan
const ReportArticles = async () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    // Kirim data menggunakan Axios
    const response = await axios.post('http://localhost:8080/user/report',
    formData.value, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    // Tampilkan respons dari server
    console.log('Response:', response.data);

    // Reset form jika berhasil
    formData.value.name = '';
    formData.value.reason = '';

    alertMessage.value = 'Pesan berhasil dikirim!';
    alertClass.value = 'alert alert-success';
    setTimeout(() => {
      window.location.href = '/user/message';
    }, 3000);
  } catch (error) {
    console.error('Error sending message:', error);
    alertMessage.value = 'Gagal mengirim pesan!';
    alertClass.value = 'alert alert-danger';
  }
};

const cancelReport = () => {
  router.push(`/user/articles/${route.params.id}`);
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
  