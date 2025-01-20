<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg">
        <div class="container">
            <NuxtLink class="logo navbar-brand" to="/user/dashboard">Raffa Mr.</NuxtLink>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ms-auto">
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/user/dashboard">Home</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/about">About</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/contact">Contact</NuxtLink>
                    </li>
                </ul>
                <div class="dropdown ms-2" v-if="isLoggedIn">
                  <a
                    class="d-flex align-items-center text-decoration-none text-white"
                    href="#"
                    role="button"
                    id="profileDropdown"
                    data-bs-toggle="dropdown"
                    aria-expanded="false"
                  >
                    <span class="me-2" style="color: #F9F6E6;">
                      <i class="bi bi-chevron-down"></i> {{ users.name }}
                    </span>
                    <img
                      :src="`http://localhost:8080/uploads/${users.foto}`"
                      alt="Profile"
                      class="profile-img"
                    />
                  </a>
                  <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="profileDropdown">
                    <li><NuxtLink class="dropdown-item" :to="`/user/profile/${user_id}`"><i class="bi bi-person"></i>  Profile</NuxtLink></li>
                    <li><a class="dropdown-item" href="#"><i class="bi bi-gear"></i>  Settings</a></li>
                    <li><hr class="hr"/></li>
                    <li><button class="dropdown-item" @click="handleLogout"><i class="bi bi-box-arrow-left"></i>  Logout</button></li>
                  </ul>
                </div>
                <li class="nav-link" v-else><NuxtLink class="btn" to="/login">Login</NuxtLink></li>
            </div>
        </div>
    </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { jwtDecode } from 'jwt-decode';
import Cookies from "js-cookie";
import axios from 'axios';

// Status login user
const isLoggedIn = ref(true);

// Email pengguna dari token
const users = ref([]);
const user_id = ref('');

const errorMessage = ref(''); // Inisialisasi errorMessage dengan ref



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
    errorMessage.value = ''; // Reset pesan error jika berhasil
  } catch (err) {
    console.error('Error fetching users:', err);
    errorMessage.value = 'Gagal memuat data pengguna. Silakan coba lagi nanti.'; // Tampilkan pesan error
  }
};

// Fungsi untuk mendekode token JWT
const decodeToken = () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const decoded = jwtDecode(token); // Dekode JWT
    user_id.value = decoded.user_id; // Ambil email dari payload
    console.log(decoded);
  } catch (error) {
    console.error('Error decoding token:', error);
    errorMessage.value = 'Terjadi kesalahan saat memproses token. Harap login kembali.'; // Tampilkan pesan error
  }
};

const handleLogout = () => {
  try {
    // Hapus token dari cookies
    Cookies.remove('token');

    alert('Logout berhasil!');

    // Redirect ke halaman login
    window.location.href = '/login';
  } catch (error) {
    console.error('Terjadi kesalahan saat logout:', error);
    alert('Logout gagal. Silakan coba lagi.');
  }
};

onMounted(() => {
  decodeToken();
  fetchUsers(user_id.value);
});
</script>

<style scoped>
.navbar {
  background-color: #1D2B53; /* Warna gelap */
  position: sticky;
  top: 0px;
  z-index: 2;
}
.navbar .nav-link {
  color: #F9F6E6; /* Warna teks */
}
.navbar .nav-link:hover {
  color: #FF004D; /* Warna teks saat hover */
}
.profile-img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.logo{
  color: #F9F6E6;
}
.logo:hover{
  color: #FF004D;
}
.btn{
  color: #1D2B53;
  background-color: #F9F6E6;
  border-color: #1D2B53;
}
.btn:hover{
  color: #F9F6E6;
  background-color: #1D2B53;
  border-color: #F9F6E6;
}
.dropdown-menu{
  background: #1D2B53;
}
.dropdown-item {
  color: #F9F6E6;
}
.dropdown-item:hover{
  color: #1D2B53;
  background-color: #F9F6E6;
  transition: 0.3s;
}
.input{
  color: #F9F6E6;
  border-color: #F9F6E6;
}
.hr{
  color: #F9F6E6;
}
</style>
