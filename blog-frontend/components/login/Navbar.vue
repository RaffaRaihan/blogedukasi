<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg">
    <!-- Alert Message -->
    <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
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
                    <span class="me-2" style="color: #F0F3FF;">
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
                    <li><NuxtLink class="dropdown-item" to="/user/message/"><i class="bi bi-envelope"></i>  Mail</NuxtLink></li>
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
const alertMessage = ref('');
const alertClass = ref('');

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

    console.log('Respons API:', response.data.data);
    users.value = response.data.data;
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

    confirm('apakah yakin ingin logout')
    alertMessage.value = 'Logout Berhasil!'
    alertClass.value = 'alert alert-success'

    // Redirect ke halaman login
    setTimeout(() => {
      window.location.href = '/login';
    }, 1500)
  } catch (error) {
    console.error('Terjadi kesalahan saat logout:', error);
    alertMessage.value = 'Logout gagal. Silakan coba lagi.'
    alertClass.value = 'alert alert-danger'
  }
};

onMounted(() => {
  decodeToken();
  fetchUsers(user_id.value);
});
</script>

<style scoped>
.navbar {
  background-color: #211951;
  position: sticky;
  top: 0px;
  z-index: 2;
}
.navbar .nav-link {
  color: #F0F3FF;
}
.navbar .nav-link:hover {
  color: #15F5BA;
  transition: 0.3s;
}
.profile-img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
}
.logo{
  color: #F0F3FF;
}
.logo:hover{
  color: #15F5BA;
  transition: 0.3s;
}
.btn{
  color: #211951;
  background-color: #15F5BA;
  border-color: #211951;
}
.btn:hover{
  color: #15F5BA;
  background-color: #211951;
  border-color: #15F5BA;
  transition: 0.3s;
}
.dropdown-menu{
  background: #211951;
}
.dropdown-item {
  color: #F0F3FF;
}
.dropdown-item:hover{
  color: #27005D;
  background-color: #15F5BA;
  transition: 0.3s;
}
.input{
  color: #F9F6E6;
  border-color: #F9F6E6;
}
.hr{
  color: #AED2FF;
}
.alert {
  position: absolute;
  z-index: 0;
  width: 400px;
  margin-left: 600px;
}
</style>
