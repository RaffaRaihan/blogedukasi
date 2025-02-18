<template>
  <div class="container mt-4">
    <NuxtLink to="/user/dashboard" class="text-decoration-none" style="color: #211951;"><i class="bi bi-arrow-left"></i>  Kembali</NuxtLink>
    <div class="card mx-auto mt-3">
      <div class="card-header" style="background-color: #211951; color: #F0F3FF;">
        <h3 class="mb-0">Profil Pengguna</h3>
      </div>
      <div class="card-body">
        <div class="d-flex">
          <img
            :src="`http://localhost:8080/uploads/${users.foto}`"
            alt="Foto Profil"
            class="rounded-1"
            style="width: 500px; height: 500px;"
          />
          <div class="text-start ms-5 mt-5">
            <div class="mb-3 d-flex">
              <h2 class="fw-bold">Nama:</h2>
              <h2 class="ms-3">{{ users.name }}</h2>
            </div>
            <div class="mb-3 d-flex">
              <h2 class="fw-bold">Email:</h2>
              <h2 class="ms-3">{{ users.email }}</h2>
            </div>
            <button class="btn btn-outline-warning" @click="navigateToEditProfile"><i class="bi bi-pencil"></i>  Edit Profil</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";

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

    console.log('Respons API:', response.data.data);
    users.value = response.data.data;
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

</script>
  