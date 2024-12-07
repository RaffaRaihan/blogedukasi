<template>
  <div class="container mt-5">
    <div class="card mx-auto" style="max-width: 600px;">
      <div class="card-header bg-success text-white">
        <h3 class="mb-0">Profil Pengguna</h3>
      </div>
      <div class="card-body">
        <div class="text-center">
          <img
            :src="users.foto ? `http://localhost:8080/uploads/${users.foto}` : 'default-avatar.png'"
            alt="Foto Profil"
            class="rounded-circle mb-3"
            style="width: 120px; height: 120px; object-fit: cover;"
          />
        </div>
        <div class="mb-3">
          <label class="form-label fw-bold">Nama:</label>
          <p class="form-control-plaintext">{{ users.name }}</p>
        </div>
        <div class="mb-3">
          <label class="form-label fw-bold">Email:</label>
          <p class="form-control-plaintext">{{ users.email }}</p>
        </div>
        <div class="text-center">
          <button class="btn btn-warning" @click="navigateToEditProfile">Edit Profil</button>
          <button class="btn btn-danger ms-2" @click="handleLogout">Logout</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import Cookies from "js-cookie";

definePageMeta({
  middleware: 'auth',
});

const users = ref([])
const error = ref([])
const router = useRouter();

const navigateToEditProfile = () => {
  router.push(`/user/profile/${route.params.id}/edit`);
};


// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];
  
  return token;
};

const fetchUsers = async (id) => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const response = await axios.get(`http://localhost:8080/user/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    console.log('Respons API:', response.data);
    users.value = response.data;
  } catch (err) {
    error.value = err.response?.data?.message || err.message;
    console.error('Error fetching users:', err);
  }
};

// Mengambil ID dari URL
const route = useRoute();
onMounted(() => {
  const userId = route.params.id;
  fetchUsers(userId);
});

const handleLogout = () => {
  try {
    // Hapus token dari cookies
    Cookies.remove('token')

    alert('Logout berhasil!')

    // Redirect ke halaman login
    window.location.href = '/login'
  } catch (error) {
    console.error('Terjadi kesalahan saat logout:', error)
    alert('Logout gagal. Silakan coba lagi.')
  }
}
</script>
  