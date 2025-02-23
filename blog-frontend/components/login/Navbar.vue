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
                        <NuxtLink class="nav-link" to="/user/dashboard">Beranda</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/about">Tentang</NuxtLink>
                    </li>
                    <li class="nav-item">
                        <NuxtLink class="nav-link" to="/contact">Kontak</NuxtLink>
                    </li>
                </ul>
                <div class="dropdown ms-2" v-if="isLoggedIn">
                  <a class="d-flex align-items-center text-decoration-none text-white" href="#" role="button" id="profileDropdown" data-bs-toggle="dropdown" aria-expanded="false">
                    <span class="me-2" style="color: #F0F3FF;">
                      <i class="bi bi-chevron-down"></i> {{ users.name }}
                    </span>
                    <img :src="`http://localhost:8080/uploads/${users.foto}`" alt="Profile" class="profile-img" />
                  </a>
                  <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="profileDropdown">
                    <li><NuxtLink class="dropdown-item" :to="`/user/profile/${users.ID}`"><i class="bi bi-person"></i> Profil</NuxtLink></li>
                    <li><NuxtLink class="dropdown-item" to="/user/message/"><i class="bi bi-envelope"></i> Pesan</NuxtLink></li>
                    <li><hr class="hr"/></li>
                    <li><button class="dropdown-item" data-bs-toggle="modal" data-bs-target="#logoutModal"><i class="bi bi-box-arrow-left"></i> Logout</button></li>
                  </ul>
                </div>
                <li class="nav-link" v-else><NuxtLink class="btn" to="/login">Login</NuxtLink></li>
            </div>
        </div>
        <div v-if="alertMessage" class="alert" :class="alertClass" role="alert">{{ alertMessage }}</div>
    </nav>
  
  <!-- Logout Modal -->
  <div class="modal fade" id="logoutModal" tabindex="-1" aria-labelledby="logoutModalLabel" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="logoutModalLabel">Konfirmasi Logout</h5>
        </div>
        <div class="modal-body">
          Apakah Anda yakin ingin logout?
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Batal</button>
          <button type="button" class="btn btn-danger" @click="handleLogout" data-bs-dismiss="modal">Logout</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { jwtDecode } from 'jwt-decode';
import useHandleLogout from '~/composables/api/useHandleLogout';
import axios from 'axios';

const isLoggedIn = ref(true);
const users = ref([]);
const user_id = ref('');

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
    users.value = response.data.data;
  } catch (err) {
    console.error('Error fetching users:', err);
  }
};

const decodeToken = () => {
  try {
    const token = getTokenFromCookies();
    if (!token) throw new Error('Token tidak ditemukan. Harap login terlebih dahulu.');

    const decoded = jwtDecode(token);
    user_id.value = decoded.user_id;
  } catch (error) {
    console.error('Error decoding token:', error);
  }
};

const { alertMessage, alertClass, handleLogout } = useHandleLogout();

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
  z-index: 10;
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
  z-index: 10;
  width: 400px;
  margin-left: 600px;
}
</style>
