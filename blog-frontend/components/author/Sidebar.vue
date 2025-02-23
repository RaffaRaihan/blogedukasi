<template>
    <!-- Sidebar -->
    <div class="sidebar">
        <div class="logo">
            <h3 class="fw-bold">Raffa Mr.</h3>
        </div>
        <nav class="nav flex-column">
            <NuxtLink to="/author/dashboard" class="nav-link">
                <i class="bi bi-house"></i> Dashboard
            </NuxtLink>
            <NuxtLink to="/author/articles" class="nav-link">
                <i class="bi bi-book"></i> Artikel
            </NuxtLink>
            <NuxtLink :to="`/author/profile/${user_id}`" class="nav-link">
              <i class="bi bi-person"></i> Profil
            </NuxtLink>
            <hr>
            <button data-bs-toggle="modal" data-bs-target="#logoutModal" class="btn-2">
              <i class="bi bi-box-arrow-left"></i> Keluar
            </button>
        </nav>
        <div v-if="alertMessage" class="alert mt-2" :class="alertClass" role="alert">{{ alertMessage }}</div>
    </div>

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
import useHandleLogout from '~/composables/api/useHandleLogout';
import { jwtDecode } from 'jwt-decode';
import 'bootstrap-icons/font/bootstrap-icons.css';

const user_id = ref('');

// Fungsi untuk mengambil token dari cookies
const getTokenFromCookies = () => {
  const token = document.cookie
    .split('; ')
    .find(row => row.startsWith('token='))
    ?.split('=')[1];

  return token;
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

const { alertMessage, alertClass, handleLogout } = useHandleLogout();

onMounted(() => {
  decodeToken();
});
</script>

<style scoped>
.sidebar {
  width: 250px;
  height: 100vh;
  background-color: #1D2B53; /* Warna latar sidebar */
  display: flex;
  flex-direction: column;
  padding: 1rem;
  border-right: 1px solid #e0e0e0;
  position: sticky;
  top: 1px;
}
.sidebar .logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
  margin-left: 1rem;
  margin-top: 1rem;
  color: #F9F6E6;
}
.sidebar .nav-item {
  margin: 0.5rem 0;
}
.sidebar .nav-link {
  color: #F9F6E6;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
.sidebar .nav-link:hover {
  background-color: #F9F6E6;
  border-radius: 5px;
  color: #212529;
}
.sidebar .notification-badge {
  background-color: #f03e3e;
  color: #fff;
  font-size: 0.75rem;
  font-weight: bold;
  border-radius: 50%;
  padding: 0.25rem 0.5rem;
}
.sidebar hr{
  color: #F9F6E6;
}
.btn-2 {
  padding: 0.5rem;
  border-radius: 5px;
  color: #F9F6E6;
  border: solid 1px;
  background-color: #1D2B53;
}
.btn-2:hover{
  color: #1D2B53;
  border-color: #1D2B53;
  background-color: #F9F6E6;
  transition: 0.3s;
}
.alert {
  position: absolute;
  z-index: 0;
  width: 400px;
  margin-left: 600px;
}
</style>
