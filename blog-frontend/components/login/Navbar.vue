<template>
  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg">
    <div class="container">
      <NuxtLink class="navbar-brand" to="/">Kang Tutor</NuxtLink>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-auto">
          <li class="nav-item"><NuxtLink class="nav-link text-white" to="/user/dashboard">Beranda</NuxtLink></li>
          <li class="nav-item" v-if="user_id">
            <NuxtLink :to="`/user/profile/${user_id}`" class="nav-link text-white">Profile</NuxtLink>
          </li>
          <!-- Logika untuk menampilkan input atau tombol login -->
          <li class="d-flex nav-item" v-if="isLoggedIn">
            <input
              class="form-control"
              type="text"
              v-model="query"
              placeholder="Cari artikel atau konten"
              aria-label="default input example"
            />
            <button @click="goToSearchPage" class="btn btn-outline-light ms-2">Search</button>
          </li>
          <li class="nav-item ms-3" v-else>
            <NuxtLink class="btn btn-outline-light" to="/login">Login</NuxtLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { jwtDecode } from 'jwt-decode';
import { useRouter } from 'vue-router';

// Status login user
const isLoggedIn = ref(true);

// Email pengguna dari token
const user_id = ref('');

// State untuk pencarian
const query = ref('');
const categoryId = ref('');
const router = useRouter();

// Navigasi ke halaman search
const goToSearchPage = () => {
  if (!query.value.trim()) {
    alert('Masukkan kata kunci pencarian.');
    return;
  }

  router.push({
    path: '/user/search', // Pastikan rute ini ada di Nuxt
    query: {
      query: query.value.trim(),
      category_id: categoryId.value || '', // Pastikan query string tidak kosong
    },
  });
};

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
    console.log(decoded)
  } catch (error) {
    console.error('Error decoding token:', error);
  }
};

onMounted(() => {
  decodeToken();
});
</script>

<style scoped>
body {
  font-family: Arial, sans-serif;
}
.navbar {
  background-color: #00A885;
}
.navbar-brand {
  font-weight: bold;
  color: white;
}
.nav-link:hover {
  color: #FFD700; /* Warna hover */
  transition: color 0.3s;
}
</style>
