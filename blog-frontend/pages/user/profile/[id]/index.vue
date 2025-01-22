<template>
  <div class="container mt-5">
    <NuxtLink to="/user/dashboard" class="text-decoration-none" style="color: #1D2B53;"><i class="bi bi-arrow-left"></i>  Back</NuxtLink>
    <div class="card mx-auto" style="max-width: 600px;">
      <div class="card-header" style="background-color: #1D2B53; color: #FF004D;">
        <h3 class="mb-0">Profil Pengguna</h3>
      </div>
      <div class="card-body">
        <div class="text-center">
          <img
            :src="`http://localhost:8080/uploads/${users.foto}`"
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
          <button class="btn btn-outline-warning" @click="navigateToEditProfile"><i class="bi bi-pencil"></i>  Edit Profil</button>
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
  